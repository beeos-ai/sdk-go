# DeployModelsListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Data** | [**[]DeployModelItem**](DeployModelItem.md) |  | 

## Methods

### NewDeployModelsListResponse

`func NewDeployModelsListResponse(success bool, data []DeployModelItem, ) *DeployModelsListResponse`

NewDeployModelsListResponse instantiates a new DeployModelsListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployModelsListResponseWithDefaults

`func NewDeployModelsListResponseWithDefaults() *DeployModelsListResponse`

NewDeployModelsListResponseWithDefaults instantiates a new DeployModelsListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *DeployModelsListResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *DeployModelsListResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *DeployModelsListResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *DeployModelsListResponse) GetData() []DeployModelItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DeployModelsListResponse) GetDataOk() (*[]DeployModelItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DeployModelsListResponse) SetData(v []DeployModelItem)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


