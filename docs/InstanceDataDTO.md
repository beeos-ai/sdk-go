# InstanceDataDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**OwnerId** | **string** |  | 
**AgentFramework** | **string** |  | 
**ProviderId** | **string** |  | 
**ExternalId** | **string** |  | 
**OsType** | **string** |  | 
**HostingType** | **string** |  | 
**CloudProvider** | **string** |  | 
**Name** | **string** |  | 
**ModelPrimary** | **string** |  | 
**Models** | **[]string** |  | 
**Status** | **string** | One of &#x60;pending&#x60; | &#x60;provisioning&#x60; | &#x60;running&#x60; | &#x60;stopped&#x60; | &#x60;terminated&#x60;.  | 
**DesiredStatus** | **string** | One of &#x60;running&#x60; | &#x60;stopped&#x60; | &#x60;terminated&#x60;.  | 
**Endpoint** | **string** |  | 
**BridgeId** | **string** |  | 
**IdentityId** | **string** |  | 
**TotalRunSeconds** | **int64** |  | 
**Connectivity** | **string** | One of &#x60;unknown&#x60; | &#x60;online&#x60; | &#x60;offline&#x60;. | 
**ConnectivityReason** | Pointer to **string** | Open-string reason annotating a non-online &#x60;connectivity&#x60; value (independent of &#x60;haltReason&#x60;). Canonical values: &#x60;device_disconnected&#x60; | &#x60;device_unauthorized&#x60; | &#x60;instance_disconnected&#x60; (legacy alias &#x60;agent_disconnected&#x60; still accepted during Phase 1 rollout). Empty when connectivity is online. Lifecycle is tied to connectivity: any transition into &#x60;online&#x60; strictly clears this field.  | [optional] 
**HaltReason** | Pointer to **string** |  | [optional] 
**ProviderConfig** | **map[string]interface{}** | Sanitized provider config (apiKeys stripped). Raw JSON object.  | 
**PublicIp** | Pointer to **string** |  | [optional] 
**ClusterId** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**ImageId** | Pointer to **string** |  | [optional] 
**ImageRef** | Pointer to **string** |  | [optional] 
**IsTrial** | Pointer to **bool** |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**SystemPrompt** | Pointer to **string** |  | [optional] 
**McpServers** | Pointer to **map[string]interface{}** |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**StoppedAt** | Pointer to **time.Time** |  | [optional] 
**StatusEnteredAt** | Pointer to **time.Time** |  | [optional] 
**ConnectivityUpdatedAt** | Pointer to **time.Time** |  | [optional] 
**MsConnectionStatus** | Pointer to **string** | Independent MessageClient/Centrifugo connection reported by the current Runtime. One of &#x60;unknown&#x60; | &#x60;connecting&#x60; | &#x60;connected&#x60; | &#x60;reconnecting&#x60; | &#x60;disconnected&#x60; | &#x60;failed&#x60;. Cluster returns &#x60;unknown&#x60; when the Runtime lease is missing, expired, or the instance is stopped/terminated. This does not replace &#x60;connectivity&#x60;.  | [optional] 
**MsConnectionUpdatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**ScreenshotUrl** | Pointer to **string** |  | [optional] 
**ScreenshotUpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Capabilities** | Pointer to **map[string]interface{}** | Frozen capability declaration snapshot (raw JSON object, nullable). Byte-for-byte passthrough of the stored &#x60;capabilities_json&#x60; (proto &#x60;InstanceData&#x60; field 98). Omitted entirely when the instance has no declared capabilities (pre-migration rows / no matching billing template) rather than emitted as &#x60;null&#x60;. Mirrors the same field on the main Gateway DTO (services/gateway/internal/dto/instance_data.go).  Note: this is the **instance-level** capability JSON object and is semantically distinct from &#x60;DeviceInfoDTO.capabilities&#x60; (a string array of device hardware / Agent-reported capabilities returned by &#x60;GET /api/v1/instances/{id}/device/info&#x60;).  | [optional] 

## Methods

### NewInstanceDataDTO

`func NewInstanceDataDTO(id string, ownerId string, agentFramework string, providerId string, externalId string, osType string, hostingType string, cloudProvider string, name string, modelPrimary string, models []string, status string, desiredStatus string, endpoint string, bridgeId string, identityId string, totalRunSeconds int64, connectivity string, providerConfig map[string]interface{}, ) *InstanceDataDTO`

