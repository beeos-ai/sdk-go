# ListMessagesResponseAllOfData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | [**[]MessageDTO**](MessageDTO.md) |  | 
**LatestOffset** | **int64** | Channel&#39;s current max offset for \&quot;caught up\&quot; detection. | 
**NextSince** | Pointer to **string** | Public offset cursor to pass as the next request&#39;s since value. | [optional] 
**HasMore** | **bool** | True when another page is available via next_since. | 

## Methods

### NewListMessagesResponseAllOfData

`func NewListMessagesResponseAllOfData(messages []MessageDTO, latestOffset int64, hasMore bool, ) *ListMessagesResponseAllOfData`

NewListMessagesResponseAllOfData instantiates a new ListMessagesResponseAllOfData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMessagesResponseAllOfDataWithDefaults

`func NewListMessagesResponseAllOfDataWithDefaults() *ListMessagesResponseAllOfData`

NewListMessagesResponseAllOfDataWithDefaults instantiates a new ListMessagesResponseAllOfData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *ListMessagesResponseAllOfData) GetMessages() []MessageDTO`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ListMessagesResponseAllOfData) GetMessagesOk() (*[]MessageDTO, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ListMessagesResponseAllOfData) SetMessages(v []MessageDTO)`

SetMessages sets Messages field to given value.


### GetLatestOffset

`func (o *ListMessagesResponseAllOfData) GetLatestOffset() int64`

GetLatestOffset returns the LatestOffset field if non-nil, zero value otherwise.

### GetLatestOffsetOk

`func (o *ListMessagesResponseAllOfData) GetLatestOffsetOk() (*int64, bool)`

GetLatestOffsetOk returns a tuple with the LatestOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestOffset

`func (o *ListMessagesResponseAllOfData) SetLatestOffset(v int64)`

SetLatestOffset sets LatestOffset field to given value.


### GetNextSince

`func (o *ListMessagesResponseAllOfData) GetNextSince() string`

GetNextSince returns the NextSince field if non-nil, zero value otherwise.

### GetNextSinceOk

`func (o *ListMessagesResponseAllOfData) GetNextSinceOk() (*string, bool)`

GetNextSinceOk returns a tuple with the NextSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextSince

`func (o *ListMessagesResponseAllOfData) SetNextSince(v string)`

SetNextSince sets NextSince field to given value.

### HasNextSince

`func (o *ListMessagesResponseAllOfData) HasNextSince() bool`

HasNextSince returns a boolean if a field has been set.

### GetHasMore

`func (o *ListMessagesResponseAllOfData) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *ListMessagesResponseAllOfData) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *ListMessagesResponseAllOfData) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


