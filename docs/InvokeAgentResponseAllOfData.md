# InvokeAgentResponseAllOfData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** | Agent reply text. | 
**ContextId** | **string** | Conversation context ID (channel ID). | 
**IsError** | **bool** | Whether the agent&#39;s reply indicated an error. | 

## Methods

### NewInvokeAgentResponseAllOfData

`func NewInvokeAgentResponseAllOfData(text string, contextId string, isError bool, ) *InvokeAgentResponseAllOfData`

NewInvokeAgentResponseAllOfData instantiates a new InvokeAgentResponseAllOfData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvokeAgentResponseAllOfDataWithDefaults

`func NewInvokeAgentResponseAllOfDataWithDefaults() *InvokeAgentResponseAllOfData`

NewInvokeAgentResponseAllOfDataWithDefaults instantiates a new InvokeAgentResponseAllOfData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *InvokeAgentResponseAllOfData) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *InvokeAgentResponseAllOfData) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *InvokeAgentResponseAllOfData) SetText(v string)`

SetText sets Text field to given value.


### GetContextId

`func (o *InvokeAgentResponseAllOfData) GetContextId() string`

GetContextId returns the ContextId field if non-nil, zero value otherwise.

### GetContextIdOk

`func (o *InvokeAgentResponseAllOfData) GetContextIdOk() (*string, bool)`

GetContextIdOk returns a tuple with the ContextId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextId

`func (o *InvokeAgentResponseAllOfData) SetContextId(v string)`

SetContextId sets ContextId field to given value.


### GetIsError

`func (o *InvokeAgentResponseAllOfData) GetIsError() bool`

GetIsError returns the IsError field if non-nil, zero value otherwise.

### GetIsErrorOk

`func (o *InvokeAgentResponseAllOfData) GetIsErrorOk() (*bool, bool)`

GetIsErrorOk returns a tuple with the IsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsError

`func (o *InvokeAgentResponseAllOfData) SetIsError(v bool)`

SetIsError sets IsError field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


