# ContinueTaskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Input** | Pointer to **interface{}** |  | [optional] 
**AuthGrant** | Pointer to **bool** | Switches the wire envelope from &#x60;user.continue&#x60; to &#x60;user.auth_grant&#x60; — used to resume an &#x60;auth_required&#x60; pause (OAuth-style permission grant).  | [optional] 

## Methods

### NewContinueTaskRequest

`func NewContinueTaskRequest() *ContinueTaskRequest`

NewContinueTaskRequest instantiates a new ContinueTaskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContinueTaskRequestWithDefaults

`func NewContinueTaskRequestWithDefaults() *ContinueTaskRequest`

NewContinueTaskRequestWithDefaults instantiates a new ContinueTaskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInput

`func (o *ContinueTaskRequest) GetInput() interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ContinueTaskRequest) GetInputOk() (*interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ContinueTaskRequest) SetInput(v interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *ContinueTaskRequest) HasInput() bool`

HasInput returns a boolean if a field has been set.

### SetInputNil

`func (o *ContinueTaskRequest) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *ContinueTaskRequest) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetAuthGrant

`func (o *ContinueTaskRequest) GetAuthGrant() bool`

GetAuthGrant returns the AuthGrant field if non-nil, zero value otherwise.

### GetAuthGrantOk

`func (o *ContinueTaskRequest) GetAuthGrantOk() (*bool, bool)`

GetAuthGrantOk returns a tuple with the AuthGrant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthGrant

`func (o *ContinueTaskRequest) SetAuthGrant(v bool)`

SetAuthGrant sets AuthGrant field to given value.

### HasAuthGrant

`func (o *ContinueTaskRequest) HasAuthGrant() bool`

HasAuthGrant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


