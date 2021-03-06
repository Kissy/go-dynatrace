// Code generated by go-swagger; DO NOT EDIT.

package dynatrace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SymbolFileList symbol file list
// swagger:model SymbolFileList
type SymbolFileList struct {

	// A list of symbolication files.
	SymbolFiles []*SymbolFile `json:"symbolFiles"`
}

// Validate validates this symbol file list
func (m *SymbolFileList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSymbolFiles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SymbolFileList) validateSymbolFiles(formats strfmt.Registry) error {

	if swag.IsZero(m.SymbolFiles) { // not required
		return nil
	}

	for i := 0; i < len(m.SymbolFiles); i++ {
		if swag.IsZero(m.SymbolFiles[i]) { // not required
			continue
		}

		if m.SymbolFiles[i] != nil {
			if err := m.SymbolFiles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("symbolFiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SymbolFileList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SymbolFileList) UnmarshalBinary(b []byte) error {
	var res SymbolFileList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
