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
// GetResourceBase struct for GetResourceBase
type GetResourceBase struct {
	CreatedTime time.Time `json:"created_time,omitempty"`
	UpdatedTime time.Time `json:"updated_time,omitempty"`
}
