# RuntimeOperationSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**InstanceId** | **string** |  | 
**Target** | [**RuntimeOperationTarget**](RuntimeOperationTarget.md) |  | 
**Method** | [**ServiceRuntimeMethod**](ServiceRuntimeMethod.md) |  | 
**Capability** | **string** |  | 
**Transport** | **string** |  | 
**Sequence** | **string** |  | 
**Cursor** | Pointer to **string** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Status** | **string** |  | 
**EffectState** | **string** |  | 
**Terminal** | **bool** |  | 
**Progress** | Pointer to [**RuntimeOperationSnapshotProgress**](RuntimeOperationSnapshotProgress.md) |  | [optional] 
**CancelRequest** | Pointer to [**RuntimeOperationSnapshotCancelRequest**](RuntimeOperationSnapshotCancelRequest.md) |  | [optional] 
**Projection** | Pointer to [**RuntimeOperationSnapshotProjection**](RuntimeOperationSnapshotProjection.md) |  | [optional] 
**Result** | Pointer to **interface{}** |  | [optional] 
**Error** | Pointer to [**RuntimeOperationError**](RuntimeOperationError.md) |  | [optional] 

## Methods

### NewRuntimeOperationSnapshot

`func NewRuntimeOperationSnapshot(id string, instanceId string, target RuntimeOperationTarget, method ServiceRuntimeMethod, capability string, transport string, sequence string, createdAt time.Time, updatedAt time.Time, status string, effectState string, terminal bool, ) *RuntimeOperationSnapshot`

NewRuntimeOperationSnapshot instantiates a new RuntimeOperationSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuntimeOperationSnapshotWithDefaults

`func NewRuntimeOperationSnapshotWithDefaults() *RuntimeOperationSnapshot`

NewRuntimeOperationSnapshotWithDefaults instantiates a new RuntimeOperationSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RuntimeOperationSnapshot) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RuntimeOperationSnapshot) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RuntimeOperationSnapshot) SetId(v string)`

SetId sets Id field to given value.


### GetInstanceId

`func (o *RuntimeOperationSnapshot) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *RuntimeOperationSnapshot) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *RuntimeOperationSnapshot) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetTarget

`func (o *RuntimeOperationSnapshot) GetTarget() RuntimeOperationTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *RuntimeOperationSnapshot) GetTargetOk() (*RuntimeOperationTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *RuntimeOperationSnapshot) SetTarget(v RuntimeOperationTarget)`

SetTarget sets Target field to given value.


### GetMethod

`func (o *RuntimeOperationSnapshot) GetMethod() ServiceRuntimeMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *RuntimeOperationSnapshot) GetMethodOk() (*ServiceRuntimeMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *RuntimeOperationSnapshot) SetMethod(v ServiceRuntimeMethod)`

SetMethod sets Method field to given value.


### GetCapability

`func (o *RuntimeOperationSnapshot) GetCapability() string`

GetCapability returns the Capability field if non-nil, zero value otherwise.

### GetCapabilityOk

`func (o *RuntimeOperationSnapshot) GetCapabilityOk() (*string, bool)`

GetCapabilityOk returns a tuple with the Capability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapability

`func (o *RuntimeOperationSnapshot) SetCapability(v string)`

SetCapability sets Capability field to given value.


### GetTransport

`func (o *RuntimeOperationSnapshot) GetTransport() string`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *RuntimeOperationSnapshot) GetTransportOk() (*string, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *RuntimeOperationSnapshot) SetTransport(v string)`

SetTransport sets Transport field to given value.


### GetSequence

`func (o *RuntimeOperationSnapshot) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *RuntimeOperationSnapshot) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *RuntimeOperationSnapshot) SetSequence(v string)`

SetSequence sets Sequence field to given value.


### GetCursor

`func (o *RuntimeOperationSnapshot) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *RuntimeOperationSnapshot) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *RuntimeOperationSnapshot) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *RuntimeOperationSnapshot) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RuntimeOperationSnapshot) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RuntimeOperationSnapshot) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RuntimeOperationSnapshot) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *RuntimeOperationSnapshot) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RuntimeOperationSnapshot) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RuntimeOperationSnapshot) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetStatus

