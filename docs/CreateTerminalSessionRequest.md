# CreateTerminalSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlatformAgentId** | **string** |  | 
**ConversationId** | Pointer to **string** |  | [optional] 
**ResumeTerminalId** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateTerminalSessionRequest

`func NewCreateTerminalSessionRequest(platformAgentId string, ) *CreateTerminalSessionRequest`

NewCreateTerminalSessionRequest instantiates a new CreateTerminalSessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTerminalSessionRequestWithDefaults

`func NewCreateTerminalSessionRequestWithDefaults() *CreateTerminalSessionRequest`

NewCreateTerminalSessionRequestWithDefaults instantiates a new CreateTerminalSessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatformAgentId

`func (o *CreateTerminalSessionRequest) GetPlatformAgentId() string`

GetPlatformAgentId returns the PlatformAgentId field if non-nil, zero value otherwise.

### GetPlatformAgentIdOk

`func (o *CreateTerminalSessionRequest) GetPlatformAgentIdOk() (*string, bool)`

GetPlatformAgentIdOk returns a tuple with the PlatformAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformAgentId

`func (o *CreateTerminalSessionRequest) SetPlatformAgentId(v string)`

SetPlatformAgentId sets PlatformAgentId field to given value.


### GetConversationId

`func (o *CreateTerminalSessionRequest) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *CreateTerminalSessionRequest) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *CreateTerminalSessionRequest) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.

### HasConversationId

`func (o *CreateTerminalSessionRequest) HasConversationId() bool`

HasConversationId returns a boolean if a field has been set.

### GetResumeTerminalId

`func (o *CreateTerminalSessionRequest) GetResumeTerminalId() string`

GetResumeTerminalId returns the ResumeTerminalId field if non-nil, zero value otherwise.

### GetResumeTerminalIdOk

`func (o *CreateTerminalSessionRequest) GetResumeTerminalIdOk() (*string, bool)`

GetResumeTerminalIdOk returns a tuple with the ResumeTerminalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeTerminalId

`func (o *CreateTerminalSessionRequest) SetResumeTerminalId(v string)`

SetResumeTerminalId sets ResumeTerminalId field to given value.

### HasResumeTerminalId

`func (o *CreateTerminalSessionRequest) HasResumeTerminalId() bool`

HasResumeTerminalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


