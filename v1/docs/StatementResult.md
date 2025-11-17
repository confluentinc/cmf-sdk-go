# StatementResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | API version for spec | 
**Kind** | **string** | Kind of resource - set to resource type | 
**Metadata** | [**StatementResultMetadata**](StatementResultMetadata.md) |  | 
**Results** | [**StatementResults**](StatementResults.md) |  | 

## Methods

### NewStatementResult

`func NewStatementResult(apiVersion string, kind string, metadata StatementResultMetadata, results StatementResults, ) *StatementResult`

NewStatementResult instantiates a new StatementResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatementResultWithDefaults

`func NewStatementResultWithDefaults() *StatementResult`

NewStatementResultWithDefaults instantiates a new StatementResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *StatementResult) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *StatementResult) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *StatementResult) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *StatementResult) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *StatementResult) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *StatementResult) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *StatementResult) GetMetadata() StatementResultMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *StatementResult) GetMetadataOk() (*StatementResultMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *StatementResult) SetMetadata(v StatementResultMetadata)`

SetMetadata sets Metadata field to given value.


### GetResults

`func (o *StatementResult) GetResults() StatementResults`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *StatementResult) GetResultsOk() (*StatementResults, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *StatementResult) SetResults(v StatementResults)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


