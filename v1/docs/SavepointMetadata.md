# SavepointMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the Savepoint | [optional] 
**CreationTimestamp** | Pointer to **string** | Timestamp when the Savepoint was created | [optional] 
**Uid** | Pointer to **string** | Unique identifier of the Savepoint | [optional] 
**Labels** | Pointer to **map[string]string** | Labels of the Savepoint | [optional] 
**Annotations** | Pointer to **map[string]string** | Annotations of the Savepoint | [optional] 

## Methods

### NewSavepointMetadata

`func NewSavepointMetadata() *SavepointMetadata`

NewSavepointMetadata instantiates a new SavepointMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSavepointMetadataWithDefaults

`func NewSavepointMetadataWithDefaults() *SavepointMetadata`

NewSavepointMetadataWithDefaults instantiates a new SavepointMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SavepointMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SavepointMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SavepointMetadata) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SavepointMetadata) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreationTimestamp

`func (o *SavepointMetadata) GetCreationTimestamp() string`

GetCreationTimestamp returns the CreationTimestamp field if non-nil, zero value otherwise.

### GetCreationTimestampOk

`func (o *SavepointMetadata) GetCreationTimestampOk() (*string, bool)`

GetCreationTimestampOk returns a tuple with the CreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimestamp

`func (o *SavepointMetadata) SetCreationTimestamp(v string)`

SetCreationTimestamp sets CreationTimestamp field to given value.

### HasCreationTimestamp

`func (o *SavepointMetadata) HasCreationTimestamp() bool`

HasCreationTimestamp returns a boolean if a field has been set.

### GetUid

`func (o *SavepointMetadata) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *SavepointMetadata) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *SavepointMetadata) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *SavepointMetadata) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetLabels

`func (o *SavepointMetadata) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *SavepointMetadata) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *SavepointMetadata) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *SavepointMetadata) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetAnnotations

`func (o *SavepointMetadata) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *SavepointMetadata) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *SavepointMetadata) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *SavepointMetadata) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


