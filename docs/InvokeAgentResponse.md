# InvokeAgentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Data** | [**InvokeAgentResponseAllOfData**](InvokeAgentResponseAllOfData.md) |  | 

## Methods

### NewInvokeAgentResponse

`func NewInvokeAgentResponse(success bool, data InvokeAgentResponseAllOfData, ) *InvokeAgentResponse`

NewInvokeAgentResponse instantiates a new InvokeAgentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvokeAgentResponseWithDefaults

`func NewInvokeAgentResponseWithDefaults() *InvokeAgentResponse`

NewInvokeAgentResponseWithDefaults instantiates a new InvokeAgentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *InvokeAgentResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *InvokeAgentResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *InvokeAgentResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *InvokeAgentResponse) GetData() InvokeAgentResponseAllOfData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InvokeAgentResponse) GetDataOk() (*InvokeAgentResponseAllOfData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InvokeAgentResponse) SetData(v InvokeAgentResponseAllOfData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


