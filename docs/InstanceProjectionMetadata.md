# InstanceProjectionMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptedEpoch** | **string** |  | 
**AcceptedRevision** | **string** |  | 
**CanonicalHash** | **string** |  | 
**LastSyncedAt** | **time.Time** |  | 
**SyncState** | **string** |  | 
**AgentsComplete** | **bool** |  | 
**SkillsComplete** | **bool** |  | 

## Methods

### NewInstanceProjectionMetadata

`func NewInstanceProjectionMetadata(acceptedEpoch string, acceptedRevision string, canonicalHash string, lastSyncedAt time.Time, syncState string, agentsComplete bool, skillsComplete bool, ) *InstanceProjectionMetadata`

NewInstanceProjectionMetadata instantiates a new InstanceProjectionMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceProjectionMetadataWithDefaults

`func NewInstanceProjectionMetadataWithDefaults() *InstanceProjectionMetadata`

NewInstanceProjectionMetadataWithDefaults instantiates a new InstanceProjectionMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptedEpoch

`func (o *InstanceProjectionMetadata) GetAcceptedEpoch() string`

GetAcceptedEpoch returns the AcceptedEpoch field if non-nil, zero value otherwise.

### GetAcceptedEpochOk

`func (o *InstanceProjectionMetadata) GetAcceptedEpochOk() (*string, bool)`

GetAcceptedEpochOk returns a tuple with the AcceptedEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedEpoch

`func (o *InstanceProjectionMetadata) SetAcceptedEpoch(v string)`

SetAcceptedEpoch sets AcceptedEpoch field to given value.


### GetAcceptedRevision

`func (o *InstanceProjectionMetadata) GetAcceptedRevision() string`

GetAcceptedRevision returns the AcceptedRevision field if non-nil, zero value otherwise.

### GetAcceptedRevisionOk

`func (o *InstanceProjectionMetadata) GetAcceptedRevisionOk() (*string, bool)`

GetAcceptedRevisionOk returns a tuple with the AcceptedRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedRevision

`func (o *InstanceProjectionMetadata) SetAcceptedRevision(v string)`

SetAcceptedRevision sets AcceptedRevision field to given value.


### GetCanonicalHash

`func (o *InstanceProjectionMetadata) GetCanonicalHash() string`

GetCanonicalHash returns the CanonicalHash field if non-nil, zero value otherwise.

### GetCanonicalHashOk

`func (o *InstanceProjectionMetadata) GetCanonicalHashOk() (*string, bool)`

GetCanonicalHashOk returns a tuple with the CanonicalHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalHash

`func (o *InstanceProjectionMetadata) SetCanonicalHash(v string)`

SetCanonicalHash sets CanonicalHash field to given value.


### GetLastSyncedAt

`func (o *InstanceProjectionMetadata) GetLastSyncedAt() time.Time`

GetLastSyncedAt returns the LastSyncedAt field if non-nil, zero value otherwise.

### GetLastSyncedAtOk

`func (o *InstanceProjectionMetadata) GetLastSyncedAtOk() (*time.Time, bool)`

GetLastSyncedAtOk returns a tuple with the LastSyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncedAt

`func (o *InstanceProjectionMetadata) SetLastSyncedAt(v time.Time)`

SetLastSyncedAt sets LastSyncedAt field to given value.


### GetSyncState

`func (o *InstanceProjectionMetadata) GetSyncState() string`

GetSyncState returns the SyncState field if non-nil, zero value otherwise.

### GetSyncStateOk

`func (o *InstanceProjectionMetadata) GetSyncStateOk() (*string, bool)`

GetSyncStateOk returns a tuple with the SyncState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncState

`func (o *InstanceProjectionMetadata) SetSyncState(v string)`

SetSyncState sets SyncState field to given value.


### GetAgentsComplete

`func (o *InstanceProjectionMetadata) GetAgentsComplete() bool`

GetAgentsComplete returns the AgentsComplete field if non-nil, zero value otherwise.

### GetAgentsCompleteOk

`func (o *InstanceProjectionMetadata) GetAgentsCompleteOk() (*bool, bool)`

GetAgentsCompleteOk returns a tuple with the AgentsComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentsComplete

`func (o *InstanceProjectionMetadata) SetAgentsComplete(v bool)`

SetAgentsComplete sets AgentsComplete field to given value.


### GetSkillsComplete

`func (o *InstanceProjectionMetadata) GetSkillsComplete() bool`

GetSkillsComplete returns the SkillsComplete field if non-nil, zero value otherwise.

### GetSkillsCompleteOk

`func (o *InstanceProjectionMetadata) GetSkillsCompleteOk() (*bool, bool)`

GetSkillsCompleteOk returns a tuple with the SkillsComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillsComplete

`func (o *InstanceProjectionMetadata) SetSkillsComplete(v bool)`

SetSkillsComplete sets SkillsComplete field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


