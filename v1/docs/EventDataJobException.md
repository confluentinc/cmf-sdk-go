# EventDataJobException

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExceptionString** | Pointer to **string** | The full exception string from the Flink job | [optional] 

## Methods

### NewEventDataJobException

`func NewEventDataJobException() *EventDataJobException`

NewEventDataJobException instantiates a new EventDataJobException object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventDataJobExceptionWithDefaults

`func NewEventDataJobExceptionWithDefaults() *EventDataJobException`

NewEventDataJobExceptionWithDefaults instantiates a new EventDataJobException object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExceptionString

`func (o *EventDataJobException) GetExceptionString() string`

GetExceptionString returns the ExceptionString field if non-nil, zero value otherwise.

### GetExceptionStringOk

`func (o *EventDataJobException) GetExceptionStringOk() (*string, bool)`

GetExceptionStringOk returns a tuple with the ExceptionString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionString

`func (o *EventDataJobException) SetExceptionString(v string)`

SetExceptionString sets ExceptionString field to given value.

### HasExceptionString

`func (o *EventDataJobException) HasExceptionString() bool`

HasExceptionString returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


