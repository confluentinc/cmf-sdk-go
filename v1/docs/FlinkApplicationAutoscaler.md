# FlinkApplicationAutoscaler

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** |  | 
**Kind** | **string** |  | 
**Metadata** | [**AutoscalerMetadata**](AutoscalerMetadata.md) |  | 
**Status** | [**AutoscalerStatus**](AutoscalerStatus.md) |  | 

## Methods

### NewFlinkApplicationAutoscaler

`func NewFlinkApplicationAutoscaler(apiVersion string, kind string, metadata AutoscalerMetadata, status AutoscalerStatus, ) *FlinkApplicationAutoscaler`

NewFlinkApplicationAutoscaler instantiates a new FlinkApplicationAutoscaler object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlinkApplicationAutoscalerWithDefaults

`func NewFlinkApplicationAutoscalerWithDefaults() *FlinkApplicationAutoscaler`

NewFlinkApplicationAutoscalerWithDefaults instantiates a new FlinkApplicationAutoscaler object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *FlinkApplicationAutoscaler) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *FlinkApplicationAutoscaler) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *FlinkApplicationAutoscaler) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *FlinkApplicationAutoscaler) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *FlinkApplicationAutoscaler) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *FlinkApplicationAutoscaler) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *FlinkApplicationAutoscaler) GetMetadata() AutoscalerMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FlinkApplicationAutoscaler) GetMetadataOk() (*AutoscalerMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FlinkApplicationAutoscaler) SetMetadata(v AutoscalerMetadata)`

SetMetadata sets Metadata field to given value.


### GetStatus

`func (o *FlinkApplicationAutoscaler) GetStatus() AutoscalerStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FlinkApplicationAutoscaler) GetStatusOk() (*AutoscalerStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FlinkApplicationAutoscaler) SetStatus(v AutoscalerStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


