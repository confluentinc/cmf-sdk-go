# ComputePoolAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**ComputePoolMetadata**](ComputePoolMetadata.md) |  | 
**Spec** | [**ComputePoolSpec**](ComputePoolSpec.md) |  | 
**Status** | Pointer to [**ComputePoolStatus**](ComputePoolStatus.md) |  | [optional] 

## Methods

### NewComputePoolAllOf

`func NewComputePoolAllOf(metadata ComputePoolMetadata, spec ComputePoolSpec, ) *ComputePoolAllOf`

NewComputePoolAllOf instantiates a new ComputePoolAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputePoolAllOfWithDefaults

`func NewComputePoolAllOfWithDefaults() *ComputePoolAllOf`

NewComputePoolAllOfWithDefaults instantiates a new ComputePoolAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *ComputePoolAllOf) GetMetadata() ComputePoolMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ComputePoolAllOf) GetMetadataOk() (*ComputePoolMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ComputePoolAllOf) SetMetadata(v ComputePoolMetadata)`

SetMetadata sets Metadata field to given value.


### GetSpec

`func (o *ComputePoolAllOf) GetSpec() ComputePoolSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ComputePoolAllOf) GetSpecOk() (*ComputePoolSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ComputePoolAllOf) SetSpec(v ComputePoolSpec)`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *ComputePoolAllOf) GetStatus() ComputePoolStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ComputePoolAllOf) GetStatusOk() (*ComputePoolStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ComputePoolAllOf) SetStatus(v ComputePoolStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ComputePoolAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


