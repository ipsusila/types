package internal

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

type VariantE interface {
	IsNil() bool
	IsZero() bool
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
	StringE() (string, error)
	DurationE() (time.Duration, error)
	TimeE() (time.Time, error)
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

// newVariantReflectE create variant from given value
func newVariantReflectE(i interface{}) VariantE {
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Int8:
		return signedVariant[int8]{val: int8(v.Int())}
	case reflect.Int16:
		return signedVariant[int16]{val: int16(v.Int())}
	case reflect.Int32:
		return signedVariant[int32]{val: int32(v.Int())}
	case reflect.Int64:
		return signedVariant[int64]{val: v.Int()}
	case reflect.Int:
		return signedVariant[int]{val: int(v.Int())}
	case reflect.Uint8:
		return unsignedVariant[uint8]{val: uint8(v.Uint())}
	case reflect.Uint16:
		return unsignedVariant[uint16]{val: uint16(v.Uint())}
	case reflect.Uint32:
		return unsignedVariant[uint32]{val: uint32(v.Uint())}
	case reflect.Uint64:
		return unsignedVariant[uint64]{val: v.Uint()}
	case reflect.Uint:
		return unsignedVariant[uint]{val: uint(v.Uint())}
	case reflect.Float32:
		return floatVariant[float32]{val: float32(v.Float()), bitSize: 32}
	case reflect.Float64:
		return floatVariant[float64]{val: v.Float(), bitSize: 64}
	default:
		if s, ok := i.(fmt.Stringer); ok {
			return stringVariant{val: s.String()}
		} else if e, ok := i.(error); ok {
			return stringVariant{val: e.Error()}
		}
	}
	return nilVariant{}
}

// NewVariantE create variant from given value
func NewVariantE(i interface{}) VariantE {
	i = indirect(i)
	if i == nil {
		return nilVariant{}
	}

	// check types
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
	case uint8:
		return unsignedVariant[uint8]{val: v}
	case uint16:
		return unsignedVariant[uint16]{val: v}
	case uint32:
		return unsignedVariant[uint32]{val: v}
	case uint64:
		return unsignedVariant[uint64]{val: v}
	case uint:
		return unsignedVariant[uint]{val: v}
	case float32:
		return floatVariant[float32]{val: v, bitSize: 32}
	case float64:
		return floatVariant[float64]{val: v, bitSize: 64}
	case bool:
		return boolVariant{val: v}
	case string:
		return stringVariant{val: v}
	case time.Time:
		return timeVariant(v)
	case json.Number:
		return stringVariant{val: string(v)}
	case []byte:
		return stringVariant{val: string(v)}
	default:
		return newVariantReflectE(i)
	}
}
