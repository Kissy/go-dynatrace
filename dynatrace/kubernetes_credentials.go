// Code generated by go-swagger; DO NOT EDIT.

package dynatrace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// KubernetesCredentials Configuration for specific Kubernetes credentials.
// swagger:model KubernetesCredentials
type KubernetesCredentials struct {

	// The monitoring is enabled (`true`) or disabled (`false`) for given credentials configuration.
	//
	// If not set on creation, the `true` value is used.
	//
	// If the field is omitted during an update, the old value is used.
	Active bool `json:"active,omitempty"`

	// The service account bearer token for the Kubernetes API server.
	//
	// Submit your token on creation or update of the configuration. For security reasons, GET requests return this field as `null`.
	//
	// If the field is omitted during an update, the old value is used.
	AuthToken string `json:"authToken,omitempty"`

	// The status of the configured endpoint.
	//
	// ASSIGNED: The credentials are assigned to an ActiveGate which is responsible for processing.
	// UNASSIGNED: The credentials are not yet assigned to an ActiveGate so there is currently no processing.
	// DISABLED: The credentials have been disabled by the user.
	// FASTCHECK_AUTH_ERROR: The credentials are invalid.
	// FASTCHECK_TLS_ERROR: The endpoint TLS certificate is invalid.
	// FASTCHECK_NO_RESPONSE: The endpoint did not return a result until the timeout was reached.
	// FASTCHECK_INVALID_ENDPOINT: The endpoint URL was invalid.
	// FASTCHECK_AUTH_LOCKED: The credentials seem to be locked.
	// UNKNOWN: An unknown error occured.
	//
	// Read Only: true
	// Enum: [ASSIGNED DISABLED FASTCHECK_AUTH_ERROR FASTCHECK_AUTH_LOCKED FASTCHECK_INVALID_ENDPOINT FASTCHECK_NO_RESPONSE FASTCHECK_TLS_ERROR UNASSIGNED UNKNOWN]
	EndpointStatus string `json:"endpointStatus,omitempty"`

	// The detailed status info of the configured endpoint.
	// Read Only: true
	EndpointStatusInfo string `json:"endpointStatusInfo,omitempty"`

	// The URL of the Kubernetes API server.
	//
	// It must be unique within a Dynatrace environment.
	//
	// The URL must valid according to RFC 2396. Leading or trailing whitespaces are not allowed.
	EndpointURL string `json:"endpointUrl,omitempty"`

	// The ID of the given credentials configuration.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The name of the Kubernetes credentials configuration.
	//
	// Allowed characters are letters, numbers, whitespaces, and the following characters: `.+-_`. Leading or trailing whitespace is not allowed.
	// Required: true
	// Min Length: 1
	// Pattern: ^([a-zA-Z0-9_ +-.]*)$
	Label *string `json:"label"`

	// Metadata of the given credentials configuration which is useful for debugging.
	// Read Only: true
	Metadata *ConfigurationMetadata `json:"metadata,omitempty"`
}

// Validate validates this kubernetes credentials
func (m *KubernetesCredentials) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndpointStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var kubernetesCredentialsTypeEndpointStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ASSIGNED","DISABLED","FASTCHECK_AUTH_ERROR","FASTCHECK_AUTH_LOCKED","FASTCHECK_INVALID_ENDPOINT","FASTCHECK_NO_RESPONSE","FASTCHECK_TLS_ERROR","UNASSIGNED","UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		kubernetesCredentialsTypeEndpointStatusPropEnum = append(kubernetesCredentialsTypeEndpointStatusPropEnum, v)
	}
}

const (

	// KubernetesCredentialsEndpointStatusASSIGNED captures enum value "ASSIGNED"
	KubernetesCredentialsEndpointStatusASSIGNED string = "ASSIGNED"

	// KubernetesCredentialsEndpointStatusDISABLED captures enum value "DISABLED"
	KubernetesCredentialsEndpointStatusDISABLED string = "DISABLED"

	// KubernetesCredentialsEndpointStatusFASTCHECKAUTHERROR captures enum value "FASTCHECK_AUTH_ERROR"
	KubernetesCredentialsEndpointStatusFASTCHECKAUTHERROR string = "FASTCHECK_AUTH_ERROR"

	// KubernetesCredentialsEndpointStatusFASTCHECKAUTHLOCKED captures enum value "FASTCHECK_AUTH_LOCKED"
	KubernetesCredentialsEndpointStatusFASTCHECKAUTHLOCKED string = "FASTCHECK_AUTH_LOCKED"

	// KubernetesCredentialsEndpointStatusFASTCHECKINVALIDENDPOINT captures enum value "FASTCHECK_INVALID_ENDPOINT"
	KubernetesCredentialsEndpointStatusFASTCHECKINVALIDENDPOINT string = "FASTCHECK_INVALID_ENDPOINT"

	// KubernetesCredentialsEndpointStatusFASTCHECKNORESPONSE captures enum value "FASTCHECK_NO_RESPONSE"
	KubernetesCredentialsEndpointStatusFASTCHECKNORESPONSE string = "FASTCHECK_NO_RESPONSE"

	// KubernetesCredentialsEndpointStatusFASTCHECKTLSERROR captures enum value "FASTCHECK_TLS_ERROR"
	KubernetesCredentialsEndpointStatusFASTCHECKTLSERROR string = "FASTCHECK_TLS_ERROR"

	// KubernetesCredentialsEndpointStatusUNASSIGNED captures enum value "UNASSIGNED"
	KubernetesCredentialsEndpointStatusUNASSIGNED string = "UNASSIGNED"

	// KubernetesCredentialsEndpointStatusUNKNOWN captures enum value "UNKNOWN"
	KubernetesCredentialsEndpointStatusUNKNOWN string = "UNKNOWN"
)

// prop value enum
func (m *KubernetesCredentials) validateEndpointStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, kubernetesCredentialsTypeEndpointStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *KubernetesCredentials) validateEndpointStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.EndpointStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateEndpointStatusEnum("endpointStatus", "body", m.EndpointStatus); err != nil {
		return err
	}

	return nil
}

func (m *KubernetesCredentials) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("label", "body", m.Label); err != nil {
		return err
	}

	if err := validate.MinLength("label", "body", string(*m.Label), 1); err != nil {
		return err
	}

	if err := validate.Pattern("label", "body", string(*m.Label), `^([a-zA-Z0-9_ +-.]*)$`); err != nil {
		return err
	}

	return nil
}

func (m *KubernetesCredentials) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KubernetesCredentials) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubernetesCredentials) UnmarshalBinary(b []byte) error {
	var res KubernetesCredentials
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
