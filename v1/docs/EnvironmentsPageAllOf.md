# EnvironmentsPageAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**EnvironmentsPageMetadata**](EnvironmentsPageMetadata.md) |  | [optional] 
**Items** | Pointer to [**[]Environment**](Environment.md) |  | [optional] [default to []]

## Methods

### NewEnvironmentsPageAllOf

`func NewEnvironmentsPageAllOf() *EnvironmentsPageAllOf`

NewEnvironmentsPageAllOf instantiates a new EnvironmentsPageAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentsPageAllOfWithDefaults

`func NewEnvironmentsPageAllOfWithDefaults() *EnvironmentsPageAllOf`

NewEnvironmentsPageAllOfWithDefaults instantiates a new EnvironmentsPageAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *EnvironmentsPageAllOf) GetMetadata() EnvironmentsPageMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EnvironmentsPageAllOf) GetMetadataOk() (*EnvironmentsPageMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EnvironmentsPageAllOf) SetMetadata(v EnvironmentsPageMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EnvironmentsPageAllOf) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetItems

`func (o *EnvironmentsPageAllOf) GetItems() []Environment`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *EnvironmentsPageAllOf) GetItemsOk() (*[]Environment, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *EnvironmentsPageAllOf) SetItems(v []Environment)`

SetItems sets Items field to given value.

### HasItems

`func (o *EnvironmentsPageAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


