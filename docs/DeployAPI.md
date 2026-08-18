# \DeployAPI

All URIs are relative to *https://openapi.beeos.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDeployModels**](DeployAPI.md#ListDeployModels) | **Get** /api/v1/deploy/models | List deploy models
[**ListDeployRegions**](DeployAPI.md#ListDeployRegions) | **Get** /api/v1/deploy/regions | List deploy regions
[**ListProviders**](DeployAPI.md#ListProviders) | **Get** /api/v1/providers | List available providers (catalog)



## ListDeployModels

> DeployModelsListResponse ListDeployModels(ctx).AgentFramework(agentFramework).Execute()

List deploy models

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
	agentFramework := "agentFramework_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeployAPI.ListDeployModels(context.Background()).AgentFramework(agentFramework).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeployAPI.ListDeployModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDeployModels`: DeployModelsListResponse
	fmt.Fprintf(os.Stdout, "Response from `DeployAPI.ListDeployModels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDeployModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agentFramework** | **string** |  | 

### Return type

[**DeployModelsListResponse**](DeployModelsListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDeployRegions

> DeployRegionsListResponse ListDeployRegions(ctx).ProviderId(providerId).Execute()

List deploy regions

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
	providerId := "providerId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeployAPI.ListDeployRegions(context.Background()).ProviderId(providerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeployAPI.ListDeployRegions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDeployRegions`: DeployRegionsListResponse
	fmt.Fprintf(os.Stdout, "Response from `DeployAPI.ListDeployRegions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDeployRegionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerId** | **string** |  | 

### Return type

[**DeployRegionsListResponse**](DeployRegionsListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProviders

> ProvidersListResponse ListProviders(ctx).Execute()

List available providers (catalog)

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeployAPI.ListProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeployAPI.ListProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProviders`: ProvidersListResponse
	fmt.Fprintf(os.Stdout, "Response from `DeployAPI.ListProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListProvidersRequest struct via the builder pattern


### Return type

[**ProvidersListResponse**](ProvidersListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

