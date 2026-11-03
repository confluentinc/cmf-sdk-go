# KubernetesPodLogs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** |  | 
**Kind** | **string** |  | 
**Metadata** | [**KubernetesPodLogsMetadata**](KubernetesPodLogsMetadata.md) |  | 
**Status** | [**KubernetesPodLogsStatus**](KubernetesPodLogsStatus.md) |  | 

## Methods

### NewKubernetesPodLogs

`func NewKubernetesPodLogs(apiVersion string, kind string, metadata KubernetesPodLogsMetadata, status KubernetesPodLogsStatus, ) *KubernetesPodLogs`

NewKubernetesPodLogs instantiates a new KubernetesPodLogs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesPodLogsWithDefaults

`func NewKubernetesPodLogsWithDefaults() *KubernetesPodLogs`

NewKubernetesPodLogsWithDefaults instantiates a new KubernetesPodLogs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *KubernetesPodLogs) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *KubernetesPodLogs) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *KubernetesPodLogs) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *KubernetesPodLogs) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *KubernetesPodLogs) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *KubernetesPodLogs) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *KubernetesPodLogs) GetMetadata() KubernetesPodLogsMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *KubernetesPodLogs) GetMetadataOk() (*KubernetesPodLogsMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *KubernetesPodLogs) SetMetadata(v KubernetesPodLogsMetadata)`

SetMetadata sets Metadata field to given value.


### GetStatus

`func (o *KubernetesPodLogs) GetStatus() KubernetesPodLogsStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KubernetesPodLogs) GetStatusOk() (*KubernetesPodLogsStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KubernetesPodLogs) SetStatus(v KubernetesPodLogsStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


