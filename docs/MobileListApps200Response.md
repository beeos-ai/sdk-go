# MobileListApps200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  |
**Data** | Pointer to [**AppListResult**](AppListResult.md) |  | [optional]

## Methods

### NewMobileListApps200Response

`func NewMobileListApps200Response(success bool, ) *MobileListApps200Response`

NewMobileListApps200Response instantiates a new MobileListApps200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileListApps200ResponseWithDefaults

`func NewMobileListApps200ResponseWithDefaults() *MobileListApps200Response`

NewMobileListApps200ResponseWithDefaults instantiates a new MobileListApps200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *MobileListApps200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *MobileListApps200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *MobileListApps200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *MobileListApps200Response) GetData() AppListResult`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MobileListApps200Response) GetDataOk() (*AppListResult, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MobileListApps200Response) SetData(v AppListResult)`

SetData sets Data field to given value.

### HasData

`func (o *MobileListApps200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


