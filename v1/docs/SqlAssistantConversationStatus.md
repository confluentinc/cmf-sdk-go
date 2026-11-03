# SqlAssistantConversationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | The latest model-suggested title. The UI shows spec.title when set, otherwise this value. | [optional] 
**TurnCount** | Pointer to **int32** | Number of turns in the conversation. | [optional] 
**LastMode** | Pointer to **string** | The mode of the most recent model turn (sql, text, or question). | [optional] 

## Methods

### NewSqlAssistantConversationStatus

`func NewSqlAssistantConversationStatus() *SqlAssistantConversationStatus`

NewSqlAssistantConversationStatus instantiates a new SqlAssistantConversationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlAssistantConversationStatusWithDefaults

`func NewSqlAssistantConversationStatusWithDefaults() *SqlAssistantConversationStatus`

NewSqlAssistantConversationStatusWithDefaults instantiates a new SqlAssistantConversationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *SqlAssistantConversationStatus) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SqlAssistantConversationStatus) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SqlAssistantConversationStatus) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SqlAssistantConversationStatus) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTurnCount

`func (o *SqlAssistantConversationStatus) GetTurnCount() int32`

GetTurnCount returns the TurnCount field if non-nil, zero value otherwise.

### GetTurnCountOk

`func (o *SqlAssistantConversationStatus) GetTurnCountOk() (*int32, bool)`

GetTurnCountOk returns a tuple with the TurnCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTurnCount

`func (o *SqlAssistantConversationStatus) SetTurnCount(v int32)`

SetTurnCount sets TurnCount field to given value.

### HasTurnCount

`func (o *SqlAssistantConversationStatus) HasTurnCount() bool`

HasTurnCount returns a boolean if a field has been set.

### GetLastMode

`func (o *SqlAssistantConversationStatus) GetLastMode() string`

GetLastMode returns the LastMode field if non-nil, zero value otherwise.

### GetLastModeOk

`func (o *SqlAssistantConversationStatus) GetLastModeOk() (*string, bool)`

GetLastModeOk returns a tuple with the LastMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMode

`func (o *SqlAssistantConversationStatus) SetLastMode(v string)`

SetLastMode sets LastMode field to given value.

### HasLastMode

`func (o *SqlAssistantConversationStatus) HasLastMode() bool`

HasLastMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


