# ResourceEventsPageAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ResourceEventsPageMetadata**](ResourceEventsPageMetadata.md) |  | [optional] 
**Items** | Pointer to [**[]ResourceEvent**](ResourceEvent.md) |  | [optional] [default to []]

## Methods

### NewResourceEventsPageAllOf

`func NewResourceEventsPageAllOf() *ResourceEventsPageAllOf`

NewResourceEventsPageAllOf instantiates a new ResourceEventsPageAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceEventsPageAllOfWithDefaults

`func NewResourceEventsPageAllOfWithDefaults() *ResourceEventsPageAllOf`

NewResourceEventsPageAllOfWithDefaults instantiates a new ResourceEventsPageAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *ResourceEventsPageAllOf) GetMetadata() ResourceEventsPageMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ResourceEventsPageAllOf) GetMetadataOk() (*ResourceEventsPageMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ResourceEventsPageAllOf) SetMetadata(v ResourceEventsPageMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ResourceEventsPageAllOf) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetItems

`func (o *ResourceEventsPageAllOf) GetItems() []ResourceEvent`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ResourceEventsPageAllOf) GetItemsOk() (*[]ResourceEvent, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ResourceEventsPageAllOf) SetItems(v []ResourceEvent)`

SetItems sets Items field to given value.

### HasItems

`func (o *ResourceEventsPageAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


