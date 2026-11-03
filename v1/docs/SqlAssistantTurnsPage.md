# SqlAssistantTurnsPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pageable** | Pointer to [**Pageable**](Pageable.md) |  | [optional] 
**Metadata** | Pointer to [**SqlAssistantTurnsPageMetadata**](SqlAssistantTurnsPageMetadata.md) |  | [optional] 
**Items** | Pointer to [**[]SqlAssistantTurn**](SqlAssistantTurn.md) |  | [optional] [default to []]

## Methods

### NewSqlAssistantTurnsPage

`func NewSqlAssistantTurnsPage() *SqlAssistantTurnsPage`

NewSqlAssistantTurnsPage instantiates a new SqlAssistantTurnsPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlAssistantTurnsPageWithDefaults

`func NewSqlAssistantTurnsPageWithDefaults() *SqlAssistantTurnsPage`

NewSqlAssistantTurnsPageWithDefaults instantiates a new SqlAssistantTurnsPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageable

`func (o *SqlAssistantTurnsPage) GetPageable() Pageable`

GetPageable returns the Pageable field if non-nil, zero value otherwise.

### GetPageableOk

`func (o *SqlAssistantTurnsPage) GetPageableOk() (*Pageable, bool)`

GetPageableOk returns a tuple with the Pageable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageable

`func (o *SqlAssistantTurnsPage) SetPageable(v Pageable)`

SetPageable sets Pageable field to given value.

### HasPageable

`func (o *SqlAssistantTurnsPage) HasPageable() bool`

HasPageable returns a boolean if a field has been set.

### GetMetadata

`func (o *SqlAssistantTurnsPage) GetMetadata() SqlAssistantTurnsPageMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SqlAssistantTurnsPage) GetMetadataOk() (*SqlAssistantTurnsPageMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SqlAssistantTurnsPage) SetMetadata(v SqlAssistantTurnsPageMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SqlAssistantTurnsPage) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetItems

`func (o *SqlAssistantTurnsPage) GetItems() []SqlAssistantTurn`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SqlAssistantTurnsPage) GetItemsOk() (*[]SqlAssistantTurn, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SqlAssistantTurnsPage) SetItems(v []SqlAssistantTurn)`

SetItems sets Items field to given value.

### HasItems

`func (o *SqlAssistantTurnsPage) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


