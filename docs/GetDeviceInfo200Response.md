# GetDeviceInfo200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Data** | Pointer to [**DeviceInfoDTO**](DeviceInfoDTO.md) |  | [optional] 

## Methods

### NewGetDeviceInfo200Response

`func NewGetDeviceInfo200Response(success bool, ) *GetDeviceInfo200Response`

NewGetDeviceInfo200Response instantiates a new GetDeviceInfo200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeviceInfo200ResponseWithDefaults

`func NewGetDeviceInfo200ResponseWithDefaults() *GetDeviceInfo200Response`

NewGetDeviceInfo200ResponseWithDefaults instantiates a new GetDeviceInfo200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *GetDeviceInfo200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetDeviceInfo200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetDeviceInfo200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *GetDeviceInfo200Response) GetData() DeviceInfoDTO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetDeviceInfo200Response) GetDataOk() (*DeviceInfoDTO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetDeviceInfo200Response) SetData(v DeviceInfoDTO)`

SetData sets Data field to given value.

### HasData

`func (o *GetDeviceInfo200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


