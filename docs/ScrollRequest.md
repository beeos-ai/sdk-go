# ScrollRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**X** | Pointer to **int32** |  | [optional] 
**Y** | Pointer to **int32** |  | [optional] 
**Direction** | **string** | Scroll direction (&#x60;up&#x60; / &#x60;down&#x60; / &#x60;left&#x60; / &#x60;right&#x60;). | 
**Amount** | Pointer to **int32** | Distance multiplier (~300px per unit). Default 1. | [optional] 

## Methods

### NewScrollRequest

`func NewScrollRequest(direction string, ) *ScrollRequest`

NewScrollRequest instantiates a new ScrollRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScrollRequestWithDefaults

`func NewScrollRequestWithDefaults() *ScrollRequest`

NewScrollRequestWithDefaults instantiates a new ScrollRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetX

`func (o *ScrollRequest) GetX() int32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *ScrollRequest) GetXOk() (*int32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *ScrollRequest) SetX(v int32)`

SetX sets X field to given value.

### HasX

`func (o *ScrollRequest) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *ScrollRequest) GetY() int32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *ScrollRequest) GetYOk() (*int32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *ScrollRequest) SetY(v int32)`

SetY sets Y field to given value.

### HasY

`func (o *ScrollRequest) HasY() bool`

HasY returns a boolean if a field has been set.

### GetDirection

`func (o *ScrollRequest) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *ScrollRequest) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *ScrollRequest) SetDirection(v string)`

SetDirection sets Direction field to given value.


### GetAmount

`func (o *ScrollRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ScrollRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ScrollRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ScrollRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


