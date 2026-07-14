# SavepointRetention

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeepLast** | Pointer to **int32** | Keep the N most recently completed periodic savepoints. | [optional] 
**KeepWithin** | Pointer to **string** | ISO-8601 duration; keep every completed periodic savepoint taken within this duration of now. Examples: \&quot;PT24H\&quot;, \&quot;P7D\&quot;.  | [optional] 
**KeepHourly** | Pointer to **int32** | Keep the latest periodic savepoint per calendar hour for the most recent N hours. | [optional] 
**KeepDaily** | Pointer to **int32** | Keep the latest periodic savepoint per calendar day for the most recent N days. | [optional] 
**KeepWeekly** | Pointer to **int32** | Keep the latest periodic savepoint per ISO week for the most recent N weeks. | [optional] 
**KeepMonthly** | Pointer to **int32** | Keep the latest periodic savepoint per calendar month for the most recent N months. | [optional] 
**KeepYearly** | Pointer to **int32** | Keep the latest periodic savepoint per calendar year for the most recent N years. | [optional] 

## Methods

### NewSavepointRetention

`func NewSavepointRetention() *SavepointRetention`

NewSavepointRetention instantiates a new SavepointRetention object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSavepointRetentionWithDefaults

`func NewSavepointRetentionWithDefaults() *SavepointRetention`

NewSavepointRetentionWithDefaults instantiates a new SavepointRetention object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeepLast

`func (o *SavepointRetention) GetKeepLast() int32`

GetKeepLast returns the KeepLast field if non-nil, zero value otherwise.

### GetKeepLastOk

`func (o *SavepointRetention) GetKeepLastOk() (*int32, bool)`

GetKeepLastOk returns a tuple with the KeepLast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepLast

`func (o *SavepointRetention) SetKeepLast(v int32)`

SetKeepLast sets KeepLast field to given value.

### HasKeepLast

`func (o *SavepointRetention) HasKeepLast() bool`

HasKeepLast returns a boolean if a field has been set.

### GetKeepWithin

`func (o *SavepointRetention) GetKeepWithin() string`

GetKeepWithin returns the KeepWithin field if non-nil, zero value otherwise.

### GetKeepWithinOk

`func (o *SavepointRetention) GetKeepWithinOk() (*string, bool)`

GetKeepWithinOk returns a tuple with the KeepWithin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepWithin

`func (o *SavepointRetention) SetKeepWithin(v string)`

SetKeepWithin sets KeepWithin field to given value.

### HasKeepWithin

`func (o *SavepointRetention) HasKeepWithin() bool`

HasKeepWithin returns a boolean if a field has been set.

### GetKeepHourly

`func (o *SavepointRetention) GetKeepHourly() int32`

GetKeepHourly returns the KeepHourly field if non-nil, zero value otherwise.

### GetKeepHourlyOk

`func (o *SavepointRetention) GetKeepHourlyOk() (*int32, bool)`

GetKeepHourlyOk returns a tuple with the KeepHourly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepHourly

`func (o *SavepointRetention) SetKeepHourly(v int32)`

SetKeepHourly sets KeepHourly field to given value.

### HasKeepHourly

`func (o *SavepointRetention) HasKeepHourly() bool`

HasKeepHourly returns a boolean if a field has been set.

### GetKeepDaily

`func (o *SavepointRetention) GetKeepDaily() int32`

GetKeepDaily returns the KeepDaily field if non-nil, zero value otherwise.

### GetKeepDailyOk

`func (o *SavepointRetention) GetKeepDailyOk() (*int32, bool)`

GetKeepDailyOk returns a tuple with the KeepDaily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepDaily

`func (o *SavepointRetention) SetKeepDaily(v int32)`

SetKeepDaily sets KeepDaily field to given value.

### HasKeepDaily

`func (o *SavepointRetention) HasKeepDaily() bool`

HasKeepDaily returns a boolean if a field has been set.

### GetKeepWeekly

`func (o *SavepointRetention) GetKeepWeekly() int32`

GetKeepWeekly returns the KeepWeekly field if non-nil, zero value otherwise.

### GetKeepWeeklyOk

`func (o *SavepointRetention) GetKeepWeeklyOk() (*int32, bool)`

GetKeepWeeklyOk returns a tuple with the KeepWeekly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepWeekly

`func (o *SavepointRetention) SetKeepWeekly(v int32)`

SetKeepWeekly sets KeepWeekly field to given value.

### HasKeepWeekly

`func (o *SavepointRetention) HasKeepWeekly() bool`

HasKeepWeekly returns a boolean if a field has been set.

### GetKeepMonthly

`func (o *SavepointRetention) GetKeepMonthly() int32`

GetKeepMonthly returns the KeepMonthly field if non-nil, zero value otherwise.

### GetKeepMonthlyOk

`func (o *SavepointRetention) GetKeepMonthlyOk() (*int32, bool)`

GetKeepMonthlyOk returns a tuple with the KeepMonthly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepMonthly

`func (o *SavepointRetention) SetKeepMonthly(v int32)`

SetKeepMonthly sets KeepMonthly field to given value.

### HasKeepMonthly

`func (o *SavepointRetention) HasKeepMonthly() bool`

HasKeepMonthly returns a boolean if a field has been set.

### GetKeepYearly

`func (o *SavepointRetention) GetKeepYearly() int32`

GetKeepYearly returns the KeepYearly field if non-nil, zero value otherwise.

### GetKeepYearlyOk

`func (o *SavepointRetention) GetKeepYearlyOk() (*int32, bool)`

GetKeepYearlyOk returns a tuple with the KeepYearly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepYearly

`func (o *SavepointRetention) SetKeepYearly(v int32)`

SetKeepYearly sets KeepYearly field to given value.

### HasKeepYearly

`func (o *SavepointRetention) HasKeepYearly() bool`

HasKeepYearly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


