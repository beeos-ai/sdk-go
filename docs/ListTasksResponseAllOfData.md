# ListTasksResponseAllOfData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tasks** | [**[]TaskSummary**](TaskSummary.md) |  | 
**NextSince** | Pointer to **string** | Pagination cursor. Empty when the caller has reached the end of results.  | [optional] 

## Methods

### NewListTasksResponseAllOfData

`func NewListTasksResponseAllOfData(tasks []TaskSummary, ) *ListTasksResponseAllOfData`

NewListTasksResponseAllOfData instantiates a new ListTasksResponseAllOfData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTasksResponseAllOfDataWithDefaults

`func NewListTasksResponseAllOfDataWithDefaults() *ListTasksResponseAllOfData`

NewListTasksResponseAllOfDataWithDefaults instantiates a new ListTasksResponseAllOfData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTasks

`func (o *ListTasksResponseAllOfData) GetTasks() []TaskSummary`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *ListTasksResponseAllOfData) GetTasksOk() (*[]TaskSummary, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *ListTasksResponseAllOfData) SetTasks(v []TaskSummary)`

SetTasks sets Tasks field to given value.


### GetNextSince

`func (o *ListTasksResponseAllOfData) GetNextSince() string`

GetNextSince returns the NextSince field if non-nil, zero value otherwise.

### GetNextSinceOk

`func (o *ListTasksResponseAllOfData) GetNextSinceOk() (*string, bool)`

GetNextSinceOk returns a tuple with the NextSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextSince

`func (o *ListTasksResponseAllOfData) SetNextSince(v string)`

SetNextSince sets NextSince field to given value.

### HasNextSince

`func (o *ListTasksResponseAllOfData) HasNextSince() bool`

HasNextSince returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


