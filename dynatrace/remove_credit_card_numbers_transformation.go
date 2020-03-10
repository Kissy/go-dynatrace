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

// RemoveCreditCardNumbersTransformation remove credit card numbers transformation
// swagger:model RemoveCreditCardNumbersTransformation
type RemoveCreditCardNumbersTransformation struct {
	RemoveCreditCardNumbersTransformationAllOf1
}

// Type gets the type of this subtype
func (m *RemoveCreditCardNumbersTransformation) Type() string {
	return "RemoveCreditCardNumbersTransformation"
}

// SetType sets the type of this subtype
func (m *RemoveCreditCardNumbersTransformation) SetType(val string) {

}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *RemoveCreditCardNumbersTransformation) UnmarshalJSON(raw []byte) error {
	var data struct {
		RemoveCreditCardNumbersTransformationAllOf1
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

	var result RemoveCreditCardNumbersTransformation

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.RemoveCreditCardNumbersTransformationAllOf1 = data.RemoveCreditCardNumbersTransformationAllOf1

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m RemoveCreditCardNumbersTransformation) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		RemoveCreditCardNumbersTransformationAllOf1
	}{

		RemoveCreditCardNumbersTransformationAllOf1: m.RemoveCreditCardNumbersTransformationAllOf1,
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

// Validate validates this remove credit card numbers transformation
func (m *RemoveCreditCardNumbersTransformation) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with RemoveCreditCardNumbersTransformationAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *RemoveCreditCardNumbersTransformation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RemoveCreditCardNumbersTransformation) UnmarshalBinary(b []byte) error {
	var res RemoveCreditCardNumbersTransformation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RemoveCreditCardNumbersTransformationAllOf1 The transformation of the `REMOVE_CREDIT_CARDS` type.
//
// The transformation automatically detects and removes credit card numbers. No additional parameters needed.
// swagger:model RemoveCreditCardNumbersTransformationAllOf1
type RemoveCreditCardNumbersTransformationAllOf1 interface{}