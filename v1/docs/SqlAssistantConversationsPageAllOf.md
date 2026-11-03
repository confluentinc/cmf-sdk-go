# SqlAssistantConversationsPageAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**SqlAssistantConversationsPageMetadata**](SqlAssistantConversationsPageMetadata.md) |  | [optional] 
**Items** | Pointer to [**[]SqlAssistantConversation**](SqlAssistantConversation.md) |  | [optional] [default to []]

## Methods

### NewSqlAssistantConversationsPageAllOf

`func NewSqlAssistantConversationsPageAllOf() *SqlAssistantConversationsPageAllOf`

NewSqlAssistantConversationsPageAllOf instantiates a new SqlAssistantConversationsPageAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlAssistantConversationsPageAllOfWithDefaults

`func NewSqlAssistantConversationsPageAllOfWithDefaults() *SqlAssistantConversationsPageAllOf`

NewSqlAssistantConversationsPageAllOfWithDefaults instantiates a new SqlAssistantConversationsPageAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *SqlAssistantConversationsPageAllOf) GetMetadata() SqlAssistantConversationsPageMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SqlAssistantConversationsPageAllOf) GetMetadataOk() (*SqlAssistantConversationsPageMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SqlAssistantConversationsPageAllOf) SetMetadata(v SqlAssistantConversationsPageMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SqlAssistantConversationsPageAllOf) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetItems

`func (o *SqlAssistantConversationsPageAllOf) GetItems() []SqlAssistantConversation`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SqlAssistantConversationsPageAllOf) GetItemsOk() (*[]SqlAssistantConversation, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SqlAssistantConversationsPageAllOf) SetItems(v []SqlAssistantConversation)`

SetItems sets Items field to given value.

### HasItems

`func (o *SqlAssistantConversationsPageAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


