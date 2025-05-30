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

// FlinkApplicationAllOf struct for FlinkApplicationAllOf
type FlinkApplicationAllOf struct {
	Metadata map[string]interface{} `json:"metadata"`
	Spec map[string]interface{} `json:"spec"`
	Status *map[string]interface{} `json:"status,omitempty"`
}

// NewFlinkApplicationAllOf instantiates a new FlinkApplicationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlinkApplicationAllOf(metadata map[string]interface{}, spec map[string]interface{}) *FlinkApplicationAllOf {
	this := FlinkApplicationAllOf{}
	this.Metadata = metadata
	this.Spec = spec
	return &this
}

// NewFlinkApplicationAllOfWithDefaults instantiates a new FlinkApplicationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlinkApplicationAllOfWithDefaults() *FlinkApplicationAllOf {
	this := FlinkApplicationAllOf{}
	return &this
}

// GetMetadata returns the Metadata field value
func (o *FlinkApplicationAllOf) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *FlinkApplicationAllOf) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *FlinkApplicationAllOf) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetSpec returns the Spec field value
func (o *FlinkApplicationAllOf) GetSpec() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Spec
}

// GetSpecOk returns a tuple with the Spec field value
// and a boolean to check if the value has been set.
func (o *FlinkApplicationAllOf) GetSpecOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Spec, true
}

// SetSpec sets field value
func (o *FlinkApplicationAllOf) SetSpec(v map[string]interface{}) {
	o.Spec = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FlinkApplicationAllOf) GetStatus() map[string]interface{} {
	if o == nil || o.Status == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlinkApplicationAllOf) GetStatusOk() (*map[string]interface{}, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FlinkApplicationAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given map[string]interface{} and assigns it to the Status field.
func (o *FlinkApplicationAllOf) SetStatus(v map[string]interface{}) {
	o.Status = &v
}

func (o FlinkApplicationAllOf) MarshalJSON() ([]byte, error) {
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

type NullableFlinkApplicationAllOf struct {
	value *FlinkApplicationAllOf
	isSet bool
}

func (v NullableFlinkApplicationAllOf) Get() *FlinkApplicationAllOf {
	return v.value
}

func (v *NullableFlinkApplicationAllOf) Set(val *FlinkApplicationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFlinkApplicationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFlinkApplicationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlinkApplicationAllOf(val *FlinkApplicationAllOf) *NullableFlinkApplicationAllOf {
	return &NullableFlinkApplicationAllOf{value: val, isSet: true}
}

func (v NullableFlinkApplicationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlinkApplicationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


