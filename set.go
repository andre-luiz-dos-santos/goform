package goform

import (
	"reflect"
)

func (f *Form) Set(target interface{}, path ...interface{}) bool {
	t := reflect.ValueOf(target)
	if t.Kind() != reflect.Ptr {
		panic(&ErrTargetIsNotPointer{target})
	}
	for t.Kind() == reflect.Ptr || t.Kind() == reflect.Interface {
		t = t.Elem()
	}
	v := f.GetValue(path...)
	if v == nil {
		return false
	}
	vv := reflect.ValueOf(v)
	if t.Kind() != vv.Kind() {
		panic(&ErrTargetHasWrongKind{t, vv})
	}
	t.Set(vv)
	return true
}
