# ListTaskWebhooksResponseAllOfDataWebhooks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookId** | **string** |  | 
**TaskId** | **string** |  | 
**Url** | **string** |  | 
**HasSecret** | **bool** |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewListTaskWebhooksResponseAllOfDataWebhooks

`func NewListTaskWebhooksResponseAllOfDataWebhooks(webhookId string, taskId string, url string, hasSecret bool, ) *ListTaskWebhooksResponseAllOfDataWebhooks`

NewListTaskWebhooksResponseAllOfDataWebhooks instantiates a new ListTaskWebhooksResponseAllOfDataWebhooks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTaskWebhooksResponseAllOfDataWebhooksWithDefaults

`func NewListTaskWebhooksResponseAllOfDataWebhooksWithDefaults() *ListTaskWebhooksResponseAllOfDataWebhooks`

NewListTaskWebhooksResponseAllOfDataWebhooksWithDefaults instantiates a new ListTaskWebhooksResponseAllOfDataWebhooks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookId

`func (o *ListTaskWebhooksResponseAllOfDataWebhooks) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *ListTaskWebhooksResponseAllOfDataWebhooks) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *ListTaskWebhooksResponseAllOfDataWebhooks) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.


### GetTaskId

`func (o *ListTaskWebhooksResponseAllOfDataWebhooks) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *ListTaskWebhooksResponseAllOfDataWebhooks) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *ListTaskWebhooksResponseAllOfDataWebhooks) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.


### GetUrl

`func (o *ListTaskWebhooksResponseAllOfDataWebhooks) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ListTaskWebhooksResponseAllOfDataWebhooks) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ListTaskWebhooksResponseAllOfDataWebhooks) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHasSecret

`func (o *ListTaskWebhooksResponseAllOfDataWebhooks) GetHasSecret() bool`

GetHasSecret returns the HasSecret field if non-nil, zero value otherwise.

### GetHasSecretOk

`func (o *ListTaskWebhooksResponseAllOfDataWebhooks) GetHasSecretOk() (*bool, bool)`

GetHasSecretOk returns a tuple with the HasSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSecret

`func (o *ListTaskWebhooksResponseAllOfDataWebhooks) SetHasSecret(v bool)`

SetHasSecret sets HasSecret field to given value.


### GetCreatedAt

`func (o *ListTaskWebhooksResponseAllOfDataWebhooks) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListTaskWebhooksResponseAllOfDataWebhooks) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListTaskWebhooksResponseAllOfDataWebhooks) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ListTaskWebhooksResponseAllOfDataWebhooks) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


