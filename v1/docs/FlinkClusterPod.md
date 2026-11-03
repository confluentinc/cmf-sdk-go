# FlinkClusterPod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** |  | 
**Kind** | **string** |  | 
**Metadata** | [**FlinkClusterPodMetadata**](FlinkClusterPodMetadata.md) |  | 
**Status** | [**FlinkClusterPodStatus**](FlinkClusterPodStatus.md) |  | 

## Methods

### NewFlinkClusterPod

`func NewFlinkClusterPod(apiVersion string, kind string, metadata FlinkClusterPodMetadata, status FlinkClusterPodStatus, ) *FlinkClusterPod`

NewFlinkClusterPod instantiates a new FlinkClusterPod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlinkClusterPodWithDefaults

`func NewFlinkClusterPodWithDefaults() *FlinkClusterPod`

NewFlinkClusterPodWithDefaults instantiates a new FlinkClusterPod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *FlinkClusterPod) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *FlinkClusterPod) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *FlinkClusterPod) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *FlinkClusterPod) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *FlinkClusterPod) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *FlinkClusterPod) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *FlinkClusterPod) GetMetadata() FlinkClusterPodMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FlinkClusterPod) GetMetadataOk() (*FlinkClusterPodMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FlinkClusterPod) SetMetadata(v FlinkClusterPodMetadata)`

SetMetadata sets Metadata field to given value.


### GetStatus

`func (o *FlinkClusterPod) GetStatus() FlinkClusterPodStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FlinkClusterPod) GetStatusOk() (*FlinkClusterPodStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FlinkClusterPod) SetStatus(v FlinkClusterPodStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


