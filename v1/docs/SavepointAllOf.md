# SavepointAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**SavepointMetadata**](SavepointMetadata.md) |  | 
**Spec** | [**SavepointSpec**](SavepointSpec.md) |  | 
**Status** | Pointer to [**SavepointStatus**](SavepointStatus.md) |  | [optional] 

## Methods

### NewSavepointAllOf

`func NewSavepointAllOf(metadata SavepointMetadata, spec SavepointSpec, ) *SavepointAllOf`

NewSavepointAllOf instantiates a new SavepointAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSavepointAllOfWithDefaults

`func NewSavepointAllOfWithDefaults() *SavepointAllOf`

NewSavepointAllOfWithDefaults instantiates a new SavepointAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *SavepointAllOf) GetMetadata() SavepointMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SavepointAllOf) GetMetadataOk() (*SavepointMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SavepointAllOf) SetMetadata(v SavepointMetadata)`

SetMetadata sets Metadata field to given value.


### GetSpec

`func (o *SavepointAllOf) GetSpec() SavepointSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *SavepointAllOf) GetSpecOk() (*SavepointSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *SavepointAllOf) SetSpec(v SavepointSpec)`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *SavepointAllOf) GetStatus() SavepointStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SavepointAllOf) GetStatusOk() (*SavepointStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SavepointAllOf) SetStatus(v SavepointStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SavepointAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


