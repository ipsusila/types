package types

import "golang.org/x/exp/constraints"

// Scalar number constraint
type Scalar interface {
	constraints.Float | constraints.Integer
}
