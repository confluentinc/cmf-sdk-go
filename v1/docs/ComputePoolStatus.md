# ComputePoolStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phase** | **string** | Phase of the ComputePool | 
**ResourceSummary** | Pointer to [**ResourceUsageSummary**](ResourceUsageSummary.md) |  | [optional] 

## Methods

### NewComputePoolStatus

`func NewComputePoolStatus(phase string, ) *ComputePoolStatus`

NewComputePoolStatus instantiates a new ComputePoolStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputePoolStatusWithDefaults

`func NewComputePoolStatusWithDefaults() *ComputePoolStatus`

NewComputePoolStatusWithDefaults instantiates a new ComputePoolStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhase

`func (o *ComputePoolStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *ComputePoolStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *ComputePoolStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.


### GetResourceSummary

`func (o *ComputePoolStatus) GetResourceSummary() ResourceUsageSummary`

GetResourceSummary returns the ResourceSummary field if non-nil, zero value otherwise.

### GetResourceSummaryOk

`func (o *ComputePoolStatus) GetResourceSummaryOk() (*ResourceUsageSummary, bool)`

GetResourceSummaryOk returns a tuple with the ResourceSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSummary

`func (o *ComputePoolStatus) SetResourceSummary(v ResourceUsageSummary)`

SetResourceSummary sets ResourceSummary field to given value.

### HasResourceSummary

`func (o *ComputePoolStatus) HasResourceSummary() bool`

HasResourceSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


