package v1

import (
	"encoding/json"
	"testing"
)

// makeComputePoolJSON builds a ComputePool JSON payload with the given status literal.
// The status argument is inserted verbatim (no quoting), so callers pass raw JSON like
// `"RUNNING"`, `null`, `{"phase":"X"}`, `[]`, or `42`.
func makeComputePoolJSON(name, statusLiteral string) []byte {
	if statusLiteral == "" {
		return []byte(`{
			"apiVersion": "cmf.confluent.io/v1",
			"kind": "ComputePool",
			"metadata": {"name": "` + name + `"},
			"spec": {"type": "DEDICATED", "clusterSpec": {}}
		}`)
	}
	return []byte(`{
		"apiVersion": "cmf.confluent.io/v1",
		"kind": "ComputePool",
		"metadata": {"name": "` + name + `"},
		"spec": {"type": "DEDICATED", "clusterSpec": {}},
		"status": ` + statusLiteral + `
	}`)
}

// wrappedValue returns status[key]["value"] and whether both keys were found.
func wrappedValue(pool ComputePool, key string) (interface{}, bool) {
	if pool.Status == nil {
		return nil, false
	}
	inner, ok := (*pool.Status)[key]
	if !ok {
		return nil, false
	}
	v, ok := inner["value"]
	return v, ok
}

func TestComputePoolUnmarshal_FlatStatus(t *testing.T) {
	data := makeComputePoolJSON("pool-a", `{"phase":"RUNNING","message":null,"reason":"x","ready":true}`)

	var pool ComputePool
	if err := json.Unmarshal(data, &pool); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if pool.Metadata.Name != "pool-a" {
		t.Errorf("metadata.name = %q, want %q", pool.Metadata.Name, "pool-a")
	}
	if pool.Spec.Type != "DEDICATED" {
		t.Errorf("spec.type = %q, want %q", pool.Spec.Type, "DEDICATED")
	}
	if pool.Status == nil {
		t.Fatal("status is nil")
	}
	// Verify every flat key got wrapped, not just phase.
	wantValues := map[string]interface{}{
		"phase":   "RUNNING",
		"message": nil,
		"reason":  "x",
		"ready":   true,
	}
	status := *pool.Status
	if len(status) != len(wantValues) {
		t.Errorf("status has %d keys, want %d", len(status), len(wantValues))
	}
	for k, want := range wantValues {
		inner, ok := status[k]
		if !ok {
			t.Errorf("status[%q] missing", k)
			continue
		}
		// Literal-assert the wrapper key name to catch format drift.
		got, hasValue := inner["value"]
		if !hasValue {
			t.Errorf("status[%q] missing %q key; got %v", k, "value", inner)
			continue
		}
		if got != want {
			t.Errorf("status[%q][\"value\"] = %v, want %v", k, got, want)
		}
	}
}

func TestComputePoolUnmarshal_AlreadyNestedStatus(t *testing.T) {
	// Object values pass through unchanged, including extra keys (no re-wrap).
	data := makeComputePoolJSON("pool-b", `{"phase":{"value":"PENDING","extra":"meta"}}`)

	var pool ComputePool
	if err := json.Unmarshal(data, &pool); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	v, ok := wrappedValue(pool, "phase")
	if !ok {
		t.Fatal("phase.value missing")
	}
	if v != "PENDING" {
		t.Errorf("phase.value = %v, want %q", v, "PENDING")
	}
	if extra := (*pool.Status)["phase"]["extra"]; extra != "meta" {
		t.Errorf("phase.extra = %v, want %q (passthrough dropped keys)", extra, "meta")
	}
}

func TestComputePoolUnmarshal_PartialNestedStatus(t *testing.T) {
	// Mixed: some values are objects (already wrapped), some are scalars.
	// The object passes through unchanged; scalars are wrapped. No double-wrapping.
	data := makeComputePoolJSON("pool-c", `{"phase":{"value":"RUNNING"},"message":"ok"}`)

	var pool ComputePool
	if err := json.Unmarshal(data, &pool); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	phase, ok := wrappedValue(pool, "phase")
	if !ok {
		t.Fatal("phase.value missing")
	}
	if phase != "RUNNING" {
		t.Errorf("phase.value = %v, want %q (not double-wrapped)", phase, "RUNNING")
	}
	message, ok := wrappedValue(pool, "message")
	if !ok {
		t.Fatal("message.value missing")
	}
	if message != "ok" {
		t.Errorf("message.value = %v, want %q", message, "ok")
	}
}

func TestComputePoolUnmarshal_NullStatus(t *testing.T) {
	data := makeComputePoolJSON("pool-d", `null`)

	var pool ComputePool
	if err := json.Unmarshal(data, &pool); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if pool.Status != nil {
		t.Errorf("status = %v, want nil", pool.Status)
	}
}

func TestComputePoolUnmarshal_MissingStatus(t *testing.T) {
	data := makeComputePoolJSON("pool-e", "") // no status field at all

	var pool ComputePool
	if err := json.Unmarshal(data, &pool); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if pool.Status != nil {
		t.Errorf("status = %v, want nil", pool.Status)
	}
	if pool.Metadata.Name != "pool-e" {
		t.Errorf("metadata.name = %q, want %q", pool.Metadata.Name, "pool-e")
	}
}

