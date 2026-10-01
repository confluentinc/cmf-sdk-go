# ServiceAccountAllowlistEntryStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | Pointer to **string** | How the entry came to be on the allowlist. &#x60;SEED&#x60; for entries grandfathered from resources already using the account; &#x60;USER_ALLOWED&#x60; for entries added deliberately by a user. | [optional] 

## Methods

### NewServiceAccountAllowlistEntryStatus

`func NewServiceAccountAllowlistEntryStatus() *ServiceAccountAllowlistEntryStatus`

NewServiceAccountAllowlistEntryStatus instantiates a new ServiceAccountAllowlistEntryStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccountAllowlistEntryStatusWithDefaults

`func NewServiceAccountAllowlistEntryStatusWithDefaults() *ServiceAccountAllowlistEntryStatus`

NewServiceAccountAllowlistEntryStatusWithDefaults instantiates a new ServiceAccountAllowlistEntryStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *ServiceAccountAllowlistEntryStatus) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ServiceAccountAllowlistEntryStatus) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ServiceAccountAllowlistEntryStatus) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ServiceAccountAllowlistEntryStatus) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


