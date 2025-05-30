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

// StatementPageMetadata struct for StatementPageMetadata
type StatementPageMetadata struct {
	Size *int64 `json:"size,omitempty"`
}

// NewStatementPageMetadata instantiates a new StatementPageMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatementPageMetadata() *StatementPageMetadata {
	this := StatementPageMetadata{}
	var size int64 = 0
	this.Size = &size
	return &this
}

// NewStatementPageMetadataWithDefaults instantiates a new StatementPageMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatementPageMetadataWithDefaults() *StatementPageMetadata {
	this := StatementPageMetadata{}
	var size int64 = 0
	this.Size = &size
	return &this
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *StatementPageMetadata) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatementPageMetadata) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *StatementPageMetadata) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *StatementPageMetadata) SetSize(v int64) {
	o.Size = &v
}

func (o StatementPageMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	return json.Marshal(toSerialize)
}

type NullableStatementPageMetadata struct {
	value *StatementPageMetadata
	isSet bool
}

func (v NullableStatementPageMetadata) Get() *StatementPageMetadata {
	return v.value
}

func (v *NullableStatementPageMetadata) Set(val *StatementPageMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableStatementPageMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableStatementPageMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatementPageMetadata(val *StatementPageMetadata) *NullableStatementPageMetadata {
	return &NullableStatementPageMetadata{value: val, isSet: true}
}

func (v NullableStatementPageMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatementPageMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


