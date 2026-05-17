# InvokeAgentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | The message to send to the agent. | 
**ContextId** | Pointer to **string** | Optional conversation context ID for multi-turn conversations. | [optional] 
**TimeoutMs** | Pointer to **int64** | Timeout in milliseconds (default 120000). | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


