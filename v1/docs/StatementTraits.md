# StatementTraits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SqlKind** | Pointer to **string** | The kind of SQL statement | [optional] 
**IsBounded** | Pointer to **bool** | Whether the result of the statement is bounded | [optional] 
**IsAppendOnly** | Pointer to **bool** | Whether the result of the statement is append only | [optional] 
**UpsertColumns** | Pointer to **[]int32** | The column indexes that are updated by the statement | [optional] 
**Schema** | Pointer to [**ResultSchema**](ResultSchema.md) |  | [optional] 
**UdfJarUris** | Pointer to **[]string** | UDF JAR URIs (pre-resolution, e.g. &#x60;cmf://&#x60;, &#x60;s3://&#x60;) that this statement references via the environment catalog. Empty or absent when the statement uses no env-catalog UDFs. Surfaced so callers can see which JARs the statement loads at runtime. | [optional] 
**ConnectorJarUris** | Pointer to **[]string** | Connector and format artifact JAR URIs (e.g. &#x60;s3://&#x60;) that this statement&#39;s referenced tables require, one entry per connector or format actually used, and the JAR URI (e.g. &#x60;file://&#x60;) of the provider of each model the statement calls, for providers that ship as a separate JAR. Empty or absent when the statement needs none of these JARs. Surfaced so callers can see which JARs the statement loads at runtime. | [optional] 
**InheritedOffsetTopics** | Pointer to **[]string** | Topics whose Kafka sources this statement starts from the offsets it inherited through &#x60;sql.tables.initial-offset-from&#x60;. Empty when it inherits offsets but every source sets its own start, for example with a &#x60;scan.startup.mode&#x60; OPTIONS hint. Absent when it inherits no offsets or still waits for them, and when CMF cannot tell which topics its sources start from, for example for a statement started by an earlier CMF version; recovery then treats a statement that inherits offsets as starting its Kafka sources from them. | [optional] [readonly] 
**UpgradeMode** | Pointer to **string** | The upgrade mode that has been applied to the statement&#39;s deployment, shown so callers can see the value that took effect without inspecting the underlying Flink resources. When &#x60;spec.upgradeMode&#x60; is set, this matches it, except on a stopped SELECT or a stopped batch statement, which use &#x60;stateless&#x60;; otherwise it is the value selected by the default rules described on &#x60;spec.upgradeMode&#x60;. Populated once the statement has been deployed at least once; absent before the first deployment attempt. | [optional] 
**JobId** | Pointer to **string** | The ID of the Flink job most recently submitted for this statement. Absent for interactive SELECT statements, and until the statement&#39;s first job is submitted.  Use it to identify the job&#39;s checkpoints, stored under &#x60;&lt;state.checkpoints.dir&gt;/&lt;jobId&gt;/&#x60;, for example to restore the statement from its latest checkpoint after a disaster. A newly submitted job may not have completed a checkpoint yet.  The ID is kept while the statement is stopped. It changes when the statement is resumed or redeployed with any upgrade mode other than &#x60;last-state&#x60;. | [optional] [readonly] 
**ParkedByRecovery** | Pointer to **string** | Set when CMF, recreating the statement&#39;s missing Flink resources after a restart, stopped the statement instead of starting it without the state it needs, or instead of letting it reprocess input. &#x60;status.detail&#x60; says how to recover it. Cleared when the statement is resumed, or set to FAILED because its compute pool was deleted. Values: &#x60;no-restore-point&#x60;: the statement has no savepoint to restore from; &#x60;restore-point-too-old&#x60;: the restore point CMF would use is older than &#x60;cmf.recovery.max-restore-point-age&#x60;; &#x60;start-from-savepoint-unresolved&#x60;: the savepoint set in &#x60;spec.startFromSavepoint&#x60; no longer resolves; &#x60;carry-over-replay&#x60;: the statement starts Kafka sources from the offsets it inherited from another statement (&#x60;sql.tables.initial-offset-from&#x60;, see &#x60;inheritedOffsetTopics&#x60;) and it has no savepoint CMF can restore it from (none exists, the &#x60;startFromSavepoint&#x60; it last started from no longer resolves, or each one is older than &#x60;cmf.recovery.max-restore-point-age&#x60; or was taken before that &#x60;startFromSavepoint&#x60; was applied), so starting it would reprocess all their input since the offsets it was created with. Resuming the statement starts it anyway; &#x60;batch-rerun&#x60;: the statement is an &#x60;INSERT INTO&#x60; or &#x60;STATEMENT SET&#x60; that runs in batch mode and its job already ran. A batch job cannot resume, so starting it again would run it from the start and could write its output a second time. The statement cannot be resumed; create a new statement to run it again. | [optional] [readonly] 
**RuntimeMode** | Pointer to **string** | The runtime mode the statement&#39;s Flink job runs in: &#x60;execution.runtime-mode&#x60; from the statement&#39;s &#x60;flinkConfiguration&#x60; merged with the Environment&#39;s statement defaults (statementDefaults.interactive for a SELECT, statementDefaults.detached for an INSERT INTO or a statement set), with the defaults&#39; value taking precedence. STREAMING when neither sets it. Set when the statement is created and does not change. Absent for statements that do not run a Flink job (for example SHOW, DESCRIBE or CREATE TABLE), for statements that failed to compile, and for statements created by an earlier CMF version. | [optional] [readonly] 

## Methods

### NewStatementTraits

`func NewStatementTraits() *StatementTraits`

NewStatementTraits instantiates a new StatementTraits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatementTraitsWithDefaults

`func NewStatementTraitsWithDefaults() *StatementTraits`

NewStatementTraitsWithDefaults instantiates a new StatementTraits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSqlKind

`func (o *StatementTraits) GetSqlKind() string`

GetSqlKind returns the SqlKind field if non-nil, zero value otherwise.

### GetSqlKindOk

`func (o *StatementTraits) GetSqlKindOk() (*string, bool)`

GetSqlKindOk returns a tuple with the SqlKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlKind

`func (o *StatementTraits) SetSqlKind(v string)`

SetSqlKind sets SqlKind field to given value.

### HasSqlKind

`func (o *StatementTraits) HasSqlKind() bool`

HasSqlKind returns a boolean if a field has been set.

### GetIsBounded

`func (o *StatementTraits) GetIsBounded() bool`

GetIsBounded returns the IsBounded field if non-nil, zero value otherwise.

### GetIsBoundedOk

`func (o *StatementTraits) GetIsBoundedOk() (*bool, bool)`

GetIsBoundedOk returns a tuple with the IsBounded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBounded

`func (o *StatementTraits) SetIsBounded(v bool)`

SetIsBounded sets IsBounded field to given value.

### HasIsBounded

`func (o *StatementTraits) HasIsBounded() bool`

HasIsBounded returns a boolean if a field has been set.

### GetIsAppendOnly

`func (o *StatementTraits) GetIsAppendOnly() bool`

GetIsAppendOnly returns the IsAppendOnly field if non-nil, zero value otherwise.

### GetIsAppendOnlyOk

`func (o *StatementTraits) GetIsAppendOnlyOk() (*bool, bool)`

GetIsAppendOnlyOk returns a tuple with the IsAppendOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAppendOnly

`func (o *StatementTraits) SetIsAppendOnly(v bool)`

SetIsAppendOnly sets IsAppendOnly field to given value.

### HasIsAppendOnly

`func (o *StatementTraits) HasIsAppendOnly() bool`

HasIsAppendOnly returns a boolean if a field has been set.

### GetUpsertColumns

`func (o *StatementTraits) GetUpsertColumns() []int32`

GetUpsertColumns returns the UpsertColumns field if non-nil, zero value otherwise.

### GetUpsertColumnsOk

`func (o *StatementTraits) GetUpsertColumnsOk() (*[]int32, bool)`

GetUpsertColumnsOk returns a tuple with the UpsertColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpsertColumns

`func (o *StatementTraits) SetUpsertColumns(v []int32)`

SetUpsertColumns sets UpsertColumns field to given value.

### HasUpsertColumns

`func (o *StatementTraits) HasUpsertColumns() bool`

HasUpsertColumns returns a boolean if a field has been set.

### GetSchema

`func (o *StatementTraits) GetSchema() ResultSchema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *StatementTraits) GetSchemaOk() (*ResultSchema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *StatementTraits) SetSchema(v ResultSchema)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *StatementTraits) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetUdfJarUris

`func (o *StatementTraits) GetUdfJarUris() []string`

GetUdfJarUris returns the UdfJarUris field if non-nil, zero value otherwise.

### GetUdfJarUrisOk

`func (o *StatementTraits) GetUdfJarUrisOk() (*[]string, bool)`

GetUdfJarUrisOk returns a tuple with the UdfJarUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdfJarUris

`func (o *StatementTraits) SetUdfJarUris(v []string)`

SetUdfJarUris sets UdfJarUris field to given value.

### HasUdfJarUris

`func (o *StatementTraits) HasUdfJarUris() bool`

HasUdfJarUris returns a boolean if a field has been set.

### GetConnectorJarUris

`func (o *StatementTraits) GetConnectorJarUris() []string`

GetConnectorJarUris returns the ConnectorJarUris field if non-nil, zero value otherwise.

### GetConnectorJarUrisOk

`func (o *StatementTraits) GetConnectorJarUrisOk() (*[]string, bool)`

GetConnectorJarUrisOk returns a tuple with the ConnectorJarUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorJarUris

`func (o *StatementTraits) SetConnectorJarUris(v []string)`

SetConnectorJarUris sets ConnectorJarUris field to given value.

### HasConnectorJarUris

`func (o *StatementTraits) HasConnectorJarUris() bool`

HasConnectorJarUris returns a boolean if a field has been set.

### GetInheritedOffsetTopics

`func (o *StatementTraits) GetInheritedOffsetTopics() []string`

GetInheritedOffsetTopics returns the InheritedOffsetTopics field if non-nil, zero value otherwise.

### GetInheritedOffsetTopicsOk

`func (o *StatementTraits) GetInheritedOffsetTopicsOk() (*[]string, bool)`

GetInheritedOffsetTopicsOk returns a tuple with the InheritedOffsetTopics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritedOffsetTopics

`func (o *StatementTraits) SetInheritedOffsetTopics(v []string)`

SetInheritedOffsetTopics sets InheritedOffsetTopics field to given value.

### HasInheritedOffsetTopics

`func (o *StatementTraits) HasInheritedOffsetTopics() bool`

HasInheritedOffsetTopics returns a boolean if a field has been set.

### SetInheritedOffsetTopicsNil

`func (o *StatementTraits) SetInheritedOffsetTopicsNil(b bool)`

 SetInheritedOffsetTopicsNil sets the value for InheritedOffsetTopics to be an explicit nil

### UnsetInheritedOffsetTopics
`func (o *StatementTraits) UnsetInheritedOffsetTopics()`

UnsetInheritedOffsetTopics ensures that no value is present for InheritedOffsetTopics, not even an explicit nil
### GetUpgradeMode

`func (o *StatementTraits) GetUpgradeMode() string`

GetUpgradeMode returns the UpgradeMode field if non-nil, zero value otherwise.

### GetUpgradeModeOk

`func (o *StatementTraits) GetUpgradeModeOk() (*string, bool)`

GetUpgradeModeOk returns a tuple with the UpgradeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeMode

`func (o *StatementTraits) SetUpgradeMode(v string)`

SetUpgradeMode sets UpgradeMode field to given value.

### HasUpgradeMode

`func (o *StatementTraits) HasUpgradeMode() bool`

HasUpgradeMode returns a boolean if a field has been set.

### GetJobId

`func (o *StatementTraits) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *StatementTraits) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *StatementTraits) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *StatementTraits) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetParkedByRecovery

