# EventData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewStatus** | Pointer to **string** | The new status | [optional] 
**ExceptionString** | Pointer to **string** | The full exception string from the Flink job | [optional] 

## Methods

### NewEventData

`func NewEventData() *EventData`

NewEventData instantiates a new EventData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventDataWithDefaults

`func NewEventDataWithDefaults() *EventData`

NewEventDataWithDefaults instantiates a new EventData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewStatus

`func (o *EventData) GetNewStatus() string`

GetNewStatus returns the NewStatus field if non-nil, zero value otherwise.

### GetNewStatusOk

`func (o *EventData) GetNewStatusOk() (*string, bool)`

GetNewStatusOk returns a tuple with the NewStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewStatus

`func (o *EventData) SetNewStatus(v string)`

SetNewStatus sets NewStatus field to given value.

### HasNewStatus

`func (o *EventData) HasNewStatus() bool`

HasNewStatus returns a boolean if a field has been set.

### GetExceptionString

`func (o *EventData) GetExceptionString() string`

GetExceptionString returns the ExceptionString field if non-nil, zero value otherwise.

### GetExceptionStringOk

`func (o *EventData) GetExceptionStringOk() (*string, bool)`

GetExceptionStringOk returns a tuple with the ExceptionString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionString

`func (o *EventData) SetExceptionString(v string)`

SetExceptionString sets ExceptionString field to given value.

### HasExceptionString

`func (o *EventData) HasExceptionString() bool`

HasExceptionString returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


