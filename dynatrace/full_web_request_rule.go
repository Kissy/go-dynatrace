// Code generated by go-swagger; DO NOT EDIT.

package dynatrace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FullWebRequestRule The service detection rule of the `FULL_WEB_REQUEST` type.
// swagger:model FullWebRequestRule
type FullWebRequestRule struct {

	// How to handle the detected ID of the application.
	ApplicationID *ApplicationID `json:"applicationId,omitempty"`

	// A list of conditions of the rule.
	//
	// If several conditions are specified, the AND logic applies.
	// Max Items: 10
	// Min Items: 0
	Conditions []*ConditionsFullWebRequestAttributeTypeDto `json:"conditions"`

	// How to handle the detected context root of the request URL.
	ContextRoot *ContextRoot `json:"contextRoot,omitempty"`

	// A short description of the rule.
	Description string `json:"description,omitempty"`

	// The rule is enabled(`true`) or disabled (`false`).
	// Required: true
	Enabled *bool `json:"enabled"`

	// The ID of the service detection rule.
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Specifies the management zones of the process group for which this service detection rule should be created.
	// Max Items: 1
	// Min Items: 0
	ManagementZones []string `json:"managementZones"`

	// Metadata useful for debugging.
	// Read Only: true
	Metadata *ConfigurationMetadata `json:"metadata,omitempty"`

	// The name of the rule.
	// Required: true
	// Min Length: 1
	Name *string `json:"name"`

	// The order of the rule in the rules list.
	//
	//  The rules are evaluated from top to bottom. The first matching rule applies.
	// Max Length: 2147483647
	// Min Length: 1
	Order string `json:"order,omitempty"`

	// How to handle the detected server name of the request URL.
	ServerName *ServerName `json:"serverName,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this full web request rule
func (m *FullWebRequestRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplicationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContextRoot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManagementZones(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrder(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FullWebRequestRule) validateApplicationID(formats strfmt.Registry) error {

	if swag.IsZero(m.ApplicationID) { // not required
		return nil
	}

	if m.ApplicationID != nil {
		if err := m.ApplicationID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("applicationId")
			}
			return err
		}
	}

	return nil
}

func (m *FullWebRequestRule) validateConditions(formats strfmt.Registry) error {

	if swag.IsZero(m.Conditions) { // not required
		return nil
	}

	iConditionsSize := int64(len(m.Conditions))

	if err := validate.MinItems("conditions", "body", iConditionsSize, 0); err != nil {
		return err
	}

	if err := validate.MaxItems("conditions", "body", iConditionsSize, 10); err != nil {
		return err
	}

	for i := 0; i < len(m.Conditions); i++ {
		if swag.IsZero(m.Conditions[i]) { // not required
			continue
		}

		if m.Conditions[i] != nil {
			if err := m.Conditions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FullWebRequestRule) validateContextRoot(formats strfmt.Registry) error {

	if swag.IsZero(m.ContextRoot) { // not required
		return nil
	}

	if m.ContextRoot != nil {
		if err := m.ContextRoot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contextRoot")
			}
			return err
		}
	}

	return nil
}

func (m *FullWebRequestRule) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *FullWebRequestRule) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *FullWebRequestRule) validateManagementZones(formats strfmt.Registry) error {

	if swag.IsZero(m.ManagementZones) { // not required
		return nil
	}

	iManagementZonesSize := int64(len(m.ManagementZones))

	if err := validate.MinItems("managementZones", "body", iManagementZonesSize, 0); err != nil {
		return err
	}

	if err := validate.MaxItems("managementZones", "body", iManagementZonesSize, 1); err != nil {
		return err
	}

	return nil
}

func (m *FullWebRequestRule) validateMetadata(formats strfmt.Registry) error {

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

func (m *FullWebRequestRule) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	return nil
}

func (m *FullWebRequestRule) validateOrder(formats strfmt.Registry) error {

	if swag.IsZero(m.Order) { // not required
		return nil
	}

	if err := validate.MinLength("order", "body", string(m.Order), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("order", "body", string(m.Order), 2147483647); err != nil {
		return err
	}

	return nil
}

func (m *FullWebRequestRule) validateServerName(formats strfmt.Registry) error {

	if swag.IsZero(m.ServerName) { // not required
		return nil
	}

	if m.ServerName != nil {
		if err := m.ServerName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("serverName")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FullWebRequestRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FullWebRequestRule) UnmarshalBinary(b []byte) error {
	var res FullWebRequestRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
