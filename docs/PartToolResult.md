# PartToolResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**ToolUseId** | **string** |  | 
**Content** | [**PartToolResultContent**](PartToolResultContent.md) |  | 
**IsError** | Pointer to **bool** | True iff the tool call errored (model should treat as such). | [optional] 

## Methods

### NewPartToolResult

`func NewPartToolResult(type_ string, toolUseId string, content PartToolResultContent, ) *PartToolResult`

NewPartToolResult instantiates a new PartToolResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartToolResultWithDefaults

`func NewPartToolResultWithDefaults() *PartToolResult`

NewPartToolResultWithDefaults instantiates a new PartToolResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PartToolResult) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PartToolResult) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PartToolResult) SetType(v string)`

SetType sets Type field to given value.


### GetToolUseId

`func (o *PartToolResult) GetToolUseId() string`

GetToolUseId returns the ToolUseId field if non-nil, zero value otherwise.

### GetToolUseIdOk

`func (o *PartToolResult) GetToolUseIdOk() (*string, bool)`

GetToolUseIdOk returns a tuple with the ToolUseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolUseId

`func (o *PartToolResult) SetToolUseId(v string)`

SetToolUseId sets ToolUseId field to given value.


### GetContent

`func (o *PartToolResult) GetContent() PartToolResultContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PartToolResult) GetContentOk() (*PartToolResultContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PartToolResult) SetContent(v PartToolResultContent)`

SetContent sets Content field to given value.


### GetIsError

`func (o *PartToolResult) GetIsError() bool`

GetIsError returns the IsError field if non-nil, zero value otherwise.

### GetIsErrorOk

`func (o *PartToolResult) GetIsErrorOk() (*bool, bool)`

GetIsErrorOk returns a tuple with the IsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsError

`func (o *PartToolResult) SetIsError(v bool)`

SetIsError sets IsError field to given value.

### HasIsError

`func (o *PartToolResult) HasIsError() bool`

HasIsError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


