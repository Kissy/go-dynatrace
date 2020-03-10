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

// NotificationConfigStub The short representation of a notification.
// swagger:model NotificationConfigStub
type NotificationConfigStub struct {

	// A short description of the Dynatrace entity.
	// Read Only: true
	Description string `json:"description,omitempty"`

	// The ID of the Dynatrace entity.
	// Required: true
	ID *string `json:"id"`

	// The name of the Dynatrace entity.
	// Read Only: true
	Name string `json:"name,omitempty"`

	// The type of the notification.
	// Enum: [ANSIBLETOWER EMAIL HIPCHAT JIRA OPS_GENIE PAGER_DUTY SERVICE_NOW SLACK TRELLO VICTOROPS WEBHOOK XMATTERS]
	Type string `json:"type,omitempty"`
}

// Validate validates this notification config stub
func (m *NotificationConfigStub) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
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

func (m *NotificationConfigStub) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

var notificationConfigStubTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ANSIBLETOWER","EMAIL","HIPCHAT","JIRA","OPS_GENIE","PAGER_DUTY","SERVICE_NOW","SLACK","TRELLO","VICTOROPS","WEBHOOK","XMATTERS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		notificationConfigStubTypeTypePropEnum = append(notificationConfigStubTypeTypePropEnum, v)
	}
}

const (

	// NotificationConfigStubTypeANSIBLETOWER captures enum value "ANSIBLETOWER"
	NotificationConfigStubTypeANSIBLETOWER string = "ANSIBLETOWER"

	// NotificationConfigStubTypeEMAIL captures enum value "EMAIL"
	NotificationConfigStubTypeEMAIL string = "EMAIL"

	// NotificationConfigStubTypeHIPCHAT captures enum value "HIPCHAT"
	NotificationConfigStubTypeHIPCHAT string = "HIPCHAT"

	// NotificationConfigStubTypeJIRA captures enum value "JIRA"
	NotificationConfigStubTypeJIRA string = "JIRA"

	// NotificationConfigStubTypeOPSGENIE captures enum value "OPS_GENIE"
	NotificationConfigStubTypeOPSGENIE string = "OPS_GENIE"

	// NotificationConfigStubTypePAGERDUTY captures enum value "PAGER_DUTY"
	NotificationConfigStubTypePAGERDUTY string = "PAGER_DUTY"

	// NotificationConfigStubTypeSERVICENOW captures enum value "SERVICE_NOW"
	NotificationConfigStubTypeSERVICENOW string = "SERVICE_NOW"

	// NotificationConfigStubTypeSLACK captures enum value "SLACK"
	NotificationConfigStubTypeSLACK string = "SLACK"

	// NotificationConfigStubTypeTRELLO captures enum value "TRELLO"
	NotificationConfigStubTypeTRELLO string = "TRELLO"

	// NotificationConfigStubTypeVICTOROPS captures enum value "VICTOROPS"
	NotificationConfigStubTypeVICTOROPS string = "VICTOROPS"

	// NotificationConfigStubTypeWEBHOOK captures enum value "WEBHOOK"
	NotificationConfigStubTypeWEBHOOK string = "WEBHOOK"

	// NotificationConfigStubTypeXMATTERS captures enum value "XMATTERS"
	NotificationConfigStubTypeXMATTERS string = "XMATTERS"
)

// prop value enum
func (m *NotificationConfigStub) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, notificationConfigStubTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *NotificationConfigStub) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NotificationConfigStub) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NotificationConfigStub) UnmarshalBinary(b []byte) error {
	var res NotificationConfigStub
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}