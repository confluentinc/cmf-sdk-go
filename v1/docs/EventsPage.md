# EventsPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pageable** | Pointer to [**Pageable**](Pageable.md) |  | [optional] 
**Metadata** | Pointer to [**EventsPageMetadata**](EventsPageMetadata.md) |  | [optional] 
**Items** | Pointer to [**[]FlinkApplicationEvent**](FlinkApplicationEvent.md) |  | [optional] [default to []]

## Methods

### NewEventsPage

`func NewEventsPage() *EventsPage`

NewEventsPage instantiates a new EventsPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsPageWithDefaults

`func NewEventsPageWithDefaults() *EventsPage`

NewEventsPageWithDefaults instantiates a new EventsPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageable

`func (o *EventsPage) GetPageable() Pageable`

GetPageable returns the Pageable field if non-nil, zero value otherwise.

### GetPageableOk

`func (o *EventsPage) GetPageableOk() (*Pageable, bool)`

GetPageableOk returns a tuple with the Pageable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageable

`func (o *EventsPage) SetPageable(v Pageable)`

SetPageable sets Pageable field to given value.

### HasPageable

`func (o *EventsPage) HasPageable() bool`

HasPageable returns a boolean if a field has been set.

### GetMetadata

`func (o *EventsPage) GetMetadata() EventsPageMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EventsPage) GetMetadataOk() (*EventsPageMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EventsPage) SetMetadata(v EventsPageMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EventsPage) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetItems

`func (o *EventsPage) GetItems() []FlinkApplicationEvent`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *EventsPage) GetItemsOk() (*[]FlinkApplicationEvent, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *EventsPage) SetItems(v []FlinkApplicationEvent)`

SetItems sets Items field to given value.

### HasItems

`func (o *EventsPage) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


