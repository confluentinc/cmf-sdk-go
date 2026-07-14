# SavepointsBulkDeleteResultFailures

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the savepoint that failed to delete. | [optional] 
**Reason** | Pointer to **string** | Human-readable reason the deletion failed. | [optional] 

## Methods

### NewSavepointsBulkDeleteResultFailures

`func NewSavepointsBulkDeleteResultFailures() *SavepointsBulkDeleteResultFailures`

NewSavepointsBulkDeleteResultFailures instantiates a new SavepointsBulkDeleteResultFailures object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSavepointsBulkDeleteResultFailuresWithDefaults

`func NewSavepointsBulkDeleteResultFailuresWithDefaults() *SavepointsBulkDeleteResultFailures`

NewSavepointsBulkDeleteResultFailuresWithDefaults instantiates a new SavepointsBulkDeleteResultFailures object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SavepointsBulkDeleteResultFailures) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SavepointsBulkDeleteResultFailures) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SavepointsBulkDeleteResultFailures) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SavepointsBulkDeleteResultFailures) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReason

`func (o *SavepointsBulkDeleteResultFailures) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SavepointsBulkDeleteResultFailures) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SavepointsBulkDeleteResultFailures) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *SavepointsBulkDeleteResultFailures) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


