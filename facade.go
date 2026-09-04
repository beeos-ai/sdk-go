package beeos

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// Client is the stable convenience facade. API exposes the full generated client.
type Client struct {
	API     *APIClient
	apiKey  string
	baseURL string
	http    *http.Client
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
	return &Client{
		API:     NewAPIClient(cfg),
		apiKey:  resolved.APIKey,
		baseURL: strings.TrimRight(resolved.BaseURL, "/"),
		http:    resolved.HTTPClient,
	}, nil
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
	for {
		response, _, err := c.API.MobileAPI.GetMobileInfo(c.authContext(ctx), c.InstanceID).Execute()
		if err != nil {
			return nil, err
		}
		if response.Data != nil && response.Data.Online {
			return response, nil
		}
		if err := waitContext(ctx, c.pollInterval); err != nil {
			return nil, err
		}
	}
}

// Run creates a durable task and waits for its terminal snapshot.
func (c *MobileClient) Run(ctx context.Context, request CreateTaskRequest) (*TaskResponse, error) {
	created, err := c.CreateTask(ctx, c.AgentID, request)
	if err != nil {
		return nil, err
	}
	for {
		snapshot, err := c.GetTask(ctx, c.AgentID, created.Data.TaskId)
		if err != nil {
			return nil, err
		}
		if isTerminalTaskStatus(snapshot.Data.Status) {
			return snapshot, nil
		}
		if err := waitContext(ctx, c.pollInterval); err != nil {
			return nil, err
		}
	}
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
	response, _, err := c.API.AgentsAPI.ListAgents(c.authContext(ctx)).Execute()
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

func (c *Client) CreateTask(ctx context.Context, agentID string, request CreateTaskRequest) (*TaskCreatedResponse, error) {
	response, _, err := c.API.TasksAPI.CreateTask(c.authContext(ctx), agentID).CreateTaskRequest(request).Execute()
	return response, err
}

func (c *Client) GetTask(ctx context.Context, agentID, taskID string) (*TaskResponse, error) {
	response, _, err := c.API.TasksAPI.GetTask(c.authContext(ctx), agentID, taskID).Execute()
	return response, err
}

func (c *Client) CancelTask(ctx context.Context, agentID, taskID string) (*TaskResponse, error) {
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
			_, _ = io.Copy(io.Discard, response.Body)
			errors <- fmt.Errorf("task event stream failed: HTTP %d", response.StatusCode)
			return
		}
		if err := scanTaskEvents(response.Body, events); err != nil {
			errors <- err
		}
	}()
	return events, errors
}

func scanTaskEvents(reader io.Reader, events chan<- TaskEvent) error {
	scanner := bufio.NewScanner(reader)
	var event TaskEvent
	var data []string
	flush := func() {
		if len(data) == 0 {
			return
		}
		if event.Event == "" {
			event.Event = "message"
		}
		event.Data = json.RawMessage(strings.Join(data, "\n"))
		events <- event
		event = TaskEvent{}
		data = nil
	}
	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), "\r")
		if line == "" {
			flush()
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
			event.ID = value
		case "data":
			data = append(data, value)
		}
	}
	flush()
	return scanner.Err()
}
