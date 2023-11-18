package internal

import (
	"database/sql/driver"
	"time"
)

type nilVariant struct {
}

func (v nilVariant) IsNil() bool {
	return true
}
func (v nilVariant) IsZero() bool {
	return false
}
func (v nilVariant) Interface() any {
	return nil
}
func (v nilVariant) Uint8E() (uint8, error) {
	return 0, errNilValue
}
func (v nilVariant) Uint16E() (uint16, error) {
	return 0, errNilValue
}
func (v nilVariant) Uint32E() (uint32, error) {
	return 0, errNilValue
}
func (v nilVariant) Uint64E() (uint64, error) {
	return 0, errNilValue
}
func (v nilVariant) UintE() (uint, error) {
	return 0, errNilValue
}
func (v nilVariant) Int8E() (int8, error) {
	return 0, errNilValue
}
func (v nilVariant) Int16E() (int16, error) {
	return 0, errNilValue
}
func (v nilVariant) Int32E() (int32, error) {
	return 0, errNilValue
}
func (v nilVariant) Int64E() (int64, error) {
	return 0, errNilValue
}
func (v nilVariant) IntE() (int, error) {
	return 0, errNilValue
}
func (v nilVariant) ByteE() (byte, error) {
	return 0, errNilValue
}
func (v nilVariant) Float32E() (float32, error) {
	return 0, errNilValue
}
func (v nilVariant) Float64E() (float64, error) {
	return 0, errNilValue
}
func (v nilVariant) BoolE() (bool, error) {
	return false, errNilValue
}
func (v nilVariant) StringE() (string, error) {
	return "", errNilValue
}
func (v nilVariant) DurationE() (time.Duration, error) {
	return 0, errNilValue
}
func (v nilVariant) TimeE() (time.Time, error) {
	return time.Time{}, errNilValue
}
func (v nilVariant) Value() (driver.Value, error) {
	return nil, nil
}
func (v nilVariant) MarshalJSON() ([]byte, error) {
	return []byte{'n', 'u', 'l', 'l'}, nil
}
