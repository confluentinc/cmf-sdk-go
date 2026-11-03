# EnvironmentCatalogModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the model. | 
**DatabaseName** | **string** | Name of the database the model belongs to. | 
**CreationTimestamp** | Pointer to **string** | Timestamp when the model was created. | [optional] 
**Options** | Pointer to **map[string]string** | The model&#39;s options from its CREATE MODEL WITH clause. Only options known not to contain credentials, such as provider, endpoint and model, are returned as written; other values are masked. In the endpoint, any text before an &#39;@&#39; (user info) and any query or fragment are masked. For a model that takes its credentials from a secret mapping, the secret&#39;s data is never included; credentials-secret-id shows only the mapping&#39;s name. | [optional] 

## Methods

### NewEnvironmentCatalogModel

`func NewEnvironmentCatalogModel(name string, databaseName string, ) *EnvironmentCatalogModel`

NewEnvironmentCatalogModel instantiates a new EnvironmentCatalogModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentCatalogModelWithDefaults

`func NewEnvironmentCatalogModelWithDefaults() *EnvironmentCatalogModel`

NewEnvironmentCatalogModelWithDefaults instantiates a new EnvironmentCatalogModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EnvironmentCatalogModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentCatalogModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentCatalogModel) SetName(v string)`

SetName sets Name field to given value.


### GetDatabaseName

`func (o *EnvironmentCatalogModel) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *EnvironmentCatalogModel) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *EnvironmentCatalogModel) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.


### GetCreationTimestamp

`func (o *EnvironmentCatalogModel) GetCreationTimestamp() string`

GetCreationTimestamp returns the CreationTimestamp field if non-nil, zero value otherwise.

### GetCreationTimestampOk

`func (o *EnvironmentCatalogModel) GetCreationTimestampOk() (*string, bool)`

GetCreationTimestampOk returns a tuple with the CreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimestamp

`func (o *EnvironmentCatalogModel) SetCreationTimestamp(v string)`

SetCreationTimestamp sets CreationTimestamp field to given value.

### HasCreationTimestamp

`func (o *EnvironmentCatalogModel) HasCreationTimestamp() bool`

HasCreationTimestamp returns a boolean if a field has been set.

### GetOptions

`func (o *EnvironmentCatalogModel) GetOptions() map[string]string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *EnvironmentCatalogModel) GetOptionsOk() (*map[string]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *EnvironmentCatalogModel) SetOptions(v map[string]string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *EnvironmentCatalogModel) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


