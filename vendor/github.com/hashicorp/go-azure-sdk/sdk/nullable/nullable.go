// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package nullable

import (
	"bytes"
	"encoding/json"
)

var _ json.Marshaler = Type[string]{}
var _ json.Unmarshaler = &Type[string]{}

type Type[T comparable] map[bool]T

// Value returns a new Type[T], setting its type and value to the provided value
func Value[T comparable](t T) Type[T] {
	var n Type[T]
	n.Set(t)
	return n
}

// NoZero returns a new Type[T], setting its type and value, whilst also nulling the value if it was set to
// its zero value. This ensures that zero values are sent as null.
func NoZero[T comparable](t T) Type[T] {
	var n Type[T]
	n.SetNoZero(t)
	return n
}

// Get retrieves the underlying value, if present, and returns nil if the value is null
func (t Type[T]) Get() *T {
	var empty T
	if t.IsNull() {
		return nil
	}
	if t.IsSet() {
		ret := t[true]
		return &ret
	}
	return &empty
}

// GetOrZero retrieves the underlying value, if present, and returns the zero value if null
func (t Type[T]) GetOrZero() T {
	var empty T
	val := t.Get()
	if val == nil {
		return empty
	}
	return *val
}

// Set sets the underlying value to a given value
func (t *Type[T]) Set(value T) {
	*t = map[bool]T{true: value}
}

// SetNoZero sets the underlying value to a given value, whilst also nulling the value if it was set to
// its zero value. This ensures that zero values are sent as null.
func (t *Type[T]) SetNoZero(value T) {
	var empty T
	*t = map[bool]T{value != empty: value}
}

// SetNull clears the value and ensures a value of `null`
func (t *Type[T]) SetNull() {
	var empty T
	*t = map[bool]T{false: empty}
}

// SetUnspecified clears the value
func (t *Type[T]) SetUnspecified() {
	*t = map[bool]T{}
}

// IsNull indicates whether the value was set to `null`
func (t Type[T]) IsNull() bool {
	_, foundNull := t[false]
	return foundNull
}

// IsSet indicates whether a value is set
func (t Type[T]) IsSet() bool {
	return len(t) != 0
}

func (t Type[T]) MarshalJSON() ([]byte, error) {
	// note: if value was unspecified, and `omitempty` is set on the field tags, `json.Marshal` will omit this field
	// if value was specified, and `null`, marshal it
	if t.IsNull() {
		return []byte("null"), nil
	}
	// otherwise, we have a value, so marshal it
	return json.Marshal(t[true])
}

func (t *Type[T]) UnmarshalJSON(data []byte) error {
	// note: if value is unspecified, UnmarshalJSON won't be called
	// if value is specified and `null`
	if bytes.Equal(data, []byte("null")) {
		t.SetNull()
		return nil
	}
	// otherwise, we have an actual value, so parse it
	var val T
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}
	t.Set(val)
	return nil
}
