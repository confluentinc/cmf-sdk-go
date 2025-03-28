# SecretSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]string** | Data of the secret | [optional] 

## Methods

### NewSecretSpec

`func NewSecretSpec() *SecretSpec`

NewSecretSpec instantiates a new SecretSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretSpecWithDefaults

`func NewSecretSpecWithDefaults() *SecretSpec`

NewSecretSpecWithDefaults instantiates a new SecretSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SecretSpec) GetData() map[string]string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SecretSpec) GetDataOk() (*map[string]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SecretSpec) SetData(v map[string]string)`

SetData sets Data field to given value.

### HasData

`func (o *SecretSpec) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


