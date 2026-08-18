# \TasksAPI

All URIs are relative to *https://openapi.beeos.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelTask**](TasksAPI.md#CancelTask) | **Post** /api/v1/agents/{agentId}/tasks/{taskId}/cancel | Cancel an in-flight task.
[**ContinueTask**](TasksAPI.md#ContinueTask) | **Post** /api/v1/agents/{agentId}/tasks/{taskId}/continue | Resume a task from input_required / auth_required.
[**CreateTask**](TasksAPI.md#CreateTask) | **Post** /api/v1/agents/{agentId}/tasks | Submit an async task to an agent.
[**DeleteTaskWebhook**](TasksAPI.md#DeleteTaskWebhook) | **Delete** /api/v1/agents/{agentId}/tasks/{taskId}/webhooks/{webhookId} | Unsubscribe a webhook.
[**GetTask**](TasksAPI.md#GetTask) | **Get** /api/v1/agents/{agentId}/tasks/{taskId} | Get a task&#39;s current state snapshot.
[**ListAgentTasks**](TasksAPI.md#ListAgentTasks) | **Get** /api/v1/agents/{agentId}/tasks | List the caller&#39;s tasks for this agent.
[**ListTaskMessages**](TasksAPI.md#ListTaskMessages) | **Get** /api/v1/agents/{agentId}/tasks/{taskId}/messages | List the persisted message log for a task (paged by offset).
[**ListTaskWebhooks**](TasksAPI.md#ListTaskWebhooks) | **Get** /api/v1/agents/{agentId}/tasks/{taskId}/webhooks | List registered webhooks for a task.
[**ListUserTasks**](TasksAPI.md#ListUserTasks) | **Get** /api/v1/tasks | List all of the caller&#39;s tasks across every agent.
[**ListWebhookDeliveries**](TasksAPI.md#ListWebhookDeliveries) | **Get** /api/v1/agents/{agentId}/tasks/{taskId}/webhooks/{webhookId}/deliveries | List webhook delivery attempts (audit log).
[**RedeliverWebhook**](TasksAPI.md#RedeliverWebhook) | **Post** /api/v1/agents/{agentId}/tasks/{taskId}/webhooks/{webhookId}/deliveries/{deliveryId}/redeliver | Manually replay a webhook delivery.
[**RegisterTaskWebhook**](TasksAPI.md#RegisterTaskWebhook) | **Post** /api/v1/agents/{agentId}/tasks/{taskId}/webhooks | Register a webhook callback for terminal task events.
[**StreamTaskEvents**](TasksAPI.md#StreamTaskEvents) | **Get** /api/v1/agents/{agentId}/tasks/{taskId}/events | Subscribe to a task&#39;s event stream (SSE).



## CancelTask

> TaskResponse CancelTask(ctx, agentId, taskId).CancelTaskRequest(cancelTaskRequest).Execute()

Cancel an in-flight task.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/beeos-ai/sdk-go"
)

func main() {
	agentId := "agentId_example" // string | 
	taskId := "taskId_example" // string | 
	cancelTaskRequest := *openapiclient.NewCancelTaskRequest() // CancelTaskRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.CancelTask(context.Background(), agentId, taskId).CancelTaskRequest(cancelTaskRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.CancelTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelTask`: TaskResponse
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.CancelTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cancelTaskRequest** | [**CancelTaskRequest**](CancelTaskRequest.md) |  | 

### Return type

[**TaskResponse**](TaskResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContinueTask

> TaskResponse ContinueTask(ctx, agentId, taskId).ContinueTaskRequest(continueTaskRequest).Execute()

Resume a task from input_required / auth_required.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/beeos-ai/sdk-go"
)

func main() {
	agentId := "agentId_example" // string | 
	taskId := "taskId_example" // string | 
	continueTaskRequest := *openapiclient.NewContinueTaskRequest() // ContinueTaskRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.ContinueTask(context.Background(), agentId, taskId).ContinueTaskRequest(continueTaskRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.ContinueTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContinueTask`: TaskResponse
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.ContinueTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContinueTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **continueTaskRequest** | [**ContinueTaskRequest**](ContinueTaskRequest.md) |  | 

### Return type

[**TaskResponse**](TaskResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTask

> TaskCreatedResponse CreateTask(ctx, agentId).CreateTaskRequest(createTaskRequest).Execute()

Submit an async task to an agent.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/beeos-ai/sdk-go"
)

func main() {
	agentId := "agentId_example" // string | 
	createTaskRequest := *openapiclient.NewCreateTaskRequest("Message_example") // CreateTaskRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.CreateTask(context.Background(), agentId).CreateTaskRequest(createTaskRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.CreateTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTask`: TaskCreatedResponse
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.CreateTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createTaskRequest** | [**CreateTaskRequest**](CreateTaskRequest.md) |  | 

### Return type

[**TaskCreatedResponse**](TaskCreatedResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTaskWebhook

> DeleteTaskWebhook(ctx, agentId, taskId, webhookId).Execute()

Unsubscribe a webhook.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/beeos-ai/sdk-go"
)

func main() {
	agentId := "agentId_example" // string | 
	taskId := "taskId_example" // string | 
	webhookId := "webhookId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TasksAPI.DeleteTaskWebhook(context.Background(), agentId, taskId, webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.DeleteTaskWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 
**taskId** | **string** |  | 
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTaskWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTask

> TaskResponse GetTask(ctx, agentId, taskId).Execute()

Get a task's current state snapshot.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/beeos-ai/sdk-go"
)

func main() {
	agentId := "agentId_example" // string | 
	taskId := "taskId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.GetTask(context.Background(), agentId, taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.GetTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTask`: TaskResponse
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.GetTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TaskResponse**](TaskResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAgentTasks

> ListTasksResponse ListAgentTasks(ctx, agentId).State(state).Since(since).Limit(limit).Execute()

List the caller's tasks for this agent.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/beeos-ai/sdk-go"
)

func main() {
	agentId := "agentId_example" // string | 
	state := "state_example" // string | Filter by channel state. `active` is the backward-compat wire alias for v2's `open`; `all` returns both. Default: `active`.  (optional) (default to "active")
	since := "since_example" // string | Pagination cursor returned by the previous call (`next_since`). (optional)
	limit := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.ListAgentTasks(context.Background(), agentId).State(state).Since(since).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.ListAgentTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAgentTasks`: ListTasksResponse
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.ListAgentTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAgentTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **state** | **string** | Filter by channel state. &#x60;active&#x60; is the backward-compat wire alias for v2&#39;s &#x60;open&#x60;; &#x60;all&#x60; returns both. Default: &#x60;active&#x60;.  | [default to &quot;active&quot;]
 **since** | **string** | Pagination cursor returned by the previous call (&#x60;next_since&#x60;). | 
 **limit** | **int32** |  | 

### Return type

[**ListTasksResponse**](ListTasksResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTaskMessages

> ListMessagesResponse ListTaskMessages(ctx, agentId, taskId).Since(since).Limit(limit).IncludeDeltas(includeDeltas).Execute()

List the persisted message log for a task (paged by offset).



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/beeos-ai/sdk-go"
)

func main() {
	agentId := "agentId_example" // string | 
	taskId := "taskId_example" // string | 
	since := int64(789) // int64 | Return messages with `offset` > `since`. Combine with `limit` for paginated reads. Default is 0 (start of task log).  (optional)
	limit := int32(56) // int32 |  (optional)
	includeDeltas := true // bool | When `true`, include ephemeral streaming chunk rows (`agent_reply_delta`, `agent_thought_chunk`, `agent_message_chunk`) in the response. Default `false`. Accepts `true|1|yes` (case-insensitive); any other value is treated as `false`.  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.ListTaskMessages(context.Background(), agentId, taskId).Since(since).Limit(limit).IncludeDeltas(includeDeltas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.ListTaskMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTaskMessages`: ListMessagesResponse
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.ListTaskMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTaskMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **since** | **int64** | Return messages with &#x60;offset&#x60; &gt; &#x60;since&#x60;. Combine with &#x60;limit&#x60; for paginated reads. Default is 0 (start of task log).  | 
 **limit** | **int32** |  | 
 **includeDeltas** | **bool** | When &#x60;true&#x60;, include ephemeral streaming chunk rows (&#x60;agent_reply_delta&#x60;, &#x60;agent_thought_chunk&#x60;, &#x60;agent_message_chunk&#x60;) in the response. Default &#x60;false&#x60;. Accepts &#x60;true|1|yes&#x60; (case-insensitive); any other value is treated as &#x60;false&#x60;.  | [default to false]

### Return type

[**ListMessagesResponse**](ListMessagesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTaskWebhooks

> ListTaskWebhooksResponse ListTaskWebhooks(ctx, agentId, taskId).Execute()

List registered webhooks for a task.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/beeos-ai/sdk-go"
)

func main() {
	agentId := "agentId_example" // string | 
	taskId := "taskId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.ListTaskWebhooks(context.Background(), agentId, taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.ListTaskWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTaskWebhooks`: ListTaskWebhooksResponse
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.ListTaskWebhooks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTaskWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListTaskWebhooksResponse**](ListTaskWebhooksResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserTasks

> ListTasksResponse ListUserTasks(ctx).AgentId(agentId).InstanceId(instanceId).State(state).Since(since).Limit(limit).Execute()

List all of the caller's tasks across every agent.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/beeos-ai/sdk-go"
)

func main() {
	agentId := "agentId_example" // string | Narrow the result to a single agent. Equivalent to calling `GET /agents/{agentId}/tasks` directly; kept on this endpoint so callers can collapse their pagination logic across the \"all agents\" and \"single agent\" paths.  (optional)
	instanceId := "instanceId_example" // string | Narrow the result to tasks running on a specific instance (across any agent on that instance). Combine with `state=open` to answer \"is this instance currently running a task?\" — a non-empty result means busy. Device control commands (computer / mobile) are NOT counted here.  (optional)
	state := "state_example" // string | Filter by channel state. `active` is the legacy wire alias for v2's `open`; `all` returns both. Default: `active`.  (optional) (default to "active")
	since := "since_example" // string | Opaque pagination cursor from the previous page's `next_since`. (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.ListUserTasks(context.Background()).AgentId(agentId).InstanceId(instanceId).State(state).Since(since).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.ListUserTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUserTasks`: ListTasksResponse
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.ListUserTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUserTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agentId** | **string** | Narrow the result to a single agent. Equivalent to calling &#x60;GET /agents/{agentId}/tasks&#x60; directly; kept on this endpoint so callers can collapse their pagination logic across the \&quot;all agents\&quot; and \&quot;single agent\&quot; paths.  | 
 **instanceId** | **string** | Narrow the result to tasks running on a specific instance (across any agent on that instance). Combine with &#x60;state&#x3D;open&#x60; to answer \&quot;is this instance currently running a task?\&quot; — a non-empty result means busy. Device control commands (computer / mobile) are NOT counted here.  | 
 **state** | **string** | Filter by channel state. &#x60;active&#x60; is the legacy wire alias for v2&#39;s &#x60;open&#x60;; &#x60;all&#x60; returns both. Default: &#x60;active&#x60;.  | [default to &quot;active&quot;]
 **since** | **string** | Opaque pagination cursor from the previous page&#39;s &#x60;next_since&#x60;. | 
 **limit** | **int32** |  | [default to 50]

### Return type

[**ListTasksResponse**](ListTasksResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWebhookDeliveries

> ListWebhookDeliveriesResponse ListWebhookDeliveries(ctx, agentId, taskId, webhookId).Limit(limit).Execute()

List webhook delivery attempts (audit log).



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/beeos-ai/sdk-go"
)

func main() {
	agentId := "agentId_example" // string | 
	taskId := "taskId_example" // string | 
	webhookId := "webhookId_example" // string | 
	limit := int32(56) // int32 | Maximum number of delivery rows to return. Clamped to [1, 200]; default 50.  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.ListWebhookDeliveries(context.Background(), agentId, taskId, webhookId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.ListWebhookDeliveries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWebhookDeliveries`: ListWebhookDeliveriesResponse
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.ListWebhookDeliveries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 
**taskId** | **string** |  | 
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWebhookDeliveriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **int32** | Maximum number of delivery rows to return. Clamped to [1, 200]; default 50.  | [default to 50]

### Return type

[**ListWebhookDeliveriesResponse**](ListWebhookDeliveriesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RedeliverWebhook

> RedeliverWebhook202Response RedeliverWebhook(ctx, agentId, taskId, webhookId, deliveryId).Execute()

Manually replay a webhook delivery.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/beeos-ai/sdk-go"
)

func main() {
	agentId := "agentId_example" // string | 
	taskId := "taskId_example" // string | 
	webhookId := "webhookId_example" // string | 
	deliveryId := "deliveryId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.RedeliverWebhook(context.Background(), agentId, taskId, webhookId, deliveryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.RedeliverWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RedeliverWebhook`: RedeliverWebhook202Response
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.RedeliverWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 
**taskId** | **string** |  | 
**webhookId** | **string** |  | 
**deliveryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRedeliverWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**RedeliverWebhook202Response**](RedeliverWebhook202Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterTaskWebhook

> TaskWebhookResponse RegisterTaskWebhook(ctx, agentId, taskId).RegisterTaskWebhookRequest(registerTaskWebhookRequest).Execute()

Register a webhook callback for terminal task events.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/beeos-ai/sdk-go"
)

func main() {
	agentId := "agentId_example" // string | 
	taskId := "taskId_example" // string | 
	registerTaskWebhookRequest := *openapiclient.NewRegisterTaskWebhookRequest("Url_example") // RegisterTaskWebhookRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.RegisterTaskWebhook(context.Background(), agentId, taskId).RegisterTaskWebhookRequest(registerTaskWebhookRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.RegisterTaskWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterTaskWebhook`: TaskWebhookResponse
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.RegisterTaskWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterTaskWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **registerTaskWebhookRequest** | [**RegisterTaskWebhookRequest**](RegisterTaskWebhookRequest.md) |  | 

### Return type

[**TaskWebhookResponse**](TaskWebhookResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StreamTaskEvents

> string StreamTaskEvents(ctx, agentId, taskId).Since(since).Execute()

Subscribe to a task's event stream (SSE).



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/beeos-ai/sdk-go"
)

func main() {
	agentId := "agentId_example" // string | 
	taskId := "taskId_example" // string | 
	since := int64(789) // int64 | Replay cursor. Matches the `offset` field on the last received event — clients persist this value and pass it back on reconnect to resume without gaps or duplicates (Last-Event-ID semantics; the wire spelling stays `since` for backward compatibility with the pre-v2 SSE clients).  Special values: * omitted / `0`  — full history replay (per MS `Subscribe`   contract). Useful for late-attaching SSE clients that want   the full task transcript. * `<offset>`     — resume strictly AFTER the given offset.   The first event the client receives has `offset` > since.  The wire contract guarantees that the `offset` field on emitted events is **strictly monotonically increasing per channel** but is **NOT guaranteed contiguous** (ADR-0022 §1.2). A producer-side storage failure may leave a small hole — e.g. you may observe `... 40, 42, ...` with `41` missing. Clients MUST treat `offset > since` as the resume invariant, never `offset == since + 1`. Per ADR-0022 the retention window is 24 h after last-touch (durable rows) and 24 h or 10 000 entries (ephemeral chunk stream, whichever is tighter); outside that window SSE replies with a `backfill_truncated` frame before resuming.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.StreamTaskEvents(context.Background(), agentId, taskId).Since(since).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.StreamTaskEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StreamTaskEvents`: string
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.StreamTaskEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStreamTaskEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **since** | **int64** | Replay cursor. Matches the &#x60;offset&#x60; field on the last received event — clients persist this value and pass it back on reconnect to resume without gaps or duplicates (Last-Event-ID semantics; the wire spelling stays &#x60;since&#x60; for backward compatibility with the pre-v2 SSE clients).  Special values: * omitted / &#x60;0&#x60;  — full history replay (per MS &#x60;Subscribe&#x60;   contract). Useful for late-attaching SSE clients that want   the full task transcript. * &#x60;&lt;offset&gt;&#x60;     — resume strictly AFTER the given offset.   The first event the client receives has &#x60;offset&#x60; &gt; since.  The wire contract guarantees that the &#x60;offset&#x60; field on emitted events is **strictly monotonically increasing per channel** but is **NOT guaranteed contiguous** (ADR-0022 §1.2). A producer-side storage failure may leave a small hole — e.g. you may observe &#x60;... 40, 42, ...&#x60; with &#x60;41&#x60; missing. Clients MUST treat &#x60;offset &gt; since&#x60; as the resume invariant, never &#x60;offset &#x3D;&#x3D; since + 1&#x60;. Per ADR-0022 the retention window is 24 h after last-touch (durable rows) and 24 h or 10 000 entries (ephemeral chunk stream, whichever is tighter); outside that window SSE replies with a &#x60;backfill_truncated&#x60; frame before resuming.  | 

### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/event-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

