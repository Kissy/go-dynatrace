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

// ApplicationAnomalyDetectionConfig The configuration of anomaly detection for applications.
// swagger:model ApplicationAnomalyDetectionConfig
type ApplicationAnomalyDetectionConfig struct {

	// How to detect failure rate increase.
	// Required: true
	FailureRateIncrease *FailureRateIncreaseDetectionConfig `json:"failureRateIncrease"`

	// Metadata useful for debugging.
	// Read Only: true
	Metadata *ConfigurationMetadata `json:"metadata,omitempty"`

	// How to detect response time degradation.
	// Required: true
	ResponseTimeDegradation *ResponseTimeDegradationDetectionConfig `json:"responseTimeDegradation"`

	// How to detect traffic drops.
	// Required: true
	TrafficDrop *TrafficDropDetectionConfig `json:"trafficDrop"`

	// How to detect traffic spikes.
	// Required: true
	TrafficSpike *TrafficSpikeDetectionConfig `json:"trafficSpike"`
}

// Validate validates this application anomaly detection config
func (m *ApplicationAnomalyDetectionConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFailureRateIncrease(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponseTimeDegradation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrafficDrop(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrafficSpike(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationAnomalyDetectionConfig) validateFailureRateIncrease(formats strfmt.Registry) error {

	if err := validate.Required("failureRateIncrease", "body", m.FailureRateIncrease); err != nil {
		return err
	}

	if m.FailureRateIncrease != nil {
		if err := m.FailureRateIncrease.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("failureRateIncrease")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationAnomalyDetectionConfig) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationAnomalyDetectionConfig) validateResponseTimeDegradation(formats strfmt.Registry) error {

	if err := validate.Required("responseTimeDegradation", "body", m.ResponseTimeDegradation); err != nil {
		return err
	}

	if m.ResponseTimeDegradation != nil {
		if err := m.ResponseTimeDegradation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("responseTimeDegradation")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationAnomalyDetectionConfig) validateTrafficDrop(formats strfmt.Registry) error {

	if err := validate.Required("trafficDrop", "body", m.TrafficDrop); err != nil {
		return err
	}

	if m.TrafficDrop != nil {
		if err := m.TrafficDrop.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trafficDrop")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationAnomalyDetectionConfig) validateTrafficSpike(formats strfmt.Registry) error {

	if err := validate.Required("trafficSpike", "body", m.TrafficSpike); err != nil {
		return err
	}

	if m.TrafficSpike != nil {
		if err := m.TrafficSpike.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trafficSpike")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationAnomalyDetectionConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationAnomalyDetectionConfig) UnmarshalBinary(b []byte) error {
	var res ApplicationAnomalyDetectionConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}