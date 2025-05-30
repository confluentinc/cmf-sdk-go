# StatementExceptionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | API version for spec | 
**Kind** | **string** | Kind of resource - set to resource type | 
**Data** | [**[]StatementException**](StatementException.md) | List of exceptions | 

## Methods

### NewStatementExceptionList

`func NewStatementExceptionList(apiVersion string, kind string, data []StatementException, ) *StatementExceptionList`

NewStatementExceptionList instantiates a new StatementExceptionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatementExceptionListWithDefaults

`func NewStatementExceptionListWithDefaults() *StatementExceptionList`

NewStatementExceptionListWithDefaults instantiates a new StatementExceptionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *StatementExceptionList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *StatementExceptionList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *StatementExceptionList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *StatementExceptionList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *StatementExceptionList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *StatementExceptionList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetData

`func (o *StatementExceptionList) GetData() []StatementException`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *StatementExceptionList) GetDataOk() (*[]StatementException, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *StatementExceptionList) SetData(v []StatementException)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


