package types

type Dictionary map[string]Variant

func (d Dictionary) Get(key string) Variant {
	v, ok := d[key]
	if !ok {
		return NewVariant()
	}
	return v
}
func (d Dictionary) Set(key string, v any) {
	d[key] = NewVariant(v)
}