func TestComputePoolUnmarshal_EmptyObjectStatus(t *testing.T) {
	data := makeComputePoolJSON("pool-f", `{}`)

	var pool ComputePool
	if err := json.Unmarshal(data, &pool); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if pool.Status == nil {
		t.Fatal("status is nil, want non-nil empty map")
	}
	if len(*pool.Status) != 0 {
		t.Errorf("status has %d keys, want 0", len(*pool.Status))
	}
}

func TestComputePoolUnmarshal_AllNullValues(t *testing.T) {
	data := makeComputePoolJSON("pool-g", `{"phase":null,"message":null}`)

	var pool ComputePool
	if err := json.Unmarshal(data, &pool); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	for _, k := range []string{"phase", "message"} {
		v, ok := wrappedValue(pool, k)
		if !ok {
			t.Errorf("%s.value missing", k)
			continue
		}
		if v != nil {
			t.Errorf("%s.value = %v, want nil", k, v)
		}
	}
}

func TestComputePoolUnmarshal_ArrayStatus(t *testing.T) {
	// Malformed status shape (array). Must NOT return an error; set Status = nil.
	data := makeComputePoolJSON("pool-h", `[1,2,3]`)

	var pool ComputePool
	if err := json.Unmarshal(data, &pool); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if pool.Status != nil {
		t.Errorf("status = %v, want nil for non-object status", pool.Status)
	}
	// Other fields must still populate.
	if pool.Metadata.Name != "pool-h" {
		t.Errorf("metadata.name = %q, want %q", pool.Metadata.Name, "pool-h")
	}
}

func TestComputePoolUnmarshal_ScalarStringStatus(t *testing.T) {
	data := makeComputePoolJSON("pool-i", `"RUNNING"`)

	var pool ComputePool
	if err := json.Unmarshal(data, &pool); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if pool.Status != nil {
		t.Errorf("status = %v, want nil for scalar status", pool.Status)
	}
}

func TestComputePoolUnmarshal_ScalarNumberStatus(t *testing.T) {
	data := makeComputePoolJSON("pool-j", `42`)

	var pool ComputePool
	if err := json.Unmarshal(data, &pool); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if pool.Status != nil {
		t.Errorf("status = %v, want nil for scalar status", pool.Status)
	}
}

func TestComputePoolUnmarshal_ReceiverReuse(t *testing.T) {
	// Unmarshal populates Status, then unmarshal a null-status payload into the
	// same receiver: Status must be reset to nil, not retain the previous value.
	var pool ComputePool
	first := makeComputePoolJSON("first", `{"phase":"RUNNING"}`)
	if err := json.Unmarshal(first, &pool); err != nil {
		t.Fatalf("first unmarshal: %v", err)
	}
	if pool.Status == nil {
		t.Fatal("status nil after first unmarshal")
	}
	second := makeComputePoolJSON("second", `null`)
	if err := json.Unmarshal(second, &pool); err != nil {
		t.Fatalf("second unmarshal: %v", err)
	}
	if pool.Status != nil {
		t.Errorf("status = %v after null payload, want nil", pool.Status)
	}
}

func TestComputePoolUnmarshal_PageResponse(t *testing.T) {
	// ComputePoolsPage.Items should invoke our UnmarshalJSON per element.
	data := []byte(`{
		"items": [
			{
				"apiVersion": "cmf.confluent.io/v1",
				"kind": "ComputePool",
				"metadata": {"name": "pool-1"},
				"spec": {"type": "DEDICATED", "clusterSpec": {}},
				"status": {"phase": "RUNNING"}
			},
			{
				"apiVersion": "cmf.confluent.io/v1",
				"kind": "ComputePool",
				"metadata": {"name": "pool-2"},
				"spec": {"type": "DEDICATED", "clusterSpec": {}},
				"status": {"phase": "PENDING"}
			}
		]
	}`)

	var page ComputePoolsPage
	if err := json.Unmarshal(data, &page); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	items := page.GetItems()
	if len(items) != 2 {
		t.Fatalf("got %d items, want 2", len(items))
	}
	wantPhases := []string{"RUNNING", "PENDING"}
	for i, item := range items {
		v, ok := wrappedValue(item, "phase")
		if !ok {
			t.Errorf("item %d: phase.value missing", i)
			continue
		}
		if v != wantPhases[i] {
			t.Errorf("item %d phase.value = %v, want %q", i, v, wantPhases[i])
		}
	}
}

func TestComputePoolUnmarshal_EmptyPageItems(t *testing.T) {
	// Empty and missing items must both yield a zero-length slice, not an error.
	for _, tc := range []struct {
		name string
		body string
	}{
		{"empty", `{"items":[]}`},
		{"missing", `{}`},
		{"null", `{"items":null}`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var page ComputePoolsPage
			if err := json.Unmarshal([]byte(tc.body), &page); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if n := len(page.GetItems()); n != 0 {
				t.Errorf("got %d items, want 0", n)
			}
		})
	}
}

func TestComputePoolUnmarshal_InvalidJSON(t *testing.T) {
	// Top-level invalid JSON must still surface an error (not silently passed).
	var pool ComputePool
	err := json.Unmarshal([]byte(`{not valid json`), &pool)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if _, isSyntax := err.(*json.SyntaxError); !isSyntax {
		t.Errorf("err type = %T, want *json.SyntaxError", err)
	}
}
