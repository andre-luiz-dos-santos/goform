package validate

import (
	"fmt"
	"reflect"
)

type Min struct {
	Size int
}

func (r Min) Validate(f Form, path []interface{}) {
	v := f.GetValue(path...)
	if v == nil {
		return
	}
	vv := reflect.ValueOf(v)
	for vv.Kind() == reflect.Ptr || vv.Kind() == reflect.Interface {
		vv = vv.Elem()
	}
	switch vv.Kind() {
	case reflect.String:
		if vv.Len() < r.Size {
			e := fmt.Sprintf("Deve ter mais de %d caracteres", r.Size)
			f.SetError(e, path...)
		}
	case reflect.Slice, reflect.Array, reflect.Map:
		if vv.Len() < r.Size {
			e := fmt.Sprintf("Deve ter mais de %d itens", r.Size)
			f.SetError(e, path...)
		}
	}
}
