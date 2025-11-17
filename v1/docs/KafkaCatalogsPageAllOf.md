# KafkaCatalogsPageAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**CatalogPageMetadata**](CatalogPageMetadata.md) |  | [optional] 
**Items** | Pointer to [**[]KafkaCatalog**](KafkaCatalog.md) |  | [optional] [default to []]

## Methods

### NewKafkaCatalogsPageAllOf

`func NewKafkaCatalogsPageAllOf() *KafkaCatalogsPageAllOf`

NewKafkaCatalogsPageAllOf instantiates a new KafkaCatalogsPageAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaCatalogsPageAllOfWithDefaults

`func NewKafkaCatalogsPageAllOfWithDefaults() *KafkaCatalogsPageAllOf`

NewKafkaCatalogsPageAllOfWithDefaults instantiates a new KafkaCatalogsPageAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *KafkaCatalogsPageAllOf) GetMetadata() CatalogPageMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *KafkaCatalogsPageAllOf) GetMetadataOk() (*CatalogPageMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *KafkaCatalogsPageAllOf) SetMetadata(v CatalogPageMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *KafkaCatalogsPageAllOf) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetItems

`func (o *KafkaCatalogsPageAllOf) GetItems() []KafkaCatalog`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *KafkaCatalogsPageAllOf) GetItemsOk() (*[]KafkaCatalog, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *KafkaCatalogsPageAllOf) SetItems(v []KafkaCatalog)`

SetItems sets Items field to given value.

### HasItems

`func (o *KafkaCatalogsPageAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


