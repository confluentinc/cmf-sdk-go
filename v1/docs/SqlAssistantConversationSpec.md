# SqlAssistantConversationSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComputePool** | **string** | The compute pool used to read the catalog context and to run the conversation&#39;s SQL. Required at creation. An update can move the conversation to another existing compute pool; an empty string keeps the current pool. | 
**Title** | Pointer to **NullableString** | The user-chosen title. Once set, the model stops suggesting titles and its suggestions no longer change it. | [optional] 

## Methods

### NewSqlAssistantConversationSpec

`func NewSqlAssistantConversationSpec(computePool string, ) *SqlAssistantConversationSpec`

NewSqlAssistantConversationSpec instantiates a new SqlAssistantConversationSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlAssistantConversationSpecWithDefaults

`func NewSqlAssistantConversationSpecWithDefaults() *SqlAssistantConversationSpec`

NewSqlAssistantConversationSpecWithDefaults instantiates a new SqlAssistantConversationSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComputePool

`func (o *SqlAssistantConversationSpec) GetComputePool() string`

GetComputePool returns the ComputePool field if non-nil, zero value otherwise.

### GetComputePoolOk

`func (o *SqlAssistantConversationSpec) GetComputePoolOk() (*string, bool)`

GetComputePoolOk returns a tuple with the ComputePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputePool

`func (o *SqlAssistantConversationSpec) SetComputePool(v string)`

SetComputePool sets ComputePool field to given value.


### GetTitle

`func (o *SqlAssistantConversationSpec) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SqlAssistantConversationSpec) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SqlAssistantConversationSpec) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SqlAssistantConversationSpec) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *SqlAssistantConversationSpec) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *SqlAssistantConversationSpec) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


