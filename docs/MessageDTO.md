# MessageDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | **string** |  | 
**Offset** | **int64** |  | 
**Type** | **string** |  | 
**PublisherId** | Pointer to **string** |  | [optional] 
**InReplyTo** | Pointer to **string** |  | [optional] 
**Payload** | Pointer to **interface{}** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Body** | Pointer to **string** | Cumulative agent reply text. Mutated in place across PATCH updates while state &#x3D;&#x3D; \&quot;streaming\&quot;; frozen when state transitions to a terminal value (&#x60;completed&#x60; / &#x60;failed&#x60; / &#x60;refused&#x60; / &#x60;cancelled&#x60;).  | [optional] 
**State** | Pointer to [**MessageState**](MessageState.md) | Envelope lifecycle (canonical vocabulary, see MessageState). Pre-v3 rows omit this field; v3 consumers should treat absent state as &#x60;completed&#x60;.  | [optional] 
**StopReason** | Pointer to [**StopReason**](StopReason.md) | Terminal reason set when state becomes non-streaming (canonical vocabulary, see StopReason). Only meaningful when state ∈ {completed, failed, refused, cancelled}.  | [optional] 
**Parts** | Pointer to [**[]Part**](Part.md) | Structured part list for rich replies (thinking / tool_use / tool_result / file / source / custom). Empty / absent for plain text replies — consumers should render &#x60;body&#x60; only in that case. Discriminated union; see the Part schema for the per-type field set.  | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Server&#39;s last-write timestamp (RFC3339 with nanosecond precision, UTC). Equals &#x60;created_at&#x60; after the initial POST and is bumped on every successful PATCH (streaming snapshot updates AND terminal state transitions). Terminal rows are immutable, so &#x60;updated_at&#x60; freezes once &#x60;state&#x60; leaves &#x60;streaming&#x60;. SSE replay clients SHOULD use this to ack progress checkpoints; REST clients SHOULD use it to detect mid-stream progress vs &#x60;created_at&#x60;-only first reads.  | [optional] 

## Methods

### NewMessageDTO

`func NewMessageDTO(messageId string, offset int64, type_ string, ) *MessageDTO`

NewMessageDTO instantiates a new MessageDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageDTOWithDefaults

`func NewMessageDTOWithDefaults() *MessageDTO`

NewMessageDTOWithDefaults instantiates a new MessageDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageId

`func (o *MessageDTO) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *MessageDTO) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *MessageDTO) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.


### GetOffset

`func (o *MessageDTO) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *MessageDTO) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *MessageDTO) SetOffset(v int64)`

SetOffset sets Offset field to given value.


### GetType

`func (o *MessageDTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MessageDTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MessageDTO) SetType(v string)`

SetType sets Type field to given value.


### GetPublisherId

`func (o *MessageDTO) GetPublisherId() string`

GetPublisherId returns the PublisherId field if non-nil, zero value otherwise.

### GetPublisherIdOk

`func (o *MessageDTO) GetPublisherIdOk() (*string, bool)`

GetPublisherIdOk returns a tuple with the PublisherId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisherId

`func (o *MessageDTO) SetPublisherId(v string)`

SetPublisherId sets PublisherId field to given value.

### HasPublisherId

`func (o *MessageDTO) HasPublisherId() bool`

HasPublisherId returns a boolean if a field has been set.

### GetInReplyTo

`func (o *MessageDTO) GetInReplyTo() string`

GetInReplyTo returns the InReplyTo field if non-nil, zero value otherwise.

### GetInReplyToOk

`func (o *MessageDTO) GetInReplyToOk() (*string, bool)`

GetInReplyToOk returns a tuple with the InReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInReplyTo

`func (o *MessageDTO) SetInReplyTo(v string)`

SetInReplyTo sets InReplyTo field to given value.

### HasInReplyTo

`func (o *MessageDTO) HasInReplyTo() bool`

HasInReplyTo returns a boolean if a field has been set.

### GetPayload

`func (o *MessageDTO) GetPayload() interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *MessageDTO) GetPayloadOk() (*interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *MessageDTO) SetPayload(v interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *MessageDTO) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### SetPayloadNil

`func (o *MessageDTO) SetPayloadNil(b bool)`

 SetPayloadNil sets the value for Payload to be an explicit nil

### UnsetPayload
`func (o *MessageDTO) UnsetPayload()`

UnsetPayload ensures that no value is present for Payload, not even an explicit nil
### GetCreatedAt

`func (o *MessageDTO) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MessageDTO) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MessageDTO) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MessageDTO) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetBody

`func (o *MessageDTO) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *MessageDTO) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *MessageDTO) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *MessageDTO) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetState

`func (o *MessageDTO) GetState() MessageState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MessageDTO) GetStateOk() (*MessageState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MessageDTO) SetState(v MessageState)`

SetState sets State field to given value.

### HasState

`func (o *MessageDTO) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStopReason

`func (o *MessageDTO) GetStopReason() StopReason`

GetStopReason returns the StopReason field if non-nil, zero value otherwise.

### GetStopReasonOk

`func (o *MessageDTO) GetStopReasonOk() (*StopReason, bool)`

GetStopReasonOk returns a tuple with the StopReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopReason

`func (o *MessageDTO) SetStopReason(v StopReason)`

SetStopReason sets StopReason field to given value.

### HasStopReason

`func (o *MessageDTO) HasStopReason() bool`

HasStopReason returns a boolean if a field has been set.

### GetParts

`func (o *MessageDTO) GetParts() []Part`

GetParts returns the Parts field if non-nil, zero value otherwise.

### GetPartsOk

`func (o *MessageDTO) GetPartsOk() (*[]Part, bool)`

GetPartsOk returns a tuple with the Parts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParts

`func (o *MessageDTO) SetParts(v []Part)`

SetParts sets Parts field to given value.

### HasParts

`func (o *MessageDTO) HasParts() bool`

HasParts returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MessageDTO) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MessageDTO) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MessageDTO) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MessageDTO) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


