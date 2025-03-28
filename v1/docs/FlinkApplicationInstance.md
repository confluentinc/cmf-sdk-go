# FlinkApplicationInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | API version for spec | 
**Kind** | **string** | Kind of resource - set to resource type | 
**Metadata** | Pointer to [**ApplicationInstanceMetadata**](ApplicationInstanceMetadata.md) |  | [optional] 
**Status** | Pointer to [**ApplicationInstanceStatus**](ApplicationInstanceStatus.md) |  | [optional] 

## Methods

### NewFlinkApplicationInstance

`func NewFlinkApplicationInstance(apiVersion string, kind string, ) *FlinkApplicationInstance`

NewFlinkApplicationInstance instantiates a new FlinkApplicationInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlinkApplicationInstanceWithDefaults

`func NewFlinkApplicationInstanceWithDefaults() *FlinkApplicationInstance`

NewFlinkApplicationInstanceWithDefaults instantiates a new FlinkApplicationInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *FlinkApplicationInstance) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *FlinkApplicationInstance) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *FlinkApplicationInstance) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *FlinkApplicationInstance) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *FlinkApplicationInstance) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *FlinkApplicationInstance) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *FlinkApplicationInstance) GetMetadata() ApplicationInstanceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FlinkApplicationInstance) GetMetadataOk() (*ApplicationInstanceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FlinkApplicationInstance) SetMetadata(v ApplicationInstanceMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FlinkApplicationInstance) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetStatus

`func (o *FlinkApplicationInstance) GetStatus() ApplicationInstanceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FlinkApplicationInstance) GetStatusOk() (*ApplicationInstanceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FlinkApplicationInstance) SetStatus(v ApplicationInstanceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FlinkApplicationInstance) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


