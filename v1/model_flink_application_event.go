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

// FlinkApplicationEvent Events from the deployment of Flink clusters
type FlinkApplicationEvent struct {
	ApiVersion string `json:"apiVersion"`
	Kind string `json:"kind"`
	Metadata EventMetadata `json:"metadata"`
	Status EventStatus `json:"status"`
}

// NewFlinkApplicationEvent instantiates a new FlinkApplicationEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlinkApplicationEvent(apiVersion string, kind string, metadata EventMetadata, status EventStatus) *FlinkApplicationEvent {
	this := FlinkApplicationEvent{}
	this.ApiVersion = apiVersion
	this.Kind = kind
	this.Metadata = metadata
	this.Status = status
	return &this
}

// NewFlinkApplicationEventWithDefaults instantiates a new FlinkApplicationEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlinkApplicationEventWithDefaults() *FlinkApplicationEvent {
	this := FlinkApplicationEvent{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *FlinkApplicationEvent) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *FlinkApplicationEvent) GetApiVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *FlinkApplicationEvent) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetKind returns the Kind field value
func (o *FlinkApplicationEvent) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *FlinkApplicationEvent) GetKindOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *FlinkApplicationEvent) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *FlinkApplicationEvent) GetMetadata() EventMetadata {
	if o == nil {
		var ret EventMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *FlinkApplicationEvent) GetMetadataOk() (*EventMetadata, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *FlinkApplicationEvent) SetMetadata(v EventMetadata) {
	o.Metadata = v
}

// GetStatus returns the Status field value
func (o *FlinkApplicationEvent) GetStatus() EventStatus {
	if o == nil {
		var ret EventStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *FlinkApplicationEvent) GetStatusOk() (*EventStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *FlinkApplicationEvent) SetStatus(v EventStatus) {
	o.Status = v
}

func (o FlinkApplicationEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableFlinkApplicationEvent struct {
	value *FlinkApplicationEvent
	isSet bool
}

func (v NullableFlinkApplicationEvent) Get() *FlinkApplicationEvent {
	return v.value
}

func (v *NullableFlinkApplicationEvent) Set(val *FlinkApplicationEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableFlinkApplicationEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableFlinkApplicationEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlinkApplicationEvent(val *FlinkApplicationEvent) *NullableFlinkApplicationEvent {
	return &NullableFlinkApplicationEvent{value: val, isSet: true}
}

func (v NullableFlinkApplicationEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlinkApplicationEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


