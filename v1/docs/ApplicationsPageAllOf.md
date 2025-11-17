# ApplicationsPageAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ApplicationPageMetadata**](ApplicationPageMetadata.md) |  | [optional] 
**Items** | Pointer to [**[]FlinkApplication**](FlinkApplication.md) |  | [optional] [default to []]

## Methods

### NewApplicationsPageAllOf

`func NewApplicationsPageAllOf() *ApplicationsPageAllOf`

NewApplicationsPageAllOf instantiates a new ApplicationsPageAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationsPageAllOfWithDefaults

`func NewApplicationsPageAllOfWithDefaults() *ApplicationsPageAllOf`

NewApplicationsPageAllOfWithDefaults instantiates a new ApplicationsPageAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *ApplicationsPageAllOf) GetMetadata() ApplicationPageMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ApplicationsPageAllOf) GetMetadataOk() (*ApplicationPageMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ApplicationsPageAllOf) SetMetadata(v ApplicationPageMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ApplicationsPageAllOf) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetItems

`func (o *ApplicationsPageAllOf) GetItems() []FlinkApplication`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ApplicationsPageAllOf) GetItemsOk() (*[]FlinkApplication, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ApplicationsPageAllOf) SetItems(v []FlinkApplication)`

SetItems sets Items field to given value.

### HasItems

`func (o *ApplicationsPageAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


