# EnvironmentSecretMappingsPageAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**EnvironmentSecretMappingsPageMetadata**](EnvironmentSecretMappingsPageMetadata.md) |  | [optional] 
**Items** | Pointer to [**[]EnvironmentSecretMapping**](EnvironmentSecretMapping.md) |  | [optional] [default to []]

## Methods

### NewEnvironmentSecretMappingsPageAllOf

`func NewEnvironmentSecretMappingsPageAllOf() *EnvironmentSecretMappingsPageAllOf`

NewEnvironmentSecretMappingsPageAllOf instantiates a new EnvironmentSecretMappingsPageAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentSecretMappingsPageAllOfWithDefaults

`func NewEnvironmentSecretMappingsPageAllOfWithDefaults() *EnvironmentSecretMappingsPageAllOf`

NewEnvironmentSecretMappingsPageAllOfWithDefaults instantiates a new EnvironmentSecretMappingsPageAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *EnvironmentSecretMappingsPageAllOf) GetMetadata() EnvironmentSecretMappingsPageMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EnvironmentSecretMappingsPageAllOf) GetMetadataOk() (*EnvironmentSecretMappingsPageMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EnvironmentSecretMappingsPageAllOf) SetMetadata(v EnvironmentSecretMappingsPageMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EnvironmentSecretMappingsPageAllOf) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetItems

`func (o *EnvironmentSecretMappingsPageAllOf) GetItems() []EnvironmentSecretMapping`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *EnvironmentSecretMappingsPageAllOf) GetItemsOk() (*[]EnvironmentSecretMapping, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *EnvironmentSecretMappingsPageAllOf) SetItems(v []EnvironmentSecretMapping)`

SetItems sets Items field to given value.

### HasItems

`func (o *EnvironmentSecretMappingsPageAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


