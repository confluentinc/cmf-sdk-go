# KafkaDatabaseSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KafkaCluster** | [**KafkaDatabaseSpecKafkaCluster**](KafkaDatabaseSpecKafkaCluster.md) |  | 
**DdlEnvironments** | Pointer to **[]string** | List of environments that have permission to execute DDL statements on the tables of this database | [optional] 

## Methods

### NewKafkaDatabaseSpec

`func NewKafkaDatabaseSpec(kafkaCluster KafkaDatabaseSpecKafkaCluster, ) *KafkaDatabaseSpec`

NewKafkaDatabaseSpec instantiates a new KafkaDatabaseSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaDatabaseSpecWithDefaults

`func NewKafkaDatabaseSpecWithDefaults() *KafkaDatabaseSpec`

NewKafkaDatabaseSpecWithDefaults instantiates a new KafkaDatabaseSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKafkaCluster

`func (o *KafkaDatabaseSpec) GetKafkaCluster() KafkaDatabaseSpecKafkaCluster`

GetKafkaCluster returns the KafkaCluster field if non-nil, zero value otherwise.

### GetKafkaClusterOk

`func (o *KafkaDatabaseSpec) GetKafkaClusterOk() (*KafkaDatabaseSpecKafkaCluster, bool)`

GetKafkaClusterOk returns a tuple with the KafkaCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaCluster

`func (o *KafkaDatabaseSpec) SetKafkaCluster(v KafkaDatabaseSpecKafkaCluster)`

SetKafkaCluster sets KafkaCluster field to given value.


### GetDdlEnvironments

`func (o *KafkaDatabaseSpec) GetDdlEnvironments() []string`

GetDdlEnvironments returns the DdlEnvironments field if non-nil, zero value otherwise.

### GetDdlEnvironmentsOk

`func (o *KafkaDatabaseSpec) GetDdlEnvironmentsOk() (*[]string, bool)`

GetDdlEnvironmentsOk returns a tuple with the DdlEnvironments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdlEnvironments

`func (o *KafkaDatabaseSpec) SetDdlEnvironments(v []string)`

SetDdlEnvironments sets DdlEnvironments field to given value.

### HasDdlEnvironments

`func (o *KafkaDatabaseSpec) HasDdlEnvironments() bool`

HasDdlEnvironments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


