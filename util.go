package goform

import (
	"reflect"
)

var (
	intf     = []interface{}(nil)
	intfType = reflect.TypeOf(intf).Elem()
)

func getValue(store reflect.Value, path []interface{}) reflect.Value {
	if len(path) == 0 {
		return store
	}
	for store.Kind() == reflect.Ptr || store.Kind() == reflect.Interface {
		store = store.Elem()
	}
	if !store.IsValid() {
		return reflect.ValueOf(nil)
	}
	switch store.Kind() {
	case reflect.Map:
		k := reflect.ValueOf(path[0])
		if k.Kind() == store.Type().Key().Kind() {
			return getValue(store.MapIndex(k), path[1:])
		}
	case reflect.Slice:
		if i, ok := path[0].(int); ok {
			if i >= 0 && i < store.Len() {
				return getValue(store.Index(i), path[1:])
			}
		}
	}
	return reflect.ValueOf(nil)
}

func setValue(store reflect.Value, path []interface{}, value reflect.Value) reflect.Value {
	if len(path) == 0 {
		return value
	}
	for store.Kind() == reflect.Ptr || store.Kind() == reflect.Interface {
		store = store.Elem()
	}
	// Create the store map.
	if !store.IsValid() {
		k := reflect.ValueOf(path[0])
		store = reflect.MakeMap(reflect.MapOf(k.Type(), intfType))
	}
	// Set the value at path[0].
	switch store.Kind() {
	case reflect.Map:
		k := reflect.ValueOf(path[0])
		v := setValue(store.MapIndex(k), path[1:], value)
		store.SetMapIndex(k, v)
	case reflect.Slice:
		i := store.Index(path[0].(int))
		v := setValue(i, path[1:], value)
		i.Set(v)
	default:
		panic(&ErrInvalidValueStoreKind{store})
	}
	return store
}

func setError(store reflect.Value, path []interface{}, error reflect.Value) reflect.Value {
	for store.Kind() == reflect.Ptr || store.Kind() == reflect.Interface {
		store = store.Elem()
	}
	if len(path) == 0 {
		if store.Kind() == reflect.String {
			// Return the error message that's already set.
			return store
		} else {
			return error
		}
	}
	// Create the store map.
	k := reflect.ValueOf(path[0])
	if !store.IsValid() {
		store = reflect.MakeMap(reflect.MapOf(k.Type(), intfType))
	} else if store.Kind() != reflect.Map {
		// Return the error message that's already set.
		return store
	}
	// Set the error at path[0].
	v := setError(store.MapIndex(k), path[1:], error)
	store.SetMapIndex(k, v)
	return store
}
