package v1

import (
	"bytes"
	"encoding/json"
)

// UnmarshalJSON is a custom unmarshaler that makes ComputePool.Status
// deserialize correctly against the real CMF API response.
//
// The generated type declares Status as *map[string]map[string]interface{}
// (due to an inline OpenAPI schema with additionalProperties: true), but the
// API returns a flat object like {"phase":"RUNNING","message":null}. The
// standard encoding/json decode fails on that shape because scalar values
// cannot fit into map[string]interface{}.
//
// To conform to the declared type, each scalar (or null) value in the status
// object is wrapped as {"value": <value>}; values that are already objects
// are passed through unchanged to avoid double-wrapping. Callers can then
// read fields uniformly via pool.GetStatus()["<key>"]["value"] for scalars,
// or traverse the nested object directly.
//
// Status shapes that are neither null nor a JSON object (e.g. an array,
// string, or number at the top level) are treated as absent — Status is set
// to nil rather than failing the whole response deserialization, since the
// whole point of this unmarshaler is to keep Execute() succeeding on valid
// 200 responses.
//
// TODO: remove this file and its entries in v1/.openapi-generator-ignore
// once the OpenAPI spec references ComputePoolStatus via $ref instead of
// defining status inline.
func (o *ComputePool) UnmarshalJSON(data []byte) error {
	// Alias prevents infinite recursion: the aliased type has no methods,
	// so json.Unmarshal uses reflection-based decoding for the embedded struct.
	type Alias ComputePool
	aux := struct {
		Status json.RawMessage `json:"status"`
		*Alias
	}{
		Alias: (*Alias)(o),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	// Reset Status so re-used receivers don't retain stale values when the
	// incoming payload has status null or missing.
	o.Status = nil
	if len(aux.Status) == 0 || bytes.Equal(aux.Status, []byte("null")) {
		return nil
	}
	// The status must decode as a JSON object. Anything else (array, scalar)
	// is treated as absent.
	var flat map[string]interface{}
	if err := json.Unmarshal(aux.Status, &flat); err != nil {
		return nil
	}
	wrapped := make(map[string]map[string]interface{}, len(flat))
	for k, v := range flat {
		if m, ok := v.(map[string]interface{}); ok {
			// Value is already an object — matches the declared inner type, pass through.
			wrapped[k] = m
		} else {
			// Scalar or null — wrap so the result conforms to the declared type.
			wrapped[k] = map[string]interface{}{"value": v}
		}
	}
	o.Status = &wrapped
	return nil
}
