# RuntimeOperationSnapshotProgress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sequence** | **string** |  | 
**Phase** | **string** |  | 
**Current** | Pointer to **float32** |  | [optional] 
**Total** | Pointer to **float32** |  | [optional] 

## Methods

### NewRuntimeOperationSnapshotProgress

`func NewRuntimeOperationSnapshotProgress(sequence string, phase string, ) *RuntimeOperationSnapshotProgress`

NewRuntimeOperationSnapshotProgress instantiates a new RuntimeOperationSnapshotProgress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuntimeOperationSnapshotProgressWithDefaults

`func NewRuntimeOperationSnapshotProgressWithDefaults() *RuntimeOperationSnapshotProgress`

NewRuntimeOperationSnapshotProgressWithDefaults instantiates a new RuntimeOperationSnapshotProgress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSequence

`func (o *RuntimeOperationSnapshotProgress) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *RuntimeOperationSnapshotProgress) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *RuntimeOperationSnapshotProgress) SetSequence(v string)`

SetSequence sets Sequence field to given value.


### GetPhase

`func (o *RuntimeOperationSnapshotProgress) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *RuntimeOperationSnapshotProgress) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *RuntimeOperationSnapshotProgress) SetPhase(v string)`

SetPhase sets Phase field to given value.


### GetCurrent

`func (o *RuntimeOperationSnapshotProgress) GetCurrent() float32`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *RuntimeOperationSnapshotProgress) GetCurrentOk() (*float32, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *RuntimeOperationSnapshotProgress) SetCurrent(v float32)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *RuntimeOperationSnapshotProgress) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### GetTotal

`func (o *RuntimeOperationSnapshotProgress) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *RuntimeOperationSnapshotProgress) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *RuntimeOperationSnapshotProgress) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *RuntimeOperationSnapshotProgress) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


