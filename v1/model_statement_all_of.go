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

// StatementAllOf struct for StatementAllOf
type StatementAllOf struct {
	Metadata StatementMetadata `json:"metadata"`
	Spec StatementSpec `json:"spec"`
	Status *StatementStatus `json:"status,omitempty"`
	Result *StatementResult `json:"result,omitempty"`
}

// NewStatementAllOf instantiates a new StatementAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatementAllOf(metadata StatementMetadata, spec StatementSpec) *StatementAllOf {
	this := StatementAllOf{}
	this.Metadata = metadata
	this.Spec = spec
	return &this
}

// NewStatementAllOfWithDefaults instantiates a new StatementAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatementAllOfWithDefaults() *StatementAllOf {
	this := StatementAllOf{}
	return &this
}

// GetMetadata returns the Metadata field value
func (o *StatementAllOf) GetMetadata() StatementMetadata {
	if o == nil {
		var ret StatementMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *StatementAllOf) GetMetadataOk() (*StatementMetadata, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *StatementAllOf) SetMetadata(v StatementMetadata) {
	o.Metadata = v
}

// GetSpec returns the Spec field value
func (o *StatementAllOf) GetSpec() StatementSpec {
	if o == nil {
		var ret StatementSpec
		return ret
	}

	return o.Spec
}

// GetSpecOk returns a tuple with the Spec field value
// and a boolean to check if the value has been set.
func (o *StatementAllOf) GetSpecOk() (*StatementSpec, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Spec, true
}

// SetSpec sets field value
func (o *StatementAllOf) SetSpec(v StatementSpec) {
	o.Spec = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *StatementAllOf) GetStatus() StatementStatus {
	if o == nil || o.Status == nil {
		var ret StatementStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatementAllOf) GetStatusOk() (*StatementStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *StatementAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given StatementStatus and assigns it to the Status field.
func (o *StatementAllOf) SetStatus(v StatementStatus) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *StatementAllOf) GetResult() StatementResult {
	if o == nil || o.Result == nil {
		var ret StatementResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatementAllOf) GetResultOk() (*StatementResult, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *StatementAllOf) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given StatementResult and assigns it to the Result field.
func (o *StatementAllOf) SetResult(v StatementResult) {
	o.Result = &v
}

func (o StatementAllOf) MarshalJSON() ([]byte, error) {
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
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	return json.Marshal(toSerialize)
}

type NullableStatementAllOf struct {
	value *StatementAllOf
	isSet bool
}

func (v NullableStatementAllOf) Get() *StatementAllOf {
	return v.value
}

func (v *NullableStatementAllOf) Set(val *StatementAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStatementAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStatementAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatementAllOf(val *StatementAllOf) *NullableStatementAllOf {
	return &NullableStatementAllOf{value: val, isSet: true}
}

func (v NullableStatementAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatementAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


