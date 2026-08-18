# TaskWebhookResponseAllOfData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookId** | **string** |  | 
**TaskId** | **string** |  | 
**Url** | **string** |  | 
**HasSecret** | **bool** | &#x60;true&#x60; when an HMAC signing secret is configured for this webhook (P2-A). The raw secret is never returned — receivers verify against the secret they originally registered.  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewTaskWebhookResponseAllOfData

`func NewTaskWebhookResponseAllOfData(webhookId string, taskId string, url string, hasSecret bool, ) *TaskWebhookResponseAllOfData`

NewTaskWebhookResponseAllOfData instantiates a new TaskWebhookResponseAllOfData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskWebhookResponseAllOfDataWithDefaults

`func NewTaskWebhookResponseAllOfDataWithDefaults() *TaskWebhookResponseAllOfData`

NewTaskWebhookResponseAllOfDataWithDefaults instantiates a new TaskWebhookResponseAllOfData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookId

`func (o *TaskWebhookResponseAllOfData) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *TaskWebhookResponseAllOfData) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *TaskWebhookResponseAllOfData) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.


### GetTaskId

`func (o *TaskWebhookResponseAllOfData) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *TaskWebhookResponseAllOfData) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *TaskWebhookResponseAllOfData) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.


### GetUrl

`func (o *TaskWebhookResponseAllOfData) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TaskWebhookResponseAllOfData) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TaskWebhookResponseAllOfData) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHasSecret

`func (o *TaskWebhookResponseAllOfData) GetHasSecret() bool`

GetHasSecret returns the HasSecret field if non-nil, zero value otherwise.

### GetHasSecretOk

`func (o *TaskWebhookResponseAllOfData) GetHasSecretOk() (*bool, bool)`

GetHasSecretOk returns a tuple with the HasSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSecret

`func (o *TaskWebhookResponseAllOfData) SetHasSecret(v bool)`

SetHasSecret sets HasSecret field to given value.


### GetCreatedAt

`func (o *TaskWebhookResponseAllOfData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TaskWebhookResponseAllOfData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TaskWebhookResponseAllOfData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TaskWebhookResponseAllOfData) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


