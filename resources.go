package beeos

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// WaitOptions controls local waiting, never the lifetime of the remote task.
type WaitOptions struct {
	Timeout           time.Duration
	PollInterval      time.Duration
	AllowUnsuccessful bool
}

func validateIDs(ids ...string) error {
	for _, id := range ids {
		if strings.TrimSpace(id) == "" {
			return fmt.Errorf("resource ID is required")
		}
	}
	return nil
}

func requestContext(ctx context.Context) (context.Context, context.CancelFunc) {
	if _, ok := ctx.Deadline(); ok {
		return context.WithCancel(ctx)
	}
	return context.WithTimeout(ctx, 30*time.Second)
}

func waitOptions(options []WaitOptions, defaultTimeout time.Duration) (WaitOptions, error) {
	if len(options) > 1 {
		return WaitOptions{}, fmt.Errorf("expected at most one WaitOptions value")
	}
	var opts WaitOptions
	if len(options) == 1 {
		opts = options[0]
	}
	if opts.Timeout < 0 || opts.PollInterval < 0 {
		return opts, fmt.Errorf("wait durations cannot be negative")
	}
	if opts.Timeout == 0 {
		opts.Timeout = defaultTimeout
	}
	if opts.PollInterval == 0 {
		opts.PollInterval = time.Second
	}
	return opts, nil
}

func resourceCall[T any](ctx context.Context, id string, execute func(context.Context) (*T, *http.Response, error)) (*T, error) {
	if err := validateIDs(id); err != nil {
		return nil, err
	}
	ctx, cancel := requestContext(ctx)
	defer cancel()
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	response, _, err := execute(ctx)
	return response, err
}

type InstancesResource struct{ client *Client }
type InstanceResource struct {
	client   *Client
	id       string
	Data     *InstanceDataDTO
	Mobile   *MobileResource
	Computer *ComputerResource
}

// Instance binds a handle locally. It does not create or fetch a resource.
func (c *Client) Instance(id string) *InstanceResource {
	return &InstanceResource{client: c, id: id, Mobile: &MobileResource{client: c, id: id}, Computer: &ComputerResource{client: c, id: id}}
}
func (i *InstanceResource) ID() string { return i.id }
func (r *InstancesResource) Get(ctx context.Context, id string) (*InstanceResource, error) {
	return r.client.Instance(id).Refresh(ctx)
}
func (r *InstancesResource) Create(ctx context.Context, request DeployInstanceRequest) (*InstanceResource, error) {
	response, err := resourceCall(ctx, "create", func(ctx context.Context) (*InstanceResponse, *http.Response, error) {
		return r.client.API.InstancesAPI.DeployInstance(r.client.authContext(ctx)).DeployInstanceRequest(request).Execute()
	})
	if err != nil {
		return nil, err
	}
	if response == nil || !response.Success || strings.TrimSpace(response.Data.Id) == "" {
		return nil, fmt.Errorf("invalid instance creation response")
	}
	instance := r.client.Instance(response.Data.Id)
	instance.Data = &response.Data
	return instance, nil
}

type InstanceListOptions struct {
	Page, PageSize                                        int32
	Status, ProviderID, AgentFramework, ClusterID, Search string
}

