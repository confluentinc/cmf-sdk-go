# TaskManager

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to **float64** | Total CPU units allocated | [optional] 
**Memory** | Pointer to **int64** | Total memory allocated in bytes | [optional] 
**Count** | Pointer to **int32** | Number of instances | [optional] 

## Methods

### NewTaskManager

`func NewTaskManager() *TaskManager`

NewTaskManager instantiates a new TaskManager object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskManagerWithDefaults

`func NewTaskManagerWithDefaults() *TaskManager`

NewTaskManagerWithDefaults instantiates a new TaskManager object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *TaskManager) GetCpu() float64`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *TaskManager) GetCpuOk() (*float64, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *TaskManager) SetCpu(v float64)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *TaskManager) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMemory

`func (o *TaskManager) GetMemory() int64`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *TaskManager) GetMemoryOk() (*int64, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *TaskManager) SetMemory(v int64)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *TaskManager) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetCount

`func (o *TaskManager) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TaskManager) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TaskManager) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *TaskManager) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


