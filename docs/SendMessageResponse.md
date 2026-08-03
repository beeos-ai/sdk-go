# SendMessageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Data** | [**SendMessageResponseAllOfData**](SendMessageResponseAllOfData.md) |  | 

## Methods

### NewSendMessageResponse

`func NewSendMessageResponse(success bool, data SendMessageResponseAllOfData, ) *SendMessageResponse`

NewSendMessageResponse instantiates a new SendMessageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendMessageResponseWithDefaults

`func NewSendMessageResponseWithDefaults() *SendMessageResponse`

NewSendMessageResponseWithDefaults instantiates a new SendMessageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *SendMessageResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *SendMessageResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *SendMessageResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *SendMessageResponse) GetData() SendMessageResponseAllOfData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SendMessageResponse) GetDataOk() (*SendMessageResponseAllOfData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SendMessageResponse) SetData(v SendMessageResponseAllOfData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


