package beeos

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// Client is the stable convenience facade. API exposes the full generated client.
type Client struct {
	API       *APIClient
	Instances *InstancesResource
	Files     *FilesResource
	apiKey    string
	baseURL   string
	http      *http.Client
}

type ClientOptions struct {
	APIKey     string
	BaseURL    string
	HTTPClient *http.Client
}

type TaskEvent struct {
	Event string
	ID    string
	Data  json.RawMessage
}

type MobileClientOptions struct {
	ClientOptions
	AgentID      string
	InstanceID   string
	PollInterval time.Duration
}

// MobileClient is the task-first facade shared by Device Agent, BeeRunner,
// and Redroid. Direct MobileAPI methods remain available through Client.API.
type MobileClient struct {
	*Client
	AgentID      string
	InstanceID   string
	pollInterval time.Duration
}

func NewClient(options ...ClientOptions) (*Client, error) {
	if len(options) > 1 {
		return nil, fmt.Errorf("expected at most one ClientOptions value")
	}
	resolved := ClientOptions{}
	if len(options) == 1 {
		resolved = options[0]
	}
	resolved.APIKey = strings.TrimSpace(resolved.APIKey)
	if resolved.APIKey == "" {
		resolved.APIKey = strings.TrimSpace(os.Getenv("BEEOS_API_KEY"))
	}
	if resolved.APIKey == "" {
		return nil, fmt.Errorf("api key is required; pass it explicitly or set BEEOS_API_KEY")
	}
	resolved.BaseURL = strings.TrimSpace(resolved.BaseURL)
	if resolved.BaseURL == "" {
		resolved.BaseURL = strings.TrimSpace(os.Getenv("BEEOS_API_URL"))
	}
	if resolved.BaseURL == "" {
		resolved.BaseURL = "https://openapi.beeos.ai"
	}
	if resolved.HTTPClient == nil {
		resolved.HTTPClient = http.DefaultClient
	}
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: strings.TrimRight(resolved.BaseURL, "/")}}
	cfg.HTTPClient = resolved.HTTPClient
	client := &Client{
		API:     NewAPIClient(cfg),
		apiKey:  resolved.APIKey,
		baseURL: strings.TrimRight(resolved.BaseURL, "/"),
		http:    resolved.HTTPClient,
	}
	client.Instances = &InstancesResource{client: client}
	client.Files = &FilesResource{client: client}
	return client, nil
}

func NewMobileClient(options MobileClientOptions) (*MobileClient, error) {
	client, err := NewClient(options.ClientOptions)
	if err != nil {
		return nil, err
	}
	if options.AgentID == "" {
		return nil, fmt.Errorf("agent id is required")
	}
	if options.InstanceID == "" {
		return nil, fmt.Errorf("instance id is required")
	}
	if options.PollInterval <= 0 {
		options.PollInterval = time.Second
	}
	return &MobileClient{
		Client: client, AgentID: options.AgentID, InstanceID: options.InstanceID,
		pollInterval: options.PollInterval,
	}, nil
}

// WaitReady polls live mobile control info until the runtime reports online.
func (c *MobileClient) WaitReady(ctx context.Context) (*GetComputerInfo200Response, error) {
	return c.Instance(c.InstanceID).Mobile.WaitReady(ctx, WaitOptions{PollInterval: c.pollInterval})
}

// Run creates a durable task and waits for its terminal snapshot.
func (c *MobileClient) Run(ctx context.Context, request CreateTaskRequest) (*TaskResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 120*time.Second)
	defer cancel()
	created, err := c.CreateTask(ctx, c.AgentID, request)
	if err != nil {
		return nil, err
	}
	return c.Task(c.AgentID, created.Data.TaskId).Wait(ctx, WaitOptions{
		PollInterval: c.pollInterval, AllowUnsuccessful: true,
	})
}

func (c *MobileClient) Watch(ctx context.Context, taskID string, since int64) (<-chan TaskEvent, <-chan error) {
	return c.TaskEvents(ctx, c.AgentID, taskID, since)
}

func (c *MobileClient) Cancel(ctx context.Context, taskID string) (*TaskResponse, error) {
	return c.CancelTask(ctx, c.AgentID, taskID)
}

func isTerminalTaskStatus(status TaskStatus) bool {
	switch status {
	case TASKSTATUS_COMPLETED, TASKSTATUS_FAILED, TASKSTATUS_CANCELED, TASKSTATUS_TIMEOUT, TASKSTATUS_REJECTED:
		return true
	default:
		return string(status) == "cancelled"
	}
}

func waitContext(ctx context.Context, interval time.Duration) error {
	timer := time.NewTimer(interval)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

func (c *Client) authContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, ContextAccessToken, c.apiKey)
}

