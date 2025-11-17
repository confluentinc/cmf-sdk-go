# KafkaCatalogsPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pageable** | Pointer to [**Pageable**](Pageable.md) |  | [optional] 
**Metadata** | Pointer to [**CatalogPageMetadata**](CatalogPageMetadata.md) |  | [optional] 
**Items** | Pointer to [**[]KafkaCatalog**](KafkaCatalog.md) |  | [optional] [default to []]

## Methods

### NewKafkaCatalogsPage

`func NewKafkaCatalogsPage() *KafkaCatalogsPage`

NewKafkaCatalogsPage instantiates a new KafkaCatalogsPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaCatalogsPageWithDefaults

`func NewKafkaCatalogsPageWithDefaults() *KafkaCatalogsPage`

NewKafkaCatalogsPageWithDefaults instantiates a new KafkaCatalogsPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageable

`func (o *KafkaCatalogsPage) GetPageable() Pageable`

GetPageable returns the Pageable field if non-nil, zero value otherwise.

### GetPageableOk

`func (o *KafkaCatalogsPage) GetPageableOk() (*Pageable, bool)`

GetPageableOk returns a tuple with the Pageable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageable

`func (o *KafkaCatalogsPage) SetPageable(v Pageable)`

SetPageable sets Pageable field to given value.

### HasPageable

`func (o *KafkaCatalogsPage) HasPageable() bool`

HasPageable returns a boolean if a field has been set.

### GetMetadata

`func (o *KafkaCatalogsPage) GetMetadata() CatalogPageMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *KafkaCatalogsPage) GetMetadataOk() (*CatalogPageMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *KafkaCatalogsPage) SetMetadata(v CatalogPageMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *KafkaCatalogsPage) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetItems

`func (o *KafkaCatalogsPage) GetItems() []KafkaCatalog`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *KafkaCatalogsPage) GetItemsOk() (*[]KafkaCatalog, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *KafkaCatalogsPage) SetItems(v []KafkaCatalog)`

SetItems sets Items field to given value.

### HasItems

`func (o *KafkaCatalogsPage) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


