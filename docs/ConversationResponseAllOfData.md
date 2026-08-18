# ConversationResponseAllOfData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConversationId** | **string** |  | 
**AgentId** | **string** |  | 
**Title** | Pointer to **string** |  | [optional] 
**State** | **string** |  | 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**LastActivityAt** | Pointer to **time.Time** |  | [optional] 
**ClosedAt** | Pointer to **time.Time** |  | [optional] 
**ModelOverrideId** | Pointer to **string** | Session model override projected by Message Service after &#x60;PUT …/conversations/{id}/model&#x60; (session/set_model). Absent when the conversation uses the agent/instance default.  | [optional] 
**EffectiveModelId** | Pointer to **string** | Resolved effective model when Core/MS projects one alongside the override. Clients may fall back to modelOverrideId when this field is absent.  | [optional] 

## Methods

### NewConversationResponseAllOfData

`func NewConversationResponseAllOfData(conversationId string, agentId string, state string, ) *ConversationResponseAllOfData`

NewConversationResponseAllOfData instantiates a new ConversationResponseAllOfData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversationResponseAllOfDataWithDefaults

`func NewConversationResponseAllOfDataWithDefaults() *ConversationResponseAllOfData`

NewConversationResponseAllOfDataWithDefaults instantiates a new ConversationResponseAllOfData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConversationId

`func (o *ConversationResponseAllOfData) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *ConversationResponseAllOfData) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *ConversationResponseAllOfData) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.


### GetAgentId

`func (o *ConversationResponseAllOfData) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *ConversationResponseAllOfData) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *ConversationResponseAllOfData) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.


### GetTitle

`func (o *ConversationResponseAllOfData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ConversationResponseAllOfData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ConversationResponseAllOfData) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ConversationResponseAllOfData) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetState

`func (o *ConversationResponseAllOfData) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ConversationResponseAllOfData) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ConversationResponseAllOfData) SetState(v string)`

SetState sets State field to given value.


### GetMetadata

`func (o *ConversationResponseAllOfData) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ConversationResponseAllOfData) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ConversationResponseAllOfData) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ConversationResponseAllOfData) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ConversationResponseAllOfData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ConversationResponseAllOfData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ConversationResponseAllOfData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ConversationResponseAllOfData) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastActivityAt

`func (o *ConversationResponseAllOfData) GetLastActivityAt() time.Time`

GetLastActivityAt returns the LastActivityAt field if non-nil, zero value otherwise.

### GetLastActivityAtOk

`func (o *ConversationResponseAllOfData) GetLastActivityAtOk() (*time.Time, bool)`

GetLastActivityAtOk returns a tuple with the LastActivityAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivityAt

`func (o *ConversationResponseAllOfData) SetLastActivityAt(v time.Time)`

SetLastActivityAt sets LastActivityAt field to given value.

### HasLastActivityAt

`func (o *ConversationResponseAllOfData) HasLastActivityAt() bool`

HasLastActivityAt returns a boolean if a field has been set.

### GetClosedAt

`func (o *ConversationResponseAllOfData) GetClosedAt() time.Time`

GetClosedAt returns the ClosedAt field if non-nil, zero value otherwise.

### GetClosedAtOk

`func (o *ConversationResponseAllOfData) GetClosedAtOk() (*time.Time, bool)`

GetClosedAtOk returns a tuple with the ClosedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedAt

`func (o *ConversationResponseAllOfData) SetClosedAt(v time.Time)`

SetClosedAt sets ClosedAt field to given value.

### HasClosedAt

`func (o *ConversationResponseAllOfData) HasClosedAt() bool`

HasClosedAt returns a boolean if a field has been set.

### GetModelOverrideId

`func (o *ConversationResponseAllOfData) GetModelOverrideId() string`

GetModelOverrideId returns the ModelOverrideId field if non-nil, zero value otherwise.

### GetModelOverrideIdOk

`func (o *ConversationResponseAllOfData) GetModelOverrideIdOk() (*string, bool)`

GetModelOverrideIdOk returns a tuple with the ModelOverrideId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelOverrideId

`func (o *ConversationResponseAllOfData) SetModelOverrideId(v string)`

SetModelOverrideId sets ModelOverrideId field to given value.

### HasModelOverrideId

`func (o *ConversationResponseAllOfData) HasModelOverrideId() bool`

HasModelOverrideId returns a boolean if a field has been set.

### GetEffectiveModelId

`func (o *ConversationResponseAllOfData) GetEffectiveModelId() string`

GetEffectiveModelId returns the EffectiveModelId field if non-nil, zero value otherwise.

### GetEffectiveModelIdOk

`func (o *ConversationResponseAllOfData) GetEffectiveModelIdOk() (*string, bool)`

GetEffectiveModelIdOk returns a tuple with the EffectiveModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveModelId

`func (o *ConversationResponseAllOfData) SetEffectiveModelId(v string)`

SetEffectiveModelId sets EffectiveModelId field to given value.

### HasEffectiveModelId

`func (o *ConversationResponseAllOfData) HasEffectiveModelId() bool`

HasEffectiveModelId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