func (c *Client) ListAgents(ctx context.Context) ([]AgentDTO, error) {
	ctx, cancel := requestContext(ctx)
	defer cancel()
	response, _, err := c.API.AgentsAPI.ListAgents(c.authContext(ctx)).Execute()
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

func (c *Client) CreateTask(ctx context.Context, agentID string, request CreateTaskRequest) (*TaskCreatedResponse, error) {
	if err := validateIDs(agentID); err != nil {
		return nil, err
	}
	ctx, cancel := requestContext(ctx)
	defer cancel()
	response, _, err := c.API.TasksAPI.CreateTask(c.authContext(ctx), agentID).CreateTaskRequest(request).Execute()
	return response, err
}

func (c *Client) GetTask(ctx context.Context, agentID, taskID string) (*TaskResponse, error) {
	if err := validateIDs(agentID, taskID); err != nil {
		return nil, err
	}
	ctx, cancel := requestContext(ctx)
	defer cancel()
	response, _, err := c.API.TasksAPI.GetTask(c.authContext(ctx), agentID, taskID).Execute()
	return response, err
}

func (c *Client) CancelTask(ctx context.Context, agentID, taskID string) (*TaskResponse, error) {
	if err := validateIDs(agentID, taskID); err != nil {
		return nil, err
	}
	ctx, cancel := requestContext(ctx)
	defer cancel()
	response, _, err := c.API.TasksAPI.CancelTask(c.authContext(ctx), agentID, taskID).Execute()
	return response, err
}

// TaskEvents streams parsed SSE events until the server closes the response.
func (c *Client) TaskEvents(ctx context.Context, agentID, taskID string, since int64) (<-chan TaskEvent, <-chan error) {
	events := make(chan TaskEvent)
	errors := make(chan error, 1)
	go func() {
		defer close(events)
		defer close(errors)
		if err := validateIDs(agentID, taskID); err != nil {
			errors <- err
			return
		}
		if since < 0 {
			errors <- fmt.Errorf("since must be nonnegative")
			return
		}
		u := fmt.Sprintf("%s/api/v1/agents/%s/tasks/%s/events", c.baseURL, url.PathEscape(agentID), url.PathEscape(taskID))
		if since > 0 {
			u += fmt.Sprintf("?since=%d", since)
		}
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
		if err != nil {
			errors <- err
			return
		}
		req.Header.Set("Accept", "text/event-stream")
		req.Header.Set("Authorization", "Bearer "+c.apiKey)
		response, err := c.http.Do(req)
		if err != nil {
			errors <- err
			return
		}
		defer response.Body.Close()
		if response.StatusCode < 200 || response.StatusCode >= 300 {
			_, _ = io.CopyN(io.Discard, response.Body, 4096)
			errors <- fmt.Errorf("task event stream failed: HTTP %d", response.StatusCode)
			return
		}
		contentType, _, _ := mime.ParseMediaType(response.Header.Get("Content-Type"))
		if contentType != "text/event-stream" {
			errors <- fmt.Errorf("task event stream did not return text/event-stream")
			return
		}
		if err := scanTaskEventsContext(ctx, response.Body, events); err != nil {
			errors <- err
		}
	}()
	return events, errors
}

func scanTaskEvents(reader io.Reader, events chan<- TaskEvent) error {
	return scanTaskEventsContext(context.Background(), reader, events)
}

func scanTaskEventsContext(ctx context.Context, reader io.Reader, events chan<- TaskEvent) error {
	scanner := bufio.NewScanner(reader)
	scanner.Buffer(make([]byte, 4096), 4*1024*1024)
	var event TaskEvent
	var data []string
	var frameBytes int
	flush := func() error {
		if len(data) == 0 {
			event.Event = ""
			frameBytes = 0
			return nil
		}
		if event.Event == "" {
			event.Event = "message"
		}
		event.Data = json.RawMessage(strings.Join(data, "\n"))
		select {
		case events <- event:
		case <-ctx.Done():
			return ctx.Err()
		}
		ended := event.Event == "end"
		event = TaskEvent{ID: event.ID}
		data = nil
		frameBytes = 0
		if ended {
			return io.EOF
		}
		return nil
	}
	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), "\r")
		frameBytes += len(line) + 1
		if frameBytes > 4*1024*1024 {
			return fmt.Errorf("task event exceeds 4 MiB")
		}
		if line == "" {
			if err := flush(); err != nil {
				if err == io.EOF {
					return nil
				}
				return err
			}
			continue
		}
		if strings.HasPrefix(line, ":") {
			continue
		}
		field, value, _ := strings.Cut(line, ":")
		value = strings.TrimPrefix(value, " ")
		switch field {
		case "event":
			event.Event = value
		case "id":
			if !strings.ContainsRune(value, '\x00') {
				event.ID = value
			}
		case "data":
			data = append(data, value)
		}
	}
	// An incomplete EOF frame must not be mistaken for a complete terminal event.
	return scanner.Err()
}
