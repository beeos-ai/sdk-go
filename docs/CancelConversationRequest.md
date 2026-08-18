# CancelConversationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetMessageId** | **string** |  | 
**Reason** | Pointer to **string** |  | [optional] 

## Methods

### NewCancelConversationRequest

`func NewCancelConversationRequest(targetMessageId string, ) *CancelConversationRequest`

NewCancelConversationRequest instantiates a new CancelConversationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelConversationRequestWithDefaults

`func NewCancelConversationRequestWithDefaults() *CancelConversationRequest`

NewCancelConversationRequestWithDefaults instantiates a new CancelConversationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetMessageId

`func (o *CancelConversationRequest) GetTargetMessageId() string`

GetTargetMessageId returns the TargetMessageId field if non-nil, zero value otherwise.

### GetTargetMessageIdOk

`func (o *CancelConversationRequest) GetTargetMessageIdOk() (*string, bool)`

GetTargetMessageIdOk returns a tuple with the TargetMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMessageId

`func (o *CancelConversationRequest) SetTargetMessageId(v string)`

SetTargetMessageId sets TargetMessageId field to given value.


### GetReason

`func (o *CancelConversationRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CancelConversationRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CancelConversationRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CancelConversationRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


