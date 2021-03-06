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

// UndersizedStorageThresholds Custom thresholds for undersized storage device. If not set then the automatic mode is used.
//
//  Fulfillment of **any** condition triggers an alert.
// swagger:model UndersizedStorageThresholds
type UndersizedStorageThresholds struct {

	// Average queue command latency is higher than *X* milliseconds in 3 out of 5 samples.
	// Required: true
	// Maximum: 10000
	// Minimum: 1
	AverageQueueCommandLatency *int32 `json:"averageQueueCommandLatency"`

	// Peak queue command latency is higher than *X* milliseconds in 3 out of 5 samples.
	// Required: true
	// Maximum: 10000
	// Minimum: 1
	PeakQueueCommandLatency *int32 `json:"peakQueueCommandLatency"`
}

// Validate validates this undersized storage thresholds
func (m *UndersizedStorageThresholds) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAverageQueueCommandLatency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeakQueueCommandLatency(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UndersizedStorageThresholds) validateAverageQueueCommandLatency(formats strfmt.Registry) error {

	if err := validate.Required("averageQueueCommandLatency", "body", m.AverageQueueCommandLatency); err != nil {
		return err
	}

	if err := validate.MinimumInt("averageQueueCommandLatency", "body", int64(*m.AverageQueueCommandLatency), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("averageQueueCommandLatency", "body", int64(*m.AverageQueueCommandLatency), 10000, false); err != nil {
		return err
	}

	return nil
}

func (m *UndersizedStorageThresholds) validatePeakQueueCommandLatency(formats strfmt.Registry) error {

	if err := validate.Required("peakQueueCommandLatency", "body", m.PeakQueueCommandLatency); err != nil {
		return err
	}

	if err := validate.MinimumInt("peakQueueCommandLatency", "body", int64(*m.PeakQueueCommandLatency), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("peakQueueCommandLatency", "body", int64(*m.PeakQueueCommandLatency), 10000, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UndersizedStorageThresholds) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UndersizedStorageThresholds) UnmarshalBinary(b []byte) error {
	var res UndersizedStorageThresholds
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
