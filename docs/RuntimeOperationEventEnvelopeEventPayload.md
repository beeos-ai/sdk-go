# RuntimeOperationEventEnvelopeEventPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClaimedAt** | Pointer to **time.Time** |  | [optional] 
**ExecutionMode** | Pointer to **string** |  | [optional] 
**Phase** | Pointer to **string** |  | [optional] 
**Current** | Pointer to **float32** |  | [optional] 
**Total** | Pointer to **float32** |  | [optional] 
**CommittedAt** | Pointer to **time.Time** |  | [optional] 
**ProjectionRevision** | Pointer to **string** |  | [optional] 
**Revision** | Pointer to **string** |  | [optional] 
**PlatformAgentId** | Pointer to **string** |  | [optional] 
**ErrorCode** | Pointer to **string** |  | [optional] 
**CauseCode** | Pointer to **string** |  | [optional] 
**RequestedBy** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 

## Methods

### NewRuntimeOperationEventEnvelopeEventPayload

`func NewRuntimeOperationEventEnvelopeEventPayload() *RuntimeOperationEventEnvelopeEventPayload`

NewRuntimeOperationEventEnvelopeEventPayload instantiates a new RuntimeOperationEventEnvelopeEventPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuntimeOperationEventEnvelopeEventPayloadWithDefaults

`func NewRuntimeOperationEventEnvelopeEventPayloadWithDefaults() *RuntimeOperationEventEnvelopeEventPayload`

NewRuntimeOperationEventEnvelopeEventPayloadWithDefaults instantiates a new RuntimeOperationEventEnvelopeEventPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaimedAt

`func (o *RuntimeOperationEventEnvelopeEventPayload) GetClaimedAt() time.Time`

GetClaimedAt returns the ClaimedAt field if non-nil, zero value otherwise.

### GetClaimedAtOk

`func (o *RuntimeOperationEventEnvelopeEventPayload) GetClaimedAtOk() (*time.Time, bool)`

GetClaimedAtOk returns a tuple with the ClaimedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimedAt

`func (o *RuntimeOperationEventEnvelopeEventPayload) SetClaimedAt(v time.Time)`

SetClaimedAt sets ClaimedAt field to given value.

### HasClaimedAt

`func (o *RuntimeOperationEventEnvelopeEventPayload) HasClaimedAt() bool`

HasClaimedAt returns a boolean if a field has been set.

### GetExecutionMode

`func (o *RuntimeOperationEventEnvelopeEventPayload) GetExecutionMode() string`

GetExecutionMode returns the ExecutionMode field if non-nil, zero value otherwise.

### GetExecutionModeOk

`func (o *RuntimeOperationEventEnvelopeEventPayload) GetExecutionModeOk() (*string, bool)`

GetExecutionModeOk returns a tuple with the ExecutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionMode

`func (o *RuntimeOperationEventEnvelopeEventPayload) SetExecutionMode(v string)`

SetExecutionMode sets ExecutionMode field to given value.

### HasExecutionMode

`func (o *RuntimeOperationEventEnvelopeEventPayload) HasExecutionMode() bool`

HasExecutionMode returns a boolean if a field has been set.

### GetPhase

`func (o *RuntimeOperationEventEnvelopeEventPayload) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *RuntimeOperationEventEnvelopeEventPayload) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *RuntimeOperationEventEnvelopeEventPayload) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *RuntimeOperationEventEnvelopeEventPayload) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetCurrent

`func (o *RuntimeOperationEventEnvelopeEventPayload) GetCurrent() float32`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *RuntimeOperationEventEnvelopeEventPayload) GetCurrentOk() (*float32, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *RuntimeOperationEventEnvelopeEventPayload) SetCurrent(v float32)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *RuntimeOperationEventEnvelopeEventPayload) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### GetTotal

`func (o *RuntimeOperationEventEnvelopeEventPayload) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *RuntimeOperationEventEnvelopeEventPayload) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *RuntimeOperationEventEnvelopeEventPayload) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *RuntimeOperationEventEnvelopeEventPayload) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetCommittedAt

`func (o *RuntimeOperationEventEnvelopeEventPayload) GetCommittedAt() time.Time`

GetCommittedAt returns the CommittedAt field if non-nil, zero value otherwise.

### GetCommittedAtOk

`func (o *RuntimeOperationEventEnvelopeEventPayload) GetCommittedAtOk() (*time.Time, bool)`

