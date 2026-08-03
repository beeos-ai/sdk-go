# InvokeAgentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | The message to send to the agent. | 
**ContextId** | Pointer to **string** | Optional conversation context ID for multi-turn conversations. | [optional] 
**TimeoutMs** | Pointer to **int64** | Per-request timeout in milliseconds. Default 120000 (2 minutes). Hard-capped at 115000 server-side so the response always has time to flush before the HTTP &#x60;WriteTimeout&#x60; (120s) fires. Values above 115000 are silently clamped; if the agent does not reply within the effective window the server returns &#x60;service_timeout&#x60; (HTTP 504) — see [docs/reference/errors.md](../../docs/reference/errors.md).  | [optional] 
**IdempotencyKey** | Pointer to **string** | Optional caller-generated idempotency key forwarded to Message Service as the chat_message&#39;s &#x60;idempotency_key&#x60;. Retries that race the same key dedup at MS (UNIQUE index on &#x60;channel_messages&#x60;). When omitted, the gateway generates a fresh UUID. The same key also doubles as the &#x60;message_id&#x60; an agent must echo back as &#x60;in_reply_to&#x60; on its reply.  | [optional] 
**Metadata** | Pointer to **map[string]string** | Opaque caller-controlled key/value pairs merged into the channel metadata (e.g. &#x60;trace_id&#x60;, &#x60;user_id&#x60;, business tags). Reserved routing keys (&#x60;protocol&#x60;, &#x60;caller_owner_id&#x60;, &#x60;target_agent_id&#x60;, &#x60;delivery_principal&#x60;) are silently scrubbed server-side and cannot be overridden by the caller.  | [optional] 
**Attachments** | Pointer to [**[]AttachmentRef**](AttachmentRef.md) | Optional list of files previously uploaded via [&#x60;POST /api/v1/files/presign-upload&#x60;](#operation/presignFileUpload). Each &#x60;file_id&#x60; is resolved server-side to a presigned download URL and embedded in the chat_message envelope so the receiving agent can fetch the bytes without further BeeOS auth.  | [optional] 

## Methods

### NewInvokeAgentRequest

`func NewInvokeAgentRequest(message string, ) *InvokeAgentRequest`

NewInvokeAgentRequest instantiates a new InvokeAgentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvokeAgentRequestWithDefaults

`func NewInvokeAgentRequestWithDefaults() *InvokeAgentRequest`

NewInvokeAgentRequestWithDefaults instantiates a new InvokeAgentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *InvokeAgentRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InvokeAgentRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InvokeAgentRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetContextId

`func (o *InvokeAgentRequest) GetContextId() string`

GetContextId returns the ContextId field if non-nil, zero value otherwise.

### GetContextIdOk

`func (o *InvokeAgentRequest) GetContextIdOk() (*string, bool)`

GetContextIdOk returns a tuple with the ContextId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextId

`func (o *InvokeAgentRequest) SetContextId(v string)`

SetContextId sets ContextId field to given value.

### HasContextId

`func (o *InvokeAgentRequest) HasContextId() bool`

HasContextId returns a boolean if a field has been set.

### GetTimeoutMs

`func (o *InvokeAgentRequest) GetTimeoutMs() int64`

GetTimeoutMs returns the TimeoutMs field if non-nil, zero value otherwise.

### GetTimeoutMsOk

`func (o *InvokeAgentRequest) GetTimeoutMsOk() (*int64, bool)`

GetTimeoutMsOk returns a tuple with the TimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutMs

`func (o *InvokeAgentRequest) SetTimeoutMs(v int64)`

SetTimeoutMs sets TimeoutMs field to given value.

### HasTimeoutMs

`func (o *InvokeAgentRequest) HasTimeoutMs() bool`

HasTimeoutMs returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *InvokeAgentRequest) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *InvokeAgentRequest) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *InvokeAgentRequest) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *InvokeAgentRequest) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### GetMetadata

`func (o *InvokeAgentRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InvokeAgentRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InvokeAgentRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *InvokeAgentRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetAttachments

`func (o *InvokeAgentRequest) GetAttachments() []AttachmentRef`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *InvokeAgentRequest) GetAttachmentsOk() (*[]AttachmentRef, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *InvokeAgentRequest) SetAttachments(v []AttachmentRef)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *InvokeAgentRequest) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


