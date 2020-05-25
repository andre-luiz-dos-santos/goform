package goform

import (
	"fmt"
	"reflect"
)

type ErrInvalidValueStoreKind struct {
	Store reflect.Value
}

func (e *ErrInvalidValueStoreKind) Error() string {
	return fmt.Sprintf("value store is %v; must be map or slice", e.Store.Kind())
}

type ErrTargetIsNotPointer struct {
	Target interface{}
}

func (e *ErrTargetIsNotPointer) Error() string {
	return fmt.Sprintf("target type is %T; must be a pointer", e.Target)
}

type ErrTargetHasWrongKind struct {
	Target, Value reflect.Value
}

func (e *ErrTargetHasWrongKind) Error() string {
	return fmt.Sprintf("target type is %v; must be %v", e.Target.Kind(), e.Value.Kind())
}
