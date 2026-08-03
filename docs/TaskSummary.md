# TaskSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskId** | **string** | Stable task identifier (&#x3D;&#x3D; underlying channel_id). | 
**AgentId** | **string** |  | 
**CallerId** | Pointer to **string** | The owner user id of this task. Always equal to the caller for tasks they themselves submitted; populated from the channel metadata so it survives across sessions.  | [optional] 
**State** | **string** | Channel-level state: &#x60;active&#x60; while the task is open (the agent may still be working / paused at &#x60;input_required&#x60;), &#x60;closed&#x60; after terminal reply / cancel / timeout. NOT the derived &#x60;TaskStatus&#x60; — use &#x60;getTask&#x60; for that.  | 
**CloseReason** | Pointer to **string** | Populated when &#x60;state&#x3D;closed&#x60;. Mirrors the underlying channel&#39;s &#x60;closed_reason&#x60; so SDKs can branch without re-pulling the message log. Canonical values include &#x60;completed&#x60;, &#x60;canceled&#x60;, &#x60;timeout&#x60;.  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**ClosedAt** | Pointer to **time.Time** |  | [optional] 
**DeadlineAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewTaskSummary

`func NewTaskSummary(taskId string, agentId string, state string, ) *TaskSummary`

NewTaskSummary instantiates a new TaskSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskSummaryWithDefaults

`func NewTaskSummaryWithDefaults() *TaskSummary`

NewTaskSummaryWithDefaults instantiates a new TaskSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskId

`func (o *TaskSummary) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *TaskSummary) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *TaskSummary) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.


### GetAgentId

`func (o *TaskSummary) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *TaskSummary) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *TaskSummary) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.


### GetCallerId

`func (o *TaskSummary) GetCallerId() string`

GetCallerId returns the CallerId field if non-nil, zero value otherwise.

### GetCallerIdOk

`func (o *TaskSummary) GetCallerIdOk() (*string, bool)`

GetCallerIdOk returns a tuple with the CallerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerId

`func (o *TaskSummary) SetCallerId(v string)`

SetCallerId sets CallerId field to given value.

### HasCallerId

`func (o *TaskSummary) HasCallerId() bool`

HasCallerId returns a boolean if a field has been set.

### GetState

`func (o *TaskSummary) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TaskSummary) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TaskSummary) SetState(v string)`

SetState sets State field to given value.


### GetCloseReason

`func (o *TaskSummary) GetCloseReason() string`

GetCloseReason returns the CloseReason field if non-nil, zero value otherwise.

### GetCloseReasonOk

`func (o *TaskSummary) GetCloseReasonOk() (*string, bool)`

GetCloseReasonOk returns a tuple with the CloseReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseReason

`func (o *TaskSummary) SetCloseReason(v string)`

SetCloseReason sets CloseReason field to given value.

### HasCloseReason

`func (o *TaskSummary) HasCloseReason() bool`

HasCloseReason returns a boolean if a field has been set.

### GetMetadata

`func (o *TaskSummary) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TaskSummary) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TaskSummary) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TaskSummary) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TaskSummary) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TaskSummary) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TaskSummary) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TaskSummary) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetClosedAt

`func (o *TaskSummary) GetClosedAt() time.Time`

GetClosedAt returns the ClosedAt field if non-nil, zero value otherwise.

### GetClosedAtOk

`func (o *TaskSummary) GetClosedAtOk() (*time.Time, bool)`

GetClosedAtOk returns a tuple with the ClosedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedAt

`func (o *TaskSummary) SetClosedAt(v time.Time)`

SetClosedAt sets ClosedAt field to given value.

### HasClosedAt

`func (o *TaskSummary) HasClosedAt() bool`

HasClosedAt returns a boolean if a field has been set.

### GetDeadlineAt

`func (o *TaskSummary) GetDeadlineAt() time.Time`

GetDeadlineAt returns the DeadlineAt field if non-nil, zero value otherwise.

### GetDeadlineAtOk

`func (o *TaskSummary) GetDeadlineAtOk() (*time.Time, bool)`

GetDeadlineAtOk returns a tuple with the DeadlineAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadlineAt

`func (o *TaskSummary) SetDeadlineAt(v time.Time)`

SetDeadlineAt sets DeadlineAt field to given value.

### HasDeadlineAt

`func (o *TaskSummary) HasDeadlineAt() bool`

HasDeadlineAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


