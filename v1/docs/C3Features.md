# C3Features

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** |  | 
**Kind** | **string** |  | 
**Metadata** | [**C3FeaturesMetadata**](C3FeaturesMetadata.md) |  | 
**Status** | [**C3FeaturesStatus**](C3FeaturesStatus.md) |  | 

## Methods

### NewC3Features

`func NewC3Features(apiVersion string, kind string, metadata C3FeaturesMetadata, status C3FeaturesStatus, ) *C3Features`

NewC3Features instantiates a new C3Features object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC3FeaturesWithDefaults

`func NewC3FeaturesWithDefaults() *C3Features`

NewC3FeaturesWithDefaults instantiates a new C3Features object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *C3Features) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *C3Features) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *C3Features) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *C3Features) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *C3Features) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *C3Features) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *C3Features) GetMetadata() C3FeaturesMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *C3Features) GetMetadataOk() (*C3FeaturesMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *C3Features) SetMetadata(v C3FeaturesMetadata)`

SetMetadata sets Metadata field to given value.


### GetStatus

`func (o *C3Features) GetStatus() C3FeaturesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *C3Features) GetStatusOk() (*C3FeaturesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *C3Features) SetStatus(v C3FeaturesStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


