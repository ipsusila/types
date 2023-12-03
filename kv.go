package types

import (
	"fmt"
	"strings"
	"time"
)

// KV key value data
type KV map[string]any

func (k KV) splitKey(key string) ([]string, error) {
	keys := []string{}
	sb := strings.Builder{}
	escape := false
	for _, c := range key {
		switch c {
		case '\\':
			if escape {
				escape = false
				sb.WriteByte('\\')
			} else {
				escape = true
			}
		case '.':
			if escape {
				escape = false
				sb.WriteByte('.')
			} else if sb.Len() > 0 {
				keys = append(keys, sb.String())
				sb.Reset()
			}
		default:
			if escape {
				return keys, fmt.Errorf("invalid KV key: %s", key)
			}
			sb.WriteRune(c)
		}
	}
	if sb.Len() > 0 {
		keys = append(keys, sb.String())
	}
	return keys, nil
}

func (k KV) get(keys []string) (any, bool) {
	nk := len(keys)
	if nk == 0 {
		return nil, false
	}
	lk := keys[nk-1]
	rkeys := keys[0 : nk-1]
	kv := map[string]any(k)
	for _, key := range rkeys {
		v, ok := kv[key]
		if !ok {
			return nil, false
		}
		kv, ok = v.(map[string]any)
		if !ok {
			return nil, false
		}
	}

	if kv == nil {
		return nil, false
	}
	v, ok := kv[lk]
	return v, ok
}

func (k KV) V(key string) RVariant {
	keys, _ := k.splitKey(key)
	val, _ := k.get(keys)
	return NewRVariant(val)
}

func (k KV) Bool(key string, def ...bool) bool {
	va := k.V(key)
	if va.IsNil() {
		dv := false
		if len(def) != 0 {
			dv = def[0]
		}
		return dv
	}

	return va.Bool()
}

func (k KV) Int(key string, def ...int) int {
	va := k.V(key)
	if va.IsNil() {
		dv := 0
		if len(def) != 0 {
			dv = def[0]
		}
		return dv
	}

	return va.Int()
}

func (k KV) Uint(key string, def ...uint) uint {
	va := k.V(key)
	if va.IsNil() {
		dv := uint(0)
		if len(def) != 0 {
			dv = def[0]
		}
		return dv
	}

	return va.Uint()
}

func (k KV) Float(key string, def ...float64) float64 {
	va := k.V(key)
	if va.IsNil() {
		dv := 0.0
		if len(def) != 0 {
			dv = def[0]
		}
		return dv
	}

	return va.Float64()
}

func (k KV) String(key string, def ...string) string {
	va := k.V(key)
	if va.IsNil() {
		dv := ""
		if len(def) != 0 {
			dv = def[0]
		}
		return dv
	}

	return va.String()
}

func (k KV) Time(key string, def ...time.Time) time.Time {
	va := k.V(key)
	if va.IsNil() {
		dv := time.Time{}
		if len(def) != 0 {
			dv = def[0]
		}
		return dv
	}

	return va.Time()
}
func (k KV) Duration(key string, def ...time.Duration) time.Duration {
	va := k.V(key)
	if va.IsNil() {
		dv := time.Duration(0)
		if len(def) != 0 {
			dv = def[0]
		}
		return dv
	}

	return va.Duration()
}
