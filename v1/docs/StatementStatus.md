# StatementStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phase** | **string** | The lifecycle phase of the statement | 
**Detail** | Pointer to **string** | Details about the execution status of the statement | [optional] 
**Traits** | Pointer to [**StatementTraits**](StatementTraits.md) |  | [optional] 

## Methods

### NewStatementStatus

`func NewStatementStatus(phase string, ) *StatementStatus`

NewStatementStatus instantiates a new StatementStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatementStatusWithDefaults

`func NewStatementStatusWithDefaults() *StatementStatus`

NewStatementStatusWithDefaults instantiates a new StatementStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhase

`func (o *StatementStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *StatementStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *StatementStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.


### GetDetail

`func (o *StatementStatus) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *StatementStatus) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *StatementStatus) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *StatementStatus) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetTraits

`func (o *StatementStatus) GetTraits() StatementTraits`

GetTraits returns the Traits field if non-nil, zero value otherwise.

### GetTraitsOk

`func (o *StatementStatus) GetTraitsOk() (*StatementTraits, bool)`

GetTraitsOk returns a tuple with the Traits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraits

`func (o *StatementStatus) SetTraits(v StatementTraits)`

SetTraits sets Traits field to given value.

### HasTraits

`func (o *StatementStatus) HasTraits() bool`

HasTraits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


