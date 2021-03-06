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

// EventWrapperSettings In addition to the event handlers, events called using `addEventListener` or `attachEvent` can be captured. Be careful with this option! Event wrappers can conflict with the JavaScript code on a web page.
// swagger:model EventWrapperSettings
type EventWrapperSettings struct {

	// Blur enabled/disabled.
	// Required: true
	Blur *bool `json:"blur"`

	// Change enabled/disabled.
	// Required: true
	Change *bool `json:"change"`

	// Click enabled/disabled.
	// Required: true
	Click *bool `json:"click"`

	// MouseUp enabled/disabled.
	// Required: true
	MouseUp *bool `json:"mouseUp"`

	// TouchEnd enabled/disabled.
	// Required: true
	TouchEnd *bool `json:"touchEnd"`

	// TouchStart enabled/disabled.
	// Required: true
	TouchStart *bool `json:"touchStart"`
}

// Validate validates this event wrapper settings
func (m *EventWrapperSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBlur(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChange(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClick(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMouseUp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTouchEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTouchStart(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EventWrapperSettings) validateBlur(formats strfmt.Registry) error {

	if err := validate.Required("blur", "body", m.Blur); err != nil {
		return err
	}

	return nil
}

func (m *EventWrapperSettings) validateChange(formats strfmt.Registry) error {

	if err := validate.Required("change", "body", m.Change); err != nil {
		return err
	}

	return nil
}

func (m *EventWrapperSettings) validateClick(formats strfmt.Registry) error {

	if err := validate.Required("click", "body", m.Click); err != nil {
		return err
	}

	return nil
}

func (m *EventWrapperSettings) validateMouseUp(formats strfmt.Registry) error {

	if err := validate.Required("mouseUp", "body", m.MouseUp); err != nil {
		return err
	}

	return nil
}

func (m *EventWrapperSettings) validateTouchEnd(formats strfmt.Registry) error {

	if err := validate.Required("touchEnd", "body", m.TouchEnd); err != nil {
		return err
	}

	return nil
}

func (m *EventWrapperSettings) validateTouchStart(formats strfmt.Registry) error {

	if err := validate.Required("touchStart", "body", m.TouchStart); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EventWrapperSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventWrapperSettings) UnmarshalBinary(b []byte) error {
	var res EventWrapperSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
