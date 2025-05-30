/*
Confluent Manager for Apache Flink / CMF

Apache Flink job lifecycle management component for Confluent Platform.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// EnvironmentSecretMappingsPageMetadata struct for EnvironmentSecretMappingsPageMetadata
type EnvironmentSecretMappingsPageMetadata struct {
	Size *int64 `json:"size,omitempty"`
}

// NewEnvironmentSecretMappingsPageMetadata instantiates a new EnvironmentSecretMappingsPageMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentSecretMappingsPageMetadata() *EnvironmentSecretMappingsPageMetadata {
	this := EnvironmentSecretMappingsPageMetadata{}
	var size int64 = 0
	this.Size = &size
	return &this
}

// NewEnvironmentSecretMappingsPageMetadataWithDefaults instantiates a new EnvironmentSecretMappingsPageMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentSecretMappingsPageMetadataWithDefaults() *EnvironmentSecretMappingsPageMetadata {
	this := EnvironmentSecretMappingsPageMetadata{}
	var size int64 = 0
	this.Size = &size
	return &this
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *EnvironmentSecretMappingsPageMetadata) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentSecretMappingsPageMetadata) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *EnvironmentSecretMappingsPageMetadata) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *EnvironmentSecretMappingsPageMetadata) SetSize(v int64) {
	o.Size = &v
}

func (o EnvironmentSecretMappingsPageMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentSecretMappingsPageMetadata struct {
	value *EnvironmentSecretMappingsPageMetadata
	isSet bool
}

func (v NullableEnvironmentSecretMappingsPageMetadata) Get() *EnvironmentSecretMappingsPageMetadata {
	return v.value
}

func (v *NullableEnvironmentSecretMappingsPageMetadata) Set(val *EnvironmentSecretMappingsPageMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentSecretMappingsPageMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentSecretMappingsPageMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentSecretMappingsPageMetadata(val *EnvironmentSecretMappingsPageMetadata) *NullableEnvironmentSecretMappingsPageMetadata {
	return &NullableEnvironmentSecretMappingsPageMetadata{value: val, isSet: true}
}

func (v NullableEnvironmentSecretMappingsPageMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentSecretMappingsPageMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


