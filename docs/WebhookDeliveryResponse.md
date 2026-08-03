# WebhookDeliveryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeliveryId** | **string** |  | 
**WebhookId** | **string** |  | 
**TaskId** | **string** |  | 
**Renderer** | **string** | Payload format used at delivery time. OpenAPI callbacks registered through this gateway use the &#x60;openapi&#x60; renderer (TaskEvent envelope).  | 
**Status** | **string** | Lifecycle of this attempt. * &#x60;pending&#x60; — queued for first attempt or awaiting retry * &#x60;succeeded&#x60; — receiver returned 2xx (terminal) * &#x60;failed&#x60; — last attempt errored / non-2xx, will be   retried per the backoff schedule (intermediate) * &#x60;dead_letter&#x60; — exhausted retry schedule (terminal).   Recoverable via &#x60;POST .../redeliver&#x60;.  | 
**AttemptNum** | **int32** | 1-indexed count of attempts already made. New rows surface as 0 until the dispatcher claims them.  | 
**LastResponseStatus** | **int32** | HTTP status of the most recent attempt. &#x60;0&#x60; means \&quot;never attempted\&quot; (transport or pre-flight error before any HTTP response).  | 
**LastError** | Pointer to **string** | Truncated copy (≤ 2 KiB) of the most recent failure reason — either a transport error (&#x60;transport: connection refused&#x60;) or &#x60;http_status: &lt;code&gt;&#x60;. Empty on success / never attempted.  | [optional] 
**NextAttemptAt** | **time.Time** | Scheduled time for the next attempt while &#x60;status&#x60; is &#x60;pending&#x60;. On terminal rows this carries the last scheduled value (no longer meaningful for retries).  | 
**LastAttemptedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**CompletedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewWebhookDeliveryResponse

`func NewWebhookDeliveryResponse(deliveryId string, webhookId string, taskId string, renderer string, status string, attemptNum int32, lastResponseStatus int32, nextAttemptAt time.Time, createdAt time.Time, ) *WebhookDeliveryResponse`

NewWebhookDeliveryResponse instantiates a new WebhookDeliveryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookDeliveryResponseWithDefaults

`func NewWebhookDeliveryResponseWithDefaults() *WebhookDeliveryResponse`

NewWebhookDeliveryResponseWithDefaults instantiates a new WebhookDeliveryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeliveryId

`func (o *WebhookDeliveryResponse) GetDeliveryId() string`

GetDeliveryId returns the DeliveryId field if non-nil, zero value otherwise.

### GetDeliveryIdOk

`func (o *WebhookDeliveryResponse) GetDeliveryIdOk() (*string, bool)`

GetDeliveryIdOk returns a tuple with the DeliveryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryId

`func (o *WebhookDeliveryResponse) SetDeliveryId(v string)`

SetDeliveryId sets DeliveryId field to given value.


### GetWebhookId

`func (o *WebhookDeliveryResponse) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *WebhookDeliveryResponse) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *WebhookDeliveryResponse) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.


### GetTaskId

`func (o *WebhookDeliveryResponse) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *WebhookDeliveryResponse) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *WebhookDeliveryResponse) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.


### GetRenderer

`func (o *WebhookDeliveryResponse) GetRenderer() string`

GetRenderer returns the Renderer field if non-nil, zero value otherwise.

### GetRendererOk

`func (o *WebhookDeliveryResponse) GetRendererOk() (*string, bool)`

GetRendererOk returns a tuple with the Renderer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderer

`func (o *WebhookDeliveryResponse) SetRenderer(v string)`

SetRenderer sets Renderer field to given value.


### GetStatus

`func (o *WebhookDeliveryResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebhookDeliveryResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebhookDeliveryResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAttemptNum

`func (o *WebhookDeliveryResponse) GetAttemptNum() int32`

GetAttemptNum returns the AttemptNum field if non-nil, zero value otherwise.

### GetAttemptNumOk

`func (o *WebhookDeliveryResponse) GetAttemptNumOk() (*int32, bool)`

GetAttemptNumOk returns a tuple with the AttemptNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptNum

`func (o *WebhookDeliveryResponse) SetAttemptNum(v int32)`

SetAttemptNum sets AttemptNum field to given value.


### GetLastResponseStatus

`func (o *WebhookDeliveryResponse) GetLastResponseStatus() int32`

GetLastResponseStatus returns the LastResponseStatus field if non-nil, zero value otherwise.

### GetLastResponseStatusOk

`func (o *WebhookDeliveryResponse) GetLastResponseStatusOk() (*int32, bool)`

GetLastResponseStatusOk returns a tuple with the LastResponseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResponseStatus

`func (o *WebhookDeliveryResponse) SetLastResponseStatus(v int32)`

SetLastResponseStatus sets LastResponseStatus field to given value.


### GetLastError

`func (o *WebhookDeliveryResponse) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *WebhookDeliveryResponse) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *WebhookDeliveryResponse) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *WebhookDeliveryResponse) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### GetNextAttemptAt

`func (o *WebhookDeliveryResponse) GetNextAttemptAt() time.Time`

GetNextAttemptAt returns the NextAttemptAt field if non-nil, zero value otherwise.

### GetNextAttemptAtOk

`func (o *WebhookDeliveryResponse) GetNextAttemptAtOk() (*time.Time, bool)`

GetNextAttemptAtOk returns a tuple with the NextAttemptAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAttemptAt

`func (o *WebhookDeliveryResponse) SetNextAttemptAt(v time.Time)`

SetNextAttemptAt sets NextAttemptAt field to given value.


### GetLastAttemptedAt

`func (o *WebhookDeliveryResponse) GetLastAttemptedAt() time.Time`

GetLastAttemptedAt returns the LastAttemptedAt field if non-nil, zero value otherwise.

### GetLastAttemptedAtOk

`func (o *WebhookDeliveryResponse) GetLastAttemptedAtOk() (*time.Time, bool)`

GetLastAttemptedAtOk returns a tuple with the LastAttemptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAttemptedAt

`func (o *WebhookDeliveryResponse) SetLastAttemptedAt(v time.Time)`

SetLastAttemptedAt sets LastAttemptedAt field to given value.

### HasLastAttemptedAt

`func (o *WebhookDeliveryResponse) HasLastAttemptedAt() bool`

HasLastAttemptedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *WebhookDeliveryResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WebhookDeliveryResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WebhookDeliveryResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCompletedAt

`func (o *WebhookDeliveryResponse) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *WebhookDeliveryResponse) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *WebhookDeliveryResponse) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *WebhookDeliveryResponse) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


