# FlinkConfigOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The configuration key. A part in angle brackets, such as &lt;name&gt;, stands for a part of the key that the user chooses. | 
**Type** | **string** | The type of the value as Flink&#39;s documentation names it, for example Boolean, Duration, Enum or Map. | 
**DefaultValue** | Pointer to **string** | The default value as Flink&#39;s documentation shows it, for example (none) or 1 s. | [optional] 
**Category** | Pointer to **string** | The documentation section of the option, for example execution_config for the table.exec options. | [optional] 
**Scope** | Pointer to **[]string** | Whether the option applies to Batch or Streaming execution, or both. Absent when the documentation does not say. | [optional] 
**DescriptionHtml** | Pointer to **string** | The description from Flink&#39;s documentation, as HTML. Render it sanitized. | [optional] 

## Methods

### NewFlinkConfigOption

`func NewFlinkConfigOption(key string, type_ string, ) *FlinkConfigOption`

NewFlinkConfigOption instantiates a new FlinkConfigOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlinkConfigOptionWithDefaults

`func NewFlinkConfigOptionWithDefaults() *FlinkConfigOption`

NewFlinkConfigOptionWithDefaults instantiates a new FlinkConfigOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *FlinkConfigOption) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FlinkConfigOption) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FlinkConfigOption) SetKey(v string)`

SetKey sets Key field to given value.


### GetType

`func (o *FlinkConfigOption) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FlinkConfigOption) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FlinkConfigOption) SetType(v string)`

SetType sets Type field to given value.


### GetDefaultValue

`func (o *FlinkConfigOption) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *FlinkConfigOption) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *FlinkConfigOption) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *FlinkConfigOption) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetCategory

`func (o *FlinkConfigOption) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *FlinkConfigOption) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *FlinkConfigOption) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *FlinkConfigOption) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetScope

`func (o *FlinkConfigOption) GetScope() []string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *FlinkConfigOption) GetScopeOk() (*[]string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *FlinkConfigOption) SetScope(v []string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *FlinkConfigOption) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetDescriptionHtml

`func (o *FlinkConfigOption) GetDescriptionHtml() string`

GetDescriptionHtml returns the DescriptionHtml field if non-nil, zero value otherwise.

### GetDescriptionHtmlOk

`func (o *FlinkConfigOption) GetDescriptionHtmlOk() (*string, bool)`

GetDescriptionHtmlOk returns a tuple with the DescriptionHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionHtml

`func (o *FlinkConfigOption) SetDescriptionHtml(v string)`

SetDescriptionHtml sets DescriptionHtml field to given value.

### HasDescriptionHtml

`func (o *FlinkConfigOption) HasDescriptionHtml() bool`

HasDescriptionHtml returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


