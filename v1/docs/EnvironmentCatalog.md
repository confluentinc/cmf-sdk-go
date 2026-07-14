# EnvironmentCatalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the default catalog. | 
**Databases** | [**[]EnvironmentCatalogDatabase**](EnvironmentCatalogDatabase.md) | Databases currently registered in the default catalog. | 

## Methods

### NewEnvironmentCatalog

`func NewEnvironmentCatalog(name string, databases []EnvironmentCatalogDatabase, ) *EnvironmentCatalog`

NewEnvironmentCatalog instantiates a new EnvironmentCatalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentCatalogWithDefaults

`func NewEnvironmentCatalogWithDefaults() *EnvironmentCatalog`

NewEnvironmentCatalogWithDefaults instantiates a new EnvironmentCatalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EnvironmentCatalog) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentCatalog) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentCatalog) SetName(v string)`

SetName sets Name field to given value.


### GetDatabases

`func (o *EnvironmentCatalog) GetDatabases() []EnvironmentCatalogDatabase`

GetDatabases returns the Databases field if non-nil, zero value otherwise.

### GetDatabasesOk

`func (o *EnvironmentCatalog) GetDatabasesOk() (*[]EnvironmentCatalogDatabase, bool)`

GetDatabasesOk returns a tuple with the Databases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabases

`func (o *EnvironmentCatalog) SetDatabases(v []EnvironmentCatalogDatabase)`

SetDatabases sets Databases field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


