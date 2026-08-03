# \ComputerAPI

All URIs are relative to *https://openapi.beeos.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ComputerClick**](ComputerAPI.md#ComputerClick) | **Post** /api/v1/instances/{id}/computer/click | Mouse click at (x, y).
[**ComputerKey**](ComputerAPI.md#ComputerKey) | **Post** /api/v1/instances/{id}/computer/key | Press a key (single) or hotkey combo (multiple keys).
[**ComputerMove**](ComputerAPI.md#ComputerMove) | **Post** /api/v1/instances/{id}/computer/move | Move the pointer to (x, y) without clicking.
[**ComputerScreenshot**](ComputerAPI.md#ComputerScreenshot) | **Post** /api/v1/instances/{id}/computer/screenshot | Capture a desktop screenshot (returns a file reference).
[**ComputerScroll**](ComputerAPI.md#ComputerScroll) | **Post** /api/v1/instances/{id}/computer/scroll | Scroll in a cardinal direction.
[**ComputerType**](ComputerAPI.md#ComputerType) | **Post** /api/v1/instances/{id}/computer/type | Type Unicode text into the focused field.
[**GetComputerInfo**](ComputerAPI.md#GetComputerInfo) | **Get** /api/v1/instances/{id}/computer | Live desktop control info (geometry + supported actions).



## ComputerClick

> InlineObject ComputerClick(ctx, id).ClickRequest(clickRequest).Execute()

Mouse click at (x, y).

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
	id := "id_example" // string | 
	clickRequest := *openapiclient.NewClickRequest(int32(123), int32(123)) // ClickRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerAPI.ComputerClick(context.Background(), id).ClickRequest(clickRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerAPI.ComputerClick``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ComputerClick`: InlineObject
	fmt.Fprintf(os.Stdout, "Response from `ComputerAPI.ComputerClick`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputerClickRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clickRequest** | [**ClickRequest**](ClickRequest.md) |  | 

### Return type

[**InlineObject**](InlineObject.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComputerKey

> InlineObject ComputerKey(ctx, id).KeyRequest(keyRequest).Execute()

Press a key (single) or hotkey combo (multiple keys).

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
	id := "id_example" // string | 
	keyRequest := *openapiclient.NewKeyRequest([]string{"Keys_example"}) // KeyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerAPI.ComputerKey(context.Background(), id).KeyRequest(keyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerAPI.ComputerKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ComputerKey`: InlineObject
	fmt.Fprintf(os.Stdout, "Response from `ComputerAPI.ComputerKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputerKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **keyRequest** | [**KeyRequest**](KeyRequest.md) |  | 

### Return type

[**InlineObject**](InlineObject.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComputerMove

> InlineObject ComputerMove(ctx, id).MoveRequest(moveRequest).Execute()

Move the pointer to (x, y) without clicking.

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
	id := "id_example" // string | 
	moveRequest := *openapiclient.NewMoveRequest(int32(123), int32(123)) // MoveRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerAPI.ComputerMove(context.Background(), id).MoveRequest(moveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerAPI.ComputerMove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ComputerMove`: InlineObject
	fmt.Fprintf(os.Stdout, "Response from `ComputerAPI.ComputerMove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputerMoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **moveRequest** | [**MoveRequest**](MoveRequest.md) |  | 

### Return type

[**InlineObject**](InlineObject.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComputerScreenshot

> ComputerScreenshot200Response ComputerScreenshot(ctx, id).Execute()

Capture a desktop screenshot (returns a file reference).



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerAPI.ComputerScreenshot(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerAPI.ComputerScreenshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ComputerScreenshot`: ComputerScreenshot200Response
	fmt.Fprintf(os.Stdout, "Response from `ComputerAPI.ComputerScreenshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputerScreenshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ComputerScreenshot200Response**](ComputerScreenshot200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComputerScroll

> InlineObject ComputerScroll(ctx, id).ScrollRequest(scrollRequest).Execute()

Scroll in a cardinal direction.

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
	id := "id_example" // string | 
	scrollRequest := *openapiclient.NewScrollRequest("Direction_example") // ScrollRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerAPI.ComputerScroll(context.Background(), id).ScrollRequest(scrollRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerAPI.ComputerScroll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ComputerScroll`: InlineObject
	fmt.Fprintf(os.Stdout, "Response from `ComputerAPI.ComputerScroll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputerScrollRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scrollRequest** | [**ScrollRequest**](ScrollRequest.md) |  | 

### Return type

[**InlineObject**](InlineObject.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComputerType

> InlineObject ComputerType(ctx, id).TypeRequest(typeRequest).Execute()

Type Unicode text into the focused field.

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
	id := "id_example" // string | 
	typeRequest := *openapiclient.NewTypeRequest("Text_example") // TypeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerAPI.ComputerType(context.Background(), id).TypeRequest(typeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerAPI.ComputerType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ComputerType`: InlineObject
	fmt.Fprintf(os.Stdout, "Response from `ComputerAPI.ComputerType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiComputerTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **typeRequest** | [**TypeRequest**](TypeRequest.md) |  | 

### Return type

[**InlineObject**](InlineObject.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComputerInfo

> GetComputerInfo200Response GetComputerInfo(ctx, id).Execute()

Live desktop control info (geometry + supported actions).



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerAPI.GetComputerInfo(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerAPI.GetComputerInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComputerInfo`: GetComputerInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `ComputerAPI.GetComputerInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetComputerInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetComputerInfo200Response**](GetComputerInfo200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

