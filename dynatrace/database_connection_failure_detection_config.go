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

// DatabaseConnectionFailureDetectionConfig Parameters of the failed database connections detection.
//
// The alert is triggered when failed connections number exceeds **connectionFailsCount** during any **timePeriodMinutes** minutes period.
// swagger:model DatabaseConnectionFailureDetectionConfig
type DatabaseConnectionFailureDetectionConfig struct {

	// Number of failed database connections during any **timePeriodMinutes** minutes period to trigger an alert.
	// Maximum: 9.999999e+06
	// Minimum: 1
	ConnectionFailsCount int32 `json:"connectionFailsCount,omitempty"`

	// The detection is enabled (`true`) or disabled (`false`).
	// Required: true
	Enabled *bool `json:"enabled"`

	// The *X* minutes time period during which the **connectionFailsCount** is evaluated.
	// Maximum: 15
	// Minimum: 1
	TimePeriodMinutes int32 `json:"timePeriodMinutes,omitempty"`
}

// Validate validates this database connection failure detection config
func (m *DatabaseConnectionFailureDetectionConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectionFailsCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimePeriodMinutes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatabaseConnectionFailureDetectionConfig) validateConnectionFailsCount(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectionFailsCount) { // not required
		return nil
	}

	if err := validate.MinimumInt("connectionFailsCount", "body", int64(m.ConnectionFailsCount), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("connectionFailsCount", "body", int64(m.ConnectionFailsCount), 9.999999e+06, false); err != nil {
		return err
	}

	return nil
}

func (m *DatabaseConnectionFailureDetectionConfig) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *DatabaseConnectionFailureDetectionConfig) validateTimePeriodMinutes(formats strfmt.Registry) error {

	if swag.IsZero(m.TimePeriodMinutes) { // not required
		return nil
	}

	if err := validate.MinimumInt("timePeriodMinutes", "body", int64(m.TimePeriodMinutes), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("timePeriodMinutes", "body", int64(m.TimePeriodMinutes), 15, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DatabaseConnectionFailureDetectionConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatabaseConnectionFailureDetectionConfig) UnmarshalBinary(b []byte) error {
	var res DatabaseConnectionFailureDetectionConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}