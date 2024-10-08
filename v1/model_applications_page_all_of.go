/*
 * Confluent Manager for Apache Flink / CMF
 *
 * Apache Flink job lifecycle management component for Confluent Platform.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1
// ApplicationsPageAllOf struct for ApplicationsPageAllOf
type ApplicationsPageAllOf struct {
	Metadata ApplicationPageMetadata `json:"metadata,omitempty"`
	Items []Application `json:"items,omitempty"`
}
