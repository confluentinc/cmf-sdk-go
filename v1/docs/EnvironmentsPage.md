# EnvironmentsPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pageable** | Pointer to [**Pageable**](Pageable.md) |  | [optional] 
**Metadata** | Pointer to [**EnvironmentsPageMetadata**](EnvironmentsPageMetadata.md) |  | [optional] 
**Items** | Pointer to [**[]Environment**](Environment.md) |  | [optional] [default to []]

## Methods

### NewEnvironmentsPage

`func NewEnvironmentsPage() *EnvironmentsPage`

NewEnvironmentsPage instantiates a new EnvironmentsPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentsPageWithDefaults

`func NewEnvironmentsPageWithDefaults() *EnvironmentsPage`

NewEnvironmentsPageWithDefaults instantiates a new EnvironmentsPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageable

`func (o *EnvironmentsPage) GetPageable() Pageable`

GetPageable returns the Pageable field if non-nil, zero value otherwise.

### GetPageableOk

`func (o *EnvironmentsPage) GetPageableOk() (*Pageable, bool)`

GetPageableOk returns a tuple with the Pageable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageable

`func (o *EnvironmentsPage) SetPageable(v Pageable)`

SetPageable sets Pageable field to given value.

### HasPageable

`func (o *EnvironmentsPage) HasPageable() bool`

HasPageable returns a boolean if a field has been set.

### GetMetadata

`func (o *EnvironmentsPage) GetMetadata() EnvironmentsPageMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EnvironmentsPage) GetMetadataOk() (*EnvironmentsPageMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EnvironmentsPage) SetMetadata(v EnvironmentsPageMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EnvironmentsPage) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetItems

`func (o *EnvironmentsPage) GetItems() []Environment`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *EnvironmentsPage) GetItemsOk() (*[]Environment, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *EnvironmentsPage) SetItems(v []Environment)`

SetItems sets Items field to given value.

### HasItems

`func (o *EnvironmentsPage) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


