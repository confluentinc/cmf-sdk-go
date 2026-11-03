# SqlAssistantConversationsPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pageable** | Pointer to [**Pageable**](Pageable.md) |  | [optional] 
**Metadata** | Pointer to [**SqlAssistantConversationsPageMetadata**](SqlAssistantConversationsPageMetadata.md) |  | [optional] 
**Items** | Pointer to [**[]SqlAssistantConversation**](SqlAssistantConversation.md) |  | [optional] [default to []]

## Methods

### NewSqlAssistantConversationsPage

`func NewSqlAssistantConversationsPage() *SqlAssistantConversationsPage`

NewSqlAssistantConversationsPage instantiates a new SqlAssistantConversationsPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlAssistantConversationsPageWithDefaults

`func NewSqlAssistantConversationsPageWithDefaults() *SqlAssistantConversationsPage`

NewSqlAssistantConversationsPageWithDefaults instantiates a new SqlAssistantConversationsPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageable

`func (o *SqlAssistantConversationsPage) GetPageable() Pageable`

GetPageable returns the Pageable field if non-nil, zero value otherwise.

### GetPageableOk

`func (o *SqlAssistantConversationsPage) GetPageableOk() (*Pageable, bool)`

GetPageableOk returns a tuple with the Pageable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageable

`func (o *SqlAssistantConversationsPage) SetPageable(v Pageable)`

SetPageable sets Pageable field to given value.

### HasPageable

`func (o *SqlAssistantConversationsPage) HasPageable() bool`

HasPageable returns a boolean if a field has been set.

### GetMetadata

`func (o *SqlAssistantConversationsPage) GetMetadata() SqlAssistantConversationsPageMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SqlAssistantConversationsPage) GetMetadataOk() (*SqlAssistantConversationsPageMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SqlAssistantConversationsPage) SetMetadata(v SqlAssistantConversationsPageMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SqlAssistantConversationsPage) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetItems

`func (o *SqlAssistantConversationsPage) GetItems() []SqlAssistantConversation`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SqlAssistantConversationsPage) GetItemsOk() (*[]SqlAssistantConversation, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SqlAssistantConversationsPage) SetItems(v []SqlAssistantConversation)`

SetItems sets Items field to given value.

### HasItems

`func (o *SqlAssistantConversationsPage) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


