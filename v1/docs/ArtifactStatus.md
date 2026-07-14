# ArtifactStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **int32** | Version number (starts at 1, auto-incremented on each new binary upload). | [optional] 
**CreationTimestamp** | Pointer to **string** | Timestamp when this version was created (when its binary was uploaded). For version 1 this equals the Artifact&#39;s &#x60;metadata.creationTimestamp&#x60;; later versions each carry their own upload time. | [optional] [readonly] 
**Path** | Pointer to **string** | Full path to the artifact in blob storage. | [optional] 
**Size** | Pointer to **int64** | File size in bytes. Omitted from the response while the version is in &#x60;UPLOADING&#x60; phase. | [optional] 
**Checksum** | Pointer to **string** | SHA-256 checksum of the artifact content (e.g. &#x60;sha256:abc...&#x60;). Omitted from the response while the version is in &#x60;UPLOADING&#x60; phase. | [optional] 
**Phase** | Pointer to **string** | Lifecycle phase of this version. &#x60;UPLOADING&#x60;: content is still being uploaded. &#x60;READY&#x60;: content is uploaded and available. &#x60;FILE_MISSING&#x60;: the backing file is absent from storage. &#x60;STORE_UNREACHABLE&#x60;: artifact storage could not be reached, so the file&#39;s presence could not be verified. | [optional] 
**Message** | Pointer to **string** | Human-readable status message, populated on error states such as &#x60;FILE_MISSING&#x60; and &#x60;STORE_UNREACHABLE&#x60;. | [optional] 

## Methods

### NewArtifactStatus

`func NewArtifactStatus() *ArtifactStatus`

NewArtifactStatus instantiates a new ArtifactStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactStatusWithDefaults

`func NewArtifactStatusWithDefaults() *ArtifactStatus`

NewArtifactStatusWithDefaults instantiates a new ArtifactStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *ArtifactStatus) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ArtifactStatus) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ArtifactStatus) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ArtifactStatus) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCreationTimestamp

`func (o *ArtifactStatus) GetCreationTimestamp() string`

GetCreationTimestamp returns the CreationTimestamp field if non-nil, zero value otherwise.

### GetCreationTimestampOk

`func (o *ArtifactStatus) GetCreationTimestampOk() (*string, bool)`

GetCreationTimestampOk returns a tuple with the CreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimestamp

`func (o *ArtifactStatus) SetCreationTimestamp(v string)`

SetCreationTimestamp sets CreationTimestamp field to given value.

### HasCreationTimestamp

`func (o *ArtifactStatus) HasCreationTimestamp() bool`

HasCreationTimestamp returns a boolean if a field has been set.

### GetPath

`func (o *ArtifactStatus) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ArtifactStatus) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ArtifactStatus) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ArtifactStatus) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetSize

`func (o *ArtifactStatus) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ArtifactStatus) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ArtifactStatus) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *ArtifactStatus) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetChecksum

`func (o *ArtifactStatus) GetChecksum() string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *ArtifactStatus) GetChecksumOk() (*string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *ArtifactStatus) SetChecksum(v string)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *ArtifactStatus) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### GetPhase

`func (o *ArtifactStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *ArtifactStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *ArtifactStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *ArtifactStatus) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetMessage

`func (o *ArtifactStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ArtifactStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ArtifactStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ArtifactStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


