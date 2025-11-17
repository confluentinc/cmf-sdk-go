# StatementResultMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationTimestamp** | Pointer to **string** | Timestamp when the StatementResult was created | [optional] 
**Annotations** | Pointer to **map[string]string** | Annotations of the StatementResult | [optional] 

## Methods

### NewStatementResultMetadata

`func NewStatementResultMetadata() *StatementResultMetadata`

NewStatementResultMetadata instantiates a new StatementResultMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatementResultMetadataWithDefaults

`func NewStatementResultMetadataWithDefaults() *StatementResultMetadata`

NewStatementResultMetadataWithDefaults instantiates a new StatementResultMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationTimestamp

`func (o *StatementResultMetadata) GetCreationTimestamp() string`

GetCreationTimestamp returns the CreationTimestamp field if non-nil, zero value otherwise.

### GetCreationTimestampOk

`func (o *StatementResultMetadata) GetCreationTimestampOk() (*string, bool)`

GetCreationTimestampOk returns a tuple with the CreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimestamp

`func (o *StatementResultMetadata) SetCreationTimestamp(v string)`

SetCreationTimestamp sets CreationTimestamp field to given value.

### HasCreationTimestamp

`func (o *StatementResultMetadata) HasCreationTimestamp() bool`

HasCreationTimestamp returns a boolean if a field has been set.

### GetAnnotations

`func (o *StatementResultMetadata) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *StatementResultMetadata) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *StatementResultMetadata) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *StatementResultMetadata) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