NewInstanceDataDTO instantiates a new InstanceDataDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceDataDTOWithDefaults

`func NewInstanceDataDTOWithDefaults() *InstanceDataDTO`

NewInstanceDataDTOWithDefaults instantiates a new InstanceDataDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceDataDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceDataDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceDataDTO) SetId(v string)`

SetId sets Id field to given value.


### GetOwnerId

`func (o *InstanceDataDTO) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *InstanceDataDTO) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *InstanceDataDTO) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### GetAgentFramework

`func (o *InstanceDataDTO) GetAgentFramework() string`

GetAgentFramework returns the AgentFramework field if non-nil, zero value otherwise.

### GetAgentFrameworkOk

`func (o *InstanceDataDTO) GetAgentFrameworkOk() (*string, bool)`

GetAgentFrameworkOk returns a tuple with the AgentFramework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentFramework

`func (o *InstanceDataDTO) SetAgentFramework(v string)`

SetAgentFramework sets AgentFramework field to given value.


### GetProviderId

`func (o *InstanceDataDTO) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *InstanceDataDTO) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *InstanceDataDTO) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.


### GetExternalId

`func (o *InstanceDataDTO) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *InstanceDataDTO) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *InstanceDataDTO) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetOsType

`func (o *InstanceDataDTO) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *InstanceDataDTO) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *InstanceDataDTO) SetOsType(v string)`

SetOsType sets OsType field to given value.


### GetHostingType

`func (o *InstanceDataDTO) GetHostingType() string`

GetHostingType returns the HostingType field if non-nil, zero value otherwise.

### GetHostingTypeOk

`func (o *InstanceDataDTO) GetHostingTypeOk() (*string, bool)`

GetHostingTypeOk returns a tuple with the HostingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostingType

`func (o *InstanceDataDTO) SetHostingType(v string)`

SetHostingType sets HostingType field to given value.


### GetCloudProvider

`func (o *InstanceDataDTO) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *InstanceDataDTO) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *InstanceDataDTO) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetName

`func (o *InstanceDataDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceDataDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceDataDTO) SetName(v string)`

SetName sets Name field to given value.


### GetModelPrimary

`func (o *InstanceDataDTO) GetModelPrimary() string`

GetModelPrimary returns the ModelPrimary field if non-nil, zero value otherwise.

### GetModelPrimaryOk

`func (o *InstanceDataDTO) GetModelPrimaryOk() (*string, bool)`

GetModelPrimaryOk returns a tuple with the ModelPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelPrimary

`func (o *InstanceDataDTO) SetModelPrimary(v string)`

SetModelPrimary sets ModelPrimary field to given value.


### GetModels

`func (o *InstanceDataDTO) GetModels() []string`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *InstanceDataDTO) GetModelsOk() (*[]string, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *InstanceDataDTO) SetModels(v []string)`

SetModels sets Models field to given value.


### GetStatus

`func (o *InstanceDataDTO) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InstanceDataDTO) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InstanceDataDTO) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDesiredStatus

`func (o *InstanceDataDTO) GetDesiredStatus() string`

GetDesiredStatus returns the DesiredStatus field if non-nil, zero value otherwise.

### GetDesiredStatusOk

`func (o *InstanceDataDTO) GetDesiredStatusOk() (*string, bool)`

GetDesiredStatusOk returns a tuple with the DesiredStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredStatus

`func (o *InstanceDataDTO) SetDesiredStatus(v string)`

SetDesiredStatus sets DesiredStatus field to given value.


### GetEndpoint

`func (o *InstanceDataDTO) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *InstanceDataDTO) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *InstanceDataDTO) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetBridgeId

`func (o *InstanceDataDTO) GetBridgeId() string`

GetBridgeId returns the BridgeId field if non-nil, zero value otherwise.

### GetBridgeIdOk

`func (o *InstanceDataDTO) GetBridgeIdOk() (*string, bool)`

GetBridgeIdOk returns a tuple with the BridgeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeId

`func (o *InstanceDataDTO) SetBridgeId(v string)`

SetBridgeId sets BridgeId field to given value.


