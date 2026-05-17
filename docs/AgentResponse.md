# AgentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Data** | [**AgentDTO**](AgentDTO.md) |  | 

## Methods

### NewAgentResponse

`func NewAgentResponse(success bool, data AgentDTO, ) *AgentResponse`

NewAgentResponse instantiates a new AgentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentResponseWithDefaults

`func NewAgentResponseWithDefaults() *AgentResponse`

NewAgentResponseWithDefaults instantiates a new AgentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *AgentResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AgentResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AgentResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *AgentResponse) GetData() AgentDTO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AgentResponse) GetDataOk() (*AgentDTO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AgentResponse) SetData(v AgentDTO)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


