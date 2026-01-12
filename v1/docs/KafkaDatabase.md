# KafkaDatabase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | API version for spec | 
**Kind** | **string** | Kind of resource - set to resource type | 
**Metadata** | [**DatabaseMetadata**](DatabaseMetadata.md) |  | 
**Spec** | [**KafkaDatabaseSpec**](KafkaDatabaseSpec.md) |  | 
**Status** | Pointer to [**KafkaDatabaseStatus**](KafkaDatabaseStatus.md) |  | [optional] 

## Methods

### NewKafkaDatabase

`func NewKafkaDatabase(apiVersion string, kind string, metadata DatabaseMetadata, spec KafkaDatabaseSpec, ) *KafkaDatabase`

NewKafkaDatabase instantiates a new KafkaDatabase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaDatabaseWithDefaults

`func NewKafkaDatabaseWithDefaults() *KafkaDatabase`

NewKafkaDatabaseWithDefaults instantiates a new KafkaDatabase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *KafkaDatabase) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *KafkaDatabase) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *KafkaDatabase) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *KafkaDatabase) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *KafkaDatabase) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *KafkaDatabase) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *KafkaDatabase) GetMetadata() DatabaseMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *KafkaDatabase) GetMetadataOk() (*DatabaseMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *KafkaDatabase) SetMetadata(v DatabaseMetadata)`

SetMetadata sets Metadata field to given value.


### GetSpec

`func (o *KafkaDatabase) GetSpec() KafkaDatabaseSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *KafkaDatabase) GetSpecOk() (*KafkaDatabaseSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *KafkaDatabase) SetSpec(v KafkaDatabaseSpec)`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *KafkaDatabase) GetStatus() KafkaDatabaseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KafkaDatabase) GetStatusOk() (*KafkaDatabaseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KafkaDatabase) SetStatus(v KafkaDatabaseStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KafkaDatabase) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


