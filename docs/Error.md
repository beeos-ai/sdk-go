# Error

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | RFC 9457 problem-type slug. Most SDK consumers key on &#x60;code&#x60; rather than &#x60;type&#x60;; this field is kept for structural parity with the wider apierror catalogue.  | [optional] 
**Code** | **string** | Stable machine-readable error code. See [docs/reference/errors.md](https://github.com/beeos-ai/beeos/blob/main/docs/reference/errors.md) for the complete catalogue.  | 
**Message** | **string** | Human-readable explanation; presentation-only, never key on this. | 
**RequestId** | Pointer to **string** | Per-request correlation id propagated end-to-end (also returned in the &#x60;X-Request-Id&#x60; response header). Include this value when reporting errors so the server-side trace can be located.  | [optional] 
**Param** | Pointer to **string** | Field path of the offending parameter for validation-style errors (&#x60;code&#x60; in &#x60;missing_param&#x60; / &#x60;invalid_param&#x60;). Empty for errors that don&#39;t reference a specific field.  | [optional] 
**Metadata** | Pointer to **map[string]string** | Optional structured key-value hints for the client (retry strategy, related resource ids, etc.). Always a string-to-string map.  | [optional] 

## Methods

### NewError

`func NewError(code string, message string, ) *Error`

NewError instantiates a new Error object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorWithDefaults

`func NewErrorWithDefaults() *Error`

NewErrorWithDefaults instantiates a new Error object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Error) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Error) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Error) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Error) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCode

`func (o *Error) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Error) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Error) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *Error) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Error) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Error) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetRequestId

`func (o *Error) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *Error) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *Error) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *Error) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetParam

`func (o *Error) GetParam() string`

GetParam returns the Param field if non-nil, zero value otherwise.

### GetParamOk

`func (o *Error) GetParamOk() (*string, bool)`

GetParamOk returns a tuple with the Param field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParam

`func (o *Error) SetParam(v string)`

SetParam sets Param field to given value.

### HasParam

`func (o *Error) HasParam() bool`

HasParam returns a boolean if a field has been set.

### GetMetadata

`func (o *Error) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Error) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Error) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Error) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


