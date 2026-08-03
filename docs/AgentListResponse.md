# AgentListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Data** | [**[]AgentDTO**](AgentDTO.md) |  | 
**Total** | **int64** |  | 

## Methods

### NewAgentListResponse

`func NewAgentListResponse(success bool, data []AgentDTO, total int64, ) *AgentListResponse`

NewAgentListResponse instantiates a new AgentListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentListResponseWithDefaults

`func NewAgentListResponseWithDefaults() *AgentListResponse`

NewAgentListResponseWithDefaults instantiates a new AgentListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *AgentListResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AgentListResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AgentListResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *AgentListResponse) GetData() []AgentDTO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AgentListResponse) GetDataOk() (*[]AgentDTO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AgentListResponse) SetData(v []AgentDTO)`

SetData sets Data field to given value.


### GetTotal

`func (o *AgentListResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AgentListResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AgentListResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


