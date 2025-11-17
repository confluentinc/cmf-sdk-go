# KafkaDatabaseSpecKafkaCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionConfig** | **map[string]string** | connection options for the Kafka client | 
**ConnectionSecretId** | Pointer to **string** | an identifier to look up a secret that contains the connection credentials | [optional] 

## Methods

### NewKafkaDatabaseSpecKafkaCluster

`func NewKafkaDatabaseSpecKafkaCluster(connectionConfig map[string]string, ) *KafkaDatabaseSpecKafkaCluster`

NewKafkaDatabaseSpecKafkaCluster instantiates a new KafkaDatabaseSpecKafkaCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaDatabaseSpecKafkaClusterWithDefaults

`func NewKafkaDatabaseSpecKafkaClusterWithDefaults() *KafkaDatabaseSpecKafkaCluster`

NewKafkaDatabaseSpecKafkaClusterWithDefaults instantiates a new KafkaDatabaseSpecKafkaCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionConfig

`func (o *KafkaDatabaseSpecKafkaCluster) GetConnectionConfig() map[string]string`

GetConnectionConfig returns the ConnectionConfig field if non-nil, zero value otherwise.

### GetConnectionConfigOk

`func (o *KafkaDatabaseSpecKafkaCluster) GetConnectionConfigOk() (*map[string]string, bool)`

GetConnectionConfigOk returns a tuple with the ConnectionConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionConfig

`func (o *KafkaDatabaseSpecKafkaCluster) SetConnectionConfig(v map[string]string)`

SetConnectionConfig sets ConnectionConfig field to given value.


### GetConnectionSecretId

`func (o *KafkaDatabaseSpecKafkaCluster) GetConnectionSecretId() string`

GetConnectionSecretId returns the ConnectionSecretId field if non-nil, zero value otherwise.

### GetConnectionSecretIdOk

`func (o *KafkaDatabaseSpecKafkaCluster) GetConnectionSecretIdOk() (*string, bool)`

GetConnectionSecretIdOk returns a tuple with the ConnectionSecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSecretId

`func (o *KafkaDatabaseSpecKafkaCluster) SetConnectionSecretId(v string)`

SetConnectionSecretId sets ConnectionSecretId field to given value.

### HasConnectionSecretId

`func (o *KafkaDatabaseSpecKafkaCluster) HasConnectionSecretId() bool`

HasConnectionSecretId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


