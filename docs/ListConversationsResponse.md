# ListConversationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Data** | [**ListConversationsResponseAllOfData**](ListConversationsResponseAllOfData.md) |  | 

## Methods

### NewListConversationsResponse

`func NewListConversationsResponse(success bool, data ListConversationsResponseAllOfData, ) *ListConversationsResponse`

NewListConversationsResponse instantiates a new ListConversationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConversationsResponseWithDefaults

`func NewListConversationsResponseWithDefaults() *ListConversationsResponse`

NewListConversationsResponseWithDefaults instantiates a new ListConversationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ListConversationsResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ListConversationsResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ListConversationsResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *ListConversationsResponse) GetData() ListConversationsResponseAllOfData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListConversationsResponse) GetDataOk() (*ListConversationsResponseAllOfData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListConversationsResponse) SetData(v ListConversationsResponseAllOfData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


