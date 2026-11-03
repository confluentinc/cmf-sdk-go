# AutoscalerStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether the autoscaler is turned on for this Application | [optional] 
**ScalingEnabled** | Pointer to **bool** | Whether the autoscaler may act on its recommendations. When false, the autoscaler still evaluates metrics and reports a recommended parallelism but never rescales the job.  | [optional] 
**Options** | Pointer to **map[string]string** | The autoscaler options the operator sees for this Application, as a map of option name to value (for example &#x60;job.autoscaler.utilization.target&#x60;). These are the Application&#39;s own options merged with the Application defaults of its Environment, so an option inherited from the Environment appears here too and cannot be told apart from one set on the Application. Options absent from this map fall back to the defaults of the Flink Kubernetes Operator in use.  | [optional] 
**State** | Pointer to **string** | &#x60;EVALUATING&#x60; - a full metric window has been collected and the reported recommendations are current. &#x60;COLLECTING&#x60; - the autoscaler is still filling its metric window, so no recommendation exists yet. &#x60;DISABLED&#x60; - the autoscaler is not turned on for this Application. &#x60;UNAVAILABLE&#x60; - the state could not be read; see &#x60;message&#x60;. The boundary between &#x60;COLLECTING&#x60; and &#x60;EVALUATING&#x60; is &#x60;windowFullAt&#x60; and carries the same approximation.  | [optional] 
**Message** | Pointer to **string** | Human readable explanation, set when the state is not EVALUATING | [optional] 
**ParallelismOverrides** | Pointer to **map[string]string** | The parallelism the autoscaler is currently applying, as a map of job vertex id to parallelism. Empty when the autoscaler has not changed the job&#39;s parallelism.  | [optional] 
**DelayedScaleDown** | Pointer to [**AutoscalerDelayedScaleDown**](AutoscalerDelayedScaleDown.md) |  | [optional] 
**MetricHistory** | Pointer to [**[]AutoscalerMetricSample**](AutoscalerMetricSample.md) | The metric samples currently in the autoscaler&#39;s window, oldest first | [optional] [default to []]
**JobRunningSince** | Pointer to **time.Time** | When the job started, as the operator reports it. The autoscaler&#39;s evaluation cycle is measured from the moment the job entered the RUNNING state, which is not reported here and is later than this by however long the job took to initialize, so treat this as the start of the cycle rather than its exact beginning. It also does not move where a rescale restarts the job in place, though the autoscaler&#39;s cycle does. Absent while the operator has reported no start time for the job.  | [optional] 
**WindowFullAt** | Pointer to **time.Time** | Approximately when the autoscaler first holds, or first held, a full metric window and so begins deciding: &#x60;jobRunningSince&#x60; plus the stabilization interval plus the metric window. Approximate in both directions - the autoscaler measures the window from its first sample after stabilization ends rather than from that instant, which puts the real moment later, while &#x60;jobRunningSince&#x60; is earlier than the state transition the autoscaler measures from, which puts it earlier. Absent while there is nothing to derive it from.  | [optional] 
**ScalingEvents** | Pointer to [**[]AutoscalerScalingEvent**](AutoscalerScalingEvent.md) | The parallelism changes the autoscaler has made, oldest first, flattened across job vertices.  | [optional] [default to []]
**ScalingTracking** | Pointer to **map[string]map[string]interface{}** | Timing of recent rescale operations. The shape of each entry mirrors the autoscaler&#39;s own bookkeeping and may change between releases.  | [optional] 

## Methods

### NewAutoscalerStatus

`func NewAutoscalerStatus() *AutoscalerStatus`

NewAutoscalerStatus instantiates a new AutoscalerStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoscalerStatusWithDefaults

`func NewAutoscalerStatusWithDefaults() *AutoscalerStatus`

NewAutoscalerStatusWithDefaults instantiates a new AutoscalerStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *AutoscalerStatus) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AutoscalerStatus) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AutoscalerStatus) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AutoscalerStatus) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetScalingEnabled

`func (o *AutoscalerStatus) GetScalingEnabled() bool`

GetScalingEnabled returns the ScalingEnabled field if non-nil, zero value otherwise.

### GetScalingEnabledOk

`func (o *AutoscalerStatus) GetScalingEnabledOk() (*bool, bool)`

GetScalingEnabledOk returns a tuple with the ScalingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalingEnabled

`func (o *AutoscalerStatus) SetScalingEnabled(v bool)`

SetScalingEnabled sets ScalingEnabled field to given value.

### HasScalingEnabled

`func (o *AutoscalerStatus) HasScalingEnabled() bool`

HasScalingEnabled returns a boolean if a field has been set.

### GetOptions

