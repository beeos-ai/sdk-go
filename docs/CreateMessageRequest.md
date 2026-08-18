# CreateMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Message intent. Defaults to &#x60;agent_reply&#x60; in the SDK. | 
**ReplyTo** | Pointer to **string** |  | [optional] 
**Body** | Pointer to **string** |  | [optional] 
**Parts** | Pointer to [**[]Part**](Part.md) |  | [optional] 
**State** | Pointer to [**MessageState**](MessageState.md) |  | [optional] 
**StopReason** | Pointer to [**StopReason**](StopReason.md) |  | [optional] 
**RequireSubscriber** | Pointer to **bool** | When true, the POST fails with 410 no_subscriber if no peer is online. | [optional] 

## Methods

### NewCreateMessageRequest

`func NewCreateMessageRequest(type_ string, ) *CreateMessageRequest`

NewCreateMessageRequest instantiates a new CreateMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMessageRequestWithDefaults

`func NewCreateMessageRequestWithDefaults() *CreateMessageRequest`

NewCreateMessageRequestWithDefaults instantiates a new CreateMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateMessageRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateMessageRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateMessageRequest) SetType(v string)`

SetType sets Type field to given value.


### GetReplyTo

`func (o *CreateMessageRequest) GetReplyTo() string`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *CreateMessageRequest) GetReplyToOk() (*string, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *CreateMessageRequest) SetReplyTo(v string)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *CreateMessageRequest) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetBody

`func (o *CreateMessageRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *CreateMessageRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *CreateMessageRequest) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *CreateMessageRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetParts

`func (o *CreateMessageRequest) GetParts() []Part`

GetParts returns the Parts field if non-nil, zero value otherwise.

### GetPartsOk

`func (o *CreateMessageRequest) GetPartsOk() (*[]Part, bool)`

GetPartsOk returns a tuple with the Parts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParts

`func (o *CreateMessageRequest) SetParts(v []Part)`

SetParts sets Parts field to given value.

### HasParts

`func (o *CreateMessageRequest) HasParts() bool`

HasParts returns a boolean if a field has been set.

### GetState

`func (o *CreateMessageRequest) GetState() MessageState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CreateMessageRequest) GetStateOk() (*MessageState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CreateMessageRequest) SetState(v MessageState)`

SetState sets State field to given value.

### HasState

`func (o *CreateMessageRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStopReason

`func (o *CreateMessageRequest) GetStopReason() StopReason`

GetStopReason returns the StopReason field if non-nil, zero value otherwise.

### GetStopReasonOk

`func (o *CreateMessageRequest) GetStopReasonOk() (*StopReason, bool)`

GetStopReasonOk returns a tuple with the StopReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopReason

`func (o *CreateMessageRequest) SetStopReason(v StopReason)`

SetStopReason sets StopReason field to given value.

### HasStopReason

`func (o *CreateMessageRequest) HasStopReason() bool`

HasStopReason returns a boolean if a field has been set.

### GetRequireSubscriber

`func (o *CreateMessageRequest) GetRequireSubscriber() bool`

GetRequireSubscriber returns the RequireSubscriber field if non-nil, zero value otherwise.

### GetRequireSubscriberOk

`func (o *CreateMessageRequest) GetRequireSubscriberOk() (*bool, bool)`

GetRequireSubscriberOk returns a tuple with the RequireSubscriber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSubscriber

`func (o *CreateMessageRequest) SetRequireSubscriber(v bool)`

SetRequireSubscriber sets RequireSubscriber field to given value.

### HasRequireSubscriber

`func (o *CreateMessageRequest) HasRequireSubscriber() bool`

HasRequireSubscriber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


