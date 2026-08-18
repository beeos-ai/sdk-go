# TaskCreatedResponseAllOfData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskId** | **string** | Stable task identifier (&#x3D;&#x3D; underlying channel_id). | 
**AgentId** | **string** |  | 
**Status** | [**TaskStatus**](TaskStatus.md) |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewTaskCreatedResponseAllOfData

`func NewTaskCreatedResponseAllOfData(taskId string, agentId string, status TaskStatus, ) *TaskCreatedResponseAllOfData`

NewTaskCreatedResponseAllOfData instantiates a new TaskCreatedResponseAllOfData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskCreatedResponseAllOfDataWithDefaults

`func NewTaskCreatedResponseAllOfDataWithDefaults() *TaskCreatedResponseAllOfData`

NewTaskCreatedResponseAllOfDataWithDefaults instantiates a new TaskCreatedResponseAllOfData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskId

`func (o *TaskCreatedResponseAllOfData) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *TaskCreatedResponseAllOfData) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *TaskCreatedResponseAllOfData) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.


### GetAgentId

`func (o *TaskCreatedResponseAllOfData) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *TaskCreatedResponseAllOfData) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *TaskCreatedResponseAllOfData) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.


### GetStatus

`func (o *TaskCreatedResponseAllOfData) GetStatus() TaskStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaskCreatedResponseAllOfData) GetStatusOk() (*TaskStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaskCreatedResponseAllOfData) SetStatus(v TaskStatus)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *TaskCreatedResponseAllOfData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TaskCreatedResponseAllOfData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TaskCreatedResponseAllOfData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TaskCreatedResponseAllOfData) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


