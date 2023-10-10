package internal

import (
	"errors"
	"strconv"
)

var (
	errNegativeNotAllowed = errors.New("unable to cast negative value")
	errCastDataLost       = errors.New("data lost between type conversion")
	errNilValue           = errors.New("can not convert nil to types")
	errCannotConvert      = errors.New("can not convert value to given type")
	errOutOfRange         = strconv.ErrRange
)
