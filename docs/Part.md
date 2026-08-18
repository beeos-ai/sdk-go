# Part

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Text** | **string** |  | 
**State** | Pointer to [**PartState**](PartState.md) |  | [optional] 
**Id** | **string** | Unique tool-call id; matches the &#x60;tool_use_id&#x60; on the corresponding tool_result. | 
**Name** | **string** |  | 
**Arguments** | Pointer to **interface{}** |  | [optional] 
**ToolUseId** | **string** |  | 
**Content** | [**PartToolResultContent**](PartToolResultContent.md) |  | 
**IsError** | Pointer to **bool** | True iff the tool call errored (model should treat as such). | [optional] 
**Url** | **string** |  | 
**MimeType** | **string** |  | 
**Size** | Pointer to **int64** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Snippet** | Pointer to **string** |  | [optional] 
**Kind** | **string** | Caller-controlled extension discriminant. | 
**Data** | **interface{}** |  | 

## Methods

### NewPart

`func NewPart(type_ string, text string, id string, name string, toolUseId string, content PartToolResultContent, url string, mimeType string, kind string, data interface{}, ) *Part`

NewPart instantiates a new Part object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartWithDefaults

`func NewPartWithDefaults() *Part`

NewPartWithDefaults instantiates a new Part object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Part) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Part) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Part) SetType(v string)`

SetType sets Type field to given value.


### GetText

`func (o *Part) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *Part) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *Part) SetText(v string)`

SetText sets Text field to given value.


### GetState

`func (o *Part) GetState() PartState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Part) GetStateOk() (*PartState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Part) SetState(v PartState)`

SetState sets State field to given value.

### HasState

`func (o *Part) HasState() bool`

HasState returns a boolean if a field has been set.

### GetId

`func (o *Part) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Part) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Part) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Part) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Part) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Part) SetName(v string)`

SetName sets Name field to given value.


### GetArguments

`func (o *Part) GetArguments() interface{}`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *Part) GetArgumentsOk() (*interface{}, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *Part) SetArguments(v interface{})`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *Part) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### SetArgumentsNil

`func (o *Part) SetArgumentsNil(b bool)`

 SetArgumentsNil sets the value for Arguments to be an explicit nil

### UnsetArguments
`func (o *Part) UnsetArguments()`

UnsetArguments ensures that no value is present for Arguments, not even an explicit nil
### GetToolUseId

`func (o *Part) GetToolUseId() string`

GetToolUseId returns the ToolUseId field if non-nil, zero value otherwise.

### GetToolUseIdOk

`func (o *Part) GetToolUseIdOk() (*string, bool)`

GetToolUseIdOk returns a tuple with the ToolUseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolUseId

`func (o *Part) SetToolUseId(v string)`

SetToolUseId sets ToolUseId field to given value.


### GetContent

`func (o *Part) GetContent() PartToolResultContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Part) GetContentOk() (*PartToolResultContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Part) SetContent(v PartToolResultContent)`

SetContent sets Content field to given value.


### GetIsError

`func (o *Part) GetIsError() bool`

GetIsError returns the IsError field if non-nil, zero value otherwise.

### GetIsErrorOk

`func (o *Part) GetIsErrorOk() (*bool, bool)`

GetIsErrorOk returns a tuple with the IsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsError

`func (o *Part) SetIsError(v bool)`

SetIsError sets IsError field to given value.

### HasIsError

`func (o *Part) HasIsError() bool`

HasIsError returns a boolean if a field has been set.

### GetUrl

`func (o *Part) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Part) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Part) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetMimeType

`func (o *Part) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *Part) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *Part) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.


### GetSize

`func (o *Part) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Part) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Part) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *Part) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTitle

`func (o *Part) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Part) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Part) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Part) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetSnippet

`func (o *Part) GetSnippet() string`

GetSnippet returns the Snippet field if non-nil, zero value otherwise.

### GetSnippetOk

`func (o *Part) GetSnippetOk() (*string, bool)`

GetSnippetOk returns a tuple with the Snippet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnippet

`func (o *Part) SetSnippet(v string)`

SetSnippet sets Snippet field to given value.

### HasSnippet

`func (o *Part) HasSnippet() bool`

HasSnippet returns a boolean if a field has been set.

### GetKind

`func (o *Part) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Part) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Part) SetKind(v string)`

SetKind sets Kind field to given value.


### GetData

`func (o *Part) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Part) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Part) SetData(v interface{})`

SetData sets Data field to given value.


### SetDataNil

`func (o *Part) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *Part) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


