package cast_test

import (
	"testing"

	"github.com/ipsusila/types/cast"
	"github.com/stretchr/testify/assert"
)

func TestSigned(t *testing.T) {
	v, err := cast.ToInt8(123)
	assert.NoError(t, err)
	assert.Equal(t, int8(123), v)

	v, err = cast.ToInt8(255)
	assert.Error(t, err)
	assert.NotEqual(t, 255, v)
}
