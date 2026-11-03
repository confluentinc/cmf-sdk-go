# KafkaDatabaseSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KafkaCluster** | [**KafkaDatabaseSpecKafkaCluster**](KafkaDatabaseSpecKafkaCluster.md) |  | 
**DdlEnvironments** | Pointer to **[]string** | List of environments that have permission to execute DDL statements on the tables of this database | [optional] 
**SubjectNameStrategy** | Pointer to **string** | How the tables of this database find their Schema Registry subjects. TopicNameStrategy, the default, uses the &lt;topic&gt;-key and &lt;topic&gt;-value subjects. AssociatedNameStrategy uses the topic-subject associations of the Kafka cluster: CREATE TABLE has the brokers register the schemas, and tables read the subjects their topic is associated with.  AssociatedNameStrategy needs Schema Registry 8.3.1 or later, Kafka brokers with topic-subject associations enabled, and Flink 2.x. CMF accepts AssociatedNameStrategy only when it has a Flink 2.x SQL compiler (the default compilers include one). Statements on Flink 1.x compute pools fail when they use the tables of such a database, whether they name the database or run in it. A statement that sets a catalog but no database runs in the catalog&#39;s earliest-created database that its environment can connect to.  If an update leaves out this field, the database keeps its current strategy. Keep the strategy once the database has tables, since it decides where those tables find their schemas.  | [optional] 

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

### GetSubjectNameStrategy

`func (o *KafkaDatabaseSpec) GetSubjectNameStrategy() string`

GetSubjectNameStrategy returns the SubjectNameStrategy field if non-nil, zero value otherwise.

### GetSubjectNameStrategyOk

`func (o *KafkaDatabaseSpec) GetSubjectNameStrategyOk() (*string, bool)`

GetSubjectNameStrategyOk returns a tuple with the SubjectNameStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectNameStrategy

`func (o *KafkaDatabaseSpec) SetSubjectNameStrategy(v string)`

SetSubjectNameStrategy sets SubjectNameStrategy field to given value.

### HasSubjectNameStrategy

`func (o *KafkaDatabaseSpec) HasSubjectNameStrategy() bool`

HasSubjectNameStrategy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


