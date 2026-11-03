# KubernetesCleanupFailure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | Kind of the deleted resource: &#x60;FlinkApplication&#x60;, &#x60;Statement&#x60;, &#x60;ComputePool&#x60;, &#x60;Savepoint&#x60; or &#x60;Environment&#x60;. | [optional] 
**Name** | Pointer to **string** | Name of the deleted resource. | [optional] 
**Namespace** | Pointer to **NullableString** | Kubernetes namespace of its objects, if known. | [optional] 
**Reason** | Pointer to **string** | Which of its Kubernetes objects might still exist, and why. | [optional] 

## Methods

### NewKubernetesCleanupFailure

`func NewKubernetesCleanupFailure() *KubernetesCleanupFailure`

NewKubernetesCleanupFailure instantiates a new KubernetesCleanupFailure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesCleanupFailureWithDefaults

`func NewKubernetesCleanupFailureWithDefaults() *KubernetesCleanupFailure`

NewKubernetesCleanupFailureWithDefaults instantiates a new KubernetesCleanupFailure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *KubernetesCleanupFailure) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *KubernetesCleanupFailure) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *KubernetesCleanupFailure) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *KubernetesCleanupFailure) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *KubernetesCleanupFailure) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesCleanupFailure) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesCleanupFailure) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubernetesCleanupFailure) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *KubernetesCleanupFailure) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *KubernetesCleanupFailure) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *KubernetesCleanupFailure) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *KubernetesCleanupFailure) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *KubernetesCleanupFailure) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *KubernetesCleanupFailure) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetReason

`func (o *KubernetesCleanupFailure) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *KubernetesCleanupFailure) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *KubernetesCleanupFailure) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *KubernetesCleanupFailure) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


