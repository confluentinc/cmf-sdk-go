# AutoscalerScalingEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **time.Time** | When the autoscaler made the change | [optional] 
**VertexId** | Pointer to **string** | The job vertex whose parallelism changed | [optional] 
**FromParallelism** | Pointer to **int32** | The parallelism the vertex ran at before the change | [optional] 
**ToParallelism** | Pointer to **int32** | The parallelism the autoscaler changed the vertex to | [optional] 
**EvaluatedMetrics** | Pointer to **map[string]float64** | The metrics the autoscaler evaluated when it took this decision, as a map of metric name to value (for example &#x60;TRUE_PROCESSING_RATE&#x60; or &#x60;SCALE_UP_RATE_THRESHOLD&#x60;). These are a snapshot taken at the moment of the decision, not live values, and are recorded only when a decision is taken. A value the autoscaler reports as infinite or undefined is returned as null, because those are not representable in JSON.  | [optional] 

## Methods

### NewAutoscalerScalingEvent

`func NewAutoscalerScalingEvent() *AutoscalerScalingEvent`

NewAutoscalerScalingEvent instantiates a new AutoscalerScalingEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoscalerScalingEventWithDefaults

`func NewAutoscalerScalingEventWithDefaults() *AutoscalerScalingEvent`

NewAutoscalerScalingEventWithDefaults instantiates a new AutoscalerScalingEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *AutoscalerScalingEvent) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AutoscalerScalingEvent) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AutoscalerScalingEvent) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AutoscalerScalingEvent) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetVertexId

`func (o *AutoscalerScalingEvent) GetVertexId() string`

GetVertexId returns the VertexId field if non-nil, zero value otherwise.

### GetVertexIdOk

`func (o *AutoscalerScalingEvent) GetVertexIdOk() (*string, bool)`

GetVertexIdOk returns a tuple with the VertexId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVertexId

`func (o *AutoscalerScalingEvent) SetVertexId(v string)`

SetVertexId sets VertexId field to given value.

### HasVertexId

`func (o *AutoscalerScalingEvent) HasVertexId() bool`

HasVertexId returns a boolean if a field has been set.

### GetFromParallelism

`func (o *AutoscalerScalingEvent) GetFromParallelism() int32`

GetFromParallelism returns the FromParallelism field if non-nil, zero value otherwise.

### GetFromParallelismOk

`func (o *AutoscalerScalingEvent) GetFromParallelismOk() (*int32, bool)`

GetFromParallelismOk returns a tuple with the FromParallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromParallelism

`func (o *AutoscalerScalingEvent) SetFromParallelism(v int32)`

SetFromParallelism sets FromParallelism field to given value.

### HasFromParallelism

`func (o *AutoscalerScalingEvent) HasFromParallelism() bool`

HasFromParallelism returns a boolean if a field has been set.

### GetToParallelism

`func (o *AutoscalerScalingEvent) GetToParallelism() int32`

GetToParallelism returns the ToParallelism field if non-nil, zero value otherwise.

### GetToParallelismOk

`func (o *AutoscalerScalingEvent) GetToParallelismOk() (*int32, bool)`

GetToParallelismOk returns a tuple with the ToParallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToParallelism

`func (o *AutoscalerScalingEvent) SetToParallelism(v int32)`

SetToParallelism sets ToParallelism field to given value.

### HasToParallelism

`func (o *AutoscalerScalingEvent) HasToParallelism() bool`

HasToParallelism returns a boolean if a field has been set.

### GetEvaluatedMetrics

`func (o *AutoscalerScalingEvent) GetEvaluatedMetrics() map[string]float64`

GetEvaluatedMetrics returns the EvaluatedMetrics field if non-nil, zero value otherwise.

### GetEvaluatedMetricsOk

`func (o *AutoscalerScalingEvent) GetEvaluatedMetricsOk() (*map[string]float64, bool)`

GetEvaluatedMetricsOk returns a tuple with the EvaluatedMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluatedMetrics

`func (o *AutoscalerScalingEvent) SetEvaluatedMetrics(v map[string]float64)`

SetEvaluatedMetrics sets EvaluatedMetrics field to given value.

### HasEvaluatedMetrics

`func (o *AutoscalerScalingEvent) HasEvaluatedMetrics() bool`

HasEvaluatedMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


