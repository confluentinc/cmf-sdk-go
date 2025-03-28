# StatementDefaults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlinkConfiguration** | Pointer to **map[string]string** | default Flink configuration for Statements | [optional] 

## Methods

### NewStatementDefaults

`func NewStatementDefaults() *StatementDefaults`

NewStatementDefaults instantiates a new StatementDefaults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatementDefaultsWithDefaults

`func NewStatementDefaultsWithDefaults() *StatementDefaults`

NewStatementDefaultsWithDefaults instantiates a new StatementDefaults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlinkConfiguration

`func (o *StatementDefaults) GetFlinkConfiguration() map[string]string`

GetFlinkConfiguration returns the FlinkConfiguration field if non-nil, zero value otherwise.

### GetFlinkConfigurationOk

`func (o *StatementDefaults) GetFlinkConfigurationOk() (*map[string]string, bool)`

GetFlinkConfigurationOk returns a tuple with the FlinkConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlinkConfiguration

`func (o *StatementDefaults) SetFlinkConfiguration(v map[string]string)`

SetFlinkConfiguration sets FlinkConfiguration field to given value.

### HasFlinkConfiguration

`func (o *StatementDefaults) HasFlinkConfiguration() bool`

HasFlinkConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


