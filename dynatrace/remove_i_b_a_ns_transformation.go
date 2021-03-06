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
)

// RemoveIBANsTransformation remove i b a ns transformation
// swagger:model RemoveIBANsTransformation
type RemoveIBANsTransformation struct {
	RemoveIBANsTransformationAllOf1
}

// Type gets the type of this subtype
func (m *RemoveIBANsTransformation) Type() string {
	return "RemoveIBANsTransformation"
}

// SetType sets the type of this subtype
func (m *RemoveIBANsTransformation) SetType(val string) {

}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *RemoveIBANsTransformation) UnmarshalJSON(raw []byte) error {
	var data struct {
		RemoveIBANsTransformationAllOf1
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Type string `json:"type"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result RemoveIBANsTransformation

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.RemoveIBANsTransformationAllOf1 = data.RemoveIBANsTransformationAllOf1

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m RemoveIBANsTransformation) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		RemoveIBANsTransformationAllOf1
	}{

		RemoveIBANsTransformationAllOf1: m.RemoveIBANsTransformationAllOf1,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Type string `json:"type"`
	}{

		Type: m.Type(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this remove i b a ns transformation
func (m *RemoveIBANsTransformation) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with RemoveIBANsTransformationAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *RemoveIBANsTransformation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RemoveIBANsTransformation) UnmarshalBinary(b []byte) error {
	var res RemoveIBANsTransformation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RemoveIBANsTransformationAllOf1 The transformation of the `REMOVE_IBANS` type.
//
// The transformation automatically detects and removes International Bank Account Numbers (IBAN). No additional parameters needed.
// swagger:model RemoveIBANsTransformationAllOf1
type RemoveIBANsTransformationAllOf1 interface{}
