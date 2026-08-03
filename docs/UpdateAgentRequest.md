# UpdateAgentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Visibility** | Pointer to **string** | Agent ACL bucket. &#x60;private&#x60; is owner-only; &#x60;public&#x60; lets any authenticated caller invoke; &#x60;unlisted&#x60; is invokable but absent from public listings. The &#x60;org&#x60; bucket is reserved for P1-C when the org-scope ACL lands.  | [optional] 
**McpEnabled** | Pointer to **bool** | Whether the agent shows up under MCP Gateway&#39;s &#x60;tools/list&#x60;. Independent of &#x60;visibility&#x60; — an &#x60;unlisted&#x60; agent can still be MCP-exposed; a &#x60;public&#x60; agent without MCP is invoke-only.  | [optional] 

## Methods

### NewUpdateAgentRequest

`func NewUpdateAgentRequest() *UpdateAgentRequest`

NewUpdateAgentRequest instantiates a new UpdateAgentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAgentRequestWithDefaults

`func NewUpdateAgentRequestWithDefaults() *UpdateAgentRequest`

NewUpdateAgentRequestWithDefaults instantiates a new UpdateAgentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVisibility

`func (o *UpdateAgentRequest) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateAgentRequest) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateAgentRequest) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateAgentRequest) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetMcpEnabled

`func (o *UpdateAgentRequest) GetMcpEnabled() bool`

GetMcpEnabled returns the McpEnabled field if non-nil, zero value otherwise.

### GetMcpEnabledOk

`func (o *UpdateAgentRequest) GetMcpEnabledOk() (*bool, bool)`

GetMcpEnabledOk returns a tuple with the McpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcpEnabled

`func (o *UpdateAgentRequest) SetMcpEnabled(v bool)`

SetMcpEnabled sets McpEnabled field to given value.

### HasMcpEnabled

`func (o *UpdateAgentRequest) HasMcpEnabled() bool`

HasMcpEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


