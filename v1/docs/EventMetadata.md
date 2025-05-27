# EventMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the Event | [optional] 
**Uid** | Pointer to **string** | Unique identifier of the Event. Identical to name. | [optional] 
**CreationTimestamp** | Pointer to **string** | Timestamp when the Event was created | [optional] 
**FlinkApplicationInstance** | Pointer to **string** | Name of the FlinkApplicationInstance which this event is related to | [optional] 
**Labels** | Pointer to **map[string]string** | Labels of the Event | [optional] 
**Annotations** | Pointer to **map[string]string** | Annotations of the Event | [optional] 

## Methods

### NewEventMetadata

`func NewEventMetadata() *EventMetadata`

NewEventMetadata instantiates a new EventMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventMetadataWithDefaults

`func NewEventMetadataWithDefaults() *EventMetadata`

NewEventMetadataWithDefaults instantiates a new EventMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EventMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventMetadata) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EventMetadata) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUid

`func (o *EventMetadata) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *EventMetadata) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *EventMetadata) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *EventMetadata) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetCreationTimestamp

`func (o *EventMetadata) GetCreationTimestamp() string`

GetCreationTimestamp returns the CreationTimestamp field if non-nil, zero value otherwise.

### GetCreationTimestampOk

`func (o *EventMetadata) GetCreationTimestampOk() (*string, bool)`

GetCreationTimestampOk returns a tuple with the CreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimestamp

`func (o *EventMetadata) SetCreationTimestamp(v string)`

SetCreationTimestamp sets CreationTimestamp field to given value.

### HasCreationTimestamp

`func (o *EventMetadata) HasCreationTimestamp() bool`

HasCreationTimestamp returns a boolean if a field has been set.

### GetFlinkApplicationInstance

`func (o *EventMetadata) GetFlinkApplicationInstance() string`

GetFlinkApplicationInstance returns the FlinkApplicationInstance field if non-nil, zero value otherwise.

### GetFlinkApplicationInstanceOk

`func (o *EventMetadata) GetFlinkApplicationInstanceOk() (*string, bool)`

GetFlinkApplicationInstanceOk returns a tuple with the FlinkApplicationInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlinkApplicationInstance

`func (o *EventMetadata) SetFlinkApplicationInstance(v string)`

SetFlinkApplicationInstance sets FlinkApplicationInstance field to given value.

### HasFlinkApplicationInstance

`func (o *EventMetadata) HasFlinkApplicationInstance() bool`

HasFlinkApplicationInstance returns a boolean if a field has been set.

### GetLabels

`func (o *EventMetadata) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *EventMetadata) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *EventMetadata) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *EventMetadata) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetAnnotations

`func (o *EventMetadata) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *EventMetadata) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *EventMetadata) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *EventMetadata) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


