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

// SecretAllOf struct for SecretAllOf
type SecretAllOf struct {
	Metadata SecretMetadata `json:"metadata"`
	Spec SecretSpec `json:"spec"`
	Status *SecretStatus `json:"status,omitempty"`
}

// NewSecretAllOf instantiates a new SecretAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecretAllOf(metadata SecretMetadata, spec SecretSpec) *SecretAllOf {
	this := SecretAllOf{}
	this.Metadata = metadata
	this.Spec = spec
	return &this
}

// NewSecretAllOfWithDefaults instantiates a new SecretAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretAllOfWithDefaults() *SecretAllOf {
	this := SecretAllOf{}
	return &this
}

// GetMetadata returns the Metadata field value
func (o *SecretAllOf) GetMetadata() SecretMetadata {
	if o == nil {
		var ret SecretMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *SecretAllOf) GetMetadataOk() (*SecretMetadata, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *SecretAllOf) SetMetadata(v SecretMetadata) {
	o.Metadata = v
}

// GetSpec returns the Spec field value
func (o *SecretAllOf) GetSpec() SecretSpec {
	if o == nil {
		var ret SecretSpec
		return ret
	}

	return o.Spec
}

// GetSpecOk returns a tuple with the Spec field value
// and a boolean to check if the value has been set.
func (o *SecretAllOf) GetSpecOk() (*SecretSpec, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Spec, true
}

// SetSpec sets field value
func (o *SecretAllOf) SetSpec(v SecretSpec) {
	o.Spec = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SecretAllOf) GetStatus() SecretStatus {
	if o == nil || o.Status == nil {
		var ret SecretStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretAllOf) GetStatusOk() (*SecretStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SecretAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given SecretStatus and assigns it to the Status field.
func (o *SecretAllOf) SetStatus(v SecretStatus) {
	o.Status = &v
}

func (o SecretAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["spec"] = o.Spec
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableSecretAllOf struct {
	value *SecretAllOf
	isSet bool
}

func (v NullableSecretAllOf) Get() *SecretAllOf {
	return v.value
}

func (v *NullableSecretAllOf) Set(val *SecretAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSecretAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSecretAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecretAllOf(val *SecretAllOf) *NullableSecretAllOf {
	return &NullableSecretAllOf{value: val, isSet: true}
}

func (v NullableSecretAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecretAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


