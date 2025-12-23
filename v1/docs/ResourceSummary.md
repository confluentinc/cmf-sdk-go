# ResourceSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to **float64** | Total CPU units allocated | [optional] 
**Memory** | Pointer to **int64** | Total memory allocated in bytes | [optional] 

## Methods

### NewResourceSummary

`func NewResourceSummary() *ResourceSummary`

NewResourceSummary instantiates a new ResourceSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceSummaryWithDefaults

`func NewResourceSummaryWithDefaults() *ResourceSummary`

NewResourceSummaryWithDefaults instantiates a new ResourceSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *ResourceSummary) GetCpu() float64`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ResourceSummary) GetCpuOk() (*float64, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ResourceSummary) SetCpu(v float64)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *ResourceSummary) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMemory

`func (o *ResourceSummary) GetMemory() int64`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ResourceSummary) GetMemoryOk() (*int64, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ResourceSummary) SetMemory(v int64)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *ResourceSummary) HasMemory() bool`

HasMemory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


