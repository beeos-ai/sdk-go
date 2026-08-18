# RuntimeOperationSnapshotCancelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestedBy** | **string** |  | 
**Reason** | Pointer to **string** |  | [optional] 
**RequestedAt** | **time.Time** |  | 

## Methods

### NewRuntimeOperationSnapshotCancelRequest

`func NewRuntimeOperationSnapshotCancelRequest(requestedBy string, requestedAt time.Time, ) *RuntimeOperationSnapshotCancelRequest`

NewRuntimeOperationSnapshotCancelRequest instantiates a new RuntimeOperationSnapshotCancelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuntimeOperationSnapshotCancelRequestWithDefaults

`func NewRuntimeOperationSnapshotCancelRequestWithDefaults() *RuntimeOperationSnapshotCancelRequest`

NewRuntimeOperationSnapshotCancelRequestWithDefaults instantiates a new RuntimeOperationSnapshotCancelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestedBy

`func (o *RuntimeOperationSnapshotCancelRequest) GetRequestedBy() string`

GetRequestedBy returns the RequestedBy field if non-nil, zero value otherwise.

### GetRequestedByOk

`func (o *RuntimeOperationSnapshotCancelRequest) GetRequestedByOk() (*string, bool)`

GetRequestedByOk returns a tuple with the RequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedBy

`func (o *RuntimeOperationSnapshotCancelRequest) SetRequestedBy(v string)`

SetRequestedBy sets RequestedBy field to given value.


### GetReason

`func (o *RuntimeOperationSnapshotCancelRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *RuntimeOperationSnapshotCancelRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *RuntimeOperationSnapshotCancelRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *RuntimeOperationSnapshotCancelRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRequestedAt

`func (o *RuntimeOperationSnapshotCancelRequest) GetRequestedAt() time.Time`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *RuntimeOperationSnapshotCancelRequest) GetRequestedAtOk() (*time.Time, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *RuntimeOperationSnapshotCancelRequest) SetRequestedAt(v time.Time)`

SetRequestedAt sets RequestedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


