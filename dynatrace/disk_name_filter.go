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

// DiskNameFilter Narrows the rule usage down to disks, matching the specified criteria.
// swagger:model DiskNameFilter
type DiskNameFilter struct {

	// Comparison operator.
	// Required: true
	// Enum: [CONTAINS DOES_NOT_CONTAIN DOES_NOT_EQUAL DOES_NOT_START_WITH EQUALS STARTS_WITH]
	Operator *string `json:"operator"`

	// Value to compare to.
	// Required: true
	// Max Length: 100
	// Min Length: 1
	Value *string `json:"value"`
}

// Validate validates this disk name filter
func (m *DiskNameFilter) Validate(formats strfmt.Registry) error {
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

var diskNameFilterTypeOperatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CONTAINS","DOES_NOT_CONTAIN","DOES_NOT_EQUAL","DOES_NOT_START_WITH","EQUALS","STARTS_WITH"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		diskNameFilterTypeOperatorPropEnum = append(diskNameFilterTypeOperatorPropEnum, v)
	}
}

const (

	// DiskNameFilterOperatorCONTAINS captures enum value "CONTAINS"
	DiskNameFilterOperatorCONTAINS string = "CONTAINS"

	// DiskNameFilterOperatorDOESNOTCONTAIN captures enum value "DOES_NOT_CONTAIN"
	DiskNameFilterOperatorDOESNOTCONTAIN string = "DOES_NOT_CONTAIN"

	// DiskNameFilterOperatorDOESNOTEQUAL captures enum value "DOES_NOT_EQUAL"
	DiskNameFilterOperatorDOESNOTEQUAL string = "DOES_NOT_EQUAL"

	// DiskNameFilterOperatorDOESNOTSTARTWITH captures enum value "DOES_NOT_START_WITH"
	DiskNameFilterOperatorDOESNOTSTARTWITH string = "DOES_NOT_START_WITH"

	// DiskNameFilterOperatorEQUALS captures enum value "EQUALS"
	DiskNameFilterOperatorEQUALS string = "EQUALS"

	// DiskNameFilterOperatorSTARTSWITH captures enum value "STARTS_WITH"
	DiskNameFilterOperatorSTARTSWITH string = "STARTS_WITH"
)

// prop value enum
func (m *DiskNameFilter) validateOperatorEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, diskNameFilterTypeOperatorPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DiskNameFilter) validateOperator(formats strfmt.Registry) error {

	if err := validate.Required("operator", "body", m.Operator); err != nil {
		return err
	}

	// value enum
	if err := m.validateOperatorEnum("operator", "body", *m.Operator); err != nil {
		return err
	}

	return nil
}

func (m *DiskNameFilter) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	if err := validate.MinLength("value", "body", string(*m.Value), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("value", "body", string(*m.Value), 100); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DiskNameFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DiskNameFilter) UnmarshalBinary(b []byte) error {
	var res DiskNameFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
