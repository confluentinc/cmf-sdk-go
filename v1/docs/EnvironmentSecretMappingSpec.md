# EnvironmentSecretMappingSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecretName** | **string** | Name of the secret to be mapped to the connection secret id of this mapping. | 

## Methods

### NewEnvironmentSecretMappingSpec

`func NewEnvironmentSecretMappingSpec(secretName string, ) *EnvironmentSecretMappingSpec`

NewEnvironmentSecretMappingSpec instantiates a new EnvironmentSecretMappingSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentSecretMappingSpecWithDefaults

`func NewEnvironmentSecretMappingSpecWithDefaults() *EnvironmentSecretMappingSpec`

NewEnvironmentSecretMappingSpecWithDefaults instantiates a new EnvironmentSecretMappingSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecretName

`func (o *EnvironmentSecretMappingSpec) GetSecretName() string`

GetSecretName returns the SecretName field if non-nil, zero value otherwise.

### GetSecretNameOk

`func (o *EnvironmentSecretMappingSpec) GetSecretNameOk() (*string, bool)`

GetSecretNameOk returns a tuple with the SecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretName

`func (o *EnvironmentSecretMappingSpec) SetSecretName(v string)`

SetSecretName sets SecretName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


