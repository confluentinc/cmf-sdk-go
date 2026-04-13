# KubernetesClustersPageAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**KubernetesClustersPageMetadata**](KubernetesClustersPageMetadata.md) |  | [optional] 
**Items** | Pointer to [**[]KubernetesCluster**](KubernetesCluster.md) |  | [optional] [default to []]

## Methods

### NewKubernetesClustersPageAllOf

`func NewKubernetesClustersPageAllOf() *KubernetesClustersPageAllOf`

NewKubernetesClustersPageAllOf instantiates a new KubernetesClustersPageAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesClustersPageAllOfWithDefaults

`func NewKubernetesClustersPageAllOfWithDefaults() *KubernetesClustersPageAllOf`

NewKubernetesClustersPageAllOfWithDefaults instantiates a new KubernetesClustersPageAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *KubernetesClustersPageAllOf) GetMetadata() KubernetesClustersPageMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *KubernetesClustersPageAllOf) GetMetadataOk() (*KubernetesClustersPageMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *KubernetesClustersPageAllOf) SetMetadata(v KubernetesClustersPageMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *KubernetesClustersPageAllOf) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetItems

`func (o *KubernetesClustersPageAllOf) GetItems() []KubernetesCluster`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *KubernetesClustersPageAllOf) GetItemsOk() (*[]KubernetesCluster, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *KubernetesClustersPageAllOf) SetItems(v []KubernetesCluster)`

SetItems sets Items field to given value.

### HasItems

`func (o *KubernetesClustersPageAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


