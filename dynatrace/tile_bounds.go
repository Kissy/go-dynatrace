// Code generated by go-swagger; DO NOT EDIT.

package dynatrace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TileBounds The position and size of a tile.
// swagger:model TileBounds
type TileBounds struct {

	// The height of the tile, in pixels.
	Height int32 `json:"height,omitempty"`

	// The horizontal distance from the top left corner of the dashboard to the top left corner of the tile, in pixels.
	Left int32 `json:"left,omitempty"`

	// The vertical distance from the top left corner of the dashboard to the top left corner of the tile, in pixels.
	Top int32 `json:"top,omitempty"`

	// The width of the tile, in pixels.
	Width int32 `json:"width,omitempty"`
}

// Validate validates this tile bounds
func (m *TileBounds) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TileBounds) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TileBounds) UnmarshalBinary(b []byte) error {
	var res TileBounds
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
