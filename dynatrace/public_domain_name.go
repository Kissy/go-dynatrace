// Code generated by go-swagger; DO NOT EDIT.

package dynatrace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PublicDomainName The contribution to the service ID calculation from the domain name where the web request has been detected.
//
//  You have two mutually exclusive options:
// * Override the detected value with a specified static value. Specify the new value in the **valueOverride** field.
// * Dynamically transform the detected value. Specify the transformation parameters in the **transformations** field.
// swagger:model PublicDomainName
type PublicDomainName struct {

	// Use (`true`) or don't use (`false`) the detected host name as base for transformation.
	//
	//  Not applicable if the override is specified.
	CopyFromHostName bool `json:"copyFromHostName,omitempty"`

	transformationsField []TransformationBase

	// The value to be used instead of the detected value.
	ValueOverride string `json:"valueOverride,omitempty"`
}

// Transformations gets the transformations of this base type
func (m *PublicDomainName) Transformations() []TransformationBase {
	return m.transformationsField
}

// SetTransformations sets the transformations of this base type
func (m *PublicDomainName) SetTransformations(val []TransformationBase) {
	m.transformationsField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *PublicDomainName) UnmarshalJSON(raw []byte) error {
	var data struct {
		CopyFromHostName bool `json:"copyFromHostName,omitempty"`

		Transformations json.RawMessage `json:"transformations"`

		ValueOverride string `json:"valueOverride,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var propTransformations []TransformationBase
	if string(data.Transformations) != "null" {
		transformations, err := UnmarshalTransformationBaseSlice(bytes.NewBuffer(data.Transformations), runtime.JSONConsumer())
		if err != nil && err != io.EOF {
			return err
		}
		propTransformations = transformations
	}

	var result PublicDomainName

	// copyFromHostName
	result.CopyFromHostName = data.CopyFromHostName

	// transformations
	result.transformationsField = propTransformations

	// valueOverride
	result.ValueOverride = data.ValueOverride

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m PublicDomainName) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		CopyFromHostName bool `json:"copyFromHostName,omitempty"`

		ValueOverride string `json:"valueOverride,omitempty"`
	}{

		CopyFromHostName: m.CopyFromHostName,

		ValueOverride: m.ValueOverride,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Transformations []TransformationBase `json:"transformations"`
	}{

		Transformations: m.transformationsField,
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this public domain name
func (m *PublicDomainName) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTransformations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublicDomainName) validateTransformations(formats strfmt.Registry) error {

	if swag.IsZero(m.Transformations()) { // not required
		return nil
	}

	iTransformationsSize := int64(len(m.Transformations()))

	if err := validate.MinItems("transformations", "body", iTransformationsSize, 0); err != nil {
		return err
	}

	if err := validate.MaxItems("transformations", "body", iTransformationsSize, 10); err != nil {
		return err
	}

	for i := 0; i < len(m.Transformations()); i++ {

		if err := m.transformationsField[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("transformations" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PublicDomainName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicDomainName) UnmarshalBinary(b []byte) error {
	var res PublicDomainName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
