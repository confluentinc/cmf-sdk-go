# SystemInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**SystemInformationStatus**](SystemInformationStatus.md) |  | [optional] 

## Methods

### NewSystemInformation

`func NewSystemInformation() *SystemInformation`

NewSystemInformation instantiates a new SystemInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemInformationWithDefaults

`func NewSystemInformationWithDefaults() *SystemInformation`

NewSystemInformationWithDefaults instantiates a new SystemInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SystemInformation) GetStatus() SystemInformationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SystemInformation) GetStatusOk() (*SystemInformationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SystemInformation) SetStatus(v SystemInformationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SystemInformation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


