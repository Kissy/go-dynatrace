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

// FailureRateIncreaseAutodetectionConfig Parameters of failure rate increase auto-detection. Required if **detectionMode** is `DETECT_AUTOMATICALLY`. Not applicable otherwise.
//
// The absolute and relative thresholds **both** must exceed to trigger an alert.
//
// Example: If the expected error rate is 1.5%, and you set an absolute increase of 1%, and a relative increase of 50%, the thresholds will be:
// Absolute: 1.5% + **1%** = 2.5%
// Relative: 1.5% + 1.5% * **50%** = 2.25%
// swagger:model FailureRateIncreaseAutodetectionConfig
type FailureRateIncreaseAutodetectionConfig struct {

	// Absolute increase of failing service calls to trigger an alert, %.
	// Required: true
	// Maximum: 1000
	// Minimum: 0
	FailingServiceCallPercentageIncreaseAbsolute *int32 `json:"failingServiceCallPercentageIncreaseAbsolute"`

	// Relative increase of failing service calls to trigger an alert, %.
	// Required: true
	// Maximum: 1000
	// Minimum: 0
	FailingServiceCallPercentageIncreaseRelative *int32 `json:"failingServiceCallPercentageIncreaseRelative"`
}

// Validate validates this failure rate increase autodetection config
func (m *FailureRateIncreaseAutodetectionConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFailingServiceCallPercentageIncreaseAbsolute(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFailingServiceCallPercentageIncreaseRelative(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FailureRateIncreaseAutodetectionConfig) validateFailingServiceCallPercentageIncreaseAbsolute(formats strfmt.Registry) error {

	if err := validate.Required("failingServiceCallPercentageIncreaseAbsolute", "body", m.FailingServiceCallPercentageIncreaseAbsolute); err != nil {
		return err
	}

	if err := validate.MinimumInt("failingServiceCallPercentageIncreaseAbsolute", "body", int64(*m.FailingServiceCallPercentageIncreaseAbsolute), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("failingServiceCallPercentageIncreaseAbsolute", "body", int64(*m.FailingServiceCallPercentageIncreaseAbsolute), 1000, false); err != nil {
		return err
	}

	return nil
}

func (m *FailureRateIncreaseAutodetectionConfig) validateFailingServiceCallPercentageIncreaseRelative(formats strfmt.Registry) error {

	if err := validate.Required("failingServiceCallPercentageIncreaseRelative", "body", m.FailingServiceCallPercentageIncreaseRelative); err != nil {
		return err
	}

	if err := validate.MinimumInt("failingServiceCallPercentageIncreaseRelative", "body", int64(*m.FailingServiceCallPercentageIncreaseRelative), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("failingServiceCallPercentageIncreaseRelative", "body", int64(*m.FailingServiceCallPercentageIncreaseRelative), 1000, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FailureRateIncreaseAutodetectionConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FailureRateIncreaseAutodetectionConfig) UnmarshalBinary(b []byte) error {
	var res FailureRateIncreaseAutodetectionConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
