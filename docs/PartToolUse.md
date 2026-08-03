# PartToolUse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** | Unique tool-call id; matches the &#x60;tool_use_id&#x60; on the corresponding tool_result. | 
**Name** | **string** |  | 
**Arguments** | Pointer to **interface{}** |  | [optional] 
**State** | Pointer to [**PartState**](PartState.md) |  | [optional] 

## Methods

### NewPartToolUse

`func NewPartToolUse(type_ string, id string, name string, ) *PartToolUse`

NewPartToolUse instantiates a new PartToolUse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartToolUseWithDefaults

`func NewPartToolUseWithDefaults() *PartToolUse`

NewPartToolUseWithDefaults instantiates a new PartToolUse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PartToolUse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PartToolUse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PartToolUse) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *PartToolUse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PartToolUse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PartToolUse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PartToolUse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PartToolUse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PartToolUse) SetName(v string)`

SetName sets Name field to given value.


### GetArguments

`func (o *PartToolUse) GetArguments() interface{}`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *PartToolUse) GetArgumentsOk() (*interface{}, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *PartToolUse) SetArguments(v interface{})`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *PartToolUse) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### SetArgumentsNil

`func (o *PartToolUse) SetArgumentsNil(b bool)`

 SetArgumentsNil sets the value for Arguments to be an explicit nil

### UnsetArguments
`func (o *PartToolUse) UnsetArguments()`

UnsetArguments ensures that no value is present for Arguments, not even an explicit nil
### GetState

`func (o *PartToolUse) GetState() PartState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PartToolUse) GetStateOk() (*PartState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PartToolUse) SetState(v PartState)`

SetState sets State field to given value.

### HasState

`func (o *PartToolUse) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


