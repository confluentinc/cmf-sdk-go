# SystemInformationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | Version of the CMF application | [optional] 
**Revision** | Pointer to **string** | Git commit hash of the CMF application | [optional] 
**CmfClusterId** | Pointer to **NullableString** | Identifier of this CMF instance in its authorization cluster scope. Pass this as the &#x60;cmf&#x60; key in the &#x60;clusters&#x60; object of any Metadata Service RBAC request addressed to this CMF instance (for example the Confluent CLI&#39;s &#x60;--cmf&#x60; flag). Null when authorizationMode is NONE. | [optional] 
**AuthorizationMode** | Pointer to **string** | How this CMF instance resolves RBAC authorization decisions. NONE -- authorization is not configured, or is configured but not enforced (cmf.authorization.enforce: false). CMF -- CMF&#39;s own embedded Metadata Service makes authorization decisions. CP_MDS -- an external Metadata Service makes authorization decisions. MIGRATION -- an external Metadata Service still makes authorization decisions, while CMF&#39;s embedded Metadata Service is also running so role bindings can be prepared ahead of switching to CMF. | [optional] 

## Methods

### NewSystemInformationStatus

`func NewSystemInformationStatus() *SystemInformationStatus`

NewSystemInformationStatus instantiates a new SystemInformationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemInformationStatusWithDefaults

`func NewSystemInformationStatusWithDefaults() *SystemInformationStatus`

NewSystemInformationStatusWithDefaults instantiates a new SystemInformationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *SystemInformationStatus) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SystemInformationStatus) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SystemInformationStatus) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SystemInformationStatus) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetRevision

`func (o *SystemInformationStatus) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *SystemInformationStatus) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *SystemInformationStatus) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *SystemInformationStatus) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetCmfClusterId

`func (o *SystemInformationStatus) GetCmfClusterId() string`

GetCmfClusterId returns the CmfClusterId field if non-nil, zero value otherwise.

### GetCmfClusterIdOk

`func (o *SystemInformationStatus) GetCmfClusterIdOk() (*string, bool)`

GetCmfClusterIdOk returns a tuple with the CmfClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmfClusterId

`func (o *SystemInformationStatus) SetCmfClusterId(v string)`

SetCmfClusterId sets CmfClusterId field to given value.

### HasCmfClusterId

`func (o *SystemInformationStatus) HasCmfClusterId() bool`

HasCmfClusterId returns a boolean if a field has been set.

### SetCmfClusterIdNil

`func (o *SystemInformationStatus) SetCmfClusterIdNil(b bool)`

 SetCmfClusterIdNil sets the value for CmfClusterId to be an explicit nil

### UnsetCmfClusterId
`func (o *SystemInformationStatus) UnsetCmfClusterId()`

UnsetCmfClusterId ensures that no value is present for CmfClusterId, not even an explicit nil
### GetAuthorizationMode

`func (o *SystemInformationStatus) GetAuthorizationMode() string`

GetAuthorizationMode returns the AuthorizationMode field if non-nil, zero value otherwise.

### GetAuthorizationModeOk

`func (o *SystemInformationStatus) GetAuthorizationModeOk() (*string, bool)`

GetAuthorizationModeOk returns a tuple with the AuthorizationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationMode

`func (o *SystemInformationStatus) SetAuthorizationMode(v string)`

SetAuthorizationMode sets AuthorizationMode field to given value.

### HasAuthorizationMode

`func (o *SystemInformationStatus) HasAuthorizationMode() bool`

HasAuthorizationMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


