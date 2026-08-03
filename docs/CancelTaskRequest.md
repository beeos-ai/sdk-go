# CancelTaskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** | Free-form caller label propagated on the &#x60;chat_cancel&#x60; envelope payload so the agent can audit-log why its turn was aborted. Empty defaults to &#x60;canceled&#x60;.  | [optional] 

## Methods

### NewCancelTaskRequest

`func NewCancelTaskRequest() *CancelTaskRequest`

NewCancelTaskRequest instantiates a new CancelTaskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelTaskRequestWithDefaults

`func NewCancelTaskRequestWithDefaults() *CancelTaskRequest`

NewCancelTaskRequestWithDefaults instantiates a new CancelTaskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *CancelTaskRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CancelTaskRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CancelTaskRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CancelTaskRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


