# TypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** |  | 
**Submit** | Pointer to **bool** | When true, press ENTER after typing. | [optional] 

## Methods

### NewTypeRequest

`func NewTypeRequest(text string, ) *TypeRequest`

NewTypeRequest instantiates a new TypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypeRequestWithDefaults

`func NewTypeRequestWithDefaults() *TypeRequest`

NewTypeRequestWithDefaults instantiates a new TypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *TypeRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *TypeRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *TypeRequest) SetText(v string)`

SetText sets Text field to given value.


### GetSubmit

`func (o *TypeRequest) GetSubmit() bool`

GetSubmit returns the Submit field if non-nil, zero value otherwise.

### GetSubmitOk

`func (o *TypeRequest) GetSubmitOk() (*bool, bool)`

GetSubmitOk returns a tuple with the Submit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmit

`func (o *TypeRequest) SetSubmit(v bool)`

SetSubmit sets Submit field to given value.

### HasSubmit

`func (o *TypeRequest) HasSubmit() bool`

HasSubmit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


