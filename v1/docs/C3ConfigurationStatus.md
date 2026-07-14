# C3ConfigurationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationProperties** | Pointer to [**[]C3ConfigurationProperty**](C3ConfigurationProperty.md) | The curated view of CMF&#39;s own settings (the &#x60;cmf.*&#x60; and &#x60;encryption.*&#x60; configuration). Sorted by key.  | [optional] 

## Methods

### NewC3ConfigurationStatus

`func NewC3ConfigurationStatus() *C3ConfigurationStatus`

NewC3ConfigurationStatus instantiates a new C3ConfigurationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC3ConfigurationStatusWithDefaults

`func NewC3ConfigurationStatusWithDefaults() *C3ConfigurationStatus`

NewC3ConfigurationStatusWithDefaults instantiates a new C3ConfigurationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurationProperties

`func (o *C3ConfigurationStatus) GetConfigurationProperties() []C3ConfigurationProperty`

GetConfigurationProperties returns the ConfigurationProperties field if non-nil, zero value otherwise.

### GetConfigurationPropertiesOk

`func (o *C3ConfigurationStatus) GetConfigurationPropertiesOk() (*[]C3ConfigurationProperty, bool)`

GetConfigurationPropertiesOk returns a tuple with the ConfigurationProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationProperties

`func (o *C3ConfigurationStatus) SetConfigurationProperties(v []C3ConfigurationProperty)`

SetConfigurationProperties sets ConfigurationProperties field to given value.

### HasConfigurationProperties

`func (o *C3ConfigurationStatus) HasConfigurationProperties() bool`

HasConfigurationProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


