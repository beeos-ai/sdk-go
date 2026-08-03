# PresignUploadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filename** | **string** | Original filename hint; stored on the file record. | 
**ContentType** | **string** | MIME type of the bytes being uploaded (e.g. &#x60;image/png&#x60;). Uploads with a different &#x60;Content-Type&#x60; may be rejected by the storage backend.  | 

## Methods

### NewPresignUploadRequest

`func NewPresignUploadRequest(filename string, contentType string, ) *PresignUploadRequest`

NewPresignUploadRequest instantiates a new PresignUploadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPresignUploadRequestWithDefaults

`func NewPresignUploadRequestWithDefaults() *PresignUploadRequest`

NewPresignUploadRequestWithDefaults instantiates a new PresignUploadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilename

`func (o *PresignUploadRequest) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *PresignUploadRequest) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *PresignUploadRequest) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetContentType

`func (o *PresignUploadRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *PresignUploadRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *PresignUploadRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


