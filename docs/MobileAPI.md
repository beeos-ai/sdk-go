# \MobileAPI

All URIs are relative to *https://openapi.beeos.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMobileInfo**](MobileAPI.md#GetMobileInfo) | **Get** /api/v1/instances/{id}/mobile | Live mobile control info (geometry + supported actions).
[**MobileDoubleTap**](MobileAPI.md#MobileDoubleTap) | **Post** /api/v1/instances/{id}/mobile/double_tap | Double-tap at (x, y).
[**MobileDrag**](MobileAPI.md#MobileDrag) | **Post** /api/v1/instances/{id}/mobile/drag | Drag an object from (x1, y1) to (x2, y2).
[**MobileGetUiTree**](MobileAPI.md#MobileGetUiTree) | **Get** /api/v1/instances/{id}/mobile/ui_tree | Read the current Android UI hierarchy.
[**MobileKey**](MobileAPI.md#MobileKey) | **Post** /api/v1/instances/{id}/mobile/key | Press an Android key (KeyCode name).
[**MobileListApps**](MobileAPI.md#MobileListApps) | **Get** /api/v1/instances/{id}/mobile/apps | List installed Android applications.
[**MobileLongPress**](MobileAPI.md#MobileLongPress) | **Post** /api/v1/instances/{id}/mobile/long_press | Long-press at (x, y).
[**MobileOpenApp**](MobileAPI.md#MobileOpenApp) | **Post** /api/v1/instances/{id}/mobile/open_app | Open an installed app by package id or display name.
[**MobilePressButton**](MobileAPI.md#MobilePressButton) | **Post** /api/v1/instances/{id}/mobile/press_button | Press a hardware / navigation button (back / home / ...).
[**MobileScreenshot**](MobileAPI.md#MobileScreenshot) | **Post** /api/v1/instances/{id}/mobile/screenshot | Capture a mobile screenshot (returns a file reference).
[**MobileScroll**](MobileAPI.md#MobileScroll) | **Post** /api/v1/instances/{id}/mobile/scroll | Scroll in a cardinal direction.
[**MobileSwipe**](MobileAPI.md#MobileSwipe) | **Post** /api/v1/instances/{id}/mobile/swipe | Swipe from (x1, y1) to (x2, y2).
[**MobileTap**](MobileAPI.md#MobileTap) | **Post** /api/v1/instances/{id}/mobile/tap | Tap at (x, y).
[**MobileType**](MobileAPI.md#MobileType) | **Post** /api/v1/instances/{id}/mobile/type | Type Unicode text into the focused field.



## GetMobileInfo

> GetComputerInfo200Response GetMobileInfo(ctx, id).Execute()

Live mobile control info (geometry + supported actions).



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
	resp, r, err := apiClient.MobileAPI.GetMobileInfo(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MobileAPI.GetMobileInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMobileInfo`: GetComputerInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `MobileAPI.GetMobileInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiGetMobileInfoRequest struct via the builder pattern


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


## MobileDoubleTap

> InlineObject MobileDoubleTap(ctx, id).TapRequest(tapRequest).Execute()

Double-tap at (x, y).

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
	tapRequest := *openapiclient.NewTapRequest(int32(123), int32(123)) // TapRequest |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MobileAPI.MobileDoubleTap(context.Background(), id).TapRequest(tapRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MobileAPI.MobileDoubleTap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MobileDoubleTap`: InlineObject
	fmt.Fprintf(os.Stdout, "Response from `MobileAPI.MobileDoubleTap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiMobileDoubleTapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tapRequest** | [**TapRequest**](TapRequest.md) |  |

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


## MobileDrag

> InlineObject MobileDrag(ctx, id).DragRequest(dragRequest).Execute()

Drag an object from (x1, y1) to (x2, y2).

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
	dragRequest := *openapiclient.NewDragRequest(int32(123), int32(123), int32(123), int32(123)) // DragRequest |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MobileAPI.MobileDrag(context.Background(), id).DragRequest(dragRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MobileAPI.MobileDrag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MobileDrag`: InlineObject
	fmt.Fprintf(os.Stdout, "Response from `MobileAPI.MobileDrag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiMobileDragRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dragRequest** | [**DragRequest**](DragRequest.md) |  |

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


## MobileGetUiTree

> MobileGetUiTree200Response MobileGetUiTree(ctx, id).Query(query).Format(format).Execute()

Read the current Android UI hierarchy.

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
	query := "query_example" // string |  (optional)
	format := "format_example" // string |  (optional) (default to "compact")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MobileAPI.MobileGetUiTree(context.Background(), id).Query(query).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MobileAPI.MobileGetUiTree``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MobileGetUiTree`: MobileGetUiTree200Response
	fmt.Fprintf(os.Stdout, "Response from `MobileAPI.MobileGetUiTree`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiMobileGetUiTreeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | **string** |  |
 **format** | **string** |  | [default to &quot;compact&quot;]

### Return type

[**MobileGetUiTree200Response**](MobileGetUiTree200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MobileKey

> InlineObject MobileKey(ctx, id).KeyRequest(keyRequest).Execute()

Press an Android key (KeyCode name).

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
	resp, r, err := apiClient.MobileAPI.MobileKey(context.Background(), id).KeyRequest(keyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MobileAPI.MobileKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MobileKey`: InlineObject
	fmt.Fprintf(os.Stdout, "Response from `MobileAPI.MobileKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiMobileKeyRequest struct via the builder pattern


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


## MobileListApps

> MobileListApps200Response MobileListApps(ctx, id).IncludeSystem(includeSystem).LaunchableOnly(launchableOnly).Execute()

List installed Android applications.

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
	includeSystem := true // bool |  (optional) (default to false)
	launchableOnly := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MobileAPI.MobileListApps(context.Background(), id).IncludeSystem(includeSystem).LaunchableOnly(launchableOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MobileAPI.MobileListApps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MobileListApps`: MobileListApps200Response
	fmt.Fprintf(os.Stdout, "Response from `MobileAPI.MobileListApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiMobileListAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeSystem** | **bool** |  | [default to false]
 **launchableOnly** | **bool** |  | [default to false]

### Return type

[**MobileListApps200Response**](MobileListApps200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MobileLongPress

> InlineObject MobileLongPress(ctx, id).LongPressRequest(longPressRequest).Execute()

Long-press at (x, y).

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
	longPressRequest := *openapiclient.NewLongPressRequest(int32(123), int32(123)) // LongPressRequest |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MobileAPI.MobileLongPress(context.Background(), id).LongPressRequest(longPressRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MobileAPI.MobileLongPress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MobileLongPress`: InlineObject
	fmt.Fprintf(os.Stdout, "Response from `MobileAPI.MobileLongPress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiMobileLongPressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **longPressRequest** | [**LongPressRequest**](LongPressRequest.md) |  |

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


## MobileOpenApp

> InlineObject MobileOpenApp(ctx, id).OpenAppRequest(openAppRequest).Execute()

Open an installed app by package id or display name.

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
	openAppRequest := *openapiclient.NewOpenAppRequest("App_example") // OpenAppRequest |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MobileAPI.MobileOpenApp(context.Background(), id).OpenAppRequest(openAppRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MobileAPI.MobileOpenApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MobileOpenApp`: InlineObject
	fmt.Fprintf(os.Stdout, "Response from `MobileAPI.MobileOpenApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiMobileOpenAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **openAppRequest** | [**OpenAppRequest**](OpenAppRequest.md) |  |

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


## MobilePressButton

> InlineObject MobilePressButton(ctx, id).PressButtonRequest(pressButtonRequest).Execute()

Press a hardware / navigation button (back / home / ...).

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
	pressButtonRequest := *openapiclient.NewPressButtonRequest("Button_example") // PressButtonRequest |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MobileAPI.MobilePressButton(context.Background(), id).PressButtonRequest(pressButtonRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MobileAPI.MobilePressButton``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MobilePressButton`: InlineObject
	fmt.Fprintf(os.Stdout, "Response from `MobileAPI.MobilePressButton`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiMobilePressButtonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pressButtonRequest** | [**PressButtonRequest**](PressButtonRequest.md) |  |

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


## MobileScreenshot

> ComputerScreenshot200Response MobileScreenshot(ctx, id).Execute()

Capture a mobile screenshot (returns a file reference).

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
	resp, r, err := apiClient.MobileAPI.MobileScreenshot(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MobileAPI.MobileScreenshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MobileScreenshot`: ComputerScreenshot200Response
	fmt.Fprintf(os.Stdout, "Response from `MobileAPI.MobileScreenshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiMobileScreenshotRequest struct via the builder pattern


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


## MobileScroll

> InlineObject MobileScroll(ctx, id).ScrollRequest(scrollRequest).Execute()

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
	resp, r, err := apiClient.MobileAPI.MobileScroll(context.Background(), id).ScrollRequest(scrollRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MobileAPI.MobileScroll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MobileScroll`: InlineObject
	fmt.Fprintf(os.Stdout, "Response from `MobileAPI.MobileScroll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiMobileScrollRequest struct via the builder pattern


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


## MobileSwipe

> InlineObject MobileSwipe(ctx, id).SwipeRequest(swipeRequest).Execute()

Swipe from (x1, y1) to (x2, y2).

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
	swipeRequest := *openapiclient.NewSwipeRequest(int32(123), int32(123), int32(123), int32(123)) // SwipeRequest |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MobileAPI.MobileSwipe(context.Background(), id).SwipeRequest(swipeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MobileAPI.MobileSwipe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MobileSwipe`: InlineObject
	fmt.Fprintf(os.Stdout, "Response from `MobileAPI.MobileSwipe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiMobileSwipeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **swipeRequest** | [**SwipeRequest**](SwipeRequest.md) |  |

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


## MobileTap

> InlineObject MobileTap(ctx, id).TapRequest(tapRequest).Execute()

Tap at (x, y).

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
	tapRequest := *openapiclient.NewTapRequest(int32(123), int32(123)) // TapRequest |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MobileAPI.MobileTap(context.Background(), id).TapRequest(tapRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MobileAPI.MobileTap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MobileTap`: InlineObject
	fmt.Fprintf(os.Stdout, "Response from `MobileAPI.MobileTap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiMobileTapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tapRequest** | [**TapRequest**](TapRequest.md) |  |

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


## MobileType

> InlineObject MobileType(ctx, id).TypeRequest(typeRequest).Execute()

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
	resp, r, err := apiClient.MobileAPI.MobileType(context.Background(), id).TypeRequest(typeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MobileAPI.MobileType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MobileType`: InlineObject
	fmt.Fprintf(os.Stdout, "Response from `MobileAPI.MobileType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiMobileTypeRequest struct via the builder pattern


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

