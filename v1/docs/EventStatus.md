# EventStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Human readable status message. | [optional] 
**Type** | Pointer to **string** | Type of the event | [optional] 
**Data** | Pointer to [**EventData**](EventData.md) |  | [optional] 

## Methods

### NewEventStatus

`func NewEventStatus() *EventStatus`

NewEventStatus instantiates a new EventStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventStatusWithDefaults

`func NewEventStatusWithDefaults() *EventStatus`

NewEventStatusWithDefaults instantiates a new EventStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *EventStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EventStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EventStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *EventStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetType

`func (o *EventStatus) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventStatus) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventStatus) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EventStatus) HasType() bool`

HasType returns a boolean if a field has been set.

### GetData

`func (o *EventStatus) GetData() EventData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EventStatus) GetDataOk() (*EventData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EventStatus) SetData(v EventData)`

SetData sets Data field to given value.

### HasData

`func (o *EventStatus) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


