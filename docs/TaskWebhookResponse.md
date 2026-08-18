# TaskWebhookResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Data** | [**TaskWebhookResponseAllOfData**](TaskWebhookResponseAllOfData.md) |  | 

## Methods

### NewTaskWebhookResponse

`func NewTaskWebhookResponse(success bool, data TaskWebhookResponseAllOfData, ) *TaskWebhookResponse`

NewTaskWebhookResponse instantiates a new TaskWebhookResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskWebhookResponseWithDefaults

`func NewTaskWebhookResponseWithDefaults() *TaskWebhookResponse`

NewTaskWebhookResponseWithDefaults instantiates a new TaskWebhookResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *TaskWebhookResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *TaskWebhookResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *TaskWebhookResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *TaskWebhookResponse) GetData() TaskWebhookResponseAllOfData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TaskWebhookResponse) GetDataOk() (*TaskWebhookResponseAllOfData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TaskWebhookResponse) SetData(v TaskWebhookResponseAllOfData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


