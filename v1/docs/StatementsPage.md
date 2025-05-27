# StatementsPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pageable** | Pointer to [**Pageable**](Pageable.md) |  | [optional] 
**Metadata** | Pointer to [**StatementPageMetadata**](StatementPageMetadata.md) |  | [optional] 
**Items** | Pointer to [**[]Statement**](Statement.md) |  | [optional] [default to []]

## Methods

### NewStatementsPage

`func NewStatementsPage() *StatementsPage`

NewStatementsPage instantiates a new StatementsPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatementsPageWithDefaults

`func NewStatementsPageWithDefaults() *StatementsPage`

NewStatementsPageWithDefaults instantiates a new StatementsPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageable

`func (o *StatementsPage) GetPageable() Pageable`

GetPageable returns the Pageable field if non-nil, zero value otherwise.

### GetPageableOk

`func (o *StatementsPage) GetPageableOk() (*Pageable, bool)`

GetPageableOk returns a tuple with the Pageable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageable

`func (o *StatementsPage) SetPageable(v Pageable)`

SetPageable sets Pageable field to given value.

### HasPageable

`func (o *StatementsPage) HasPageable() bool`

HasPageable returns a boolean if a field has been set.

### GetMetadata

`func (o *StatementsPage) GetMetadata() StatementPageMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *StatementsPage) GetMetadataOk() (*StatementPageMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *StatementsPage) SetMetadata(v StatementPageMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *StatementsPage) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetItems

`func (o *StatementsPage) GetItems() []Statement`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *StatementsPage) GetItemsOk() (*[]Statement, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *StatementsPage) SetItems(v []Statement)`

SetItems sets Items field to given value.

### HasItems

`func (o *StatementsPage) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


