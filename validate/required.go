package validate

import (
	"reflect"
)

type Required struct {
}

func (r Required) Validate(f Form, path []interface{}) {
	v := f.GetValue(path...)
	if v == nil {
		f.SetError("É obrigatório", path...)
		return
	}
	vv := reflect.ValueOf(v)
	for vv.Kind() == reflect.Ptr || vv.Kind() == reflect.Interface {
		vv = vv.Elem()
	}
	switch vv.Kind() {
	case reflect.String, reflect.Slice, reflect.Array, reflect.Map:
		if vv.Len() == 0 {
			f.SetError("Não pode ser vazio", path...)
		}
	}
}
