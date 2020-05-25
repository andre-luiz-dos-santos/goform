package goform

import (
	"encoding/json"
	"github.com/andre-luiz-dos-santos/goform/validate"
	"strings"
	"testing"
)

func TestForm_GetValue_UnknownKeys(t *testing.T) {
	f := New()
	for _, path := range [][]interface{}{
		{"unknown"},
		{"x", "y", "z"},
		{"x", 1, 2},
		{0},
		{0, 1, 2},
		{0, "x"},
	} {
		v := f.GetValue(path...)
		if v != nil {
			t.Fatalf("Form.GetValue(%#v) = %#v; want nil", path, v)
		}
	}
}

func TestForm_GetValue_Object(t *testing.T) {
	f := New()
	err := f.ReadJSON(strings.NewReader(`{ "str": "string", "num": 123456 }`))
	if err != nil {
		t.Fatalf("Form.ReadJSON failed: %v", err)
	}
	t.Run("string", func(t *testing.T) {
		p := []interface{}{"str"}
		w := "string"
		v := f.GetValue(p...)
		s, ok := v.(string)
		if !ok || s != w {
			t.Fatalf("Form.GetValue(%#v) = %#v; want %#v", p, v, w)
		}
	})
	t.Run("number", func(t *testing.T) {
		p := []interface{}{"num"}
		w := float64(123456)
		v := f.GetValue(p...)
		n, ok := v.(float64)
		if !ok || n != w {
			t.Fatalf("Form.GetValue(%#v) = %#v; want %#v", p, v, w)
		}
	})
}

func TestForm_GetValue_Slice(t *testing.T) {
	f := New()
	err := f.ReadJSON(strings.NewReader(`{ "str": [ "TEST" ] }`))
	if err != nil {
		t.Fatalf("Form.ReadJSON failed: %v", err)
	}
	p := []interface{}{"str", 0}
	w := "TEST"
	v := f.GetValue(p...)
	s, ok := v.(string)
	if !ok || s != w {
		t.Fatalf("Form.GetValue(%#v) = %#v; want %#v", p, v, w)
	}
}

func TestForm_SetValue_Slice(t *testing.T) {
	// Create form.
	w := "v1"
	f := New()
	f.Values = map[string]interface{}{"a": []string{w}}
	p := []interface{}{"a", 0}
	v := f.GetValue(p...)
	s, ok := v.(string)
	if !ok || s != w {
		t.Fatalf("Form.GetValue(%#v) = %#v; want %#v", p, v, w)
	}
	// Change value.
	w = "v2"
	err := f.SetValue(w, p...)
	if err != nil {
		t.Fatalf("Failed to Form.SetValue(%#v): %v", p, err)
	}
	v = f.GetValue(p...)
	s, ok = v.(string)
	if !ok || s != w {
		t.Fatalf("Form.GetValue(%#v) = %#v; want %#v", p, v, w)
	}
	if f.Err != nil {
		t.Fatalf("Form.Err = %v; want nil", f.Err)
	}
	// Invalid path.
	p = []interface{}{"a", 0, "v2"}
	err = f.SetValue(w, p...)
	if err == nil {
		t.Fatalf("Form.SetValue(%#v) = nil; want invalid path error", p)
	}
	if f.Err != err {
		t.Fatalf("Form.Err = %v; want Form.SetValue(%#v)'s return value: %#v", f.Err, p, err)
	}
}

func TestForm_Validate(t *testing.T) {
	// Create form.
	f := New()
	err := f.ReadJSON(strings.NewReader(`{ "nome": 1, "bairro": "Lima", "emails": [ { "email": 2 } ] }`))
	if err != nil {
		t.Fatalf("Form.ReadJSON failed: %v", err)
	}
	if f.HasError() {
		t.Fatalf("Form.HasError() = true; want false")
	}
	// Validate form.
	f.For("nome", "bairro").Run(validate.IsString{})
	f.In("emails", 0).For("email").Run(validate.IsString{})
	// Check errors.
	w := `{"emails":{"0":{"email":"Não é string"}},"nome":"Não é string"}`
	b, err := json.Marshal(f.Errors)
	if err != nil {
		t.Fatalf("Failed to json.Marshal Form.Errors: %v", err)
	}
	if string(b) != w {
		t.Fatalf("Form.Errors = %s; want %s", b, w)
	}
	if !f.HasError() {
		t.Fatalf("Form.HasError() = false; want true")
	}
}

func TestForm_Set(t *testing.T) {
	f := New()
	f.Values = map[string]interface{}{"a": []string{"b"}}
	t.Run("unknown path", func(t *testing.T) {
		target := "str"
		unknownPath := []interface{}{"a", 999999}
		f.Set(&target, unknownPath...)
		if target != "str" {
			t.Fatalf("Form.Set(target, unknownPath) changed target")
		}
	})
	t.Run("valid path", func(t *testing.T) {
		target := "str"
		validPath := []interface{}{"a", 0}
		f.Set(&target, validPath...)
		if target != "b" {
			t.Fatalf("Form.Set(target, validPath) did not set target")
		}
	})
}

func TestForm_Set_WrongType(t *testing.T) {
	defer func() { recover() }()
	f := New()
	f.Values = map[string]interface{}{"a": 123456}
	s := "str"
	f.Set(&s, "a")
	t.Fatalf("Form.Set(string, path to int) worked; want panic(wrong type)")
}
