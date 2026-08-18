# ServiceOperationsListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operations** | [**[]RuntimeOperationSnapshot**](RuntimeOperationSnapshot.md) |  | 
**NextCursor** | Pointer to **string** |  | [optional] 

## Methods

### NewServiceOperationsListResponse

`func NewServiceOperationsListResponse(operations []RuntimeOperationSnapshot, ) *ServiceOperationsListResponse`

NewServiceOperationsListResponse instantiates a new ServiceOperationsListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceOperationsListResponseWithDefaults

`func NewServiceOperationsListResponseWithDefaults() *ServiceOperationsListResponse`

NewServiceOperationsListResponseWithDefaults instantiates a new ServiceOperationsListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperations

`func (o *ServiceOperationsListResponse) GetOperations() []RuntimeOperationSnapshot`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *ServiceOperationsListResponse) GetOperationsOk() (*[]RuntimeOperationSnapshot, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *ServiceOperationsListResponse) SetOperations(v []RuntimeOperationSnapshot)`

SetOperations sets Operations field to given value.


### GetNextCursor

`func (o *ServiceOperationsListResponse) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *ServiceOperationsListResponse) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *ServiceOperationsListResponse) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *ServiceOperationsListResponse) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


