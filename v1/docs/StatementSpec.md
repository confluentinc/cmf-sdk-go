# StatementSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statement** | **string** | SQL statement | 
**Properties** | Pointer to **map[string]string** | Properties of the client session | [optional] 
**FlinkConfiguration** | Pointer to **map[string]string** | Flink configuration for the statement | [optional] 
**ComputePoolName** | **string** | Name of the ComputePool | 
**Parallelism** | Pointer to **int32** | Parallelism of the statement | [optional] 
**Stopped** | Pointer to **bool** | Whether the statement is stopped | [optional] 
**StartFromSavepoint** | Pointer to [**StatementStartFromSavepoint**](StatementStartFromSavepoint.md) |  | [optional] 

## Methods

### NewStatementSpec

`func NewStatementSpec(statement string, computePoolName string, ) *StatementSpec`

NewStatementSpec instantiates a new StatementSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatementSpecWithDefaults

`func NewStatementSpecWithDefaults() *StatementSpec`

NewStatementSpecWithDefaults instantiates a new StatementSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatement

`func (o *StatementSpec) GetStatement() string`

GetStatement returns the Statement field if non-nil, zero value otherwise.

### GetStatementOk

`func (o *StatementSpec) GetStatementOk() (*string, bool)`

GetStatementOk returns a tuple with the Statement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatement

`func (o *StatementSpec) SetStatement(v string)`

SetStatement sets Statement field to given value.


### GetProperties

`func (o *StatementSpec) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *StatementSpec) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *StatementSpec) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *StatementSpec) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetFlinkConfiguration

`func (o *StatementSpec) GetFlinkConfiguration() map[string]string`

GetFlinkConfiguration returns the FlinkConfiguration field if non-nil, zero value otherwise.

### GetFlinkConfigurationOk

`func (o *StatementSpec) GetFlinkConfigurationOk() (*map[string]string, bool)`

GetFlinkConfigurationOk returns a tuple with the FlinkConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlinkConfiguration

`func (o *StatementSpec) SetFlinkConfiguration(v map[string]string)`

SetFlinkConfiguration sets FlinkConfiguration field to given value.

### HasFlinkConfiguration

`func (o *StatementSpec) HasFlinkConfiguration() bool`

HasFlinkConfiguration returns a boolean if a field has been set.

### GetComputePoolName

`func (o *StatementSpec) GetComputePoolName() string`

GetComputePoolName returns the ComputePoolName field if non-nil, zero value otherwise.

### GetComputePoolNameOk

`func (o *StatementSpec) GetComputePoolNameOk() (*string, bool)`

GetComputePoolNameOk returns a tuple with the ComputePoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputePoolName

`func (o *StatementSpec) SetComputePoolName(v string)`

SetComputePoolName sets ComputePoolName field to given value.


### GetParallelism

`func (o *StatementSpec) GetParallelism() int32`

GetParallelism returns the Parallelism field if non-nil, zero value otherwise.

### GetParallelismOk

`func (o *StatementSpec) GetParallelismOk() (*int32, bool)`

GetParallelismOk returns a tuple with the Parallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelism

`func (o *StatementSpec) SetParallelism(v int32)`

SetParallelism sets Parallelism field to given value.

### HasParallelism

`func (o *StatementSpec) HasParallelism() bool`

HasParallelism returns a boolean if a field has been set.

### GetStopped

`func (o *StatementSpec) GetStopped() bool`

GetStopped returns the Stopped field if non-nil, zero value otherwise.

### GetStoppedOk

`func (o *StatementSpec) GetStoppedOk() (*bool, bool)`

GetStoppedOk returns a tuple with the Stopped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopped

`func (o *StatementSpec) SetStopped(v bool)`

SetStopped sets Stopped field to given value.

### HasStopped

`func (o *StatementSpec) HasStopped() bool`

HasStopped returns a boolean if a field has been set.

### GetStartFromSavepoint

`func (o *StatementSpec) GetStartFromSavepoint() StatementStartFromSavepoint`

GetStartFromSavepoint returns the StartFromSavepoint field if non-nil, zero value otherwise.

### GetStartFromSavepointOk

`func (o *StatementSpec) GetStartFromSavepointOk() (*StatementStartFromSavepoint, bool)`

GetStartFromSavepointOk returns a tuple with the StartFromSavepoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartFromSavepoint

`func (o *StatementSpec) SetStartFromSavepoint(v StatementStartFromSavepoint)`

SetStartFromSavepoint sets StartFromSavepoint field to given value.

### HasStartFromSavepoint

`func (o *StatementSpec) HasStartFromSavepoint() bool`

HasStartFromSavepoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


