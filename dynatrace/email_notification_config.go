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

// EmailNotificationConfig email notification config
// swagger:model EmailNotificationConfig
type EmailNotificationConfig struct {
	activeField *bool

	alertingProfileField *strfmt.UUID

	idField strfmt.UUID

	nameField *string

	// The template of the email notification.
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
	// Min Length: 1
	Body *string `json:"body"`

	// The list of the email CC-recipients.
	// Unique: true
	CcReceivers []string `json:"ccReceivers"`

	// The list of the email recipients.
	// Required: true
	// Min Items: 1
	// Unique: true
	Receivers []string `json:"receivers"`

	// Send (`true`) or doesn't send (`false`) an email, confirming problem resolution.
	// Required: true
	ShouldSendForResolvedProblems *bool `json:"shouldSendForResolvedProblems"`

	// The subject of the email notifications.
	// Required: true
	// Max Length: 1000
	// Min Length: 1
	Subject *string `json:"subject"`
}

// Active gets the active of this subtype
func (m *EmailNotificationConfig) Active() *bool {
	return m.activeField
}

// SetActive sets the active of this subtype
func (m *EmailNotificationConfig) SetActive(val *bool) {
	m.activeField = val
}

// AlertingProfile gets the alerting profile of this subtype
func (m *EmailNotificationConfig) AlertingProfile() *strfmt.UUID {
	return m.alertingProfileField
}

// SetAlertingProfile sets the alerting profile of this subtype
func (m *EmailNotificationConfig) SetAlertingProfile(val *strfmt.UUID) {
	m.alertingProfileField = val
}

// ID gets the id of this subtype
func (m *EmailNotificationConfig) ID() strfmt.UUID {
	return m.idField
}

// SetID sets the id of this subtype
func (m *EmailNotificationConfig) SetID(val strfmt.UUID) {
	m.idField = val
}

// Name gets the name of this subtype
func (m *EmailNotificationConfig) Name() *string {
	return m.nameField
}

// SetName sets the name of this subtype
func (m *EmailNotificationConfig) SetName(val *string) {
	m.nameField = val
}

// Type gets the type of this subtype
func (m *EmailNotificationConfig) Type() string {
	return "EmailNotificationConfig"
}

// SetType sets the type of this subtype
func (m *EmailNotificationConfig) SetType(val string) {

}

// Body gets the body of this subtype

// CcReceivers gets the cc receivers of this subtype

// Receivers gets the receivers of this subtype

// ShouldSendForResolvedProblems gets the should send for resolved problems of this subtype

// Subject gets the subject of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *EmailNotificationConfig) UnmarshalJSON(raw []byte) error {
	var data struct {

		// The template of the email notification.
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
		// Min Length: 1
		Body *string `json:"body"`

		// The list of the email CC-recipients.
		// Unique: true
		CcReceivers []string `json:"ccReceivers"`

		// The list of the email recipients.
		// Required: true
		// Min Items: 1
		// Unique: true
		Receivers []string `json:"receivers"`

		// Send (`true`) or doesn't send (`false`) an email, confirming problem resolution.
		// Required: true
		ShouldSendForResolvedProblems *bool `json:"shouldSendForResolvedProblems"`

		// The subject of the email notifications.
		// Required: true
		// Max Length: 1000
		// Min Length: 1
		Subject *string `json:"subject"`
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

	var result EmailNotificationConfig

	result.activeField = base.Active

	result.alertingProfileField = base.AlertingProfile

	result.idField = base.ID

	result.nameField = base.Name

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.Body = data.Body

	result.CcReceivers = data.CcReceivers

	result.Receivers = data.Receivers

	result.ShouldSendForResolvedProblems = data.ShouldSendForResolvedProblems

	result.Subject = data.Subject

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m EmailNotificationConfig) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// The template of the email notification.
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
		// Min Length: 1
		Body *string `json:"body"`

		// The list of the email CC-recipients.
		// Unique: true
		CcReceivers []string `json:"ccReceivers"`

		// The list of the email recipients.
		// Required: true
		// Min Items: 1
		// Unique: true
		Receivers []string `json:"receivers"`

		// Send (`true`) or doesn't send (`false`) an email, confirming problem resolution.
		// Required: true
		ShouldSendForResolvedProblems *bool `json:"shouldSendForResolvedProblems"`

		// The subject of the email notifications.
		// Required: true
		// Max Length: 1000
		// Min Length: 1
		Subject *string `json:"subject"`
	}{

		Body: m.Body,

		CcReceivers: m.CcReceivers,

		Receivers: m.Receivers,

		ShouldSendForResolvedProblems: m.ShouldSendForResolvedProblems,

		Subject: m.Subject,
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

// Validate validates this email notification config
func (m *EmailNotificationConfig) Validate(formats strfmt.Registry) error {
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

	if err := m.validateBody(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCcReceivers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReceivers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShouldSendForResolvedProblems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubject(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmailNotificationConfig) validateActive(formats strfmt.Registry) error {

	if err := validate.Required("active", "body", m.Active()); err != nil {
		return err
	}

	return nil
}

func (m *EmailNotificationConfig) validateAlertingProfile(formats strfmt.Registry) error {

	if err := validate.Required("alertingProfile", "body", m.AlertingProfile()); err != nil {
		return err
	}

	if err := validate.FormatOf("alertingProfile", "body", "uuid", m.AlertingProfile().String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EmailNotificationConfig) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID()) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID().String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EmailNotificationConfig) validateName(formats strfmt.Registry) error {

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

func (m *EmailNotificationConfig) validateBody(formats strfmt.Registry) error {

	if err := validate.Required("body", "body", m.Body); err != nil {
		return err
	}

	if err := validate.MinLength("body", "body", string(*m.Body), 1); err != nil {
		return err
	}

	return nil
}

func (m *EmailNotificationConfig) validateCcReceivers(formats strfmt.Registry) error {

	if swag.IsZero(m.CcReceivers) { // not required
		return nil
	}

	if err := validate.UniqueItems("ccReceivers", "body", m.CcReceivers); err != nil {
		return err
	}

	return nil
}

func (m *EmailNotificationConfig) validateReceivers(formats strfmt.Registry) error {

	if err := validate.Required("receivers", "body", m.Receivers); err != nil {
		return err
	}

	iReceiversSize := int64(len(m.Receivers))

	if err := validate.MinItems("receivers", "body", iReceiversSize, 1); err != nil {
		return err
	}

	if err := validate.UniqueItems("receivers", "body", m.Receivers); err != nil {
		return err
	}

	return nil
}

func (m *EmailNotificationConfig) validateShouldSendForResolvedProblems(formats strfmt.Registry) error {

	if err := validate.Required("shouldSendForResolvedProblems", "body", m.ShouldSendForResolvedProblems); err != nil {
		return err
	}

	return nil
}

func (m *EmailNotificationConfig) validateSubject(formats strfmt.Registry) error {

	if err := validate.Required("subject", "body", m.Subject); err != nil {
		return err
	}

	if err := validate.MinLength("subject", "body", string(*m.Subject), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("subject", "body", string(*m.Subject), 1000); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmailNotificationConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmailNotificationConfig) UnmarshalBinary(b []byte) error {
	var res EmailNotificationConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
