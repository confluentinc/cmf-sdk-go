# MetadataSyncStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** | Whether the restore is still running. | 
**StartedAt** | **time.Time** | When the restore began. | 
**CompletedAt** | Pointer to **time.Time** | Absent while the restore is still running. | [optional] 

## Methods

### NewMetadataSyncStatus

`func NewMetadataSyncStatus(state string, startedAt time.Time, ) *MetadataSyncStatus`

NewMetadataSyncStatus instantiates a new MetadataSyncStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataSyncStatusWithDefaults

`func NewMetadataSyncStatusWithDefaults() *MetadataSyncStatus`

NewMetadataSyncStatusWithDefaults instantiates a new MetadataSyncStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *MetadataSyncStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MetadataSyncStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MetadataSyncStatus) SetState(v string)`

SetState sets State field to given value.


### GetStartedAt

`func (o *MetadataSyncStatus) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *MetadataSyncStatus) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *MetadataSyncStatus) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetCompletedAt

`func (o *MetadataSyncStatus) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *MetadataSyncStatus) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *MetadataSyncStatus) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *MetadataSyncStatus) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


