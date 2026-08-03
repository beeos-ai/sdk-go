# CreateTaskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Initial user prompt. Wrapped into a &#x60;chat_message&#x60; envelope. | 
**ContextId** | Pointer to **string** | Optional existing conversation/channel ID to reuse instead of creating a fresh task channel. Rare; most callers omit this.  | [optional] 
**DeadlineMs** | Pointer to **int64** | Task deadline in milliseconds. When &gt; 0, the channel auto-closes with &#x60;status&#x3D;timeout&#x60; if no terminal reply lands in this window. Max 7 days (server-side cap).  | [optional] 
**IdempotencyKey** | Pointer to **string** | Forwarded to MS&#39;s &#x60;channel_messages&#x60; UNIQUE index. Re-submitting with the same key returns the original task_id rather than spawning a duplicate. Recommended for retry-prone callers.  | [optional] 
**Metadata** | Pointer to **map[string]string** | Caller-supplied correlation tags (trace_id, business labels). Reserved chatinvoke routing keys (&#x60;protocol&#x60;, &#x60;caller_owner_id&#x60;, &#x60;target_agent_id&#x60;, &#x60;delivery_principal&#x60;) are overwritten by the gateway and stripped from the response.  | [optional] 
**Attachments** | Pointer to [**[]AttachmentRef**](AttachmentRef.md) | Optional list of files previously uploaded via [&#x60;POST /api/v1/files/presign-upload&#x60;](#operation/presignFileUpload). Resolved server-side and embedded in the initial chat_message envelope.  | [optional] 

## Methods

### NewCreateTaskRequest

`func NewCreateTaskRequest(message string, ) *CreateTaskRequest`

NewCreateTaskRequest instantiates a new CreateTaskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTaskRequestWithDefaults

`func NewCreateTaskRequestWithDefaults() *CreateTaskRequest`

NewCreateTaskRequestWithDefaults instantiates a new CreateTaskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *CreateTaskRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateTaskRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateTaskRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetContextId

`func (o *CreateTaskRequest) GetContextId() string`

GetContextId returns the ContextId field if non-nil, zero value otherwise.

### GetContextIdOk

`func (o *CreateTaskRequest) GetContextIdOk() (*string, bool)`

GetContextIdOk returns a tuple with the ContextId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextId

`func (o *CreateTaskRequest) SetContextId(v string)`

SetContextId sets ContextId field to given value.

### HasContextId

`func (o *CreateTaskRequest) HasContextId() bool`

HasContextId returns a boolean if a field has been set.

### GetDeadlineMs

`func (o *CreateTaskRequest) GetDeadlineMs() int64`

GetDeadlineMs returns the DeadlineMs field if non-nil, zero value otherwise.

### GetDeadlineMsOk

`func (o *CreateTaskRequest) GetDeadlineMsOk() (*int64, bool)`

GetDeadlineMsOk returns a tuple with the DeadlineMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadlineMs

`func (o *CreateTaskRequest) SetDeadlineMs(v int64)`

SetDeadlineMs sets DeadlineMs field to given value.

### HasDeadlineMs

`func (o *CreateTaskRequest) HasDeadlineMs() bool`

HasDeadlineMs returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *CreateTaskRequest) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *CreateTaskRequest) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *CreateTaskRequest) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *CreateTaskRequest) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateTaskRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateTaskRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateTaskRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateTaskRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetAttachments

`func (o *CreateTaskRequest) GetAttachments() []AttachmentRef`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *CreateTaskRequest) GetAttachmentsOk() (*[]AttachmentRef, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *CreateTaskRequest) SetAttachments(v []AttachmentRef)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *CreateTaskRequest) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


