# GetFile200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Data** | Pointer to [**FileMetaResponse**](FileMetaResponse.md) |  | [optional] 

## Methods

### NewGetFile200Response

`func NewGetFile200Response(success bool, ) *GetFile200Response`

NewGetFile200Response instantiates a new GetFile200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFile200ResponseWithDefaults

`func NewGetFile200ResponseWithDefaults() *GetFile200Response`

NewGetFile200ResponseWithDefaults instantiates a new GetFile200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *GetFile200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetFile200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetFile200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *GetFile200Response) GetData() FileMetaResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetFile200Response) GetDataOk() (*FileMetaResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetFile200Response) SetData(v FileMetaResponse)`

SetData sets Data field to given value.

### HasData

`func (o *GetFile200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


