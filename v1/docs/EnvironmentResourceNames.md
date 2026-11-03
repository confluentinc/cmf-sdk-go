# EnvironmentResourceNames

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applications** | Pointer to **[]string** | Names of the applications. | [optional] 
**Statements** | Pointer to **[]string** | Names of the statements. | [optional] 
**ComputePools** | Pointer to **[]string** | Names of the compute pools. | [optional] 
**Savepoints** | Pointer to [**[]SavepointReference**](SavepointReference.md) | Savepoints of the applications and statements. | [optional] 
**SecretMappings** | Pointer to **[]string** | Names of the secret mappings. | [optional] 

## Methods

### NewEnvironmentResourceNames

`func NewEnvironmentResourceNames() *EnvironmentResourceNames`

NewEnvironmentResourceNames instantiates a new EnvironmentResourceNames object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentResourceNamesWithDefaults

`func NewEnvironmentResourceNamesWithDefaults() *EnvironmentResourceNames`

NewEnvironmentResourceNamesWithDefaults instantiates a new EnvironmentResourceNames object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplications

`func (o *EnvironmentResourceNames) GetApplications() []string`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *EnvironmentResourceNames) GetApplicationsOk() (*[]string, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *EnvironmentResourceNames) SetApplications(v []string)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *EnvironmentResourceNames) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetStatements

`func (o *EnvironmentResourceNames) GetStatements() []string`

GetStatements returns the Statements field if non-nil, zero value otherwise.

### GetStatementsOk

`func (o *EnvironmentResourceNames) GetStatementsOk() (*[]string, bool)`

GetStatementsOk returns a tuple with the Statements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatements

`func (o *EnvironmentResourceNames) SetStatements(v []string)`

SetStatements sets Statements field to given value.

### HasStatements

`func (o *EnvironmentResourceNames) HasStatements() bool`

HasStatements returns a boolean if a field has been set.

### GetComputePools

`func (o *EnvironmentResourceNames) GetComputePools() []string`

GetComputePools returns the ComputePools field if non-nil, zero value otherwise.

### GetComputePoolsOk

`func (o *EnvironmentResourceNames) GetComputePoolsOk() (*[]string, bool)`

GetComputePoolsOk returns a tuple with the ComputePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputePools

`func (o *EnvironmentResourceNames) SetComputePools(v []string)`

SetComputePools sets ComputePools field to given value.

### HasComputePools

`func (o *EnvironmentResourceNames) HasComputePools() bool`

HasComputePools returns a boolean if a field has been set.

### GetSavepoints

`func (o *EnvironmentResourceNames) GetSavepoints() []SavepointReference`

GetSavepoints returns the Savepoints field if non-nil, zero value otherwise.

### GetSavepointsOk

`func (o *EnvironmentResourceNames) GetSavepointsOk() (*[]SavepointReference, bool)`

GetSavepointsOk returns a tuple with the Savepoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavepoints

`func (o *EnvironmentResourceNames) SetSavepoints(v []SavepointReference)`

SetSavepoints sets Savepoints field to given value.

### HasSavepoints

`func (o *EnvironmentResourceNames) HasSavepoints() bool`

HasSavepoints returns a boolean if a field has been set.

### GetSecretMappings

`func (o *EnvironmentResourceNames) GetSecretMappings() []string`

GetSecretMappings returns the SecretMappings field if non-nil, zero value otherwise.

### GetSecretMappingsOk

`func (o *EnvironmentResourceNames) GetSecretMappingsOk() (*[]string, bool)`

GetSecretMappingsOk returns a tuple with the SecretMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretMappings

`func (o *EnvironmentResourceNames) SetSecretMappings(v []string)`

SetSecretMappings sets SecretMappings field to given value.

### HasSecretMappings

`func (o *EnvironmentResourceNames) HasSecretMappings() bool`

HasSecretMappings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


