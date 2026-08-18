# InvokeRuntimeMethodRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jsonrpc** | **string** |  | 
**Id** | **interface{}** |  | 
**Method** | **string** |  | 
**Params** | **map[string]interface{}** |  | 

## Methods

### NewInvokeRuntimeMethodRequest

`func NewInvokeRuntimeMethodRequest(jsonrpc string, id interface{}, method string, params map[string]interface{}, ) *InvokeRuntimeMethodRequest`

NewInvokeRuntimeMethodRequest instantiates a new InvokeRuntimeMethodRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvokeRuntimeMethodRequestWithDefaults

`func NewInvokeRuntimeMethodRequestWithDefaults() *InvokeRuntimeMethodRequest`

NewInvokeRuntimeMethodRequestWithDefaults instantiates a new InvokeRuntimeMethodRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsonrpc

`func (o *InvokeRuntimeMethodRequest) GetJsonrpc() string`

GetJsonrpc returns the Jsonrpc field if non-nil, zero value otherwise.

### GetJsonrpcOk

`func (o *InvokeRuntimeMethodRequest) GetJsonrpcOk() (*string, bool)`

GetJsonrpcOk returns a tuple with the Jsonrpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonrpc

`func (o *InvokeRuntimeMethodRequest) SetJsonrpc(v string)`

SetJsonrpc sets Jsonrpc field to given value.


### GetId

`func (o *InvokeRuntimeMethodRequest) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvokeRuntimeMethodRequest) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvokeRuntimeMethodRequest) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *InvokeRuntimeMethodRequest) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *InvokeRuntimeMethodRequest) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetMethod

`func (o *InvokeRuntimeMethodRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *InvokeRuntimeMethodRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *InvokeRuntimeMethodRequest) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetParams

`func (o *InvokeRuntimeMethodRequest) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *InvokeRuntimeMethodRequest) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *InvokeRuntimeMethodRequest) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


