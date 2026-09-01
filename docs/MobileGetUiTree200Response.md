# MobileGetUiTree200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  |
**Data** | Pointer to [**UITreeResult**](UITreeResult.md) |  | [optional]

## Methods

### NewMobileGetUiTree200Response

`func NewMobileGetUiTree200Response(success bool, ) *MobileGetUiTree200Response`

NewMobileGetUiTree200Response instantiates a new MobileGetUiTree200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileGetUiTree200ResponseWithDefaults

`func NewMobileGetUiTree200ResponseWithDefaults() *MobileGetUiTree200Response`

NewMobileGetUiTree200ResponseWithDefaults instantiates a new MobileGetUiTree200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *MobileGetUiTree200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *MobileGetUiTree200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *MobileGetUiTree200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *MobileGetUiTree200Response) GetData() UITreeResult`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MobileGetUiTree200Response) GetDataOk() (*UITreeResult, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MobileGetUiTree200Response) SetData(v UITreeResult)`

SetData sets Data field to given value.

### HasData

`func (o *MobileGetUiTree200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


