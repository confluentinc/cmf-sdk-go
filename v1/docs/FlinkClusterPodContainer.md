# FlinkClusterPodContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the container. Use it with the &#x60;container&#x60; query parameter of the logs endpoint. | 
**Ready** | **bool** | Whether the container currently passes its readiness check. | 
**RestartCount** | **int32** | How often Kubernetes has restarted this container. | 
**State** | Pointer to **string** | What the container is doing right now, one of &#x60;running&#x60;, &#x60;waiting&#x60; or &#x60;terminated&#x60;. Absent while Kubernetes has not reported a state yet. | [optional] 
**StateReason** | Pointer to **string** | The reason Kubernetes gives for the current state, for example &#x60;CrashLoopBackOff&#x60; or &#x60;OOMKilled&#x60;. Absent when Kubernetes reports no reason. | [optional] 
**PreviousInstanceAvailable** | **bool** | Whether a previous, terminated instance of this container exists whose logs can be read with the &#x60;previous&#x60; query parameter of the logs endpoint. | 

## Methods

### NewFlinkClusterPodContainer

`func NewFlinkClusterPodContainer(name string, ready bool, restartCount int32, previousInstanceAvailable bool, ) *FlinkClusterPodContainer`

NewFlinkClusterPodContainer instantiates a new FlinkClusterPodContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlinkClusterPodContainerWithDefaults

`func NewFlinkClusterPodContainerWithDefaults() *FlinkClusterPodContainer`

NewFlinkClusterPodContainerWithDefaults instantiates a new FlinkClusterPodContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FlinkClusterPodContainer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlinkClusterPodContainer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlinkClusterPodContainer) SetName(v string)`

SetName sets Name field to given value.


### GetReady

`func (o *FlinkClusterPodContainer) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *FlinkClusterPodContainer) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *FlinkClusterPodContainer) SetReady(v bool)`

SetReady sets Ready field to given value.


### GetRestartCount

`func (o *FlinkClusterPodContainer) GetRestartCount() int32`

GetRestartCount returns the RestartCount field if non-nil, zero value otherwise.

### GetRestartCountOk

`func (o *FlinkClusterPodContainer) GetRestartCountOk() (*int32, bool)`

GetRestartCountOk returns a tuple with the RestartCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartCount

`func (o *FlinkClusterPodContainer) SetRestartCount(v int32)`

SetRestartCount sets RestartCount field to given value.


### GetState

`func (o *FlinkClusterPodContainer) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FlinkClusterPodContainer) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FlinkClusterPodContainer) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *FlinkClusterPodContainer) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateReason

`func (o *FlinkClusterPodContainer) GetStateReason() string`

GetStateReason returns the StateReason field if non-nil, zero value otherwise.

### GetStateReasonOk

`func (o *FlinkClusterPodContainer) GetStateReasonOk() (*string, bool)`

GetStateReasonOk returns a tuple with the StateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateReason

`func (o *FlinkClusterPodContainer) SetStateReason(v string)`

SetStateReason sets StateReason field to given value.

### HasStateReason

`func (o *FlinkClusterPodContainer) HasStateReason() bool`

HasStateReason returns a boolean if a field has been set.

### GetPreviousInstanceAvailable

`func (o *FlinkClusterPodContainer) GetPreviousInstanceAvailable() bool`

GetPreviousInstanceAvailable returns the PreviousInstanceAvailable field if non-nil, zero value otherwise.

### GetPreviousInstanceAvailableOk

`func (o *FlinkClusterPodContainer) GetPreviousInstanceAvailableOk() (*bool, bool)`

GetPreviousInstanceAvailableOk returns a tuple with the PreviousInstanceAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousInstanceAvailable

`func (o *FlinkClusterPodContainer) SetPreviousInstanceAvailable(v bool)`

SetPreviousInstanceAvailable sets PreviousInstanceAvailable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


