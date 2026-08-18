package beeos

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
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

func NewClient(options ClientOptions) (*Client, error) {
	if options.APIKey == "" {
		return nil, fmt.Errorf("api key is required")
	}
	if options.BaseURL == "" {
		options.BaseURL = "https://openapi.beeos.ai"
	}
	if options.HTTPClient == nil {
		options.HTTPClient = http.DefaultClient
	}
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: strings.TrimRight(options.BaseURL, "/")}}
	cfg.HTTPClient = options.HTTPClient
	return &Client{
		API:     NewAPIClient(cfg),
		apiKey:  options.APIKey,
		baseURL: strings.TrimRight(options.BaseURL, "/"),
		http:    options.HTTPClient,
	}, nil
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
