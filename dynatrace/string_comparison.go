// Code generated by go-swagger; DO NOT EDIT.

package dynatrace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StringComparison string comparison
// swagger:model StringComparison
type StringComparison struct {
	negateField *bool

	operatorField Enum

	valueField interface{}

	// The comparison is case-sensitive (`true`) or insensitive (`false`).
	CaseSensitive bool `json:"caseSensitive,omitempty"`
}

// Negate gets the negate of this subtype
func (m *StringComparison) Negate() *bool {
	return m.negateField
}

// SetNegate sets the negate of this subtype
func (m *StringComparison) SetNegate(val *bool) {
	m.negateField = val
}

// Operator gets the operator of this subtype
func (m *StringComparison) Operator() Enum {
	return m.operatorField
}

// SetOperator sets the operator of this subtype
func (m *StringComparison) SetOperator(val Enum) {
	m.operatorField = val
}

// Type gets the type of this subtype
func (m *StringComparison) Type() string {
	return "StringComparison"
}

// SetType sets the type of this subtype
func (m *StringComparison) SetType(val string) {

}

// Value gets the value of this subtype
func (m *StringComparison) Value() interface{} {
	return m.valueField
}

// SetValue sets the value of this subtype
func (m *StringComparison) SetValue(val interface{}) {
	m.valueField = val
}

// CaseSensitive gets the case sensitive of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *StringComparison) UnmarshalJSON(raw []byte) error {
	var data struct {

		// The comparison is case-sensitive (`true`) or insensitive (`false`).
		CaseSensitive bool `json:"caseSensitive,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Negate *bool `json:"negate"`

		Operator Enum `json:"operator"`

		Type string `json:"type"`

		Value interface{} `json:"value,omitempty"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result StringComparison

	result.negateField = base.Negate

	result.operatorField = base.Operator

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.valueField = base.Value

	result.CaseSensitive = data.CaseSensitive

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m StringComparison) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// The comparison is case-sensitive (`true`) or insensitive (`false`).
		CaseSensitive bool `json:"caseSensitive,omitempty"`
	}{

		CaseSensitive: m.CaseSensitive,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Negate *bool `json:"negate"`

		Operator Enum `json:"operator"`

		Type string `json:"type"`

		Value interface{} `json:"value,omitempty"`
	}{

		Negate: m.Negate(),

		Operator: m.Operator(),

		Type: m.Type(),

		Value: m.Value(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this string comparison
func (m *StringComparison) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNegate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperator(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StringComparison) validateNegate(formats strfmt.Registry) error {

	if err := validate.Required("negate", "body", m.Negate()); err != nil {
		return err
	}

	return nil
}

func (m *StringComparison) validateOperator(formats strfmt.Registry) error {

	if err := validate.Required("operator", "body", m.Operator()); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StringComparison) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StringComparison) UnmarshalBinary(b []byte) error {
	var res StringComparison
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
