# ResourceEventMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the Event | [optional] 
**Uid** | Pointer to **string** | Unique identifier of the Event. Identical to name. | [optional] 
**ResourceKind** | Pointer to **string** | Kind of the owning resource this event is about, for example Statement, ComputePool, or FlinkApplication. | [optional] 
**CreationTimestamp** | Pointer to **string** | Timestamp when the Event was created | [optional] 
**UpdateTimestamp** | Pointer to **time.Time** | Timestamp of the most recent occurrence of this event | [optional] 
**FlinkApplicationInstance** | Pointer to **string** | UID of the FlinkApplicationInstance this event is related to. Set only for FlinkApplication events, and absent on events recorded before the application was first deployed. | [optional] 

## Methods

### NewResourceEventMetadata

`func NewResourceEventMetadata() *ResourceEventMetadata`

NewResourceEventMetadata instantiates a new ResourceEventMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceEventMetadataWithDefaults

`func NewResourceEventMetadataWithDefaults() *ResourceEventMetadata`

NewResourceEventMetadataWithDefaults instantiates a new ResourceEventMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ResourceEventMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceEventMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceEventMetadata) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResourceEventMetadata) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUid

`func (o *ResourceEventMetadata) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ResourceEventMetadata) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ResourceEventMetadata) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *ResourceEventMetadata) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetResourceKind

`func (o *ResourceEventMetadata) GetResourceKind() string`

GetResourceKind returns the ResourceKind field if non-nil, zero value otherwise.

### GetResourceKindOk

`func (o *ResourceEventMetadata) GetResourceKindOk() (*string, bool)`

GetResourceKindOk returns a tuple with the ResourceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceKind

`func (o *ResourceEventMetadata) SetResourceKind(v string)`

SetResourceKind sets ResourceKind field to given value.

### HasResourceKind

`func (o *ResourceEventMetadata) HasResourceKind() bool`

HasResourceKind returns a boolean if a field has been set.

### GetCreationTimestamp

`func (o *ResourceEventMetadata) GetCreationTimestamp() string`

GetCreationTimestamp returns the CreationTimestamp field if non-nil, zero value otherwise.

### GetCreationTimestampOk

`func (o *ResourceEventMetadata) GetCreationTimestampOk() (*string, bool)`

GetCreationTimestampOk returns a tuple with the CreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimestamp

`func (o *ResourceEventMetadata) SetCreationTimestamp(v string)`

SetCreationTimestamp sets CreationTimestamp field to given value.

### HasCreationTimestamp

`func (o *ResourceEventMetadata) HasCreationTimestamp() bool`

HasCreationTimestamp returns a boolean if a field has been set.

### GetUpdateTimestamp

`func (o *ResourceEventMetadata) GetUpdateTimestamp() time.Time`

GetUpdateTimestamp returns the UpdateTimestamp field if non-nil, zero value otherwise.

### GetUpdateTimestampOk

`func (o *ResourceEventMetadata) GetUpdateTimestampOk() (*time.Time, bool)`

GetUpdateTimestampOk returns a tuple with the UpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTimestamp

`func (o *ResourceEventMetadata) SetUpdateTimestamp(v time.Time)`

SetUpdateTimestamp sets UpdateTimestamp field to given value.

### HasUpdateTimestamp

`func (o *ResourceEventMetadata) HasUpdateTimestamp() bool`

HasUpdateTimestamp returns a boolean if a field has been set.

### GetFlinkApplicationInstance

`func (o *ResourceEventMetadata) GetFlinkApplicationInstance() string`

GetFlinkApplicationInstance returns the FlinkApplicationInstance field if non-nil, zero value otherwise.

### GetFlinkApplicationInstanceOk

`func (o *ResourceEventMetadata) GetFlinkApplicationInstanceOk() (*string, bool)`

GetFlinkApplicationInstanceOk returns a tuple with the FlinkApplicationInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlinkApplicationInstance

`func (o *ResourceEventMetadata) SetFlinkApplicationInstance(v string)`

SetFlinkApplicationInstance sets FlinkApplicationInstance field to given value.

### HasFlinkApplicationInstance

`func (o *ResourceEventMetadata) HasFlinkApplicationInstance() bool`

HasFlinkApplicationInstance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


