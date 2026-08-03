# AttachmentRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | **string** | File identifier returned from [&#x60;POST /api/v1/files/presign-upload&#x60;](#operation/presignFileUpload).  | 
**Filename** | Pointer to **string** | Optional override for the file_record&#39;s filename. | [optional] 
**ContentType** | Pointer to **string** | Optional override for the file_record&#39;s MIME type. | [optional] 

## Methods

### NewAttachmentRef

`func NewAttachmentRef(fileId string, ) *AttachmentRef`

NewAttachmentRef instantiates a new AttachmentRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentRefWithDefaults

`func NewAttachmentRefWithDefaults() *AttachmentRef`

NewAttachmentRefWithDefaults instantiates a new AttachmentRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileId

`func (o *AttachmentRef) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *AttachmentRef) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *AttachmentRef) SetFileId(v string)`

SetFileId sets FileId field to given value.


### GetFilename

`func (o *AttachmentRef) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *AttachmentRef) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *AttachmentRef) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *AttachmentRef) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetContentType

`func (o *AttachmentRef) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *AttachmentRef) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *AttachmentRef) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *AttachmentRef) HasContentType() bool`

HasContentType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


