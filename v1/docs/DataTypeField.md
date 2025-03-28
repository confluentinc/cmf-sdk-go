# DataTypeField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the field | 
**FieldType** | [**DataType**](DataType.md) |  | 
**Description** | Pointer to **string** | Description of the field | [optional] 

## Methods

### NewDataTypeField

`func NewDataTypeField(name string, fieldType DataType, ) *DataTypeField`

NewDataTypeField instantiates a new DataTypeField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataTypeFieldWithDefaults

`func NewDataTypeFieldWithDefaults() *DataTypeField`

NewDataTypeFieldWithDefaults instantiates a new DataTypeField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DataTypeField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataTypeField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataTypeField) SetName(v string)`

SetName sets Name field to given value.


### GetFieldType

`func (o *DataTypeField) GetFieldType() DataType`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *DataTypeField) GetFieldTypeOk() (*DataType, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *DataTypeField) SetFieldType(v DataType)`

SetFieldType sets FieldType field to given value.


### GetDescription

`func (o *DataTypeField) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DataTypeField) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DataTypeField) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DataTypeField) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


