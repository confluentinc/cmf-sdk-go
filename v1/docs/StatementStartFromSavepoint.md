# StatementStartFromSavepoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SavepointName** | Pointer to **string** | The name of the Savepoint resource to start Statement from. The request will be rejected if savepoint that has not completed is referenced. | [optional] 
**Uid** | Pointer to **string** | The uuid of the Savepoint resource to start from. | [optional] 
**InitialSavepointPath** | Pointer to **string** | The path of the savepoint to start the Statement from. This could be an external path too. | [optional] 
**AllowNonRestoredState** | Pointer to **bool** | A boolean flag to allow the job to start even if some state could not be restored. | [optional] 
**SavepointRedeployNonce** | Pointer to **int64** | Nonce used to trigger a full redeployment of the job from the savepoint. In order to trigger redeployment, change the number to a different non-null value. Rollback is not possible after redeployment. | [optional] 

## Methods

### NewStatementStartFromSavepoint

`func NewStatementStartFromSavepoint() *StatementStartFromSavepoint`

NewStatementStartFromSavepoint instantiates a new StatementStartFromSavepoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatementStartFromSavepointWithDefaults

`func NewStatementStartFromSavepointWithDefaults() *StatementStartFromSavepoint`

NewStatementStartFromSavepointWithDefaults instantiates a new StatementStartFromSavepoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSavepointName

`func (o *StatementStartFromSavepoint) GetSavepointName() string`

GetSavepointName returns the SavepointName field if non-nil, zero value otherwise.

### GetSavepointNameOk

`func (o *StatementStartFromSavepoint) GetSavepointNameOk() (*string, bool)`

GetSavepointNameOk returns a tuple with the SavepointName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavepointName

`func (o *StatementStartFromSavepoint) SetSavepointName(v string)`

SetSavepointName sets SavepointName field to given value.

### HasSavepointName

`func (o *StatementStartFromSavepoint) HasSavepointName() bool`

HasSavepointName returns a boolean if a field has been set.

### GetUid

`func (o *StatementStartFromSavepoint) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *StatementStartFromSavepoint) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *StatementStartFromSavepoint) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *StatementStartFromSavepoint) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetInitialSavepointPath

`func (o *StatementStartFromSavepoint) GetInitialSavepointPath() string`

GetInitialSavepointPath returns the InitialSavepointPath field if non-nil, zero value otherwise.

### GetInitialSavepointPathOk

`func (o *StatementStartFromSavepoint) GetInitialSavepointPathOk() (*string, bool)`

GetInitialSavepointPathOk returns a tuple with the InitialSavepointPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialSavepointPath

`func (o *StatementStartFromSavepoint) SetInitialSavepointPath(v string)`

SetInitialSavepointPath sets InitialSavepointPath field to given value.

### HasInitialSavepointPath

`func (o *StatementStartFromSavepoint) HasInitialSavepointPath() bool`

HasInitialSavepointPath returns a boolean if a field has been set.

### GetAllowNonRestoredState

`func (o *StatementStartFromSavepoint) GetAllowNonRestoredState() bool`

GetAllowNonRestoredState returns the AllowNonRestoredState field if non-nil, zero value otherwise.

### GetAllowNonRestoredStateOk

`func (o *StatementStartFromSavepoint) GetAllowNonRestoredStateOk() (*bool, bool)`

GetAllowNonRestoredStateOk returns a tuple with the AllowNonRestoredState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowNonRestoredState

`func (o *StatementStartFromSavepoint) SetAllowNonRestoredState(v bool)`

SetAllowNonRestoredState sets AllowNonRestoredState field to given value.

### HasAllowNonRestoredState

`func (o *StatementStartFromSavepoint) HasAllowNonRestoredState() bool`

HasAllowNonRestoredState returns a boolean if a field has been set.

### GetSavepointRedeployNonce

`func (o *StatementStartFromSavepoint) GetSavepointRedeployNonce() int64`

GetSavepointRedeployNonce returns the SavepointRedeployNonce field if non-nil, zero value otherwise.

### GetSavepointRedeployNonceOk

`func (o *StatementStartFromSavepoint) GetSavepointRedeployNonceOk() (*int64, bool)`

GetSavepointRedeployNonceOk returns a tuple with the SavepointRedeployNonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavepointRedeployNonce

`func (o *StatementStartFromSavepoint) SetSavepointRedeployNonce(v int64)`

SetSavepointRedeployNonce sets SavepointRedeployNonce field to given value.

### HasSavepointRedeployNonce

`func (o *StatementStartFromSavepoint) HasSavepointRedeployNonce() bool`

HasSavepointRedeployNonce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


