# SavepointScheduleConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CronExpression** | Pointer to **string** | Six-field cron expression with seconds resolution, in the order: second, minute, hour, day-of-month, month, day-of-week. Minimum interval between triggers is 5 minutes.  | [optional] 
**Timezone** | Pointer to **string** | IANA time zone ID used to evaluate the cron expression. When omitted, UTC is used. No schema-level default is declared on purpose: environment-level defaults are merged over resource-level settings field by field, so a materialized default here would silently override the resource&#39;s explicit value.  | [optional] 
**Paused** | Pointer to **bool** | When true, the schedule is registered but no new savepoints are triggered. When omitted, the schedule is not paused.  | [optional] 
**CatchUpOnRestart** | Pointer to **bool** | When true, if CMF restarts after missing at least one scheduled trigger, fire exactly one catch-up trigger immediately on startup. When omitted, no catch-up is performed.  | [optional] 
**JitterSeconds** | Pointer to **int32** | Upper bound (in seconds) for the random offset added to each trigger time. Null or 0 means the service auto-computes a bound based on the cron interval; a positive value sets the bound explicitly.  | [optional] 
**BlackoutWindows** | Pointer to [**[]BlackoutWindow**](BlackoutWindow.md) | Time ranges during which periodic savepoint triggers are suppressed. Missed triggers during blackouts are skipped, not retried.  | [optional] 
**FreshnessThresholdHours** | Pointer to **int32** | Maximum acceptable gap in hours between successful periodic savepoints. CMF logs a SAVEPOINT_SLA_VIOLATION warning (visible in the CMF server log, not via the API) when the threshold is exceeded. When omitted or 0, the check is disabled.  | [optional] 
**FormatType** | Pointer to **string** | Savepoint format (see the Flink savepoint documentation) used for every scheduled savepoint this schedule triggers. No schema-level default is declared on purpose: environment-level defaults are merged over resource-level settings field by field, so a materialized default here would silently override the resource&#39;s explicit value. A null value is treated as CANONICAL when the savepoint is triggered.  | [optional] 
**Retention** | Pointer to [**SavepointRetention**](SavepointRetention.md) |  | [optional] 

## Methods

### NewSavepointScheduleConfig

`func NewSavepointScheduleConfig() *SavepointScheduleConfig`

NewSavepointScheduleConfig instantiates a new SavepointScheduleConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSavepointScheduleConfigWithDefaults

`func NewSavepointScheduleConfigWithDefaults() *SavepointScheduleConfig`

NewSavepointScheduleConfigWithDefaults instantiates a new SavepointScheduleConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCronExpression

`func (o *SavepointScheduleConfig) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *SavepointScheduleConfig) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *SavepointScheduleConfig) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.

### HasCronExpression

`func (o *SavepointScheduleConfig) HasCronExpression() bool`

HasCronExpression returns a boolean if a field has been set.

### GetTimezone

`func (o *SavepointScheduleConfig) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *SavepointScheduleConfig) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *SavepointScheduleConfig) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *SavepointScheduleConfig) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetPaused

`func (o *SavepointScheduleConfig) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *SavepointScheduleConfig) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *SavepointScheduleConfig) SetPaused(v bool)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *SavepointScheduleConfig) HasPaused() bool`

HasPaused returns a boolean if a field has been set.

### GetCatchUpOnRestart

`func (o *SavepointScheduleConfig) GetCatchUpOnRestart() bool`

GetCatchUpOnRestart returns the CatchUpOnRestart field if non-nil, zero value otherwise.

### GetCatchUpOnRestartOk

`func (o *SavepointScheduleConfig) GetCatchUpOnRestartOk() (*bool, bool)`

GetCatchUpOnRestartOk returns a tuple with the CatchUpOnRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatchUpOnRestart

`func (o *SavepointScheduleConfig) SetCatchUpOnRestart(v bool)`

SetCatchUpOnRestart sets CatchUpOnRestart field to given value.

### HasCatchUpOnRestart

`func (o *SavepointScheduleConfig) HasCatchUpOnRestart() bool`

HasCatchUpOnRestart returns a boolean if a field has been set.

### GetJitterSeconds

`func (o *SavepointScheduleConfig) GetJitterSeconds() int32`

GetJitterSeconds returns the JitterSeconds field if non-nil, zero value otherwise.

### GetJitterSecondsOk

`func (o *SavepointScheduleConfig) GetJitterSecondsOk() (*int32, bool)`

GetJitterSecondsOk returns a tuple with the JitterSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitterSeconds

`func (o *SavepointScheduleConfig) SetJitterSeconds(v int32)`

SetJitterSeconds sets JitterSeconds field to given value.

### HasJitterSeconds

`func (o *SavepointScheduleConfig) HasJitterSeconds() bool`

HasJitterSeconds returns a boolean if a field has been set.

### GetBlackoutWindows

`func (o *SavepointScheduleConfig) GetBlackoutWindows() []BlackoutWindow`

GetBlackoutWindows returns the BlackoutWindows field if non-nil, zero value otherwise.

### GetBlackoutWindowsOk

`func (o *SavepointScheduleConfig) GetBlackoutWindowsOk() (*[]BlackoutWindow, bool)`

GetBlackoutWindowsOk returns a tuple with the BlackoutWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackoutWindows

`func (o *SavepointScheduleConfig) SetBlackoutWindows(v []BlackoutWindow)`

SetBlackoutWindows sets BlackoutWindows field to given value.

### HasBlackoutWindows

`func (o *SavepointScheduleConfig) HasBlackoutWindows() bool`

HasBlackoutWindows returns a boolean if a field has been set.

### GetFreshnessThresholdHours

`func (o *SavepointScheduleConfig) GetFreshnessThresholdHours() int32`

GetFreshnessThresholdHours returns the FreshnessThresholdHours field if non-nil, zero value otherwise.

### GetFreshnessThresholdHoursOk

`func (o *SavepointScheduleConfig) GetFreshnessThresholdHoursOk() (*int32, bool)`

GetFreshnessThresholdHoursOk returns a tuple with the FreshnessThresholdHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreshnessThresholdHours

`func (o *SavepointScheduleConfig) SetFreshnessThresholdHours(v int32)`

SetFreshnessThresholdHours sets FreshnessThresholdHours field to given value.

### HasFreshnessThresholdHours

`func (o *SavepointScheduleConfig) HasFreshnessThresholdHours() bool`

HasFreshnessThresholdHours returns a boolean if a field has been set.

### GetFormatType

`func (o *SavepointScheduleConfig) GetFormatType() string`

GetFormatType returns the FormatType field if non-nil, zero value otherwise.

### GetFormatTypeOk

`func (o *SavepointScheduleConfig) GetFormatTypeOk() (*string, bool)`

GetFormatTypeOk returns a tuple with the FormatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatType

`func (o *SavepointScheduleConfig) SetFormatType(v string)`

SetFormatType sets FormatType field to given value.

### HasFormatType

`func (o *SavepointScheduleConfig) HasFormatType() bool`

HasFormatType returns a boolean if a field has been set.

### GetRetention

`func (o *SavepointScheduleConfig) GetRetention() SavepointRetention`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *SavepointScheduleConfig) GetRetentionOk() (*SavepointRetention, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *SavepointScheduleConfig) SetRetention(v SavepointRetention)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *SavepointScheduleConfig) HasRetention() bool`

HasRetention returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


