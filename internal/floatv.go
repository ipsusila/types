package internal

import (
	"math"
	"strconv"
	"time"

	"golang.org/x/exp/constraints"
)

type floatVariant[T constraints.Float] struct {
	val     T
	bitSize int
}

func (f floatVariant[T]) IsNil() bool {
	return false
}
func (f floatVariant[T]) IsZero() bool {
	return f.val == 0
}

func (f floatVariant[T]) Uint8E() (uint8, error) {
	if f.val < 0 {
		return 0, errNegativeNotAllowed
	}
	v := float64(f.val)
	t := math.Trunc(v)
	if v != t || v > math.MaxUint8 {
		return uint8(f.val), errCastDataLost
	}

	return uint8(f.val), nil
}

func (f floatVariant[T]) Uint16E() (uint16, error) {
	if f.val < 0 {
		return 0, errNegativeNotAllowed
	}
	v := float64(f.val)
	t := math.Trunc(v)
	if v != t || v > math.MaxUint16 {
		return uint16(f.val), errCastDataLost
	}
	return uint16(f.val), nil
}
func (f floatVariant[T]) Uint32E() (uint32, error) {
	if f.val < 0 {
		return 0, errNegativeNotAllowed
	}
	v := float64(f.val)
	t := math.Trunc(v)
	if v != t || v > math.MaxUint32 {
		return uint32(f.val), errCastDataLost
	}
	return uint32(f.val), nil
}
func (f floatVariant[T]) Uint64E() (uint64, error) {
	if f.val < 0 {
		return 0, errNegativeNotAllowed
	}
	v := float64(f.val)
	t := math.Trunc(v)
	if v != t {
		return uint64(f.val), errCastDataLost
	}
	return uint64(f.val), nil
}
func (f floatVariant[T]) UintE() (uint, error) {
	if f.val < 0 {
		return 0, errNegativeNotAllowed
	}
	v := float64(f.val)
	t := math.Trunc(v)
	if v != t || v > math.MaxUint {
		return uint(f.val), errCastDataLost
	}
	return uint(f.val), nil
}

func (f floatVariant[T]) Int8E() (int8, error) {
	v := float64(f.val)
	t := math.Trunc(v)
	if v != t || v < math.MinInt8 || v > math.MaxInt8 {
		return int8(f.val), errCastDataLost
	}
	return int8(f.val), nil
}

func (f floatVariant[T]) Int16E() (int16, error) {
	v := float64(f.val)
	t := math.Trunc(v)
	if v != t || v < math.MinInt16 || v > math.MaxInt16 {
		return int16(f.val), errCastDataLost
	}
	return int16(f.val), nil
}
func (f floatVariant[T]) Int32E() (int32, error) {
	v := float64(f.val)
	t := math.Trunc(v)
	if v != t || v < math.MinInt32 || v > math.MaxInt32 {
		return int32(f.val), errCastDataLost
	}
	return int32(f.val), nil
}
func (f floatVariant[T]) Int64E() (int64, error) {
	v := float64(f.val)
	t := math.Trunc(v)
	if v != t {
		return int64(f.val), errCastDataLost
	}
	return int64(f.val), nil
}
func (f floatVariant[T]) IntE() (int, error) {
	v := float64(f.val)
	t := math.Trunc(v)
	if v != t || v < math.MinInt || v > math.MaxInt {
		return int(f.val), errCastDataLost
	}
	return int(f.val), nil
}
func (f floatVariant[T]) ByteE() (byte, error) {
	if f.val < 0 {
		return 0, errNegativeNotAllowed
	}
	v := float64(f.val)
	t := math.Trunc(v)
	if v != t || v > math.MaxUint8 {
		return byte(f.val), errCastDataLost
	}

	return byte(f.val), nil
}
func (f floatVariant[T]) Float32E() (float32, error) {
	if f.bitSize == 64 {
		if f.val > math.MaxFloat32 {
			return float32(f.val), errCastDataLost
		}
	}
	return float32(f.val), nil
}
func (f floatVariant[T]) Float64E() (float64, error) {
	return float64(f.val), nil
}
func (f floatVariant[T]) BoolE() (bool, error) {
	if f.val == 0 {
		return false, nil
	}
	return true, nil
}
func (f floatVariant[T]) StringE() (string, error) {
	return strconv.FormatFloat(float64(f.val), 'f', -1, f.bitSize), nil
}
func (f floatVariant[T]) DurationE() (time.Duration, error) {
	sec, nano := f.secNano()
	return time.Duration(sec)*time.Second + time.Duration(nano), nil
}
func (f floatVariant[T]) TimeE() (time.Time, error) {
	sec, nano := f.secNano()
	return time.Unix(sec, nano), nil
}

func (f floatVariant[T]) secNano() (int64, int64) {
	f64 := float64(f.val)
	sec := math.Trunc(f64)
	nano := (f64 - sec) * 1000_000_000

	return int64(sec), int64(nano)
}
