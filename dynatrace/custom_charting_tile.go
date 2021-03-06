// Code generated by go-swagger; DO NOT EDIT.

package dynatrace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CustomChartingTile custom charting tile
// swagger:model CustomChartingTile
type CustomChartingTile struct {
	boundsField *TileBounds

	configuredField bool

	nameField *string

	tileFilterField *TileFilter

	// How the data is filtered for charting.
	FilterConfig *CustomFilterConfig `json:"filterConfig,omitempty"`
}

// Bounds gets the bounds of this subtype
func (m *CustomChartingTile) Bounds() *TileBounds {
	return m.boundsField
}

// SetBounds sets the bounds of this subtype
func (m *CustomChartingTile) SetBounds(val *TileBounds) {
	m.boundsField = val
}

// Configured gets the configured of this subtype
func (m *CustomChartingTile) Configured() bool {
	return m.configuredField
}

// SetConfigured sets the configured of this subtype
func (m *CustomChartingTile) SetConfigured(val bool) {
	m.configuredField = val
}

// Name gets the name of this subtype
func (m *CustomChartingTile) Name() *string {
	return m.nameField
}

// SetName sets the name of this subtype
func (m *CustomChartingTile) SetName(val *string) {
	m.nameField = val
}

// TileFilter gets the tile filter of this subtype
func (m *CustomChartingTile) TileFilter() *TileFilter {
	return m.tileFilterField
}

// SetTileFilter sets the tile filter of this subtype
func (m *CustomChartingTile) SetTileFilter(val *TileFilter) {
	m.tileFilterField = val
}

// TileType gets the tile type of this subtype
func (m *CustomChartingTile) TileType() string {
	return "CustomChartingTile"
}

// SetTileType sets the tile type of this subtype
func (m *CustomChartingTile) SetTileType(val string) {

}

// FilterConfig gets the filter config of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *CustomChartingTile) UnmarshalJSON(raw []byte) error {
	var data struct {

		// How the data is filtered for charting.
		FilterConfig *CustomFilterConfig `json:"filterConfig,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Bounds *TileBounds `json:"bounds"`

		Configured bool `json:"configured,omitempty"`

		Name *string `json:"name"`

		TileFilter *TileFilter `json:"tileFilter,omitempty"`

		TileType string `json:"tileType"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result CustomChartingTile

	result.boundsField = base.Bounds

	result.configuredField = base.Configured

	result.nameField = base.Name

	result.tileFilterField = base.TileFilter

	if base.TileType != result.TileType() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid tileType value: %q", base.TileType)
	}

	result.FilterConfig = data.FilterConfig

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m CustomChartingTile) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// How the data is filtered for charting.
		FilterConfig *CustomFilterConfig `json:"filterConfig,omitempty"`
	}{

		FilterConfig: m.FilterConfig,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Bounds *TileBounds `json:"bounds"`

		Configured bool `json:"configured,omitempty"`

		Name *string `json:"name"`

		TileFilter *TileFilter `json:"tileFilter,omitempty"`

		TileType string `json:"tileType"`
	}{

		Bounds: m.Bounds(),

		Configured: m.Configured(),

		Name: m.Name(),

		TileFilter: m.TileFilter(),

		TileType: m.TileType(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this custom charting tile
func (m *CustomChartingTile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBounds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTileFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilterConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomChartingTile) validateBounds(formats strfmt.Registry) error {

	if err := validate.Required("bounds", "body", m.Bounds()); err != nil {
		return err
	}

	if m.Bounds() != nil {
		if err := m.Bounds().Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bounds")
			}
			return err
		}
	}

	return nil
}

func (m *CustomChartingTile) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name()); err != nil {
		return err
	}

	return nil
}

func (m *CustomChartingTile) validateTileFilter(formats strfmt.Registry) error {

	if swag.IsZero(m.TileFilter()) { // not required
		return nil
	}

	if m.TileFilter() != nil {
		if err := m.TileFilter().Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tileFilter")
			}
			return err
		}
	}

	return nil
}

func (m *CustomChartingTile) validateFilterConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.FilterConfig) { // not required
		return nil
	}

	if m.FilterConfig != nil {
		if err := m.FilterConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filterConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustomChartingTile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomChartingTile) UnmarshalBinary(b []byte) error {
	var res CustomChartingTile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
