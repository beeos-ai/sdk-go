# PresignUploadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | **string** | Stable identifier — attach via &#x60;attachments[].file_id&#x60;. | 
**UploadUrl** | **string** | Presigned PUT URL. Client uploads bytes directly with the same &#x60;Content-Type&#x60; supplied in the request. Single-use; expires per storage-backend policy.  | 
**DownloadUrl** | Pointer to **string** | Presigned GET URL for the same file, when the storage backend supports it (S3 / MinIO). Empty for backends that don&#39;t presign downloads — fall back to &#x60;GET /api/v1/files/{id}&#x60; in that case.  | [optional] 
**DownloadExpiresIn** | Pointer to **int64** | TTL of &#x60;download_url&#x60; in seconds. | [optional] 

## Methods

### NewPresignUploadResponse

`func NewPresignUploadResponse(fileId string, uploadUrl string, ) *PresignUploadResponse`

NewPresignUploadResponse instantiates a new PresignUploadResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPresignUploadResponseWithDefaults

`func NewPresignUploadResponseWithDefaults() *PresignUploadResponse`

NewPresignUploadResponseWithDefaults instantiates a new PresignUploadResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileId

`func (o *PresignUploadResponse) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *PresignUploadResponse) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *PresignUploadResponse) SetFileId(v string)`

SetFileId sets FileId field to given value.


### GetUploadUrl

`func (o *PresignUploadResponse) GetUploadUrl() string`

GetUploadUrl returns the UploadUrl field if non-nil, zero value otherwise.

### GetUploadUrlOk

`func (o *PresignUploadResponse) GetUploadUrlOk() (*string, bool)`

GetUploadUrlOk returns a tuple with the UploadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadUrl

`func (o *PresignUploadResponse) SetUploadUrl(v string)`

SetUploadUrl sets UploadUrl field to given value.


### GetDownloadUrl

`func (o *PresignUploadResponse) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *PresignUploadResponse) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *PresignUploadResponse) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *PresignUploadResponse) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.

### GetDownloadExpiresIn

`func (o *PresignUploadResponse) GetDownloadExpiresIn() int64`

GetDownloadExpiresIn returns the DownloadExpiresIn field if non-nil, zero value otherwise.

### GetDownloadExpiresInOk

`func (o *PresignUploadResponse) GetDownloadExpiresInOk() (*int64, bool)`

GetDownloadExpiresInOk returns a tuple with the DownloadExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadExpiresIn

`func (o *PresignUploadResponse) SetDownloadExpiresIn(v int64)`

SetDownloadExpiresIn sets DownloadExpiresIn field to given value.

### HasDownloadExpiresIn

`func (o *PresignUploadResponse) HasDownloadExpiresIn() bool`

HasDownloadExpiresIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