func (r *InstancesResource) List(ctx context.Context, opts InstanceListOptions) (*InstancePageResponse, error) {
	return resourceCall(ctx, "list", func(ctx context.Context) (*InstancePageResponse, *http.Response, error) {
		request := r.client.API.InstancesAPI.ListInstances(r.client.authContext(ctx))
		if opts.Page > 0 {
			request = request.Page(opts.Page)
		}
		if opts.PageSize > 0 {
			request = request.PageSize(opts.PageSize)
		}
		if opts.Status != "" {
			request = request.Status(opts.Status)
		}
		if opts.ProviderID != "" {
			request = request.ProviderId(opts.ProviderID)
		}
		if opts.AgentFramework != "" {
			request = request.AgentFramework(opts.AgentFramework)
		}
		if opts.ClusterID != "" {
			request = request.ClusterId(opts.ClusterID)
		}
		if opts.Search != "" {
			request = request.Search(opts.Search)
		}
		return request.Execute()
	})
}
func (i *InstanceResource) Refresh(ctx context.Context) (*InstanceResource, error) {
	response, err := resourceCall(ctx, i.id, func(ctx context.Context) (*InstanceResponse, *http.Response, error) {
		return i.client.API.InstancesAPI.GetInstance(i.client.authContext(ctx), i.id).Execute()
	})
	if err != nil {
		return nil, err
	}
	if response == nil || !response.Success || response.Data.Id != i.id {
		return nil, fmt.Errorf("instance response identity does not match the requested instance")
	}
	i.Data = &response.Data
	return i, nil
}
func (i *InstanceResource) WaitReady(ctx context.Context, options ...WaitOptions) (*InstanceResource, error) {
	opts, err := waitOptions(options, 60*time.Second)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(ctx, opts.Timeout)
	defer cancel()
	for {
		if _, err := i.Refresh(ctx); err != nil {
			return nil, err
		}
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		if i.Data.Status == "running" && i.Data.Connectivity == "online" {
			return i, nil
		}
		switch i.Data.Status {
		case "failed", "error", "terminated", "destroyed":
			return nil, fmt.Errorf("instance %s cannot become ready: %s", i.id, i.Data.Status)
		}
		if err := waitContext(ctx, opts.PollInterval); err != nil {
			return nil, err
		}
	}
}
func (i *InstanceResource) Agents(ctx context.Context) (*AgentListResponse, error) {
	return resourceCall(ctx, i.id, func(ctx context.Context) (*AgentListResponse, *http.Response, error) {
		return i.client.API.AgentsAPI.ListAgents(i.client.authContext(ctx)).InstanceId(i.id).Execute()
	})
}
func (i *InstanceResource) Start(ctx context.Context) (*InstanceResponse, error) {
	return resourceCall(ctx, i.id, func(ctx context.Context) (*InstanceResponse, *http.Response, error) {
		return i.client.API.InstancesAPI.StartInstance(i.client.authContext(ctx), i.id).Execute()
	})
}
func (i *InstanceResource) Stop(ctx context.Context) (*InstanceResponse, error) {
	return resourceCall(ctx, i.id, func(ctx context.Context) (*InstanceResponse, *http.Response, error) {
		return i.client.API.InstancesAPI.StopInstance(i.client.authContext(ctx), i.id).Execute()
	})
}
func (i *InstanceResource) Restart(ctx context.Context) (*InstanceResponse, error) {
	return resourceCall(ctx, i.id, func(ctx context.Context) (*InstanceResponse, *http.Response, error) {
		return i.client.API.InstancesAPI.RestartInstance(i.client.authContext(ctx), i.id).Execute()
	})
}
func (i *InstanceResource) Destroy(ctx context.Context) (*StringResponse, error) {
	return resourceCall(ctx, i.id, func(ctx context.Context) (*StringResponse, *http.Response, error) {
		return i.client.API.InstancesAPI.DestroyInstance(i.client.authContext(ctx), i.id).Execute()
	})
}

type TaskExecutionError struct{ Snapshot *TaskResponse }

func (e *TaskExecutionError) Error() string {
	return fmt.Sprintf("task %s ended with %s", e.Snapshot.Data.TaskId, e.Snapshot.Data.Status)
}

type TaskResource struct {
	client      *Client
	agentID, id string
}

func (c *Client) Task(agentID, taskID string) *TaskResource {
	return &TaskResource{client: c, agentID: agentID, id: taskID}
}
func (t *TaskResource) ID() string      { return t.id }
func (t *TaskResource) AgentID() string { return t.agentID }
func (t *TaskResource) Get(ctx context.Context) (*TaskResponse, error) {
	return t.client.GetTask(ctx, t.agentID, t.id)
}
func (t *TaskResource) Cancel(ctx context.Context) (*TaskResponse, error) {
	return t.client.CancelTask(ctx, t.agentID, t.id)
}
func (t *TaskResource) Events(ctx context.Context, since int64) (<-chan TaskEvent, <-chan error) {
	return t.client.TaskEvents(ctx, t.agentID, t.id, since)
}
func (t *TaskResource) Wait(ctx context.Context, options ...WaitOptions) (*TaskResponse, error) {
	opts, err := waitOptions(options, 120*time.Second)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(ctx, opts.Timeout)
	defer cancel()
	for {
		snapshot, err := t.Get(ctx)
		if err != nil {
			return nil, err
		}
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		if snapshot == nil || snapshot.Data.TaskId != t.id || snapshot.Data.AgentId != t.agentID {
			return nil, fmt.Errorf("task response identity does not match the requested task")
		}
		if snapshot.Data.GetTruncated() {
			return nil, fmt.Errorf("task snapshot is truncated; resume events for the authoritative tail")
		}
		if isTerminalTaskStatus(snapshot.Data.Status) {
			if !opts.AllowUnsuccessful && snapshot.Data.Status != TASKSTATUS_COMPLETED {
				return snapshot, &TaskExecutionError{Snapshot: snapshot}
			}
			return snapshot, nil
		}
		if err := waitContext(ctx, opts.PollInterval); err != nil {
			return nil, err
		}
	}
}

type MessageListOptions struct {
	Since         int64
	Limit         int32
	IncludeDeltas bool
}

