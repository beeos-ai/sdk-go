# TaskSSEMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Message type — see the TaskEventStream type table. | 
**MessageId** | Pointer to **string** |  | [optional] 
**Offset** | Pointer to **int64** | Strictly monotonic per-channel offset, used as the SSE resume cursor. **NOT guaranteed contiguous** — small holes can appear after producer-side storage failures (ADR-0022 §1.2). Clients use it via &#x60;offset &gt; since&#x60; only.  | [optional] 
**InReplyTo** | Pointer to **string** |  | [optional] 
**PublisherId** | Pointer to **string** |  | [optional] 
**Payload** | Pointer to **interface{}** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Server&#39;s last-write timestamp on this frame (RFC3339 with nanosecond precision, UTC). Equals &#x60;created_at&#x60; on the initial &#x60;message.created&#x60; frame and is bumped on every successful PATCH that produces a &#x60;message.updated&#x60; frame. SSE replay clients SHOULD use it to ack progress checkpoints; omitted only on legacy pre-v3 rows from before migration 006.  | [optional] 
**State** | Pointer to [**MessageState**](MessageState.md) | Message Envelope v3 lifecycle state. When present, SSE clients SHOULD use this as the primary terminal-detection signal: &#x60;streaming&#x60; keeps the stream open; the four terminal states (&#x60;completed&#x60;/&#x60;failed&#x60;/&#x60;refused&#x60;/&#x60;cancelled&#x60;) cause the gateway to emit &#x60;event: end&#x60; with &#x60;reason: task_terminal&#x60;. Omitted on legacy pre-v3 rows; consumers fall back to the legacy &#x60;type&#x60;-based classification (&#x60;agent_reply&#x60;, &#x60;agent_reply_error&#x60;, &#x60;agent.refuse&#x60;, &#x60;agent_busy&#x60;).  | [optional] 
**StopReason** | Pointer to [**StopReason**](StopReason.md) | Producer-supplied terminal reason, paired with a terminal &#x60;state&#x60; (omitted for &#x60;streaming&#x60;). Surfaces the same vocabulary as Message Envelope v3 (&#x60;end_turn&#x60;, &#x60;max_tokens&#x60;, &#x60;tool_use&#x60;, &#x60;user_stop&#x60;, &#x60;timeout&#x60;, &#x60;error&#x60;, &#x60;refused&#x60;, &#x60;content_filter&#x60;).  | [optional] 
**Body** | Pointer to **string** | Message Envelope v3 cumulative main-text column. On a &#x60;state&#x3D;streaming&#x60; frame, &#x60;body&#x60; is the latest snapshot the producer has PATCHed (monotonically growing). On a terminal frame, &#x60;body&#x60; is the assembled final text — clients SHOULD prefer it over the legacy &#x60;payload.text&#x60; field for v3 rows. Omitted on legacy pre-v3 rows.  | [optional] 
**Parts** | Pointer to [**[]Part**](Part.md) | Message Envelope v3 structured supplementary content (thinking / tool_use / tool_result / file / source / custom). Mutates with &#x60;body&#x60; during streaming; sealed once &#x60;state&#x60; becomes terminal. Omitted on legacy pre-v3 rows.  | [optional] 

## Methods

### NewTaskSSEMessage

`func NewTaskSSEMessage(type_ string, ) *TaskSSEMessage`

NewTaskSSEMessage instantiates a new TaskSSEMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskSSEMessageWithDefaults

`func NewTaskSSEMessageWithDefaults() *TaskSSEMessage`

NewTaskSSEMessageWithDefaults instantiates a new TaskSSEMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TaskSSEMessage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaskSSEMessage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaskSSEMessage) SetType(v string)`

SetType sets Type field to given value.


### GetMessageId

`func (o *TaskSSEMessage) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *TaskSSEMessage) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *TaskSSEMessage) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *TaskSSEMessage) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetOffset

`func (o *TaskSSEMessage) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *TaskSSEMessage) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *TaskSSEMessage) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *TaskSSEMessage) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetInReplyTo

`func (o *TaskSSEMessage) GetInReplyTo() string`

GetInReplyTo returns the InReplyTo field if non-nil, zero value otherwise.

### GetInReplyToOk

`func (o *TaskSSEMessage) GetInReplyToOk() (*string, bool)`

GetInReplyToOk returns a tuple with the InReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInReplyTo

`func (o *TaskSSEMessage) SetInReplyTo(v string)`

SetInReplyTo sets InReplyTo field to given value.

### HasInReplyTo

`func (o *TaskSSEMessage) HasInReplyTo() bool`

HasInReplyTo returns a boolean if a field has been set.

### GetPublisherId

`func (o *TaskSSEMessage) GetPublisherId() string`

GetPublisherId returns the PublisherId field if non-nil, zero value otherwise.

### GetPublisherIdOk

`func (o *TaskSSEMessage) GetPublisherIdOk() (*string, bool)`

GetPublisherIdOk returns a tuple with the PublisherId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisherId

`func (o *TaskSSEMessage) SetPublisherId(v string)`

SetPublisherId sets PublisherId field to given value.

### HasPublisherId

`func (o *TaskSSEMessage) HasPublisherId() bool`

HasPublisherId returns a boolean if a field has been set.

### GetPayload

`func (o *TaskSSEMessage) GetPayload() interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *TaskSSEMessage) GetPayloadOk() (*interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *TaskSSEMessage) SetPayload(v interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *TaskSSEMessage) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### SetPayloadNil

`func (o *TaskSSEMessage) SetPayloadNil(b bool)`

 SetPayloadNil sets the value for Payload to be an explicit nil

### UnsetPayload
`func (o *TaskSSEMessage) UnsetPayload()`

UnsetPayload ensures that no value is present for Payload, not even an explicit nil
### GetCreatedAt

`func (o *TaskSSEMessage) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TaskSSEMessage) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TaskSSEMessage) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TaskSSEMessage) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TaskSSEMessage) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TaskSSEMessage) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TaskSSEMessage) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TaskSSEMessage) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetState

`func (o *TaskSSEMessage) GetState() MessageState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TaskSSEMessage) GetStateOk() (*MessageState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TaskSSEMessage) SetState(v MessageState)`

SetState sets State field to given value.

### HasState

`func (o *TaskSSEMessage) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStopReason

`func (o *TaskSSEMessage) GetStopReason() StopReason`

GetStopReason returns the StopReason field if non-nil, zero value otherwise.

### GetStopReasonOk

`func (o *TaskSSEMessage) GetStopReasonOk() (*StopReason, bool)`

GetStopReasonOk returns a tuple with the StopReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopReason

`func (o *TaskSSEMessage) SetStopReason(v StopReason)`

SetStopReason sets StopReason field to given value.

### HasStopReason

`func (o *TaskSSEMessage) HasStopReason() bool`

HasStopReason returns a boolean if a field has been set.

### GetBody

`func (o *TaskSSEMessage) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *TaskSSEMessage) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *TaskSSEMessage) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *TaskSSEMessage) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetParts

`func (o *TaskSSEMessage) GetParts() []Part`

GetParts returns the Parts field if non-nil, zero value otherwise.

### GetPartsOk

`func (o *TaskSSEMessage) GetPartsOk() (*[]Part, bool)`

GetPartsOk returns a tuple with the Parts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParts

`func (o *TaskSSEMessage) SetParts(v []Part)`

SetParts sets Parts field to given value.

### HasParts

`func (o *TaskSSEMessage) HasParts() bool`

HasParts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


