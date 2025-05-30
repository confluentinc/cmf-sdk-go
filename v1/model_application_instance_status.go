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

// ApplicationInstanceStatus struct for ApplicationInstanceStatus
type ApplicationInstanceStatus struct {
	// The environment defaults merged with the FlinkApplication spec at instance creation time
	Spec *map[string]interface{} `json:"spec,omitempty"`
	JobStatus *ApplicationInstanceStatusJobStatus `json:"jobStatus,omitempty"`
}

// NewApplicationInstanceStatus instantiates a new ApplicationInstanceStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationInstanceStatus() *ApplicationInstanceStatus {
	this := ApplicationInstanceStatus{}
	return &this
}

// NewApplicationInstanceStatusWithDefaults instantiates a new ApplicationInstanceStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationInstanceStatusWithDefaults() *ApplicationInstanceStatus {
	this := ApplicationInstanceStatus{}
	return &this
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *ApplicationInstanceStatus) GetSpec() map[string]interface{} {
	if o == nil || o.Spec == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationInstanceStatus) GetSpecOk() (*map[string]interface{}, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *ApplicationInstanceStatus) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given map[string]interface{} and assigns it to the Spec field.
func (o *ApplicationInstanceStatus) SetSpec(v map[string]interface{}) {
	o.Spec = &v
}

// GetJobStatus returns the JobStatus field value if set, zero value otherwise.
func (o *ApplicationInstanceStatus) GetJobStatus() ApplicationInstanceStatusJobStatus {
	if o == nil || o.JobStatus == nil {
		var ret ApplicationInstanceStatusJobStatus
		return ret
	}
	return *o.JobStatus
}

// GetJobStatusOk returns a tuple with the JobStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationInstanceStatus) GetJobStatusOk() (*ApplicationInstanceStatusJobStatus, bool) {
	if o == nil || o.JobStatus == nil {
		return nil, false
	}
	return o.JobStatus, true
}

// HasJobStatus returns a boolean if a field has been set.
func (o *ApplicationInstanceStatus) HasJobStatus() bool {
	if o != nil && o.JobStatus != nil {
		return true
	}

	return false
}

// SetJobStatus gets a reference to the given ApplicationInstanceStatusJobStatus and assigns it to the JobStatus field.
func (o *ApplicationInstanceStatus) SetJobStatus(v ApplicationInstanceStatusJobStatus) {
	o.JobStatus = &v
}

func (o ApplicationInstanceStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Spec != nil {
		toSerialize["spec"] = o.Spec
	}
	if o.JobStatus != nil {
		toSerialize["jobStatus"] = o.JobStatus
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationInstanceStatus struct {
	value *ApplicationInstanceStatus
	isSet bool
}

func (v NullableApplicationInstanceStatus) Get() *ApplicationInstanceStatus {
	return v.value
}

func (v *NullableApplicationInstanceStatus) Set(val *ApplicationInstanceStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationInstanceStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationInstanceStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationInstanceStatus(val *ApplicationInstanceStatus) *NullableApplicationInstanceStatus {
	return &NullableApplicationInstanceStatus{value: val, isSet: true}
}

func (v NullableApplicationInstanceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationInstanceStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


