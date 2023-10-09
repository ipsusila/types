package internal

import "reflect"

type VariantE interface {
	Uint8E() (uint8, error)
	Uint16E() (uint16, error)
	Uint32E() (uint32, error)
	Uint64E() (uint64, error)
	UintE() (uint, error)
	Int8E() (int8, error)
	Int16E() (int16, error)
	Int32E() (int32, error)
	Int64E() (int64, error)
	IntE() (int, error)
	ByteE() (byte, error)
	Float32E() (float32, error)
	Float64E() (float64, error)
	BoolE() (bool, error)
}

// indirect get pointed value as interface.
//
// From html/template/content.go
// Copyright 2011 The Go Authors. All rights reserved.
// indirect returns the value, after dereferencing as many times
// as necessary to reach the base type (or nil).
func indirect(i interface{}) interface{} {
	if i == nil {
		return nil
	}
	if t := reflect.TypeOf(i); t.Kind() != reflect.Ptr {
		return i
	}

	v := reflect.ValueOf(i)
	for v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

// NewVariantE create variant from given value
func NewVariantE(i interface{}) VariantE {
	i = indirect(i)
	switch v := i.(type) {
	case int8:
		return signedVariant[int8]{val: v}
	case int16:
		return signedVariant[int16]{val: v}
	case int32:
		return signedVariant[int32]{val: v}
	case int64:
		return signedVariant[int64]{val: v}
	case int:
		return signedVariant[int]{val: v}
	}

	return nil
}
