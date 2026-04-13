# KubernetesClusterStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | The effective state of the cluster. CONNECTED when the cluster is reachable, DISCONNECTED when unreachable, or DECOMMISSIONED when intentionally taken offline. | [optional] 
**Message** | Pointer to **string** | A human-readable message providing details about the current state (e.g., connection error details). | [optional] 
**LastHeartbeatTimestamp** | Pointer to **time.Time** | Timestamp of the last successful heartbeat check. | [optional] 
**KubernetesVersion** | Pointer to **string** | The Kubernetes server version reported by the cluster (e.g., \&quot;v1.28.3\&quot;). Only available when the cluster is CONNECTED. | [optional] 

## Methods

### NewKubernetesClusterStatus

`func NewKubernetesClusterStatus() *KubernetesClusterStatus`

NewKubernetesClusterStatus instantiates a new KubernetesClusterStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesClusterStatusWithDefaults

`func NewKubernetesClusterStatusWithDefaults() *KubernetesClusterStatus`

NewKubernetesClusterStatusWithDefaults instantiates a new KubernetesClusterStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *KubernetesClusterStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *KubernetesClusterStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *KubernetesClusterStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *KubernetesClusterStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetMessage

`func (o *KubernetesClusterStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *KubernetesClusterStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *KubernetesClusterStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *KubernetesClusterStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetLastHeartbeatTimestamp

`func (o *KubernetesClusterStatus) GetLastHeartbeatTimestamp() time.Time`

GetLastHeartbeatTimestamp returns the LastHeartbeatTimestamp field if non-nil, zero value otherwise.

### GetLastHeartbeatTimestampOk

`func (o *KubernetesClusterStatus) GetLastHeartbeatTimestampOk() (*time.Time, bool)`

GetLastHeartbeatTimestampOk returns a tuple with the LastHeartbeatTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHeartbeatTimestamp

`func (o *KubernetesClusterStatus) SetLastHeartbeatTimestamp(v time.Time)`

SetLastHeartbeatTimestamp sets LastHeartbeatTimestamp field to given value.

### HasLastHeartbeatTimestamp

`func (o *KubernetesClusterStatus) HasLastHeartbeatTimestamp() bool`

HasLastHeartbeatTimestamp returns a boolean if a field has been set.

### GetKubernetesVersion

`func (o *KubernetesClusterStatus) GetKubernetesVersion() string`

GetKubernetesVersion returns the KubernetesVersion field if non-nil, zero value otherwise.

### GetKubernetesVersionOk

`func (o *KubernetesClusterStatus) GetKubernetesVersionOk() (*string, bool)`

GetKubernetesVersionOk returns a tuple with the KubernetesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesVersion

`func (o *KubernetesClusterStatus) SetKubernetesVersion(v string)`

SetKubernetesVersion sets KubernetesVersion field to given value.

### HasKubernetesVersion

`func (o *KubernetesClusterStatus) HasKubernetesVersion() bool`

HasKubernetesVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


