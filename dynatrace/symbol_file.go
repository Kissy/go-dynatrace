// Code generated by go-swagger; DO NOT EDIT.

package dynatrace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SymbolFile symbol file
// swagger:model SymbolFile
type SymbolFile struct {

	// The appId, the app version and the bundle id which uniquely identify the app
	AppID *AppIdentifier `json:"appId,omitempty"`

	// The name of the application this file belongs to
	ApplicationName string `json:"applicationName,omitempty"`

	// Is the file pinned and therefore cannot be deleted.
	Pinned bool `json:"pinned,omitempty"`

	// The size of the file in KB
	Size int32 `json:"size,omitempty"`

	// The timestamp of the upload time of the file, in UTC milliseconds
	UploadTimestamp int64 `json:"uploadTimestamp,omitempty"`
}

// Validate validates this symbol file
func (m *SymbolFile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SymbolFile) validateAppID(formats strfmt.Registry) error {

	if swag.IsZero(m.AppID) { // not required
		return nil
	}

	if m.AppID != nil {
		if err := m.AppID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appId")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SymbolFile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SymbolFile) UnmarshalBinary(b []byte) error {
	var res SymbolFile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
