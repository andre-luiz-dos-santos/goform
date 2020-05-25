package goform

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestForm_SetValue_EmptyForm(t *testing.T) {
	for index, path := range [][]interface{}{
		{},
		{"path_1"},
		{"path_1", "path_2"},
		{"path_1", "path_2", "path_3"},
		{0},
		{0, 1},
		{0, 1, 2},
		{"path_1", 0},
		{0, "path_1"},
	} {
		t.Run(fmt.Sprintf("%v", index), func(t *testing.T) {
			t.Logf("> path %q", path)
			f := New()
			v := f.GetValue(path...)
			if v != nil {
				t.Fatalf("Form.GetValue returned %q; want nil", v)
			}
			err := f.SetValue("TEST", path...)
			if err != nil {
				t.Fatalf("Failed to Form.SetValue(%q): %v", path, err)
			}
			v = f.GetValue(path...)
			if v != "TEST" {
				t.Fatalf("Form.SetValue(%q) set wrong value: %q", path, v)
			}
			t.Logf("< values %q", f.Values)
		})
	}
}

func TestForm_SetValue_Replace(t *testing.T) {
	for index, path := range [][]interface{}{
		// Works.
		{},
		{"path_1"},
		// All below must return error and set Form.Err.
		{"path_1", "path_2"},
		{"path_1", "path_2", "path_3"},
		{0},
		{0, 1},
		{0, 1, 2},
		{"path_1", 0},
		{0, "path_1"},
	} {
		t.Run(fmt.Sprintf("%v", path), func(t *testing.T) {
			t.Logf("> path %q", path)
			f := New()
			f.Values = map[string]string{"path_1": "OLD"}
			err := f.SetValue("NEW", path...)
			v := f.GetValue(path...)
			if index < 2 {
				if err != nil {
					t.Fatalf("Failed to Form.SetValue: %v", err)
				}
				if f.Err != nil {
					t.Fatalf("Form.Err = %v; want nil", f.Err)
				}
				if v != "NEW" {
					t.Fatalf("Form.SetValue set value to %q; want `NEW`", v)
				}
			} else {
				if err == nil {
					t.Fatalf("Form.SetValue(wrong path type) must return error")
				}
				if f.Err == nil {
					t.Fatalf("Form.SetValue(wrong path type) must set Form.Err")
				}
				if v != nil {
					t.Fatalf("Form.GetValue(wrong path type) = %q; want nil", v)
				}
			}
			t.Logf("< values %q", f.Values)
		})
	}
}

func TestForm_SetError_EmptyForm(t *testing.T) {
	for _, d := range [][]interface{}{
		{`"TEST"`},
		{`{"path_1":"TEST"}`, "path_1"},
		{`{"path_1":{"path_2":"TEST"}}`, "path_1", "path_2"},
		{`{"path_1":{"path_2":{"path_3":"TEST"}}}`, "path_1", "path_2", "path_3"},
		{`{"0":"TEST"}`, 0},
		{`{"0":{"1":"TEST"}}`, 0, 1},
		{`{"0":{"1":{"2":"TEST"}}}`, 0, 1, 2},
		{`{"path_1":{"0":"TEST"}}`, "path_1", 0},
		{`{"0":{"path_1":"TEST"}}`, 0, "path_1"},
	} {
		t.Run(fmt.Sprintf("%v", d[1:]), func(t *testing.T) {
			want, path := d[0], d[1:]
			t.Logf("> path %q", path)
			f := New()
			f.SetError("TEST", path...)
			j, err := json.Marshal(f.Errors)
			if err != nil {
				t.Fatalf("Failed to json.Marshal(Form.Errors): %v", err)
			}
			s := string(j)
			if s != want {
				t.Fatalf("Form.Errors = %v; want %v", s, want)
			}
			t.Logf("< errors %q", f.Errors)
		})
	}
}

func TestForm_SetError_Replace(t *testing.T) {
	for _, d := range [][]interface{}{
		{`"NEW"`},
		{`{"a":"OLD"}`, "a"},
		{`{"a":"OLD"}`, "a", "b"},
		{`{"a":"OLD","b":"NEW"}`, "b"},
		{`{"a":"OLD","b":{"c":"NEW"}}`, "b", "c"},
	} {
		t.Run(fmt.Sprintf("%v", d[1:]), func(t *testing.T) {
			want, path := d[0], d[1:]
			t.Logf("> path %q", path)
			f := New()
			f.SetError("OLD", "a")
			f.SetError("NEW", path...)
			j, err := json.Marshal(f.Errors)
			if err != nil {
				t.Fatalf("Failed to json.Marshal(f.Errors): %v", err)
			}
			s := string(j)
			if s != want {
				t.Fatalf("Form.Errors = %v; want %v", s, want)
			}
			t.Logf("< errors %q", f.Errors)
		})
	}
}
