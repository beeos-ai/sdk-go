# RuntimeJSONRPCBusinessErrorError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **int32** |  | 
**Message** | **string** |  | 
**Data** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewRuntimeJSONRPCBusinessErrorError

`func NewRuntimeJSONRPCBusinessErrorError(code int32, message string, ) *RuntimeJSONRPCBusinessErrorError`

NewRuntimeJSONRPCBusinessErrorError instantiates a new RuntimeJSONRPCBusinessErrorError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuntimeJSONRPCBusinessErrorErrorWithDefaults

`func NewRuntimeJSONRPCBusinessErrorErrorWithDefaults() *RuntimeJSONRPCBusinessErrorError`

NewRuntimeJSONRPCBusinessErrorErrorWithDefaults instantiates a new RuntimeJSONRPCBusinessErrorError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *RuntimeJSONRPCBusinessErrorError) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *RuntimeJSONRPCBusinessErrorError) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *RuntimeJSONRPCBusinessErrorError) SetCode(v int32)`

SetCode sets Code field to given value.


### GetMessage

`func (o *RuntimeJSONRPCBusinessErrorError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RuntimeJSONRPCBusinessErrorError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RuntimeJSONRPCBusinessErrorError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *RuntimeJSONRPCBusinessErrorError) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RuntimeJSONRPCBusinessErrorError) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RuntimeJSONRPCBusinessErrorError) SetData(v interface{})`

SetData sets Data field to given value.

### HasData

`func (o *RuntimeJSONRPCBusinessErrorError) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *RuntimeJSONRPCBusinessErrorError) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *RuntimeJSONRPCBusinessErrorError) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


