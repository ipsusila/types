package internal

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

type stringVariant struct {
	val string
}

func (s stringVariant) IsNil() bool {
	return false
}
func (s stringVariant) IsZero() bool {
	return s.val == ""
}
func (s stringVariant) Interface() any {
	return s.val
}
func (s stringVariant) Uint8E() (uint8, error) {
	u, err := strconv.ParseUint(s.val, 0, 8)
	return uint8(u), err
}
func (s stringVariant) Uint16E() (uint16, error) {
	u, err := strconv.ParseUint(s.val, 0, 16)
	return uint16(u), err
}
func (s stringVariant) Uint32E() (uint32, error) {
	u, err := strconv.ParseUint(s.val, 0, 32)
	return uint32(u), err
}
func (s stringVariant) Uint64E() (uint64, error) {
	return strconv.ParseUint(s.val, 0, 64)
}
func (s stringVariant) UintE() (uint, error) {
	u, err := strconv.ParseUint(s.val, 0, 0)
	return uint(u), err
}
func (s stringVariant) Int8E() (int8, error) {
	i, err := strconv.ParseInt(s.val, 0, 8)
	return int8(i), err
}
func (s stringVariant) Int16E() (int16, error) {
	i, err := strconv.ParseInt(s.val, 0, 16)
	return int16(i), err
}
func (s stringVariant) Int32E() (int32, error) {
	i, err := strconv.ParseInt(s.val, 0, 32)
	return int32(i), err
}
func (s stringVariant) Int64E() (int64, error) {
	return strconv.ParseInt(s.val, 0, 64)
}
func (s stringVariant) IntE() (int, error) {
	i, err := strconv.ParseInt(s.val, 0, 0)
	return int(i), err
}
func (s stringVariant) ByteE() (byte, error) {
	i, err := strconv.ParseUint(s.val, 0, 8)
	return byte(i), err
}
func (s stringVariant) Float32E() (float32, error) {
	f, err := strconv.ParseFloat(s.val, 32)
	if err != nil && errors.Is(err, strconv.ErrSyntax) {
		i, err := strconv.ParseInt(s.val, 0, 64)
		if err == nil {
			return float32(i), nil
		} else if errors.Is(err, strconv.ErrRange) {
			u, err := strconv.ParseUint(s.val, 0, 64)
			if err == nil {
				return float32(u), nil
			}
		}
		return float32(f), err
	}
	return float32(f), err
}
func (s stringVariant) Float64E() (float64, error) {
	f, err := strconv.ParseFloat(s.val, 64)
	if err != nil && errors.Is(err, strconv.ErrSyntax) {
		i, err := strconv.ParseInt(s.val, 0, 64)
		if err == nil {
			return float64(i), nil
		} else if errors.Is(err, strconv.ErrRange) {
			u, err := strconv.ParseUint(s.val, 0, 64)
			if err == nil {
				return float64(u), nil
			}
		}
		return f, err
	}
	return f, err
}
func (s stringVariant) BoolE() (bool, error) {
	return strconv.ParseBool(s.val)
}
func (s stringVariant) StringE() (string, error) {
	return s.val, nil
}
func (s stringVariant) DurationE() (time.Duration, error) {
	return time.ParseDuration(s.val)
}
func (s stringVariant) TimeE() (time.Time, error) {
	// TODO:
	return time.Time{}, nil
}
func (s stringVariant) Value() (driver.Value, error) {
	return driver.Value(s.val), nil
}
func (s stringVariant) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.val)
}
