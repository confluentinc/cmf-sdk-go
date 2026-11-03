# C3FeaturesStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PodLogsEnabled** | Pointer to **bool** | Whether reading the JobManager/TaskManager pod logs of an Application is enabled. When false, the pod-log endpoints answer 404 and the SPA shows an explanatory screen instead of an empty log view. Controlled by cmf.logs.enabled. | [optional] 
**RoleBindingEnabled** | Pointer to **bool** | Whether managing RBAC role bindings is enabled in the UI. Controlled by cmf.ui.role-binding.enabled. | [optional] 
**SqlAssistant** | Pointer to [**C3SqlAssistantFeature**](C3SqlAssistantFeature.md) |  | [optional] 

## Methods

### NewC3FeaturesStatus

`func NewC3FeaturesStatus() *C3FeaturesStatus`

NewC3FeaturesStatus instantiates a new C3FeaturesStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC3FeaturesStatusWithDefaults

`func NewC3FeaturesStatusWithDefaults() *C3FeaturesStatus`

NewC3FeaturesStatusWithDefaults instantiates a new C3FeaturesStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPodLogsEnabled

`func (o *C3FeaturesStatus) GetPodLogsEnabled() bool`

GetPodLogsEnabled returns the PodLogsEnabled field if non-nil, zero value otherwise.

### GetPodLogsEnabledOk

`func (o *C3FeaturesStatus) GetPodLogsEnabledOk() (*bool, bool)`

GetPodLogsEnabledOk returns a tuple with the PodLogsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodLogsEnabled

`func (o *C3FeaturesStatus) SetPodLogsEnabled(v bool)`

SetPodLogsEnabled sets PodLogsEnabled field to given value.

### HasPodLogsEnabled

`func (o *C3FeaturesStatus) HasPodLogsEnabled() bool`

HasPodLogsEnabled returns a boolean if a field has been set.

### GetRoleBindingEnabled

`func (o *C3FeaturesStatus) GetRoleBindingEnabled() bool`

GetRoleBindingEnabled returns the RoleBindingEnabled field if non-nil, zero value otherwise.

### GetRoleBindingEnabledOk

`func (o *C3FeaturesStatus) GetRoleBindingEnabledOk() (*bool, bool)`

GetRoleBindingEnabledOk returns a tuple with the RoleBindingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleBindingEnabled

`func (o *C3FeaturesStatus) SetRoleBindingEnabled(v bool)`

SetRoleBindingEnabled sets RoleBindingEnabled field to given value.

### HasRoleBindingEnabled

`func (o *C3FeaturesStatus) HasRoleBindingEnabled() bool`

HasRoleBindingEnabled returns a boolean if a field has been set.

### GetSqlAssistant

`func (o *C3FeaturesStatus) GetSqlAssistant() C3SqlAssistantFeature`

GetSqlAssistant returns the SqlAssistant field if non-nil, zero value otherwise.

### GetSqlAssistantOk

`func (o *C3FeaturesStatus) GetSqlAssistantOk() (*C3SqlAssistantFeature, bool)`

GetSqlAssistantOk returns a tuple with the SqlAssistant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlAssistant

`func (o *C3FeaturesStatus) SetSqlAssistant(v C3SqlAssistantFeature)`

SetSqlAssistant sets SqlAssistant field to given value.

### HasSqlAssistant

`func (o *C3FeaturesStatus) HasSqlAssistant() bool`

HasSqlAssistant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


