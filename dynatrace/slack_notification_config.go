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

// SlackNotificationConfig slack notification config
// swagger:model SlackNotificationConfig
type SlackNotificationConfig struct {
	activeField *bool

	alertingProfileField *strfmt.UUID

	idField strfmt.UUID

	nameField *string

	// The channel (for example, `#general`) or the user (for example, `@john.smith`) to send the message to.
	// Required: true
	// Max Length: 1000
	// Min Length: 1
	Channel *string `json:"channel"`

	// The content of the message.
	//
	//  You can use the following placeholders:
	// * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.
	// * `{PID}`: The ID of the reported problem.
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
	Title *string `json:"title"`

	// The URL of the Slack WebHook.
	//
	//  This is confidential information, therefore GET requests return this field with the `null` value, and it is optional for PUT requests.
	// Max Length: 1000
	// Min Length: 1
	URL string `json:"url,omitempty"`
}

// Active gets the active of this subtype
func (m *SlackNotificationConfig) Active() *bool {
	return m.activeField
}

// SetActive sets the active of this subtype
func (m *SlackNotificationConfig) SetActive(val *bool) {
	m.activeField = val
}

// AlertingProfile gets the alerting profile of this subtype
func (m *SlackNotificationConfig) AlertingProfile() *strfmt.UUID {
	return m.alertingProfileField
}

// SetAlertingProfile sets the alerting profile of this subtype
func (m *SlackNotificationConfig) SetAlertingProfile(val *strfmt.UUID) {
	m.alertingProfileField = val
}

// ID gets the id of this subtype
func (m *SlackNotificationConfig) ID() strfmt.UUID {
	return m.idField
}

// SetID sets the id of this subtype
func (m *SlackNotificationConfig) SetID(val strfmt.UUID) {
	m.idField = val
}

// Name gets the name of this subtype
func (m *SlackNotificationConfig) Name() *string {
	return m.nameField
}

// SetName sets the name of this subtype
func (m *SlackNotificationConfig) SetName(val *string) {
	m.nameField = val
}

// Type gets the type of this subtype
func (m *SlackNotificationConfig) Type() string {
	return "SlackNotificationConfig"
}

// SetType sets the type of this subtype
func (m *SlackNotificationConfig) SetType(val string) {

}

// Channel gets the channel of this subtype

// Title gets the title of this subtype

// URL gets the url of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *SlackNotificationConfig) UnmarshalJSON(raw []byte) error {
	var data struct {

		// The channel (for example, `#general`) or the user (for example, `@john.smith`) to send the message to.
		// Required: true
		// Max Length: 1000
		// Min Length: 1
		Channel *string `json:"channel"`

		// The content of the message.
		//
		//  You can use the following placeholders:
		// * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.
		// * `{PID}`: The ID of the reported problem.
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
		Title *string `json:"title"`

		// The URL of the Slack WebHook.
		//
		//  This is confidential information, therefore GET requests return this field with the `null` value, and it is optional for PUT requests.
		// Max Length: 1000
		// Min Length: 1
		URL string `json:"url,omitempty"`
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

	var result SlackNotificationConfig

	result.activeField = base.Active

	result.alertingProfileField = base.AlertingProfile

	result.idField = base.ID

	result.nameField = base.Name

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.Channel = data.Channel

	result.Title = data.Title

	result.URL = data.URL

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m SlackNotificationConfig) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// The channel (for example, `#general`) or the user (for example, `@john.smith`) to send the message to.
		// Required: true
		// Max Length: 1000
		// Min Length: 1
		Channel *string `json:"channel"`

		// The content of the message.
		//
		//  You can use the following placeholders:
		// * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.
		// * `{PID}`: The ID of the reported problem.
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
		Title *string `json:"title"`

		// The URL of the Slack WebHook.
		//
		//  This is confidential information, therefore GET requests return this field with the `null` value, and it is optional for PUT requests.
		// Max Length: 1000
		// Min Length: 1
		URL string `json:"url,omitempty"`
	}{

		Channel: m.Channel,

		Title: m.Title,

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

// Validate validates this slack notification config
func (m *SlackNotificationConfig) Validate(formats strfmt.Registry) error {
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

	if err := m.validateChannel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
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

func (m *SlackNotificationConfig) validateActive(formats strfmt.Registry) error {

	if err := validate.Required("active", "body", m.Active()); err != nil {
		return err
	}

	return nil
}

func (m *SlackNotificationConfig) validateAlertingProfile(formats strfmt.Registry) error {

	if err := validate.Required("alertingProfile", "body", m.AlertingProfile()); err != nil {
		return err
	}

	if err := validate.FormatOf("alertingProfile", "body", "uuid", m.AlertingProfile().String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SlackNotificationConfig) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID()) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID().String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SlackNotificationConfig) validateName(formats strfmt.Registry) error {

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

func (m *SlackNotificationConfig) validateChannel(formats strfmt.Registry) error {

	if err := validate.Required("channel", "body", m.Channel); err != nil {
		return err
	}

	if err := validate.MinLength("channel", "body", string(*m.Channel), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("channel", "body", string(*m.Channel), 1000); err != nil {
		return err
	}

	return nil
}

func (m *SlackNotificationConfig) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	if err := validate.MinLength("title", "body", string(*m.Title), 1); err != nil {
		return err
	}

	return nil
}

func (m *SlackNotificationConfig) validateURL(formats strfmt.Registry) error {

	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.MinLength("url", "body", string(m.URL), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("url", "body", string(m.URL), 1000); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SlackNotificationConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SlackNotificationConfig) UnmarshalBinary(b []byte) error {
	var res SlackNotificationConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}