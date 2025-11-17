# ResultSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | [**[]ResultSchemaColumn**](ResultSchemaColumn.md) | Properites of all columns in the schema | 

## Methods

### NewResultSchema

`func NewResultSchema(columns []ResultSchemaColumn, ) *ResultSchema`

NewResultSchema instantiates a new ResultSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultSchemaWithDefaults

`func NewResultSchemaWithDefaults() *ResultSchema`

NewResultSchemaWithDefaults instantiates a new ResultSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *ResultSchema) GetColumns() []ResultSchemaColumn`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *ResultSchema) GetColumnsOk() (*[]ResultSchemaColumn, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *ResultSchema) SetColumns(v []ResultSchemaColumn)`

SetColumns sets Columns field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


