package internal

import (
	"database/sql/driver"
	"math"
	"strconv"
	"time"

	"golang.org/x/exp/constraints"
)

type unsignedVariant[T constraints.Unsigned] struct {
	val T
}

func (u unsignedVariant[T]) IsNil() bool {
	return false
}
func (u unsignedVariant[T]) IsZero() bool {
	return u.val == 0
}
func (u unsignedVariant[T]) Interface() any {
	return u.val
}

func (u unsignedVariant[T]) Uint8E() (uint8, error) {
	if uint64(u.val) > math.MaxUint8 {
		return uint8(u.val), errOutOfRange
	}
	return uint8(u.val), nil
}

func (u unsignedVariant[T]) Uint16E() (uint16, error) {
	if uint64(u.val) > math.MaxUint16 {
		return uint16(u.val), errOutOfRange
	}
	return uint16(u.val), nil
}
func (u unsignedVariant[T]) Uint32E() (uint32, error) {
	if uint64(u.val) > math.MaxUint32 {
		return uint32(u.val), errOutOfRange
	}
	return uint32(u.val), nil
}
func (u unsignedVariant[T]) Uint64E() (uint64, error) {
	return uint64(u.val), nil
}
func (u unsignedVariant[T]) UintE() (uint, error) {
	if uint64(u.val) > math.MaxUint {
		return uint(u.val), errOutOfRange
	}
	return uint(u.val), nil
}

func (u unsignedVariant[T]) Int8E() (int8, error) {
	if uint64(u.val) > math.MaxInt8 {
		return int8(u.val), errOutOfRange
	}
	return int8(u.val), nil
}

func (u unsignedVariant[T]) Int16E() (int16, error) {
	if uint64(u.val) > math.MaxInt16 {
		return int16(u.val), errOutOfRange
	}
	return int16(u.val), nil
}
func (u unsignedVariant[T]) Int32E() (int32, error) {
	if uint64(u.val) > math.MaxInt32 {
		return int32(u.val), errOutOfRange
	}
	return int32(u.val), nil
}
func (u unsignedVariant[T]) Int64E() (int64, error) {
	if uint64(u.val) > math.MaxInt64 {
		return int64(u.val), errOutOfRange
	}
	return int64(u.val), nil
}
func (u unsignedVariant[T]) IntE() (int, error) {
	if int64(u.val) > math.MaxInt {
		return int(u.val), errOutOfRange
	}
	return int(u.val), nil
}
func (u unsignedVariant[T]) ByteE() (byte, error) {
	if int64(u.val) > math.MaxUint8 {
		return byte(u.val), errOutOfRange
	}
	return byte(u.val), nil
}
func (u unsignedVariant[T]) Float32E() (float32, error) {
	return float32(u.val), nil
}
func (u unsignedVariant[T]) Float64E() (float64, error) {
	return float64(u.val), nil
}
func (u unsignedVariant[T]) BoolE() (bool, error) {
	if u.val == 0 {
		return false, nil
	}
	return true, nil
}
func (u unsignedVariant[T]) StringE() (string, error) {
	return strconv.FormatUint(uint64(u.val), 10), nil
}
func (u unsignedVariant[T]) DurationE() (time.Duration, error) {
	return time.Duration(u.val), nil
}
func (u unsignedVariant[T]) TimeE() (time.Time, error) {
	i64, err := u.Int64E()
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(i64, 0), nil
}
func (u unsignedVariant[T]) Value() (driver.Value, error) {
	i64, err := u.Int64E()
	if err != nil {
		return nil, err
	}
	return driver.Value(i64), nil
}
func (u unsignedVariant[T]) MarshalJSON() ([]byte, error) {
	return strconv.AppendUint(nil, uint64(u.val), 10), nil
}
