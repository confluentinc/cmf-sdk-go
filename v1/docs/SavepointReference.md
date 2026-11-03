# SavepointReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the savepoint. | [optional] 
**Application** | Pointer to **string** | Name of the application the savepoint was taken of, if it was taken of one. | [optional] 
**Statement** | Pointer to **string** | Name of the statement the savepoint was taken of, if it was taken of one. | [optional] 
**Path** | Pointer to **string** | Path of the savepoint, if it has one. | [optional] 

## Methods

### NewSavepointReference

`func NewSavepointReference() *SavepointReference`

NewSavepointReference instantiates a new SavepointReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSavepointReferenceWithDefaults

`func NewSavepointReferenceWithDefaults() *SavepointReference`

NewSavepointReferenceWithDefaults instantiates a new SavepointReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SavepointReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SavepointReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SavepointReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SavepointReference) HasName() bool`

HasName returns a boolean if a field has been set.

### GetApplication

`func (o *SavepointReference) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *SavepointReference) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *SavepointReference) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *SavepointReference) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetStatement

`func (o *SavepointReference) GetStatement() string`

GetStatement returns the Statement field if non-nil, zero value otherwise.

### GetStatementOk

`func (o *SavepointReference) GetStatementOk() (*string, bool)`

GetStatementOk returns a tuple with the Statement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatement

`func (o *SavepointReference) SetStatement(v string)`

SetStatement sets Statement field to given value.

### HasStatement

`func (o *SavepointReference) HasStatement() bool`

HasStatement returns a boolean if a field has been set.

### GetPath

`func (o *SavepointReference) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SavepointReference) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SavepointReference) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SavepointReference) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


