# Savepoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | API version for spec | 
**Kind** | **string** | Kind of resource - set to resource type | 
**Metadata** | [**SavepointMetadata**](SavepointMetadata.md) |  | 
**Spec** | [**SavepointSpec**](SavepointSpec.md) |  | 
**Status** | Pointer to [**SavepointStatus**](SavepointStatus.md) |  | [optional] 

## Methods

### NewSavepoint

`func NewSavepoint(apiVersion string, kind string, metadata SavepointMetadata, spec SavepointSpec, ) *Savepoint`

NewSavepoint instantiates a new Savepoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSavepointWithDefaults

`func NewSavepointWithDefaults() *Savepoint`

NewSavepointWithDefaults instantiates a new Savepoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *Savepoint) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *Savepoint) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *Savepoint) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *Savepoint) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Savepoint) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Savepoint) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *Savepoint) GetMetadata() SavepointMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Savepoint) GetMetadataOk() (*SavepointMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Savepoint) SetMetadata(v SavepointMetadata)`

SetMetadata sets Metadata field to given value.


### GetSpec

`func (o *Savepoint) GetSpec() SavepointSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *Savepoint) GetSpecOk() (*SavepointSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *Savepoint) SetSpec(v SavepointSpec)`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *Savepoint) GetStatus() SavepointStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Savepoint) GetStatusOk() (*SavepointStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Savepoint) SetStatus(v SavepointStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Savepoint) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


