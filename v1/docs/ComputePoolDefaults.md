# ComputePoolDefaults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComputePoolDefaults** | Pointer to **map[string]interface{}** | the defaults as YAML or JSON for ComputePools  Some keys in \&quot;clusterSpec.flinkConfiguration\&quot; are reserved and cannot be set: &#x60;web.submit.jar-sources&#x60;, &#x60;pipeline.jars&#x60;, &#x60;pipeline.classpaths&#x60;, every key beginning with &#x60;kubernetes.operator.cmf.&#x60;, and every key beginning with &#x60;user.artifacts.&#x60; except &#x60;user.artifacts.max-concurrent-fetches&#x60;, &#x60;user.artifacts.max-size&#x60;, &#x60;user.artifacts.fetch-timeout&#x60;, &#x60;user.artifacts.base-dir&#x60;, &#x60;user.artifacts.http-headers&#x60; and &#x60;user.artifacts.raw-http-enabled&#x60;. An environment that sets one is rejected with an error naming the key.  An update that changes the Flink version of a compute pool (&#x60;clusterSpec.flinkVersion&#x60;, or the Flink line of a &#x60;cp-flink-sql&#x60; image in &#x60;clusterSpec.image&#x60;) is rejected while Statements on that pool can still be resumed or rescaled; the error names them.  | [optional] 

## Methods

### NewComputePoolDefaults

`func NewComputePoolDefaults() *ComputePoolDefaults`

NewComputePoolDefaults instantiates a new ComputePoolDefaults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputePoolDefaultsWithDefaults

`func NewComputePoolDefaultsWithDefaults() *ComputePoolDefaults`

NewComputePoolDefaultsWithDefaults instantiates a new ComputePoolDefaults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComputePoolDefaults

`func (o *ComputePoolDefaults) GetComputePoolDefaults() map[string]interface{}`

GetComputePoolDefaults returns the ComputePoolDefaults field if non-nil, zero value otherwise.

### GetComputePoolDefaultsOk

`func (o *ComputePoolDefaults) GetComputePoolDefaultsOk() (*map[string]interface{}, bool)`

GetComputePoolDefaultsOk returns a tuple with the ComputePoolDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputePoolDefaults

`func (o *ComputePoolDefaults) SetComputePoolDefaults(v map[string]interface{})`

SetComputePoolDefaults sets ComputePoolDefaults field to given value.

### HasComputePoolDefaults

`func (o *ComputePoolDefaults) HasComputePoolDefaults() bool`

HasComputePoolDefaults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


