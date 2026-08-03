# ClickRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**X** | **int32** |  | 
**Y** | **int32** |  | 
**Button** | Pointer to **string** | Mouse button (&#x60;left&#x60; / &#x60;right&#x60; / &#x60;middle&#x60;). Default &#x60;left&#x60;. | [optional] 
**Count** | Pointer to **int32** | Click count (1 &#x3D; single, 2 &#x3D; double). Default 1. | [optional] 

## Methods

### NewClickRequest

`func NewClickRequest(x int32, y int32, ) *ClickRequest`

NewClickRequest instantiates a new ClickRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClickRequestWithDefaults

`func NewClickRequestWithDefaults() *ClickRequest`

NewClickRequestWithDefaults instantiates a new ClickRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetX

`func (o *ClickRequest) GetX() int32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *ClickRequest) GetXOk() (*int32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *ClickRequest) SetX(v int32)`

SetX sets X field to given value.


### GetY

`func (o *ClickRequest) GetY() int32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *ClickRequest) GetYOk() (*int32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *ClickRequest) SetY(v int32)`

SetY sets Y field to given value.


### GetButton

`func (o *ClickRequest) GetButton() string`

GetButton returns the Button field if non-nil, zero value otherwise.

### GetButtonOk

`func (o *ClickRequest) GetButtonOk() (*string, bool)`

GetButtonOk returns a tuple with the Button field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButton

`func (o *ClickRequest) SetButton(v string)`

SetButton sets Button field to given value.

### HasButton

`func (o *ClickRequest) HasButton() bool`

HasButton returns a boolean if a field has been set.

### GetCount

`func (o *ClickRequest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ClickRequest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ClickRequest) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ClickRequest) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


