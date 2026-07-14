# BlackoutWindow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CronExpression** | **string** | Six-field cron expression with seconds resolution (second, minute, hour, day-of-month, month, day-of-week) whose firing times mark the START of each blackout window.  | 
**DurationMinutes** | Pointer to **int32** | Duration of the blackout in minutes starting from the cron fire time. | [optional] [default to 60]
**Timezone** | Pointer to **string** | IANA time zone for evaluating this window&#39;s cron expression. Defaults to the parent schedule&#39;s timezone when omitted.  | [optional] 

## Methods

### NewBlackoutWindow

`func NewBlackoutWindow(cronExpression string, ) *BlackoutWindow`

NewBlackoutWindow instantiates a new BlackoutWindow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlackoutWindowWithDefaults

`func NewBlackoutWindowWithDefaults() *BlackoutWindow`

NewBlackoutWindowWithDefaults instantiates a new BlackoutWindow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCronExpression

`func (o *BlackoutWindow) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *BlackoutWindow) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *BlackoutWindow) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.


### GetDurationMinutes

`func (o *BlackoutWindow) GetDurationMinutes() int32`

GetDurationMinutes returns the DurationMinutes field if non-nil, zero value otherwise.

### GetDurationMinutesOk

`func (o *BlackoutWindow) GetDurationMinutesOk() (*int32, bool)`

GetDurationMinutesOk returns a tuple with the DurationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMinutes

`func (o *BlackoutWindow) SetDurationMinutes(v int32)`

SetDurationMinutes sets DurationMinutes field to given value.

### HasDurationMinutes

`func (o *BlackoutWindow) HasDurationMinutes() bool`

HasDurationMinutes returns a boolean if a field has been set.

### GetTimezone

`func (o *BlackoutWindow) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *BlackoutWindow) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *BlackoutWindow) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *BlackoutWindow) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


