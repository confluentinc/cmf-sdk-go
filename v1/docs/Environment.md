# Environment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Secrets** | Pointer to **map[string]string** | The secrets mapping for the environment. This is a mapping between connection_secret_id and the secret name. | [optional] [default to {}]
**Name** | **string** | A unique name for the resource. | 
**CreatedTime** | Pointer to **time.Time** |  | [optional] 
**UpdatedTime** | Pointer to **time.Time** |  | [optional] 
**FlinkApplicationDefaults** | Pointer to **map[string]interface{}** |  | [optional] 
**KubernetesNamespace** | **string** |  | 
**ComputePoolDefaults** | Pointer to **map[string]interface{}** | the defaults as YAML or JSON for ComputePools | [optional] 
**StatementDefaults** | Pointer to [**AllStatementDefaults1**](AllStatementDefaults1.md) |  | [optional] 

## Methods

### NewEnvironment

`func NewEnvironment(name string, kubernetesNamespace string, ) *Environment`

NewEnvironment instantiates a new Environment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentWithDefaults

`func NewEnvironmentWithDefaults() *Environment`

NewEnvironmentWithDefaults instantiates a new Environment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecrets

`func (o *Environment) GetSecrets() map[string]string`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *Environment) GetSecretsOk() (*map[string]string, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *Environment) SetSecrets(v map[string]string)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *Environment) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.

### GetName

`func (o *Environment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Environment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Environment) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedTime

`func (o *Environment) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Environment) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Environment) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Environment) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *Environment) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *Environment) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *Environment) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *Environment) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.

### GetFlinkApplicationDefaults

`func (o *Environment) GetFlinkApplicationDefaults() map[string]interface{}`

GetFlinkApplicationDefaults returns the FlinkApplicationDefaults field if non-nil, zero value otherwise.

### GetFlinkApplicationDefaultsOk

`func (o *Environment) GetFlinkApplicationDefaultsOk() (*map[string]interface{}, bool)`

GetFlinkApplicationDefaultsOk returns a tuple with the FlinkApplicationDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlinkApplicationDefaults

`func (o *Environment) SetFlinkApplicationDefaults(v map[string]interface{})`

SetFlinkApplicationDefaults sets FlinkApplicationDefaults field to given value.

### HasFlinkApplicationDefaults

`func (o *Environment) HasFlinkApplicationDefaults() bool`

HasFlinkApplicationDefaults returns a boolean if a field has been set.

### GetKubernetesNamespace

`func (o *Environment) GetKubernetesNamespace() string`

GetKubernetesNamespace returns the KubernetesNamespace field if non-nil, zero value otherwise.

### GetKubernetesNamespaceOk

`func (o *Environment) GetKubernetesNamespaceOk() (*string, bool)`

GetKubernetesNamespaceOk returns a tuple with the KubernetesNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesNamespace

`func (o *Environment) SetKubernetesNamespace(v string)`

SetKubernetesNamespace sets KubernetesNamespace field to given value.


### GetComputePoolDefaults

`func (o *Environment) GetComputePoolDefaults() map[string]interface{}`

GetComputePoolDefaults returns the ComputePoolDefaults field if non-nil, zero value otherwise.

### GetComputePoolDefaultsOk

`func (o *Environment) GetComputePoolDefaultsOk() (*map[string]interface{}, bool)`

GetComputePoolDefaultsOk returns a tuple with the ComputePoolDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputePoolDefaults

`func (o *Environment) SetComputePoolDefaults(v map[string]interface{})`

SetComputePoolDefaults sets ComputePoolDefaults field to given value.

### HasComputePoolDefaults

`func (o *Environment) HasComputePoolDefaults() bool`

HasComputePoolDefaults returns a boolean if a field has been set.

### GetStatementDefaults

`func (o *Environment) GetStatementDefaults() AllStatementDefaults1`

GetStatementDefaults returns the StatementDefaults field if non-nil, zero value otherwise.

### GetStatementDefaultsOk

`func (o *Environment) GetStatementDefaultsOk() (*AllStatementDefaults1, bool)`

GetStatementDefaultsOk returns a tuple with the StatementDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementDefaults

`func (o *Environment) SetStatementDefaults(v AllStatementDefaults1)`

SetStatementDefaults sets StatementDefaults field to given value.

### HasStatementDefaults

`func (o *Environment) HasStatementDefaults() bool`

HasStatementDefaults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


