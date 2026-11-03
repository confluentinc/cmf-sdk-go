# StatementStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phase** | **string** | The lifecycle phase of the statement. &#x60;PENDING&#x60;: the statement has been submitted, resumed or changed, and is waiting to start executing. &#x60;RECONCILING&#x60;: the statement&#39;s job is being restored, cannot be found, or its state cannot be read, for example after its JobManager restarted or when the statement is redeployed from a savepoint. &#x60;RESTARTING&#x60;: the statement is restarting. &#x60;RUNNING&#x60;: the statement is executing. &#x60;COMPLETED&#x60;: the statement finished executing successfully. &#x60;DELETING&#x60;: the statement is being deleted. &#x60;FAILING&#x60;: the statement has failed and cleanup is in progress. &#x60;FAILED&#x60;: the statement has failed. &#x60;STOPPED&#x60;: the statement has been stopped. &#x60;UNKNOWN&#x60;: the statement&#39;s phase could not be determined. | 
**Detail** | Pointer to **string** | Details about the execution status of the statement | [optional] 
**Traits** | Pointer to [**StatementTraits**](StatementTraits.md) |  | [optional] 
**ResourceSummary** | Pointer to [**ResourceUsageSummary**](ResourceUsageSummary.md) |  | [optional] 
**Warning** | Pointer to **string** | Warning message indicating the displayed status may be stale, for example, when the backing Kubernetes cluster is disconnected or decommissioned. | [optional] 
**SavepointSchedule** | Pointer to [**SavepointScheduleStatus**](SavepointScheduleStatus.md) |  | [optional] 
**LatestOffsets** | Pointer to **map[string]string** | The Kafka read positions this statement had reached, as a map of topic name to an offset string in the same format as the &#x60;scan.startup.specific-offsets&#x60; table option, for example &#x60;partition:0,offset:100;partition:1,offset:200&#x60;.  Not populated when the statement stops. The positions are read from the savepoint taken when the statement was last stopped, and only once another statement names this one in its &#x60;sql.tables.initial-offset-from&#x60; property; absent until then, even when that savepoint has completed. After the statement runs and is stopped again, absent until the newer savepoint is read the same way.  Only topics the statement read appear. A partition it had not read from yet appears as &#x60;offset:-2&#x60; when it was to start from the earliest offset; one that was to start from the latest or the committed offset is left out, and a topic whose partitions are all left out does not appear. A statement inheriting these offsets starts a partition left out of a listed topic from its committed group offset if there is one, otherwise from the earliest offset.  Available only while carry-over offsets is enabled (&#x60;cmf.sql.carry-over-offsets.enabled&#x60;), for detached statements reading Kafka tables, and when the statement&#39;s compute pool is configured to take a savepoint when a statement is stopped. Absent otherwise. | [optional] 
**LatestOffsetsTimestamp** | Pointer to **time.Time** | When &#x60;latestOffsets&#x60; was read from the savepoint. Absent while &#x60;latestOffsets&#x60; is absent. | [optional] 

## Methods

### NewStatementStatus

`func NewStatementStatus(phase string, ) *StatementStatus`

NewStatementStatus instantiates a new StatementStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatementStatusWithDefaults

`func NewStatementStatusWithDefaults() *StatementStatus`

NewStatementStatusWithDefaults instantiates a new StatementStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhase

`func (o *StatementStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *StatementStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *StatementStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.


### GetDetail

`func (o *StatementStatus) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *StatementStatus) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *StatementStatus) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *StatementStatus) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetTraits

`func (o *StatementStatus) GetTraits() StatementTraits`

GetTraits returns the Traits field if non-nil, zero value otherwise.

### GetTraitsOk

`func (o *StatementStatus) GetTraitsOk() (*StatementTraits, bool)`

GetTraitsOk returns a tuple with the Traits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraits

`func (o *StatementStatus) SetTraits(v StatementTraits)`

SetTraits sets Traits field to given value.

### HasTraits

`func (o *StatementStatus) HasTraits() bool`

HasTraits returns a boolean if a field has been set.

### GetResourceSummary

`func (o *StatementStatus) GetResourceSummary() ResourceUsageSummary`

GetResourceSummary returns the ResourceSummary field if non-nil, zero value otherwise.

### GetResourceSummaryOk

`func (o *StatementStatus) GetResourceSummaryOk() (*ResourceUsageSummary, bool)`

GetResourceSummaryOk returns a tuple with the ResourceSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSummary

`func (o *StatementStatus) SetResourceSummary(v ResourceUsageSummary)`

SetResourceSummary sets ResourceSummary field to given value.

### HasResourceSummary

`func (o *StatementStatus) HasResourceSummary() bool`

HasResourceSummary returns a boolean if a field has been set.

### GetWarning

`func (o *StatementStatus) GetWarning() string`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *StatementStatus) GetWarningOk() (*string, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *StatementStatus) SetWarning(v string)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *StatementStatus) HasWarning() bool`

HasWarning returns a boolean if a field has been set.

### GetSavepointSchedule

`func (o *StatementStatus) GetSavepointSchedule() SavepointScheduleStatus`

GetSavepointSchedule returns the SavepointSchedule field if non-nil, zero value otherwise.

### GetSavepointScheduleOk

`func (o *StatementStatus) GetSavepointScheduleOk() (*SavepointScheduleStatus, bool)`

GetSavepointScheduleOk returns a tuple with the SavepointSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavepointSchedule

`func (o *StatementStatus) SetSavepointSchedule(v SavepointScheduleStatus)`

SetSavepointSchedule sets SavepointSchedule field to given value.

### HasSavepointSchedule

`func (o *StatementStatus) HasSavepointSchedule() bool`

HasSavepointSchedule returns a boolean if a field has been set.

### GetLatestOffsets

`func (o *StatementStatus) GetLatestOffsets() map[string]string`

GetLatestOffsets returns the LatestOffsets field if non-nil, zero value otherwise.

### GetLatestOffsetsOk

`func (o *StatementStatus) GetLatestOffsetsOk() (*map[string]string, bool)`

GetLatestOffsetsOk returns a tuple with the LatestOffsets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestOffsets

`func (o *StatementStatus) SetLatestOffsets(v map[string]string)`

SetLatestOffsets sets LatestOffsets field to given value.

### HasLatestOffsets

`func (o *StatementStatus) HasLatestOffsets() bool`

HasLatestOffsets returns a boolean if a field has been set.

### GetLatestOffsetsTimestamp

`func (o *StatementStatus) GetLatestOffsetsTimestamp() time.Time`

GetLatestOffsetsTimestamp returns the LatestOffsetsTimestamp field if non-nil, zero value otherwise.

### GetLatestOffsetsTimestampOk

`func (o *StatementStatus) GetLatestOffsetsTimestampOk() (*time.Time, bool)`

GetLatestOffsetsTimestampOk returns a tuple with the LatestOffsetsTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestOffsetsTimestamp

`func (o *StatementStatus) SetLatestOffsetsTimestamp(v time.Time)`

SetLatestOffsetsTimestamp sets LatestOffsetsTimestamp field to given value.

### HasLatestOffsetsTimestamp

`func (o *StatementStatus) HasLatestOffsetsTimestamp() bool`

HasLatestOffsetsTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


