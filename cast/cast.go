package cast

import "github.com/ipsusila/types"

// Number convert argument v to type T
func Number[V types.Scalar, T types.Scalar](v V) (T, error) {
	res := T(v)
	if V(res) != v {
		return res, ErrDataLost
	}
	return res, nil
}
