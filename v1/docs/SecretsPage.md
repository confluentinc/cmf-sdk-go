# SecretsPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pageable** | Pointer to [**Pageable**](Pageable.md) |  | [optional] 
**Metadata** | Pointer to [**SecretsPageMetadata**](SecretsPageMetadata.md) |  | [optional] 
**Items** | Pointer to [**[]Secret**](Secret.md) |  | [optional] [default to []]

## Methods

### NewSecretsPage

`func NewSecretsPage() *SecretsPage`

NewSecretsPage instantiates a new SecretsPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretsPageWithDefaults

`func NewSecretsPageWithDefaults() *SecretsPage`

NewSecretsPageWithDefaults instantiates a new SecretsPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageable

`func (o *SecretsPage) GetPageable() Pageable`

GetPageable returns the Pageable field if non-nil, zero value otherwise.

### GetPageableOk

`func (o *SecretsPage) GetPageableOk() (*Pageable, bool)`

GetPageableOk returns a tuple with the Pageable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageable

`func (o *SecretsPage) SetPageable(v Pageable)`

SetPageable sets Pageable field to given value.

### HasPageable

`func (o *SecretsPage) HasPageable() bool`

HasPageable returns a boolean if a field has been set.

### GetMetadata

`func (o *SecretsPage) GetMetadata() SecretsPageMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SecretsPage) GetMetadataOk() (*SecretsPageMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SecretsPage) SetMetadata(v SecretsPageMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SecretsPage) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetItems

`func (o *SecretsPage) GetItems() []Secret`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SecretsPage) GetItemsOk() (*[]Secret, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SecretsPage) SetItems(v []Secret)`

SetItems sets Items field to given value.

### HasItems

`func (o *SecretsPage) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


