package overwrite

import (
	"errors"
	"reflect"
)

// Error types of package overwrite
var (
	ErrNotSameType = errors.New("Dst and Src are not the same type")
	ErrSrcNil      = errors.New("Input src type is nil")
	ErrDstNil      = errors.New("Input dst type is nil")
)

// Do write here
func Do(dst, src interface{}) error {
	if dst == nil {
		return ErrDstNil
	}

	if src == nil {
		return ErrSrcNil
	}

	if reflect.TypeOf(dst) != reflect.TypeOf(src) {
		return ErrNotSameType
	}

	return nil
}
