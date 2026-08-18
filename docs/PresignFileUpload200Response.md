# PresignFileUpload200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Data** | Pointer to [**PresignUploadResponse**](PresignUploadResponse.md) |  | [optional] 

## Methods

### NewPresignFileUpload200Response

`func NewPresignFileUpload200Response(success bool, ) *PresignFileUpload200Response`

NewPresignFileUpload200Response instantiates a new PresignFileUpload200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPresignFileUpload200ResponseWithDefaults

`func NewPresignFileUpload200ResponseWithDefaults() *PresignFileUpload200Response`

NewPresignFileUpload200ResponseWithDefaults instantiates a new PresignFileUpload200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *PresignFileUpload200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *PresignFileUpload200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *PresignFileUpload200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *PresignFileUpload200Response) GetData() PresignUploadResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PresignFileUpload200Response) GetDataOk() (*PresignUploadResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PresignFileUpload200Response) SetData(v PresignUploadResponse)`

SetData sets Data field to given value.

### HasData

`func (o *PresignFileUpload200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


