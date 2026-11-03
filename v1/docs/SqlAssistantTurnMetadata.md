# SqlAssistantTurnMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The turn&#39;s sequence number within the conversation, as a string. | [optional] [readonly] 
**CreationTimestamp** | Pointer to **time.Time** | Timestamp when the turn was created. | [optional] [readonly] 

## Methods

### NewSqlAssistantTurnMetadata

`func NewSqlAssistantTurnMetadata() *SqlAssistantTurnMetadata`

NewSqlAssistantTurnMetadata instantiates a new SqlAssistantTurnMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlAssistantTurnMetadataWithDefaults

`func NewSqlAssistantTurnMetadataWithDefaults() *SqlAssistantTurnMetadata`

NewSqlAssistantTurnMetadataWithDefaults instantiates a new SqlAssistantTurnMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SqlAssistantTurnMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SqlAssistantTurnMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SqlAssistantTurnMetadata) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SqlAssistantTurnMetadata) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreationTimestamp

`func (o *SqlAssistantTurnMetadata) GetCreationTimestamp() time.Time`

GetCreationTimestamp returns the CreationTimestamp field if non-nil, zero value otherwise.

### GetCreationTimestampOk

`func (o *SqlAssistantTurnMetadata) GetCreationTimestampOk() (*time.Time, bool)`

GetCreationTimestampOk returns a tuple with the CreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimestamp

`func (o *SqlAssistantTurnMetadata) SetCreationTimestamp(v time.Time)`

SetCreationTimestamp sets CreationTimestamp field to given value.

### HasCreationTimestamp

`func (o *SqlAssistantTurnMetadata) HasCreationTimestamp() bool`

HasCreationTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


