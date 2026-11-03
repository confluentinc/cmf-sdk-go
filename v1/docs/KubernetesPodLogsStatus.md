# KubernetesPodLogsStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Container** | **string** | Name of the container within that pod the returned logs were read from. | 
**Logs** | **string** | The log lines, newline separated, oldest first. Empty when the container has not written any output yet. | 
**AvailablePods** | **[]string** | Names of all pods of the requested component that are currently known, including the one the logs were read from. Use these with the &#x60;podName&#x60; query parameter to read the logs of another pod. | [default to []]
**Truncated** | **bool** | Whether the log output was cut short because it exceeded the size a single response may carry. Only &#x60;fromStart&#x60; requests can hit this cap in practice. | [default to false]

## Methods

### NewKubernetesPodLogsStatus

`func NewKubernetesPodLogsStatus(container string, logs string, availablePods []string, truncated bool, ) *KubernetesPodLogsStatus`

NewKubernetesPodLogsStatus instantiates a new KubernetesPodLogsStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesPodLogsStatusWithDefaults

`func NewKubernetesPodLogsStatusWithDefaults() *KubernetesPodLogsStatus`

NewKubernetesPodLogsStatusWithDefaults instantiates a new KubernetesPodLogsStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainer

`func (o *KubernetesPodLogsStatus) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *KubernetesPodLogsStatus) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *KubernetesPodLogsStatus) SetContainer(v string)`

SetContainer sets Container field to given value.


### GetLogs

`func (o *KubernetesPodLogsStatus) GetLogs() string`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *KubernetesPodLogsStatus) GetLogsOk() (*string, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *KubernetesPodLogsStatus) SetLogs(v string)`

SetLogs sets Logs field to given value.


### GetAvailablePods

`func (o *KubernetesPodLogsStatus) GetAvailablePods() []string`

GetAvailablePods returns the AvailablePods field if non-nil, zero value otherwise.

### GetAvailablePodsOk

`func (o *KubernetesPodLogsStatus) GetAvailablePodsOk() (*[]string, bool)`

GetAvailablePodsOk returns a tuple with the AvailablePods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailablePods

`func (o *KubernetesPodLogsStatus) SetAvailablePods(v []string)`

SetAvailablePods sets AvailablePods field to given value.


### GetTruncated

`func (o *KubernetesPodLogsStatus) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *KubernetesPodLogsStatus) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *KubernetesPodLogsStatus) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


