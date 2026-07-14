# ArtifactsPageAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ArtifactPageMetadata**](ArtifactPageMetadata.md) |  | [optional] 
**Items** | Pointer to [**[]Artifact**](Artifact.md) |  | [optional] [default to []]

## Methods

### NewArtifactsPageAllOf

`func NewArtifactsPageAllOf() *ArtifactsPageAllOf`

NewArtifactsPageAllOf instantiates a new ArtifactsPageAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactsPageAllOfWithDefaults

`func NewArtifactsPageAllOfWithDefaults() *ArtifactsPageAllOf`

NewArtifactsPageAllOfWithDefaults instantiates a new ArtifactsPageAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *ArtifactsPageAllOf) GetMetadata() ArtifactPageMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ArtifactsPageAllOf) GetMetadataOk() (*ArtifactPageMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ArtifactsPageAllOf) SetMetadata(v ArtifactPageMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ArtifactsPageAllOf) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetItems

`func (o *ArtifactsPageAllOf) GetItems() []Artifact`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ArtifactsPageAllOf) GetItemsOk() (*[]Artifact, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ArtifactsPageAllOf) SetItems(v []Artifact)`

SetItems sets Items field to given value.

### HasItems

`func (o *ArtifactsPageAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


