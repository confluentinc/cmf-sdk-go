# StatementsPageAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**StatementPageMetadata**](StatementPageMetadata.md) |  | [optional] 
**Items** | Pointer to [**[]Statement**](Statement.md) |  | [optional] [default to []]

## Methods

### NewStatementsPageAllOf

`func NewStatementsPageAllOf() *StatementsPageAllOf`

NewStatementsPageAllOf instantiates a new StatementsPageAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatementsPageAllOfWithDefaults

`func NewStatementsPageAllOfWithDefaults() *StatementsPageAllOf`

NewStatementsPageAllOfWithDefaults instantiates a new StatementsPageAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *StatementsPageAllOf) GetMetadata() StatementPageMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *StatementsPageAllOf) GetMetadataOk() (*StatementPageMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *StatementsPageAllOf) SetMetadata(v StatementPageMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *StatementsPageAllOf) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetItems

`func (o *StatementsPageAllOf) GetItems() []Statement`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *StatementsPageAllOf) GetItemsOk() (*[]Statement, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *StatementsPageAllOf) SetItems(v []Statement)`

SetItems sets Items field to given value.

### HasItems

`func (o *StatementsPageAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


