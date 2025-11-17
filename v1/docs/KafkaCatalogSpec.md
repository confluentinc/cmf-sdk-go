# KafkaCatalogSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SrInstance** | [**KafkaCatalogSpecSrInstance**](KafkaCatalogSpecSrInstance.md) |  | 
**KafkaClusters** | Pointer to [**[]KafkaCatalogSpecKafkaClusters**](KafkaCatalogSpecKafkaClusters.md) |  | [optional] 

## Methods

### NewKafkaCatalogSpec

`func NewKafkaCatalogSpec(srInstance KafkaCatalogSpecSrInstance, ) *KafkaCatalogSpec`

NewKafkaCatalogSpec instantiates a new KafkaCatalogSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaCatalogSpecWithDefaults

`func NewKafkaCatalogSpecWithDefaults() *KafkaCatalogSpec`

NewKafkaCatalogSpecWithDefaults instantiates a new KafkaCatalogSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSrInstance

`func (o *KafkaCatalogSpec) GetSrInstance() KafkaCatalogSpecSrInstance`

GetSrInstance returns the SrInstance field if non-nil, zero value otherwise.

### GetSrInstanceOk

`func (o *KafkaCatalogSpec) GetSrInstanceOk() (*KafkaCatalogSpecSrInstance, bool)`

GetSrInstanceOk returns a tuple with the SrInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrInstance

`func (o *KafkaCatalogSpec) SetSrInstance(v KafkaCatalogSpecSrInstance)`

SetSrInstance sets SrInstance field to given value.


### GetKafkaClusters

`func (o *KafkaCatalogSpec) GetKafkaClusters() []KafkaCatalogSpecKafkaClusters`

GetKafkaClusters returns the KafkaClusters field if non-nil, zero value otherwise.

### GetKafkaClustersOk

`func (o *KafkaCatalogSpec) GetKafkaClustersOk() (*[]KafkaCatalogSpecKafkaClusters, bool)`

GetKafkaClustersOk returns a tuple with the KafkaClusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaClusters

`func (o *KafkaCatalogSpec) SetKafkaClusters(v []KafkaCatalogSpecKafkaClusters)`

SetKafkaClusters sets KafkaClusters field to given value.

### HasKafkaClusters

`func (o *KafkaCatalogSpec) HasKafkaClusters() bool`

HasKafkaClusters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


