# ComputePoolSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the ComputePool | 
**State** | Pointer to **string** | Desired state of the compute pool. Can only be specified for shared compute pools. Defaults to RUNNING if not specified. | [optional] 
**AllowedWorkloadKinds** | Pointer to **[]string** | Reserved for a future release. Do not set this field. On a shared compute pool it limits which kinds of workload the pool accepts, and its behavior may change without notice. It cannot be set on dedicated compute pools. | [optional] 
**AllowedJarUriPrefixes** | Pointer to **[]string** | Reserved for a future release. Do not set this field. On a shared compute pool whose allowedWorkloadKinds includes APPLICATION it lists the URI prefixes that Application jars may be fetched from, and its behavior may change without notice. It cannot be set on other compute pools, and is cleared when APPLICATION is removed from allowedWorkloadKinds. | [optional] 
**ClusterSpec** | **map[string]interface{}** | Cluster Spec  Some keys in \&quot;flinkConfiguration\&quot; are reserved and cannot be set. For every compute pool, these are &#x60;web.submit.jar-sources&#x60; and every key beginning with &#x60;kubernetes.operator.cmf.&#x60;. For a shared compute pool, they also include &#x60;pipeline.jars&#x60;, &#x60;pipeline.classpaths&#x60;, and every key beginning with &#x60;user.artifacts.&#x60; except &#x60;user.artifacts.max-concurrent-fetches&#x60;, &#x60;user.artifacts.max-size&#x60;, &#x60;user.artifacts.fetch-timeout&#x60;, &#x60;user.artifacts.base-dir&#x60;, &#x60;user.artifacts.http-headers&#x60; and &#x60;user.artifacts.raw-http-enabled&#x60;. A compute pool that sets one is rejected with an error naming the key.  | 
**CpFlinkSqlVersion** | Pointer to **string** | CP Flink SQL version that statements on this compute pool are compiled against. Releases for CMF 2.5 and later have the format &#x60;&lt;major&gt;.&lt;minor&gt;-cmf&lt;X.Y&gt;-cp&lt;build&gt;&#x60;, for example &#x60;1.19-cmf2.5-cp1&#x60;, where &#x60;&lt;X.Y&gt;&#x60; is the CMF minor version and &#x60;&lt;build&gt;&#x60; counts from 1 within it. The earlier releases &#x60;1.19-cp1&#x60; to &#x60;1.19-cp12&#x60; are also accepted, and are older than every other 1.19 version. A release candidate, for example &#x60;1.19-rc1-cmf2.5-cp1&#x60;, is the same version as its release. No other format is accepted. A compiler compiles for the versions of its Flink line up to every build of its own CMF minor version. A version of a later CMF minor version than this one is rejected, and so is a version older than the oldest version of its Flink line that CMF supports (&#x60;1.19-cp8&#x60; for 1.19, &#x60;2.3-cmf2.5-cp1&#x60; for 2.3), and a version of another Flink major version than the compute pool&#39;s &#x60;clusterSpec.flinkVersion&#x60;. Only used when CMF is configured with shared SQL compilers or compile workers. Creating a statement fails if no configured compiler supports this version, or if this version has another Flink major version than the compute pool&#39;s Flink version. When unset, CMF uses the version in a &#x60;cp-flink-sql&#x60; image tag if a configured compiler supports it and it has the compute pool&#39;s Flink major version, and otherwise a compiler matching the compute pool&#39;s Flink version. An update that leaves this field out keeps it; an empty string clears it. | [optional] 

## Methods

### NewComputePoolSpec

`func NewComputePoolSpec(type_ string, clusterSpec map[string]interface{}, ) *ComputePoolSpec`

NewComputePoolSpec instantiates a new ComputePoolSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputePoolSpecWithDefaults

`func NewComputePoolSpecWithDefaults() *ComputePoolSpec`

