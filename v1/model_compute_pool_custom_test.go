package v1

import (
	"encoding/json"
	"strings"
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

// Non-object status shapes (array, top-level scalar) must surface as decode
// errors rather than be silently dropped — the backend never returns these
// shapes, so they indicate an API regression worth flagging.
func TestComputePoolUnmarshal_NonObjectStatusFails(t *testing.T) {
	cases := []struct {
		name   string
		status string
	}{
		{"array", `[1,2,3]`},
		{"string", `"RUNNING"`},
		{"number", `42`},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			data := makeComputePoolJSON("pool-"+tc.name, tc.status)
			var pool ComputePool
			err := json.Unmarshal(data, &pool)
			if err == nil {
				t.Fatalf("expected error for %s status, got nil", tc.name)
			}
			if !strings.Contains(err.Error(), "ComputePool.Status") {
				t.Errorf("err = %q, want it to mention ComputePool.Status", err)
			}
		})
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
