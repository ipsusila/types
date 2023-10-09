package internal

import "errors"

var (
	errNegativeNotAllowed = errors.New("unable to cast negative value")
	errCastDataLost       = errors.New("data lost between type conversion")
)
