package goform

import (
	"encoding/json"
	"github.com/andre-luiz-dos-santos/goform/validate"
	"testing"
)

func TestRunner_Range(t *testing.T) {
	f := New()
	f.Values = []string{
		"PASS",
		"", // FAIL
	}
	f.Range().Run(validate.Required{})
	if !f.HasError() {
		t.Fatalf("Form.HasError() = false; want true")
	}
	b, _ := json.Marshal(f.Errors)
	s := string(b)
	w := `{"1":"Não pode ser vazio"}`
	if s != w {
		t.Fatalf("Form.Errors = %v; want = %v", s, w)
	}
}

func TestRunner_For(t *testing.T) {
	f := New()
	f.Values = map[string]string{
		"a": "PASS",
		"b": "PASS",
		"c": "", // FAIL
		"d": "PASS",
	}
	f.For("a", "c").Run(validate.Required{})
	if !f.HasError() {
		t.Fatalf("Form.HasError() = false; want true")
	}
	b, _ := json.Marshal(f.Errors)
	s := string(b)
	w := `{"c":"Não pode ser vazio"}`
	if s != w {
		t.Fatalf("Form.Errors = %v; want = %v", s, w)
	}
}

func TestRunner_In(t *testing.T) {
	f := New()
	f.Values = map[string]interface{}{
		"a": map[string]interface{}{
			"b": map[string]interface{}{
				"c": []interface{}{
					map[string]interface{}{
						"k": "PASS",
					},
				},
			},
		},
	}
	f.In("a", "b", "c", 0, "k").Run(validate.Required{})
	if f.HasError() {
		t.Fatalf("Form.HasError() = true; want false")
	}
}
