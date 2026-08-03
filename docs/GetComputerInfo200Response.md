# GetComputerInfo200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Data** | Pointer to [**ControlInfo**](ControlInfo.md) |  | [optional] 

## Methods

### NewGetComputerInfo200Response

`func NewGetComputerInfo200Response(success bool, ) *GetComputerInfo200Response`

NewGetComputerInfo200Response instantiates a new GetComputerInfo200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetComputerInfo200ResponseWithDefaults

`func NewGetComputerInfo200ResponseWithDefaults() *GetComputerInfo200Response`

NewGetComputerInfo200ResponseWithDefaults instantiates a new GetComputerInfo200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *GetComputerInfo200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetComputerInfo200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetComputerInfo200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *GetComputerInfo200Response) GetData() ControlInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetComputerInfo200Response) GetDataOk() (*ControlInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetComputerInfo200Response) SetData(v ControlInfo)`

SetData sets Data field to given value.

### HasData

`func (o *GetComputerInfo200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


