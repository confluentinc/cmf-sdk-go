# ResourceEventStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Human readable status message. | [optional] 
**Type** | Pointer to **string** | Type of the event | [optional] 
**Severity** | Pointer to **string** | Severity of the event. Currently &#x60;INFO&#x60; or &#x60;WARNING&#x60;; other values may be added later. &#x60;WARNING&#x60; indicates a problem, such as a Kubernetes Warning event or a change into a failing or failed state. | [optional] 
**Count** | Pointer to **int32** | Number of times this event has occurred. 1 for events that are not aggregated (e.g. status events). | [optional] 
**Data** | Pointer to [**EventData**](EventData.md) |  | [optional] 

## Methods

### NewResourceEventStatus

`func NewResourceEventStatus() *ResourceEventStatus`

NewResourceEventStatus instantiates a new ResourceEventStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceEventStatusWithDefaults

`func NewResourceEventStatusWithDefaults() *ResourceEventStatus`

NewResourceEventStatusWithDefaults instantiates a new ResourceEventStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ResourceEventStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ResourceEventStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ResourceEventStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ResourceEventStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetType

`func (o *ResourceEventStatus) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResourceEventStatus) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResourceEventStatus) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResourceEventStatus) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSeverity

`func (o *ResourceEventStatus) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ResourceEventStatus) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ResourceEventStatus) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ResourceEventStatus) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetCount

`func (o *ResourceEventStatus) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ResourceEventStatus) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ResourceEventStatus) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ResourceEventStatus) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetData

`func (o *ResourceEventStatus) GetData() EventData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResourceEventStatus) GetDataOk() (*EventData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResourceEventStatus) SetData(v EventData)`

SetData sets Data field to given value.

### HasData

`func (o *ResourceEventStatus) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


