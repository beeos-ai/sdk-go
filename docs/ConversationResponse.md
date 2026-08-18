# ConversationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Data** | [**ConversationResponseAllOfData**](ConversationResponseAllOfData.md) |  | 

## Methods

### NewConversationResponse

`func NewConversationResponse(success bool, data ConversationResponseAllOfData, ) *ConversationResponse`

NewConversationResponse instantiates a new ConversationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversationResponseWithDefaults

`func NewConversationResponseWithDefaults() *ConversationResponse`

NewConversationResponseWithDefaults instantiates a new ConversationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ConversationResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ConversationResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ConversationResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *ConversationResponse) GetData() ConversationResponseAllOfData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ConversationResponse) GetDataOk() (*ConversationResponseAllOfData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ConversationResponse) SetData(v ConversationResponseAllOfData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


