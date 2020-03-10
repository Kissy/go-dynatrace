// Code generated by go-swagger; DO NOT EDIT.

package dynatrace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SharingInfo Sharing configuration of a dashboard.
// swagger:model SharingInfo
type SharingInfo struct {

	// If `true`, the dashboard is shared via link and authenticated users with the link can view.
	LinkShared bool `json:"linkShared,omitempty"`

	// If `true`, the dashboard is published to anyone on this environment.
	Published bool `json:"published,omitempty"`
}

// Validate validates this sharing info
func (m *SharingInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SharingInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SharingInfo) UnmarshalBinary(b []byte) error {
	var res SharingInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}