`func (o *AutoscalerStatus) GetOptions() map[string]string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *AutoscalerStatus) GetOptionsOk() (*map[string]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *AutoscalerStatus) SetOptions(v map[string]string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *AutoscalerStatus) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetState

`func (o *AutoscalerStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AutoscalerStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AutoscalerStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AutoscalerStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetMessage

`func (o *AutoscalerStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AutoscalerStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AutoscalerStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AutoscalerStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetParallelismOverrides

`func (o *AutoscalerStatus) GetParallelismOverrides() map[string]string`

GetParallelismOverrides returns the ParallelismOverrides field if non-nil, zero value otherwise.

### GetParallelismOverridesOk

`func (o *AutoscalerStatus) GetParallelismOverridesOk() (*map[string]string, bool)`

GetParallelismOverridesOk returns a tuple with the ParallelismOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelismOverrides

`func (o *AutoscalerStatus) SetParallelismOverrides(v map[string]string)`

SetParallelismOverrides sets ParallelismOverrides field to given value.

### HasParallelismOverrides

`func (o *AutoscalerStatus) HasParallelismOverrides() bool`

HasParallelismOverrides returns a boolean if a field has been set.

### GetDelayedScaleDown

`func (o *AutoscalerStatus) GetDelayedScaleDown() AutoscalerDelayedScaleDown`

GetDelayedScaleDown returns the DelayedScaleDown field if non-nil, zero value otherwise.

### GetDelayedScaleDownOk

`func (o *AutoscalerStatus) GetDelayedScaleDownOk() (*AutoscalerDelayedScaleDown, bool)`

GetDelayedScaleDownOk returns a tuple with the DelayedScaleDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayedScaleDown

`func (o *AutoscalerStatus) SetDelayedScaleDown(v AutoscalerDelayedScaleDown)`

SetDelayedScaleDown sets DelayedScaleDown field to given value.

### HasDelayedScaleDown

`func (o *AutoscalerStatus) HasDelayedScaleDown() bool`

HasDelayedScaleDown returns a boolean if a field has been set.

### GetMetricHistory

`func (o *AutoscalerStatus) GetMetricHistory() []AutoscalerMetricSample`

GetMetricHistory returns the MetricHistory field if non-nil, zero value otherwise.

### GetMetricHistoryOk

`func (o *AutoscalerStatus) GetMetricHistoryOk() (*[]AutoscalerMetricSample, bool)`

GetMetricHistoryOk returns a tuple with the MetricHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricHistory

`func (o *AutoscalerStatus) SetMetricHistory(v []AutoscalerMetricSample)`

SetMetricHistory sets MetricHistory field to given value.

### HasMetricHistory

`func (o *AutoscalerStatus) HasMetricHistory() bool`

HasMetricHistory returns a boolean if a field has been set.

### GetJobRunningSince

`func (o *AutoscalerStatus) GetJobRunningSince() time.Time`

GetJobRunningSince returns the JobRunningSince field if non-nil, zero value otherwise.

### GetJobRunningSinceOk

`func (o *AutoscalerStatus) GetJobRunningSinceOk() (*time.Time, bool)`

GetJobRunningSinceOk returns a tuple with the JobRunningSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobRunningSince

`func (o *AutoscalerStatus) SetJobRunningSince(v time.Time)`

SetJobRunningSince sets JobRunningSince field to given value.

### HasJobRunningSince

`func (o *AutoscalerStatus) HasJobRunningSince() bool`

HasJobRunningSince returns a boolean if a field has been set.

### GetWindowFullAt

`func (o *AutoscalerStatus) GetWindowFullAt() time.Time`

GetWindowFullAt returns the WindowFullAt field if non-nil, zero value otherwise.

### GetWindowFullAtOk

`func (o *AutoscalerStatus) GetWindowFullAtOk() (*time.Time, bool)`

GetWindowFullAtOk returns a tuple with the WindowFullAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowFullAt

`func (o *AutoscalerStatus) SetWindowFullAt(v time.Time)`

SetWindowFullAt sets WindowFullAt field to given value.

### HasWindowFullAt

`func (o *AutoscalerStatus) HasWindowFullAt() bool`

HasWindowFullAt returns a boolean if a field has been set.

### GetScalingEvents

`func (o *AutoscalerStatus) GetScalingEvents() []AutoscalerScalingEvent`

GetScalingEvents returns the ScalingEvents field if non-nil, zero value otherwise.

### GetScalingEventsOk

`func (o *AutoscalerStatus) GetScalingEventsOk() (*[]AutoscalerScalingEvent, bool)`

GetScalingEventsOk returns a tuple with the ScalingEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalingEvents

`func (o *AutoscalerStatus) SetScalingEvents(v []AutoscalerScalingEvent)`

SetScalingEvents sets ScalingEvents field to given value.

### HasScalingEvents

`func (o *AutoscalerStatus) HasScalingEvents() bool`

HasScalingEvents returns a boolean if a field has been set.

### GetScalingTracking

`func (o *AutoscalerStatus) GetScalingTracking() map[string]map[string]interface{}`

GetScalingTracking returns the ScalingTracking field if non-nil, zero value otherwise.

### GetScalingTrackingOk

`func (o *AutoscalerStatus) GetScalingTrackingOk() (*map[string]map[string]interface{}, bool)`

GetScalingTrackingOk returns a tuple with the ScalingTracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalingTracking

`func (o *AutoscalerStatus) SetScalingTracking(v map[string]map[string]interface{})`

SetScalingTracking sets ScalingTracking field to given value.

### HasScalingTracking

`func (o *AutoscalerStatus) HasScalingTracking() bool`

HasScalingTracking returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


