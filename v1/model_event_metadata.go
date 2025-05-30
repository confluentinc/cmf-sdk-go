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

// EventMetadata Metadata about the event
type EventMetadata struct {
	// Name of the Event
	Name *string `json:"name,omitempty"`
	// Unique identifier of the Event. Identical to name.
	Uid *string `json:"uid,omitempty"`
	// Timestamp when the Event was created
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`
	// Name of the FlinkApplicationInstance which this event is related to
	FlinkApplicationInstance *string `json:"flinkApplicationInstance,omitempty"`
	// Labels of the Event
	Labels *map[string]string `json:"labels,omitempty"`
	// Annotations of the Event
	Annotations *map[string]string `json:"annotations,omitempty"`
}

// NewEventMetadata instantiates a new EventMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventMetadata() *EventMetadata {
	this := EventMetadata{}
	return &this
}

// NewEventMetadataWithDefaults instantiates a new EventMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventMetadataWithDefaults() *EventMetadata {
	this := EventMetadata{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EventMetadata) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventMetadata) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EventMetadata) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EventMetadata) SetName(v string) {
	o.Name = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *EventMetadata) GetUid() string {
	if o == nil || o.Uid == nil {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventMetadata) GetUidOk() (*string, bool) {
	if o == nil || o.Uid == nil {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *EventMetadata) HasUid() bool {
	if o != nil && o.Uid != nil {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *EventMetadata) SetUid(v string) {
	o.Uid = &v
}

// GetCreationTimestamp returns the CreationTimestamp field value if set, zero value otherwise.
func (o *EventMetadata) GetCreationTimestamp() string {
	if o == nil || o.CreationTimestamp == nil {
		var ret string
		return ret
	}
	return *o.CreationTimestamp
}

// GetCreationTimestampOk returns a tuple with the CreationTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventMetadata) GetCreationTimestampOk() (*string, bool) {
	if o == nil || o.CreationTimestamp == nil {
		return nil, false
	}
	return o.CreationTimestamp, true
}

// HasCreationTimestamp returns a boolean if a field has been set.
func (o *EventMetadata) HasCreationTimestamp() bool {
	if o != nil && o.CreationTimestamp != nil {
		return true
	}

	return false
}

// SetCreationTimestamp gets a reference to the given string and assigns it to the CreationTimestamp field.
func (o *EventMetadata) SetCreationTimestamp(v string) {
	o.CreationTimestamp = &v
}

// GetFlinkApplicationInstance returns the FlinkApplicationInstance field value if set, zero value otherwise.
func (o *EventMetadata) GetFlinkApplicationInstance() string {
	if o == nil || o.FlinkApplicationInstance == nil {
		var ret string
		return ret
	}
	return *o.FlinkApplicationInstance
}

// GetFlinkApplicationInstanceOk returns a tuple with the FlinkApplicationInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventMetadata) GetFlinkApplicationInstanceOk() (*string, bool) {
	if o == nil || o.FlinkApplicationInstance == nil {
		return nil, false
	}
	return o.FlinkApplicationInstance, true
}

// HasFlinkApplicationInstance returns a boolean if a field has been set.
func (o *EventMetadata) HasFlinkApplicationInstance() bool {
	if o != nil && o.FlinkApplicationInstance != nil {
		return true
	}

	return false
}

// SetFlinkApplicationInstance gets a reference to the given string and assigns it to the FlinkApplicationInstance field.
func (o *EventMetadata) SetFlinkApplicationInstance(v string) {
	o.FlinkApplicationInstance = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *EventMetadata) GetLabels() map[string]string {
	if o == nil || o.Labels == nil {
		var ret map[string]string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventMetadata) GetLabelsOk() (*map[string]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *EventMetadata) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]string and assigns it to the Labels field.
func (o *EventMetadata) SetLabels(v map[string]string) {
	o.Labels = &v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *EventMetadata) GetAnnotations() map[string]string {
	if o == nil || o.Annotations == nil {
		var ret map[string]string
		return ret
	}
	return *o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventMetadata) GetAnnotationsOk() (*map[string]string, bool) {
	if o == nil || o.Annotations == nil {
		return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *EventMetadata) HasAnnotations() bool {
	if o != nil && o.Annotations != nil {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given map[string]string and assigns it to the Annotations field.
func (o *EventMetadata) SetAnnotations(v map[string]string) {
	o.Annotations = &v
}

func (o EventMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Uid != nil {
		toSerialize["uid"] = o.Uid
	}
	if o.CreationTimestamp != nil {
		toSerialize["creationTimestamp"] = o.CreationTimestamp
	}
	if o.FlinkApplicationInstance != nil {
		toSerialize["flinkApplicationInstance"] = o.FlinkApplicationInstance
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.Annotations != nil {
		toSerialize["annotations"] = o.Annotations
	}
	return json.Marshal(toSerialize)
}

type NullableEventMetadata struct {
	value *EventMetadata
	isSet bool
}

func (v NullableEventMetadata) Get() *EventMetadata {
	return v.value
}

func (v *NullableEventMetadata) Set(val *EventMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableEventMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableEventMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventMetadata(val *EventMetadata) *NullableEventMetadata {
	return &NullableEventMetadata{value: val, isSet: true}
}

func (v NullableEventMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


