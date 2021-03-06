// Code generated by go-swagger; DO NOT EDIT.

package dynatrace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// Credentials A set of credentials for synthetic monitors.
// swagger:discriminator Credentials type
type Credentials interface {
	runtime.Validatable

	// A short description of the credentials set..
	// Required: true
	Description() *string
	SetDescription(*string)

	// The ID of the credentials set.
	ID() string
	SetID(string)

	// The name of the credentials set.
	// Required: true
	Name() *string
	SetName(*string)

	// The credentials set is available to every user (`false`) or to owner only (`true`).
	OwnerAccessOnly() bool
	SetOwnerAccessOnly(bool)

	// Defines the actual set of fields depending on the value:
	//
	// CERTIFICATE -> CertificateCredentials
	// USERNAME_PASSWORD -> UserPasswordCredentials
	// TOKEN -> TokenCredentials
	//
	// Required: true
	// Enum: [CERTIFICATE USERNAME_PASSWORD TOKEN]
	Type() string
	SetType(string)
}

type credentials struct {
	descriptionField *string

	idField string

	nameField *string

	ownerAccessOnlyField bool

	typeField string
}

// Description gets the description of this polymorphic type
func (m *credentials) Description() *string {
	return m.descriptionField
}

// SetDescription sets the description of this polymorphic type
func (m *credentials) SetDescription(val *string) {
	m.descriptionField = val
}

// ID gets the id of this polymorphic type
func (m *credentials) ID() string {
	return m.idField
}

// SetID sets the id of this polymorphic type
func (m *credentials) SetID(val string) {
	m.idField = val
}

// Name gets the name of this polymorphic type
func (m *credentials) Name() *string {
	return m.nameField
}

// SetName sets the name of this polymorphic type
func (m *credentials) SetName(val *string) {
	m.nameField = val
}

// OwnerAccessOnly gets the owner access only of this polymorphic type
func (m *credentials) OwnerAccessOnly() bool {
	return m.ownerAccessOnlyField
}

// SetOwnerAccessOnly sets the owner access only of this polymorphic type
func (m *credentials) SetOwnerAccessOnly(val bool) {
	m.ownerAccessOnlyField = val
}

// Type gets the type of this polymorphic type
func (m *credentials) Type() string {
	return "Credentials"
}

// SetType sets the type of this polymorphic type
func (m *credentials) SetType(val string) {

}

// UnmarshalCredentialsSlice unmarshals polymorphic slices of Credentials
func UnmarshalCredentialsSlice(reader io.Reader, consumer runtime.Consumer) ([]Credentials, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []Credentials
	for _, element := range elements {
		obj, err := unmarshalCredentials(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalCredentials unmarshals polymorphic Credentials
func UnmarshalCredentials(reader io.Reader, consumer runtime.Consumer) (Credentials, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalCredentials(data, consumer)
}

func unmarshalCredentials(data []byte, consumer runtime.Consumer) (Credentials, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the type property.
	var getType struct {
		Type string `json:"type"`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("type", "body", getType.Type); err != nil {
		return nil, err
	}

	// The value of type is used to determine which type to create and unmarshal the data into
	switch getType.Type {
	case "CertificateCredentials":
		var result CertificateCredentials
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "Credentials":
		var result credentials
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "TokenCredentials":
		var result TokenCredentials
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "UserPasswordCredentials":
		var result UserPasswordCredentials
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	}
	return nil, errors.New(422, "invalid type value: %q", getType.Type)

}

// Validate validates this credentials
func (m *credentials) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *credentials) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description()); err != nil {
		return err
	}

	return nil
}

func (m *credentials) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name()); err != nil {
		return err
	}

	return nil
}

var credentialsTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CERTIFICATE","USERNAME_PASSWORD","TOKEN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		credentialsTypeTypePropEnum = append(credentialsTypeTypePropEnum, v)
	}
}

const (

	// CredentialsTypeCERTIFICATE captures enum value "CERTIFICATE"
	CredentialsTypeCERTIFICATE string = "CERTIFICATE"

	// CredentialsTypeUSERNAMEPASSWORD captures enum value "USERNAME_PASSWORD"
	CredentialsTypeUSERNAMEPASSWORD string = "USERNAME_PASSWORD"

	// CredentialsTypeTOKEN captures enum value "TOKEN"
	CredentialsTypeTOKEN string = "TOKEN"
)

// prop value enum
func (m *credentials) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, credentialsTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}
