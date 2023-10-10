package internal

import (
	"strconv"
	"time"
)

type boolVariant struct {
	val bool
}

func (v boolVariant) IsNil() bool {
	return false
}
func (v boolVariant) IsZero() bool {
	return v.val == false
}
func (v boolVariant) Uint8E() (uint8, error) {
	if v.val {
		return 1, nil
	}
	return 0, nil
}
func (v boolVariant) Uint16E() (uint16, error) {
	if v.val {
		return 1, nil
	}
	return 0, nil
}
func (v boolVariant) Uint32E() (uint32, error) {
	if v.val {
		return 1, nil
	}
	return 0, nil
}
func (v boolVariant) Uint64E() (uint64, error) {
	if v.val {
		return 1, nil
	}
	return 0, nil
}
func (v boolVariant) UintE() (uint, error) {
	if v.val {
		return 1, nil
	}
	return 0, nil
}
func (v boolVariant) Int8E() (int8, error) {
	if v.val {
		return 1, nil
	}
	return 0, nil
}
func (v boolVariant) Int16E() (int16, error) {
	if v.val {
		return 1, nil
	}
	return 0, nil
}
func (v boolVariant) Int32E() (int32, error) {
	if v.val {
		return 1, nil
	}
	return 0, nil
}
func (v boolVariant) Int64E() (int64, error) {
	if v.val {
		return 1, nil
	}
	return 0, nil
}
func (v boolVariant) IntE() (int, error) {
	if v.val {
		return 1, nil
	}
	return 0, nil
}
func (v boolVariant) ByteE() (byte, error) {
	if v.val {
		return 1, nil
	}
	return 0, nil
}
func (v boolVariant) Float32E() (float32, error) {
	if v.val {
		return 1, nil
	}
	return 0, nil
}
func (v boolVariant) Float64E() (float64, error) {
	if v.val {
		return 1, nil
	}
	return 0, nil
}
func (v boolVariant) BoolE() (bool, error) {
	return v.val, nil
}
func (v boolVariant) StringE() (string, error) {
	return strconv.FormatBool(v.val), nil
}
func (v boolVariant) DurationE() (time.Duration, error) {
	return 0, errCannotConvert
}
func (v boolVariant) TimeE() (time.Time, error) {
	return time.Time{}, errCannotConvert
}
