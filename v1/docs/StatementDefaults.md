# StatementDefaults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlinkConfiguration** | Pointer to **map[string]string** | Default Flink configuration for Statements.  Some keys are reserved and cannot be set: &#x60;web.submit.jar-sources&#x60;, &#x60;pipeline.jars&#x60;, &#x60;pipeline.classpaths&#x60;, and every key beginning with &#x60;user.artifacts.&#x60; or &#x60;kubernetes.operator.cmf.&#x60;. An environment that sets one is rejected with an error naming the key.  A Statement keeps the &#x60;execution.runtime-mode&#x60; it was created with. An update that changes this key is rejected while Statements that can still be resumed or rescaled would run with the new value; the error names them.  | [optional] 
**SavepointSchedule** | Pointer to [**SavepointScheduleConfig**](SavepointScheduleConfig.md) |  | [optional] 

## Methods

### NewStatementDefaults

`func NewStatementDefaults() *StatementDefaults`

NewStatementDefaults instantiates a new StatementDefaults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatementDefaultsWithDefaults

`func NewStatementDefaultsWithDefaults() *StatementDefaults`

NewStatementDefaultsWithDefaults instantiates a new StatementDefaults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlinkConfiguration

`func (o *StatementDefaults) GetFlinkConfiguration() map[string]string`

GetFlinkConfiguration returns the FlinkConfiguration field if non-nil, zero value otherwise.

### GetFlinkConfigurationOk

`func (o *StatementDefaults) GetFlinkConfigurationOk() (*map[string]string, bool)`

GetFlinkConfigurationOk returns a tuple with the FlinkConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlinkConfiguration

`func (o *StatementDefaults) SetFlinkConfiguration(v map[string]string)`

SetFlinkConfiguration sets FlinkConfiguration field to given value.

### HasFlinkConfiguration

`func (o *StatementDefaults) HasFlinkConfiguration() bool`

HasFlinkConfiguration returns a boolean if a field has been set.

### GetSavepointSchedule

`func (o *StatementDefaults) GetSavepointSchedule() SavepointScheduleConfig`

GetSavepointSchedule returns the SavepointSchedule field if non-nil, zero value otherwise.

### GetSavepointScheduleOk

`func (o *StatementDefaults) GetSavepointScheduleOk() (*SavepointScheduleConfig, bool)`

GetSavepointScheduleOk returns a tuple with the SavepointSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavepointSchedule

`func (o *StatementDefaults) SetSavepointSchedule(v SavepointScheduleConfig)`

SetSavepointSchedule sets SavepointSchedule field to given value.

### HasSavepointSchedule

`func (o *StatementDefaults) HasSavepointSchedule() bool`

HasSavepointSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


