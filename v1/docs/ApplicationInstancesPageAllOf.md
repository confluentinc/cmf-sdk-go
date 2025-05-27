# ApplicationInstancesPageAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ApplicationInstancesPageMetadata**](ApplicationInstancesPageMetadata.md) |  | [optional] 
**Items** | Pointer to [**[]FlinkApplicationInstance**](FlinkApplicationInstance.md) |  | [optional] [default to []]

## Methods

### NewApplicationInstancesPageAllOf

`func NewApplicationInstancesPageAllOf() *ApplicationInstancesPageAllOf`

NewApplicationInstancesPageAllOf instantiates a new ApplicationInstancesPageAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationInstancesPageAllOfWithDefaults

`func NewApplicationInstancesPageAllOfWithDefaults() *ApplicationInstancesPageAllOf`

NewApplicationInstancesPageAllOfWithDefaults instantiates a new ApplicationInstancesPageAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *ApplicationInstancesPageAllOf) GetMetadata() ApplicationInstancesPageMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ApplicationInstancesPageAllOf) GetMetadataOk() (*ApplicationInstancesPageMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ApplicationInstancesPageAllOf) SetMetadata(v ApplicationInstancesPageMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ApplicationInstancesPageAllOf) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetItems

`func (o *ApplicationInstancesPageAllOf) GetItems() []FlinkApplicationInstance`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ApplicationInstancesPageAllOf) GetItemsOk() (*[]FlinkApplicationInstance, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ApplicationInstancesPageAllOf) SetItems(v []FlinkApplicationInstance)`

SetItems sets Items field to given value.

### HasItems

`func (o *ApplicationInstancesPageAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


