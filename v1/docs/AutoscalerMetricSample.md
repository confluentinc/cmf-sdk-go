# AutoscalerMetricSample

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **time.Time** | When the sample was collected | [optional] 
**Vertices** | Pointer to **map[string]map[string]float64** | Per job vertex metrics, as a map of job vertex id to a map of metric name to value (for example &#x60;LOAD&#x60; or &#x60;TRUE_PROCESSING_RATE&#x60;). A value the autoscaler reports as infinite or undefined is returned as null, because those are not representable in JSON.  | [optional] 
**Global** | Pointer to **map[string]float64** | Job-wide metrics such as heap usage and garbage collection pressure, as a map of metric name to value. Infinite or undefined values are returned as null.  | [optional] 

## Methods

### NewAutoscalerMetricSample

`func NewAutoscalerMetricSample() *AutoscalerMetricSample`

NewAutoscalerMetricSample instantiates a new AutoscalerMetricSample object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoscalerMetricSampleWithDefaults

`func NewAutoscalerMetricSampleWithDefaults() *AutoscalerMetricSample`

NewAutoscalerMetricSampleWithDefaults instantiates a new AutoscalerMetricSample object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *AutoscalerMetricSample) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AutoscalerMetricSample) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AutoscalerMetricSample) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AutoscalerMetricSample) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetVertices

`func (o *AutoscalerMetricSample) GetVertices() map[string]map[string]float64`

GetVertices returns the Vertices field if non-nil, zero value otherwise.

### GetVerticesOk

`func (o *AutoscalerMetricSample) GetVerticesOk() (*map[string]map[string]float64, bool)`

GetVerticesOk returns a tuple with the Vertices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVertices

`func (o *AutoscalerMetricSample) SetVertices(v map[string]map[string]float64)`

SetVertices sets Vertices field to given value.

### HasVertices

`func (o *AutoscalerMetricSample) HasVertices() bool`

HasVertices returns a boolean if a field has been set.

### GetGlobal

`func (o *AutoscalerMetricSample) GetGlobal() map[string]float64`

GetGlobal returns the Global field if non-nil, zero value otherwise.

### GetGlobalOk

`func (o *AutoscalerMetricSample) GetGlobalOk() (*map[string]float64, bool)`

GetGlobalOk returns a tuple with the Global field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobal

`func (o *AutoscalerMetricSample) SetGlobal(v map[string]float64)`

SetGlobal sets Global field to given value.

### HasGlobal

`func (o *AutoscalerMetricSample) HasGlobal() bool`

HasGlobal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


