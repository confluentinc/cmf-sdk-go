# PostEnvironment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A unique name for the resource. | 
**FlinkApplicationDefaults** | Pointer to **map[string]interface{}** |  | [optional] 
**KubernetesNamespace** | Pointer to **string** |  | [optional] 
**ComputePoolDefaults** | Pointer to **map[string]interface{}** | the defaults as YAML or JSON for ComputePools | [optional] 
**StatementDefaults** | Pointer to [**AllStatementDefaults1**](AllStatementDefaults1.md) |  | [optional] 

## Methods

### NewPostEnvironment

`func NewPostEnvironment(name string, ) *PostEnvironment`

NewPostEnvironment instantiates a new PostEnvironment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostEnvironmentWithDefaults

`func NewPostEnvironmentWithDefaults() *PostEnvironment`

NewPostEnvironmentWithDefaults instantiates a new PostEnvironment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PostEnvironment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostEnvironment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostEnvironment) SetName(v string)`

SetName sets Name field to given value.


### GetFlinkApplicationDefaults

`func (o *PostEnvironment) GetFlinkApplicationDefaults() map[string]interface{}`

GetFlinkApplicationDefaults returns the FlinkApplicationDefaults field if non-nil, zero value otherwise.

### GetFlinkApplicationDefaultsOk

`func (o *PostEnvironment) GetFlinkApplicationDefaultsOk() (*map[string]interface{}, bool)`

GetFlinkApplicationDefaultsOk returns a tuple with the FlinkApplicationDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlinkApplicationDefaults

`func (o *PostEnvironment) SetFlinkApplicationDefaults(v map[string]interface{})`

SetFlinkApplicationDefaults sets FlinkApplicationDefaults field to given value.

### HasFlinkApplicationDefaults

`func (o *PostEnvironment) HasFlinkApplicationDefaults() bool`

HasFlinkApplicationDefaults returns a boolean if a field has been set.

### GetKubernetesNamespace

`func (o *PostEnvironment) GetKubernetesNamespace() string`

GetKubernetesNamespace returns the KubernetesNamespace field if non-nil, zero value otherwise.

### GetKubernetesNamespaceOk

`func (o *PostEnvironment) GetKubernetesNamespaceOk() (*string, bool)`

GetKubernetesNamespaceOk returns a tuple with the KubernetesNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesNamespace

`func (o *PostEnvironment) SetKubernetesNamespace(v string)`

SetKubernetesNamespace sets KubernetesNamespace field to given value.

### HasKubernetesNamespace

`func (o *PostEnvironment) HasKubernetesNamespace() bool`

HasKubernetesNamespace returns a boolean if a field has been set.

### GetComputePoolDefaults

`func (o *PostEnvironment) GetComputePoolDefaults() map[string]interface{}`

GetComputePoolDefaults returns the ComputePoolDefaults field if non-nil, zero value otherwise.

### GetComputePoolDefaultsOk

`func (o *PostEnvironment) GetComputePoolDefaultsOk() (*map[string]interface{}, bool)`

GetComputePoolDefaultsOk returns a tuple with the ComputePoolDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputePoolDefaults

`func (o *PostEnvironment) SetComputePoolDefaults(v map[string]interface{})`

SetComputePoolDefaults sets ComputePoolDefaults field to given value.

### HasComputePoolDefaults

`func (o *PostEnvironment) HasComputePoolDefaults() bool`

HasComputePoolDefaults returns a boolean if a field has been set.

### GetStatementDefaults

`func (o *PostEnvironment) GetStatementDefaults() AllStatementDefaults1`

GetStatementDefaults returns the StatementDefaults field if non-nil, zero value otherwise.

### GetStatementDefaultsOk

`func (o *PostEnvironment) GetStatementDefaultsOk() (*AllStatementDefaults1, bool)`

GetStatementDefaultsOk returns a tuple with the StatementDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementDefaults

`func (o *PostEnvironment) SetStatementDefaults(v AllStatementDefaults1)`

SetStatementDefaults sets StatementDefaults field to given value.

### HasStatementDefaults

`func (o *PostEnvironment) HasStatementDefaults() bool`

HasStatementDefaults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


