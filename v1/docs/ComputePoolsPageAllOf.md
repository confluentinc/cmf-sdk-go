# ComputePoolsPageAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ComputePoolPageMetadata**](ComputePoolPageMetadata.md) |  | [optional] 
**Items** | Pointer to [**[]ComputePool**](ComputePool.md) |  | [optional] [default to []]

## Methods

### NewComputePoolsPageAllOf

`func NewComputePoolsPageAllOf() *ComputePoolsPageAllOf`

NewComputePoolsPageAllOf instantiates a new ComputePoolsPageAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputePoolsPageAllOfWithDefaults

`func NewComputePoolsPageAllOfWithDefaults() *ComputePoolsPageAllOf`

NewComputePoolsPageAllOfWithDefaults instantiates a new ComputePoolsPageAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *ComputePoolsPageAllOf) GetMetadata() ComputePoolPageMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ComputePoolsPageAllOf) GetMetadataOk() (*ComputePoolPageMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ComputePoolsPageAllOf) SetMetadata(v ComputePoolPageMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ComputePoolsPageAllOf) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetItems

`func (o *ComputePoolsPageAllOf) GetItems() []ComputePool`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ComputePoolsPageAllOf) GetItemsOk() (*[]ComputePool, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ComputePoolsPageAllOf) SetItems(v []ComputePool)`

SetItems sets Items field to given value.

### HasItems

`func (o *ComputePoolsPageAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


