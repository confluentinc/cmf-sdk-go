# ServiceAccountAllowlistEntryMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the permitted service account. | 
**CreationTimestamp** | Pointer to **string** | Timestamp at which the entry was added to the allowlist. | [optional] [readonly] 

## Methods

### NewServiceAccountAllowlistEntryMetadata

`func NewServiceAccountAllowlistEntryMetadata(name string, ) *ServiceAccountAllowlistEntryMetadata`

NewServiceAccountAllowlistEntryMetadata instantiates a new ServiceAccountAllowlistEntryMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccountAllowlistEntryMetadataWithDefaults

`func NewServiceAccountAllowlistEntryMetadataWithDefaults() *ServiceAccountAllowlistEntryMetadata`

NewServiceAccountAllowlistEntryMetadataWithDefaults instantiates a new ServiceAccountAllowlistEntryMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ServiceAccountAllowlistEntryMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceAccountAllowlistEntryMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceAccountAllowlistEntryMetadata) SetName(v string)`

SetName sets Name field to given value.


### GetCreationTimestamp

`func (o *ServiceAccountAllowlistEntryMetadata) GetCreationTimestamp() string`

GetCreationTimestamp returns the CreationTimestamp field if non-nil, zero value otherwise.

### GetCreationTimestampOk

`func (o *ServiceAccountAllowlistEntryMetadata) GetCreationTimestampOk() (*string, bool)`

GetCreationTimestampOk returns a tuple with the CreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimestamp

`func (o *ServiceAccountAllowlistEntryMetadata) SetCreationTimestamp(v string)`

SetCreationTimestamp sets CreationTimestamp field to given value.

### HasCreationTimestamp

`func (o *ServiceAccountAllowlistEntryMetadata) HasCreationTimestamp() bool`

HasCreationTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


