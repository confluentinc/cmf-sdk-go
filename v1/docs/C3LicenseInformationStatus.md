# C3LicenseInformationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | License validity status: VALID, EXPIRED, or INVALID | [optional] 
**LicenseType** | Pointer to **string** | Type of the Confluent license (e.g. ENTERPRISE, TRIAL, FREE_TIER) | [optional] 
**Expiration** | Pointer to **int64** | License expiration as Unix epoch timestamp in milliseconds (milliseconds since 1970-01-01T00:00:00Z) | [optional] 

## Methods

### NewC3LicenseInformationStatus

`func NewC3LicenseInformationStatus() *C3LicenseInformationStatus`

NewC3LicenseInformationStatus instantiates a new C3LicenseInformationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC3LicenseInformationStatusWithDefaults

`func NewC3LicenseInformationStatusWithDefaults() *C3LicenseInformationStatus`

NewC3LicenseInformationStatusWithDefaults instantiates a new C3LicenseInformationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *C3LicenseInformationStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *C3LicenseInformationStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *C3LicenseInformationStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *C3LicenseInformationStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLicenseType

`func (o *C3LicenseInformationStatus) GetLicenseType() string`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *C3LicenseInformationStatus) GetLicenseTypeOk() (*string, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *C3LicenseInformationStatus) SetLicenseType(v string)`

SetLicenseType sets LicenseType field to given value.

### HasLicenseType

`func (o *C3LicenseInformationStatus) HasLicenseType() bool`

HasLicenseType returns a boolean if a field has been set.

### GetExpiration

`func (o *C3LicenseInformationStatus) GetExpiration() int64`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *C3LicenseInformationStatus) GetExpirationOk() (*int64, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *C3LicenseInformationStatus) SetExpiration(v int64)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *C3LicenseInformationStatus) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


