# FlinkClusterPodsPageAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**FlinkClusterPodsPageMetadata**](FlinkClusterPodsPageMetadata.md) |  | [optional] 
**Items** | Pointer to [**[]FlinkClusterPod**](FlinkClusterPod.md) |  | [optional] [default to []]

## Methods

### NewFlinkClusterPodsPageAllOf

`func NewFlinkClusterPodsPageAllOf() *FlinkClusterPodsPageAllOf`

NewFlinkClusterPodsPageAllOf instantiates a new FlinkClusterPodsPageAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlinkClusterPodsPageAllOfWithDefaults

`func NewFlinkClusterPodsPageAllOfWithDefaults() *FlinkClusterPodsPageAllOf`

NewFlinkClusterPodsPageAllOfWithDefaults instantiates a new FlinkClusterPodsPageAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *FlinkClusterPodsPageAllOf) GetMetadata() FlinkClusterPodsPageMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FlinkClusterPodsPageAllOf) GetMetadataOk() (*FlinkClusterPodsPageMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FlinkClusterPodsPageAllOf) SetMetadata(v FlinkClusterPodsPageMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FlinkClusterPodsPageAllOf) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetItems

`func (o *FlinkClusterPodsPageAllOf) GetItems() []FlinkClusterPod`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *FlinkClusterPodsPageAllOf) GetItemsOk() (*[]FlinkClusterPod, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *FlinkClusterPodsPageAllOf) SetItems(v []FlinkClusterPod)`

SetItems sets Items field to given value.

### HasItems

`func (o *FlinkClusterPodsPageAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


