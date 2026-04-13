# EnvironmentMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the Environment | 
**CreationTimestamp** | Pointer to **string** | Timestamp when the Environment was created | [optional] [readonly] 
**UpdateTimestamp** | Pointer to **string** | Timestamp when the Environment was updated the last time | [optional] [readonly] 
**Labels** | Pointer to **map[string]string** | Labels of the Environment | [optional] 
**Annotations** | Pointer to **map[string]string** | Annotations of the Environment | [optional] 

## Methods

### NewEnvironmentMetadata

`func NewEnvironmentMetadata(name string, ) *EnvironmentMetadata`

NewEnvironmentMetadata instantiates a new EnvironmentMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentMetadataWithDefaults

`func NewEnvironmentMetadataWithDefaults() *EnvironmentMetadata`

NewEnvironmentMetadataWithDefaults instantiates a new EnvironmentMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EnvironmentMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentMetadata) SetName(v string)`

SetName sets Name field to given value.


### GetCreationTimestamp

`func (o *EnvironmentMetadata) GetCreationTimestamp() string`

GetCreationTimestamp returns the CreationTimestamp field if non-nil, zero value otherwise.

### GetCreationTimestampOk

`func (o *EnvironmentMetadata) GetCreationTimestampOk() (*string, bool)`

GetCreationTimestampOk returns a tuple with the CreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimestamp

`func (o *EnvironmentMetadata) SetCreationTimestamp(v string)`

SetCreationTimestamp sets CreationTimestamp field to given value.

### HasCreationTimestamp

`func (o *EnvironmentMetadata) HasCreationTimestamp() bool`

HasCreationTimestamp returns a boolean if a field has been set.

### GetUpdateTimestamp

`func (o *EnvironmentMetadata) GetUpdateTimestamp() string`

GetUpdateTimestamp returns the UpdateTimestamp field if non-nil, zero value otherwise.

### GetUpdateTimestampOk

`func (o *EnvironmentMetadata) GetUpdateTimestampOk() (*string, bool)`

GetUpdateTimestampOk returns a tuple with the UpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTimestamp

`func (o *EnvironmentMetadata) SetUpdateTimestamp(v string)`

SetUpdateTimestamp sets UpdateTimestamp field to given value.

### HasUpdateTimestamp

`func (o *EnvironmentMetadata) HasUpdateTimestamp() bool`

HasUpdateTimestamp returns a boolean if a field has been set.

### GetLabels

`func (o *EnvironmentMetadata) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *EnvironmentMetadata) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *EnvironmentMetadata) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *EnvironmentMetadata) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetAnnotations

`func (o *EnvironmentMetadata) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *EnvironmentMetadata) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *EnvironmentMetadata) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *EnvironmentMetadata) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