func (t *TaskResource) Messages(ctx context.Context, opts MessageListOptions) (*ListMessagesResponse, error) {
	if err := validateIDs(t.agentID, t.id); err != nil {
		return nil, err
	}
	if opts.Since < 0 || opts.Limit < 0 {
		return nil, fmt.Errorf("message pagination cannot be negative")
	}
	return resourceCall(ctx, t.id, func(ctx context.Context) (*ListMessagesResponse, *http.Response, error) {
		request := t.client.API.TasksAPI.ListTaskMessages(t.client.authContext(ctx), t.agentID, t.id).Since(opts.Since).IncludeDeltas(opts.IncludeDeltas)
		if opts.Limit > 0 {
			request = request.Limit(opts.Limit)
		}
		return request.Execute()
	})
}

type MobileResource struct {
	client *Client
	id     string
}

func (m *MobileResource) Info(ctx context.Context) (*GetComputerInfo200Response, error) {
	return resourceCall(ctx, m.id, func(ctx context.Context) (*GetComputerInfo200Response, *http.Response, error) {
		return m.client.API.MobileAPI.GetMobileInfo(m.client.authContext(ctx), m.id).Execute()
	})
}
func (m *MobileResource) WaitReady(ctx context.Context, options ...WaitOptions) (*GetComputerInfo200Response, error) {
	opts, err := waitOptions(options, 60*time.Second)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(ctx, opts.Timeout)
	defer cancel()
	for {
		response, err := m.Info(ctx)
		if err != nil {
			return nil, err
		}
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		if response != nil && response.Data != nil && response.Data.Online {
			return response, nil
		}
		if err := waitContext(ctx, opts.PollInterval); err != nil {
			return nil, err
		}
	}
}
func (m *MobileResource) Screenshot(ctx context.Context) (*ComputerScreenshot200Response, error) {
	return resourceCall(ctx, m.id, func(ctx context.Context) (*ComputerScreenshot200Response, *http.Response, error) {
		return m.client.API.MobileAPI.MobileScreenshot(m.client.authContext(ctx), m.id).Execute()
	})
}

type UITreeOptions struct{ Query, Format string }

func (m *MobileResource) UITree(ctx context.Context, opts UITreeOptions) (*MobileGetUiTree200Response, error) {
	return resourceCall(ctx, m.id, func(ctx context.Context) (*MobileGetUiTree200Response, *http.Response, error) {
		request := m.client.API.MobileAPI.MobileGetUiTree(m.client.authContext(ctx), m.id)
		if opts.Query != "" {
			request = request.Query(opts.Query)
		}
		if opts.Format != "" {
			request = request.Format(opts.Format)
		}
		return request.Execute()
	})
}

type AppListOptions struct {
	IncludeSystem        bool
	IncludeNonLaunchable bool
}

