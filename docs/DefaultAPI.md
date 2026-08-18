# \DefaultAPI

All URIs are relative to *https://openapi.beeos.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelRuntimeOperation**](DefaultAPI.md#CancelRuntimeOperation) | **Post** /api/v1/instances/{instanceId}/operations/{operationId}/cancel | Cancel a runtime operation
[**CreateCanvasSession**](DefaultAPI.md#CreateCanvasSession) | **Post** /api/v1/instances/{instanceId}/canvas-sessions | Create a canvas session
[**CreateTerminalSession**](DefaultAPI.md#CreateTerminalSession) | **Post** /api/v1/instances/{instanceId}/terminal-sessions | Create a terminal session
[**GetRuntimeCapabilities**](DefaultAPI.md#GetRuntimeCapabilities) | **Get** /api/v1/instances/{instanceId}/runtime-capabilities | Get runtime capabilities
[**GetRuntimeOperation**](DefaultAPI.md#GetRuntimeOperation) | **Get** /api/v1/instances/{instanceId}/operations/{operationId} | Get a runtime operation
[**InvokeRuntimeMethod**](DefaultAPI.md#InvokeRuntimeMethod) | **Post** /api/v1/instances/{instanceId}/methods | Invoke a runtime method
[**ListRuntimeOperations**](DefaultAPI.md#ListRuntimeOperations) | **Get** /api/v1/instances/{instanceId}/operations | List active runtime operations
[**StreamRuntimeOperationEvents**](DefaultAPI.md#StreamRuntimeOperationEvents) | **Get** /api/v1/instances/{instanceId}/operations/{operationId}/events | Stream runtime operation events



## CancelRuntimeOperation

> ServiceOperationCancelAccepted CancelRuntimeOperation(ctx, instanceId, operationId).IdempotencyKey(idempotencyKey).XBeeOSOperationId(xBeeOSOperationId).Execute()

Cancel a runtime operation

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
	idempotencyKey := "idempotencyKey_example" // string | 
	instanceId := "instanceId_example" // string | 
	operationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xBeeOSOperationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CancelRuntimeOperation(context.Background(), instanceId, operationId).IdempotencyKey(idempotencyKey).XBeeOSOperationId(xBeeOSOperationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CancelRuntimeOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelRuntimeOperation`: ServiceOperationCancelAccepted
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CancelRuntimeOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**operationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelRuntimeOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** |  | 


 **xBeeOSOperationId** | **string** |  | 

### Return type

[**ServiceOperationCancelAccepted**](ServiceOperationCancelAccepted.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCanvasSession

> map[string]interface{} CreateCanvasSession(ctx, instanceId).CreateCanvasSessionRequest(createCanvasSessionRequest).Execute()

Create a canvas session

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
	instanceId := "instanceId_example" // string | 
	createCanvasSessionRequest := *openapiclient.NewCreateCanvasSessionRequest("PlatformAgentId_example", "ConversationId_example") // CreateCanvasSessionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateCanvasSession(context.Background(), instanceId).CreateCanvasSessionRequest(createCanvasSessionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateCanvasSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCanvasSession`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateCanvasSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCanvasSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createCanvasSessionRequest** | [**CreateCanvasSessionRequest**](CreateCanvasSessionRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTerminalSession

> map[string]interface{} CreateTerminalSession(ctx, instanceId).CreateTerminalSessionRequest(createTerminalSessionRequest).Execute()

Create a terminal session

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
	instanceId := "instanceId_example" // string | 
	createTerminalSessionRequest := *openapiclient.NewCreateTerminalSessionRequest("PlatformAgentId_example") // CreateTerminalSessionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateTerminalSession(context.Background(), instanceId).CreateTerminalSessionRequest(createTerminalSessionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateTerminalSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTerminalSession`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateTerminalSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTerminalSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createTerminalSessionRequest** | [**CreateTerminalSessionRequest**](CreateTerminalSessionRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuntimeCapabilities

> map[string]interface{} GetRuntimeCapabilities(ctx, instanceId).Execute()

Get runtime capabilities

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
	instanceId := "instanceId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRuntimeCapabilities(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRuntimeCapabilities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRuntimeCapabilities`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRuntimeCapabilities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuntimeCapabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuntimeOperation

> RuntimeOperationSnapshot GetRuntimeOperation(ctx, instanceId, operationId).Execute()

Get a runtime operation

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
	instanceId := "instanceId_example" // string | 
	operationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetRuntimeOperation(context.Background(), instanceId, operationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRuntimeOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRuntimeOperation`: RuntimeOperationSnapshot
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRuntimeOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**operationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuntimeOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RuntimeOperationSnapshot**](RuntimeOperationSnapshot.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvokeRuntimeMethod

> RuntimeJSONRPCBusinessError InvokeRuntimeMethod(ctx, instanceId).InvokeRuntimeMethodRequest(invokeRuntimeMethodRequest).IdempotencyKey(idempotencyKey).XBeeOSOperationId(xBeeOSOperationId).Execute()

Invoke a runtime method

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
	instanceId := "instanceId_example" // string | 
	invokeRuntimeMethodRequest := *openapiclient.NewInvokeRuntimeMethodRequest("Jsonrpc_example", interface{}(123), "Method_example", map[string]interface{}(123)) // InvokeRuntimeMethodRequest | 
	idempotencyKey := "idempotencyKey_example" // string |  (optional)
	xBeeOSOperationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.InvokeRuntimeMethod(context.Background(), instanceId).InvokeRuntimeMethodRequest(invokeRuntimeMethodRequest).IdempotencyKey(idempotencyKey).XBeeOSOperationId(xBeeOSOperationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.InvokeRuntimeMethod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvokeRuntimeMethod`: RuntimeJSONRPCBusinessError
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.InvokeRuntimeMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvokeRuntimeMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invokeRuntimeMethodRequest** | [**InvokeRuntimeMethodRequest**](InvokeRuntimeMethodRequest.md) |  | 
 **idempotencyKey** | **string** |  | 
 **xBeeOSOperationId** | **string** |  | 

### Return type

[**RuntimeJSONRPCBusinessError**](RuntimeJSONRPCBusinessError.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRuntimeOperations

> ServiceOperationsListResponse ListRuntimeOperations(ctx, instanceId).Status(status).Cursor(cursor).Limit(limit).Execute()

List active runtime operations

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
	status := "status_example" // string | 
	instanceId := "instanceId_example" // string | 
	cursor := "cursor_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListRuntimeOperations(context.Background(), instanceId).Status(status).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListRuntimeOperations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRuntimeOperations`: ServiceOperationsListResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListRuntimeOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRuntimeOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** |  | 

 **cursor** | **string** |  | 
 **limit** | **int32** |  | 

### Return type

[**ServiceOperationsListResponse**](ServiceOperationsListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StreamRuntimeOperationEvents

> string StreamRuntimeOperationEvents(ctx, instanceId, operationId).LastEventID(lastEventID).Execute()

Stream runtime operation events

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
	instanceId := "instanceId_example" // string | 
	operationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	lastEventID := "lastEventID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.StreamRuntimeOperationEvents(context.Background(), instanceId, operationId).LastEventID(lastEventID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.StreamRuntimeOperationEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StreamRuntimeOperationEvents`: string
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.StreamRuntimeOperationEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**operationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStreamRuntimeOperationEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **lastEventID** | **string** |  | 

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

