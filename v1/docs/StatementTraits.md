# StatementTraits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SqlKind** | Pointer to **string** | The kind of SQL statement | [optional] 
**IsBounded** | Pointer to **bool** | Whether the result of the statement is bounded | [optional] 
**IsAppendOnly** | Pointer to **bool** | Whether the result of the statement is append only | [optional] 
**UpsertColumns** | Pointer to **[]int32** | The column indexes that are updated by the statement | [optional] 
**Schema** | Pointer to [**ResultSchema**](ResultSchema.md) |  | [optional] 

## Methods

### NewStatementTraits

`func NewStatementTraits() *StatementTraits`

NewStatementTraits instantiates a new StatementTraits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatementTraitsWithDefaults

`func NewStatementTraitsWithDefaults() *StatementTraits`

NewStatementTraitsWithDefaults instantiates a new StatementTraits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSqlKind

`func (o *StatementTraits) GetSqlKind() string`

GetSqlKind returns the SqlKind field if non-nil, zero value otherwise.

### GetSqlKindOk

`func (o *StatementTraits) GetSqlKindOk() (*string, bool)`

GetSqlKindOk returns a tuple with the SqlKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlKind

`func (o *StatementTraits) SetSqlKind(v string)`

SetSqlKind sets SqlKind field to given value.

### HasSqlKind

`func (o *StatementTraits) HasSqlKind() bool`

HasSqlKind returns a boolean if a field has been set.

### GetIsBounded

`func (o *StatementTraits) GetIsBounded() bool`

GetIsBounded returns the IsBounded field if non-nil, zero value otherwise.

### GetIsBoundedOk

`func (o *StatementTraits) GetIsBoundedOk() (*bool, bool)`

GetIsBoundedOk returns a tuple with the IsBounded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBounded

`func (o *StatementTraits) SetIsBounded(v bool)`

SetIsBounded sets IsBounded field to given value.

### HasIsBounded

`func (o *StatementTraits) HasIsBounded() bool`

HasIsBounded returns a boolean if a field has been set.

### GetIsAppendOnly

`func (o *StatementTraits) GetIsAppendOnly() bool`

GetIsAppendOnly returns the IsAppendOnly field if non-nil, zero value otherwise.

### GetIsAppendOnlyOk

`func (o *StatementTraits) GetIsAppendOnlyOk() (*bool, bool)`

GetIsAppendOnlyOk returns a tuple with the IsAppendOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAppendOnly

`func (o *StatementTraits) SetIsAppendOnly(v bool)`

SetIsAppendOnly sets IsAppendOnly field to given value.

### HasIsAppendOnly

`func (o *StatementTraits) HasIsAppendOnly() bool`

HasIsAppendOnly returns a boolean if a field has been set.

### GetUpsertColumns

`func (o *StatementTraits) GetUpsertColumns() []int32`

GetUpsertColumns returns the UpsertColumns field if non-nil, zero value otherwise.

### GetUpsertColumnsOk

`func (o *StatementTraits) GetUpsertColumnsOk() (*[]int32, bool)`

GetUpsertColumnsOk returns a tuple with the UpsertColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpsertColumns

`func (o *StatementTraits) SetUpsertColumns(v []int32)`

SetUpsertColumns sets UpsertColumns field to given value.

### HasUpsertColumns

`func (o *StatementTraits) HasUpsertColumns() bool`

HasUpsertColumns returns a boolean if a field has been set.

### GetSchema

`func (o *StatementTraits) GetSchema() ResultSchema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *StatementTraits) GetSchemaOk() (*ResultSchema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *StatementTraits) SetSchema(v ResultSchema)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *StatementTraits) HasSchema() bool`

HasSchema returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


