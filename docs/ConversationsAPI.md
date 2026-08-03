# \ConversationsAPI

All URIs are relative to *https://openapi.beeos.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConversation**](ConversationsAPI.md#CreateConversation) | **Post** /api/v1/agents/{agentId}/conversations | Open a new conversation.
[**DeleteConversation**](ConversationsAPI.md#DeleteConversation) | **Delete** /api/v1/agents/{agentId}/conversations/{convId} | Close a conversation.
[**GetConversation**](ConversationsAPI.md#GetConversation) | **Get** /api/v1/agents/{agentId}/conversations/{convId} | Get a single conversation&#39;s metadata.
[**ListConversationMessages**](ConversationsAPI.md#ListConversationMessages) | **Get** /api/v1/agents/{agentId}/conversations/{convId}/messages | List messages in a conversation (paged by offset).
[**ListConversations**](ConversationsAPI.md#ListConversations) | **Get** /api/v1/agents/{agentId}/conversations | List conversations for the calling user against this agent.
[**SendConversationMessage**](ConversationsAPI.md#SendConversationMessage) | **Post** /api/v1/agents/{agentId}/conversations/{convId}/messages | Post a new caller message into the conversation.
[**StreamConversationEvents**](ConversationsAPI.md#StreamConversationEvents) | **Get** /api/v1/agents/{agentId}/conversations/{convId}/events | SSE stream of new messages on the conversation.



## CreateConversation

> ConversationResponse CreateConversation(ctx, agentId).CreateConversationRequest(createConversationRequest).Execute()

Open a new conversation.



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
	createConversationRequest := *openapiclient.NewCreateConversationRequest() // CreateConversationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConversationsAPI.CreateConversation(context.Background(), agentId).CreateConversationRequest(createConversationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConversationsAPI.CreateConversation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConversation`: ConversationResponse
	fmt.Fprintf(os.Stdout, "Response from `ConversationsAPI.CreateConversation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateConversationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createConversationRequest** | [**CreateConversationRequest**](CreateConversationRequest.md) |  | 

### Return type

[**ConversationResponse**](ConversationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConversation

> DeleteConversation(ctx, agentId, convId).Execute()

Close a conversation.



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
	convId := "convId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConversationsAPI.DeleteConversation(context.Background(), agentId, convId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConversationsAPI.DeleteConversation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 
**convId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConversationRequest struct via the builder pattern


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


## GetConversation

> ConversationResponse GetConversation(ctx, agentId, convId).Execute()

Get a single conversation's metadata.

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
	convId := "convId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConversationsAPI.GetConversation(context.Background(), agentId, convId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConversationsAPI.GetConversation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConversation`: ConversationResponse
	fmt.Fprintf(os.Stdout, "Response from `ConversationsAPI.GetConversation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 
**convId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConversationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ConversationResponse**](ConversationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConversationMessages

> ListMessagesResponse ListConversationMessages(ctx, agentId, convId).Since(since).Limit(limit).IncludeDeltas(includeDeltas).Execute()

List messages in a conversation (paged by offset).



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
	convId := "convId_example" // string | 
	since := int64(789) // int64 | Return messages with offset > `since`. Combine with `limit` for paginated reads. Default is 0 (start of conversation).  (optional)
	limit := int32(56) // int32 |  (optional)
	includeDeltas := true // bool | When `true`, include ephemeral streaming chunk rows (`agent_reply_delta`, `agent_thought_chunk`, `agent_message_chunk`) in the response. Default `false`. Accepts `true|1|yes` (case-insensitive); any other value is treated as `false`.  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConversationsAPI.ListConversationMessages(context.Background(), agentId, convId).Since(since).Limit(limit).IncludeDeltas(includeDeltas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConversationsAPI.ListConversationMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListConversationMessages`: ListMessagesResponse
	fmt.Fprintf(os.Stdout, "Response from `ConversationsAPI.ListConversationMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 
**convId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConversationMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **since** | **int64** | Return messages with offset &gt; &#x60;since&#x60;. Combine with &#x60;limit&#x60; for paginated reads. Default is 0 (start of conversation).  | 
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


## ListConversations

> ListConversationsResponse ListConversations(ctx, agentId).State(state).Since(since).Limit(limit).Execute()

List conversations for the calling user against this agent.



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
	state := "state_example" // string | Filter by conversation state. `active` is a backward-compat wire alias for v2's `open` — both return non-closed conversations. `all` returns both open and closed. Default: `active`.  (optional) (default to "active")
	since := "since_example" // string | Pagination cursor returned by the previous call (`next_since`). (optional)
	limit := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConversationsAPI.ListConversations(context.Background(), agentId).State(state).Since(since).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConversationsAPI.ListConversations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListConversations`: ListConversationsResponse
	fmt.Fprintf(os.Stdout, "Response from `ConversationsAPI.ListConversations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConversationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **state** | **string** | Filter by conversation state. &#x60;active&#x60; is a backward-compat wire alias for v2&#39;s &#x60;open&#x60; — both return non-closed conversations. &#x60;all&#x60; returns both open and closed. Default: &#x60;active&#x60;.  | [default to &quot;active&quot;]
 **since** | **string** | Pagination cursor returned by the previous call (&#x60;next_since&#x60;). | 
 **limit** | **int32** |  | 

### Return type

[**ListConversationsResponse**](ListConversationsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendConversationMessage

> SendMessageResponse SendConversationMessage(ctx, agentId, convId).SendMessageRequest(sendMessageRequest).Execute()

Post a new caller message into the conversation.



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
	convId := "convId_example" // string | 
	sendMessageRequest := *openapiclient.NewSendMessageRequest("Message_example") // SendMessageRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConversationsAPI.SendConversationMessage(context.Background(), agentId, convId).SendMessageRequest(sendMessageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConversationsAPI.SendConversationMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendConversationMessage`: SendMessageResponse
	fmt.Fprintf(os.Stdout, "Response from `ConversationsAPI.SendConversationMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 
**convId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendConversationMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sendMessageRequest** | [**SendMessageRequest**](SendMessageRequest.md) |  | 

### Return type

[**SendMessageResponse**](SendMessageResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StreamConversationEvents

> string StreamConversationEvents(ctx, agentId, convId).Since(since).Execute()

SSE stream of new messages on the conversation.



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
	convId := "convId_example" // string | 
	since := int64(789) // int64 | Replay cursor. Matches the `offset` field on the last received event — clients persist this value and pass it back on reconnect to resume without gaps or duplicates (Last-Event-ID semantics; the wire spelling stays `since` for backward compatibility with the pre-v2 SSE clients).  Special values: * omitted / `0`  — full history replay. * `<offset>`     — resume strictly AFTER the given offset.  Unlike the task variant, this stream never auto-closes on a terminal envelope — conversations stay open across many turns and emit further `agent_reply` events. Persist `since` across reconnect ranges to avoid dropping turns during network blips.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConversationsAPI.StreamConversationEvents(context.Background(), agentId, convId).Since(since).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConversationsAPI.StreamConversationEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StreamConversationEvents`: string
	fmt.Fprintf(os.Stdout, "Response from `ConversationsAPI.StreamConversationEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 
**convId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStreamConversationEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **since** | **int64** | Replay cursor. Matches the &#x60;offset&#x60; field on the last received event — clients persist this value and pass it back on reconnect to resume without gaps or duplicates (Last-Event-ID semantics; the wire spelling stays &#x60;since&#x60; for backward compatibility with the pre-v2 SSE clients).  Special values: * omitted / &#x60;0&#x60;  — full history replay. * &#x60;&lt;offset&gt;&#x60;     — resume strictly AFTER the given offset.  Unlike the task variant, this stream never auto-closes on a terminal envelope — conversations stay open across many turns and emit further &#x60;agent_reply&#x60; events. Persist &#x60;since&#x60; across reconnect ranges to avoid dropping turns during network blips.  | 

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

