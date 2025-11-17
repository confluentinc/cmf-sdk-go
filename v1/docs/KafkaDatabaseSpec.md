# KafkaDatabaseSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KafkaCluster** | [**KafkaDatabaseSpecKafkaCluster**](KafkaDatabaseSpecKafkaCluster.md) |  | 
**AlterEnvironments** | Pointer to **[]string** | List of environments that have permission to alter the tables of this database | [optional] 

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


### GetAlterEnvironments

`func (o *KafkaDatabaseSpec) GetAlterEnvironments() []string`

GetAlterEnvironments returns the AlterEnvironments field if non-nil, zero value otherwise.

### GetAlterEnvironmentsOk

`func (o *KafkaDatabaseSpec) GetAlterEnvironmentsOk() (*[]string, bool)`

GetAlterEnvironmentsOk returns a tuple with the AlterEnvironments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlterEnvironments

`func (o *KafkaDatabaseSpec) SetAlterEnvironments(v []string)`

SetAlterEnvironments sets AlterEnvironments field to given value.

### HasAlterEnvironments

`func (o *KafkaDatabaseSpec) HasAlterEnvironments() bool`

HasAlterEnvironments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


