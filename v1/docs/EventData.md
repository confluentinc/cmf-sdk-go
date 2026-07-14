# EventData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewStatus** | Pointer to **string** | The new status | [optional] 
**ExceptionString** | Pointer to **string** | The full exception string from the Flink job | [optional] 
**Reason** | Pointer to **string** | The FKO event reason (e.g., ValidationError, CrashLoopBackOff, Missing, ScalingReport, AutoscalerError). The event category is conveyed by status.type. | [optional] 
**Message** | Pointer to **string** | The full raw K8s event message content | [optional] 

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

### GetReason

`func (o *EventData) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *EventData) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *EventData) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *EventData) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetMessage

`func (o *EventData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EventData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EventData) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *EventData) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


