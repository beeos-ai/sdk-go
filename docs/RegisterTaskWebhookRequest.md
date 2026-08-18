# RegisterTaskWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | Absolute HTTPS endpoint to POST callbacks to. The receiver gets the protocol-specific payload (OpenAPI: a task event envelope; A2A: JSON-RPC StreamResponse; MCP: notification).  | 
**Token** | Pointer to **string** | Optional bearer token sent as &#x60;Authorization: Bearer &lt;token&gt;&#x60; on each callback so the receiver can verify the call.  | [optional] 
**Secret** | Pointer to **string** | P2-A — HMAC-SHA256 signing key. When set, the deliverer signs every callback body and emits &#x60;X-BeeOS-Signature: t&#x3D;&lt;unix&gt;,v1&#x3D;&lt;hex&gt;&#x60; where &#x60;hex &#x3D; hmac_sha256(secret, t || \&quot;.\&quot; || body)&#x60;. The receiver should re-run the same HMAC to authenticate; mismatched timestamps (±5 min recommended tolerance) should be rejected to defend against replay.  **Write-only**. Never echoed on Get / List responses; the &#x60;has_secret&#x60; boolean on &#x60;TaskWebhookResponse&#x60; is the only observable signal. Rotate by Set-ing a fresh value; passing &#x60;\&quot;\&quot;&#x60; on update preserves the existing secret (use DELETE + re-register to clear).  | [optional] 

## Methods

### NewRegisterTaskWebhookRequest

`func NewRegisterTaskWebhookRequest(url string, ) *RegisterTaskWebhookRequest`

NewRegisterTaskWebhookRequest instantiates a new RegisterTaskWebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterTaskWebhookRequestWithDefaults

`func NewRegisterTaskWebhookRequestWithDefaults() *RegisterTaskWebhookRequest`

NewRegisterTaskWebhookRequestWithDefaults instantiates a new RegisterTaskWebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *RegisterTaskWebhookRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RegisterTaskWebhookRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RegisterTaskWebhookRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetToken

`func (o *RegisterTaskWebhookRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RegisterTaskWebhookRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RegisterTaskWebhookRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *RegisterTaskWebhookRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetSecret

`func (o *RegisterTaskWebhookRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *RegisterTaskWebhookRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *RegisterTaskWebhookRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *RegisterTaskWebhookRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


