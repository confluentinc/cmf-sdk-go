# EnvironmentSecretMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** |  | 
**Kind** | **string** |  | 
**Metadata** | Pointer to [**EnvironmentSecretMappingMetadata**](EnvironmentSecretMappingMetadata.md) |  | [optional] 
**Spec** | Pointer to [**EnvironmentSecretMappingSpec**](EnvironmentSecretMappingSpec.md) |  | [optional] 

## Methods

### NewEnvironmentSecretMapping

`func NewEnvironmentSecretMapping(apiVersion string, kind string, ) *EnvironmentSecretMapping`

NewEnvironmentSecretMapping instantiates a new EnvironmentSecretMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentSecretMappingWithDefaults

`func NewEnvironmentSecretMappingWithDefaults() *EnvironmentSecretMapping`

NewEnvironmentSecretMappingWithDefaults instantiates a new EnvironmentSecretMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *EnvironmentSecretMapping) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *EnvironmentSecretMapping) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *EnvironmentSecretMapping) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *EnvironmentSecretMapping) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *EnvironmentSecretMapping) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *EnvironmentSecretMapping) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *EnvironmentSecretMapping) GetMetadata() EnvironmentSecretMappingMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EnvironmentSecretMapping) GetMetadataOk() (*EnvironmentSecretMappingMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EnvironmentSecretMapping) SetMetadata(v EnvironmentSecretMappingMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EnvironmentSecretMapping) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *EnvironmentSecretMapping) GetSpec() EnvironmentSecretMappingSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *EnvironmentSecretMapping) GetSpecOk() (*EnvironmentSecretMappingSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *EnvironmentSecretMapping) SetSpec(v EnvironmentSecretMappingSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *EnvironmentSecretMapping) HasSpec() bool`

HasSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


