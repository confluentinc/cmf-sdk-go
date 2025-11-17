# FlinkApplicationEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** |  | 
**Kind** | **string** |  | 
**Metadata** | [**EventMetadata**](EventMetadata.md) |  | 
**Status** | [**EventStatus**](EventStatus.md) |  | 

## Methods

### NewFlinkApplicationEvent

`func NewFlinkApplicationEvent(apiVersion string, kind string, metadata EventMetadata, status EventStatus, ) *FlinkApplicationEvent`

NewFlinkApplicationEvent instantiates a new FlinkApplicationEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlinkApplicationEventWithDefaults

`func NewFlinkApplicationEventWithDefaults() *FlinkApplicationEvent`

NewFlinkApplicationEventWithDefaults instantiates a new FlinkApplicationEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *FlinkApplicationEvent) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *FlinkApplicationEvent) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *FlinkApplicationEvent) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *FlinkApplicationEvent) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *FlinkApplicationEvent) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *FlinkApplicationEvent) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *FlinkApplicationEvent) GetMetadata() EventMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FlinkApplicationEvent) GetMetadataOk() (*EventMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FlinkApplicationEvent) SetMetadata(v EventMetadata)`

SetMetadata sets Metadata field to given value.


### GetStatus

`func (o *FlinkApplicationEvent) GetStatus() EventStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FlinkApplicationEvent) GetStatusOk() (*EventStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FlinkApplicationEvent) SetStatus(v EventStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


