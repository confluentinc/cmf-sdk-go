# SqlAssistantConversationMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Server-assigned unique identifier of the conversation. Immutable. | [optional] [readonly] 
**Uid** | Pointer to **string** | Unique identifier of the conversation. Identical to name. | [optional] [readonly] 
**CreationTimestamp** | Pointer to **time.Time** | Timestamp when the conversation was created. | [optional] [readonly] 
**UpdateTimestamp** | Pointer to **time.Time** | Timestamp when the conversation was last updated. | [optional] [readonly] 
**Labels** | Pointer to **map[string]string** | Labels of the conversation. Omit the field (or send null) to leave existing labels unchanged; send an empty object to clear them. | [optional] 
**Annotations** | Pointer to **map[string]string** | Annotations of the conversation. Omit the field (or send null) to leave existing annotations unchanged; send an empty object to clear them. | [optional] 

## Methods

### NewSqlAssistantConversationMetadata

`func NewSqlAssistantConversationMetadata() *SqlAssistantConversationMetadata`

NewSqlAssistantConversationMetadata instantiates a new SqlAssistantConversationMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlAssistantConversationMetadataWithDefaults

`func NewSqlAssistantConversationMetadataWithDefaults() *SqlAssistantConversationMetadata`

NewSqlAssistantConversationMetadataWithDefaults instantiates a new SqlAssistantConversationMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SqlAssistantConversationMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SqlAssistantConversationMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SqlAssistantConversationMetadata) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SqlAssistantConversationMetadata) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUid

`func (o *SqlAssistantConversationMetadata) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *SqlAssistantConversationMetadata) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *SqlAssistantConversationMetadata) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *SqlAssistantConversationMetadata) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetCreationTimestamp

`func (o *SqlAssistantConversationMetadata) GetCreationTimestamp() time.Time`

GetCreationTimestamp returns the CreationTimestamp field if non-nil, zero value otherwise.

### GetCreationTimestampOk

`func (o *SqlAssistantConversationMetadata) GetCreationTimestampOk() (*time.Time, bool)`

GetCreationTimestampOk returns a tuple with the CreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimestamp

`func (o *SqlAssistantConversationMetadata) SetCreationTimestamp(v time.Time)`

SetCreationTimestamp sets CreationTimestamp field to given value.

### HasCreationTimestamp

`func (o *SqlAssistantConversationMetadata) HasCreationTimestamp() bool`

HasCreationTimestamp returns a boolean if a field has been set.

### GetUpdateTimestamp

`func (o *SqlAssistantConversationMetadata) GetUpdateTimestamp() time.Time`

GetUpdateTimestamp returns the UpdateTimestamp field if non-nil, zero value otherwise.

### GetUpdateTimestampOk

`func (o *SqlAssistantConversationMetadata) GetUpdateTimestampOk() (*time.Time, bool)`

GetUpdateTimestampOk returns a tuple with the UpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTimestamp

`func (o *SqlAssistantConversationMetadata) SetUpdateTimestamp(v time.Time)`

SetUpdateTimestamp sets UpdateTimestamp field to given value.

### HasUpdateTimestamp

`func (o *SqlAssistantConversationMetadata) HasUpdateTimestamp() bool`

HasUpdateTimestamp returns a boolean if a field has been set.

### GetLabels

`func (o *SqlAssistantConversationMetadata) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *SqlAssistantConversationMetadata) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *SqlAssistantConversationMetadata) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *SqlAssistantConversationMetadata) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *SqlAssistantConversationMetadata) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *SqlAssistantConversationMetadata) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetAnnotations

`func (o *SqlAssistantConversationMetadata) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *SqlAssistantConversationMetadata) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *SqlAssistantConversationMetadata) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *SqlAssistantConversationMetadata) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### SetAnnotationsNil

`func (o *SqlAssistantConversationMetadata) SetAnnotationsNil(b bool)`

 SetAnnotationsNil sets the value for Annotations to be an explicit nil

### UnsetAnnotations
`func (o *SqlAssistantConversationMetadata) UnsetAnnotations()`

UnsetAnnotations ensures that no value is present for Annotations, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


