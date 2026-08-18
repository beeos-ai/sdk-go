# PatchMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** |  | [optional] 
**Parts** | Pointer to [**[]Part**](Part.md) |  | [optional] 
**State** | Pointer to [**MessageState**](MessageState.md) |  | [optional] 
**StopReason** | Pointer to [**StopReason**](StopReason.md) |  | [optional] 

## Methods

### NewPatchMessageRequest

`func NewPatchMessageRequest() *PatchMessageRequest`

NewPatchMessageRequest instantiates a new PatchMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchMessageRequestWithDefaults

`func NewPatchMessageRequestWithDefaults() *PatchMessageRequest`

NewPatchMessageRequestWithDefaults instantiates a new PatchMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *PatchMessageRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *PatchMessageRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *PatchMessageRequest) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *PatchMessageRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetParts

`func (o *PatchMessageRequest) GetParts() []Part`

GetParts returns the Parts field if non-nil, zero value otherwise.

### GetPartsOk

`func (o *PatchMessageRequest) GetPartsOk() (*[]Part, bool)`

GetPartsOk returns a tuple with the Parts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParts

`func (o *PatchMessageRequest) SetParts(v []Part)`

SetParts sets Parts field to given value.

### HasParts

`func (o *PatchMessageRequest) HasParts() bool`

HasParts returns a boolean if a field has been set.

### GetState

`func (o *PatchMessageRequest) GetState() MessageState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PatchMessageRequest) GetStateOk() (*MessageState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PatchMessageRequest) SetState(v MessageState)`

SetState sets State field to given value.

### HasState

`func (o *PatchMessageRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStopReason

`func (o *PatchMessageRequest) GetStopReason() StopReason`

GetStopReason returns the StopReason field if non-nil, zero value otherwise.

### GetStopReasonOk

`func (o *PatchMessageRequest) GetStopReasonOk() (*StopReason, bool)`

GetStopReasonOk returns a tuple with the StopReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopReason

`func (o *PatchMessageRequest) SetStopReason(v StopReason)`

SetStopReason sets StopReason field to given value.

### HasStopReason

`func (o *PatchMessageRequest) HasStopReason() bool`

HasStopReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


