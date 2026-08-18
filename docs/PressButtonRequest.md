# PressButtonRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Button** | **string** | Hardware / nav button name (&#x60;back&#x60; / &#x60;home&#x60; / &#x60;recent&#x60; / &#x60;power&#x60; / &#x60;volume_up&#x60; / &#x60;volume_down&#x60; / &#x60;enter&#x60;).  | 

## Methods

### NewPressButtonRequest

`func NewPressButtonRequest(button string, ) *PressButtonRequest`

NewPressButtonRequest instantiates a new PressButtonRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPressButtonRequestWithDefaults

`func NewPressButtonRequestWithDefaults() *PressButtonRequest`

NewPressButtonRequestWithDefaults instantiates a new PressButtonRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetButton

`func (o *PressButtonRequest) GetButton() string`

GetButton returns the Button field if non-nil, zero value otherwise.

### GetButtonOk

`func (o *PressButtonRequest) GetButtonOk() (*string, bool)`

GetButtonOk returns a tuple with the Button field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButton

`func (o *PressButtonRequest) SetButton(v string)`

SetButton sets Button field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


