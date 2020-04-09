package overwrite

import (
	"errors"
	"reflect"
)

// Error types of package overwrite
var (
	ErrNotSameType  = errors.New("Dst and Src are not the same type")
	ErrSrcNil       = errors.New("Input src type is nil")
	ErrDstNil       = errors.New("Input dst type is nil")
	ErrDstNotPtr    = errors.New("Input dst must be pointer")
	ErrSrcNotStruct = errors.New("Input src must be struct")
)

// Do write here
func Do(dst, src interface{}) error {
	if err := checkInput(dst, src); err != nil {
		return err
	}

	return nil
}

func checkInput(dst, src interface{}) error {
	if dst == nil {
		return ErrDstNil
	}

	if src == nil {
		return ErrSrcNil
	}

	if reflect.TypeOf(dst).Kind() != reflect.Ptr {
		return ErrDstNotPtr
	}

	if reflect.TypeOf(src).Kind() != reflect.Struct {
		return ErrSrcNotStruct
	}

	if reflect.ValueOf(dst).Elem().Type().String() != reflect.TypeOf(src).String() {
		return ErrNotSameType
	}

	return nil
}
