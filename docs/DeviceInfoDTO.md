# DeviceInfoDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Device type (e.g. &#x60;android&#x60;, &#x60;windows&#x60;, &#x60;linux&#x60;, &#x60;macos&#x60;). | 
**Model** | **string** | Device model / hardware name. | 
**Os** | **string** | Operating system name. | 
**OsVersion** | **string** | Operating system version. | 
**Width** | **int32** | Registered screen width in device pixels. | 
**Height** | **int32** | Registered screen height in device pixels. | 
**Density** | **int32** | Registered screen density (DPI). | 
**Capabilities** | **[]string** | Device hardware / Agent-reported capability strings. Semantically different from &#x60;InstanceDataDTO.capabilities&#x60; (an instance-level capability JSON object).  | 
**Metadata** | Pointer to **map[string]interface{}** | Optional caller-/device-supplied extra metadata (raw JSON object). Omitted entirely when absent or malformed.  | [optional] 

## Methods

### NewDeviceInfoDTO

`func NewDeviceInfoDTO(type_ string, model string, os string, osVersion string, width int32, height int32, density int32, capabilities []string, ) *DeviceInfoDTO`

NewDeviceInfoDTO instantiates a new DeviceInfoDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceInfoDTOWithDefaults

`func NewDeviceInfoDTOWithDefaults() *DeviceInfoDTO`

NewDeviceInfoDTOWithDefaults instantiates a new DeviceInfoDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DeviceInfoDTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeviceInfoDTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeviceInfoDTO) SetType(v string)`

SetType sets Type field to given value.


### GetModel

`func (o *DeviceInfoDTO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DeviceInfoDTO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DeviceInfoDTO) SetModel(v string)`

SetModel sets Model field to given value.


### GetOs

`func (o *DeviceInfoDTO) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *DeviceInfoDTO) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *DeviceInfoDTO) SetOs(v string)`

SetOs sets Os field to given value.


### GetOsVersion

`func (o *DeviceInfoDTO) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *DeviceInfoDTO) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *DeviceInfoDTO) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.


### GetWidth

`func (o *DeviceInfoDTO) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *DeviceInfoDTO) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *DeviceInfoDTO) SetWidth(v int32)`

SetWidth sets Width field to given value.


### GetHeight

`func (o *DeviceInfoDTO) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *DeviceInfoDTO) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *DeviceInfoDTO) SetHeight(v int32)`

SetHeight sets Height field to given value.


### GetDensity

`func (o *DeviceInfoDTO) GetDensity() int32`

GetDensity returns the Density field if non-nil, zero value otherwise.

### GetDensityOk

`func (o *DeviceInfoDTO) GetDensityOk() (*int32, bool)`

GetDensityOk returns a tuple with the Density field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDensity

`func (o *DeviceInfoDTO) SetDensity(v int32)`

SetDensity sets Density field to given value.


### GetCapabilities

`func (o *DeviceInfoDTO) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *DeviceInfoDTO) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *DeviceInfoDTO) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.


### GetMetadata

`func (o *DeviceInfoDTO) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DeviceInfoDTO) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DeviceInfoDTO) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DeviceInfoDTO) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


