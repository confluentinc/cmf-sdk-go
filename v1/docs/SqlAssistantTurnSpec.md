# SqlAssistantTurnSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | **string** | user - a prompt or an answer to a question; sql - a SQL statement the client runs directly without the model; result - the outcome of running a prior SQL turn. | 
**Text** | Pointer to **string** | The user&#39;s prompt (role user). | [optional] 
**Sql** | Pointer to **string** | A Flink SQL statement to run directly, without the model (role sql), or the statement as it ran, including the user&#39;s edits (role result). In CREATE MODEL and ALTER MODEL statements, option values are stored and returned masked as SHOW CREATE MODEL shows them, for example &#39;api-key&#39; &#x3D; &#39;******&#39;. | [optional] 
**FlinkConfiguration** | Pointer to **map[string]string** | The Flink configuration the statement ran with, as in a Statement&#39;s spec.flinkConfiguration (roles sql and result). | [optional] 
**Parallelism** | Pointer to **int32** | The parallelism the statement ran with, as in a Statement&#39;s spec.parallelism (roles sql and result). | [optional] 
**Properties** | Pointer to **map[string]string** | The statement properties the statement ran with, as in a Statement&#39;s spec.properties (roles sql and result). | [optional] 
**SelectedOptions** | Pointer to **[]string** | Option ids chosen in answer to a preceding question turn (role user). | [optional] 
**ForTurn** | Pointer to **string** | The sequence number of the SQL turn whose result this is (role result). | [optional] 
**StatementName** | Pointer to **string** | Name of the Statement that produced these rows (role result). | [optional] 
**Schema** | Pointer to **[]string** | Column names of the result (role result). | [optional] 
**Rows** | Pointer to **[][]interface{}** | Result rows, each an array of column values (role result). Bounded by cmf.sql.assistant.max-result-rows. | [optional] 
**Error** | Pointer to **string** | The statement&#39;s error text when the run failed (role result). | [optional] 
**RequestFix** | Pointer to **bool** | true to ask the model for a corrected statement after a failed run (role result, requires error). The result of an agent step always gets a model reply; any other result gets one only with requestFix. | [optional] 

## Methods

### NewSqlAssistantTurnSpec

`func NewSqlAssistantTurnSpec(role string, ) *SqlAssistantTurnSpec`

NewSqlAssistantTurnSpec instantiates a new SqlAssistantTurnSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlAssistantTurnSpecWithDefaults

`func NewSqlAssistantTurnSpecWithDefaults() *SqlAssistantTurnSpec`

NewSqlAssistantTurnSpecWithDefaults instantiates a new SqlAssistantTurnSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *SqlAssistantTurnSpec) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *SqlAssistantTurnSpec) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *SqlAssistantTurnSpec) SetRole(v string)`

SetRole sets Role field to given value.


### GetText

`func (o *SqlAssistantTurnSpec) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *SqlAssistantTurnSpec) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *SqlAssistantTurnSpec) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *SqlAssistantTurnSpec) HasText() bool`

HasText returns a boolean if a field has been set.

### GetSql

`func (o *SqlAssistantTurnSpec) GetSql() string`

GetSql returns the Sql field if non-nil, zero value otherwise.

### GetSqlOk

`func (o *SqlAssistantTurnSpec) GetSqlOk() (*string, bool)`

GetSqlOk returns a tuple with the Sql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSql

`func (o *SqlAssistantTurnSpec) SetSql(v string)`

SetSql sets Sql field to given value.

### HasSql

`func (o *SqlAssistantTurnSpec) HasSql() bool`

HasSql returns a boolean if a field has been set.

### GetFlinkConfiguration

`func (o *SqlAssistantTurnSpec) GetFlinkConfiguration() map[string]string`

GetFlinkConfiguration returns the FlinkConfiguration field if non-nil, zero value otherwise.

### GetFlinkConfigurationOk

`func (o *SqlAssistantTurnSpec) GetFlinkConfigurationOk() (*map[string]string, bool)`

GetFlinkConfigurationOk returns a tuple with the FlinkConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlinkConfiguration

`func (o *SqlAssistantTurnSpec) SetFlinkConfiguration(v map[string]string)`

SetFlinkConfiguration sets FlinkConfiguration field to given value.

### HasFlinkConfiguration

`func (o *SqlAssistantTurnSpec) HasFlinkConfiguration() bool`

HasFlinkConfiguration returns a boolean if a field has been set.

### GetParallelism

`func (o *SqlAssistantTurnSpec) GetParallelism() int32`

GetParallelism returns the Parallelism field if non-nil, zero value otherwise.

### GetParallelismOk

`func (o *SqlAssistantTurnSpec) GetParallelismOk() (*int32, bool)`

GetParallelismOk returns a tuple with the Parallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelism

`func (o *SqlAssistantTurnSpec) SetParallelism(v int32)`

SetParallelism sets Parallelism field to given value.

### HasParallelism

`func (o *SqlAssistantTurnSpec) HasParallelism() bool`

HasParallelism returns a boolean if a field has been set.

### GetProperties

`func (o *SqlAssistantTurnSpec) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SqlAssistantTurnSpec) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SqlAssistantTurnSpec) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *SqlAssistantTurnSpec) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetSelectedOptions

