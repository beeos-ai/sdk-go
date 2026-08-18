# FileMetaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | **string** |  | 
**Filename** | Pointer to **string** |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**DownloadUrl** | Pointer to **string** | Presigned GET URL (24h default TTL on S3). Agents receive the same URL when the file is attached to an invoke / task envelope.  | [optional] 

## Methods

### NewFileMetaResponse

`func NewFileMetaResponse(fileId string, ) *FileMetaResponse`

NewFileMetaResponse instantiates a new FileMetaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileMetaResponseWithDefaults

`func NewFileMetaResponseWithDefaults() *FileMetaResponse`

NewFileMetaResponseWithDefaults instantiates a new FileMetaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileId

`func (o *FileMetaResponse) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *FileMetaResponse) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *FileMetaResponse) SetFileId(v string)`

SetFileId sets FileId field to given value.


### GetFilename

`func (o *FileMetaResponse) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *FileMetaResponse) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *FileMetaResponse) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *FileMetaResponse) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetContentType

`func (o *FileMetaResponse) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *FileMetaResponse) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *FileMetaResponse) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *FileMetaResponse) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetSize

`func (o *FileMetaResponse) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FileMetaResponse) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FileMetaResponse) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *FileMetaResponse) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetDownloadUrl

`func (o *FileMetaResponse) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *FileMetaResponse) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *FileMetaResponse) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *FileMetaResponse) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


