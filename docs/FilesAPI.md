# \FilesAPI

All URIs are relative to *https://openapi.beeos.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFile**](FilesAPI.md#GetFile) | **Get** /api/v1/files/{id} | Resolve a file_id to metadata + presigned download URL.
[**PresignFileUpload**](FilesAPI.md#PresignFileUpload) | **Post** /api/v1/files/presign-upload | Request a presigned upload URL.



## GetFile

> GetFile200Response GetFile(ctx, id).Execute()

Resolve a file_id to metadata + presigned download URL.



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
	resp, r, err := apiClient.FilesAPI.GetFile(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.GetFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFile`: GetFile200Response
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.GetFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetFile200Response**](GetFile200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PresignFileUpload

> PresignFileUpload200Response PresignFileUpload(ctx).PresignUploadRequest(presignUploadRequest).Execute()

Request a presigned upload URL.



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
	presignUploadRequest := *openapiclient.NewPresignUploadRequest("Filename_example", "ContentType_example") // PresignUploadRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.PresignFileUpload(context.Background()).PresignUploadRequest(presignUploadRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.PresignFileUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PresignFileUpload`: PresignFileUpload200Response
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.PresignFileUpload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPresignFileUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **presignUploadRequest** | [**PresignUploadRequest**](PresignUploadRequest.md) |  | 

### Return type

[**PresignFileUpload200Response**](PresignFileUpload200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

