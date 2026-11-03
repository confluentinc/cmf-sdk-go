# AutoscalerDelayedScaleDown

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vertices** | Pointer to [**map[string]AutoscalerDelayedVertex**](AutoscalerDelayedVertex.md) | Held scale-downs, as a map of job vertex id to the decision for that vertex | [optional] 

## Methods

### NewAutoscalerDelayedScaleDown

`func NewAutoscalerDelayedScaleDown() *AutoscalerDelayedScaleDown`

NewAutoscalerDelayedScaleDown instantiates a new AutoscalerDelayedScaleDown object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoscalerDelayedScaleDownWithDefaults

`func NewAutoscalerDelayedScaleDownWithDefaults() *AutoscalerDelayedScaleDown`

NewAutoscalerDelayedScaleDownWithDefaults instantiates a new AutoscalerDelayedScaleDown object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVertices

`func (o *AutoscalerDelayedScaleDown) GetVertices() map[string]AutoscalerDelayedVertex`

GetVertices returns the Vertices field if non-nil, zero value otherwise.

### GetVerticesOk

`func (o *AutoscalerDelayedScaleDown) GetVerticesOk() (*map[string]AutoscalerDelayedVertex, bool)`

GetVerticesOk returns a tuple with the Vertices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVertices

`func (o *AutoscalerDelayedScaleDown) SetVertices(v map[string]AutoscalerDelayedVertex)`

SetVertices sets Vertices field to given value.

### HasVertices

`func (o *AutoscalerDelayedScaleDown) HasVertices() bool`

HasVertices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


