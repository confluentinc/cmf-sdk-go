# ResourceWithFlinkApplicationDefaults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlinkApplicationDefaults** | Pointer to **map[string]interface{}** | Environment-level defaults for FlinkApplication specs. The structure mirrors a FlinkApplication itself: place a \&quot;spec\&quot; object here whose fields are merged into every application spec at deploy time. May include a \&quot;savepointSchedule\&quot; inside \&quot;spec\&quot; (see SavepointScheduleConfig) to set the default periodic savepoint schedule for every application in the environment; env-level schedule fields take precedence over any value set on the resource itself.  | [optional] 

## Methods

### NewResourceWithFlinkApplicationDefaults

`func NewResourceWithFlinkApplicationDefaults() *ResourceWithFlinkApplicationDefaults`

NewResourceWithFlinkApplicationDefaults instantiates a new ResourceWithFlinkApplicationDefaults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceWithFlinkApplicationDefaultsWithDefaults

`func NewResourceWithFlinkApplicationDefaultsWithDefaults() *ResourceWithFlinkApplicationDefaults`

NewResourceWithFlinkApplicationDefaultsWithDefaults instantiates a new ResourceWithFlinkApplicationDefaults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlinkApplicationDefaults

`func (o *ResourceWithFlinkApplicationDefaults) GetFlinkApplicationDefaults() map[string]interface{}`

GetFlinkApplicationDefaults returns the FlinkApplicationDefaults field if non-nil, zero value otherwise.

### GetFlinkApplicationDefaultsOk

`func (o *ResourceWithFlinkApplicationDefaults) GetFlinkApplicationDefaultsOk() (*map[string]interface{}, bool)`

GetFlinkApplicationDefaultsOk returns a tuple with the FlinkApplicationDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlinkApplicationDefaults

`func (o *ResourceWithFlinkApplicationDefaults) SetFlinkApplicationDefaults(v map[string]interface{})`

SetFlinkApplicationDefaults sets FlinkApplicationDefaults field to given value.

### HasFlinkApplicationDefaults

`func (o *ResourceWithFlinkApplicationDefaults) HasFlinkApplicationDefaults() bool`

HasFlinkApplicationDefaults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


