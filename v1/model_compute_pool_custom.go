package v1

import (
	"bytes"
	"encoding/json"
)

// UnmarshalJSON handles ComputePool.Status when the API returns a flat object
// (e.g. {"phase":"RUNNING"}). The generated type *map[string]map[string]interface{}
// cannot hold scalar values directly, so scalars and nulls are wrapped as
// {"value": v}; object values pass through unchanged. Non-object shapes
// (array, scalar) yield a nil Status instead of failing the decode.
func (o *ComputePool) UnmarshalJSON(data []byte) error {
	// Alias strips methods so the inner decode doesn't re-enter UnmarshalJSON.
	type Alias ComputePool
	aux := struct {
		Status json.RawMessage `json:"status"`
		*Alias
	}{Alias: (*Alias)(o)}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	o.Status = nil // reset for receiver reuse; repopulated below on success
	if len(aux.Status) == 0 || bytes.Equal(aux.Status, []byte("null")) {
		return nil
	}
	var flat map[string]interface{}
	if err := json.Unmarshal(aux.Status, &flat); err != nil {
		return nil // not an object (array/scalar) — treat as absent
	}
	wrapped := make(map[string]map[string]interface{}, len(flat))
	for k, v := range flat {
		if m, ok := v.(map[string]interface{}); ok {
			wrapped[k] = m // object — already satisfies declared inner type
		} else {
			wrapped[k] = map[string]interface{}{"value": v} // scalar/null — wrap
		}
	}
	o.Status = &wrapped
	return nil
}
