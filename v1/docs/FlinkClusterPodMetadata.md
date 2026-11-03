# FlinkClusterPodMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the pod. Use it with the &#x60;podName&#x60; query parameter of the logs endpoint. | 
**CreationTimestamp** | Pointer to **time.Time** | When Kubernetes started the pod. Absent while the pod is still being scheduled. | [optional] 

## Methods

### NewFlinkClusterPodMetadata

`func NewFlinkClusterPodMetadata(name string, ) *FlinkClusterPodMetadata`

NewFlinkClusterPodMetadata instantiates a new FlinkClusterPodMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlinkClusterPodMetadataWithDefaults

`func NewFlinkClusterPodMetadataWithDefaults() *FlinkClusterPodMetadata`

NewFlinkClusterPodMetadataWithDefaults instantiates a new FlinkClusterPodMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FlinkClusterPodMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlinkClusterPodMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlinkClusterPodMetadata) SetName(v string)`

SetName sets Name field to given value.


### GetCreationTimestamp

`func (o *FlinkClusterPodMetadata) GetCreationTimestamp() time.Time`

GetCreationTimestamp returns the CreationTimestamp field if non-nil, zero value otherwise.

### GetCreationTimestampOk

`func (o *FlinkClusterPodMetadata) GetCreationTimestampOk() (*time.Time, bool)`

GetCreationTimestampOk returns a tuple with the CreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimestamp

`func (o *FlinkClusterPodMetadata) SetCreationTimestamp(v time.Time)`

SetCreationTimestamp sets CreationTimestamp field to given value.

### HasCreationTimestamp

`func (o *FlinkClusterPodMetadata) HasCreationTimestamp() bool`

HasCreationTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


