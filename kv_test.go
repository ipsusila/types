package types_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ipsusila/types"
	"github.com/stretchr/testify/assert"
)

func TestKV(t *testing.T) {
	strList := []string{
		`{}`,
		`{"key": 123}`,
		`{"hello":{"world": 456}}`,
		`{"say.hello": "Hello World!"}`,
	}
	keyList := []string{
		"",
		"none",
		"key",
		"hello.world",
		"say",
		"say.hello",
		"say\\.hello",
	}

	for _, str := range strList {
		var obj map[string]any
		err := json.Unmarshal([]byte(str), &obj)
		assert.NoError(t, err)

		kv := types.KV(obj)
		for _, key := range keyList {
			v := kv.V(key)
			if !v.IsNil() {
				fmt.Printf("Variant(%s > %s) %#v\n", str, key, v)
			}
		}
	}
}
