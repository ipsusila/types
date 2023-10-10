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
)

type stringer struct {
	v int
}

func (s stringer) String() string {
	return strconv.Itoa(s.v)
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
