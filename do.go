// Copyright 2020 Alexander Félix. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package overwrite

import (
	"errors"
	"fmt"
	"reflect"
)

// Error types of package overwrite
var (
	ErrNotSameType    = errors.New("Dst and Src are not the same type")
	ErrSrcNil         = errors.New("Input src type is nil")
	ErrDstNil         = errors.New("Input dst type is nil")
	ErrDstNotPtr      = errors.New("Input dst must be pointer")
	ErrSrcNotStruct   = errors.New("Input src must be struct")
	ErrCannotSetField = errors.New("Field not addressable and/or cannot be set")

	tagOverwrite = "overwrite"
)

// Do copies tagged fields of arguments <src> into <dst>
//
// dst needs to be a pointer to same type of struct that src is
// src needs to be passed as value and not a pointer
//
// Do traverses the value src recursively. If an encountered field is tagged with
// overwrite: true it tries to copy the value of that field into the dst counterpart field.
//
// Supported types
// Simple types (string, intX, uintX, floatX, boolean)
// Arrays, slices, maps of simple types
// Structs with simple supported types (recursively)
// Pointers are not supported atm
//
// The "omitempty" option specifies that the field should be omitted
// from the overwrite if the field has an empty value, defined as
// false, 0, and any empty array, slice, map, or string.
//
// Examples of struct field tags and their meanings:
//
//   // Field in dst will be overwritten by the value in src
//   Field int `overwrite:"true"`
//
//   // Field in dst will be overwritten by the value in src
//   // the field is omitted from the overwrite if its src value is empty,
//   // as defined above.
//   Field int `overwrite:"true,omitempty"`
//
//   // Field will not be overwritten, same as not setting tag
//   Field int `overwrite:"false"`
//
// Channel, complex, and function values cannot be overwritten.
// Attempting to overwrite such a value will be silently ignored.
func Do(dst, src interface{}) error {
	if err := checkInput(dst, src); err != nil {
		return err
	}

	t := reflect.TypeOf(src)
	vDst := reflect.ValueOf(dst)
	vSrc := reflect.ValueOf(src)

	// Iterate over all available fields and read the tag value
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if field.Type.Kind() == reflect.Struct {
			// Pass the address of the interface type to recursion so all supported values get set
			// in entire structure tree
			return Do(vDst.Elem().Field(i).Addr().Interface(), vSrc.Field(i).Interface())
		}

		// Skip kinds that are not supported
		if !checkSupportedKind(field.Type.Kind()) {
			continue
		}

		// Get the field tag value
		tag, err := newTags(field.Tag.Get(tagOverwrite))
		if err != nil {
			return err
		}

		// Ignore non tagged fields
		if !tag.overwrite {
			continue
		}

		if err := checkCanSet(vDst.Elem(), field.Name); err != nil {
			return err
		}

		if tag.omitempty && vSrc.Field(i).IsZero() {
			continue
		}

		// Overwrite value
		vDst.Elem().FieldByName(field.Name).Set(vSrc.Field(i))
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

func checkSupportedKind(kind reflect.Kind) bool {
	switch kind {
	case reflect.String, reflect.Bool, reflect.Slice, reflect.Array, reflect.Map,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Float32, reflect.Float64:
		return true
	default:
		return false
	}
}

func checkCanSet(elem reflect.Value, name string) error {
	value := elem.FieldByName(name)
	if !(value.IsValid() && value.CanAddr() && value.CanSet()) {
		return fmt.Errorf("Value (%v) Name (%s): %w", value, name, ErrCannotSetField)
	}
	return nil
}