`func (o *SqlAssistantTurnSpec) GetSelectedOptions() []string`

GetSelectedOptions returns the SelectedOptions field if non-nil, zero value otherwise.

### GetSelectedOptionsOk

`func (o *SqlAssistantTurnSpec) GetSelectedOptionsOk() (*[]string, bool)`

GetSelectedOptionsOk returns a tuple with the SelectedOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedOptions

`func (o *SqlAssistantTurnSpec) SetSelectedOptions(v []string)`

SetSelectedOptions sets SelectedOptions field to given value.

### HasSelectedOptions

`func (o *SqlAssistantTurnSpec) HasSelectedOptions() bool`

HasSelectedOptions returns a boolean if a field has been set.

### GetForTurn

`func (o *SqlAssistantTurnSpec) GetForTurn() string`

GetForTurn returns the ForTurn field if non-nil, zero value otherwise.

### GetForTurnOk

`func (o *SqlAssistantTurnSpec) GetForTurnOk() (*string, bool)`

GetForTurnOk returns a tuple with the ForTurn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForTurn

`func (o *SqlAssistantTurnSpec) SetForTurn(v string)`

SetForTurn sets ForTurn field to given value.

### HasForTurn

`func (o *SqlAssistantTurnSpec) HasForTurn() bool`

HasForTurn returns a boolean if a field has been set.

### GetStatementName

`func (o *SqlAssistantTurnSpec) GetStatementName() string`

GetStatementName returns the StatementName field if non-nil, zero value otherwise.

### GetStatementNameOk

`func (o *SqlAssistantTurnSpec) GetStatementNameOk() (*string, bool)`

GetStatementNameOk returns a tuple with the StatementName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementName

`func (o *SqlAssistantTurnSpec) SetStatementName(v string)`

SetStatementName sets StatementName field to given value.

### HasStatementName

`func (o *SqlAssistantTurnSpec) HasStatementName() bool`

HasStatementName returns a boolean if a field has been set.

### GetSchema

`func (o *SqlAssistantTurnSpec) GetSchema() []string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *SqlAssistantTurnSpec) GetSchemaOk() (*[]string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *SqlAssistantTurnSpec) SetSchema(v []string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *SqlAssistantTurnSpec) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetRows

`func (o *SqlAssistantTurnSpec) GetRows() [][]interface{}`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *SqlAssistantTurnSpec) GetRowsOk() (*[][]interface{}, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *SqlAssistantTurnSpec) SetRows(v [][]interface{})`

SetRows sets Rows field to given value.

### HasRows

`func (o *SqlAssistantTurnSpec) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetError

`func (o *SqlAssistantTurnSpec) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SqlAssistantTurnSpec) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SqlAssistantTurnSpec) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *SqlAssistantTurnSpec) HasError() bool`

HasError returns a boolean if a field has been set.

### GetRequestFix

`func (o *SqlAssistantTurnSpec) GetRequestFix() bool`

GetRequestFix returns the RequestFix field if non-nil, zero value otherwise.

### GetRequestFixOk

`func (o *SqlAssistantTurnSpec) GetRequestFixOk() (*bool, bool)`

GetRequestFixOk returns a tuple with the RequestFix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestFix

`func (o *SqlAssistantTurnSpec) SetRequestFix(v bool)`

SetRequestFix sets RequestFix field to given value.

### HasRequestFix

`func (o *SqlAssistantTurnSpec) HasRequestFix() bool`

HasRequestFix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


