# Statement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | API version for spec | 
**Kind** | **string** | Kind of resource - set to resource type | 
**Metadata** | [**StatementMetadata**](StatementMetadata.md) |  | 
**Spec** | [**StatementSpec**](StatementSpec.md) |  | 
**Status** | Pointer to [**StatementStatus**](StatementStatus.md) |  | [optional] 
**Result** | Pointer to [**StatementResult**](StatementResult.md) |  | [optional] 

## Methods

### NewStatement

`func NewStatement(apiVersion string, kind string, metadata StatementMetadata, spec StatementSpec, ) *Statement`

NewStatement instantiates a new Statement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatementWithDefaults

`func NewStatementWithDefaults() *Statement`

NewStatementWithDefaults instantiates a new Statement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *Statement) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *Statement) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *Statement) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *Statement) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Statement) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Statement) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *Statement) GetMetadata() StatementMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Statement) GetMetadataOk() (*StatementMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Statement) SetMetadata(v StatementMetadata)`

SetMetadata sets Metadata field to given value.


### GetSpec

`func (o *Statement) GetSpec() StatementSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *Statement) GetSpecOk() (*StatementSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *Statement) SetSpec(v StatementSpec)`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *Statement) GetStatus() StatementStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Statement) GetStatusOk() (*StatementStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Statement) SetStatus(v StatementStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Statement) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *Statement) GetResult() StatementResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *Statement) GetResultOk() (*StatementResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *Statement) SetResult(v StatementResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *Statement) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


