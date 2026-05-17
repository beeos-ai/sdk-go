# InstancePageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Data** | [**[]InstanceDataDTO**](InstanceDataDTO.md) |  | 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewInstancePageResponse

`func NewInstancePageResponse(success bool, data []InstanceDataDTO, ) *InstancePageResponse`

NewInstancePageResponse instantiates a new InstancePageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstancePageResponseWithDefaults

`func NewInstancePageResponseWithDefaults() *InstancePageResponse`

NewInstancePageResponseWithDefaults instantiates a new InstancePageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *InstancePageResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *InstancePageResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *InstancePageResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *InstancePageResponse) GetData() []InstanceDataDTO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InstancePageResponse) GetDataOk() (*[]InstanceDataDTO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InstancePageResponse) SetData(v []InstanceDataDTO)`

SetData sets Data field to given value.


### GetTotal

`func (o *InstancePageResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InstancePageResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InstancePageResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *InstancePageResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


