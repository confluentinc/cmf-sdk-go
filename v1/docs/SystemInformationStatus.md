# SystemInformationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | Version of the CMF application | [optional] 
**Revision** | Pointer to **string** | Git commit hash of the CMF application | [optional] 

## Methods

### NewSystemInformationStatus

`func NewSystemInformationStatus() *SystemInformationStatus`

NewSystemInformationStatus instantiates a new SystemInformationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemInformationStatusWithDefaults

`func NewSystemInformationStatusWithDefaults() *SystemInformationStatus`

NewSystemInformationStatusWithDefaults instantiates a new SystemInformationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *SystemInformationStatus) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SystemInformationStatus) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SystemInformationStatus) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SystemInformationStatus) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetRevision

`func (o *SystemInformationStatus) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *SystemInformationStatus) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *SystemInformationStatus) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *SystemInformationStatus) HasRevision() bool`

HasRevision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