### GetIdentityId

`func (o *InstanceDataDTO) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *InstanceDataDTO) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *InstanceDataDTO) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.


### GetTotalRunSeconds

`func (o *InstanceDataDTO) GetTotalRunSeconds() int64`

GetTotalRunSeconds returns the TotalRunSeconds field if non-nil, zero value otherwise.

### GetTotalRunSecondsOk

`func (o *InstanceDataDTO) GetTotalRunSecondsOk() (*int64, bool)`

GetTotalRunSecondsOk returns a tuple with the TotalRunSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRunSeconds

`func (o *InstanceDataDTO) SetTotalRunSeconds(v int64)`

SetTotalRunSeconds sets TotalRunSeconds field to given value.


### GetConnectivity

`func (o *InstanceDataDTO) GetConnectivity() string`

GetConnectivity returns the Connectivity field if non-nil, zero value otherwise.

### GetConnectivityOk

`func (o *InstanceDataDTO) GetConnectivityOk() (*string, bool)`

GetConnectivityOk returns a tuple with the Connectivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectivity

`func (o *InstanceDataDTO) SetConnectivity(v string)`

SetConnectivity sets Connectivity field to given value.


### GetConnectivityReason

`func (o *InstanceDataDTO) GetConnectivityReason() string`

GetConnectivityReason returns the ConnectivityReason field if non-nil, zero value otherwise.

### GetConnectivityReasonOk

`func (o *InstanceDataDTO) GetConnectivityReasonOk() (*string, bool)`

GetConnectivityReasonOk returns a tuple with the ConnectivityReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectivityReason

`func (o *InstanceDataDTO) SetConnectivityReason(v string)`

SetConnectivityReason sets ConnectivityReason field to given value.

### HasConnectivityReason

`func (o *InstanceDataDTO) HasConnectivityReason() bool`

HasConnectivityReason returns a boolean if a field has been set.

### GetHaltReason

`func (o *InstanceDataDTO) GetHaltReason() string`

GetHaltReason returns the HaltReason field if non-nil, zero value otherwise.

### GetHaltReasonOk

`func (o *InstanceDataDTO) GetHaltReasonOk() (*string, bool)`

GetHaltReasonOk returns a tuple with the HaltReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaltReason

`func (o *InstanceDataDTO) SetHaltReason(v string)`

SetHaltReason sets HaltReason field to given value.

### HasHaltReason

`func (o *InstanceDataDTO) HasHaltReason() bool`

HasHaltReason returns a boolean if a field has been set.

### GetProviderConfig

`func (o *InstanceDataDTO) GetProviderConfig() map[string]interface{}`

GetProviderConfig returns the ProviderConfig field if non-nil, zero value otherwise.

### GetProviderConfigOk

`func (o *InstanceDataDTO) GetProviderConfigOk() (*map[string]interface{}, bool)`

GetProviderConfigOk returns a tuple with the ProviderConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderConfig

`func (o *InstanceDataDTO) SetProviderConfig(v map[string]interface{})`

SetProviderConfig sets ProviderConfig field to given value.


### GetPublicIp

`func (o *InstanceDataDTO) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *InstanceDataDTO) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *InstanceDataDTO) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *InstanceDataDTO) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetClusterId

`func (o *InstanceDataDTO) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *InstanceDataDTO) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *InstanceDataDTO) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *InstanceDataDTO) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetRegion

`func (o *InstanceDataDTO) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *InstanceDataDTO) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *InstanceDataDTO) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *InstanceDataDTO) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetImageId

