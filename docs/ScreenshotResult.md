# ScreenshotResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | **string** | Stable BeeOS file id (resolvable via &#x60;GET /files/{id}&#x60;). | 
**Url** | **string** | Short-lived presigned download URL for the image. | 
**Format** | Pointer to **string** | Image codec (&#x60;png&#x60; / &#x60;jpeg&#x60;). | [optional] 
**Width** | Pointer to **int32** |  | [optional] 
**Height** | Pointer to **int32** |  | [optional] 

## Methods

### NewScreenshotResult

`func NewScreenshotResult(fileId string, url string, ) *ScreenshotResult`

NewScreenshotResult instantiates a new ScreenshotResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScreenshotResultWithDefaults

`func NewScreenshotResultWithDefaults() *ScreenshotResult`

NewScreenshotResultWithDefaults instantiates a new ScreenshotResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileId

`func (o *ScreenshotResult) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *ScreenshotResult) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *ScreenshotResult) SetFileId(v string)`

SetFileId sets FileId field to given value.


### GetUrl

`func (o *ScreenshotResult) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ScreenshotResult) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ScreenshotResult) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetFormat

`func (o *ScreenshotResult) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ScreenshotResult) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ScreenshotResult) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *ScreenshotResult) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetWidth

`func (o *ScreenshotResult) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *ScreenshotResult) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *ScreenshotResult) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *ScreenshotResult) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *ScreenshotResult) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ScreenshotResult) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ScreenshotResult) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ScreenshotResult) HasHeight() bool`

HasHeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


