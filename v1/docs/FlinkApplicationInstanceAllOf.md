# FlinkApplicationInstanceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ApplicationInstanceMetadata**](ApplicationInstanceMetadata.md) |  | [optional] 
**Status** | Pointer to [**ApplicationInstanceStatus**](ApplicationInstanceStatus.md) |  | [optional] 

## Methods

### NewFlinkApplicationInstanceAllOf

`func NewFlinkApplicationInstanceAllOf() *FlinkApplicationInstanceAllOf`

NewFlinkApplicationInstanceAllOf instantiates a new FlinkApplicationInstanceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlinkApplicationInstanceAllOfWithDefaults

`func NewFlinkApplicationInstanceAllOfWithDefaults() *FlinkApplicationInstanceAllOf`

NewFlinkApplicationInstanceAllOfWithDefaults instantiates a new FlinkApplicationInstanceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *FlinkApplicationInstanceAllOf) GetMetadata() ApplicationInstanceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FlinkApplicationInstanceAllOf) GetMetadataOk() (*ApplicationInstanceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FlinkApplicationInstanceAllOf) SetMetadata(v ApplicationInstanceMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FlinkApplicationInstanceAllOf) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetStatus

`func (o *FlinkApplicationInstanceAllOf) GetStatus() ApplicationInstanceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FlinkApplicationInstanceAllOf) GetStatusOk() (*ApplicationInstanceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FlinkApplicationInstanceAllOf) SetStatus(v ApplicationInstanceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FlinkApplicationInstanceAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


