# PartThinking

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Text** | **string** |  | 
**State** | Pointer to [**PartState**](PartState.md) |  | [optional] 

## Methods

### NewPartThinking

`func NewPartThinking(type_ string, text string, ) *PartThinking`

NewPartThinking instantiates a new PartThinking object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartThinkingWithDefaults

`func NewPartThinkingWithDefaults() *PartThinking`

NewPartThinkingWithDefaults instantiates a new PartThinking object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PartThinking) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PartThinking) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PartThinking) SetType(v string)`

SetType sets Type field to given value.


### GetText

`func (o *PartThinking) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *PartThinking) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *PartThinking) SetText(v string)`

SetText sets Text field to given value.


### GetState

`func (o *PartThinking) GetState() PartState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PartThinking) GetStateOk() (*PartState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PartThinking) SetState(v PartState)`

SetState sets State field to given value.

### HasState

`func (o *PartThinking) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


