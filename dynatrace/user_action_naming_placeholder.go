// Code generated by go-swagger; DO NOT EDIT.

package dynatrace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UserActionNamingPlaceholder The placeholder settings.
// swagger:model UserActionNamingPlaceholder
type UserActionNamingPlaceholder struct {

	// Input.
	// Required: true
	// Enum: [ELEMENT_IDENTIFIER INPUT_TYPE METADATA PAGE_TITLE PAGE_URL SOURCE_URL TOP_XHR_URL XHR_URL]
	Input *string `json:"input"`

	// Id of the metadata.
	MetadataID int32 `json:"metadataId,omitempty"`

	// Placeholder name.
	// Required: true
	// Max Length: 50
	// Min Length: 1
	Name *string `json:"name"`

	// Part.
	// Required: true
	// Enum: [ALL ANCHOR PATH]
	ProcessingPart *string `json:"processingPart"`

	// Processing actions.
	ProcessingSteps []*UserActionNamingPlaceholderProcessingStep `json:"processingSteps"`

	// Use the element identifier that was selected by Dynatrace.
	// Required: true
	UseGuessedElementIdentifier *bool `json:"useGuessedElementIdentifier"`
}

// Validate validates this user action naming placeholder
func (m *UserActionNamingPlaceholder) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInput(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcessingPart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcessingSteps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUseGuessedElementIdentifier(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var userActionNamingPlaceholderTypeInputPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ELEMENT_IDENTIFIER","INPUT_TYPE","METADATA","PAGE_TITLE","PAGE_URL","SOURCE_URL","TOP_XHR_URL","XHR_URL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userActionNamingPlaceholderTypeInputPropEnum = append(userActionNamingPlaceholderTypeInputPropEnum, v)
	}
}

const (

	// UserActionNamingPlaceholderInputELEMENTIDENTIFIER captures enum value "ELEMENT_IDENTIFIER"
	UserActionNamingPlaceholderInputELEMENTIDENTIFIER string = "ELEMENT_IDENTIFIER"

	// UserActionNamingPlaceholderInputINPUTTYPE captures enum value "INPUT_TYPE"
	UserActionNamingPlaceholderInputINPUTTYPE string = "INPUT_TYPE"

	// UserActionNamingPlaceholderInputMETADATA captures enum value "METADATA"
	UserActionNamingPlaceholderInputMETADATA string = "METADATA"

	// UserActionNamingPlaceholderInputPAGETITLE captures enum value "PAGE_TITLE"
	UserActionNamingPlaceholderInputPAGETITLE string = "PAGE_TITLE"

	// UserActionNamingPlaceholderInputPAGEURL captures enum value "PAGE_URL"
	UserActionNamingPlaceholderInputPAGEURL string = "PAGE_URL"

	// UserActionNamingPlaceholderInputSOURCEURL captures enum value "SOURCE_URL"
	UserActionNamingPlaceholderInputSOURCEURL string = "SOURCE_URL"

	// UserActionNamingPlaceholderInputTOPXHRURL captures enum value "TOP_XHR_URL"
	UserActionNamingPlaceholderInputTOPXHRURL string = "TOP_XHR_URL"

	// UserActionNamingPlaceholderInputXHRURL captures enum value "XHR_URL"
	UserActionNamingPlaceholderInputXHRURL string = "XHR_URL"
)

// prop value enum
func (m *UserActionNamingPlaceholder) validateInputEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, userActionNamingPlaceholderTypeInputPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *UserActionNamingPlaceholder) validateInput(formats strfmt.Registry) error {

	if err := validate.Required("input", "body", m.Input); err != nil {
		return err
	}

	// value enum
	if err := m.validateInputEnum("input", "body", *m.Input); err != nil {
		return err
	}

	return nil
}

func (m *UserActionNamingPlaceholder) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 50); err != nil {
		return err
	}

	return nil
}

var userActionNamingPlaceholderTypeProcessingPartPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ALL","ANCHOR","PATH"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userActionNamingPlaceholderTypeProcessingPartPropEnum = append(userActionNamingPlaceholderTypeProcessingPartPropEnum, v)
	}
}

const (

	// UserActionNamingPlaceholderProcessingPartALL captures enum value "ALL"
	UserActionNamingPlaceholderProcessingPartALL string = "ALL"

	// UserActionNamingPlaceholderProcessingPartANCHOR captures enum value "ANCHOR"
	UserActionNamingPlaceholderProcessingPartANCHOR string = "ANCHOR"

	// UserActionNamingPlaceholderProcessingPartPATH captures enum value "PATH"
	UserActionNamingPlaceholderProcessingPartPATH string = "PATH"
)

// prop value enum
func (m *UserActionNamingPlaceholder) validateProcessingPartEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, userActionNamingPlaceholderTypeProcessingPartPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *UserActionNamingPlaceholder) validateProcessingPart(formats strfmt.Registry) error {

	if err := validate.Required("processingPart", "body", m.ProcessingPart); err != nil {
		return err
	}

	// value enum
	if err := m.validateProcessingPartEnum("processingPart", "body", *m.ProcessingPart); err != nil {
		return err
	}

	return nil
}

func (m *UserActionNamingPlaceholder) validateProcessingSteps(formats strfmt.Registry) error {

	if swag.IsZero(m.ProcessingSteps) { // not required
		return nil
	}

	for i := 0; i < len(m.ProcessingSteps); i++ {
		if swag.IsZero(m.ProcessingSteps[i]) { // not required
			continue
		}

		if m.ProcessingSteps[i] != nil {
			if err := m.ProcessingSteps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("processingSteps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UserActionNamingPlaceholder) validateUseGuessedElementIdentifier(formats strfmt.Registry) error {

	if err := validate.Required("useGuessedElementIdentifier", "body", m.UseGuessedElementIdentifier); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserActionNamingPlaceholder) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserActionNamingPlaceholder) UnmarshalBinary(b []byte) error {
	var res UserActionNamingPlaceholder
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
