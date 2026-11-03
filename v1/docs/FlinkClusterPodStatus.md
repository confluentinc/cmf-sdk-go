# FlinkClusterPodStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Component** | **string** | Which component of the Flink cluster this pod runs. | 
**Phase** | **string** | The pod&#39;s lifecycle phase as reported by Kubernetes, for example &#x60;Pending&#x60;, &#x60;Running&#x60;, &#x60;Succeeded&#x60;, &#x60;Failed&#x60; or &#x60;Unknown&#x60;. | 
**Ready** | **bool** | Whether every container of the pod currently reports itself as ready. | 
**NodeName** | Pointer to **string** | Name of the Kubernetes node the pod is scheduled on, when it has been scheduled. | [optional] 
**Containers** | [**[]FlinkClusterPodContainer**](FlinkClusterPodContainer.md) |  | [default to []]

## Methods

### NewFlinkClusterPodStatus

`func NewFlinkClusterPodStatus(component string, phase string, ready bool, containers []FlinkClusterPodContainer, ) *FlinkClusterPodStatus`

NewFlinkClusterPodStatus instantiates a new FlinkClusterPodStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlinkClusterPodStatusWithDefaults

`func NewFlinkClusterPodStatusWithDefaults() *FlinkClusterPodStatus`

NewFlinkClusterPodStatusWithDefaults instantiates a new FlinkClusterPodStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponent

`func (o *FlinkClusterPodStatus) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *FlinkClusterPodStatus) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *FlinkClusterPodStatus) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetPhase

`func (o *FlinkClusterPodStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *FlinkClusterPodStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *FlinkClusterPodStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.


### GetReady

`func (o *FlinkClusterPodStatus) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *FlinkClusterPodStatus) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *FlinkClusterPodStatus) SetReady(v bool)`

SetReady sets Ready field to given value.


### GetNodeName

`func (o *FlinkClusterPodStatus) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *FlinkClusterPodStatus) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *FlinkClusterPodStatus) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.

### HasNodeName

`func (o *FlinkClusterPodStatus) HasNodeName() bool`

HasNodeName returns a boolean if a field has been set.

### GetContainers

`func (o *FlinkClusterPodStatus) GetContainers() []FlinkClusterPodContainer`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *FlinkClusterPodStatus) GetContainersOk() (*[]FlinkClusterPodContainer, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *FlinkClusterPodStatus) SetContainers(v []FlinkClusterPodContainer)`

SetContainers sets Containers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


