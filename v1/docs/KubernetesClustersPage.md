# KubernetesClustersPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pageable** | Pointer to [**Pageable**](Pageable.md) |  | [optional] 
**Metadata** | Pointer to [**KubernetesClustersPageMetadata**](KubernetesClustersPageMetadata.md) |  | [optional] 
**Items** | Pointer to [**[]KubernetesCluster**](KubernetesCluster.md) |  | [optional] [default to []]

## Methods

### NewKubernetesClustersPage

`func NewKubernetesClustersPage() *KubernetesClustersPage`

NewKubernetesClustersPage instantiates a new KubernetesClustersPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesClustersPageWithDefaults

`func NewKubernetesClustersPageWithDefaults() *KubernetesClustersPage`

NewKubernetesClustersPageWithDefaults instantiates a new KubernetesClustersPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageable

`func (o *KubernetesClustersPage) GetPageable() Pageable`

GetPageable returns the Pageable field if non-nil, zero value otherwise.

### GetPageableOk

`func (o *KubernetesClustersPage) GetPageableOk() (*Pageable, bool)`

GetPageableOk returns a tuple with the Pageable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageable

`func (o *KubernetesClustersPage) SetPageable(v Pageable)`

SetPageable sets Pageable field to given value.

### HasPageable

`func (o *KubernetesClustersPage) HasPageable() bool`

HasPageable returns a boolean if a field has been set.

### GetMetadata

`func (o *KubernetesClustersPage) GetMetadata() KubernetesClustersPageMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *KubernetesClustersPage) GetMetadataOk() (*KubernetesClustersPageMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *KubernetesClustersPage) SetMetadata(v KubernetesClustersPageMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *KubernetesClustersPage) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetItems

`func (o *KubernetesClustersPage) GetItems() []KubernetesCluster`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *KubernetesClustersPage) GetItemsOk() (*[]KubernetesCluster, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *KubernetesClustersPage) SetItems(v []KubernetesCluster)`

SetItems sets Items field to given value.

### HasItems

`func (o *KubernetesClustersPage) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


