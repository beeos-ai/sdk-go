# RuntimeOperationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** |  | 
**Message** | **string** |  | 
**OperationId** | **string** |  | 
**RequestId** | Pointer to **string** |  | [optional] 
**EffectState** | **string** |  | 
**RetryMode** | **string** |  | 
**Context** | Pointer to [**RuntimeOperationErrorContext**](RuntimeOperationErrorContext.md) |  | [optional] 

## Methods

### NewRuntimeOperationError

`func NewRuntimeOperationError(code string, message string, operationId string, effectState string, retryMode string, ) *RuntimeOperationError`

NewRuntimeOperationError instantiates a new RuntimeOperationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuntimeOperationErrorWithDefaults

`func NewRuntimeOperationErrorWithDefaults() *RuntimeOperationError`

NewRuntimeOperationErrorWithDefaults instantiates a new RuntimeOperationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *RuntimeOperationError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *RuntimeOperationError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *RuntimeOperationError) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *RuntimeOperationError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RuntimeOperationError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RuntimeOperationError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetOperationId

`func (o *RuntimeOperationError) GetOperationId() string`

GetOperationId returns the OperationId field if non-nil, zero value otherwise.

### GetOperationIdOk

`func (o *RuntimeOperationError) GetOperationIdOk() (*string, bool)`

GetOperationIdOk returns a tuple with the OperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationId

`func (o *RuntimeOperationError) SetOperationId(v string)`

SetOperationId sets OperationId field to given value.


### GetRequestId

`func (o *RuntimeOperationError) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *RuntimeOperationError) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *RuntimeOperationError) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *RuntimeOperationError) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetEffectState

`func (o *RuntimeOperationError) GetEffectState() string`

GetEffectState returns the EffectState field if non-nil, zero value otherwise.

### GetEffectStateOk

`func (o *RuntimeOperationError) GetEffectStateOk() (*string, bool)`

GetEffectStateOk returns a tuple with the EffectState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectState

`func (o *RuntimeOperationError) SetEffectState(v string)`

SetEffectState sets EffectState field to given value.


### GetRetryMode

`func (o *RuntimeOperationError) GetRetryMode() string`

GetRetryMode returns the RetryMode field if non-nil, zero value otherwise.

### GetRetryModeOk

`func (o *RuntimeOperationError) GetRetryModeOk() (*string, bool)`

GetRetryModeOk returns a tuple with the RetryMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryMode

`func (o *RuntimeOperationError) SetRetryMode(v string)`

SetRetryMode sets RetryMode field to given value.


### GetContext

`func (o *RuntimeOperationError) GetContext() RuntimeOperationErrorContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *RuntimeOperationError) GetContextOk() (*RuntimeOperationErrorContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *RuntimeOperationError) SetContext(v RuntimeOperationErrorContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *RuntimeOperationError) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


