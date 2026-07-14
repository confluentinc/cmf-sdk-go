# ArtifactMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique identifier of the Artifact within the Environment. Lowercase alphanumeric characters and hyphens, starting with a letter and ending with a letter or digit, optionally followed by a single &#x60;.&#x60; file extension (for example &#x60;my-udf&#x60; or &#x60;my-udf.jar&#x60;). Maximum 45 characters. The extension is reflected in the stored object&#39;s path; an Artifact referenced as a SQL UDF must be named with a &#x60;.jar&#x60; extension.  | 
**Uid** | Pointer to **string** | Unique identifier of the Artifact | [optional] [readonly] 
**CreationTimestamp** | Pointer to **string** | Timestamp when the Artifact was first uploaded | [optional] [readonly] 
**UpdateTimestamp** | Pointer to **string** | Timestamp when the Artifact was last updated | [optional] [readonly] 
**Labels** | Pointer to **map[string]string** | Labels of the Artifact. Omit the field (or send null) to leave existing labels unchanged; send an empty object to clear them.  | [optional] 
**Annotations** | Pointer to **map[string]string** | Annotations of the Artifact. Omit the field (or send null) to leave existing annotations unchanged; send an empty object to clear them.  | [optional] 

## Methods

### NewArtifactMetadata

`func NewArtifactMetadata(name string, ) *ArtifactMetadata`

NewArtifactMetadata instantiates a new ArtifactMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactMetadataWithDefaults

`func NewArtifactMetadataWithDefaults() *ArtifactMetadata`

NewArtifactMetadataWithDefaults instantiates a new ArtifactMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ArtifactMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArtifactMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArtifactMetadata) SetName(v string)`

SetName sets Name field to given value.


### GetUid

`func (o *ArtifactMetadata) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ArtifactMetadata) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ArtifactMetadata) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *ArtifactMetadata) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetCreationTimestamp

`func (o *ArtifactMetadata) GetCreationTimestamp() string`

GetCreationTimestamp returns the CreationTimestamp field if non-nil, zero value otherwise.

### GetCreationTimestampOk

`func (o *ArtifactMetadata) GetCreationTimestampOk() (*string, bool)`

GetCreationTimestampOk returns a tuple with the CreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimestamp

`func (o *ArtifactMetadata) SetCreationTimestamp(v string)`

SetCreationTimestamp sets CreationTimestamp field to given value.

### HasCreationTimestamp

`func (o *ArtifactMetadata) HasCreationTimestamp() bool`

HasCreationTimestamp returns a boolean if a field has been set.

### GetUpdateTimestamp

`func (o *ArtifactMetadata) GetUpdateTimestamp() string`

GetUpdateTimestamp returns the UpdateTimestamp field if non-nil, zero value otherwise.

### GetUpdateTimestampOk

`func (o *ArtifactMetadata) GetUpdateTimestampOk() (*string, bool)`

GetUpdateTimestampOk returns a tuple with the UpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTimestamp

`func (o *ArtifactMetadata) SetUpdateTimestamp(v string)`

SetUpdateTimestamp sets UpdateTimestamp field to given value.

### HasUpdateTimestamp

`func (o *ArtifactMetadata) HasUpdateTimestamp() bool`

HasUpdateTimestamp returns a boolean if a field has been set.

### GetLabels

`func (o *ArtifactMetadata) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ArtifactMetadata) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ArtifactMetadata) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ArtifactMetadata) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *ArtifactMetadata) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *ArtifactMetadata) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetAnnotations

`func (o *ArtifactMetadata) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *ArtifactMetadata) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *ArtifactMetadata) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *ArtifactMetadata) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### SetAnnotationsNil

`func (o *ArtifactMetadata) SetAnnotationsNil(b bool)`

 SetAnnotationsNil sets the value for Annotations to be an explicit nil

### UnsetAnnotations
`func (o *ArtifactMetadata) UnsetAnnotations()`

UnsetAnnotations ensures that no value is present for Annotations, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


