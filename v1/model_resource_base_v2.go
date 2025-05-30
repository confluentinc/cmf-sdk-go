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

// ResourceBaseV2 struct for ResourceBaseV2
type ResourceBaseV2 struct {
	// API version for spec
	ApiVersion string `json:"apiVersion"`
	// Kind of resource - set to resource type
	Kind string `json:"kind"`
}

// NewResourceBaseV2 instantiates a new ResourceBaseV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceBaseV2(apiVersion string, kind string) *ResourceBaseV2 {
	this := ResourceBaseV2{}
	this.ApiVersion = apiVersion
	this.Kind = kind
	return &this
}

// NewResourceBaseV2WithDefaults instantiates a new ResourceBaseV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceBaseV2WithDefaults() *ResourceBaseV2 {
	this := ResourceBaseV2{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *ResourceBaseV2) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *ResourceBaseV2) GetApiVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *ResourceBaseV2) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetKind returns the Kind field value
func (o *ResourceBaseV2) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *ResourceBaseV2) GetKindOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *ResourceBaseV2) SetKind(v string) {
	o.Kind = v
}

func (o ResourceBaseV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if true {
		toSerialize["kind"] = o.Kind
	}
	return json.Marshal(toSerialize)
}

type NullableResourceBaseV2 struct {
	value *ResourceBaseV2
	isSet bool
}

func (v NullableResourceBaseV2) Get() *ResourceBaseV2 {
	return v.value
}

func (v *NullableResourceBaseV2) Set(val *ResourceBaseV2) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceBaseV2) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceBaseV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceBaseV2(val *ResourceBaseV2) *NullableResourceBaseV2 {
	return &NullableResourceBaseV2{value: val, isSet: true}
}

func (v NullableResourceBaseV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceBaseV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


