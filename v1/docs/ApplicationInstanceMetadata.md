# ApplicationInstanceMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the Instance - a uuid. | [optional] 
**Uid** | Pointer to **string** | Unique identifier of the instance. Identical to name. | [optional] 
**CreationTimestamp** | Pointer to **string** | Timestamp when the Instance was created | [optional] 
**UpdateTimestamp** | Pointer to **string** | Timestamp when the Instance status was last updated | [optional] 
**Labels** | Pointer to **map[string]string** | Labels of the instance | [optional] 
**Annotations** | Pointer to **map[string]string** | Annotations of the instance | [optional] 

## Methods

### NewApplicationInstanceMetadata

`func NewApplicationInstanceMetadata() *ApplicationInstanceMetadata`

NewApplicationInstanceMetadata instantiates a new ApplicationInstanceMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationInstanceMetadataWithDefaults

`func NewApplicationInstanceMetadataWithDefaults() *ApplicationInstanceMetadata`

NewApplicationInstanceMetadataWithDefaults instantiates a new ApplicationInstanceMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApplicationInstanceMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationInstanceMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationInstanceMetadata) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplicationInstanceMetadata) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUid

`func (o *ApplicationInstanceMetadata) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ApplicationInstanceMetadata) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ApplicationInstanceMetadata) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *ApplicationInstanceMetadata) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetCreationTimestamp

`func (o *ApplicationInstanceMetadata) GetCreationTimestamp() string`

GetCreationTimestamp returns the CreationTimestamp field if non-nil, zero value otherwise.

### GetCreationTimestampOk

`func (o *ApplicationInstanceMetadata) GetCreationTimestampOk() (*string, bool)`

GetCreationTimestampOk returns a tuple with the CreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimestamp

`func (o *ApplicationInstanceMetadata) SetCreationTimestamp(v string)`

SetCreationTimestamp sets CreationTimestamp field to given value.

### HasCreationTimestamp

`func (o *ApplicationInstanceMetadata) HasCreationTimestamp() bool`

HasCreationTimestamp returns a boolean if a field has been set.

### GetUpdateTimestamp

`func (o *ApplicationInstanceMetadata) GetUpdateTimestamp() string`

GetUpdateTimestamp returns the UpdateTimestamp field if non-nil, zero value otherwise.

### GetUpdateTimestampOk

`func (o *ApplicationInstanceMetadata) GetUpdateTimestampOk() (*string, bool)`

GetUpdateTimestampOk returns a tuple with the UpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTimestamp

`func (o *ApplicationInstanceMetadata) SetUpdateTimestamp(v string)`

SetUpdateTimestamp sets UpdateTimestamp field to given value.

### HasUpdateTimestamp

`func (o *ApplicationInstanceMetadata) HasUpdateTimestamp() bool`

HasUpdateTimestamp returns a boolean if a field has been set.

### GetLabels

`func (o *ApplicationInstanceMetadata) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ApplicationInstanceMetadata) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ApplicationInstanceMetadata) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ApplicationInstanceMetadata) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetAnnotations

`func (o *ApplicationInstanceMetadata) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *ApplicationInstanceMetadata) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *ApplicationInstanceMetadata) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *ApplicationInstanceMetadata) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


