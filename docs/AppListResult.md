# AppListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apps** | [**[]AppInfo**](AppInfo.md) |  |

## Methods

### NewAppListResult

`func NewAppListResult(apps []AppInfo, ) *AppListResult`

NewAppListResult instantiates a new AppListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppListResultWithDefaults

`func NewAppListResultWithDefaults() *AppListResult`

NewAppListResultWithDefaults instantiates a new AppListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApps

`func (o *AppListResult) GetApps() []AppInfo`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *AppListResult) GetAppsOk() (*[]AppInfo, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *AppListResult) SetApps(v []AppInfo)`

SetApps sets Apps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


