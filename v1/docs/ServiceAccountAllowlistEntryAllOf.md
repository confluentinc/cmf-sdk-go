# ServiceAccountAllowlistEntryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**ServiceAccountAllowlistEntryMetadata**](ServiceAccountAllowlistEntryMetadata.md) |  | 
**Spec** | **map[string]interface{}** | Reserved for future write-time fields; currently empty. | 
**Status** | Pointer to [**ServiceAccountAllowlistEntryStatus**](ServiceAccountAllowlistEntryStatus.md) |  | [optional] 

## Methods

### NewServiceAccountAllowlistEntryAllOf

`func NewServiceAccountAllowlistEntryAllOf(metadata ServiceAccountAllowlistEntryMetadata, spec map[string]interface{}, ) *ServiceAccountAllowlistEntryAllOf`

NewServiceAccountAllowlistEntryAllOf instantiates a new ServiceAccountAllowlistEntryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccountAllowlistEntryAllOfWithDefaults

`func NewServiceAccountAllowlistEntryAllOfWithDefaults() *ServiceAccountAllowlistEntryAllOf`

NewServiceAccountAllowlistEntryAllOfWithDefaults instantiates a new ServiceAccountAllowlistEntryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *ServiceAccountAllowlistEntryAllOf) GetMetadata() ServiceAccountAllowlistEntryMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ServiceAccountAllowlistEntryAllOf) GetMetadataOk() (*ServiceAccountAllowlistEntryMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ServiceAccountAllowlistEntryAllOf) SetMetadata(v ServiceAccountAllowlistEntryMetadata)`

SetMetadata sets Metadata field to given value.


### GetSpec

`func (o *ServiceAccountAllowlistEntryAllOf) GetSpec() map[string]interface{}`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ServiceAccountAllowlistEntryAllOf) GetSpecOk() (*map[string]interface{}, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ServiceAccountAllowlistEntryAllOf) SetSpec(v map[string]interface{})`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *ServiceAccountAllowlistEntryAllOf) GetStatus() ServiceAccountAllowlistEntryStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServiceAccountAllowlistEntryAllOf) GetStatusOk() (*ServiceAccountAllowlistEntryStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServiceAccountAllowlistEntryAllOf) SetStatus(v ServiceAccountAllowlistEntryStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ServiceAccountAllowlistEntryAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


