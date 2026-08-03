# InvokeAgentSSEDone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Text** | Pointer to **string** | Concatenated full reply text (echoes the deltas). | [optional] 
**ContextId** | Pointer to **string** | Conversation context ID (channel ID). | [optional] 
**IsError** | Pointer to **bool** | Whether the reply terminated in an error. | [optional] 
**Error** | Pointer to **string** | Human-readable failure cause (only when &#x60;is_error&#x3D;true&#x60;). | [optional] 
**Code** | Pointer to **string** | Stable sentinel code (only when &#x60;is_error&#x3D;true&#x60;). Matches the &#x60;code&#x60; on the preceding &#x60;InvokeAgentSSEError&#x60; frame.  | [optional] 

## Methods

### NewInvokeAgentSSEDone

`func NewInvokeAgentSSEDone(type_ string, ) *InvokeAgentSSEDone`

NewInvokeAgentSSEDone instantiates a new InvokeAgentSSEDone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvokeAgentSSEDoneWithDefaults

`func NewInvokeAgentSSEDoneWithDefaults() *InvokeAgentSSEDone`

NewInvokeAgentSSEDoneWithDefaults instantiates a new InvokeAgentSSEDone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InvokeAgentSSEDone) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InvokeAgentSSEDone) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InvokeAgentSSEDone) SetType(v string)`

SetType sets Type field to given value.


### GetText

`func (o *InvokeAgentSSEDone) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *InvokeAgentSSEDone) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *InvokeAgentSSEDone) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *InvokeAgentSSEDone) HasText() bool`

HasText returns a boolean if a field has been set.

### GetContextId

`func (o *InvokeAgentSSEDone) GetContextId() string`

GetContextId returns the ContextId field if non-nil, zero value otherwise.

### GetContextIdOk

`func (o *InvokeAgentSSEDone) GetContextIdOk() (*string, bool)`

GetContextIdOk returns a tuple with the ContextId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextId

`func (o *InvokeAgentSSEDone) SetContextId(v string)`

SetContextId sets ContextId field to given value.

### HasContextId

`func (o *InvokeAgentSSEDone) HasContextId() bool`

HasContextId returns a boolean if a field has been set.

### GetIsError

`func (o *InvokeAgentSSEDone) GetIsError() bool`

GetIsError returns the IsError field if non-nil, zero value otherwise.

### GetIsErrorOk

`func (o *InvokeAgentSSEDone) GetIsErrorOk() (*bool, bool)`

GetIsErrorOk returns a tuple with the IsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsError

`func (o *InvokeAgentSSEDone) SetIsError(v bool)`

SetIsError sets IsError field to given value.

### HasIsError

`func (o *InvokeAgentSSEDone) HasIsError() bool`

HasIsError returns a boolean if a field has been set.

### GetError

`func (o *InvokeAgentSSEDone) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *InvokeAgentSSEDone) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *InvokeAgentSSEDone) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *InvokeAgentSSEDone) HasError() bool`

HasError returns a boolean if a field has been set.

### GetCode

`func (o *InvokeAgentSSEDone) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InvokeAgentSSEDone) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InvokeAgentSSEDone) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *InvokeAgentSSEDone) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