NewComputePoolSpecWithDefaults instantiates a new ComputePoolSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ComputePoolSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ComputePoolSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ComputePoolSpec) SetType(v string)`

SetType sets Type field to given value.


### GetState

`func (o *ComputePoolSpec) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ComputePoolSpec) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ComputePoolSpec) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ComputePoolSpec) HasState() bool`

HasState returns a boolean if a field has been set.

### GetAllowedWorkloadKinds

`func (o *ComputePoolSpec) GetAllowedWorkloadKinds() []string`

GetAllowedWorkloadKinds returns the AllowedWorkloadKinds field if non-nil, zero value otherwise.

### GetAllowedWorkloadKindsOk

`func (o *ComputePoolSpec) GetAllowedWorkloadKindsOk() (*[]string, bool)`

GetAllowedWorkloadKindsOk returns a tuple with the AllowedWorkloadKinds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedWorkloadKinds

`func (o *ComputePoolSpec) SetAllowedWorkloadKinds(v []string)`

SetAllowedWorkloadKinds sets AllowedWorkloadKinds field to given value.

### HasAllowedWorkloadKinds

`func (o *ComputePoolSpec) HasAllowedWorkloadKinds() bool`

HasAllowedWorkloadKinds returns a boolean if a field has been set.

### SetAllowedWorkloadKindsNil

`func (o *ComputePoolSpec) SetAllowedWorkloadKindsNil(b bool)`

 SetAllowedWorkloadKindsNil sets the value for AllowedWorkloadKinds to be an explicit nil

### UnsetAllowedWorkloadKinds
`func (o *ComputePoolSpec) UnsetAllowedWorkloadKinds()`

UnsetAllowedWorkloadKinds ensures that no value is present for AllowedWorkloadKinds, not even an explicit nil
### GetAllowedJarUriPrefixes

`func (o *ComputePoolSpec) GetAllowedJarUriPrefixes() []string`

GetAllowedJarUriPrefixes returns the AllowedJarUriPrefixes field if non-nil, zero value otherwise.

### GetAllowedJarUriPrefixesOk

`func (o *ComputePoolSpec) GetAllowedJarUriPrefixesOk() (*[]string, bool)`

GetAllowedJarUriPrefixesOk returns a tuple with the AllowedJarUriPrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedJarUriPrefixes

`func (o *ComputePoolSpec) SetAllowedJarUriPrefixes(v []string)`

SetAllowedJarUriPrefixes sets AllowedJarUriPrefixes field to given value.

### HasAllowedJarUriPrefixes

`func (o *ComputePoolSpec) HasAllowedJarUriPrefixes() bool`

HasAllowedJarUriPrefixes returns a boolean if a field has been set.

### SetAllowedJarUriPrefixesNil

`func (o *ComputePoolSpec) SetAllowedJarUriPrefixesNil(b bool)`

 SetAllowedJarUriPrefixesNil sets the value for AllowedJarUriPrefixes to be an explicit nil

### UnsetAllowedJarUriPrefixes
`func (o *ComputePoolSpec) UnsetAllowedJarUriPrefixes()`

UnsetAllowedJarUriPrefixes ensures that no value is present for AllowedJarUriPrefixes, not even an explicit nil
### GetClusterSpec

`func (o *ComputePoolSpec) GetClusterSpec() map[string]interface{}`

GetClusterSpec returns the ClusterSpec field if non-nil, zero value otherwise.

### GetClusterSpecOk

`func (o *ComputePoolSpec) GetClusterSpecOk() (*map[string]interface{}, bool)`

GetClusterSpecOk returns a tuple with the ClusterSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterSpec

`func (o *ComputePoolSpec) SetClusterSpec(v map[string]interface{})`

SetClusterSpec sets ClusterSpec field to given value.


### GetCpFlinkSqlVersion

`func (o *ComputePoolSpec) GetCpFlinkSqlVersion() string`

GetCpFlinkSqlVersion returns the CpFlinkSqlVersion field if non-nil, zero value otherwise.

### GetCpFlinkSqlVersionOk

`func (o *ComputePoolSpec) GetCpFlinkSqlVersionOk() (*string, bool)`

GetCpFlinkSqlVersionOk returns a tuple with the CpFlinkSqlVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpFlinkSqlVersion

`func (o *ComputePoolSpec) SetCpFlinkSqlVersion(v string)`

SetCpFlinkSqlVersion sets CpFlinkSqlVersion field to given value.

### HasCpFlinkSqlVersion

`func (o *ComputePoolSpec) HasCpFlinkSqlVersion() bool`

HasCpFlinkSqlVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


