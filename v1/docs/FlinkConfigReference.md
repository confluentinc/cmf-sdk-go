# FlinkConfigReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** |  | 
**Kind** | **string** |  | 
**Metadata** | [**FlinkConfigReferenceMetadata**](FlinkConfigReferenceMetadata.md) |  | 
**Status** | [**FlinkConfigReferenceStatus**](FlinkConfigReferenceStatus.md) |  | 

## Methods

### NewFlinkConfigReference

`func NewFlinkConfigReference(apiVersion string, kind string, metadata FlinkConfigReferenceMetadata, status FlinkConfigReferenceStatus, ) *FlinkConfigReference`

NewFlinkConfigReference instantiates a new FlinkConfigReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlinkConfigReferenceWithDefaults

`func NewFlinkConfigReferenceWithDefaults() *FlinkConfigReference`

NewFlinkConfigReferenceWithDefaults instantiates a new FlinkConfigReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *FlinkConfigReference) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *FlinkConfigReference) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *FlinkConfigReference) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *FlinkConfigReference) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *FlinkConfigReference) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *FlinkConfigReference) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *FlinkConfigReference) GetMetadata() FlinkConfigReferenceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FlinkConfigReference) GetMetadataOk() (*FlinkConfigReferenceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FlinkConfigReference) SetMetadata(v FlinkConfigReferenceMetadata)`

SetMetadata sets Metadata field to given value.


### GetStatus

`func (o *FlinkConfigReference) GetStatus() FlinkConfigReferenceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FlinkConfigReference) GetStatusOk() (*FlinkConfigReferenceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FlinkConfigReference) SetStatus(v FlinkConfigReferenceStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


