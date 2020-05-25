package goform

import (
	"fmt"
	"reflect"
)

type Form struct {
	Values interface{}
	Errors interface{}
	Err    error
}

func New() *Form {
	return &Form{}
}

func (f *Form) GetValue(path ...interface{}) interface{} {
	v := getValue(reflect.ValueOf(f.Values), path)
	if v.IsValid() && v.CanInterface() {
		return v.Interface()
	} else {
		return nil
	}
}

func (f *Form) SetValue(value interface{}, path ...interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if err, _ = r.(error); err == nil {
				err = fmt.Errorf("%v", r)
			}
			if f.Err == nil {
				f.Err = err
			}
		}
	}()
	f.Values = setValue(reflect.ValueOf(f.Values), path, reflect.ValueOf(value)).Interface()
	return
}

func (f *Form) SetError(error string, path ...interface{}) {
	f.Errors = setError(reflect.ValueOf(f.Errors), path, reflect.ValueOf(error)).Interface()
}

func (f *Form) HasError() bool {
	return f.Errors != nil
}

func (f *Form) ResetError() {
	f.Errors = nil
}
