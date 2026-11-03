# SqlAssistantCollectBound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxRows** | Pointer to **int32** | Row cap, capped by cmf.sql.assistant.max-result-rows. | [optional] 
**MaxSeconds** | Pointer to **int32** | Time cap in seconds, at most 300. | [optional] 

## Methods

### NewSqlAssistantCollectBound

`func NewSqlAssistantCollectBound() *SqlAssistantCollectBound`

NewSqlAssistantCollectBound instantiates a new SqlAssistantCollectBound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlAssistantCollectBoundWithDefaults

`func NewSqlAssistantCollectBoundWithDefaults() *SqlAssistantCollectBound`

NewSqlAssistantCollectBoundWithDefaults instantiates a new SqlAssistantCollectBound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxRows

`func (o *SqlAssistantCollectBound) GetMaxRows() int32`

GetMaxRows returns the MaxRows field if non-nil, zero value otherwise.

### GetMaxRowsOk

`func (o *SqlAssistantCollectBound) GetMaxRowsOk() (*int32, bool)`

GetMaxRowsOk returns a tuple with the MaxRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRows

`func (o *SqlAssistantCollectBound) SetMaxRows(v int32)`

SetMaxRows sets MaxRows field to given value.

### HasMaxRows

`func (o *SqlAssistantCollectBound) HasMaxRows() bool`

HasMaxRows returns a boolean if a field has been set.

### GetMaxSeconds

`func (o *SqlAssistantCollectBound) GetMaxSeconds() int32`

GetMaxSeconds returns the MaxSeconds field if non-nil, zero value otherwise.

### GetMaxSecondsOk

`func (o *SqlAssistantCollectBound) GetMaxSecondsOk() (*int32, bool)`

GetMaxSecondsOk returns a tuple with the MaxSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSeconds

`func (o *SqlAssistantCollectBound) SetMaxSeconds(v int32)`

SetMaxSeconds sets MaxSeconds field to given value.

### HasMaxSeconds

`func (o *SqlAssistantCollectBound) HasMaxSeconds() bool`

HasMaxSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


