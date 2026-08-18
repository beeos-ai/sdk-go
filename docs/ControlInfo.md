# ControlInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Online** | **bool** | Whether the device is currently reachable for control. | 
**Os** | **string** | Device OS (e.g. &#x60;android&#x60;, &#x60;linux&#x60;, &#x60;windows&#x60;, &#x60;macos&#x60;). | 
**ScreenWidth** | **int32** | Current screen width in device pixels. | 
**ScreenHeight** | **int32** | Current screen height in device pixels. | 
**CoordinateSystem** | **string** | Always &#x60;absolute_pixels&#x60; (origin top-left). | 
**SupportedActions** | **[]string** | Action verbs this device supports (e.g. &#x60;screenshot&#x60;, &#x60;tap&#x60;). | 
**CurrentApp** | Pointer to **string** | Foreground app / window, when known. | [optional] 

## Methods

### NewControlInfo

`func NewControlInfo(online bool, os string, screenWidth int32, screenHeight int32, coordinateSystem string, supportedActions []string, ) *ControlInfo`

NewControlInfo instantiates a new ControlInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControlInfoWithDefaults

`func NewControlInfoWithDefaults() *ControlInfo`

NewControlInfoWithDefaults instantiates a new ControlInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOnline

`func (o *ControlInfo) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *ControlInfo) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *ControlInfo) SetOnline(v bool)`

SetOnline sets Online field to given value.


### GetOs

`func (o *ControlInfo) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *ControlInfo) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *ControlInfo) SetOs(v string)`

SetOs sets Os field to given value.


### GetScreenWidth

`func (o *ControlInfo) GetScreenWidth() int32`

GetScreenWidth returns the ScreenWidth field if non-nil, zero value otherwise.

### GetScreenWidthOk

`func (o *ControlInfo) GetScreenWidthOk() (*int32, bool)`

GetScreenWidthOk returns a tuple with the ScreenWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenWidth

`func (o *ControlInfo) SetScreenWidth(v int32)`

SetScreenWidth sets ScreenWidth field to given value.


### GetScreenHeight

`func (o *ControlInfo) GetScreenHeight() int32`

GetScreenHeight returns the ScreenHeight field if non-nil, zero value otherwise.

### GetScreenHeightOk

`func (o *ControlInfo) GetScreenHeightOk() (*int32, bool)`

GetScreenHeightOk returns a tuple with the ScreenHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenHeight

`func (o *ControlInfo) SetScreenHeight(v int32)`

SetScreenHeight sets ScreenHeight field to given value.


### GetCoordinateSystem

`func (o *ControlInfo) GetCoordinateSystem() string`

GetCoordinateSystem returns the CoordinateSystem field if non-nil, zero value otherwise.

### GetCoordinateSystemOk

`func (o *ControlInfo) GetCoordinateSystemOk() (*string, bool)`

GetCoordinateSystemOk returns a tuple with the CoordinateSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinateSystem

`func (o *ControlInfo) SetCoordinateSystem(v string)`

SetCoordinateSystem sets CoordinateSystem field to given value.


### GetSupportedActions

`func (o *ControlInfo) GetSupportedActions() []string`

GetSupportedActions returns the SupportedActions field if non-nil, zero value otherwise.

### GetSupportedActionsOk

`func (o *ControlInfo) GetSupportedActionsOk() (*[]string, bool)`

GetSupportedActionsOk returns a tuple with the SupportedActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedActions

`func (o *ControlInfo) SetSupportedActions(v []string)`

SetSupportedActions sets SupportedActions field to given value.


### GetCurrentApp

`func (o *ControlInfo) GetCurrentApp() string`

GetCurrentApp returns the CurrentApp field if non-nil, zero value otherwise.

### GetCurrentAppOk

`func (o *ControlInfo) GetCurrentAppOk() (*string, bool)`

GetCurrentAppOk returns a tuple with the CurrentApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentApp

`func (o *ControlInfo) SetCurrentApp(v string)`

SetCurrentApp sets CurrentApp field to given value.

### HasCurrentApp

`func (o *ControlInfo) HasCurrentApp() bool`

HasCurrentApp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


