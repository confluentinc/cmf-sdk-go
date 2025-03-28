# SecretsPageAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**SecretsPageMetadata**](SecretsPageMetadata.md) |  | [optional] 
**Items** | Pointer to [**[]Secret**](Secret.md) |  | [optional] [default to []]

## Methods

### NewSecretsPageAllOf

`func NewSecretsPageAllOf() *SecretsPageAllOf`

NewSecretsPageAllOf instantiates a new SecretsPageAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretsPageAllOfWithDefaults

`func NewSecretsPageAllOfWithDefaults() *SecretsPageAllOf`

NewSecretsPageAllOfWithDefaults instantiates a new SecretsPageAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *SecretsPageAllOf) GetMetadata() SecretsPageMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SecretsPageAllOf) GetMetadataOk() (*SecretsPageMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SecretsPageAllOf) SetMetadata(v SecretsPageMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SecretsPageAllOf) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetItems

`func (o *SecretsPageAllOf) GetItems() []Secret`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SecretsPageAllOf) GetItemsOk() (*[]Secret, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SecretsPageAllOf) SetItems(v []Secret)`

SetItems sets Items field to given value.

### HasItems

`func (o *SecretsPageAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