`func (o *InstanceDataDTO) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *InstanceDataDTO) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *InstanceDataDTO) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *InstanceDataDTO) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetImageRef

`func (o *InstanceDataDTO) GetImageRef() string`

GetImageRef returns the ImageRef field if non-nil, zero value otherwise.

### GetImageRefOk

`func (o *InstanceDataDTO) GetImageRefOk() (*string, bool)`

GetImageRefOk returns a tuple with the ImageRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageRef

`func (o *InstanceDataDTO) SetImageRef(v string)`

SetImageRef sets ImageRef field to given value.

### HasImageRef

`func (o *InstanceDataDTO) HasImageRef() bool`

HasImageRef returns a boolean if a field has been set.

### GetIsTrial

`func (o *InstanceDataDTO) GetIsTrial() bool`

GetIsTrial returns the IsTrial field if non-nil, zero value otherwise.

### GetIsTrialOk

`func (o *InstanceDataDTO) GetIsTrialOk() (*bool, bool)`

GetIsTrialOk returns a tuple with the IsTrial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrial

`func (o *InstanceDataDTO) SetIsTrial(v bool)`

SetIsTrial sets IsTrial field to given value.

### HasIsTrial

`func (o *InstanceDataDTO) HasIsTrial() bool`

HasIsTrial returns a boolean if a field has been set.

### GetExpiresAt

`func (o *InstanceDataDTO) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *InstanceDataDTO) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *InstanceDataDTO) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *InstanceDataDTO) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetErrorMessage

`func (o *InstanceDataDTO) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *InstanceDataDTO) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *InstanceDataDTO) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *InstanceDataDTO) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetSystemPrompt

`func (o *InstanceDataDTO) GetSystemPrompt() string`

GetSystemPrompt returns the SystemPrompt field if non-nil, zero value otherwise.

### GetSystemPromptOk

`func (o *InstanceDataDTO) GetSystemPromptOk() (*string, bool)`

GetSystemPromptOk returns a tuple with the SystemPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemPrompt

`func (o *InstanceDataDTO) SetSystemPrompt(v string)`

SetSystemPrompt sets SystemPrompt field to given value.

### HasSystemPrompt

`func (o *InstanceDataDTO) HasSystemPrompt() bool`

HasSystemPrompt returns a boolean if a field has been set.

### GetMcpServers

`func (o *InstanceDataDTO) GetMcpServers() map[string]interface{}`

GetMcpServers returns the McpServers field if non-nil, zero value otherwise.

### GetMcpServersOk

`func (o *InstanceDataDTO) GetMcpServersOk() (*map[string]interface{}, bool)`

GetMcpServersOk returns a tuple with the McpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcpServers

`func (o *InstanceDataDTO) SetMcpServers(v map[string]interface{})`

SetMcpServers sets McpServers field to given value.

### HasMcpServers

`func (o *InstanceDataDTO) HasMcpServers() bool`

HasMcpServers returns a boolean if a field has been set.

### GetStartedAt

`func (o *InstanceDataDTO) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *InstanceDataDTO) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *InstanceDataDTO) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *InstanceDataDTO) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStoppedAt

`func (o *InstanceDataDTO) GetStoppedAt() time.Time`

GetStoppedAt returns the StoppedAt field if non-nil, zero value otherwise.

### GetStoppedAtOk

`func (o *InstanceDataDTO) GetStoppedAtOk() (*time.Time, bool)`

GetStoppedAtOk returns a tuple with the StoppedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoppedAt

`func (o *InstanceDataDTO) SetStoppedAt(v time.Time)`

SetStoppedAt sets StoppedAt field to given value.

### HasStoppedAt

`func (o *InstanceDataDTO) HasStoppedAt() bool`

HasStoppedAt returns a boolean if a field has been set.

### GetStatusEnteredAt

`func (o *InstanceDataDTO) GetStatusEnteredAt() time.Time`

GetStatusEnteredAt returns the StatusEnteredAt field if non-nil, zero value otherwise.

### GetStatusEnteredAtOk

`func (o *InstanceDataDTO) GetStatusEnteredAtOk() (*time.Time, bool)`

GetStatusEnteredAtOk returns a tuple with the StatusEnteredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusEnteredAt

`func (o *InstanceDataDTO) SetStatusEnteredAt(v time.Time)`

SetStatusEnteredAt sets StatusEnteredAt field to given value.

### HasStatusEnteredAt

`func (o *InstanceDataDTO) HasStatusEnteredAt() bool`

HasStatusEnteredAt returns a boolean if a field has been set.

### GetConnectivityUpdatedAt

`func (o *InstanceDataDTO) GetConnectivityUpdatedAt() time.Time`

GetConnectivityUpdatedAt returns the ConnectivityUpdatedAt field if non-nil, zero value otherwise.

### GetConnectivityUpdatedAtOk

`func (o *InstanceDataDTO) GetConnectivityUpdatedAtOk() (*time.Time, bool)`

