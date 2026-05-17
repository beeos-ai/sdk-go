# DeployInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VariantId** | Pointer to **string** |  | [optional] 
**AgentFramework** | Pointer to **string** |  | [optional] 
**ProviderId** | Pointer to **string** |  | [optional] 
**ProviderConfig** | Pointer to **map[string]interface{}** | Free-form provider-specific configuration merged with the catalog config (catalog wins on key conflicts).  | [optional] 
**Name** | **string** |  | 
**ModelPrimary** | Pointer to **string** |  | [optional] 
**Models** | Pointer to **[]string** |  | [optional] 
**SystemPrompt** | Pointer to **string** |  | [optional] 
**McpServers** | Pointer to **map[string]interface{}** |  | [optional] 
**BridgeId** | Pointer to **string** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** | Alternate name for &#x60;providerConfig&#x60;; same semantics. Used by the variant-driven path to layer user extras (apiKeys / apiBaseUrls) on top of the catalog config.  | [optional] 
**ClusterId** | Pointer to **string** |  | [optional] 
**PreferredRegion** | Pointer to **string** |  | [optional] 

## Methods

### NewDeployInstanceRequest

`func NewDeployInstanceRequest(name string, ) *DeployInstanceRequest`

NewDeployInstanceRequest instantiates a new DeployInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployInstanceRequestWithDefaults

`func NewDeployInstanceRequestWithDefaults() *DeployInstanceRequest`

NewDeployInstanceRequestWithDefaults instantiates a new DeployInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVariantId

`func (o *DeployInstanceRequest) GetVariantId() string`

GetVariantId returns the VariantId field if non-nil, zero value otherwise.

### GetVariantIdOk

`func (o *DeployInstanceRequest) GetVariantIdOk() (*string, bool)`

GetVariantIdOk returns a tuple with the VariantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantId

`func (o *DeployInstanceRequest) SetVariantId(v string)`

SetVariantId sets VariantId field to given value.

### HasVariantId

`func (o *DeployInstanceRequest) HasVariantId() bool`

HasVariantId returns a boolean if a field has been set.

### GetAgentFramework

`func (o *DeployInstanceRequest) GetAgentFramework() string`

GetAgentFramework returns the AgentFramework field if non-nil, zero value otherwise.

### GetAgentFrameworkOk

`func (o *DeployInstanceRequest) GetAgentFrameworkOk() (*string, bool)`

GetAgentFrameworkOk returns a tuple with the AgentFramework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentFramework

`func (o *DeployInstanceRequest) SetAgentFramework(v string)`

SetAgentFramework sets AgentFramework field to given value.

### HasAgentFramework

`func (o *DeployInstanceRequest) HasAgentFramework() bool`

HasAgentFramework returns a boolean if a field has been set.

### GetProviderId

`func (o *DeployInstanceRequest) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *DeployInstanceRequest) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *DeployInstanceRequest) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *DeployInstanceRequest) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetProviderConfig

`func (o *DeployInstanceRequest) GetProviderConfig() map[string]interface{}`

GetProviderConfig returns the ProviderConfig field if non-nil, zero value otherwise.

### GetProviderConfigOk

`func (o *DeployInstanceRequest) GetProviderConfigOk() (*map[string]interface{}, bool)`

GetProviderConfigOk returns a tuple with the ProviderConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderConfig

`func (o *DeployInstanceRequest) SetProviderConfig(v map[string]interface{})`

SetProviderConfig sets ProviderConfig field to given value.

### HasProviderConfig

`func (o *DeployInstanceRequest) HasProviderConfig() bool`

HasProviderConfig returns a boolean if a field has been set.

### GetName

`func (o *DeployInstanceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeployInstanceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeployInstanceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetModelPrimary

`func (o *DeployInstanceRequest) GetModelPrimary() string`

GetModelPrimary returns the ModelPrimary field if non-nil, zero value otherwise.

### GetModelPrimaryOk

`func (o *DeployInstanceRequest) GetModelPrimaryOk() (*string, bool)`

GetModelPrimaryOk returns a tuple with the ModelPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelPrimary

`func (o *DeployInstanceRequest) SetModelPrimary(v string)`

SetModelPrimary sets ModelPrimary field to given value.

### HasModelPrimary

`func (o *DeployInstanceRequest) HasModelPrimary() bool`

HasModelPrimary returns a boolean if a field has been set.

### GetModels

`func (o *DeployInstanceRequest) GetModels() []string`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *DeployInstanceRequest) GetModelsOk() (*[]string, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *DeployInstanceRequest) SetModels(v []string)`

SetModels sets Models field to given value.

### HasModels

`func (o *DeployInstanceRequest) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetSystemPrompt

`func (o *DeployInstanceRequest) GetSystemPrompt() string`

GetSystemPrompt returns the SystemPrompt field if non-nil, zero value otherwise.

### GetSystemPromptOk

`func (o *DeployInstanceRequest) GetSystemPromptOk() (*string, bool)`

GetSystemPromptOk returns a tuple with the SystemPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemPrompt

`func (o *DeployInstanceRequest) SetSystemPrompt(v string)`

SetSystemPrompt sets SystemPrompt field to given value.

### HasSystemPrompt

`func (o *DeployInstanceRequest) HasSystemPrompt() bool`

HasSystemPrompt returns a boolean if a field has been set.

### GetMcpServers

`func (o *DeployInstanceRequest) GetMcpServers() map[string]interface{}`

GetMcpServers returns the McpServers field if non-nil, zero value otherwise.

### GetMcpServersOk

`func (o *DeployInstanceRequest) GetMcpServersOk() (*map[string]interface{}, bool)`

GetMcpServersOk returns a tuple with the McpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcpServers

`func (o *DeployInstanceRequest) SetMcpServers(v map[string]interface{})`

SetMcpServers sets McpServers field to given value.

### HasMcpServers

`func (o *DeployInstanceRequest) HasMcpServers() bool`

HasMcpServers returns a boolean if a field has been set.

### GetBridgeId

`func (o *DeployInstanceRequest) GetBridgeId() string`

GetBridgeId returns the BridgeId field if non-nil, zero value otherwise.

### GetBridgeIdOk

`func (o *DeployInstanceRequest) GetBridgeIdOk() (*string, bool)`

GetBridgeIdOk returns a tuple with the BridgeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeId

`func (o *DeployInstanceRequest) SetBridgeId(v string)`

SetBridgeId sets BridgeId field to given value.

### HasBridgeId

`func (o *DeployInstanceRequest) HasBridgeId() bool`

HasBridgeId returns a boolean if a field has been set.

### GetConfig

`func (o *DeployInstanceRequest) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *DeployInstanceRequest) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *DeployInstanceRequest) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *DeployInstanceRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetClusterId

`func (o *DeployInstanceRequest) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *DeployInstanceRequest) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *DeployInstanceRequest) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *DeployInstanceRequest) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetPreferredRegion

`func (o *DeployInstanceRequest) GetPreferredRegion() string`

GetPreferredRegion returns the PreferredRegion field if non-nil, zero value otherwise.

### GetPreferredRegionOk

`func (o *DeployInstanceRequest) GetPreferredRegionOk() (*string, bool)`

GetPreferredRegionOk returns a tuple with the PreferredRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredRegion

`func (o *DeployInstanceRequest) SetPreferredRegion(v string)`

SetPreferredRegion sets PreferredRegion field to given value.

### HasPreferredRegion

`func (o *DeployInstanceRequest) HasPreferredRegion() bool`

HasPreferredRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


