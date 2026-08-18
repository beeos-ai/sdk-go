# MessageEnvelope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Stable v3 message id (&#x3D;&#x3D; Idempotency-Key on POST). | 
**ConversationId** | **string** |  | 
**Type** | **string** | Message intent. &#x60;agent_reply&#x60; for terminal model output; &#x60;chat_message&#x60; for inbound user/agent requests; custom strings for application-specific intents. | 
**Sender** | **string** | Publisher principal id (agent / user / service). | 
**ReplyTo** | Pointer to **string** | Message id this envelope replies to. Required for request/reply matching via the Message Service &#x60;wait&#x60; primitive.  | [optional] 
**Body** | **string** | Cumulative main text. Empty string on parts-only envelopes. | 
**Parts** | Pointer to [**[]Part**](Part.md) |  | [optional] 
**State** | [**MessageState**](MessageState.md) |  | 
**StopReason** | Pointer to [**StopReason**](StopReason.md) |  | [optional] 
**Content** | Pointer to **interface{}** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | Pointer to **time.Time** | Server&#39;s last-write timestamp. Bumped on every successful PATCH (streaming or terminal). Equals &#x60;created_at&#x60; after the initial POST; diverges as subsequent PATCHes accumulate (per-token streaming snapshots and the final terminal transition). Use this for SSE replay checkpoints and client-side freshness decisions.  | [optional] 

## Methods

### NewMessageEnvelope

`func NewMessageEnvelope(id string, conversationId string, type_ string, sender string, body string, state MessageState, createdAt time.Time, ) *MessageEnvelope`

NewMessageEnvelope instantiates a new MessageEnvelope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageEnvelopeWithDefaults

`func NewMessageEnvelopeWithDefaults() *MessageEnvelope`

NewMessageEnvelopeWithDefaults instantiates a new MessageEnvelope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MessageEnvelope) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessageEnvelope) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessageEnvelope) SetId(v string)`

SetId sets Id field to given value.


### GetConversationId

`func (o *MessageEnvelope) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *MessageEnvelope) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *MessageEnvelope) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.


### GetType

`func (o *MessageEnvelope) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MessageEnvelope) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MessageEnvelope) SetType(v string)`

SetType sets Type field to given value.


### GetSender

`func (o *MessageEnvelope) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *MessageEnvelope) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *MessageEnvelope) SetSender(v string)`

SetSender sets Sender field to given value.


### GetReplyTo

`func (o *MessageEnvelope) GetReplyTo() string`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *MessageEnvelope) GetReplyToOk() (*string, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *MessageEnvelope) SetReplyTo(v string)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *MessageEnvelope) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetBody

`func (o *MessageEnvelope) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *MessageEnvelope) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *MessageEnvelope) SetBody(v string)`

SetBody sets Body field to given value.


### GetParts

`func (o *MessageEnvelope) GetParts() []Part`

GetParts returns the Parts field if non-nil, zero value otherwise.

### GetPartsOk

`func (o *MessageEnvelope) GetPartsOk() (*[]Part, bool)`

GetPartsOk returns a tuple with the Parts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParts

`func (o *MessageEnvelope) SetParts(v []Part)`

SetParts sets Parts field to given value.

### HasParts

`func (o *MessageEnvelope) HasParts() bool`

HasParts returns a boolean if a field has been set.

### GetState

`func (o *MessageEnvelope) GetState() MessageState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MessageEnvelope) GetStateOk() (*MessageState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MessageEnvelope) SetState(v MessageState)`

SetState sets State field to given value.


### GetStopReason

`func (o *MessageEnvelope) GetStopReason() StopReason`

GetStopReason returns the StopReason field if non-nil, zero value otherwise.

### GetStopReasonOk

`func (o *MessageEnvelope) GetStopReasonOk() (*StopReason, bool)`

GetStopReasonOk returns a tuple with the StopReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopReason

`func (o *MessageEnvelope) SetStopReason(v StopReason)`

SetStopReason sets StopReason field to given value.

### HasStopReason

`func (o *MessageEnvelope) HasStopReason() bool`

HasStopReason returns a boolean if a field has been set.

### GetContent

`func (o *MessageEnvelope) GetContent() interface{}`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MessageEnvelope) GetContentOk() (*interface{}, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MessageEnvelope) SetContent(v interface{})`

SetContent sets Content field to given value.

### HasContent

`func (o *MessageEnvelope) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *MessageEnvelope) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *MessageEnvelope) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetCreatedAt

`func (o *MessageEnvelope) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MessageEnvelope) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MessageEnvelope) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *MessageEnvelope) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MessageEnvelope) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MessageEnvelope) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MessageEnvelope) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