`func (o *RuntimeOperationSnapshot) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RuntimeOperationSnapshot) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RuntimeOperationSnapshot) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetEffectState

`func (o *RuntimeOperationSnapshot) GetEffectState() string`

GetEffectState returns the EffectState field if non-nil, zero value otherwise.

### GetEffectStateOk

`func (o *RuntimeOperationSnapshot) GetEffectStateOk() (*string, bool)`

GetEffectStateOk returns a tuple with the EffectState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectState

`func (o *RuntimeOperationSnapshot) SetEffectState(v string)`

SetEffectState sets EffectState field to given value.


### GetTerminal

`func (o *RuntimeOperationSnapshot) GetTerminal() bool`

GetTerminal returns the Terminal field if non-nil, zero value otherwise.

### GetTerminalOk

`func (o *RuntimeOperationSnapshot) GetTerminalOk() (*bool, bool)`

GetTerminalOk returns a tuple with the Terminal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminal

`func (o *RuntimeOperationSnapshot) SetTerminal(v bool)`

SetTerminal sets Terminal field to given value.


### GetProgress

`func (o *RuntimeOperationSnapshot) GetProgress() RuntimeOperationSnapshotProgress`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *RuntimeOperationSnapshot) GetProgressOk() (*RuntimeOperationSnapshotProgress, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *RuntimeOperationSnapshot) SetProgress(v RuntimeOperationSnapshotProgress)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *RuntimeOperationSnapshot) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetCancelRequest

`func (o *RuntimeOperationSnapshot) GetCancelRequest() RuntimeOperationSnapshotCancelRequest`

GetCancelRequest returns the CancelRequest field if non-nil, zero value otherwise.

### GetCancelRequestOk

`func (o *RuntimeOperationSnapshot) GetCancelRequestOk() (*RuntimeOperationSnapshotCancelRequest, bool)`

GetCancelRequestOk returns a tuple with the CancelRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelRequest

`func (o *RuntimeOperationSnapshot) SetCancelRequest(v RuntimeOperationSnapshotCancelRequest)`

SetCancelRequest sets CancelRequest field to given value.

### HasCancelRequest

`func (o *RuntimeOperationSnapshot) HasCancelRequest() bool`

HasCancelRequest returns a boolean if a field has been set.

### GetProjection

`func (o *RuntimeOperationSnapshot) GetProjection() RuntimeOperationSnapshotProjection`

GetProjection returns the Projection field if non-nil, zero value otherwise.

### GetProjectionOk

`func (o *RuntimeOperationSnapshot) GetProjectionOk() (*RuntimeOperationSnapshotProjection, bool)`

GetProjectionOk returns a tuple with the Projection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjection

`func (o *RuntimeOperationSnapshot) SetProjection(v RuntimeOperationSnapshotProjection)`

SetProjection sets Projection field to given value.

### HasProjection

`func (o *RuntimeOperationSnapshot) HasProjection() bool`

HasProjection returns a boolean if a field has been set.

### GetResult

`func (o *RuntimeOperationSnapshot) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *RuntimeOperationSnapshot) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *RuntimeOperationSnapshot) SetResult(v interface{})`

SetResult sets Result field to given value.

### HasResult

`func (o *RuntimeOperationSnapshot) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResultNil

`func (o *RuntimeOperationSnapshot) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *RuntimeOperationSnapshot) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetError

`func (o *RuntimeOperationSnapshot) GetError() RuntimeOperationError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RuntimeOperationSnapshot) GetErrorOk() (*RuntimeOperationError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RuntimeOperationSnapshot) SetError(v RuntimeOperationError)`

SetError sets Error field to given value.

### HasError

`func (o *RuntimeOperationSnapshot) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


