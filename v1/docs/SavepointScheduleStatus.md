# SavepointScheduleStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | Lifecycle state of the schedule. | [optional] [readonly] 
**LastTriggerTime** | Pointer to **time.Time** | Time at which the trigger task last attempted to create a savepoint, i.e. the last fire that passed the blackout, resource-ready and in-flight guards. Skipped fires do not update this field, and it reflects the attempt regardless of whether the savepoint request succeeded. Use lastCompletedTime to know when a savepoint last completed.  | [optional] [readonly] 
**NextTriggerTime** | Pointer to **time.Time** | Next time at which the schedule will be evaluated for a trigger. | [optional] [readonly] 
**LastCompletedTime** | Pointer to **time.Time** | Time at which the most recent savepoint for this resource completed successfully, across all sources (scheduled, manual, or adopted), or null if none has completed yet. This is the freshness signal the schedule&#39;s SLA check is measured against.  | [optional] [readonly] 
**LastErrorMessage** | Pointer to **string** | Last failure observed by the trigger task, or null if healthy. | [optional] [readonly] 

## Methods

### NewSavepointScheduleStatus

`func NewSavepointScheduleStatus() *SavepointScheduleStatus`

NewSavepointScheduleStatus instantiates a new SavepointScheduleStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSavepointScheduleStatusWithDefaults

`func NewSavepointScheduleStatusWithDefaults() *SavepointScheduleStatus`

NewSavepointScheduleStatusWithDefaults instantiates a new SavepointScheduleStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *SavepointScheduleStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SavepointScheduleStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SavepointScheduleStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *SavepointScheduleStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetLastTriggerTime

`func (o *SavepointScheduleStatus) GetLastTriggerTime() time.Time`

GetLastTriggerTime returns the LastTriggerTime field if non-nil, zero value otherwise.

### GetLastTriggerTimeOk

`func (o *SavepointScheduleStatus) GetLastTriggerTimeOk() (*time.Time, bool)`

GetLastTriggerTimeOk returns a tuple with the LastTriggerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTriggerTime

`func (o *SavepointScheduleStatus) SetLastTriggerTime(v time.Time)`

SetLastTriggerTime sets LastTriggerTime field to given value.

### HasLastTriggerTime

`func (o *SavepointScheduleStatus) HasLastTriggerTime() bool`

HasLastTriggerTime returns a boolean if a field has been set.

### GetNextTriggerTime

`func (o *SavepointScheduleStatus) GetNextTriggerTime() time.Time`

GetNextTriggerTime returns the NextTriggerTime field if non-nil, zero value otherwise.

### GetNextTriggerTimeOk

`func (o *SavepointScheduleStatus) GetNextTriggerTimeOk() (*time.Time, bool)`

GetNextTriggerTimeOk returns a tuple with the NextTriggerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextTriggerTime

`func (o *SavepointScheduleStatus) SetNextTriggerTime(v time.Time)`

SetNextTriggerTime sets NextTriggerTime field to given value.

### HasNextTriggerTime

`func (o *SavepointScheduleStatus) HasNextTriggerTime() bool`

HasNextTriggerTime returns a boolean if a field has been set.

### GetLastCompletedTime

`func (o *SavepointScheduleStatus) GetLastCompletedTime() time.Time`

GetLastCompletedTime returns the LastCompletedTime field if non-nil, zero value otherwise.

### GetLastCompletedTimeOk

`func (o *SavepointScheduleStatus) GetLastCompletedTimeOk() (*time.Time, bool)`

GetLastCompletedTimeOk returns a tuple with the LastCompletedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCompletedTime

`func (o *SavepointScheduleStatus) SetLastCompletedTime(v time.Time)`

SetLastCompletedTime sets LastCompletedTime field to given value.

### HasLastCompletedTime

`func (o *SavepointScheduleStatus) HasLastCompletedTime() bool`

HasLastCompletedTime returns a boolean if a field has been set.

### GetLastErrorMessage

`func (o *SavepointScheduleStatus) GetLastErrorMessage() string`

GetLastErrorMessage returns the LastErrorMessage field if non-nil, zero value otherwise.

### GetLastErrorMessageOk

`func (o *SavepointScheduleStatus) GetLastErrorMessageOk() (*string, bool)`

GetLastErrorMessageOk returns a tuple with the LastErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorMessage

`func (o *SavepointScheduleStatus) SetLastErrorMessage(v string)`

SetLastErrorMessage sets LastErrorMessage field to given value.

### HasLastErrorMessage

`func (o *SavepointScheduleStatus) HasLastErrorMessage() bool`

HasLastErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


