# SavepointSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | Path of the Savepoint | [optional] 
**BackoffLimit** | Pointer to **int32** | Backoff limit for the Savepoint | [optional] [default to -1]
**FormatType** | Pointer to **string** | Format type of the Savepoint | [optional] [default to "CANONICAL"]

## Methods

### NewSavepointSpec

`func NewSavepointSpec() *SavepointSpec`

NewSavepointSpec instantiates a new SavepointSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSavepointSpecWithDefaults

`func NewSavepointSpecWithDefaults() *SavepointSpec`

NewSavepointSpecWithDefaults instantiates a new SavepointSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *SavepointSpec) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SavepointSpec) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SavepointSpec) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SavepointSpec) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetBackoffLimit

`func (o *SavepointSpec) GetBackoffLimit() int32`

GetBackoffLimit returns the BackoffLimit field if non-nil, zero value otherwise.

### GetBackoffLimitOk

`func (o *SavepointSpec) GetBackoffLimitOk() (*int32, bool)`

GetBackoffLimitOk returns a tuple with the BackoffLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackoffLimit

`func (o *SavepointSpec) SetBackoffLimit(v int32)`

SetBackoffLimit sets BackoffLimit field to given value.

### HasBackoffLimit

`func (o *SavepointSpec) HasBackoffLimit() bool`

HasBackoffLimit returns a boolean if a field has been set.

### GetFormatType

`func (o *SavepointSpec) GetFormatType() string`

GetFormatType returns the FormatType field if non-nil, zero value otherwise.

### GetFormatTypeOk

`func (o *SavepointSpec) GetFormatTypeOk() (*string, bool)`

GetFormatTypeOk returns a tuple with the FormatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatType

`func (o *SavepointSpec) SetFormatType(v string)`

SetFormatType sets FormatType field to given value.

### HasFormatType

`func (o *SavepointSpec) HasFormatType() bool`

HasFormatType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


