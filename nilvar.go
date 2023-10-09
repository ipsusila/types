package types

import "errors"

var errNilValue = errors.New("can not convert nil to types")

type nilVariant struct {
}

func (v nilVariant) IsNil() bool {
	return true
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

func (v nilVariant) Uint8() uint8 {
	return 0
}
func (v nilVariant) Uint16() uint16 {
	return 0
}
func (v nilVariant) Uint32() uint32 {
	return 0
}
func (v nilVariant) Uint64() uint64 {
	return 0
}
func (v nilVariant) Uint() uint {
	return 0
}
func (v nilVariant) Int8() int8 {
	return 0
}
func (v nilVariant) Int16() int16 {
	return 0
}
func (v nilVariant) Int32() int32 {
	return 0
}
func (v nilVariant) Int64() int64 {
	return 0
}
func (v nilVariant) Int() int {
	return 0
}
func (v nilVariant) Byte() byte {
	return 0
}
func (v nilVariant) Float32() float32 {
	return 0
}
func (v nilVariant) Float64() float64 {
	return 0
}
func (v nilVariant) Bool() bool {
	return false
}
