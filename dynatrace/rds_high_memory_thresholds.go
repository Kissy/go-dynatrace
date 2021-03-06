// Code generated by go-swagger; DO NOT EDIT.

package dynatrace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RdsHighMemoryThresholds Custom thresholds for RDS running out of memory. If not set, automatic mode is used.
//
//  **All** conditions must be fulfilled to trigger an alert.
// swagger:model RdsHighMemoryThresholds
type RdsHighMemoryThresholds struct {

	// Freeable memory is lower than *X* Megabytes in 3 out of 5 samples.
	// Required: true
	// Maximum: 1000
	// Minimum: 0.01
	FreeMemory *float64 `json:"freeMemory"`

	// Swap usage is higher than *X* Gigabytes in 3 out of 5 samples.
	// Required: true
	// Maximum: 1000
	// Minimum: 0.01
	SwapUsage *float64 `json:"swapUsage"`
}

// Validate validates this rds high memory thresholds
func (m *RdsHighMemoryThresholds) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFreeMemory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSwapUsage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RdsHighMemoryThresholds) validateFreeMemory(formats strfmt.Registry) error {

	if err := validate.Required("freeMemory", "body", m.FreeMemory); err != nil {
		return err
	}

	if err := validate.Minimum("freeMemory", "body", float64(*m.FreeMemory), 0.01, false); err != nil {
		return err
	}

	if err := validate.Maximum("freeMemory", "body", float64(*m.FreeMemory), 1000, false); err != nil {
		return err
	}

	return nil
}

func (m *RdsHighMemoryThresholds) validateSwapUsage(formats strfmt.Registry) error {

	if err := validate.Required("swapUsage", "body", m.SwapUsage); err != nil {
		return err
	}

	if err := validate.Minimum("swapUsage", "body", float64(*m.SwapUsage), 0.01, false); err != nil {
		return err
	}

	if err := validate.Maximum("swapUsage", "body", float64(*m.SwapUsage), 1000, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RdsHighMemoryThresholds) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RdsHighMemoryThresholds) UnmarshalBinary(b []byte) error {
	var res RdsHighMemoryThresholds
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
