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

// MetricEventTextFilterMetricEventDimensionsFilterOperatorDto A filter for a string value based on the given operator.
// swagger:model MetricEventTextFilterMetricEventDimensionsFilterOperatorDto
type MetricEventTextFilterMetricEventDimensionsFilterOperatorDto struct {

	// The operator to match on.
	// Required: true
	// Enum: [EQUALS]
	Operator *string `json:"operator"`

	// The value to match on.
	// Required: true
	// Max Length: 1000
	// Min Length: 1
	Value *string `json:"value"`
}

// Validate validates this metric event text filter metric event dimensions filter operator dto
func (m *MetricEventTextFilterMetricEventDimensionsFilterOperatorDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var metricEventTextFilterMetricEventDimensionsFilterOperatorDtoTypeOperatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["EQUALS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		metricEventTextFilterMetricEventDimensionsFilterOperatorDtoTypeOperatorPropEnum = append(metricEventTextFilterMetricEventDimensionsFilterOperatorDtoTypeOperatorPropEnum, v)
	}
}

const (

	// MetricEventTextFilterMetricEventDimensionsFilterOperatorDtoOperatorEQUALS captures enum value "EQUALS"
	MetricEventTextFilterMetricEventDimensionsFilterOperatorDtoOperatorEQUALS string = "EQUALS"
)

// prop value enum
func (m *MetricEventTextFilterMetricEventDimensionsFilterOperatorDto) validateOperatorEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, metricEventTextFilterMetricEventDimensionsFilterOperatorDtoTypeOperatorPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MetricEventTextFilterMetricEventDimensionsFilterOperatorDto) validateOperator(formats strfmt.Registry) error {

	if err := validate.Required("operator", "body", m.Operator); err != nil {
		return err
	}

	// value enum
	if err := m.validateOperatorEnum("operator", "body", *m.Operator); err != nil {
		return err
	}

	return nil
}

func (m *MetricEventTextFilterMetricEventDimensionsFilterOperatorDto) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	if err := validate.MinLength("value", "body", string(*m.Value), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("value", "body", string(*m.Value), 1000); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MetricEventTextFilterMetricEventDimensionsFilterOperatorDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetricEventTextFilterMetricEventDimensionsFilterOperatorDto) UnmarshalBinary(b []byte) error {
	var res MetricEventTextFilterMetricEventDimensionsFilterOperatorDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}