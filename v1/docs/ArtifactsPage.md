# ArtifactsPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pageable** | Pointer to [**Pageable**](Pageable.md) |  | [optional] 
**Metadata** | Pointer to [**ArtifactPageMetadata**](ArtifactPageMetadata.md) |  | [optional] 
**Items** | Pointer to [**[]Artifact**](Artifact.md) |  | [optional] [default to []]

## Methods

### NewArtifactsPage

`func NewArtifactsPage() *ArtifactsPage`

NewArtifactsPage instantiates a new ArtifactsPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactsPageWithDefaults

`func NewArtifactsPageWithDefaults() *ArtifactsPage`

NewArtifactsPageWithDefaults instantiates a new ArtifactsPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageable

`func (o *ArtifactsPage) GetPageable() Pageable`

GetPageable returns the Pageable field if non-nil, zero value otherwise.

### GetPageableOk

`func (o *ArtifactsPage) GetPageableOk() (*Pageable, bool)`

GetPageableOk returns a tuple with the Pageable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageable

`func (o *ArtifactsPage) SetPageable(v Pageable)`

SetPageable sets Pageable field to given value.

### HasPageable

`func (o *ArtifactsPage) HasPageable() bool`

HasPageable returns a boolean if a field has been set.

### GetMetadata

`func (o *ArtifactsPage) GetMetadata() ArtifactPageMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ArtifactsPage) GetMetadataOk() (*ArtifactPageMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ArtifactsPage) SetMetadata(v ArtifactPageMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ArtifactsPage) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetItems

`func (o *ArtifactsPage) GetItems() []Artifact`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ArtifactsPage) GetItemsOk() (*[]Artifact, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ArtifactsPage) SetItems(v []Artifact)`

SetItems sets Items field to given value.

### HasItems

`func (o *ArtifactsPage) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


