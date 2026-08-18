# TaskSSEBackfillTruncated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OldestRedisOffset** | **int64** | Smallest offset still readable from the per-channel Redis Stream. Clients can resume from any &#x60;offset &gt;&#x3D; oldest_redis_offset&#x60; without missing chunks.  | 
**Since** | **int64** | The cursor the client reconnected with (echoed back for client-side correlation).  | 
**Hint** | Pointer to **string** | Human-readable English string explaining the gap. Informational only — SDKs should not pattern-match on it.  | [optional] 

## Methods

### NewTaskSSEBackfillTruncated

`func NewTaskSSEBackfillTruncated(oldestRedisOffset int64, since int64, ) *TaskSSEBackfillTruncated`

NewTaskSSEBackfillTruncated instantiates a new TaskSSEBackfillTruncated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskSSEBackfillTruncatedWithDefaults

`func NewTaskSSEBackfillTruncatedWithDefaults() *TaskSSEBackfillTruncated`

NewTaskSSEBackfillTruncatedWithDefaults instantiates a new TaskSSEBackfillTruncated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOldestRedisOffset

`func (o *TaskSSEBackfillTruncated) GetOldestRedisOffset() int64`

GetOldestRedisOffset returns the OldestRedisOffset field if non-nil, zero value otherwise.

### GetOldestRedisOffsetOk

`func (o *TaskSSEBackfillTruncated) GetOldestRedisOffsetOk() (*int64, bool)`

GetOldestRedisOffsetOk returns a tuple with the OldestRedisOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldestRedisOffset

`func (o *TaskSSEBackfillTruncated) SetOldestRedisOffset(v int64)`

SetOldestRedisOffset sets OldestRedisOffset field to given value.


### GetSince

`func (o *TaskSSEBackfillTruncated) GetSince() int64`

GetSince returns the Since field if non-nil, zero value otherwise.

### GetSinceOk

`func (o *TaskSSEBackfillTruncated) GetSinceOk() (*int64, bool)`

GetSinceOk returns a tuple with the Since field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSince

`func (o *TaskSSEBackfillTruncated) SetSince(v int64)`

SetSince sets Since field to given value.


### GetHint

`func (o *TaskSSEBackfillTruncated) GetHint() string`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *TaskSSEBackfillTruncated) GetHintOk() (*string, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *TaskSSEBackfillTruncated) SetHint(v string)`

SetHint sets Hint field to given value.

### HasHint

`func (o *TaskSSEBackfillTruncated) HasHint() bool`

HasHint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


