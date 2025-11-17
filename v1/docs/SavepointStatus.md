# SavepointStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | State of the Savepoint | [optional] 
**Path** | Pointer to **string** | Path of the Savepoint | [optional] 
**TriggerTimestamp** | Pointer to **string** | Timestamp when the Savepoint was triggered | [optional] 
**ResultTimestamp** | Pointer to **string** | Timestamp when the Savepoint result was received | [optional] 
**Failures** | Pointer to **int32** | The number of failures of the Savepoint | [optional] 
**Error** | Pointer to **string** | The error message for the Savepoint | [optional] 
**PendingDeletion** | Pointer to **bool** | Whether the Savepoint is pending deletion | [optional] 

## Methods

### NewSavepointStatus

`func NewSavepointStatus() *SavepointStatus`

NewSavepointStatus instantiates a new SavepointStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSavepointStatusWithDefaults

`func NewSavepointStatusWithDefaults() *SavepointStatus`

NewSavepointStatusWithDefaults instantiates a new SavepointStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *SavepointStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SavepointStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SavepointStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *SavepointStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPath

`func (o *SavepointStatus) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SavepointStatus) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SavepointStatus) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SavepointStatus) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetTriggerTimestamp

`func (o *SavepointStatus) GetTriggerTimestamp() string`

GetTriggerTimestamp returns the TriggerTimestamp field if non-nil, zero value otherwise.

### GetTriggerTimestampOk

`func (o *SavepointStatus) GetTriggerTimestampOk() (*string, bool)`

GetTriggerTimestampOk returns a tuple with the TriggerTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerTimestamp

`func (o *SavepointStatus) SetTriggerTimestamp(v string)`

SetTriggerTimestamp sets TriggerTimestamp field to given value.

### HasTriggerTimestamp

`func (o *SavepointStatus) HasTriggerTimestamp() bool`

HasTriggerTimestamp returns a boolean if a field has been set.

### GetResultTimestamp

`func (o *SavepointStatus) GetResultTimestamp() string`

GetResultTimestamp returns the ResultTimestamp field if non-nil, zero value otherwise.

### GetResultTimestampOk

`func (o *SavepointStatus) GetResultTimestampOk() (*string, bool)`

GetResultTimestampOk returns a tuple with the ResultTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultTimestamp

`func (o *SavepointStatus) SetResultTimestamp(v string)`

SetResultTimestamp sets ResultTimestamp field to given value.

### HasResultTimestamp

`func (o *SavepointStatus) HasResultTimestamp() bool`

HasResultTimestamp returns a boolean if a field has been set.

### GetFailures

`func (o *SavepointStatus) GetFailures() int32`

GetFailures returns the Failures field if non-nil, zero value otherwise.

### GetFailuresOk

`func (o *SavepointStatus) GetFailuresOk() (*int32, bool)`

GetFailuresOk returns a tuple with the Failures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailures

`func (o *SavepointStatus) SetFailures(v int32)`

SetFailures sets Failures field to given value.

### HasFailures

`func (o *SavepointStatus) HasFailures() bool`

HasFailures returns a boolean if a field has been set.

### GetError

`func (o *SavepointStatus) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SavepointStatus) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SavepointStatus) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *SavepointStatus) HasError() bool`

HasError returns a boolean if a field has been set.

### GetPendingDeletion

`func (o *SavepointStatus) GetPendingDeletion() bool`

GetPendingDeletion returns the PendingDeletion field if non-nil, zero value otherwise.

### GetPendingDeletionOk

`func (o *SavepointStatus) GetPendingDeletionOk() (*bool, bool)`

GetPendingDeletionOk returns a tuple with the PendingDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingDeletion

`func (o *SavepointStatus) SetPendingDeletion(v bool)`

SetPendingDeletion sets PendingDeletion field to given value.

### HasPendingDeletion

`func (o *SavepointStatus) HasPendingDeletion() bool`

HasPendingDeletion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


