# SavepointsPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pageable** | Pointer to [**Pageable**](Pageable.md) |  | [optional] 
**Metadata** | Pointer to [**SavepointPageMetadata**](SavepointPageMetadata.md) |  | [optional] 
**Items** | Pointer to [**[]Savepoint**](Savepoint.md) |  | [optional] [default to []]

## Methods

### NewSavepointsPage

`func NewSavepointsPage() *SavepointsPage`

NewSavepointsPage instantiates a new SavepointsPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSavepointsPageWithDefaults

`func NewSavepointsPageWithDefaults() *SavepointsPage`

NewSavepointsPageWithDefaults instantiates a new SavepointsPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageable

`func (o *SavepointsPage) GetPageable() Pageable`

GetPageable returns the Pageable field if non-nil, zero value otherwise.

### GetPageableOk

`func (o *SavepointsPage) GetPageableOk() (*Pageable, bool)`

GetPageableOk returns a tuple with the Pageable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageable

`func (o *SavepointsPage) SetPageable(v Pageable)`

SetPageable sets Pageable field to given value.

### HasPageable

`func (o *SavepointsPage) HasPageable() bool`

HasPageable returns a boolean if a field has been set.

### GetMetadata

`func (o *SavepointsPage) GetMetadata() SavepointPageMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SavepointsPage) GetMetadataOk() (*SavepointPageMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SavepointsPage) SetMetadata(v SavepointPageMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SavepointsPage) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetItems

`func (o *SavepointsPage) GetItems() []Savepoint`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SavepointsPage) GetItemsOk() (*[]Savepoint, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SavepointsPage) SetItems(v []Savepoint)`

SetItems sets Items field to given value.

### HasItems

`func (o *SavepointsPage) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


