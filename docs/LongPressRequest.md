# LongPressRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**X** | **int32** |  |
**Y** | **int32** |  |
**DurationMs** | Pointer to **int32** | Hold duration; clamped device-side to [1, 10000]. | [optional]

## Methods

### NewLongPressRequest

`func NewLongPressRequest(x int32, y int32, ) *LongPressRequest`

NewLongPressRequest instantiates a new LongPressRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLongPressRequestWithDefaults

`func NewLongPressRequestWithDefaults() *LongPressRequest`

NewLongPressRequestWithDefaults instantiates a new LongPressRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetX

`func (o *LongPressRequest) GetX() int32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *LongPressRequest) GetXOk() (*int32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *LongPressRequest) SetX(v int32)`

SetX sets X field to given value.


### GetY

`func (o *LongPressRequest) GetY() int32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *LongPressRequest) GetYOk() (*int32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *LongPressRequest) SetY(v int32)`

SetY sets Y field to given value.


### GetDurationMs

`func (o *LongPressRequest) GetDurationMs() int32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *LongPressRequest) GetDurationMsOk() (*int32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *LongPressRequest) SetDurationMs(v int32)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *LongPressRequest) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


