# RedeliverWebhook202Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Data** | [**WebhookDeliveryResponse**](WebhookDeliveryResponse.md) |  | 

## Methods

### NewRedeliverWebhook202Response

`func NewRedeliverWebhook202Response(success bool, data WebhookDeliveryResponse, ) *RedeliverWebhook202Response`

NewRedeliverWebhook202Response instantiates a new RedeliverWebhook202Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedeliverWebhook202ResponseWithDefaults

`func NewRedeliverWebhook202ResponseWithDefaults() *RedeliverWebhook202Response`

NewRedeliverWebhook202ResponseWithDefaults instantiates a new RedeliverWebhook202Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *RedeliverWebhook202Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *RedeliverWebhook202Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *RedeliverWebhook202Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *RedeliverWebhook202Response) GetData() WebhookDeliveryResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RedeliverWebhook202Response) GetDataOk() (*WebhookDeliveryResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RedeliverWebhook202Response) SetData(v WebhookDeliveryResponse)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


