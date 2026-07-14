# KubernetesClusterMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the Kubernetes cluster. | 
**CreationTimestamp** | Pointer to **string** | Timestamp when the cluster was registered. | [optional] [readonly] 
**UpdateTimestamp** | Pointer to **string** | Timestamp when the cluster was last updated. | [optional] [readonly] 
**Uid** | Pointer to **string** | Unique identifier of the Kubernetes cluster. | [optional] [readonly] 
**Labels** | Pointer to **map[string]string** | Labels of the Kubernetes cluster. Omit the field (or send null) to leave existing labels unchanged; send an empty object to clear them.  | [optional] 
**Annotations** | Pointer to **map[string]string** | Annotations of the Kubernetes cluster. Omit the field (or send null) to leave existing annotations unchanged; send an empty object to clear them.  | [optional] 

## Methods

### NewKubernetesClusterMetadata

`func NewKubernetesClusterMetadata(name string, ) *KubernetesClusterMetadata`

NewKubernetesClusterMetadata instantiates a new KubernetesClusterMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesClusterMetadataWithDefaults

`func NewKubernetesClusterMetadataWithDefaults() *KubernetesClusterMetadata`

NewKubernetesClusterMetadataWithDefaults instantiates a new KubernetesClusterMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KubernetesClusterMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesClusterMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesClusterMetadata) SetName(v string)`

SetName sets Name field to given value.


### GetCreationTimestamp

`func (o *KubernetesClusterMetadata) GetCreationTimestamp() string`

GetCreationTimestamp returns the CreationTimestamp field if non-nil, zero value otherwise.

### GetCreationTimestampOk

`func (o *KubernetesClusterMetadata) GetCreationTimestampOk() (*string, bool)`

GetCreationTimestampOk returns a tuple with the CreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimestamp

`func (o *KubernetesClusterMetadata) SetCreationTimestamp(v string)`

SetCreationTimestamp sets CreationTimestamp field to given value.

### HasCreationTimestamp

`func (o *KubernetesClusterMetadata) HasCreationTimestamp() bool`

HasCreationTimestamp returns a boolean if a field has been set.

### GetUpdateTimestamp

`func (o *KubernetesClusterMetadata) GetUpdateTimestamp() string`

GetUpdateTimestamp returns the UpdateTimestamp field if non-nil, zero value otherwise.

### GetUpdateTimestampOk

`func (o *KubernetesClusterMetadata) GetUpdateTimestampOk() (*string, bool)`

GetUpdateTimestampOk returns a tuple with the UpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTimestamp

`func (o *KubernetesClusterMetadata) SetUpdateTimestamp(v string)`

SetUpdateTimestamp sets UpdateTimestamp field to given value.

### HasUpdateTimestamp

`func (o *KubernetesClusterMetadata) HasUpdateTimestamp() bool`

HasUpdateTimestamp returns a boolean if a field has been set.

### GetUid

`func (o *KubernetesClusterMetadata) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *KubernetesClusterMetadata) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *KubernetesClusterMetadata) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *KubernetesClusterMetadata) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetLabels

`func (o *KubernetesClusterMetadata) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *KubernetesClusterMetadata) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *KubernetesClusterMetadata) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *KubernetesClusterMetadata) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *KubernetesClusterMetadata) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *KubernetesClusterMetadata) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetAnnotations

`func (o *KubernetesClusterMetadata) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *KubernetesClusterMetadata) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *KubernetesClusterMetadata) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *KubernetesClusterMetadata) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### SetAnnotationsNil

`func (o *KubernetesClusterMetadata) SetAnnotationsNil(b bool)`

 SetAnnotationsNil sets the value for Annotations to be an explicit nil

### UnsetAnnotations
`func (o *KubernetesClusterMetadata) UnsetAnnotations()`

UnsetAnnotations ensures that no value is present for Annotations, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


