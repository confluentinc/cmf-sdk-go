# EventDataReasonMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** | The FKO event reason (e.g., ValidationError, CrashLoopBackOff, Missing, ScalingReport, AutoscalerError). The event category is conveyed by status.type. | [optional] 
**Message** | Pointer to **string** | The full raw K8s event message content | [optional] 

## Methods

### NewEventDataReasonMessage

`func NewEventDataReasonMessage() *EventDataReasonMessage`

NewEventDataReasonMessage instantiates a new EventDataReasonMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventDataReasonMessageWithDefaults

`func NewEventDataReasonMessageWithDefaults() *EventDataReasonMessage`

NewEventDataReasonMessageWithDefaults instantiates a new EventDataReasonMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *EventDataReasonMessage) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *EventDataReasonMessage) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *EventDataReasonMessage) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *EventDataReasonMessage) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetMessage

`func (o *EventDataReasonMessage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EventDataReasonMessage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EventDataReasonMessage) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *EventDataReasonMessage) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


