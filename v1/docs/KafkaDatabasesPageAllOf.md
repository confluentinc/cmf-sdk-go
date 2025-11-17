# KafkaDatabasesPageAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**DatabasePageMetadata**](DatabasePageMetadata.md) |  | [optional] 
**Items** | Pointer to [**[]KafkaDatabase**](KafkaDatabase.md) |  | [optional] [default to []]

## Methods

### NewKafkaDatabasesPageAllOf

`func NewKafkaDatabasesPageAllOf() *KafkaDatabasesPageAllOf`

NewKafkaDatabasesPageAllOf instantiates a new KafkaDatabasesPageAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaDatabasesPageAllOfWithDefaults

`func NewKafkaDatabasesPageAllOfWithDefaults() *KafkaDatabasesPageAllOf`

NewKafkaDatabasesPageAllOfWithDefaults instantiates a new KafkaDatabasesPageAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *KafkaDatabasesPageAllOf) GetMetadata() DatabasePageMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *KafkaDatabasesPageAllOf) GetMetadataOk() (*DatabasePageMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *KafkaDatabasesPageAllOf) SetMetadata(v DatabasePageMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *KafkaDatabasesPageAllOf) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetItems

`func (o *KafkaDatabasesPageAllOf) GetItems() []KafkaDatabase`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *KafkaDatabasesPageAllOf) GetItemsOk() (*[]KafkaDatabase, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *KafkaDatabasesPageAllOf) SetItems(v []KafkaDatabase)`

SetItems sets Items field to given value.

### HasItems

`func (o *KafkaDatabasesPageAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