GetConnectivityUpdatedAtOk returns a tuple with the ConnectivityUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectivityUpdatedAt

`func (o *InstanceDataDTO) SetConnectivityUpdatedAt(v time.Time)`

SetConnectivityUpdatedAt sets ConnectivityUpdatedAt field to given value.

### HasConnectivityUpdatedAt

`func (o *InstanceDataDTO) HasConnectivityUpdatedAt() bool`

HasConnectivityUpdatedAt returns a boolean if a field has been set.

### GetMsConnectionStatus

`func (o *InstanceDataDTO) GetMsConnectionStatus() string`

GetMsConnectionStatus returns the MsConnectionStatus field if non-nil, zero value otherwise.

### GetMsConnectionStatusOk

`func (o *InstanceDataDTO) GetMsConnectionStatusOk() (*string, bool)`

GetMsConnectionStatusOk returns a tuple with the MsConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsConnectionStatus

`func (o *InstanceDataDTO) SetMsConnectionStatus(v string)`

SetMsConnectionStatus sets MsConnectionStatus field to given value.

### HasMsConnectionStatus

`func (o *InstanceDataDTO) HasMsConnectionStatus() bool`

HasMsConnectionStatus returns a boolean if a field has been set.

### GetMsConnectionUpdatedAt

`func (o *InstanceDataDTO) GetMsConnectionUpdatedAt() time.Time`

GetMsConnectionUpdatedAt returns the MsConnectionUpdatedAt field if non-nil, zero value otherwise.

### GetMsConnectionUpdatedAtOk

`func (o *InstanceDataDTO) GetMsConnectionUpdatedAtOk() (*time.Time, bool)`

GetMsConnectionUpdatedAtOk returns a tuple with the MsConnectionUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsConnectionUpdatedAt

`func (o *InstanceDataDTO) SetMsConnectionUpdatedAt(v time.Time)`

SetMsConnectionUpdatedAt sets MsConnectionUpdatedAt field to given value.

### HasMsConnectionUpdatedAt

`func (o *InstanceDataDTO) HasMsConnectionUpdatedAt() bool`

HasMsConnectionUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InstanceDataDTO) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InstanceDataDTO) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InstanceDataDTO) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InstanceDataDTO) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *InstanceDataDTO) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InstanceDataDTO) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InstanceDataDTO) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *InstanceDataDTO) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetScreenshotUrl

`func (o *InstanceDataDTO) GetScreenshotUrl() string`

GetScreenshotUrl returns the ScreenshotUrl field if non-nil, zero value otherwise.

### GetScreenshotUrlOk

`func (o *InstanceDataDTO) GetScreenshotUrlOk() (*string, bool)`

GetScreenshotUrlOk returns a tuple with the ScreenshotUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenshotUrl

`func (o *InstanceDataDTO) SetScreenshotUrl(v string)`

SetScreenshotUrl sets ScreenshotUrl field to given value.

### HasScreenshotUrl

`func (o *InstanceDataDTO) HasScreenshotUrl() bool`

HasScreenshotUrl returns a boolean if a field has been set.

### GetScreenshotUpdatedAt

`func (o *InstanceDataDTO) GetScreenshotUpdatedAt() time.Time`

GetScreenshotUpdatedAt returns the ScreenshotUpdatedAt field if non-nil, zero value otherwise.

### GetScreenshotUpdatedAtOk

`func (o *InstanceDataDTO) GetScreenshotUpdatedAtOk() (*time.Time, bool)`

GetScreenshotUpdatedAtOk returns a tuple with the ScreenshotUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenshotUpdatedAt

`func (o *InstanceDataDTO) SetScreenshotUpdatedAt(v time.Time)`

SetScreenshotUpdatedAt sets ScreenshotUpdatedAt field to given value.

### HasScreenshotUpdatedAt

`func (o *InstanceDataDTO) HasScreenshotUpdatedAt() bool`

HasScreenshotUpdatedAt returns a boolean if a field has been set.

### GetCapabilities

`func (o *InstanceDataDTO) GetCapabilities() map[string]interface{}`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *InstanceDataDTO) GetCapabilitiesOk() (*map[string]interface{}, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *InstanceDataDTO) SetCapabilities(v map[string]interface{})`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *InstanceDataDTO) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


