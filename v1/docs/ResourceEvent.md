# ResourceEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** |  | 
**Kind** | **string** |  | 
**Metadata** | [**ResourceEventMetadata**](ResourceEventMetadata.md) |  | 
**Status** | [**ResourceEventStatus**](ResourceEventStatus.md) |  | 

## Methods

### NewResourceEvent

`func NewResourceEvent(apiVersion string, kind string, metadata ResourceEventMetadata, status ResourceEventStatus, ) *ResourceEvent`

NewResourceEvent instantiates a new ResourceEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceEventWithDefaults

`func NewResourceEventWithDefaults() *ResourceEvent`

NewResourceEventWithDefaults instantiates a new ResourceEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ResourceEvent) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ResourceEvent) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ResourceEvent) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *ResourceEvent) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ResourceEvent) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ResourceEvent) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ResourceEvent) GetMetadata() ResourceEventMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ResourceEvent) GetMetadataOk() (*ResourceEventMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ResourceEvent) SetMetadata(v ResourceEventMetadata)`

SetMetadata sets Metadata field to given value.


### GetStatus

`func (o *ResourceEvent) GetStatus() ResourceEventStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResourceEvent) GetStatusOk() (*ResourceEventStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResourceEvent) SetStatus(v ResourceEventStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


