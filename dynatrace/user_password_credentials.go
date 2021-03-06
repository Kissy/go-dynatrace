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

// UserPasswordCredentials user password credentials
// swagger:model UserPasswordCredentials
type UserPasswordCredentials struct {
	descriptionField *string

	idField string

	nameField *string

	ownerAccessOnlyField bool

	// The password of the credential.
	// Required: true
	Password *string `json:"password"`

	// The username of the credentials set.
	// Required: true
	User *string `json:"user"`
}

// Description gets the description of this subtype
func (m *UserPasswordCredentials) Description() *string {
	return m.descriptionField
}

// SetDescription sets the description of this subtype
func (m *UserPasswordCredentials) SetDescription(val *string) {
	m.descriptionField = val
}

// ID gets the id of this subtype
func (m *UserPasswordCredentials) ID() string {
	return m.idField
}

// SetID sets the id of this subtype
func (m *UserPasswordCredentials) SetID(val string) {
	m.idField = val
}

// Name gets the name of this subtype
func (m *UserPasswordCredentials) Name() *string {
	return m.nameField
}

// SetName sets the name of this subtype
func (m *UserPasswordCredentials) SetName(val *string) {
	m.nameField = val
}

// OwnerAccessOnly gets the owner access only of this subtype
func (m *UserPasswordCredentials) OwnerAccessOnly() bool {
	return m.ownerAccessOnlyField
}

// SetOwnerAccessOnly sets the owner access only of this subtype
func (m *UserPasswordCredentials) SetOwnerAccessOnly(val bool) {
	m.ownerAccessOnlyField = val
}

// Type gets the type of this subtype
func (m *UserPasswordCredentials) Type() string {
	return "UserPasswordCredentials"
}

// SetType sets the type of this subtype
func (m *UserPasswordCredentials) SetType(val string) {

}

// Password gets the password of this subtype

// User gets the user of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *UserPasswordCredentials) UnmarshalJSON(raw []byte) error {
	var data struct {

		// The password of the credential.
		// Required: true
		Password *string `json:"password"`

		// The username of the credentials set.
		// Required: true
		User *string `json:"user"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Description *string `json:"description"`

		ID string `json:"id,omitempty"`

		Name *string `json:"name"`

		OwnerAccessOnly bool `json:"ownerAccessOnly,omitempty"`

		Type string `json:"type"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result UserPasswordCredentials

	result.descriptionField = base.Description

	result.idField = base.ID

	result.nameField = base.Name

	result.ownerAccessOnlyField = base.OwnerAccessOnly

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.Password = data.Password

	result.User = data.User

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m UserPasswordCredentials) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// The password of the credential.
		// Required: true
		Password *string `json:"password"`

		// The username of the credentials set.
		// Required: true
		User *string `json:"user"`
	}{

		Password: m.Password,

		User: m.User,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Description *string `json:"description"`

		ID string `json:"id,omitempty"`

		Name *string `json:"name"`

		OwnerAccessOnly bool `json:"ownerAccessOnly,omitempty"`

		Type string `json:"type"`
	}{

		Description: m.Description(),

		ID: m.ID(),

		Name: m.Name(),

		OwnerAccessOnly: m.OwnerAccessOnly(),

		Type: m.Type(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this user password credentials
func (m *UserPasswordCredentials) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserPasswordCredentials) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description()); err != nil {
		return err
	}

	return nil
}

func (m *UserPasswordCredentials) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name()); err != nil {
		return err
	}

	return nil
}

func (m *UserPasswordCredentials) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

func (m *UserPasswordCredentials) validateUser(formats strfmt.Registry) error {

	if err := validate.Required("user", "body", m.User); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserPasswordCredentials) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserPasswordCredentials) UnmarshalBinary(b []byte) error {
	var res UserPasswordCredentials
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
