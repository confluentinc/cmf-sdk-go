# ResourceEventsPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pageable** | Pointer to [**Pageable**](Pageable.md) |  | [optional] 
**Metadata** | Pointer to [**ResourceEventsPageMetadata**](ResourceEventsPageMetadata.md) |  | [optional] 
**Items** | Pointer to [**[]ResourceEvent**](ResourceEvent.md) |  | [optional] [default to []]

## Methods

### NewResourceEventsPage

`func NewResourceEventsPage() *ResourceEventsPage`

NewResourceEventsPage instantiates a new ResourceEventsPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceEventsPageWithDefaults

`func NewResourceEventsPageWithDefaults() *ResourceEventsPage`

NewResourceEventsPageWithDefaults instantiates a new ResourceEventsPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageable

`func (o *ResourceEventsPage) GetPageable() Pageable`

GetPageable returns the Pageable field if non-nil, zero value otherwise.

### GetPageableOk

`func (o *ResourceEventsPage) GetPageableOk() (*Pageable, bool)`

GetPageableOk returns a tuple with the Pageable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageable

`func (o *ResourceEventsPage) SetPageable(v Pageable)`

SetPageable sets Pageable field to given value.

### HasPageable

`func (o *ResourceEventsPage) HasPageable() bool`

HasPageable returns a boolean if a field has been set.

### GetMetadata

`func (o *ResourceEventsPage) GetMetadata() ResourceEventsPageMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ResourceEventsPage) GetMetadataOk() (*ResourceEventsPageMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ResourceEventsPage) SetMetadata(v ResourceEventsPageMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ResourceEventsPage) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetItems

`func (o *ResourceEventsPage) GetItems() []ResourceEvent`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ResourceEventsPage) GetItemsOk() (*[]ResourceEvent, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ResourceEventsPage) SetItems(v []ResourceEvent)`

SetItems sets Items field to given value.

### HasItems

`func (o *ResourceEventsPage) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


