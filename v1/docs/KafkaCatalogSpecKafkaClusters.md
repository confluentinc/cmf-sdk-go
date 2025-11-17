# KafkaCatalogSpecKafkaClusters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseName** | **string** | the database name under which the Kafka cluster is listed in the Catalog | 
**ConnectionConfig** | **map[string]string** | connection options for the Kafka client | 
**ConnectionSecretId** | Pointer to **string** | an identifier to look up a Kubernetes secret that contains the connection credentials | [optional] 

## Methods

### NewKafkaCatalogSpecKafkaClusters

`func NewKafkaCatalogSpecKafkaClusters(databaseName string, connectionConfig map[string]string, ) *KafkaCatalogSpecKafkaClusters`

NewKafkaCatalogSpecKafkaClusters instantiates a new KafkaCatalogSpecKafkaClusters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaCatalogSpecKafkaClustersWithDefaults

`func NewKafkaCatalogSpecKafkaClustersWithDefaults() *KafkaCatalogSpecKafkaClusters`

NewKafkaCatalogSpecKafkaClustersWithDefaults instantiates a new KafkaCatalogSpecKafkaClusters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseName

`func (o *KafkaCatalogSpecKafkaClusters) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *KafkaCatalogSpecKafkaClusters) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *KafkaCatalogSpecKafkaClusters) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.


### GetConnectionConfig

`func (o *KafkaCatalogSpecKafkaClusters) GetConnectionConfig() map[string]string`

GetConnectionConfig returns the ConnectionConfig field if non-nil, zero value otherwise.

### GetConnectionConfigOk

`func (o *KafkaCatalogSpecKafkaClusters) GetConnectionConfigOk() (*map[string]string, bool)`

GetConnectionConfigOk returns a tuple with the ConnectionConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionConfig

`func (o *KafkaCatalogSpecKafkaClusters) SetConnectionConfig(v map[string]string)`

SetConnectionConfig sets ConnectionConfig field to given value.


### GetConnectionSecretId

`func (o *KafkaCatalogSpecKafkaClusters) GetConnectionSecretId() string`

GetConnectionSecretId returns the ConnectionSecretId field if non-nil, zero value otherwise.

### GetConnectionSecretIdOk

`func (o *KafkaCatalogSpecKafkaClusters) GetConnectionSecretIdOk() (*string, bool)`

GetConnectionSecretIdOk returns a tuple with the ConnectionSecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSecretId

`func (o *KafkaCatalogSpecKafkaClusters) SetConnectionSecretId(v string)`

SetConnectionSecretId sets ConnectionSecretId field to given value.

### HasConnectionSecretId

`func (o *KafkaCatalogSpecKafkaClusters) HasConnectionSecretId() bool`

HasConnectionSecretId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


