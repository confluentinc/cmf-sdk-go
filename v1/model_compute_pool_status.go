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

// ComputePoolStatus Status for ComputePool
type ComputePoolStatus struct {
	// Phase of the ComputePool
	Phase string `json:"phase"`
}

// NewComputePoolStatus instantiates a new ComputePoolStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputePoolStatus(phase string) *ComputePoolStatus {
	this := ComputePoolStatus{}
	this.Phase = phase
	return &this
}

// NewComputePoolStatusWithDefaults instantiates a new ComputePoolStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputePoolStatusWithDefaults() *ComputePoolStatus {
	this := ComputePoolStatus{}
	return &this
}

// GetPhase returns the Phase field value
func (o *ComputePoolStatus) GetPhase() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value
// and a boolean to check if the value has been set.
func (o *ComputePoolStatus) GetPhaseOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Phase, true
}

// SetPhase sets field value
func (o *ComputePoolStatus) SetPhase(v string) {
	o.Phase = v
}

func (o ComputePoolStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["phase"] = o.Phase
	}
	return json.Marshal(toSerialize)
}

type NullableComputePoolStatus struct {
	value *ComputePoolStatus
	isSet bool
}

func (v NullableComputePoolStatus) Get() *ComputePoolStatus {
	return v.value
}

func (v *NullableComputePoolStatus) Set(val *ComputePoolStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableComputePoolStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableComputePoolStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputePoolStatus(val *ComputePoolStatus) *NullableComputePoolStatus {
	return &NullableComputePoolStatus{value: val, isSet: true}
}

func (v NullableComputePoolStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputePoolStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


