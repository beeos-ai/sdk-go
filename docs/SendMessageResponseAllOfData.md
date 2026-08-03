# SendMessageResponseAllOfData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | **string** |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSendMessageResponseAllOfData

`func NewSendMessageResponseAllOfData(messageId string, ) *SendMessageResponseAllOfData`

NewSendMessageResponseAllOfData instantiates a new SendMessageResponseAllOfData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendMessageResponseAllOfDataWithDefaults

`func NewSendMessageResponseAllOfDataWithDefaults() *SendMessageResponseAllOfData`

NewSendMessageResponseAllOfDataWithDefaults instantiates a new SendMessageResponseAllOfData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageId

`func (o *SendMessageResponseAllOfData) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *SendMessageResponseAllOfData) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *SendMessageResponseAllOfData) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.


### GetCreatedAt

`func (o *SendMessageResponseAllOfData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SendMessageResponseAllOfData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SendMessageResponseAllOfData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SendMessageResponseAllOfData) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


