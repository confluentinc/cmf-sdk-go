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

// StatementResultMetadata Metadata about the StatementResult
type StatementResultMetadata struct {
	// Timestamp when the StatementResult was created
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`
	// Annotations of the StatementResult
	Annotations *map[string]string `json:"annotations,omitempty"`
}

// NewStatementResultMetadata instantiates a new StatementResultMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatementResultMetadata() *StatementResultMetadata {
	this := StatementResultMetadata{}
	return &this
}

// NewStatementResultMetadataWithDefaults instantiates a new StatementResultMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatementResultMetadataWithDefaults() *StatementResultMetadata {
	this := StatementResultMetadata{}
	return &this
}

// GetCreationTimestamp returns the CreationTimestamp field value if set, zero value otherwise.
func (o *StatementResultMetadata) GetCreationTimestamp() string {
	if o == nil || o.CreationTimestamp == nil {
		var ret string
		return ret
	}
	return *o.CreationTimestamp
}

// GetCreationTimestampOk returns a tuple with the CreationTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatementResultMetadata) GetCreationTimestampOk() (*string, bool) {
	if o == nil || o.CreationTimestamp == nil {
		return nil, false
	}
	return o.CreationTimestamp, true
}

// HasCreationTimestamp returns a boolean if a field has been set.
func (o *StatementResultMetadata) HasCreationTimestamp() bool {
	if o != nil && o.CreationTimestamp != nil {
		return true
	}

	return false
}

// SetCreationTimestamp gets a reference to the given string and assigns it to the CreationTimestamp field.
func (o *StatementResultMetadata) SetCreationTimestamp(v string) {
	o.CreationTimestamp = &v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *StatementResultMetadata) GetAnnotations() map[string]string {
	if o == nil || o.Annotations == nil {
		var ret map[string]string
		return ret
	}
	return *o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatementResultMetadata) GetAnnotationsOk() (*map[string]string, bool) {
	if o == nil || o.Annotations == nil {
		return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *StatementResultMetadata) HasAnnotations() bool {
	if o != nil && o.Annotations != nil {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given map[string]string and assigns it to the Annotations field.
func (o *StatementResultMetadata) SetAnnotations(v map[string]string) {
	o.Annotations = &v
}

func (o StatementResultMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreationTimestamp != nil {
		toSerialize["creationTimestamp"] = o.CreationTimestamp
	}
	if o.Annotations != nil {
		toSerialize["annotations"] = o.Annotations
	}
	return json.Marshal(toSerialize)
}

type NullableStatementResultMetadata struct {
	value *StatementResultMetadata
	isSet bool
}

func (v NullableStatementResultMetadata) Get() *StatementResultMetadata {
	return v.value
}

func (v *NullableStatementResultMetadata) Set(val *StatementResultMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableStatementResultMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableStatementResultMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatementResultMetadata(val *StatementResultMetadata) *NullableStatementResultMetadata {
	return &NullableStatementResultMetadata{value: val, isSet: true}
}

func (v NullableStatementResultMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatementResultMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


