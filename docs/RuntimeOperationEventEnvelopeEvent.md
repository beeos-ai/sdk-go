# RuntimeOperationEventEnvelopeEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Sequence** | **string** |  | 
**RecordedAt** | **time.Time** |  | 
**Payload** | Pointer to [**RuntimeOperationEventEnvelopeEventPayload**](RuntimeOperationEventEnvelopeEventPayload.md) |  | [optional] 
**Outcome** | Pointer to **string** |  | [optional] 
**Result** | Pointer to **interface{}** |  | [optional] 
**Revision** | Pointer to **string** |  | [optional] 
**Error** | Pointer to [**RuntimeOperationError**](RuntimeOperationError.md) |  | [optional] 

## Methods

### NewRuntimeOperationEventEnvelopeEvent

`func NewRuntimeOperationEventEnvelopeEvent(type_ string, sequence string, recordedAt time.Time, ) *RuntimeOperationEventEnvelopeEvent`

NewRuntimeOperationEventEnvelopeEvent instantiates a new RuntimeOperationEventEnvelopeEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuntimeOperationEventEnvelopeEventWithDefaults

`func NewRuntimeOperationEventEnvelopeEventWithDefaults() *RuntimeOperationEventEnvelopeEvent`

NewRuntimeOperationEventEnvelopeEventWithDefaults instantiates a new RuntimeOperationEventEnvelopeEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RuntimeOperationEventEnvelopeEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RuntimeOperationEventEnvelopeEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RuntimeOperationEventEnvelopeEvent) SetType(v string)`

SetType sets Type field to given value.


### GetSequence

`func (o *RuntimeOperationEventEnvelopeEvent) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *RuntimeOperationEventEnvelopeEvent) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *RuntimeOperationEventEnvelopeEvent) SetSequence(v string)`

SetSequence sets Sequence field to given value.


### GetRecordedAt

`func (o *RuntimeOperationEventEnvelopeEvent) GetRecordedAt() time.Time`

GetRecordedAt returns the RecordedAt field if non-nil, zero value otherwise.

### GetRecordedAtOk

`func (o *RuntimeOperationEventEnvelopeEvent) GetRecordedAtOk() (*time.Time, bool)`

GetRecordedAtOk returns a tuple with the RecordedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordedAt

`func (o *RuntimeOperationEventEnvelopeEvent) SetRecordedAt(v time.Time)`

SetRecordedAt sets RecordedAt field to given value.


### GetPayload

`func (o *RuntimeOperationEventEnvelopeEvent) GetPayload() RuntimeOperationEventEnvelopeEventPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *RuntimeOperationEventEnvelopeEvent) GetPayloadOk() (*RuntimeOperationEventEnvelopeEventPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *RuntimeOperationEventEnvelopeEvent) SetPayload(v RuntimeOperationEventEnvelopeEventPayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *RuntimeOperationEventEnvelopeEvent) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetOutcome

`func (o *RuntimeOperationEventEnvelopeEvent) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *RuntimeOperationEventEnvelopeEvent) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *RuntimeOperationEventEnvelopeEvent) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *RuntimeOperationEventEnvelopeEvent) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### GetResult

`func (o *RuntimeOperationEventEnvelopeEvent) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *RuntimeOperationEventEnvelopeEvent) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *RuntimeOperationEventEnvelopeEvent) SetResult(v interface{})`

SetResult sets Result field to given value.

### HasResult

`func (o *RuntimeOperationEventEnvelopeEvent) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResultNil

`func (o *RuntimeOperationEventEnvelopeEvent) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *RuntimeOperationEventEnvelopeEvent) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetRevision

`func (o *RuntimeOperationEventEnvelopeEvent) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *RuntimeOperationEventEnvelopeEvent) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *RuntimeOperationEventEnvelopeEvent) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *RuntimeOperationEventEnvelopeEvent) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetError

`func (o *RuntimeOperationEventEnvelopeEvent) GetError() RuntimeOperationError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RuntimeOperationEventEnvelopeEvent) GetErrorOk() (*RuntimeOperationError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RuntimeOperationEventEnvelopeEvent) SetError(v RuntimeOperationError)`

SetError sets Error field to given value.

### HasError

`func (o *RuntimeOperationEventEnvelopeEvent) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


