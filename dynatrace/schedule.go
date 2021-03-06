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

// Schedule The schedule of the maintenance window.
// swagger:model Schedule
type Schedule struct {

	// The end date and time of the maintenance window validity period in yyyy-mm-dd HH:mm format.
	// Required: true
	End *string `json:"end"`

	// The recurrence of the schedule.
	//
	//  Not applicable if **recurrenceType** is `Once`.
	Recurrence *Recurrence `json:"recurrence,omitempty"`

	// The type of the schedule recurrence.
	// Required: true
	// Enum: [DAILY MONTHLY ONCE WEEKLY]
	RecurrenceType *string `json:"recurrenceType"`

	// The start date and time of the maintenance window validity period in yyyy-mm-dd HH:mm format.
	// Required: true
	Start *string `json:"start"`

	// The time zone of the start and end time. Default time zone is UTC.
	//
	// You can use either UTC offset `UTC+01:00` format or the IANA Time Zone Database format (for example, `Europe/Vienna`).
	// Required: true
	ZoneID *string `json:"zoneId"`
}

// Validate validates this schedule
func (m *Schedule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecurrence(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecurrenceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZoneID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Schedule) validateEnd(formats strfmt.Registry) error {

	if err := validate.Required("end", "body", m.End); err != nil {
		return err
	}

	return nil
}

func (m *Schedule) validateRecurrence(formats strfmt.Registry) error {

	if swag.IsZero(m.Recurrence) { // not required
		return nil
	}

	if m.Recurrence != nil {
		if err := m.Recurrence.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recurrence")
			}
			return err
		}
	}

	return nil
}

var scheduleTypeRecurrenceTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DAILY","MONTHLY","ONCE","WEEKLY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		scheduleTypeRecurrenceTypePropEnum = append(scheduleTypeRecurrenceTypePropEnum, v)
	}
}

const (

	// ScheduleRecurrenceTypeDAILY captures enum value "DAILY"
	ScheduleRecurrenceTypeDAILY string = "DAILY"

	// ScheduleRecurrenceTypeMONTHLY captures enum value "MONTHLY"
	ScheduleRecurrenceTypeMONTHLY string = "MONTHLY"

	// ScheduleRecurrenceTypeONCE captures enum value "ONCE"
	ScheduleRecurrenceTypeONCE string = "ONCE"

	// ScheduleRecurrenceTypeWEEKLY captures enum value "WEEKLY"
	ScheduleRecurrenceTypeWEEKLY string = "WEEKLY"
)

// prop value enum
func (m *Schedule) validateRecurrenceTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, scheduleTypeRecurrenceTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Schedule) validateRecurrenceType(formats strfmt.Registry) error {

	if err := validate.Required("recurrenceType", "body", m.RecurrenceType); err != nil {
		return err
	}

	// value enum
	if err := m.validateRecurrenceTypeEnum("recurrenceType", "body", *m.RecurrenceType); err != nil {
		return err
	}

	return nil
}

func (m *Schedule) validateStart(formats strfmt.Registry) error {

	if err := validate.Required("start", "body", m.Start); err != nil {
		return err
	}

	return nil
}

func (m *Schedule) validateZoneID(formats strfmt.Registry) error {

	if err := validate.Required("zoneId", "body", m.ZoneID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Schedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Schedule) UnmarshalBinary(b []byte) error {
	var res Schedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
