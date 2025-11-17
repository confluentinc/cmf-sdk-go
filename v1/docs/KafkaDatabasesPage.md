# KafkaDatabasesPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pageable** | Pointer to [**Pageable**](Pageable.md) |  | [optional] 
**Metadata** | Pointer to [**DatabasePageMetadata**](DatabasePageMetadata.md) |  | [optional] 
**Items** | Pointer to [**[]KafkaDatabase**](KafkaDatabase.md) |  | [optional] [default to []]

## Methods

### NewKafkaDatabasesPage

`func NewKafkaDatabasesPage() *KafkaDatabasesPage`

NewKafkaDatabasesPage instantiates a new KafkaDatabasesPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaDatabasesPageWithDefaults

`func NewKafkaDatabasesPageWithDefaults() *KafkaDatabasesPage`

NewKafkaDatabasesPageWithDefaults instantiates a new KafkaDatabasesPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageable

`func (o *KafkaDatabasesPage) GetPageable() Pageable`

GetPageable returns the Pageable field if non-nil, zero value otherwise.

### GetPageableOk

`func (o *KafkaDatabasesPage) GetPageableOk() (*Pageable, bool)`

GetPageableOk returns a tuple with the Pageable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageable

`func (o *KafkaDatabasesPage) SetPageable(v Pageable)`

SetPageable sets Pageable field to given value.

### HasPageable

`func (o *KafkaDatabasesPage) HasPageable() bool`

HasPageable returns a boolean if a field has been set.

### GetMetadata

`func (o *KafkaDatabasesPage) GetMetadata() DatabasePageMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *KafkaDatabasesPage) GetMetadataOk() (*DatabasePageMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *KafkaDatabasesPage) SetMetadata(v DatabasePageMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *KafkaDatabasesPage) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetItems

`func (o *KafkaDatabasesPage) GetItems() []KafkaDatabase`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *KafkaDatabasesPage) GetItemsOk() (*[]KafkaDatabase, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *KafkaDatabasesPage) SetItems(v []KafkaDatabase)`

SetItems sets Items field to given value.

### HasItems

`func (o *KafkaDatabasesPage) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


