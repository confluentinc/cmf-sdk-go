# KafkaCatalogAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**CatalogMetadata**](CatalogMetadata.md) |  | 
**Spec** | [**KafkaCatalogSpec**](KafkaCatalogSpec.md) |  | 
**Status** | Pointer to [**KafkaCatalogStatus**](KafkaCatalogStatus.md) |  | [optional] 

## Methods

### NewKafkaCatalogAllOf

`func NewKafkaCatalogAllOf(metadata CatalogMetadata, spec KafkaCatalogSpec, ) *KafkaCatalogAllOf`

NewKafkaCatalogAllOf instantiates a new KafkaCatalogAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaCatalogAllOfWithDefaults

`func NewKafkaCatalogAllOfWithDefaults() *KafkaCatalogAllOf`

NewKafkaCatalogAllOfWithDefaults instantiates a new KafkaCatalogAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *KafkaCatalogAllOf) GetMetadata() CatalogMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *KafkaCatalogAllOf) GetMetadataOk() (*CatalogMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *KafkaCatalogAllOf) SetMetadata(v CatalogMetadata)`

SetMetadata sets Metadata field to given value.


### GetSpec

`func (o *KafkaCatalogAllOf) GetSpec() KafkaCatalogSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *KafkaCatalogAllOf) GetSpecOk() (*KafkaCatalogSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *KafkaCatalogAllOf) SetSpec(v KafkaCatalogSpec)`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *KafkaCatalogAllOf) GetStatus() KafkaCatalogStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KafkaCatalogAllOf) GetStatusOk() (*KafkaCatalogStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KafkaCatalogAllOf) SetStatus(v KafkaCatalogStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KafkaCatalogAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


