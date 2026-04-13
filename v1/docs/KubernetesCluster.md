# KubernetesCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | API version for spec | 
**Kind** | **string** | Kind of resource - set to resource type | 
**Metadata** | [**KubernetesClusterMetadata**](KubernetesClusterMetadata.md) |  | 
**Spec** | [**KubernetesClusterSpec**](KubernetesClusterSpec.md) |  | 
**Status** | Pointer to [**KubernetesClusterStatus**](KubernetesClusterStatus.md) |  | [optional] 

## Methods

### NewKubernetesCluster

`func NewKubernetesCluster(apiVersion string, kind string, metadata KubernetesClusterMetadata, spec KubernetesClusterSpec, ) *KubernetesCluster`

NewKubernetesCluster instantiates a new KubernetesCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesClusterWithDefaults

`func NewKubernetesClusterWithDefaults() *KubernetesCluster`

NewKubernetesClusterWithDefaults instantiates a new KubernetesCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *KubernetesCluster) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *KubernetesCluster) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *KubernetesCluster) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *KubernetesCluster) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *KubernetesCluster) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *KubernetesCluster) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *KubernetesCluster) GetMetadata() KubernetesClusterMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *KubernetesCluster) GetMetadataOk() (*KubernetesClusterMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *KubernetesCluster) SetMetadata(v KubernetesClusterMetadata)`

SetMetadata sets Metadata field to given value.


### GetSpec

`func (o *KubernetesCluster) GetSpec() KubernetesClusterSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *KubernetesCluster) GetSpecOk() (*KubernetesClusterSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *KubernetesCluster) SetSpec(v KubernetesClusterSpec)`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *KubernetesCluster) GetStatus() KubernetesClusterStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KubernetesCluster) GetStatusOk() (*KubernetesClusterStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KubernetesCluster) SetStatus(v KubernetesClusterStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KubernetesCluster) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


