# RuntimeOperationEventEnvelope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**OperationId** | **string** |  | 
**Method** | [**ServiceRuntimeMethod**](ServiceRuntimeMethod.md) |  | 
**Event** | [**RuntimeOperationEventEnvelopeEvent**](RuntimeOperationEventEnvelopeEvent.md) |  | 

## Methods

### NewRuntimeOperationEventEnvelope

`func NewRuntimeOperationEventEnvelope(id string, operationId string, method ServiceRuntimeMethod, event RuntimeOperationEventEnvelopeEvent, ) *RuntimeOperationEventEnvelope`

NewRuntimeOperationEventEnvelope instantiates a new RuntimeOperationEventEnvelope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuntimeOperationEventEnvelopeWithDefaults

`func NewRuntimeOperationEventEnvelopeWithDefaults() *RuntimeOperationEventEnvelope`

NewRuntimeOperationEventEnvelopeWithDefaults instantiates a new RuntimeOperationEventEnvelope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RuntimeOperationEventEnvelope) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RuntimeOperationEventEnvelope) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RuntimeOperationEventEnvelope) SetId(v string)`

SetId sets Id field to given value.


### GetOperationId

`func (o *RuntimeOperationEventEnvelope) GetOperationId() string`

GetOperationId returns the OperationId field if non-nil, zero value otherwise.

### GetOperationIdOk

`func (o *RuntimeOperationEventEnvelope) GetOperationIdOk() (*string, bool)`

GetOperationIdOk returns a tuple with the OperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationId

`func (o *RuntimeOperationEventEnvelope) SetOperationId(v string)`

SetOperationId sets OperationId field to given value.


### GetMethod

`func (o *RuntimeOperationEventEnvelope) GetMethod() ServiceRuntimeMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *RuntimeOperationEventEnvelope) GetMethodOk() (*ServiceRuntimeMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *RuntimeOperationEventEnvelope) SetMethod(v ServiceRuntimeMethod)`

SetMethod sets Method field to given value.


### GetEvent

`func (o *RuntimeOperationEventEnvelope) GetEvent() RuntimeOperationEventEnvelopeEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *RuntimeOperationEventEnvelope) GetEventOk() (*RuntimeOperationEventEnvelopeEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *RuntimeOperationEventEnvelope) SetEvent(v RuntimeOperationEventEnvelopeEvent)`

SetEvent sets Event field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


