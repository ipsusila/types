package internal

import (
	"database/sql/driver"
	"encoding/json"
	"math"
	"time"
)

type timeVariant time.Time

func (t timeVariant) IsNil() bool {
	return false
}
func (t timeVariant) IsZero() bool {
	return time.Time(t).IsZero()
}
func (t timeVariant) Interface() any {
	return time.Time(t)
}

func (t timeVariant) Uint8E() (uint8, error) {
	en := t.epochNano()
	if en < 0 {
		return 0, errNegativeNotAllowed
	} else if en > math.MaxUint8 {
		return uint8(en), errOutOfRange
	}
	return uint8(en), nil
}
func (t timeVariant) Uint16E() (uint16, error) {
	en := t.epochNano()
	if en < 0 {
		return 0, errNegativeNotAllowed
	} else if en > math.MaxUint16 {
		return uint16(en), errOutOfRange
	}
	return uint16(en), nil
}
func (t timeVariant) Uint32E() (uint32, error) {
	en := t.epochNano()
	if en < 0 {
		return 0, errNegativeNotAllowed
	} else if en > math.MaxUint32 {
		return uint32(en), errOutOfRange
	}
	return uint32(en), nil
}
func (t timeVariant) Uint64E() (uint64, error) {
	en := t.epochNano()
	if en < 0 {
		return 0, errNegativeNotAllowed
	}
	return uint64(en), nil
}
func (t timeVariant) UintE() (uint, error) {
	en := t.epochNano()
	if en < 0 {
		return 0, errNegativeNotAllowed
	}
	return uint(en), nil
}
func (t timeVariant) Int8E() (int8, error) {
	en := t.epochNano()
	if en < math.MinInt8 || en > math.MaxInt8 {
		return int8(en), errOutOfRange
	}
	return int8(en), nil
}
func (t timeVariant) Int16E() (int16, error) {
	en := t.epochNano()
	if en < math.MinInt16 || en > math.MaxInt16 {
		return int16(en), errOutOfRange
	}
	return int16(en), nil
}
func (t timeVariant) Int32E() (int32, error) {
	en := t.epochNano()
	if en < math.MinInt32 || en > math.MaxInt32 {
		return int32(en), errOutOfRange
	}
	return int32(en), nil
}
func (t timeVariant) Int64E() (int64, error) {
	en := t.epochNano()
	return int64(en), nil
}
func (t timeVariant) IntE() (int, error) {
	en := t.epochNano()
	return int(en), nil
}
func (t timeVariant) ByteE() (byte, error) {
	en := t.epochNano()
	if en < 0 {
		return 0, errNegativeNotAllowed
	} else if en > math.MaxUint8 {
		return byte(en), errOutOfRange
	}
	return byte(en), nil
}
func (t timeVariant) Float32E() (float32, error) {
	nano := float32(time.Time(t).Nanosecond()) / 1000_000_000
	epoch := float32(time.Time(t).Unix())
	return epoch + nano, nil

}
func (t timeVariant) Float64E() (float64, error) {
	nano := float64(time.Time(t).Nanosecond()) / 1000_000_000
	epoch := float64(time.Time(t).Unix())
	return epoch + nano, nil
}
func (t timeVariant) BoolE() (bool, error) {
	if t.IsZero() {
		return false, nil
	}
	return true, nil
}
func (t timeVariant) StringE() (string, error) {
	return time.Time(t).Format(time.RFC3339Nano), nil
}
func (t timeVariant) DurationE() (time.Duration, error) {
	return time.Duration(t.epochNano()), nil
}
func (t timeVariant) TimeE() (time.Time, error) {
	return time.Time(t), nil
}
func (t timeVariant) epochNano() int64 {
	return time.Time(t).UnixNano()
}
func (t timeVariant) Value() (driver.Value, error) {
	return driver.Value(time.Time(t)), nil
}
func (t timeVariant) MarshalJSON() ([]byte, error) {
	str := time.Time(t).Format(time.RFC3339Nano)
	return json.Marshal(str)
}
