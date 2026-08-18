# ServiceOperationCursorConflict

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **interface{}** |  | 
**OperationId** | **string** |  | 
**Cursor** | **string** |  | 
**LastSequence** | **string** |  | 
**HttpStatus** | **interface{}** |  | 

## Methods

### NewServiceOperationCursorConflict

`func NewServiceOperationCursorConflict(code interface{}, operationId string, cursor string, lastSequence string, httpStatus interface{}, ) *ServiceOperationCursorConflict`

NewServiceOperationCursorConflict instantiates a new ServiceOperationCursorConflict object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceOperationCursorConflictWithDefaults

`func NewServiceOperationCursorConflictWithDefaults() *ServiceOperationCursorConflict`

NewServiceOperationCursorConflictWithDefaults instantiates a new ServiceOperationCursorConflict object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ServiceOperationCursorConflict) GetCode() interface{}`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ServiceOperationCursorConflict) GetCodeOk() (*interface{}, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ServiceOperationCursorConflict) SetCode(v interface{})`

SetCode sets Code field to given value.


### SetCodeNil

`func (o *ServiceOperationCursorConflict) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *ServiceOperationCursorConflict) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetOperationId

`func (o *ServiceOperationCursorConflict) GetOperationId() string`

GetOperationId returns the OperationId field if non-nil, zero value otherwise.

### GetOperationIdOk

`func (o *ServiceOperationCursorConflict) GetOperationIdOk() (*string, bool)`

GetOperationIdOk returns a tuple with the OperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationId

`func (o *ServiceOperationCursorConflict) SetOperationId(v string)`

SetOperationId sets OperationId field to given value.


### GetCursor

`func (o *ServiceOperationCursorConflict) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *ServiceOperationCursorConflict) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *ServiceOperationCursorConflict) SetCursor(v string)`

SetCursor sets Cursor field to given value.


### GetLastSequence

`func (o *ServiceOperationCursorConflict) GetLastSequence() string`

GetLastSequence returns the LastSequence field if non-nil, zero value otherwise.

### GetLastSequenceOk

`func (o *ServiceOperationCursorConflict) GetLastSequenceOk() (*string, bool)`

GetLastSequenceOk returns a tuple with the LastSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSequence

`func (o *ServiceOperationCursorConflict) SetLastSequence(v string)`

SetLastSequence sets LastSequence field to given value.


### GetHttpStatus

`func (o *ServiceOperationCursorConflict) GetHttpStatus() interface{}`

GetHttpStatus returns the HttpStatus field if non-nil, zero value otherwise.

### GetHttpStatusOk

`func (o *ServiceOperationCursorConflict) GetHttpStatusOk() (*interface{}, bool)`

GetHttpStatusOk returns a tuple with the HttpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpStatus

`func (o *ServiceOperationCursorConflict) SetHttpStatus(v interface{})`

SetHttpStatus sets HttpStatus field to given value.


### SetHttpStatusNil

`func (o *ServiceOperationCursorConflict) SetHttpStatusNil(b bool)`

 SetHttpStatusNil sets the value for HttpStatus to be an explicit nil

### UnsetHttpStatus
`func (o *ServiceOperationCursorConflict) UnsetHttpStatus()`

UnsetHttpStatus ensures that no value is present for HttpStatus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


