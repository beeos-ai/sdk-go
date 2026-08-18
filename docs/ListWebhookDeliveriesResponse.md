# ListWebhookDeliveriesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Data** | [**ListWebhookDeliveriesResponseAllOfData**](ListWebhookDeliveriesResponseAllOfData.md) |  | 

## Methods

### NewListWebhookDeliveriesResponse

`func NewListWebhookDeliveriesResponse(success bool, data ListWebhookDeliveriesResponseAllOfData, ) *ListWebhookDeliveriesResponse`

NewListWebhookDeliveriesResponse instantiates a new ListWebhookDeliveriesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWebhookDeliveriesResponseWithDefaults

`func NewListWebhookDeliveriesResponseWithDefaults() *ListWebhookDeliveriesResponse`

NewListWebhookDeliveriesResponseWithDefaults instantiates a new ListWebhookDeliveriesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ListWebhookDeliveriesResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ListWebhookDeliveriesResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ListWebhookDeliveriesResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *ListWebhookDeliveriesResponse) GetData() ListWebhookDeliveriesResponseAllOfData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListWebhookDeliveriesResponse) GetDataOk() (*ListWebhookDeliveriesResponseAllOfData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListWebhookDeliveriesResponse) SetData(v ListWebhookDeliveriesResponseAllOfData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


