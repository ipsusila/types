package cast

import (
	"errors"
	"strconv"
)

var (
	ErrNegativeNotAllowed = errors.New("unable to cast negative value")
	ErrDataLost           = errors.New("data lost between type conversion")
	ErrNilValue           = errors.New("can not convert nil to types")
	ErrCannotConvert      = errors.New("can not convert value to given type")
	ErrUnsupportedFormat  = errors.New("unsupporte format")
	ErrOutOfRange         = strconv.ErrRange
)
