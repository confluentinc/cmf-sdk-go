# SavepointsPageAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**SavepointPageMetadata**](SavepointPageMetadata.md) |  | [optional] 
**Items** | Pointer to [**[]Savepoint**](Savepoint.md) |  | [optional] [default to []]

## Methods

### NewSavepointsPageAllOf

`func NewSavepointsPageAllOf() *SavepointsPageAllOf`

NewSavepointsPageAllOf instantiates a new SavepointsPageAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSavepointsPageAllOfWithDefaults

`func NewSavepointsPageAllOfWithDefaults() *SavepointsPageAllOf`

NewSavepointsPageAllOfWithDefaults instantiates a new SavepointsPageAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *SavepointsPageAllOf) GetMetadata() SavepointPageMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SavepointsPageAllOf) GetMetadataOk() (*SavepointPageMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SavepointsPageAllOf) SetMetadata(v SavepointPageMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SavepointsPageAllOf) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetItems

`func (o *SavepointsPageAllOf) GetItems() []Savepoint`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SavepointsPageAllOf) GetItemsOk() (*[]Savepoint, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SavepointsPageAllOf) SetItems(v []Savepoint)`

SetItems sets Items field to given value.

### HasItems

`func (o *SavepointsPageAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


