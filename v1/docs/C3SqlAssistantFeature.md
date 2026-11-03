# C3SqlAssistantFeature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether the SQL Assistant tab and its endpoints are served (cmf.sql.assistant.enabled). | [optional] 
**AiEnabled** | Pointer to **bool** | Whether a model is configured. When false the tab is a manual SQL shell. | [optional] 
**Model** | Pointer to **string** | The configured model. Absent when no model is configured. | [optional] 
**InputTokenLimit** | Pointer to **int64** | The configured model&#39;s context-window size, used as the token-gauge denominator. Absent when not configured. | [optional] 
**MaxResultRows** | Pointer to **int64** | The maximum rows a result turn may carry (cmf.sql.assistant.max-result-rows); the client caps a collected result to this before posting it. | [optional] 

## Methods

### NewC3SqlAssistantFeature

`func NewC3SqlAssistantFeature() *C3SqlAssistantFeature`

NewC3SqlAssistantFeature instantiates a new C3SqlAssistantFeature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC3SqlAssistantFeatureWithDefaults

`func NewC3SqlAssistantFeatureWithDefaults() *C3SqlAssistantFeature`

NewC3SqlAssistantFeatureWithDefaults instantiates a new C3SqlAssistantFeature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *C3SqlAssistantFeature) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *C3SqlAssistantFeature) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *C3SqlAssistantFeature) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *C3SqlAssistantFeature) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAiEnabled

`func (o *C3SqlAssistantFeature) GetAiEnabled() bool`

GetAiEnabled returns the AiEnabled field if non-nil, zero value otherwise.

### GetAiEnabledOk

`func (o *C3SqlAssistantFeature) GetAiEnabledOk() (*bool, bool)`

GetAiEnabledOk returns a tuple with the AiEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiEnabled

`func (o *C3SqlAssistantFeature) SetAiEnabled(v bool)`

SetAiEnabled sets AiEnabled field to given value.

### HasAiEnabled

`func (o *C3SqlAssistantFeature) HasAiEnabled() bool`

HasAiEnabled returns a boolean if a field has been set.

### GetModel

`func (o *C3SqlAssistantFeature) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *C3SqlAssistantFeature) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *C3SqlAssistantFeature) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *C3SqlAssistantFeature) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetInputTokenLimit

`func (o *C3SqlAssistantFeature) GetInputTokenLimit() int64`

GetInputTokenLimit returns the InputTokenLimit field if non-nil, zero value otherwise.

### GetInputTokenLimitOk

`func (o *C3SqlAssistantFeature) GetInputTokenLimitOk() (*int64, bool)`

GetInputTokenLimitOk returns a tuple with the InputTokenLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTokenLimit

`func (o *C3SqlAssistantFeature) SetInputTokenLimit(v int64)`

SetInputTokenLimit sets InputTokenLimit field to given value.

### HasInputTokenLimit

`func (o *C3SqlAssistantFeature) HasInputTokenLimit() bool`

HasInputTokenLimit returns a boolean if a field has been set.

### GetMaxResultRows

`func (o *C3SqlAssistantFeature) GetMaxResultRows() int64`

GetMaxResultRows returns the MaxResultRows field if non-nil, zero value otherwise.

### GetMaxResultRowsOk

`func (o *C3SqlAssistantFeature) GetMaxResultRowsOk() (*int64, bool)`

GetMaxResultRowsOk returns a tuple with the MaxResultRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResultRows

`func (o *C3SqlAssistantFeature) SetMaxResultRows(v int64)`

SetMaxResultRows sets MaxResultRows field to given value.

### HasMaxResultRows

`func (o *C3SqlAssistantFeature) HasMaxResultRows() bool`

HasMaxResultRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