func (m *MobileResource) ListApps(ctx context.Context, opts AppListOptions) (*MobileListApps200Response, error) {
	return resourceCall(ctx, m.id, func(ctx context.Context) (*MobileListApps200Response, *http.Response, error) {
		return m.client.API.MobileAPI.MobileListApps(m.client.authContext(ctx), m.id).IncludeSystem(opts.IncludeSystem).LaunchableOnly(!opts.IncludeNonLaunchable).Execute()
	})
}
func (m *MobileResource) Tap(ctx context.Context, request TapRequest) (*InlineObject, error) {
	return resourceCall(ctx, m.id, func(ctx context.Context) (*InlineObject, *http.Response, error) {
		return m.client.API.MobileAPI.MobileTap(m.client.authContext(ctx), m.id).TapRequest(request).Execute()
	})
}
func (m *MobileResource) DoubleTap(ctx context.Context, request TapRequest) (*InlineObject, error) {
	return resourceCall(ctx, m.id, func(ctx context.Context) (*InlineObject, *http.Response, error) {
		return m.client.API.MobileAPI.MobileDoubleTap(m.client.authContext(ctx), m.id).TapRequest(request).Execute()
	})
}
func (m *MobileResource) LongPress(ctx context.Context, request LongPressRequest) (*InlineObject, error) {
	return resourceCall(ctx, m.id, func(ctx context.Context) (*InlineObject, *http.Response, error) {
		return m.client.API.MobileAPI.MobileLongPress(m.client.authContext(ctx), m.id).LongPressRequest(request).Execute()
	})
}
func (m *MobileResource) Drag(ctx context.Context, request DragRequest) (*InlineObject, error) {
	return resourceCall(ctx, m.id, func(ctx context.Context) (*InlineObject, *http.Response, error) {
		return m.client.API.MobileAPI.MobileDrag(m.client.authContext(ctx), m.id).DragRequest(request).Execute()
	})
}
func (m *MobileResource) Swipe(ctx context.Context, request SwipeRequest) (*InlineObject, error) {
	return resourceCall(ctx, m.id, func(ctx context.Context) (*InlineObject, *http.Response, error) {
		return m.client.API.MobileAPI.MobileSwipe(m.client.authContext(ctx), m.id).SwipeRequest(request).Execute()
	})
}
func (m *MobileResource) Scroll(ctx context.Context, request ScrollRequest) (*InlineObject, error) {
	return resourceCall(ctx, m.id, func(ctx context.Context) (*InlineObject, *http.Response, error) {
		return m.client.API.MobileAPI.MobileScroll(m.client.authContext(ctx), m.id).ScrollRequest(request).Execute()
	})
}
func (m *MobileResource) Type(ctx context.Context, text string) (*InlineObject, error) {
	return resourceCall(ctx, m.id, func(ctx context.Context) (*InlineObject, *http.Response, error) {
		return m.client.API.MobileAPI.MobileType(m.client.authContext(ctx), m.id).TypeRequest(TypeRequest{Text: text}).Execute()
	})
}
func (m *MobileResource) Key(ctx context.Context, key string) (*InlineObject, error) {
	return resourceCall(ctx, m.id, func(ctx context.Context) (*InlineObject, *http.Response, error) {
		return m.client.API.MobileAPI.MobileKey(m.client.authContext(ctx), m.id).KeyRequest(KeyRequest{Keys: []string{key}}).Execute()
	})
}
func (m *MobileResource) PressButton(ctx context.Context, button string) (*InlineObject, error) {
	return resourceCall(ctx, m.id, func(ctx context.Context) (*InlineObject, *http.Response, error) {
		return m.client.API.MobileAPI.MobilePressButton(m.client.authContext(ctx), m.id).PressButtonRequest(PressButtonRequest{Button: button}).Execute()
	})
}
func (m *MobileResource) OpenApp(ctx context.Context, app string) (*InlineObject, error) {
	return resourceCall(ctx, m.id, func(ctx context.Context) (*InlineObject, *http.Response, error) {
		return m.client.API.MobileAPI.MobileOpenApp(m.client.authContext(ctx), m.id).OpenAppRequest(OpenAppRequest{App: app}).Execute()
	})
}

type ComputerResource struct {
	client *Client
	id     string
}

func (m *ComputerResource) Info(ctx context.Context) (*GetComputerInfo200Response, error) {
	return resourceCall(ctx, m.id, func(ctx context.Context) (*GetComputerInfo200Response, *http.Response, error) {
		return m.client.API.ComputerAPI.GetComputerInfo(m.client.authContext(ctx), m.id).Execute()
	})
}
func (m *ComputerResource) Screenshot(ctx context.Context) (*ComputerScreenshot200Response, error) {
	return resourceCall(ctx, m.id, func(ctx context.Context) (*ComputerScreenshot200Response, *http.Response, error) {
		return m.client.API.ComputerAPI.ComputerScreenshot(m.client.authContext(ctx), m.id).Execute()
	})
}
func (m *ComputerResource) Click(ctx context.Context, request ClickRequest) (*InlineObject, error) {
	return resourceCall(ctx, m.id, func(ctx context.Context) (*InlineObject, *http.Response, error) {
		return m.client.API.ComputerAPI.ComputerClick(m.client.authContext(ctx), m.id).ClickRequest(request).Execute()
	})
}
func (m *ComputerResource) Move(ctx context.Context, request MoveRequest) (*InlineObject, error) {
	return resourceCall(ctx, m.id, func(ctx context.Context) (*InlineObject, *http.Response, error) {
		return m.client.API.ComputerAPI.ComputerMove(m.client.authContext(ctx), m.id).MoveRequest(request).Execute()
	})
}
func (m *ComputerResource) Scroll(ctx context.Context, request ScrollRequest) (*InlineObject, error) {
	return resourceCall(ctx, m.id, func(ctx context.Context) (*InlineObject, *http.Response, error) {
		return m.client.API.ComputerAPI.ComputerScroll(m.client.authContext(ctx), m.id).ScrollRequest(request).Execute()
	})
}
func (m *ComputerResource) Type(ctx context.Context, text string) (*InlineObject, error) {
	return resourceCall(ctx, m.id, func(ctx context.Context) (*InlineObject, *http.Response, error) {
		return m.client.API.ComputerAPI.ComputerType(m.client.authContext(ctx), m.id).TypeRequest(TypeRequest{Text: text}).Execute()
	})
}
func (m *ComputerResource) Key(ctx context.Context, keys ...string) (*InlineObject, error) {
	return resourceCall(ctx, m.id, func(ctx context.Context) (*InlineObject, *http.Response, error) {
		return m.client.API.ComputerAPI.ComputerKey(m.client.authContext(ctx), m.id).KeyRequest(KeyRequest{Keys: keys}).Execute()
	})
}
