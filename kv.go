package types

import (
	"time"
)

// KV key value data
type KV map[string]any

func (k KV) V(key string) RVariant {
	return NewRVariant(k[key])
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
