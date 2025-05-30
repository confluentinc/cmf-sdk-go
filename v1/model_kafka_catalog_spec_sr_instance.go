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

// KafkaCatalogSpecSrInstance Details about the SchemaRegistry instance of the Catalog
type KafkaCatalogSpecSrInstance struct {
	// connection options for the SR client
	ConnectionConfig map[string]string `json:"connectionConfig"`
	// an identifier to look up a Kubernetes secret that contains the connection credentials
	ConnectionSecretId *string `json:"connectionSecretId,omitempty"`
}

// NewKafkaCatalogSpecSrInstance instantiates a new KafkaCatalogSpecSrInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKafkaCatalogSpecSrInstance(connectionConfig map[string]string) *KafkaCatalogSpecSrInstance {
	this := KafkaCatalogSpecSrInstance{}
	this.ConnectionConfig = connectionConfig
	return &this
}

// NewKafkaCatalogSpecSrInstanceWithDefaults instantiates a new KafkaCatalogSpecSrInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKafkaCatalogSpecSrInstanceWithDefaults() *KafkaCatalogSpecSrInstance {
	this := KafkaCatalogSpecSrInstance{}
	return &this
}

// GetConnectionConfig returns the ConnectionConfig field value
func (o *KafkaCatalogSpecSrInstance) GetConnectionConfig() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.ConnectionConfig
}

// GetConnectionConfigOk returns a tuple with the ConnectionConfig field value
// and a boolean to check if the value has been set.
func (o *KafkaCatalogSpecSrInstance) GetConnectionConfigOk() (*map[string]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConnectionConfig, true
}

// SetConnectionConfig sets field value
func (o *KafkaCatalogSpecSrInstance) SetConnectionConfig(v map[string]string) {
	o.ConnectionConfig = v
}

// GetConnectionSecretId returns the ConnectionSecretId field value if set, zero value otherwise.
func (o *KafkaCatalogSpecSrInstance) GetConnectionSecretId() string {
	if o == nil || o.ConnectionSecretId == nil {
		var ret string
		return ret
	}
	return *o.ConnectionSecretId
}

// GetConnectionSecretIdOk returns a tuple with the ConnectionSecretId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaCatalogSpecSrInstance) GetConnectionSecretIdOk() (*string, bool) {
	if o == nil || o.ConnectionSecretId == nil {
		return nil, false
	}
	return o.ConnectionSecretId, true
}

// HasConnectionSecretId returns a boolean if a field has been set.
func (o *KafkaCatalogSpecSrInstance) HasConnectionSecretId() bool {
	if o != nil && o.ConnectionSecretId != nil {
		return true
	}

	return false
}

// SetConnectionSecretId gets a reference to the given string and assigns it to the ConnectionSecretId field.
func (o *KafkaCatalogSpecSrInstance) SetConnectionSecretId(v string) {
	o.ConnectionSecretId = &v
}

func (o KafkaCatalogSpecSrInstance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["connectionConfig"] = o.ConnectionConfig
	}
	if o.ConnectionSecretId != nil {
		toSerialize["connectionSecretId"] = o.ConnectionSecretId
	}
	return json.Marshal(toSerialize)
}

type NullableKafkaCatalogSpecSrInstance struct {
	value *KafkaCatalogSpecSrInstance
	isSet bool
}

func (v NullableKafkaCatalogSpecSrInstance) Get() *KafkaCatalogSpecSrInstance {
	return v.value
}

func (v *NullableKafkaCatalogSpecSrInstance) Set(val *KafkaCatalogSpecSrInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableKafkaCatalogSpecSrInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableKafkaCatalogSpecSrInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKafkaCatalogSpecSrInstance(val *KafkaCatalogSpecSrInstance) *NullableKafkaCatalogSpecSrInstance {
	return &NullableKafkaCatalogSpecSrInstance{value: val, isSet: true}
}

func (v NullableKafkaCatalogSpecSrInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKafkaCatalogSpecSrInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


