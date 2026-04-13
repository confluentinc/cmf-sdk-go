# KubernetesClusterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LifecycleState** | Pointer to **string** | The desired lifecycle state of the cluster. | [optional] 

## Methods

### NewKubernetesClusterSpec

`func NewKubernetesClusterSpec() *KubernetesClusterSpec`

NewKubernetesClusterSpec instantiates a new KubernetesClusterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesClusterSpecWithDefaults

`func NewKubernetesClusterSpecWithDefaults() *KubernetesClusterSpec`

NewKubernetesClusterSpecWithDefaults instantiates a new KubernetesClusterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLifecycleState

`func (o *KubernetesClusterSpec) GetLifecycleState() string`

GetLifecycleState returns the LifecycleState field if non-nil, zero value otherwise.

### GetLifecycleStateOk

`func (o *KubernetesClusterSpec) GetLifecycleStateOk() (*string, bool)`

GetLifecycleStateOk returns a tuple with the LifecycleState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleState

`func (o *KubernetesClusterSpec) SetLifecycleState(v string)`

SetLifecycleState sets LifecycleState field to given value.

### HasLifecycleState

`func (o *KubernetesClusterSpec) HasLifecycleState() bool`

HasLifecycleState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


