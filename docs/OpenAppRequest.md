# OpenAppRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | **string** | Android package id or installed app display name. |
**Activity** | Pointer to **string** |  | [optional]
**Locale** | Pointer to **string** | Optional BCP-47 per-app locale (Android 13+, best effort). | [optional]

## Methods

### NewOpenAppRequest

`func NewOpenAppRequest(app string, ) *OpenAppRequest`

NewOpenAppRequest instantiates a new OpenAppRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenAppRequestWithDefaults

`func NewOpenAppRequestWithDefaults() *OpenAppRequest`

NewOpenAppRequestWithDefaults instantiates a new OpenAppRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *OpenAppRequest) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *OpenAppRequest) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *OpenAppRequest) SetApp(v string)`

SetApp sets App field to given value.


### GetActivity

`func (o *OpenAppRequest) GetActivity() string`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *OpenAppRequest) GetActivityOk() (*string, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *OpenAppRequest) SetActivity(v string)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *OpenAppRequest) HasActivity() bool`

HasActivity returns a boolean if a field has been set.

### GetLocale

`func (o *OpenAppRequest) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *OpenAppRequest) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *OpenAppRequest) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *OpenAppRequest) HasLocale() bool`

HasLocale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


