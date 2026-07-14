# ArtifactAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**ArtifactMetadata**](ArtifactMetadata.md) |  | 
**Spec** | **map[string]interface{}** | Spec of an Artifact (reserved for future use; empty in V1). | 
**Status** | Pointer to [**ArtifactStatus**](ArtifactStatus.md) |  | [optional] 

## Methods

### NewArtifactAllOf

`func NewArtifactAllOf(metadata ArtifactMetadata, spec map[string]interface{}, ) *ArtifactAllOf`

NewArtifactAllOf instantiates a new ArtifactAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactAllOfWithDefaults

`func NewArtifactAllOfWithDefaults() *ArtifactAllOf`

NewArtifactAllOfWithDefaults instantiates a new ArtifactAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *ArtifactAllOf) GetMetadata() ArtifactMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ArtifactAllOf) GetMetadataOk() (*ArtifactMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ArtifactAllOf) SetMetadata(v ArtifactMetadata)`

SetMetadata sets Metadata field to given value.


### GetSpec

`func (o *ArtifactAllOf) GetSpec() map[string]interface{}`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ArtifactAllOf) GetSpecOk() (*map[string]interface{}, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ArtifactAllOf) SetSpec(v map[string]interface{})`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *ArtifactAllOf) GetStatus() ArtifactStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ArtifactAllOf) GetStatusOk() (*ArtifactStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ArtifactAllOf) SetStatus(v ArtifactStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ArtifactAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


