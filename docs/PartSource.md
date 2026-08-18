# PartSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Url** | **string** |  | 
**Title** | Pointer to **string** |  | [optional] 
**Snippet** | Pointer to **string** |  | [optional] 

## Methods

### NewPartSource

`func NewPartSource(type_ string, url string, ) *PartSource`

NewPartSource instantiates a new PartSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartSourceWithDefaults

`func NewPartSourceWithDefaults() *PartSource`

NewPartSourceWithDefaults instantiates a new PartSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PartSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PartSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PartSource) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *PartSource) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PartSource) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PartSource) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetTitle

`func (o *PartSource) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PartSource) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PartSource) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PartSource) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetSnippet

`func (o *PartSource) GetSnippet() string`

GetSnippet returns the Snippet field if non-nil, zero value otherwise.

### GetSnippetOk

`func (o *PartSource) GetSnippetOk() (*string, bool)`

GetSnippetOk returns a tuple with the Snippet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnippet

`func (o *PartSource) SetSnippet(v string)`

SetSnippet sets Snippet field to given value.

### HasSnippet

`func (o *PartSource) HasSnippet() bool`

HasSnippet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


