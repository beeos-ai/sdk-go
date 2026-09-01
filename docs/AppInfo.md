# AppInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageId** | **string** |  |
**Label** | Pointer to **string** |  | [optional]

## Methods

### NewAppInfo

`func NewAppInfo(packageId string, ) *AppInfo`

NewAppInfo instantiates a new AppInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppInfoWithDefaults

`func NewAppInfoWithDefaults() *AppInfo`

NewAppInfoWithDefaults instantiates a new AppInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageId

`func (o *AppInfo) GetPackageId() string`

GetPackageId returns the PackageId field if non-nil, zero value otherwise.

### GetPackageIdOk

`func (o *AppInfo) GetPackageIdOk() (*string, bool)`

GetPackageIdOk returns a tuple with the PackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageId

`func (o *AppInfo) SetPackageId(v string)`

SetPackageId sets PackageId field to given value.


### GetLabel

`func (o *AppInfo) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AppInfo) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AppInfo) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *AppInfo) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


