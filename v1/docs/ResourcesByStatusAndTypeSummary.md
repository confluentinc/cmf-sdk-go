# ResourcesByStatusAndTypeSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applications** | Pointer to [**ApplicationsByStatusSummary**](ApplicationsByStatusSummary.md) |  | [optional] 
**Statements** | Pointer to [**StatementsByStatusSummary**](StatementsByStatusSummary.md) |  | [optional] 

## Methods

### NewResourcesByStatusAndTypeSummary

`func NewResourcesByStatusAndTypeSummary() *ResourcesByStatusAndTypeSummary`

NewResourcesByStatusAndTypeSummary instantiates a new ResourcesByStatusAndTypeSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcesByStatusAndTypeSummaryWithDefaults

`func NewResourcesByStatusAndTypeSummaryWithDefaults() *ResourcesByStatusAndTypeSummary`

NewResourcesByStatusAndTypeSummaryWithDefaults instantiates a new ResourcesByStatusAndTypeSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplications

`func (o *ResourcesByStatusAndTypeSummary) GetApplications() ApplicationsByStatusSummary`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *ResourcesByStatusAndTypeSummary) GetApplicationsOk() (*ApplicationsByStatusSummary, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *ResourcesByStatusAndTypeSummary) SetApplications(v ApplicationsByStatusSummary)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *ResourcesByStatusAndTypeSummary) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetStatements

`func (o *ResourcesByStatusAndTypeSummary) GetStatements() StatementsByStatusSummary`

GetStatements returns the Statements field if non-nil, zero value otherwise.

### GetStatementsOk

`func (o *ResourcesByStatusAndTypeSummary) GetStatementsOk() (*StatementsByStatusSummary, bool)`

GetStatementsOk returns a tuple with the Statements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatements

`func (o *ResourcesByStatusAndTypeSummary) SetStatements(v StatementsByStatusSummary)`

SetStatements sets Statements field to given value.

### HasStatements

`func (o *ResourcesByStatusAndTypeSummary) HasStatements() bool`

HasStatements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


