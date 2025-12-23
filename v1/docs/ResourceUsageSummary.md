# ResourceUsageSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobManager** | Pointer to [**JobManager**](JobManager.md) |  | [optional] 
**TaskManager** | Pointer to [**TaskManager**](TaskManager.md) |  | [optional] 

## Methods

### NewResourceUsageSummary

`func NewResourceUsageSummary() *ResourceUsageSummary`

NewResourceUsageSummary instantiates a new ResourceUsageSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceUsageSummaryWithDefaults

`func NewResourceUsageSummaryWithDefaults() *ResourceUsageSummary`

NewResourceUsageSummaryWithDefaults instantiates a new ResourceUsageSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobManager

`func (o *ResourceUsageSummary) GetJobManager() JobManager`

GetJobManager returns the JobManager field if non-nil, zero value otherwise.

### GetJobManagerOk

`func (o *ResourceUsageSummary) GetJobManagerOk() (*JobManager, bool)`

GetJobManagerOk returns a tuple with the JobManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobManager

`func (o *ResourceUsageSummary) SetJobManager(v JobManager)`

SetJobManager sets JobManager field to given value.

### HasJobManager

`func (o *ResourceUsageSummary) HasJobManager() bool`

HasJobManager returns a boolean if a field has been set.

### GetTaskManager

`func (o *ResourceUsageSummary) GetTaskManager() TaskManager`

GetTaskManager returns the TaskManager field if non-nil, zero value otherwise.

### GetTaskManagerOk

`func (o *ResourceUsageSummary) GetTaskManagerOk() (*TaskManager, bool)`

GetTaskManagerOk returns a tuple with the TaskManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskManager

`func (o *ResourceUsageSummary) SetTaskManager(v TaskManager)`

SetTaskManager sets TaskManager field to given value.

### HasTaskManager

`func (o *ResourceUsageSummary) HasTaskManager() bool`

HasTaskManager returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


