# SwipeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**X1** | **int32** |  | 
**Y1** | **int32** |  | 
**X2** | **int32** |  | 
**Y2** | **int32** |  | 
**DurationMs** | Pointer to **int32** | Gesture duration; clamped device-side to [1, 10000]. | [optional] 

## Methods

### NewSwipeRequest

`func NewSwipeRequest(x1 int32, y1 int32, x2 int32, y2 int32, ) *SwipeRequest`

NewSwipeRequest instantiates a new SwipeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwipeRequestWithDefaults

`func NewSwipeRequestWithDefaults() *SwipeRequest`

NewSwipeRequestWithDefaults instantiates a new SwipeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetX1

`func (o *SwipeRequest) GetX1() int32`

GetX1 returns the X1 field if non-nil, zero value otherwise.

### GetX1Ok

`func (o *SwipeRequest) GetX1Ok() (*int32, bool)`

GetX1Ok returns a tuple with the X1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX1

`func (o *SwipeRequest) SetX1(v int32)`

SetX1 sets X1 field to given value.


### GetY1

`func (o *SwipeRequest) GetY1() int32`

GetY1 returns the Y1 field if non-nil, zero value otherwise.

### GetY1Ok

`func (o *SwipeRequest) GetY1Ok() (*int32, bool)`

GetY1Ok returns a tuple with the Y1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY1

`func (o *SwipeRequest) SetY1(v int32)`

SetY1 sets Y1 field to given value.


### GetX2

`func (o *SwipeRequest) GetX2() int32`

GetX2 returns the X2 field if non-nil, zero value otherwise.

### GetX2Ok

`func (o *SwipeRequest) GetX2Ok() (*int32, bool)`

GetX2Ok returns a tuple with the X2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX2

`func (o *SwipeRequest) SetX2(v int32)`

SetX2 sets X2 field to given value.


### GetY2

`func (o *SwipeRequest) GetY2() int32`

GetY2 returns the Y2 field if non-nil, zero value otherwise.

### GetY2Ok

`func (o *SwipeRequest) GetY2Ok() (*int32, bool)`

GetY2Ok returns a tuple with the Y2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY2

`func (o *SwipeRequest) SetY2(v int32)`

SetY2 sets Y2 field to given value.


### GetDurationMs

`func (o *SwipeRequest) GetDurationMs() int32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *SwipeRequest) GetDurationMsOk() (*int32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *SwipeRequest) SetDurationMs(v int32)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *SwipeRequest) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


