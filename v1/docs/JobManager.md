# JobManager

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to **float64** | Total CPU units allocated | [optional] 
**Memory** | Pointer to **int64** | Total memory allocated in bytes | [optional] 
**Count** | Pointer to **int32** | Number of instances | [optional] 

## Methods

### NewJobManager

`func NewJobManager() *JobManager`

NewJobManager instantiates a new JobManager object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobManagerWithDefaults

`func NewJobManagerWithDefaults() *JobManager`

NewJobManagerWithDefaults instantiates a new JobManager object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *JobManager) GetCpu() float64`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *JobManager) GetCpuOk() (*float64, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *JobManager) SetCpu(v float64)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *JobManager) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMemory

`func (o *JobManager) GetMemory() int64`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *JobManager) GetMemoryOk() (*int64, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *JobManager) SetMemory(v int64)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *JobManager) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetCount

`func (o *JobManager) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *JobManager) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *JobManager) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *JobManager) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


