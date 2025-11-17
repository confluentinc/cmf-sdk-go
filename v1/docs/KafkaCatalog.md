# KafkaCatalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | API version for spec | 
**Kind** | **string** | Kind of resource - set to resource type | 
**Metadata** | [**CatalogMetadata**](CatalogMetadata.md) |  | 
**Spec** | [**KafkaCatalogSpec**](KafkaCatalogSpec.md) |  | 

## Methods

### NewKafkaCatalog

`func NewKafkaCatalog(apiVersion string, kind string, metadata CatalogMetadata, spec KafkaCatalogSpec, ) *KafkaCatalog`

NewKafkaCatalog instantiates a new KafkaCatalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaCatalogWithDefaults

`func NewKafkaCatalogWithDefaults() *KafkaCatalog`

NewKafkaCatalogWithDefaults instantiates a new KafkaCatalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *KafkaCatalog) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *KafkaCatalog) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *KafkaCatalog) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *KafkaCatalog) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *KafkaCatalog) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *KafkaCatalog) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *KafkaCatalog) GetMetadata() CatalogMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *KafkaCatalog) GetMetadataOk() (*CatalogMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *KafkaCatalog) SetMetadata(v CatalogMetadata)`

SetMetadata sets Metadata field to given value.


### GetSpec

`func (o *KafkaCatalog) GetSpec() KafkaCatalogSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *KafkaCatalog) GetSpecOk() (*KafkaCatalogSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *KafkaCatalog) SetSpec(v KafkaCatalogSpec)`

SetSpec sets Spec field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


