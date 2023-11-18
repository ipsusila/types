package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ipsusila/types/internal"
)

var _ internal.VariantE = Variant(nil)

// RVariant holds various value and handle type conversion.
//
// For conversion to time.Duration and time.Time,
// if value type is floating point, e.g. nnn.fff
// then nnn will be treated as seconds,
// while fff will be treated as nano seconds.
type RVariant interface {
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
	Interface() any

	Uint8() uint8
	Uint16() uint16
	Uint32() uint32
	Uint64() uint64
	Uint() uint
	Int8() int8
	Int16() int16
	Int32() int32
	Int64() int64
	Int() int
	Byte() byte
	Float32() float32
	Float64() float64
	Bool() bool
	String() string
	Duration() time.Duration
	Time() time.Time

	Value() (driver.Value, error)
	MarshalJSON() ([]byte, error)
}

// Variant is a variant that can be modified
type Variant interface {
	RVariant

	Set(v any) error
	UnmarshalJSON(data []byte) error
	Scan(src any) error
}

type variant struct {
	v internal.VariantE
}

// NewRVariant create readonly variant from given value
func NewRVariant(i interface{}) RVariant {
	if v, ok := i.(RVariant); ok {
		return v
	}
	return variant{v: internal.NewVariantE(i)}
}

// NewVariant create r/w variant from optionally given value
func NewVariant(i ...interface{}) Variant {
	if len(i) == 0 {
		return &variant{v: internal.NewVariantE(nil)}
	}

	// use first argument only
	if v, ok := i[0].(Variant); ok {
		return v
	}
	return &variant{v: internal.NewVariantE(i[0])}
}

func (v variant) IsNil() bool {
	return v.v == nil || v.v.IsNil()
}
func (v variant) IsZero() bool {
	return v.v.IsZero()
}
func (v variant) Interface() any {
	return v.v.Interface()
}
func (v variant) Uint8E() (uint8, error) {
	return v.v.Uint8E()

}
func (v variant) Uint16E() (uint16, error) {
	return v.v.Uint16E()

}
func (v variant) Uint32E() (uint32, error) {
	return v.v.Uint32E()
}
func (v variant) Uint64E() (uint64, error) {
	return v.v.Uint64E()
}
func (v variant) UintE() (uint, error) {
	return v.v.UintE()
}
func (v variant) Int8E() (int8, error) {
	return v.v.Int8E()
}
func (v variant) Int16E() (int16, error) {
	return v.v.Int16E()
}
func (v variant) Int32E() (int32, error) {
	return v.v.Int32E()
}
func (v variant) Int64E() (int64, error) {
	return v.v.Int64E()
}
func (v variant) IntE() (int, error) {
	return v.v.IntE()
}
func (v variant) ByteE() (byte, error) {
	return v.v.ByteE()
}
func (v variant) Float32E() (float32, error) {
	return v.v.Float32E()
}
func (v variant) Float64E() (float64, error) {
	return v.v.Float64E()
}
func (v variant) BoolE() (bool, error) {
	return v.v.BoolE()
}
func (v variant) StringE() (string, error) {
	return v.v.StringE()
}
func (v variant) DurationE() (time.Duration, error) {
	return v.v.DurationE()
}
func (v variant) TimeE() (time.Time, error) {
	return v.v.TimeE()
}

func (v variant) Uint8() uint8 {
	u, _ := v.v.Uint8E()
	return u
}
func (v variant) Uint16() uint16 {
	u, _ := v.v.Uint16E()
	return u
}
func (v variant) Uint32() uint32 {
	u, _ := v.v.Uint32E()
	return u
}
func (v variant) Uint64() uint64 {
	u, _ := v.v.Uint64E()
	return u
}
func (v variant) Uint() uint {
	u, _ := v.v.UintE()
	return u
}
func (v variant) Int8() int8 {
	i, _ := v.v.Int8E()
	return i
}
func (v variant) Int16() int16 {
	i, _ := v.v.Int16E()
	return i
}
func (v variant) Int32() int32 {
	i, _ := v.v.Int32E()
	return i
}
func (v variant) Int64() int64 {
	i, _ := v.v.Int64E()
	return i
}
func (v variant) Int() int {
	i, _ := v.v.IntE()
	return i
}
func (v variant) Byte() byte {
	b, _ := v.v.ByteE()
	return b
}
func (v variant) Float32() float32 {
	f, _ := v.v.Float32E()
	return f
}
func (v variant) Float64() float64 {
	f, _ := v.v.Float64E()
	return f
}
func (v variant) Bool() bool {
	b, _ := v.v.BoolE()
	return b
}
func (v variant) String() string {
	s, _ := v.v.StringE()
	return s
}
func (v variant) Duration() time.Duration {
	d, _ := v.v.DurationE()
	return d
}
func (v variant) Time() time.Time {
	t, _ := v.v.TimeE()
	return t
}
func (v variant) Value() (driver.Value, error) {
	return v.v.Value()
}
func (v variant) MarshalJSON() ([]byte, error) {
	return v.v.MarshalJSON()
}

func (v *variant) Set(value any) error {
	v.v = internal.NewVariantE(value)
	if value != nil && v.v.IsNil() {
		return fmt.Errorf("unsupported variant type %T -> %v", value, value)
	}
	return nil
}
func (v *variant) Scan(src any) error {
	return v.Set(src)
}

func (v *variant) UnmarshalJSON(data []byte) error {
	var dest any
	if err := json.Unmarshal(data, &dest); err != nil {
		return err
	}
	return v.Set(dest)
}
