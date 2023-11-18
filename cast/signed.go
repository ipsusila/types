package cast

import (
	"math"

	"golang.org/x/exp/constraints"
)

func s2i8[T constraints.Signed](v T) (int8, error) {
	if v < math.MinInt8 || v > math.MaxInt8 {
		return int8(v), ErrOutOfRange
	}
	return int8(v), nil
}

func s2s[V constraints.Signed, T constraints.Signed](v V, mi, ma int64) (T, error) {
	x := int64(v)
	if x < mi || x > ma {
		return T(v), ErrOutOfRange
	}
	return T(v), nil
}
func s2u[V constraints.Signed, T constraints.Unsigned](v V, mi, ma int64) (T, error) {
	if v < 0 {
		return T(v), ErrNegativeNotAllowed
	}
	x := int64(v)
	if x > ma {
		return T(v), ErrOutOfRange
	}
	return T(v), nil
}
func u2s[V constraints.Unsigned, T constraints.Signed](v V, ma uint64) (T, error) {
	x := uint64(v)
	if x > ma {
		return T(v), ErrOutOfRange
	}
	return T(v), nil
}

func u2u[V constraints.Unsigned, T constraints.Unsigned](v V, ma uint64) (T, error) {
	x := uint64(v)
	if x > ma {
		return T(v), ErrOutOfRange
	}
	return T(v), nil
}

func ToInt8(v any) (int8, error) {
	switch v := v.(type) {
	case int8:
		return v, nil
	case int16:
		return s2s[int16, int8](v, math.MaxInt8, math.MaxInt8)
	case int32:
		return s2s[int32, int8](v, math.MaxInt8, math.MaxInt8)
	case int64:
		return s2s[int64, int8](v, math.MaxInt8, math.MaxInt8)
	case int:
		return s2s[int, int8](v, math.MaxInt8, math.MaxInt8)
	case uint8:
		return u2s[uint8, int8](v, math.MaxInt8)
	case uint16:
		return u2s[uint16, int8](v, math.MaxInt8)
	case uint32:
		return u2s[uint32, int8](v, math.MaxInt8)
	case uint64:
		return u2s[uint64, int8](v, math.MaxInt8)
	case uint:
		return u2s[uint, int8](v, math.MaxInt8)
	}
	return 0, nil
}
