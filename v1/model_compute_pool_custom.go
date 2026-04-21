package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// UnmarshalJSON wraps every value in ComputePool.Status as {"value": v} so the
// CMF API's flat status object (e.g. {"phase":"RUNNING"}) fits the generated
// type *map[string]map[string]interface{}. Non-object status shapes surface a
// decode error rather than being silently dropped.
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
	o.Status = nil // reset for receiver reuse
	if len(aux.Status) == 0 || bytes.Equal(aux.Status, []byte("null")) {
		return nil
	}
	var flat map[string]interface{}
	if err := json.Unmarshal(aux.Status, &flat); err != nil {
		return fmt.Errorf("unmarshal ComputePool.Status: %w", err)
	}
	wrapped := make(map[string]map[string]interface{}, len(flat))
	for k, v := range flat {
		wrapped[k] = map[string]interface{}{"value": v}
	}
	o.Status = &wrapped
	return nil
}
