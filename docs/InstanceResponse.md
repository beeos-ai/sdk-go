# InstanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Data** | [**InstanceDataDTO**](InstanceDataDTO.md) |  | 

## Methods

### NewInstanceResponse

`func NewInstanceResponse(success bool, data InstanceDataDTO, ) *InstanceResponse`

NewInstanceResponse instantiates a new InstanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceResponseWithDefaults

`func NewInstanceResponseWithDefaults() *InstanceResponse`

NewInstanceResponseWithDefaults instantiates a new InstanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *InstanceResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *InstanceResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *InstanceResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *InstanceResponse) GetData() InstanceDataDTO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InstanceResponse) GetDataOk() (*InstanceDataDTO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InstanceResponse) SetData(v InstanceDataDTO)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


