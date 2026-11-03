# SqlAssistantQuestion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | Pointer to **string** |  | [optional] 
**Selection** | Pointer to **string** | Whether the user may pick one option or several. | [optional] 
**Options** | Pointer to [**[]SqlAssistantQuestionOption**](SqlAssistantQuestionOption.md) |  | [optional] 

## Methods

### NewSqlAssistantQuestion

`func NewSqlAssistantQuestion() *SqlAssistantQuestion`

NewSqlAssistantQuestion instantiates a new SqlAssistantQuestion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlAssistantQuestionWithDefaults

`func NewSqlAssistantQuestionWithDefaults() *SqlAssistantQuestion`

NewSqlAssistantQuestionWithDefaults instantiates a new SqlAssistantQuestion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *SqlAssistantQuestion) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *SqlAssistantQuestion) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *SqlAssistantQuestion) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *SqlAssistantQuestion) HasText() bool`

HasText returns a boolean if a field has been set.

### GetSelection

`func (o *SqlAssistantQuestion) GetSelection() string`

GetSelection returns the Selection field if non-nil, zero value otherwise.

### GetSelectionOk

`func (o *SqlAssistantQuestion) GetSelectionOk() (*string, bool)`

GetSelectionOk returns a tuple with the Selection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelection

`func (o *SqlAssistantQuestion) SetSelection(v string)`

SetSelection sets Selection field to given value.

### HasSelection

`func (o *SqlAssistantQuestion) HasSelection() bool`

HasSelection returns a boolean if a field has been set.

### GetOptions

`func (o *SqlAssistantQuestion) GetOptions() []SqlAssistantQuestionOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SqlAssistantQuestion) GetOptionsOk() (*[]SqlAssistantQuestionOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *SqlAssistantQuestion) SetOptions(v []SqlAssistantQuestionOption)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *SqlAssistantQuestion) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


