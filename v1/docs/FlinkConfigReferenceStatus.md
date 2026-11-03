# FlinkConfigReferenceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlinkVersionLabel** | Pointer to **string** | The Flink version as Flink&#39;s documentation names it, for example 1.19. | [optional] 
**Options** | [**[]FlinkConfigOption**](FlinkConfigOption.md) |  | 

## Methods

### NewFlinkConfigReferenceStatus

`func NewFlinkConfigReferenceStatus(options []FlinkConfigOption, ) *FlinkConfigReferenceStatus`

NewFlinkConfigReferenceStatus instantiates a new FlinkConfigReferenceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlinkConfigReferenceStatusWithDefaults

`func NewFlinkConfigReferenceStatusWithDefaults() *FlinkConfigReferenceStatus`

NewFlinkConfigReferenceStatusWithDefaults instantiates a new FlinkConfigReferenceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlinkVersionLabel

`func (o *FlinkConfigReferenceStatus) GetFlinkVersionLabel() string`

GetFlinkVersionLabel returns the FlinkVersionLabel field if non-nil, zero value otherwise.

### GetFlinkVersionLabelOk

`func (o *FlinkConfigReferenceStatus) GetFlinkVersionLabelOk() (*string, bool)`

GetFlinkVersionLabelOk returns a tuple with the FlinkVersionLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlinkVersionLabel

`func (o *FlinkConfigReferenceStatus) SetFlinkVersionLabel(v string)`

SetFlinkVersionLabel sets FlinkVersionLabel field to given value.

### HasFlinkVersionLabel

`func (o *FlinkConfigReferenceStatus) HasFlinkVersionLabel() bool`

HasFlinkVersionLabel returns a boolean if a field has been set.

### GetOptions

`func (o *FlinkConfigReferenceStatus) GetOptions() []FlinkConfigOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *FlinkConfigReferenceStatus) GetOptionsOk() (*[]FlinkConfigOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *FlinkConfigReferenceStatus) SetOptions(v []FlinkConfigOption)`

SetOptions sets Options field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


