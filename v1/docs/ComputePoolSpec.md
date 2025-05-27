# ComputePoolSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the ComputePool | 
**ClusterSpec** | **map[string]interface{}** | Cluster Spec | 

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


