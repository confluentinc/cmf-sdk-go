# ApplicationInstanceStatusJobStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | Pointer to **string** | Flink job id inside the Flink cluster | [optional] 
**State** | Pointer to **string** | Tracks the final Flink JobStatus of the instance | [optional] 

## Methods

### NewApplicationInstanceStatusJobStatus

`func NewApplicationInstanceStatusJobStatus() *ApplicationInstanceStatusJobStatus`

NewApplicationInstanceStatusJobStatus instantiates a new ApplicationInstanceStatusJobStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationInstanceStatusJobStatusWithDefaults

`func NewApplicationInstanceStatusJobStatusWithDefaults() *ApplicationInstanceStatusJobStatus`

NewApplicationInstanceStatusJobStatusWithDefaults instantiates a new ApplicationInstanceStatusJobStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *ApplicationInstanceStatusJobStatus) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *ApplicationInstanceStatusJobStatus) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *ApplicationInstanceStatusJobStatus) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *ApplicationInstanceStatusJobStatus) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetState

`func (o *ApplicationInstanceStatusJobStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ApplicationInstanceStatusJobStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ApplicationInstanceStatusJobStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ApplicationInstanceStatusJobStatus) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


