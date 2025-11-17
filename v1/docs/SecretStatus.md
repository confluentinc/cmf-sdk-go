# SecretStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | The version of the secret | [optional] 
**Environments** | Pointer to **[]string** | The environments to which the secret is attached to. | [optional] 

## Methods

### NewSecretStatus

`func NewSecretStatus() *SecretStatus`

NewSecretStatus instantiates a new SecretStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretStatusWithDefaults

`func NewSecretStatusWithDefaults() *SecretStatus`

NewSecretStatusWithDefaults instantiates a new SecretStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *SecretStatus) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SecretStatus) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SecretStatus) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SecretStatus) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetEnvironments

`func (o *SecretStatus) GetEnvironments() []string`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *SecretStatus) GetEnvironmentsOk() (*[]string, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *SecretStatus) SetEnvironments(v []string)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *SecretStatus) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


