# C3ConfigurationProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | Property key in dotted form (e.g. cmf.k8s.enabled). | [optional] 
**Value** | Pointer to **NullableString** | Effective value as a string. &#x60;null&#x60; when the property has no resolvable value. Sensitive values are replaced with &#x60;******&#x60;.  | [optional] 
**Origin** | Pointer to **string** | Best-effort, human-readable source detail for display (e.g. a config file with line:column, or &#x60;default&#x60; when unset). Formats may change between releases; use &#x60;originType&#x60; for logic.  | [optional] 
**OriginType** | Pointer to **string** | Coarse, stable category of &#x60;origin&#x60; (how the value was supplied), for grouping or filtering. | [optional] 
**Sensitive** | Pointer to **bool** | True if the value was masked because it looks like a credential. Use this flag rather than comparing the value to &#x60;******&#x60;, so a value that legitimately equals &#x60;******&#x60; is not mistaken for a masked one.  | [optional] 

## Methods

### NewC3ConfigurationProperty

`func NewC3ConfigurationProperty() *C3ConfigurationProperty`

NewC3ConfigurationProperty instantiates a new C3ConfigurationProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC3ConfigurationPropertyWithDefaults

`func NewC3ConfigurationPropertyWithDefaults() *C3ConfigurationProperty`

NewC3ConfigurationPropertyWithDefaults instantiates a new C3ConfigurationProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *C3ConfigurationProperty) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *C3ConfigurationProperty) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *C3ConfigurationProperty) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *C3ConfigurationProperty) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *C3ConfigurationProperty) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *C3ConfigurationProperty) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *C3ConfigurationProperty) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *C3ConfigurationProperty) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *C3ConfigurationProperty) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *C3ConfigurationProperty) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetOrigin

`func (o *C3ConfigurationProperty) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *C3ConfigurationProperty) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *C3ConfigurationProperty) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *C3ConfigurationProperty) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetOriginType

`func (o *C3ConfigurationProperty) GetOriginType() string`

GetOriginType returns the OriginType field if non-nil, zero value otherwise.

### GetOriginTypeOk

`func (o *C3ConfigurationProperty) GetOriginTypeOk() (*string, bool)`

GetOriginTypeOk returns a tuple with the OriginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginType

`func (o *C3ConfigurationProperty) SetOriginType(v string)`

SetOriginType sets OriginType field to given value.

### HasOriginType

`func (o *C3ConfigurationProperty) HasOriginType() bool`

HasOriginType returns a boolean if a field has been set.

### GetSensitive

`func (o *C3ConfigurationProperty) GetSensitive() bool`

GetSensitive returns the Sensitive field if non-nil, zero value otherwise.

### GetSensitiveOk

`func (o *C3ConfigurationProperty) GetSensitiveOk() (*bool, bool)`

GetSensitiveOk returns a tuple with the Sensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitive

`func (o *C3ConfigurationProperty) SetSensitive(v bool)`

SetSensitive sets Sensitive field to given value.

### HasSensitive

`func (o *C3ConfigurationProperty) HasSensitive() bool`

HasSensitive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