GetCommittedAtOk returns a tuple with the CommittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommittedAt

`func (o *RuntimeOperationEventEnvelopeEventPayload) SetCommittedAt(v time.Time)`

SetCommittedAt sets CommittedAt field to given value.

### HasCommittedAt

`func (o *RuntimeOperationEventEnvelopeEventPayload) HasCommittedAt() bool`

HasCommittedAt returns a boolean if a field has been set.

### GetProjectionRevision

`func (o *RuntimeOperationEventEnvelopeEventPayload) GetProjectionRevision() string`

GetProjectionRevision returns the ProjectionRevision field if non-nil, zero value otherwise.

### GetProjectionRevisionOk

`func (o *RuntimeOperationEventEnvelopeEventPayload) GetProjectionRevisionOk() (*string, bool)`

GetProjectionRevisionOk returns a tuple with the ProjectionRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectionRevision

`func (o *RuntimeOperationEventEnvelopeEventPayload) SetProjectionRevision(v string)`

SetProjectionRevision sets ProjectionRevision field to given value.

### HasProjectionRevision

`func (o *RuntimeOperationEventEnvelopeEventPayload) HasProjectionRevision() bool`

HasProjectionRevision returns a boolean if a field has been set.

### GetRevision

`func (o *RuntimeOperationEventEnvelopeEventPayload) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *RuntimeOperationEventEnvelopeEventPayload) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *RuntimeOperationEventEnvelopeEventPayload) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *RuntimeOperationEventEnvelopeEventPayload) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetPlatformAgentId

`func (o *RuntimeOperationEventEnvelopeEventPayload) GetPlatformAgentId() string`

GetPlatformAgentId returns the PlatformAgentId field if non-nil, zero value otherwise.

### GetPlatformAgentIdOk

`func (o *RuntimeOperationEventEnvelopeEventPayload) GetPlatformAgentIdOk() (*string, bool)`

GetPlatformAgentIdOk returns a tuple with the PlatformAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformAgentId

`func (o *RuntimeOperationEventEnvelopeEventPayload) SetPlatformAgentId(v string)`

SetPlatformAgentId sets PlatformAgentId field to given value.

### HasPlatformAgentId

`func (o *RuntimeOperationEventEnvelopeEventPayload) HasPlatformAgentId() bool`

HasPlatformAgentId returns a boolean if a field has been set.

### GetErrorCode

`func (o *RuntimeOperationEventEnvelopeEventPayload) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *RuntimeOperationEventEnvelopeEventPayload) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *RuntimeOperationEventEnvelopeEventPayload) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *RuntimeOperationEventEnvelopeEventPayload) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetCauseCode

`func (o *RuntimeOperationEventEnvelopeEventPayload) GetCauseCode() string`

GetCauseCode returns the CauseCode field if non-nil, zero value otherwise.

### GetCauseCodeOk

`func (o *RuntimeOperationEventEnvelopeEventPayload) GetCauseCodeOk() (*string, bool)`

GetCauseCodeOk returns a tuple with the CauseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCauseCode

`func (o *RuntimeOperationEventEnvelopeEventPayload) SetCauseCode(v string)`

SetCauseCode sets CauseCode field to given value.

### HasCauseCode

`func (o *RuntimeOperationEventEnvelopeEventPayload) HasCauseCode() bool`

HasCauseCode returns a boolean if a field has been set.

### GetRequestedBy

`func (o *RuntimeOperationEventEnvelopeEventPayload) GetRequestedBy() string`

GetRequestedBy returns the RequestedBy field if non-nil, zero value otherwise.

### GetRequestedByOk

`func (o *RuntimeOperationEventEnvelopeEventPayload) GetRequestedByOk() (*string, bool)`

GetRequestedByOk returns a tuple with the RequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedBy

`func (o *RuntimeOperationEventEnvelopeEventPayload) SetRequestedBy(v string)`

SetRequestedBy sets RequestedBy field to given value.

### HasRequestedBy

`func (o *RuntimeOperationEventEnvelopeEventPayload) HasRequestedBy() bool`

HasRequestedBy returns a boolean if a field has been set.

### GetReason

`func (o *RuntimeOperationEventEnvelopeEventPayload) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *RuntimeOperationEventEnvelopeEventPayload) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *RuntimeOperationEventEnvelopeEventPayload) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *RuntimeOperationEventEnvelopeEventPayload) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


