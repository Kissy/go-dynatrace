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

// AdditionalEventHandlers Additional event handlers and wrappers.
// swagger:model AdditionalEventHandlers
type AdditionalEventHandlers struct {

	// Blur event handler enabled/disabled.
	// Required: true
	BlurEventHandler *bool `json:"blurEventHandler"`

	// Change event handler enabled/disabled.
	// Required: true
	ChangeEventHandler *bool `json:"changeEventHandler"`

	// Click event handler enabled/disabled.
	// Required: true
	ClickEventHandler *bool `json:"clickEventHandler"`

	// Max. number of DOM nodes to instrument. Valid values range from 0 to 100000.
	// Required: true
	// Maximum: 100000
	// Minimum: 0
	MaxDomNodesToInstrument *int32 `json:"maxDomNodesToInstrument"`

	// Mouseup event handler enabled/disabled.
	// Required: true
	MouseupEventHandler *bool `json:"mouseupEventHandler"`

	// toString method enabled/disabled.
	// Required: true
	ToStringMethod *bool `json:"toStringMethod"`

	// Use mouseup event for clicks enabled/disabled.
	// Required: true
	UserMouseupEventForClicks *bool `json:"userMouseupEventForClicks"`
}

// Validate validates this additional event handlers
func (m *AdditionalEventHandlers) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBlurEventHandler(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChangeEventHandler(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClickEventHandler(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxDomNodesToInstrument(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMouseupEventHandler(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToStringMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserMouseupEventForClicks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AdditionalEventHandlers) validateBlurEventHandler(formats strfmt.Registry) error {

	if err := validate.Required("blurEventHandler", "body", m.BlurEventHandler); err != nil {
		return err
	}

	return nil
}

func (m *AdditionalEventHandlers) validateChangeEventHandler(formats strfmt.Registry) error {

	if err := validate.Required("changeEventHandler", "body", m.ChangeEventHandler); err != nil {
		return err
	}

	return nil
}

func (m *AdditionalEventHandlers) validateClickEventHandler(formats strfmt.Registry) error {

	if err := validate.Required("clickEventHandler", "body", m.ClickEventHandler); err != nil {
		return err
	}

	return nil
}

func (m *AdditionalEventHandlers) validateMaxDomNodesToInstrument(formats strfmt.Registry) error {

	if err := validate.Required("maxDomNodesToInstrument", "body", m.MaxDomNodesToInstrument); err != nil {
		return err
	}

	if err := validate.MinimumInt("maxDomNodesToInstrument", "body", int64(*m.MaxDomNodesToInstrument), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("maxDomNodesToInstrument", "body", int64(*m.MaxDomNodesToInstrument), 100000, false); err != nil {
		return err
	}

	return nil
}

func (m *AdditionalEventHandlers) validateMouseupEventHandler(formats strfmt.Registry) error {

	if err := validate.Required("mouseupEventHandler", "body", m.MouseupEventHandler); err != nil {
		return err
	}

	return nil
}

func (m *AdditionalEventHandlers) validateToStringMethod(formats strfmt.Registry) error {

	if err := validate.Required("toStringMethod", "body", m.ToStringMethod); err != nil {
		return err
	}

	return nil
}

func (m *AdditionalEventHandlers) validateUserMouseupEventForClicks(formats strfmt.Registry) error {

	if err := validate.Required("userMouseupEventForClicks", "body", m.UserMouseupEventForClicks); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AdditionalEventHandlers) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdditionalEventHandlers) UnmarshalBinary(b []byte) error {
	var res AdditionalEventHandlers
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
