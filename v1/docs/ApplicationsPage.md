# ApplicationsPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pageable** | Pointer to [**Pageable**](Pageable.md) |  | [optional] 
**Metadata** | Pointer to [**ApplicationPageMetadata**](ApplicationPageMetadata.md) |  | [optional] 
**Items** | Pointer to [**[]FlinkApplication**](FlinkApplication.md) |  | [optional] [default to []]

## Methods

### NewApplicationsPage

`func NewApplicationsPage() *ApplicationsPage`

NewApplicationsPage instantiates a new ApplicationsPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationsPageWithDefaults

`func NewApplicationsPageWithDefaults() *ApplicationsPage`

NewApplicationsPageWithDefaults instantiates a new ApplicationsPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageable

`func (o *ApplicationsPage) GetPageable() Pageable`

GetPageable returns the Pageable field if non-nil, zero value otherwise.

### GetPageableOk

`func (o *ApplicationsPage) GetPageableOk() (*Pageable, bool)`

GetPageableOk returns a tuple with the Pageable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageable

`func (o *ApplicationsPage) SetPageable(v Pageable)`

SetPageable sets Pageable field to given value.

### HasPageable

`func (o *ApplicationsPage) HasPageable() bool`

HasPageable returns a boolean if a field has been set.

### GetMetadata

`func (o *ApplicationsPage) GetMetadata() ApplicationPageMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ApplicationsPage) GetMetadataOk() (*ApplicationPageMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ApplicationsPage) SetMetadata(v ApplicationPageMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ApplicationsPage) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetItems

`func (o *ApplicationsPage) GetItems() []FlinkApplication`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ApplicationsPage) GetItemsOk() (*[]FlinkApplication, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ApplicationsPage) SetItems(v []FlinkApplication)`

SetItems sets Items field to given value.

### HasItems

`func (o *ApplicationsPage) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


