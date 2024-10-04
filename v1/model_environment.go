/*
 * Confluent Manager for Apache Flink / CMF
 *
 * Apache Flink job lifecycle management component for Confluent Platform.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1
import (
	"time"
)
// Environment Environment
type Environment struct {
	// A unique name for the resource.
	Name string `json:"name"`
	CreatedTime time.Time `json:"created_time,omitempty"`
	UpdatedTime time.Time `json:"updated_time,omitempty"`
	FlinkApplicationDefaults map[string]interface{} `json:"flinkApplicationDefaults,omitempty"`
	KubernetesNamespace string `json:"kubernetesNamespace"`
}
