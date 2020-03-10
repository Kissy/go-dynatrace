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

// MaintenanceWindow Configuration of a maintenance window.
// swagger:model MaintenanceWindow
type MaintenanceWindow struct {

	// A short description of the maintenance purpose.
	// Required: true
	Description *string `json:"description"`

	// The ID of the maintenance window.
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Metadata useful for debugging.
	// Read Only: true
	Metadata *ConfigurationMetadata `json:"metadata,omitempty"`

	// The name of the maintenance window, displayed in the UI.
	// Required: true
	// Min Length: 1
	Name *string `json:"name"`

	// The date, time, and recurrence of the maintenance window.
	// Required: true
	Schedule *Schedule `json:"schedule"`

	// The scope of the maintenance window.
	//
	//  The scope restricts the alert/problem detection suppression to certain Dynatrace entities. It can contain a list of entities and/or matching rules for dynamic formation of the scope.
	//
	//  If no scope is specified, the alert/problem detection suppression applies to the entire environment.
	Scope *Scope `json:"scope,omitempty"`

	// The type of suppression of alerting and problem detection during the maintenance.
	// Required: true
	// Enum: [DETECT_PROBLEMS_AND_ALERT DETECT_PROBLEMS_DONT_ALERT DONT_DETECT_PROBLEMS]
	Suppression *string `json:"suppression"`

	// The type of the maintenance: planned or unplanned.
	// Required: true
	// Enum: [PLANNED UNPLANNED]
	Type *string `json:"type"`
}

// Validate validates this maintenance window
func (m *MaintenanceWindow) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchedule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSuppression(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MaintenanceWindow) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *MaintenanceWindow) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MaintenanceWindow) validateMetadata(formats strfmt.Registry) error {

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

func (m *MaintenanceWindow) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	return nil
}

func (m *MaintenanceWindow) validateSchedule(formats strfmt.Registry) error {

	if err := validate.Required("schedule", "body", m.Schedule); err != nil {
		return err
	}

	if m.Schedule != nil {
		if err := m.Schedule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

func (m *MaintenanceWindow) validateScope(formats strfmt.Registry) error {

	if swag.IsZero(m.Scope) { // not required
		return nil
	}

	if m.Scope != nil {
		if err := m.Scope.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scope")
			}
			return err
		}
	}

	return nil
}

var maintenanceWindowTypeSuppressionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DETECT_PROBLEMS_AND_ALERT","DETECT_PROBLEMS_DONT_ALERT","DONT_DETECT_PROBLEMS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		maintenanceWindowTypeSuppressionPropEnum = append(maintenanceWindowTypeSuppressionPropEnum, v)
	}
}

const (

	// MaintenanceWindowSuppressionDETECTPROBLEMSANDALERT captures enum value "DETECT_PROBLEMS_AND_ALERT"
	MaintenanceWindowSuppressionDETECTPROBLEMSANDALERT string = "DETECT_PROBLEMS_AND_ALERT"

	// MaintenanceWindowSuppressionDETECTPROBLEMSDONTALERT captures enum value "DETECT_PROBLEMS_DONT_ALERT"
	MaintenanceWindowSuppressionDETECTPROBLEMSDONTALERT string = "DETECT_PROBLEMS_DONT_ALERT"

	// MaintenanceWindowSuppressionDONTDETECTPROBLEMS captures enum value "DONT_DETECT_PROBLEMS"
	MaintenanceWindowSuppressionDONTDETECTPROBLEMS string = "DONT_DETECT_PROBLEMS"
)

// prop value enum
func (m *MaintenanceWindow) validateSuppressionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, maintenanceWindowTypeSuppressionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MaintenanceWindow) validateSuppression(formats strfmt.Registry) error {

	if err := validate.Required("suppression", "body", m.Suppression); err != nil {
		return err
	}

	// value enum
	if err := m.validateSuppressionEnum("suppression", "body", *m.Suppression); err != nil {
		return err
	}

	return nil
}

var maintenanceWindowTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PLANNED","UNPLANNED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		maintenanceWindowTypeTypePropEnum = append(maintenanceWindowTypeTypePropEnum, v)
	}
}

const (

	// MaintenanceWindowTypePLANNED captures enum value "PLANNED"
	MaintenanceWindowTypePLANNED string = "PLANNED"

	// MaintenanceWindowTypeUNPLANNED captures enum value "UNPLANNED"
	MaintenanceWindowTypeUNPLANNED string = "UNPLANNED"
)

// prop value enum
func (m *MaintenanceWindow) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, maintenanceWindowTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MaintenanceWindow) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MaintenanceWindow) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MaintenanceWindow) UnmarshalBinary(b []byte) error {
	var res MaintenanceWindow
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}