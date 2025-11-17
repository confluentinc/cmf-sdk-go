# ComputePoolsPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pageable** | Pointer to [**Pageable**](Pageable.md) |  | [optional] 
**Metadata** | Pointer to [**ComputePoolPageMetadata**](ComputePoolPageMetadata.md) |  | [optional] 
**Items** | Pointer to [**[]ComputePool**](ComputePool.md) |  | [optional] [default to []]

## Methods

### NewComputePoolsPage

`func NewComputePoolsPage() *ComputePoolsPage`

NewComputePoolsPage instantiates a new ComputePoolsPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputePoolsPageWithDefaults

`func NewComputePoolsPageWithDefaults() *ComputePoolsPage`

NewComputePoolsPageWithDefaults instantiates a new ComputePoolsPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageable

`func (o *ComputePoolsPage) GetPageable() Pageable`

GetPageable returns the Pageable field if non-nil, zero value otherwise.

### GetPageableOk

`func (o *ComputePoolsPage) GetPageableOk() (*Pageable, bool)`

GetPageableOk returns a tuple with the Pageable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageable

`func (o *ComputePoolsPage) SetPageable(v Pageable)`

SetPageable sets Pageable field to given value.

### HasPageable

`func (o *ComputePoolsPage) HasPageable() bool`

HasPageable returns a boolean if a field has been set.

### GetMetadata

`func (o *ComputePoolsPage) GetMetadata() ComputePoolPageMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ComputePoolsPage) GetMetadataOk() (*ComputePoolPageMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ComputePoolsPage) SetMetadata(v ComputePoolPageMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ComputePoolsPage) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetItems

`func (o *ComputePoolsPage) GetItems() []ComputePool`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ComputePoolsPage) GetItemsOk() (*[]ComputePool, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ComputePoolsPage) SetItems(v []ComputePool)`

SetItems sets Items field to given value.

### HasItems

`func (o *ComputePoolsPage) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


