package internal

import (
	"database/sql/driver"
	"math"
	"strconv"
	"time"

	"golang.org/x/exp/constraints"
)

type signedVariant[T constraints.Signed] struct {
	val T
}

func (s signedVariant[T]) IsNil() bool {
	return false
}
func (s signedVariant[T]) IsZero() bool {
	return s.val == 0
}
func (s signedVariant[T]) Interface() any {
	return s.val
}

func (s signedVariant[T]) Uint8E() (uint8, error) {
	if s.val < 0 {
		return 0, errNegativeNotAllowed
	}
	if uint64(s.val) > math.MaxUint8 {
		return uint8(s.val), errCastDataLost
	}

	return uint8(s.val), nil
}

func (s signedVariant[T]) Uint16E() (uint16, error) {
	if s.val < 0 {
		return 0, errNegativeNotAllowed
	}
	if uint64(s.val) > math.MaxUint16 {
		return uint16(s.val), errCastDataLost
	}
	return uint16(s.val), nil
}
func (s signedVariant[T]) Uint32E() (uint32, error) {
	if s.val < 0 {
		return 0, errNegativeNotAllowed
	}
	if uint64(s.val) > math.MaxUint32 {
		return uint32(s.val), errCastDataLost
	}
	return uint32(s.val), nil
}
func (s signedVariant[T]) Uint64E() (uint64, error) {
	if s.val < 0 {
		return 0, errNegativeNotAllowed
	}
	return uint64(s.val), nil
}
func (s signedVariant[T]) UintE() (uint, error) {
	if s.val < 0 {
		return 0, errNegativeNotAllowed
	}
	if uint64(s.val) > math.MaxUint {
		return uint(s.val), errCastDataLost
	}
	return uint(s.val), nil
}

func (s signedVariant[T]) Int8E() (int8, error) {
	v := int64(s.val)
	if v < math.MinInt8 || v > math.MaxInt8 {
		return int8(s.val), errCastDataLost
	}
	return int8(s.val), nil
}

func (s signedVariant[T]) Int16E() (int16, error) {
	v := int64(s.val)
	if v < math.MinInt16 || v > math.MaxInt16 {
		return int16(s.val), errCastDataLost
	}
	return int16(s.val), nil
}
func (s signedVariant[T]) Int32E() (int32, error) {
	v := int64(s.val)
	if v < math.MinInt32 || v > math.MaxInt32 {
		return int32(s.val), errCastDataLost
	}
	return int32(s.val), nil
}
func (s signedVariant[T]) Int64E() (int64, error) {
	return int64(s.val), nil
}
func (s signedVariant[T]) IntE() (int, error) {
	v := int64(s.val)
	if v < math.MinInt || v > math.MaxInt {
		return int(s.val), errCastDataLost
	}
	return int(s.val), nil
}
func (s signedVariant[T]) ByteE() (byte, error) {
	if s.val < 0 {
		return 0, errNegativeNotAllowed
	}
	if uint64(s.val) > math.MaxUint8 {
		return byte(s.val), errCastDataLost
	}
	return byte(s.val), nil
}
func (s signedVariant[T]) Float32E() (float32, error) {
	return float32(s.val), nil
}
func (s signedVariant[T]) Float64E() (float64, error) {
	return float64(s.val), nil
}
func (s signedVariant[T]) BoolE() (bool, error) {
	if s.val == 0 {
		return false, nil
	}
	return true, nil
}
func (s signedVariant[T]) StringE() (string, error) {
	return strconv.FormatInt(int64(s.val), 10), nil
}
func (s signedVariant[T]) DurationE() (time.Duration, error) {
	return time.Duration(s.val), nil
}
func (s signedVariant[T]) TimeE() (time.Time, error) {
	i64, err := s.Int64E()
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(i64, 0), nil
}
func (s signedVariant[T]) Value() (driver.Value, error) {
	return driver.Value(int64(s.val)), nil
}
func (s signedVariant[T]) MarshalJSON() ([]byte, error) {
	i64, err := s.Int64E()
	if err != nil {
		return nil, err
	}
	return strconv.AppendInt(nil, i64, 10), nil
}
