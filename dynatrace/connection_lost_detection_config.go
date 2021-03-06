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

// ConnectionLostDetectionConfig Configuration of lost connection detection.
// swagger:model ConnectionLostDetectionConfig
type ConnectionLostDetectionConfig struct {

	// The detection is enabled (`true`) or disabled (`false`).
	// Required: true
	Enabled *bool `json:"enabled"`

	// Alert (`true`) on graceful host shutdowns.
	// Required: true
	EnabledOnGracefulShutdowns *bool `json:"enabledOnGracefulShutdowns"`
}

// Validate validates this connection lost detection config
func (m *ConnectionLostDetectionConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabledOnGracefulShutdowns(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConnectionLostDetectionConfig) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *ConnectionLostDetectionConfig) validateEnabledOnGracefulShutdowns(formats strfmt.Registry) error {

	if err := validate.Required("enabledOnGracefulShutdowns", "body", m.EnabledOnGracefulShutdowns); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConnectionLostDetectionConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConnectionLostDetectionConfig) UnmarshalBinary(b []byte) error {
	var res ConnectionLostDetectionConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
