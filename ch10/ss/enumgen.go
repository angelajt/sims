// Code generated by "core generate -add-types"; DO NOT EDIT.

package main

import (
	"cogentcore.org/core/enums"
)

var _EnvTypeValues = []EnvType{0, 1, 2, 3}

// EnvTypeN is the highest valid value for type EnvType, plus one.
const EnvTypeN EnvType = 4

var _EnvTypeValueMap = map[string]EnvType{`Probe`: 0, `Besner`: 1, `Glushko`: 2, `Taraban`: 3}

var _EnvTypeDescMap = map[EnvType]string{0: ``, 1: ``, 2: ``, 3: ``}

var _EnvTypeMap = map[EnvType]string{0: `Probe`, 1: `Besner`, 2: `Glushko`, 3: `Taraban`}

// String returns the string representation of this EnvType value.
func (i EnvType) String() string { return enums.String(i, _EnvTypeMap) }

// SetString sets the EnvType value from its string representation,
// and returns an error if the string is invalid.
func (i *EnvType) SetString(s string) error {
	return enums.SetString(i, s, _EnvTypeValueMap, "EnvType")
}

// Int64 returns the EnvType value as an int64.
func (i EnvType) Int64() int64 { return int64(i) }

// SetInt64 sets the EnvType value from an int64.
func (i *EnvType) SetInt64(in int64) { *i = EnvType(in) }

// Desc returns the description of the EnvType value.
func (i EnvType) Desc() string { return enums.Desc(i, _EnvTypeDescMap) }

// EnvTypeValues returns all possible values for the type EnvType.
func EnvTypeValues() []EnvType { return _EnvTypeValues }

// Values returns all possible values for the type EnvType.
func (i EnvType) Values() []enums.Enum { return enums.Values(_EnvTypeValues) }

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i EnvType) MarshalText() ([]byte, error) { return []byte(i.String()), nil }

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *EnvType) UnmarshalText(text []byte) error { return enums.UnmarshalText(i, text, "EnvType") }