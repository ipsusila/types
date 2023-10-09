package types_test

import (
	"fmt"
	"testing"

	"github.com/ipsusila/types"
)

func TestVariant(t *testing.T) {
	v := types.NewVariant(nil)
	fmt.Println(v)

	v = types.NewVariant(1000)
	u8, err := v.Uint8E()
	fmt.Println(u8, err)
}
