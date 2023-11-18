package cast_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/ipsusila/types/cast"
)

func testNumberU(t *testing.T, v uint64) {
	i8, err := cast.Number[uint64, int8](v)
	fmt.Printf("Res: %v, error: %v\n", i8, err)
	i16, err := cast.Number[uint64, int16](v)
	fmt.Printf("Res: %v, error: %v\n", i16, err)
	i32, err := cast.Number[uint64, int32](v)
	fmt.Printf("Res: %v, error: %v\n", i32, err)
	i64, err := cast.Number[uint64, int64](v)
	fmt.Printf("Res: %v, error: %v\n", i64, err)
	i, err := cast.Number[uint64, int](v)
	fmt.Printf("Res: %v, error: %v\n", i, err)

	i8, err = cast.ToInt8(v)
	fmt.Printf("Res: %v, error: %v\n", i8, err)
}

func TestNumber(t *testing.T) {
	testNumberU(t, math.MaxUint64)
}
