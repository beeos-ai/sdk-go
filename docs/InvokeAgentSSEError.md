# InvokeAgentSSEError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Code** | Pointer to **string** | Stable sentinel string code. Canonical values include &#x60;agent_not_found&#x60;, &#x60;agent_unavailable&#x60;, &#x60;agent_timeout&#x60;, &#x60;forbidden&#x60;, &#x60;invalid_param&#x60;, &#x60;internal&#x60;.  | [optional] 
**StatusCode** | Pointer to **int32** | HTTP-equivalent status code (so SDKs can branch on either field). | [optional] 
**Message** | **string** | Human-readable error detail. | 

## Methods

### NewInvokeAgentSSEError

`func NewInvokeAgentSSEError(type_ string, message string, ) *InvokeAgentSSEError`

NewInvokeAgentSSEError instantiates a new InvokeAgentSSEError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvokeAgentSSEErrorWithDefaults

`func NewInvokeAgentSSEErrorWithDefaults() *InvokeAgentSSEError`

NewInvokeAgentSSEErrorWithDefaults instantiates a new InvokeAgentSSEError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InvokeAgentSSEError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InvokeAgentSSEError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InvokeAgentSSEError) SetType(v string)`

SetType sets Type field to given value.


### GetCode

`func (o *InvokeAgentSSEError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InvokeAgentSSEError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InvokeAgentSSEError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *InvokeAgentSSEError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetStatusCode

`func (o *InvokeAgentSSEError) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *InvokeAgentSSEError) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *InvokeAgentSSEError) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *InvokeAgentSSEError) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetMessage

`func (o *InvokeAgentSSEError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InvokeAgentSSEError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InvokeAgentSSEError) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


