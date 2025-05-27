# ApplicationInstanceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | Pointer to **map[string]interface{}** | The environment defaults merged with the FlinkApplication spec at instance creation time | [optional] 
**JobStatus** | Pointer to [**ApplicationInstanceStatusJobStatus**](ApplicationInstanceStatusJobStatus.md) |  | [optional] 

## Methods

### NewApplicationInstanceStatus

`func NewApplicationInstanceStatus() *ApplicationInstanceStatus`

NewApplicationInstanceStatus instantiates a new ApplicationInstanceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationInstanceStatusWithDefaults

`func NewApplicationInstanceStatusWithDefaults() *ApplicationInstanceStatus`

NewApplicationInstanceStatusWithDefaults instantiates a new ApplicationInstanceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *ApplicationInstanceStatus) GetSpec() map[string]interface{}`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ApplicationInstanceStatus) GetSpecOk() (*map[string]interface{}, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ApplicationInstanceStatus) SetSpec(v map[string]interface{})`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ApplicationInstanceStatus) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetJobStatus

`func (o *ApplicationInstanceStatus) GetJobStatus() ApplicationInstanceStatusJobStatus`

GetJobStatus returns the JobStatus field if non-nil, zero value otherwise.

### GetJobStatusOk

`func (o *ApplicationInstanceStatus) GetJobStatusOk() (*ApplicationInstanceStatusJobStatus, bool)`

GetJobStatusOk returns a tuple with the JobStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobStatus

`func (o *ApplicationInstanceStatus) SetJobStatus(v ApplicationInstanceStatusJobStatus)`

SetJobStatus sets JobStatus field to given value.

### HasJobStatus

`func (o *ApplicationInstanceStatus) HasJobStatus() bool`

HasJobStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


