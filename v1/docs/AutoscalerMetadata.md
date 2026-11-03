# AutoscalerMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdateTimestamp** | Pointer to **time.Time** | Timestamp at which this state was read from the running Application. The autoscaler keeps updating its state afterwards, so a later read can report something different.  | [optional] 

## Methods

### NewAutoscalerMetadata

`func NewAutoscalerMetadata() *AutoscalerMetadata`

NewAutoscalerMetadata instantiates a new AutoscalerMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoscalerMetadataWithDefaults

`func NewAutoscalerMetadataWithDefaults() *AutoscalerMetadata`

NewAutoscalerMetadataWithDefaults instantiates a new AutoscalerMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdateTimestamp

`func (o *AutoscalerMetadata) GetUpdateTimestamp() time.Time`

GetUpdateTimestamp returns the UpdateTimestamp field if non-nil, zero value otherwise.

### GetUpdateTimestampOk

`func (o *AutoscalerMetadata) GetUpdateTimestampOk() (*time.Time, bool)`

GetUpdateTimestampOk returns a tuple with the UpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTimestamp

`func (o *AutoscalerMetadata) SetUpdateTimestamp(v time.Time)`

SetUpdateTimestamp sets UpdateTimestamp field to given value.

### HasUpdateTimestamp

`func (o *AutoscalerMetadata) HasUpdateTimestamp() bool`

HasUpdateTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


