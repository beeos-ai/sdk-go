# AgentDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**InstanceId** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 
**Slug** | Pointer to **string** |  | [optional] 
**McpEnabled** | **bool** |  | 
**DefaultModelId** | Pointer to **string** | Safe canonical provider/model selector with an optional tag. Runtime-local keys, credentials, endpoint configuration, and opaque high-entropy values are never exposed. | [optional] 
**Visibility** | **string** |  | 
**Capabilities** | [**AgentCapabilitiesDTO**](AgentCapabilitiesDTO.md) |  | 
**ConversationTransport** | **string** | Conversation/message transport selected by the agent directory. | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAgentDTO

`func NewAgentDTO(id string, instanceId string, name string, status string, mcpEnabled bool, visibility string, capabilities AgentCapabilitiesDTO, conversationTransport string, ) *AgentDTO`

NewAgentDTO instantiates a new AgentDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentDTOWithDefaults

`func NewAgentDTOWithDefaults() *AgentDTO`

NewAgentDTOWithDefaults instantiates a new AgentDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgentDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentDTO) SetId(v string)`

SetId sets Id field to given value.


### GetInstanceId

`func (o *AgentDTO) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *AgentDTO) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *AgentDTO) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetName

`func (o *AgentDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentDTO) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AgentDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AgentDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AgentDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AgentDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *AgentDTO) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AgentDTO) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AgentDTO) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AgentDTO) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetStatus

`func (o *AgentDTO) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AgentDTO) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AgentDTO) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSlug

`func (o *AgentDTO) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *AgentDTO) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *AgentDTO) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *AgentDTO) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetMcpEnabled

`func (o *AgentDTO) GetMcpEnabled() bool`

GetMcpEnabled returns the McpEnabled field if non-nil, zero value otherwise.

### GetMcpEnabledOk

`func (o *AgentDTO) GetMcpEnabledOk() (*bool, bool)`

GetMcpEnabledOk returns a tuple with the McpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcpEnabled

`func (o *AgentDTO) SetMcpEnabled(v bool)`

SetMcpEnabled sets McpEnabled field to given value.


### GetDefaultModelId

`func (o *AgentDTO) GetDefaultModelId() string`

GetDefaultModelId returns the DefaultModelId field if non-nil, zero value otherwise.

### GetDefaultModelIdOk

`func (o *AgentDTO) GetDefaultModelIdOk() (*string, bool)`

GetDefaultModelIdOk returns a tuple with the DefaultModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultModelId

`func (o *AgentDTO) SetDefaultModelId(v string)`

SetDefaultModelId sets DefaultModelId field to given value.

### HasDefaultModelId

`func (o *AgentDTO) HasDefaultModelId() bool`

HasDefaultModelId returns a boolean if a field has been set.

### GetVisibility

`func (o *AgentDTO) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *AgentDTO) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *AgentDTO) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.


### GetCapabilities

`func (o *AgentDTO) GetCapabilities() AgentCapabilitiesDTO`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *AgentDTO) GetCapabilitiesOk() (*AgentCapabilitiesDTO, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *AgentDTO) SetCapabilities(v AgentCapabilitiesDTO)`

SetCapabilities sets Capabilities field to given value.


### GetConversationTransport

`func (o *AgentDTO) GetConversationTransport() string`

GetConversationTransport returns the ConversationTransport field if non-nil, zero value otherwise.

### GetConversationTransportOk

`func (o *AgentDTO) GetConversationTransportOk() (*string, bool)`

GetConversationTransportOk returns a tuple with the ConversationTransport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationTransport

`func (o *AgentDTO) SetConversationTransport(v string)`

SetConversationTransport sets ConversationTransport field to given value.


### GetCreatedAt

`func (o *AgentDTO) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AgentDTO) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AgentDTO) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AgentDTO) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AgentDTO) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AgentDTO) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AgentDTO) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AgentDTO) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


