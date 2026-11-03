# SqlAssistantTurnStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | The kind of response the model chose. | [optional] 
**Text** | Pointer to **string** | A markdown explanation (sql mode) or the final answer (text mode). | [optional] 
**TextPosition** | Pointer to **string** | sql mode - whether the explanation in text is shown before (above) or after (below) the SQL statement. Absent means before. | [optional] 
**Sql** | Pointer to **string** | The Flink SQL statement to run (sql mode). | [optional] 
**AgentStep** | Pointer to **bool** | sql mode - true to run the statement and post the rows back for another turn; false when the statement itself is the deliverable. | [optional] 
**Collect** | Pointer to [**SqlAssistantCollectBound**](SqlAssistantCollectBound.md) |  | [optional] 
**FlinkConfiguration** | Pointer to **map[string]string** | sql mode - the Flink configuration to run the statement with, as in a Statement&#39;s spec.flinkConfiguration. | [optional] 
**Parallelism** | Pointer to **int32** | sql mode - the parallelism to run the statement with, as in a Statement&#39;s spec.parallelism. Absent means 1. Never set for a SELECT. | [optional] 
**Properties** | Pointer to **map[string]string** | sql mode - the statement properties to run the statement with, as in a Statement&#39;s spec.properties. | [optional] 
**SqlKind** | Pointer to **string** | sql mode - the kind of the statement, with the values of a statement&#39;s &#x60;status.traits.sqlKind&#x60;, for example SELECT, SHOW_TABLES, CREATE_TABLE or INSERT_INTO. INSERT_INTO and EXECUTE_STATEMENT_SET run in the background, return no rows, and are never an agent step. Absent when the statement does not parse or when CMF does not support it. Clients should run only SELECT, SHOW_* and DESCRIBE* statements without the user&#39;s approval. | [optional] 
**Skills** | Pointer to **[]string** | The reference docs the assistant loaded to reply to this turn, for example flink-sql-config (the SQL configuration options of the compute pool&#39;s Flink version). They stay available to the assistant for the rest of the conversation. | [optional] 
**Question** | Pointer to [**SqlAssistantQuestion**](SqlAssistantQuestion.md) |  | [optional] 
**SuggestedTitle** | Pointer to **string** | A model-suggested conversation title. Omitted once spec.title is set on the conversation. | [optional] 
**Model** | Pointer to **string** | The model that produced this turn. | [optional] 
**PromptTokens** | Pointer to **int32** | The input (prompt) token count the model reported for this turn. | [optional] 

## Methods

### NewSqlAssistantTurnStatus

`func NewSqlAssistantTurnStatus() *SqlAssistantTurnStatus`

NewSqlAssistantTurnStatus instantiates a new SqlAssistantTurnStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlAssistantTurnStatusWithDefaults

`func NewSqlAssistantTurnStatusWithDefaults() *SqlAssistantTurnStatus`

NewSqlAssistantTurnStatusWithDefaults instantiates a new SqlAssistantTurnStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *SqlAssistantTurnStatus) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *SqlAssistantTurnStatus) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *SqlAssistantTurnStatus) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *SqlAssistantTurnStatus) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetText

`func (o *SqlAssistantTurnStatus) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *SqlAssistantTurnStatus) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *SqlAssistantTurnStatus) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *SqlAssistantTurnStatus) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTextPosition

`func (o *SqlAssistantTurnStatus) GetTextPosition() string`

GetTextPosition returns the TextPosition field if non-nil, zero value otherwise.

### GetTextPositionOk

`func (o *SqlAssistantTurnStatus) GetTextPositionOk() (*string, bool)`

GetTextPositionOk returns a tuple with the TextPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextPosition

`func (o *SqlAssistantTurnStatus) SetTextPosition(v string)`

SetTextPosition sets TextPosition field to given value.

### HasTextPosition

`func (o *SqlAssistantTurnStatus) HasTextPosition() bool`

HasTextPosition returns a boolean if a field has been set.

### GetSql

`func (o *SqlAssistantTurnStatus) GetSql() string`

GetSql returns the Sql field if non-nil, zero value otherwise.

### GetSqlOk

`func (o *SqlAssistantTurnStatus) GetSqlOk() (*string, bool)`

GetSqlOk returns a tuple with the Sql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSql

`func (o *SqlAssistantTurnStatus) SetSql(v string)`

SetSql sets Sql field to given value.

### HasSql

`func (o *SqlAssistantTurnStatus) HasSql() bool`

HasSql returns a boolean if a field has been set.

### GetAgentStep

`func (o *SqlAssistantTurnStatus) GetAgentStep() bool`

GetAgentStep returns the AgentStep field if non-nil, zero value otherwise.

### GetAgentStepOk

`func (o *SqlAssistantTurnStatus) GetAgentStepOk() (*bool, bool)`

GetAgentStepOk returns a tuple with the AgentStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentStep

`func (o *SqlAssistantTurnStatus) SetAgentStep(v bool)`

SetAgentStep sets AgentStep field to given value.

### HasAgentStep

`func (o *SqlAssistantTurnStatus) HasAgentStep() bool`

HasAgentStep returns a boolean if a field has been set.

### GetCollect

`func (o *SqlAssistantTurnStatus) GetCollect() SqlAssistantCollectBound`

GetCollect returns the Collect field if non-nil, zero value otherwise.

### GetCollectOk

`func (o *SqlAssistantTurnStatus) GetCollectOk() (*SqlAssistantCollectBound, bool)`

GetCollectOk returns a tuple with the Collect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollect

`func (o *SqlAssistantTurnStatus) SetCollect(v SqlAssistantCollectBound)`

SetCollect sets Collect field to given value.

### HasCollect

`func (o *SqlAssistantTurnStatus) HasCollect() bool`

HasCollect returns a boolean if a field has been set.

### GetFlinkConfiguration

`func (o *SqlAssistantTurnStatus) GetFlinkConfiguration() map[string]string`

GetFlinkConfiguration returns the FlinkConfiguration field if non-nil, zero value otherwise.

### GetFlinkConfigurationOk

`func (o *SqlAssistantTurnStatus) GetFlinkConfigurationOk() (*map[string]string, bool)`

GetFlinkConfigurationOk returns a tuple with the FlinkConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlinkConfiguration

`func (o *SqlAssistantTurnStatus) SetFlinkConfiguration(v map[string]string)`

SetFlinkConfiguration sets FlinkConfiguration field to given value.

### HasFlinkConfiguration

`func (o *SqlAssistantTurnStatus) HasFlinkConfiguration() bool`

HasFlinkConfiguration returns a boolean if a field has been set.

### GetParallelism

`func (o *SqlAssistantTurnStatus) GetParallelism() int32`

GetParallelism returns the Parallelism field if non-nil, zero value otherwise.

### GetParallelismOk

`func (o *SqlAssistantTurnStatus) GetParallelismOk() (*int32, bool)`

GetParallelismOk returns a tuple with the Parallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelism

`func (o *SqlAssistantTurnStatus) SetParallelism(v int32)`

SetParallelism sets Parallelism field to given value.

### HasParallelism

`func (o *SqlAssistantTurnStatus) HasParallelism() bool`

HasParallelism returns a boolean if a field has been set.

### GetProperties

`func (o *SqlAssistantTurnStatus) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SqlAssistantTurnStatus) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SqlAssistantTurnStatus) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *SqlAssistantTurnStatus) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetSqlKind

`func (o *SqlAssistantTurnStatus) GetSqlKind() string`

GetSqlKind returns the SqlKind field if non-nil, zero value otherwise.

### GetSqlKindOk

`func (o *SqlAssistantTurnStatus) GetSqlKindOk() (*string, bool)`

GetSqlKindOk returns a tuple with the SqlKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlKind

`func (o *SqlAssistantTurnStatus) SetSqlKind(v string)`

SetSqlKind sets SqlKind field to given value.

### HasSqlKind

`func (o *SqlAssistantTurnStatus) HasSqlKind() bool`

HasSqlKind returns a boolean if a field has been set.

### GetSkills

`func (o *SqlAssistantTurnStatus) GetSkills() []string`

GetSkills returns the Skills field if non-nil, zero value otherwise.

### GetSkillsOk

`func (o *SqlAssistantTurnStatus) GetSkillsOk() (*[]string, bool)`

GetSkillsOk returns a tuple with the Skills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkills

`func (o *SqlAssistantTurnStatus) SetSkills(v []string)`

SetSkills sets Skills field to given value.

### HasSkills

`func (o *SqlAssistantTurnStatus) HasSkills() bool`

HasSkills returns a boolean if a field has been set.

### GetQuestion

`func (o *SqlAssistantTurnStatus) GetQuestion() SqlAssistantQuestion`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *SqlAssistantTurnStatus) GetQuestionOk() (*SqlAssistantQuestion, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *SqlAssistantTurnStatus) SetQuestion(v SqlAssistantQuestion)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *SqlAssistantTurnStatus) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### GetSuggestedTitle

`func (o *SqlAssistantTurnStatus) GetSuggestedTitle() string`

GetSuggestedTitle returns the SuggestedTitle field if non-nil, zero value otherwise.

### GetSuggestedTitleOk

`func (o *SqlAssistantTurnStatus) GetSuggestedTitleOk() (*string, bool)`

GetSuggestedTitleOk returns a tuple with the SuggestedTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedTitle

`func (o *SqlAssistantTurnStatus) SetSuggestedTitle(v string)`

SetSuggestedTitle sets SuggestedTitle field to given value.

### HasSuggestedTitle

`func (o *SqlAssistantTurnStatus) HasSuggestedTitle() bool`

HasSuggestedTitle returns a boolean if a field has been set.

### GetModel

`func (o *SqlAssistantTurnStatus) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *SqlAssistantTurnStatus) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *SqlAssistantTurnStatus) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *SqlAssistantTurnStatus) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetPromptTokens

`func (o *SqlAssistantTurnStatus) GetPromptTokens() int32`

GetPromptTokens returns the PromptTokens field if non-nil, zero value otherwise.

### GetPromptTokensOk

`func (o *SqlAssistantTurnStatus) GetPromptTokensOk() (*int32, bool)`

GetPromptTokensOk returns a tuple with the PromptTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokens

`func (o *SqlAssistantTurnStatus) SetPromptTokens(v int32)`

SetPromptTokens sets PromptTokens field to given value.

### HasPromptTokens

`func (o *SqlAssistantTurnStatus) HasPromptTokens() bool`

HasPromptTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


