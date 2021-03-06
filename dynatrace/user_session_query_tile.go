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

// UserSessionQueryTile user session query tile
// swagger:model UserSessionQueryTile
type UserSessionQueryTile struct {
	boundsField *TileBounds

	configuredField bool

	nameField *string

	tileFilterField *TileFilter

	// The name of the tile, set by user.
	// Required: true
	CustomName *string `json:"customName"`

	// The limit of the results, if not set will use the default value of the system
	Limit int32 `json:"limit,omitempty"`

	// A [user session query](https://www.dynatrace.com/support/help/shortlink/usql-info) executed by the tile.
	// Required: true
	Query *string `json:"query"`

	// The comparison timeframe of the query.
	//
	//  If specified, you additionally get the results of the same query with the specified time shift.
	TimeFrameShift string `json:"timeFrameShift,omitempty"`

	// The visualization of the tile.
	// Required: true
	// Enum: [COLUMN_CHART FUNNEL LINE_CHART PIE_CHART SINGLE_VALUE TABLE]
	Type *string `json:"type"`
}

// Bounds gets the bounds of this subtype
func (m *UserSessionQueryTile) Bounds() *TileBounds {
	return m.boundsField
}

// SetBounds sets the bounds of this subtype
func (m *UserSessionQueryTile) SetBounds(val *TileBounds) {
	m.boundsField = val
}

// Configured gets the configured of this subtype
func (m *UserSessionQueryTile) Configured() bool {
	return m.configuredField
}

// SetConfigured sets the configured of this subtype
func (m *UserSessionQueryTile) SetConfigured(val bool) {
	m.configuredField = val
}

// Name gets the name of this subtype
func (m *UserSessionQueryTile) Name() *string {
	return m.nameField
}

// SetName sets the name of this subtype
func (m *UserSessionQueryTile) SetName(val *string) {
	m.nameField = val
}

// TileFilter gets the tile filter of this subtype
func (m *UserSessionQueryTile) TileFilter() *TileFilter {
	return m.tileFilterField
}

// SetTileFilter sets the tile filter of this subtype
func (m *UserSessionQueryTile) SetTileFilter(val *TileFilter) {
	m.tileFilterField = val
}

// TileType gets the tile type of this subtype
func (m *UserSessionQueryTile) TileType() string {
	return "UserSessionQueryTile"
}

// SetTileType sets the tile type of this subtype
func (m *UserSessionQueryTile) SetTileType(val string) {

}

// CustomName gets the custom name of this subtype

// Limit gets the limit of this subtype

// Query gets the query of this subtype

// TimeFrameShift gets the time frame shift of this subtype

// Type gets the type of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *UserSessionQueryTile) UnmarshalJSON(raw []byte) error {
	var data struct {

		// The name of the tile, set by user.
		// Required: true
		CustomName *string `json:"customName"`

		// The limit of the results, if not set will use the default value of the system
		Limit int32 `json:"limit,omitempty"`

		// A [user session query](https://www.dynatrace.com/support/help/shortlink/usql-info) executed by the tile.
		// Required: true
		Query *string `json:"query"`

		// The comparison timeframe of the query.
		//
		//  If specified, you additionally get the results of the same query with the specified time shift.
		TimeFrameShift string `json:"timeFrameShift,omitempty"`

		// The visualization of the tile.
		// Required: true
		// Enum: [COLUMN_CHART FUNNEL LINE_CHART PIE_CHART SINGLE_VALUE TABLE]
		Type *string `json:"type"`
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

	var result UserSessionQueryTile

	result.boundsField = base.Bounds

	result.configuredField = base.Configured

	result.nameField = base.Name

	result.tileFilterField = base.TileFilter

	if base.TileType != result.TileType() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid tileType value: %q", base.TileType)
	}

	result.CustomName = data.CustomName

	result.Limit = data.Limit

	result.Query = data.Query

	result.TimeFrameShift = data.TimeFrameShift

	result.Type = data.Type

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m UserSessionQueryTile) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// The name of the tile, set by user.
		// Required: true
		CustomName *string `json:"customName"`

		// The limit of the results, if not set will use the default value of the system
		Limit int32 `json:"limit,omitempty"`

		// A [user session query](https://www.dynatrace.com/support/help/shortlink/usql-info) executed by the tile.
		// Required: true
		Query *string `json:"query"`

		// The comparison timeframe of the query.
		//
		//  If specified, you additionally get the results of the same query with the specified time shift.
		TimeFrameShift string `json:"timeFrameShift,omitempty"`

		// The visualization of the tile.
		// Required: true
		// Enum: [COLUMN_CHART FUNNEL LINE_CHART PIE_CHART SINGLE_VALUE TABLE]
		Type *string `json:"type"`
	}{

		CustomName: m.CustomName,

		Limit: m.Limit,

		Query: m.Query,

		TimeFrameShift: m.TimeFrameShift,

		Type: m.Type,
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

// Validate validates this user session query tile
func (m *UserSessionQueryTile) Validate(formats strfmt.Registry) error {
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

	if err := m.validateCustomName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuery(formats); err != nil {
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

func (m *UserSessionQueryTile) validateBounds(formats strfmt.Registry) error {

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

func (m *UserSessionQueryTile) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name()); err != nil {
		return err
	}

	return nil
}

func (m *UserSessionQueryTile) validateTileFilter(formats strfmt.Registry) error {

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

func (m *UserSessionQueryTile) validateCustomName(formats strfmt.Registry) error {

	if err := validate.Required("customName", "body", m.CustomName); err != nil {
		return err
	}

	return nil
}

func (m *UserSessionQueryTile) validateQuery(formats strfmt.Registry) error {

	if err := validate.Required("query", "body", m.Query); err != nil {
		return err
	}

	return nil
}

var userSessionQueryTileTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["COLUMN_CHART","FUNNEL","LINE_CHART","PIE_CHART","SINGLE_VALUE","TABLE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userSessionQueryTileTypeTypePropEnum = append(userSessionQueryTileTypeTypePropEnum, v)
	}
}

// property enum
func (m *UserSessionQueryTile) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, userSessionQueryTileTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *UserSessionQueryTile) validateType(formats strfmt.Registry) error {

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
func (m *UserSessionQueryTile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserSessionQueryTile) UnmarshalBinary(b []byte) error {
	var res UserSessionQueryTile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
