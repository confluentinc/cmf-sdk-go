# SecretAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**SecretMetadata**](SecretMetadata.md) |  | 
**Spec** | [**SecretSpec**](SecretSpec.md) |  | 
**Status** | Pointer to [**SecretStatus**](SecretStatus.md) |  | [optional] 

## Methods

### NewSecretAllOf

`func NewSecretAllOf(metadata SecretMetadata, spec SecretSpec, ) *SecretAllOf`

NewSecretAllOf instantiates a new SecretAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretAllOfWithDefaults

`func NewSecretAllOfWithDefaults() *SecretAllOf`

NewSecretAllOfWithDefaults instantiates a new SecretAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *SecretAllOf) GetMetadata() SecretMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SecretAllOf) GetMetadataOk() (*SecretMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SecretAllOf) SetMetadata(v SecretMetadata)`

SetMetadata sets Metadata field to given value.


### GetSpec

`func (o *SecretAllOf) GetSpec() SecretSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *SecretAllOf) GetSpecOk() (*SecretSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *SecretAllOf) SetSpec(v SecretSpec)`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *SecretAllOf) GetStatus() SecretStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SecretAllOf) GetStatusOk() (*SecretStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SecretAllOf) SetStatus(v SecretStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SecretAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


