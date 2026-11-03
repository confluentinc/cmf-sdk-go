# EnvironmentDeletionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ToBeDeleted** | Pointer to [**EnvironmentResourceNames**](EnvironmentResourceNames.md) |  | [optional] 
**Deleted** | Pointer to [**EnvironmentResourceNames**](EnvironmentResourceNames.md) |  | [optional] 
**NotCleanedUp** | Pointer to [**[]KubernetesCleanupFailure**](KubernetesCleanupFailure.md) | Kubernetes objects of the deleted resources that might still exist, because they couldn&#39;t be deleted or were already being deleted. | [optional] 

## Methods

### NewEnvironmentDeletionResult

`func NewEnvironmentDeletionResult() *EnvironmentDeletionResult`

NewEnvironmentDeletionResult instantiates a new EnvironmentDeletionResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentDeletionResultWithDefaults

`func NewEnvironmentDeletionResultWithDefaults() *EnvironmentDeletionResult`

NewEnvironmentDeletionResultWithDefaults instantiates a new EnvironmentDeletionResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToBeDeleted

`func (o *EnvironmentDeletionResult) GetToBeDeleted() EnvironmentResourceNames`

GetToBeDeleted returns the ToBeDeleted field if non-nil, zero value otherwise.

### GetToBeDeletedOk

`func (o *EnvironmentDeletionResult) GetToBeDeletedOk() (*EnvironmentResourceNames, bool)`

GetToBeDeletedOk returns a tuple with the ToBeDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToBeDeleted

`func (o *EnvironmentDeletionResult) SetToBeDeleted(v EnvironmentResourceNames)`

SetToBeDeleted sets ToBeDeleted field to given value.

### HasToBeDeleted

`func (o *EnvironmentDeletionResult) HasToBeDeleted() bool`

HasToBeDeleted returns a boolean if a field has been set.

### GetDeleted

`func (o *EnvironmentDeletionResult) GetDeleted() EnvironmentResourceNames`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *EnvironmentDeletionResult) GetDeletedOk() (*EnvironmentResourceNames, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *EnvironmentDeletionResult) SetDeleted(v EnvironmentResourceNames)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *EnvironmentDeletionResult) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetNotCleanedUp

`func (o *EnvironmentDeletionResult) GetNotCleanedUp() []KubernetesCleanupFailure`

GetNotCleanedUp returns the NotCleanedUp field if non-nil, zero value otherwise.

### GetNotCleanedUpOk

`func (o *EnvironmentDeletionResult) GetNotCleanedUpOk() (*[]KubernetesCleanupFailure, bool)`

GetNotCleanedUpOk returns a tuple with the NotCleanedUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotCleanedUp

`func (o *EnvironmentDeletionResult) SetNotCleanedUp(v []KubernetesCleanupFailure)`

SetNotCleanedUp sets NotCleanedUp field to given value.

### HasNotCleanedUp

`func (o *EnvironmentDeletionResult) HasNotCleanedUp() bool`

HasNotCleanedUp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


