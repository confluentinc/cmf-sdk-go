# KafkaCatalogSpecSrInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionConfig** | **map[string]string** | connection options for the SR client | 
**ConnectionSecretId** | Pointer to **string** | an identifier to look up a Kubernetes secret that contains the connection credentials | [optional] 

## Methods

### NewKafkaCatalogSpecSrInstance

`func NewKafkaCatalogSpecSrInstance(connectionConfig map[string]string, ) *KafkaCatalogSpecSrInstance`

NewKafkaCatalogSpecSrInstance instantiates a new KafkaCatalogSpecSrInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaCatalogSpecSrInstanceWithDefaults

`func NewKafkaCatalogSpecSrInstanceWithDefaults() *KafkaCatalogSpecSrInstance`

NewKafkaCatalogSpecSrInstanceWithDefaults instantiates a new KafkaCatalogSpecSrInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionConfig

`func (o *KafkaCatalogSpecSrInstance) GetConnectionConfig() map[string]string`

GetConnectionConfig returns the ConnectionConfig field if non-nil, zero value otherwise.

### GetConnectionConfigOk

`func (o *KafkaCatalogSpecSrInstance) GetConnectionConfigOk() (*map[string]string, bool)`

GetConnectionConfigOk returns a tuple with the ConnectionConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionConfig

`func (o *KafkaCatalogSpecSrInstance) SetConnectionConfig(v map[string]string)`

SetConnectionConfig sets ConnectionConfig field to given value.


### GetConnectionSecretId

`func (o *KafkaCatalogSpecSrInstance) GetConnectionSecretId() string`

GetConnectionSecretId returns the ConnectionSecretId field if non-nil, zero value otherwise.

### GetConnectionSecretIdOk

`func (o *KafkaCatalogSpecSrInstance) GetConnectionSecretIdOk() (*string, bool)`

GetConnectionSecretIdOk returns a tuple with the ConnectionSecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSecretId

`func (o *KafkaCatalogSpecSrInstance) SetConnectionSecretId(v string)`

SetConnectionSecretId sets ConnectionSecretId field to given value.

### HasConnectionSecretId

`func (o *KafkaCatalogSpecSrInstance) HasConnectionSecretId() bool`

HasConnectionSecretId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


