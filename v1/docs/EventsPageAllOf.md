# EventsPageAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**EventsPageMetadata**](EventsPageMetadata.md) |  | [optional] 
**Items** | Pointer to [**[]FlinkApplicationEvent**](FlinkApplicationEvent.md) |  | [optional] [default to []]

## Methods

### NewEventsPageAllOf

`func NewEventsPageAllOf() *EventsPageAllOf`

NewEventsPageAllOf instantiates a new EventsPageAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsPageAllOfWithDefaults

`func NewEventsPageAllOfWithDefaults() *EventsPageAllOf`

NewEventsPageAllOfWithDefaults instantiates a new EventsPageAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *EventsPageAllOf) GetMetadata() EventsPageMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EventsPageAllOf) GetMetadataOk() (*EventsPageMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EventsPageAllOf) SetMetadata(v EventsPageMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EventsPageAllOf) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetItems

`func (o *EventsPageAllOf) GetItems() []FlinkApplicationEvent`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *EventsPageAllOf) GetItemsOk() (*[]FlinkApplicationEvent, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *EventsPageAllOf) SetItems(v []FlinkApplicationEvent)`

SetItems sets Items field to given value.

### HasItems

`func (o *EventsPageAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


