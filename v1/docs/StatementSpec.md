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
**SavepointSchedule** | Pointer to [**SavepointScheduleConfig**](SavepointScheduleConfig.md) |  | [optional] 
**UpgradeMode** | Pointer to **string** | Controls whether the statement&#39;s execution state is preserved across a stop and resume, or when the statement&#39;s job has to be resubmitted (for example, after the compute pool loses track of a running job and has to recreate it).  This is separate from Flink&#39;s own recovery from task failures while the job keeps running: that recovery is governed by the job&#39;s checkpointing and restart-strategy settings and happens regardless of this field.  * &#x60;savepoint&#x60;: CMF takes a savepoint when the statement is stopped or   resubmitted, and restores from it on resume. Requires both   &#x60;state.checkpoints.dir&#x60; and &#x60;state.savepoints.dir&#x60; to be configured on   the compute pool the statement runs on (directly or via the   environment&#39;s &#x60;computePoolDefaults&#x60;); setting these directories on the   statement itself has no effect on this field. * &#x60;stateless&#x60;: no state is preserved across a stop and resume, or a   resubmission. The statement comes back up with empty state even if a   checkpoint or savepoint is available. * &#x60;last-state&#x60;: on resubmission, the statement resumes from its most   recent checkpoint without CMF taking a new savepoint. Requires   &#x60;state.checkpoints.dir&#x60; to be configured on the compute pool the   statement runs on (directly or via the environment&#39;s   &#x60;computePoolDefaults&#x60;). Not supported for statements that run on a   shared compute pool.  Optional. When omitted, CMF selects &#x60;savepoint&#x60; if both a checkpoint directory and a savepoint directory are configured on the compute pool, and &#x60;stateless&#x60; otherwise.  Changing this field alone does not restart a running statement. The new value takes effect on the next stop and resume, or the next time the statement&#39;s job is resubmitted.  | [optional] 

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

### GetSavepointSchedule

`func (o *StatementSpec) GetSavepointSchedule() SavepointScheduleConfig`

GetSavepointSchedule returns the SavepointSchedule field if non-nil, zero value otherwise.

### GetSavepointScheduleOk

`func (o *StatementSpec) GetSavepointScheduleOk() (*SavepointScheduleConfig, bool)`

GetSavepointScheduleOk returns a tuple with the SavepointSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavepointSchedule

`func (o *StatementSpec) SetSavepointSchedule(v SavepointScheduleConfig)`

SetSavepointSchedule sets SavepointSchedule field to given value.

### HasSavepointSchedule

`func (o *StatementSpec) HasSavepointSchedule() bool`

HasSavepointSchedule returns a boolean if a field has been set.

### GetUpgradeMode

`func (o *StatementSpec) GetUpgradeMode() string`

GetUpgradeMode returns the UpgradeMode field if non-nil, zero value otherwise.

### GetUpgradeModeOk

`func (o *StatementSpec) GetUpgradeModeOk() (*string, bool)`

GetUpgradeModeOk returns a tuple with the UpgradeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeMode

`func (o *StatementSpec) SetUpgradeMode(v string)`

SetUpgradeMode sets UpgradeMode field to given value.

### HasUpgradeMode

`func (o *StatementSpec) HasUpgradeMode() bool`

HasUpgradeMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


