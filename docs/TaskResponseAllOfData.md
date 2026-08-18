# TaskResponseAllOfData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskId** | **string** |  | 
**AgentId** | **string** |  | 
**CallerId** | Pointer to **string** |  | [optional] 
**Status** | [**TaskStatus**](TaskStatus.md) |  | 
**Result** | Pointer to **interface{}** |  | [optional] 
**Error** | Pointer to **string** | Human-readable failure cause, present when &#x60;status&#x3D;failed&#x60;. | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**StartedAt** | Pointer to **time.Time** | First time the agent emitted any streaming output. | [optional] 
**CompletedAt** | Pointer to **time.Time** | Timestamp the task reached a terminal state. | [optional] 
**DeadlineAt** | Pointer to **time.Time** |  | [optional] 
**Truncated** | Pointer to **bool** | True when the server&#39;s message scan hit its internal cap (1000 messages) before reaching the end of the channel log. The reported &#x60;status&#x60; / &#x60;result&#x60; / &#x60;error&#x60; reflect only the scanned prefix; clients seeing this flag should follow up with &#x60;GET /tasks/{id}/events&#x60; (SSE) for the authoritative tail. Absent (false) on typical tasks.  | [optional] 

## Methods

### NewTaskResponseAllOfData

`func NewTaskResponseAllOfData(taskId string, agentId string, status TaskStatus, ) *TaskResponseAllOfData`

NewTaskResponseAllOfData instantiates a new TaskResponseAllOfData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskResponseAllOfDataWithDefaults

`func NewTaskResponseAllOfDataWithDefaults() *TaskResponseAllOfData`

NewTaskResponseAllOfDataWithDefaults instantiates a new TaskResponseAllOfData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskId

`func (o *TaskResponseAllOfData) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *TaskResponseAllOfData) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *TaskResponseAllOfData) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.


### GetAgentId

`func (o *TaskResponseAllOfData) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *TaskResponseAllOfData) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *TaskResponseAllOfData) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.


### GetCallerId

`func (o *TaskResponseAllOfData) GetCallerId() string`

GetCallerId returns the CallerId field if non-nil, zero value otherwise.

### GetCallerIdOk

`func (o *TaskResponseAllOfData) GetCallerIdOk() (*string, bool)`

GetCallerIdOk returns a tuple with the CallerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerId

`func (o *TaskResponseAllOfData) SetCallerId(v string)`

SetCallerId sets CallerId field to given value.

### HasCallerId

`func (o *TaskResponseAllOfData) HasCallerId() bool`

HasCallerId returns a boolean if a field has been set.

### GetStatus

`func (o *TaskResponseAllOfData) GetStatus() TaskStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaskResponseAllOfData) GetStatusOk() (*TaskStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaskResponseAllOfData) SetStatus(v TaskStatus)`

SetStatus sets Status field to given value.


### GetResult

`func (o *TaskResponseAllOfData) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *TaskResponseAllOfData) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *TaskResponseAllOfData) SetResult(v interface{})`

SetResult sets Result field to given value.

### HasResult

`func (o *TaskResponseAllOfData) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResultNil

`func (o *TaskResponseAllOfData) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *TaskResponseAllOfData) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetError

`func (o *TaskResponseAllOfData) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *TaskResponseAllOfData) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *TaskResponseAllOfData) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *TaskResponseAllOfData) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMetadata

`func (o *TaskResponseAllOfData) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TaskResponseAllOfData) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TaskResponseAllOfData) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TaskResponseAllOfData) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TaskResponseAllOfData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TaskResponseAllOfData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TaskResponseAllOfData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TaskResponseAllOfData) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *TaskResponseAllOfData) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *TaskResponseAllOfData) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *TaskResponseAllOfData) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *TaskResponseAllOfData) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *TaskResponseAllOfData) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *TaskResponseAllOfData) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *TaskResponseAllOfData) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *TaskResponseAllOfData) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetDeadlineAt

`func (o *TaskResponseAllOfData) GetDeadlineAt() time.Time`

GetDeadlineAt returns the DeadlineAt field if non-nil, zero value otherwise.

### GetDeadlineAtOk

`func (o *TaskResponseAllOfData) GetDeadlineAtOk() (*time.Time, bool)`

GetDeadlineAtOk returns a tuple with the DeadlineAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadlineAt

`func (o *TaskResponseAllOfData) SetDeadlineAt(v time.Time)`

SetDeadlineAt sets DeadlineAt field to given value.

### HasDeadlineAt

`func (o *TaskResponseAllOfData) HasDeadlineAt() bool`

HasDeadlineAt returns a boolean if a field has been set.

### GetTruncated

`func (o *TaskResponseAllOfData) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *TaskResponseAllOfData) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *TaskResponseAllOfData) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *TaskResponseAllOfData) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


