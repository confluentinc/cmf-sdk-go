# AutoscalerParallelismRecommendation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TriggerTime** | Pointer to **time.Time** | When this parallelism was recommended | [optional] 
**Parallelism** | Pointer to **int32** | The recommended parallelism | [optional] 
**OutsideUtilizationBound** | Pointer to **bool** | Whether the measured utilization was outside the configured target band when this recommendation was made.  | [optional] 

## Methods

### NewAutoscalerParallelismRecommendation

`func NewAutoscalerParallelismRecommendation() *AutoscalerParallelismRecommendation`

NewAutoscalerParallelismRecommendation instantiates a new AutoscalerParallelismRecommendation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoscalerParallelismRecommendationWithDefaults

`func NewAutoscalerParallelismRecommendationWithDefaults() *AutoscalerParallelismRecommendation`

NewAutoscalerParallelismRecommendationWithDefaults instantiates a new AutoscalerParallelismRecommendation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTriggerTime

`func (o *AutoscalerParallelismRecommendation) GetTriggerTime() time.Time`

GetTriggerTime returns the TriggerTime field if non-nil, zero value otherwise.

### GetTriggerTimeOk

`func (o *AutoscalerParallelismRecommendation) GetTriggerTimeOk() (*time.Time, bool)`

GetTriggerTimeOk returns a tuple with the TriggerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerTime

`func (o *AutoscalerParallelismRecommendation) SetTriggerTime(v time.Time)`

SetTriggerTime sets TriggerTime field to given value.

### HasTriggerTime

`func (o *AutoscalerParallelismRecommendation) HasTriggerTime() bool`

HasTriggerTime returns a boolean if a field has been set.

### GetParallelism

`func (o *AutoscalerParallelismRecommendation) GetParallelism() int32`

GetParallelism returns the Parallelism field if non-nil, zero value otherwise.

### GetParallelismOk

`func (o *AutoscalerParallelismRecommendation) GetParallelismOk() (*int32, bool)`

GetParallelismOk returns a tuple with the Parallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelism

`func (o *AutoscalerParallelismRecommendation) SetParallelism(v int32)`

SetParallelism sets Parallelism field to given value.

### HasParallelism

`func (o *AutoscalerParallelismRecommendation) HasParallelism() bool`

HasParallelism returns a boolean if a field has been set.

### GetOutsideUtilizationBound

`func (o *AutoscalerParallelismRecommendation) GetOutsideUtilizationBound() bool`

GetOutsideUtilizationBound returns the OutsideUtilizationBound field if non-nil, zero value otherwise.

### GetOutsideUtilizationBoundOk

`func (o *AutoscalerParallelismRecommendation) GetOutsideUtilizationBoundOk() (*bool, bool)`

GetOutsideUtilizationBoundOk returns a tuple with the OutsideUtilizationBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutsideUtilizationBound

`func (o *AutoscalerParallelismRecommendation) SetOutsideUtilizationBound(v bool)`

SetOutsideUtilizationBound sets OutsideUtilizationBound field to given value.

### HasOutsideUtilizationBound

`func (o *AutoscalerParallelismRecommendation) HasOutsideUtilizationBound() bool`

HasOutsideUtilizationBound returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


