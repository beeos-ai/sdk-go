# CreateCanvasSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlatformAgentId** | **string** |  | 
**ConversationId** | **string** |  | 
**CanvasId** | Pointer to **string** | Opaque Canvas/A2UI surface identifier; not a UUID | [optional] 

## Methods

### NewCreateCanvasSessionRequest

`func NewCreateCanvasSessionRequest(platformAgentId string, conversationId string, ) *CreateCanvasSessionRequest`

NewCreateCanvasSessionRequest instantiates a new CreateCanvasSessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCanvasSessionRequestWithDefaults

`func NewCreateCanvasSessionRequestWithDefaults() *CreateCanvasSessionRequest`

NewCreateCanvasSessionRequestWithDefaults instantiates a new CreateCanvasSessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatformAgentId

`func (o *CreateCanvasSessionRequest) GetPlatformAgentId() string`

GetPlatformAgentId returns the PlatformAgentId field if non-nil, zero value otherwise.

### GetPlatformAgentIdOk

`func (o *CreateCanvasSessionRequest) GetPlatformAgentIdOk() (*string, bool)`

GetPlatformAgentIdOk returns a tuple with the PlatformAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformAgentId

`func (o *CreateCanvasSessionRequest) SetPlatformAgentId(v string)`

SetPlatformAgentId sets PlatformAgentId field to given value.


### GetConversationId

`func (o *CreateCanvasSessionRequest) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *CreateCanvasSessionRequest) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *CreateCanvasSessionRequest) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.


### GetCanvasId

`func (o *CreateCanvasSessionRequest) GetCanvasId() string`

GetCanvasId returns the CanvasId field if non-nil, zero value otherwise.

### GetCanvasIdOk

`func (o *CreateCanvasSessionRequest) GetCanvasIdOk() (*string, bool)`

GetCanvasIdOk returns a tuple with the CanvasId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanvasId

`func (o *CreateCanvasSessionRequest) SetCanvasId(v string)`

SetCanvasId sets CanvasId field to given value.

### HasCanvasId

`func (o *CreateCanvasSessionRequest) HasCanvasId() bool`

HasCanvasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


