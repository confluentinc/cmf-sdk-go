# ApplicationInstancesPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pageable** | Pointer to [**Pageable**](Pageable.md) |  | [optional] 
**Metadata** | Pointer to [**ApplicationInstancesPageMetadata**](ApplicationInstancesPageMetadata.md) |  | [optional] 
**Items** | Pointer to [**[]FlinkApplicationInstance**](FlinkApplicationInstance.md) |  | [optional] [default to []]

## Methods

### NewApplicationInstancesPage

`func NewApplicationInstancesPage() *ApplicationInstancesPage`

NewApplicationInstancesPage instantiates a new ApplicationInstancesPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationInstancesPageWithDefaults

`func NewApplicationInstancesPageWithDefaults() *ApplicationInstancesPage`

NewApplicationInstancesPageWithDefaults instantiates a new ApplicationInstancesPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageable

`func (o *ApplicationInstancesPage) GetPageable() Pageable`

GetPageable returns the Pageable field if non-nil, zero value otherwise.

### GetPageableOk

`func (o *ApplicationInstancesPage) GetPageableOk() (*Pageable, bool)`

GetPageableOk returns a tuple with the Pageable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageable

`func (o *ApplicationInstancesPage) SetPageable(v Pageable)`

SetPageable sets Pageable field to given value.

### HasPageable

`func (o *ApplicationInstancesPage) HasPageable() bool`

HasPageable returns a boolean if a field has been set.

### GetMetadata

`func (o *ApplicationInstancesPage) GetMetadata() ApplicationInstancesPageMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ApplicationInstancesPage) GetMetadataOk() (*ApplicationInstancesPageMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ApplicationInstancesPage) SetMetadata(v ApplicationInstancesPageMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ApplicationInstancesPage) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetItems

`func (o *ApplicationInstancesPage) GetItems() []FlinkApplicationInstance`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ApplicationInstancesPage) GetItemsOk() (*[]FlinkApplicationInstance, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ApplicationInstancesPage) SetItems(v []FlinkApplicationInstance)`

SetItems sets Items field to given value.

### HasItems

`func (o *ApplicationInstancesPage) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


