# SavepointsBulkDeleteResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeletedCount** | Pointer to **int32** | Number of savepoints that were successfully deleted. | [optional] 
**FailedCount** | Pointer to **int32** | Number of savepoints that failed to delete. | [optional] 
**Failures** | Pointer to [**[]SavepointsBulkDeleteResultFailures**](SavepointsBulkDeleteResultFailures.md) | Per-savepoint failure details for each savepoint that could not be deleted. | [optional] 

## Methods

### NewSavepointsBulkDeleteResult

`func NewSavepointsBulkDeleteResult() *SavepointsBulkDeleteResult`

NewSavepointsBulkDeleteResult instantiates a new SavepointsBulkDeleteResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSavepointsBulkDeleteResultWithDefaults

`func NewSavepointsBulkDeleteResultWithDefaults() *SavepointsBulkDeleteResult`

NewSavepointsBulkDeleteResultWithDefaults instantiates a new SavepointsBulkDeleteResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletedCount

`func (o *SavepointsBulkDeleteResult) GetDeletedCount() int32`

GetDeletedCount returns the DeletedCount field if non-nil, zero value otherwise.

### GetDeletedCountOk

`func (o *SavepointsBulkDeleteResult) GetDeletedCountOk() (*int32, bool)`

GetDeletedCountOk returns a tuple with the DeletedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedCount

`func (o *SavepointsBulkDeleteResult) SetDeletedCount(v int32)`

SetDeletedCount sets DeletedCount field to given value.

### HasDeletedCount

`func (o *SavepointsBulkDeleteResult) HasDeletedCount() bool`

HasDeletedCount returns a boolean if a field has been set.

### GetFailedCount

`func (o *SavepointsBulkDeleteResult) GetFailedCount() int32`

GetFailedCount returns the FailedCount field if non-nil, zero value otherwise.

### GetFailedCountOk

`func (o *SavepointsBulkDeleteResult) GetFailedCountOk() (*int32, bool)`

GetFailedCountOk returns a tuple with the FailedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedCount

`func (o *SavepointsBulkDeleteResult) SetFailedCount(v int32)`

SetFailedCount sets FailedCount field to given value.

### HasFailedCount

`func (o *SavepointsBulkDeleteResult) HasFailedCount() bool`

HasFailedCount returns a boolean if a field has been set.

### GetFailures

`func (o *SavepointsBulkDeleteResult) GetFailures() []SavepointsBulkDeleteResultFailures`

GetFailures returns the Failures field if non-nil, zero value otherwise.

### GetFailuresOk

`func (o *SavepointsBulkDeleteResult) GetFailuresOk() (*[]SavepointsBulkDeleteResultFailures, bool)`

GetFailuresOk returns a tuple with the Failures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailures

`func (o *SavepointsBulkDeleteResult) SetFailures(v []SavepointsBulkDeleteResultFailures)`

SetFailures sets Failures field to given value.

### HasFailures

`func (o *SavepointsBulkDeleteResult) HasFailures() bool`

HasFailures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