`func (o *StatementTraits) GetParkedByRecovery() string`

GetParkedByRecovery returns the ParkedByRecovery field if non-nil, zero value otherwise.

### GetParkedByRecoveryOk

`func (o *StatementTraits) GetParkedByRecoveryOk() (*string, bool)`

GetParkedByRecoveryOk returns a tuple with the ParkedByRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParkedByRecovery

`func (o *StatementTraits) SetParkedByRecovery(v string)`

SetParkedByRecovery sets ParkedByRecovery field to given value.

### HasParkedByRecovery

`func (o *StatementTraits) HasParkedByRecovery() bool`

HasParkedByRecovery returns a boolean if a field has been set.

### GetRuntimeMode

`func (o *StatementTraits) GetRuntimeMode() string`

GetRuntimeMode returns the RuntimeMode field if non-nil, zero value otherwise.

### GetRuntimeModeOk

`func (o *StatementTraits) GetRuntimeModeOk() (*string, bool)`

GetRuntimeModeOk returns a tuple with the RuntimeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeMode

`func (o *StatementTraits) SetRuntimeMode(v string)`

SetRuntimeMode sets RuntimeMode field to given value.

### HasRuntimeMode

`func (o *StatementTraits) HasRuntimeMode() bool`

HasRuntimeMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


