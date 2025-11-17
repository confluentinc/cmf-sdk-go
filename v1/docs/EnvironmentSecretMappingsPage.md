# EnvironmentSecretMappingsPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pageable** | Pointer to [**Pageable**](Pageable.md) |  | [optional] 
**Metadata** | Pointer to [**EnvironmentSecretMappingsPageMetadata**](EnvironmentSecretMappingsPageMetadata.md) |  | [optional] 
**Items** | Pointer to [**[]EnvironmentSecretMapping**](EnvironmentSecretMapping.md) |  | [optional] [default to []]

## Methods

### NewEnvironmentSecretMappingsPage

`func NewEnvironmentSecretMappingsPage() *EnvironmentSecretMappingsPage`

NewEnvironmentSecretMappingsPage instantiates a new EnvironmentSecretMappingsPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentSecretMappingsPageWithDefaults

`func NewEnvironmentSecretMappingsPageWithDefaults() *EnvironmentSecretMappingsPage`

NewEnvironmentSecretMappingsPageWithDefaults instantiates a new EnvironmentSecretMappingsPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageable

`func (o *EnvironmentSecretMappingsPage) GetPageable() Pageable`

GetPageable returns the Pageable field if non-nil, zero value otherwise.

### GetPageableOk

`func (o *EnvironmentSecretMappingsPage) GetPageableOk() (*Pageable, bool)`

GetPageableOk returns a tuple with the Pageable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageable

`func (o *EnvironmentSecretMappingsPage) SetPageable(v Pageable)`

SetPageable sets Pageable field to given value.

### HasPageable

`func (o *EnvironmentSecretMappingsPage) HasPageable() bool`

HasPageable returns a boolean if a field has been set.

### GetMetadata

`func (o *EnvironmentSecretMappingsPage) GetMetadata() EnvironmentSecretMappingsPageMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EnvironmentSecretMappingsPage) GetMetadataOk() (*EnvironmentSecretMappingsPageMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EnvironmentSecretMappingsPage) SetMetadata(v EnvironmentSecretMappingsPageMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EnvironmentSecretMappingsPage) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetItems

`func (o *EnvironmentSecretMappingsPage) GetItems() []EnvironmentSecretMapping`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *EnvironmentSecretMappingsPage) GetItemsOk() (*[]EnvironmentSecretMapping, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *EnvironmentSecretMappingsPage) SetItems(v []EnvironmentSecretMapping)`

SetItems sets Items field to given value.

### HasItems

`func (o *EnvironmentSecretMappingsPage) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


