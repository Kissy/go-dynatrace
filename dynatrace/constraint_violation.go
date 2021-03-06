// Code generated by go-swagger; DO NOT EDIT.

package dynatrace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConstraintViolation constraint violation
// swagger:model ConstraintViolation
type ConstraintViolation struct {

	// location
	Location string `json:"location,omitempty"`

	// message
	// Read Only: true
	Message string `json:"message,omitempty"`

	// parameter location
	// Read Only: true
	// Enum: [PATH PAYLOAD_BODY QUERY]
	ParameterLocation string `json:"parameterLocation,omitempty"`

	// path
	// Read Only: true
	Path string `json:"path,omitempty"`
}

// Validate validates this constraint violation
func (m *ConstraintViolation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateParameterLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var constraintViolationTypeParameterLocationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PATH","PAYLOAD_BODY","QUERY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		constraintViolationTypeParameterLocationPropEnum = append(constraintViolationTypeParameterLocationPropEnum, v)
	}
}

const (

	// ConstraintViolationParameterLocationPATH captures enum value "PATH"
	ConstraintViolationParameterLocationPATH string = "PATH"

	// ConstraintViolationParameterLocationPAYLOADBODY captures enum value "PAYLOAD_BODY"
	ConstraintViolationParameterLocationPAYLOADBODY string = "PAYLOAD_BODY"

	// ConstraintViolationParameterLocationQUERY captures enum value "QUERY"
	ConstraintViolationParameterLocationQUERY string = "QUERY"
)

// prop value enum
func (m *ConstraintViolation) validateParameterLocationEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, constraintViolationTypeParameterLocationPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ConstraintViolation) validateParameterLocation(formats strfmt.Registry) error {

	if swag.IsZero(m.ParameterLocation) { // not required
		return nil
	}

	// value enum
	if err := m.validateParameterLocationEnum("parameterLocation", "body", m.ParameterLocation); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConstraintViolation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConstraintViolation) UnmarshalBinary(b []byte) error {
	var res ConstraintViolation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
