// Code generated by go-swagger; DO NOT EDIT.

package dynatrace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// XMattersNotificationConfig x matters notification config
// swagger:model XMattersNotificationConfig
type XMattersNotificationConfig struct {
	activeField *bool

	alertingProfileField *strfmt.UUID

	idField strfmt.UUID

	nameField *string

	// Accept any, including self-signed and invalid, SSL certificate (`true`) or only trusted (`false`) certificates.
	// Required: true
	AcceptAnyCertificate *bool `json:"acceptAnyCertificate"`

	// A list of the additional HTTP headers.
	Headers []*HTTPHeader `json:"headers"`

	// The content of the message.
	//
	//  You can use the following placeholders:
	// * `{ImpactedEntities}`: Details about the entities impacted by the problem in form of a JSON array.
	// * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.
	// * `{PID}`: The ID of the reported problem.
	// * `{ProblemDetailsHTML}`: All problem event details, including root cause, as an HTML-formatted string.
	// * `{ProblemDetailsJSON}`: All problem event details, including root cause, as a JSON object.
	// * `{ProblemDetailsMarkdown}`: All problem event details, including root cause, as a [Markdown-formatted](https://www.markdownguide.org/cheat-sheet/) string.
	// * `{ProblemDetailsText}`: All problem event details, including root cause, as a text-formatted string.
	// * `{ProblemID}`: The display number of the reported problem.
	// * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.
	// * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.
	// * `{ProblemTitle}`: A short description of the problem.
	// * `{ProblemURL}`: The URL of the problem within Dynatrace.
	// * `{State}`: The state of the problem. Possible values are `OPEN` and `RESOLVED`. If the problem has been merged into another problem, it has the `MERGED` value.
	// * `{Tags}`: The list of tags that are defined for all impacted entities, separated by commas.
	//
	// Required: true
	// Max Length: 600
	// Min Length: 1
	Payload *string `json:"payload"`

	// The URL of the xMatters WebHook.
	// Required: true
	// Max Length: 1000
	// Min Length: 1
	URL *string `json:"url"`
}

// Active gets the active of this subtype
func (m *XMattersNotificationConfig) Active() *bool {
	return m.activeField
}

// SetActive sets the active of this subtype
func (m *XMattersNotificationConfig) SetActive(val *bool) {
	m.activeField = val
}

// AlertingProfile gets the alerting profile of this subtype
func (m *XMattersNotificationConfig) AlertingProfile() *strfmt.UUID {
	return m.alertingProfileField
}

// SetAlertingProfile sets the alerting profile of this subtype
func (m *XMattersNotificationConfig) SetAlertingProfile(val *strfmt.UUID) {
	m.alertingProfileField = val
}

// ID gets the id of this subtype
func (m *XMattersNotificationConfig) ID() strfmt.UUID {
	return m.idField
}

// SetID sets the id of this subtype
func (m *XMattersNotificationConfig) SetID(val strfmt.UUID) {
	m.idField = val
}

// Name gets the name of this subtype
func (m *XMattersNotificationConfig) Name() *string {
	return m.nameField
}

// SetName sets the name of this subtype
func (m *XMattersNotificationConfig) SetName(val *string) {
	m.nameField = val
}

// Type gets the type of this subtype
func (m *XMattersNotificationConfig) Type() string {
	return "XMattersNotificationConfig"
}

// SetType sets the type of this subtype
func (m *XMattersNotificationConfig) SetType(val string) {

}

// AcceptAnyCertificate gets the accept any certificate of this subtype

// Headers gets the headers of this subtype

// Payload gets the payload of this subtype

