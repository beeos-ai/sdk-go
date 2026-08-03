# PartFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Url** | **string** |  | 
**MimeType** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 

## Methods

### NewPartFile

`func NewPartFile(type_ string, url string, mimeType string, ) *PartFile`

NewPartFile instantiates a new PartFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartFileWithDefaults

`func NewPartFileWithDefaults() *PartFile`

NewPartFileWithDefaults instantiates a new PartFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PartFile) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PartFile) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PartFile) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *PartFile) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PartFile) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PartFile) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetMimeType

`func (o *PartFile) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *PartFile) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *PartFile) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.


### GetName

`func (o *PartFile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PartFile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PartFile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PartFile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *PartFile) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PartFile) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PartFile) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *PartFile) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


