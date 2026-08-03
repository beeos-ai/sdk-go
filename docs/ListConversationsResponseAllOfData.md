# ListConversationsResponseAllOfData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conversations** | [**[]ConversationResponseAllOfData**](ConversationResponseAllOfData.md) |  | 
**NextSince** | Pointer to **string** | Pagination cursor. Empty when the caller has reached the end of results.  | [optional] 

## Methods

### NewListConversationsResponseAllOfData

`func NewListConversationsResponseAllOfData(conversations []ConversationResponseAllOfData, ) *ListConversationsResponseAllOfData`

NewListConversationsResponseAllOfData instantiates a new ListConversationsResponseAllOfData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConversationsResponseAllOfDataWithDefaults

`func NewListConversationsResponseAllOfDataWithDefaults() *ListConversationsResponseAllOfData`

NewListConversationsResponseAllOfDataWithDefaults instantiates a new ListConversationsResponseAllOfData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConversations

`func (o *ListConversationsResponseAllOfData) GetConversations() []ConversationResponseAllOfData`

GetConversations returns the Conversations field if non-nil, zero value otherwise.

### GetConversationsOk

`func (o *ListConversationsResponseAllOfData) GetConversationsOk() (*[]ConversationResponseAllOfData, bool)`

GetConversationsOk returns a tuple with the Conversations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversations

`func (o *ListConversationsResponseAllOfData) SetConversations(v []ConversationResponseAllOfData)`

SetConversations sets Conversations field to given value.


### GetNextSince

`func (o *ListConversationsResponseAllOfData) GetNextSince() string`

GetNextSince returns the NextSince field if non-nil, zero value otherwise.

### GetNextSinceOk

`func (o *ListConversationsResponseAllOfData) GetNextSinceOk() (*string, bool)`

GetNextSinceOk returns a tuple with the NextSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextSince

`func (o *ListConversationsResponseAllOfData) SetNextSince(v string)`

SetNextSince sets NextSince field to given value.

### HasNextSince

`func (o *ListConversationsResponseAllOfData) HasNextSince() bool`

HasNextSince returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