// URL gets the url of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *XMattersNotificationConfig) UnmarshalJSON(raw []byte) error {
	var data struct {

		// Accept any, including self-signed and invalid, SSL certificate (`true`) or only trusted (`false`) certificates.
		// Required: true
		AcceptAnyCertificate *bool `json:"acceptAnyCertificate"`

		// A list of the additional HTTP headers.
		Headers []*HTTPHeader `json:"headers"`

		// The content of the message.
		//
		//  You can use the following placeholders:
		// * `{ImpactedEntities}`: Details about the entities impacted by the problem in form of a JSON array.
		// * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.
		// * `{PID}`: The ID of the reported problem.
		// * `{ProblemDetailsHTML}`: All problem event details, including root cause, as an HTML-formatted string.
		// * `{ProblemDetailsJSON}`: All problem event details, including root cause, as a JSON object.
		// * `{ProblemDetailsMarkdown}`: All problem event details, including root cause, as a [Markdown-formatted](https://www.markdownguide.org/cheat-sheet/) string.
		// * `{ProblemDetailsText}`: All problem event details, including root cause, as a text-formatted string.
		// * `{ProblemID}`: The display number of the reported problem.
		// * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.
		// * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.
		// * `{ProblemTitle}`: A short description of the problem.
		// * `{ProblemURL}`: The URL of the problem within Dynatrace.
		// * `{State}`: The state of the problem. Possible values are `OPEN` and `RESOLVED`. If the problem has been merged into another problem, it has the `MERGED` value.
		// * `{Tags}`: The list of tags that are defined for all impacted entities, separated by commas.
		//
		// Required: true
		// Max Length: 600
		// Min Length: 1
		Payload *string `json:"payload"`

		// The URL of the xMatters WebHook.
		// Required: true
		// Max Length: 1000
		// Min Length: 1
		URL *string `json:"url"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Active *bool `json:"active"`

		AlertingProfile *strfmt.UUID `json:"alertingProfile"`

		ID strfmt.UUID `json:"id,omitempty"`

		Name *string `json:"name"`

		Type string `json:"type"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result XMattersNotificationConfig

	result.activeField = base.Active

	result.alertingProfileField = base.AlertingProfile

	result.idField = base.ID

	result.nameField = base.Name

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.AcceptAnyCertificate = data.AcceptAnyCertificate

	result.Headers = data.Headers

	result.Payload = data.Payload

	result.URL = data.URL

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m XMattersNotificationConfig) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// Accept any, including self-signed and invalid, SSL certificate (`true`) or only trusted (`false`) certificates.
		// Required: true
		AcceptAnyCertificate *bool `json:"acceptAnyCertificate"`

		// A list of the additional HTTP headers.
		Headers []*HTTPHeader `json:"headers"`

		// The content of the message.
		//
		//  You can use the following placeholders:
		// * `{ImpactedEntities}`: Details about the entities impacted by the problem in form of a JSON array.
		// * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.
		// * `{PID}`: The ID of the reported problem.
		// * `{ProblemDetailsHTML}`: All problem event details, including root cause, as an HTML-formatted string.
		// * `{ProblemDetailsJSON}`: All problem event details, including root cause, as a JSON object.
		// * `{ProblemDetailsMarkdown}`: All problem event details, including root cause, as a [Markdown-formatted](https://www.markdownguide.org/cheat-sheet/) string.
		// * `{ProblemDetailsText}`: All problem event details, including root cause, as a text-formatted string.
		// * `{ProblemID}`: The display number of the reported problem.
		// * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.
		// * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.
		// * `{ProblemTitle}`: A short description of the problem.
		// * `{ProblemURL}`: The URL of the problem within Dynatrace.
		// * `{State}`: The state of the problem. Possible values are `OPEN` and `RESOLVED`. If the problem has been merged into another problem, it has the `MERGED` value.
		// * `{Tags}`: The list of tags that are defined for all impacted entities, separated by commas.
		//
		// Required: true
		// Max Length: 600
		// Min Length: 1
		Payload *string `json:"payload"`

		// The URL of the xMatters WebHook.
		// Required: true
		// Max Length: 1000
		// Min Length: 1
		URL *string `json:"url"`
	}{

		AcceptAnyCertificate: m.AcceptAnyCertificate,

		Headers: m.Headers,

		Payload: m.Payload,

		URL: m.URL,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Active *bool `json:"active"`

		AlertingProfile *strfmt.UUID `json:"alertingProfile"`

		ID strfmt.UUID `json:"id,omitempty"`

		Name *string `json:"name"`

		Type string `json:"type"`
	}{

		Active: m.Active(),

		AlertingProfile: m.AlertingProfile(),

		ID: m.ID(),

		Name: m.Name(),

		Type: m.Type(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this x matters notification config
func (m *XMattersNotificationConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAlertingProfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAcceptAnyCertificate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHeaders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePayload(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *XMattersNotificationConfig) validateActive(formats strfmt.Registry) error {

	if err := validate.Required("active", "body", m.Active()); err != nil {
		return err
	}

	return nil
}

func (m *XMattersNotificationConfig) validateAlertingProfile(formats strfmt.Registry) error {

	if err := validate.Required("alertingProfile", "body", m.AlertingProfile()); err != nil {
		return err
	}

	if err := validate.FormatOf("alertingProfile", "body", "uuid", m.AlertingProfile().String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *XMattersNotificationConfig) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID()) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID().String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *XMattersNotificationConfig) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name()); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name()), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name()), 100); err != nil {
		return err
	}

	return nil
}

func (m *XMattersNotificationConfig) validateAcceptAnyCertificate(formats strfmt.Registry) error {

	if err := validate.Required("acceptAnyCertificate", "body", m.AcceptAnyCertificate); err != nil {
		return err
	}

	return nil
}

func (m *XMattersNotificationConfig) validateHeaders(formats strfmt.Registry) error {

	if swag.IsZero(m.Headers) { // not required
		return nil
	}

	for i := 0; i < len(m.Headers); i++ {
		if swag.IsZero(m.Headers[i]) { // not required
			continue
		}

		if m.Headers[i] != nil {
			if err := m.Headers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("headers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *XMattersNotificationConfig) validatePayload(formats strfmt.Registry) error {

	if err := validate.Required("payload", "body", m.Payload); err != nil {
		return err
	}

	if err := validate.MinLength("payload", "body", string(*m.Payload), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("payload", "body", string(*m.Payload), 600); err != nil {
		return err
	}

	return nil
}

func (m *XMattersNotificationConfig) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	if err := validate.MinLength("url", "body", string(*m.URL), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("url", "body", string(*m.URL), 1000); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *XMattersNotificationConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *XMattersNotificationConfig) UnmarshalBinary(b []byte) error {
	var res XMattersNotificationConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
