# AutoscalerDelayedVertex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstTriggerTime** | Pointer to **time.Time** | When the autoscaler first wanted to scale this vertex down. The scale-down becomes eligible one scale-down interval after this time.  | [optional] 
**Recommendations** | Pointer to [**[]AutoscalerParallelismRecommendation**](AutoscalerParallelismRecommendation.md) | The parallelism values recommended for this vertex while the scale-down is held | [optional] [default to []]

## Methods

### NewAutoscalerDelayedVertex

`func NewAutoscalerDelayedVertex() *AutoscalerDelayedVertex`

NewAutoscalerDelayedVertex instantiates a new AutoscalerDelayedVertex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoscalerDelayedVertexWithDefaults

`func NewAutoscalerDelayedVertexWithDefaults() *AutoscalerDelayedVertex`

NewAutoscalerDelayedVertexWithDefaults instantiates a new AutoscalerDelayedVertex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstTriggerTime

`func (o *AutoscalerDelayedVertex) GetFirstTriggerTime() time.Time`

GetFirstTriggerTime returns the FirstTriggerTime field if non-nil, zero value otherwise.

### GetFirstTriggerTimeOk

`func (o *AutoscalerDelayedVertex) GetFirstTriggerTimeOk() (*time.Time, bool)`

GetFirstTriggerTimeOk returns a tuple with the FirstTriggerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstTriggerTime

`func (o *AutoscalerDelayedVertex) SetFirstTriggerTime(v time.Time)`

SetFirstTriggerTime sets FirstTriggerTime field to given value.

### HasFirstTriggerTime

`func (o *AutoscalerDelayedVertex) HasFirstTriggerTime() bool`

HasFirstTriggerTime returns a boolean if a field has been set.

### GetRecommendations

`func (o *AutoscalerDelayedVertex) GetRecommendations() []AutoscalerParallelismRecommendation`

GetRecommendations returns the Recommendations field if non-nil, zero value otherwise.

### GetRecommendationsOk

`func (o *AutoscalerDelayedVertex) GetRecommendationsOk() (*[]AutoscalerParallelismRecommendation, bool)`

GetRecommendationsOk returns a tuple with the Recommendations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendations

`func (o *AutoscalerDelayedVertex) SetRecommendations(v []AutoscalerParallelismRecommendation)`

SetRecommendations sets Recommendations field to given value.

### HasRecommendations

`func (o *AutoscalerDelayedVertex) HasRecommendations() bool`

HasRecommendations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


