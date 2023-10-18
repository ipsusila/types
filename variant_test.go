package types_test

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"strconv"
	"testing"
	"time"

	"github.com/ipsusila/types"
	"github.com/stretchr/testify/assert"
)

type stringer struct {
	v int
}

func (s stringer) String() string {
	return strconv.Itoa(s.v)
}

func TestMarshalValue(t *testing.T) {
	type valExp struct {
		val any
		exp string
	}

	ts := "2023-11-12T01:02:03.123+07:00"
	tm, err := time.Parse(time.RFC3339Nano, ts)
	assert.NoError(t, err)
	vals := []valExp{
		{nil, "null"},
		{true, "true"},
		{false, "false"},
		{uint8(1), "1"},
		{int8(1), "1"},
		{byte(1), "1"},
		{uint16(1), "1"},
		{int16(1), "1"},
		{uint32(1), "1"},
		{int32(1), "1"},
		{uint(1), "1"},
		{int(1), "1"},
		{uint64(1), "1"},
		{int64(1), "1"},
		{float32(123.100), "123.1"},
		{123.100, "123.1"},
		{"String", `"String"`},
		{"", `""`},
		{tm, strconv.Quote(ts)},
	}

	var v types.Variant
	for _, val := range vals {
		v = types.NewVariant(val.val)
		data, err := v.MarshalJSON()
		assert.NoError(t, err)
		assert.Equal(t, val.exp, string(data))
	}
}

func TestVariant(t *testing.T) {
	v := types.NewVariant(nil)
	fmt.Printf("%#v\n", v)

	v = types.NewVariant(1000)
	fmt.Printf("%#v\n", v)
	u8, err := v.Uint8E()
	fmt.Println(u8, err)

	v = types.NewVariant(110.123)
	fmt.Printf("%#v\n", v)
	u8, err = v.Uint8E()
	fmt.Println(u8, err)

	buf := bytes.Buffer{}
	binary.Write(&buf, binary.BigEndian, float32(1))
	fmt.Println(buf.Len())

	vi := 123
	v = types.NewVariant(&vi)
	fmt.Printf("%#v\n", v)
	u8, err = v.Uint8E()
	fmt.Println(u8, err)

	v = types.NewVariant(float32(math.MaxFloat32))
	fmt.Printf("%#v\n", v)
	u8, err = v.Uint8E()
	fmt.Println(u8, err)

	v = types.NewVariant(1 + float64(math.MaxFloat32))
	fmt.Printf("%#v\n", v)
	f32, err := v.Float32E()
	fmt.Println(f32, err)

	v = types.NewVariant(true)
	fmt.Println(v.String())

	v = types.NewVariant("0xFF")
	fmt.Printf("%#v\n", v)
	f32, err = v.Float32E()
	fmt.Println(f32, err)

	v = types.NewVariant("0xFFFF_FFFF_FFFF_FFFF")
	fmt.Printf("%#v\n", v)
	f32, err = v.Float32E()
	fmt.Println(f32, err)

	v = types.NewVariant(".123")
	fmt.Printf("%#v\n", v)
	f32, err = v.Float32E()
	fmt.Println(f32, err)

	v = types.NewVariant(stringer{v: 123})
	fmt.Printf("%#v\n", v)
	f32, err = v.Float32E()
	fmt.Println(f32, err)

	v = types.NewVariant(time.Second)
	fmt.Printf("%#v\n", v)
	f32, err = v.Float32E()
	fmt.Println(f32, err)
	fmt.Println(v.Duration())
}
