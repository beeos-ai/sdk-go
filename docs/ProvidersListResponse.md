# ProvidersListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Data** | [**[]ProviderCatalogItem**](ProviderCatalogItem.md) |  | 

## Methods

### NewProvidersListResponse

`func NewProvidersListResponse(success bool, data []ProviderCatalogItem, ) *ProvidersListResponse`

NewProvidersListResponse instantiates a new ProvidersListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvidersListResponseWithDefaults

`func NewProvidersListResponseWithDefaults() *ProvidersListResponse`

NewProvidersListResponseWithDefaults instantiates a new ProvidersListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ProvidersListResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ProvidersListResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ProvidersListResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *ProvidersListResponse) GetData() []ProviderCatalogItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ProvidersListResponse) GetDataOk() (*[]ProviderCatalogItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ProvidersListResponse) SetData(v []ProviderCatalogItem)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


