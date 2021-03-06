// Code generated by go-swagger; DO NOT EDIT.

package dynatrace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MetricEvent The configuration of the metric event.
// swagger:model MetricEvent
type MetricEvent struct {

	// How the metric data points are aggregated for the evaluation.
	//
	//  The timeseries must support this aggregation.
	// Enum: [AVG COUNT MAX MEDIAN MIN OF_INTEREST OF_INTEREST_RATIO OTHER OTHER_RATIO P90 SUM VALUE]
	AggregationType string `json:"aggregationType,omitempty"`

	// The condition for the **threshold** value check: above or below.
	// Required: true
	// Enum: [ABOVE BELOW]
	AlertCondition *string `json:"alertCondition"`

	alertingScopeField []MetricEventAlertingScope

	// The number of one-minute samples within the evaluation window that must go back to normal to close the event.
	// Required: true
	// Minimum: 3
	DealertingSamples *int32 `json:"dealertingSamples"`

	// The description of the metric event.
	// Required: true
	// Max Length: 4000
	// Min Length: 0
	Description *string `json:"description"`

	// The metric event is enabled (`true`) or disabled (`false`).
	// Required: true
	Enabled *bool `json:"enabled"`

	// The ID of the metric event.
	// Max Length: 256
	// Min Length: 0
	// Pattern: ^[a-zA-Z0-9\.\-_]+$
	ID *string `json:"id,omitempty"`

	// Metadata useful for debugging.
	// Read Only: true
	Metadata *ConfigurationMetadata `json:"metadata,omitempty"`

	metricDimensionsField []MetricEventDimensions

	// The ID of the metric evaluated by the metric event.
	// Required: true
	MetricID *string `json:"metricId"`

	// The name of the metric event displayed in the UI.
	// Required: true
	// Max Length: 1024
	// Min Length: 0
	Name *string `json:"name"`

	// The number of one-minute samples that form sliding evaluation window.
	// Required: true
	// Maximum: 20
	// Minimum: 3
	Samples *int32 `json:"samples"`

	// The type of the event to trigger on the threshold violation.
	//
	// The `CUSTOM_ALERT` type is not correlated with other alerts.
	// The `INFO` type does not open a problem.
	// Enum: [AVAILABILITY CUSTOM_ALERT ERROR INFO PERFORMANCE RESOURCE_CONTENTION]
	Severity string `json:"severity,omitempty"`

	// The value of the threshold.
	// Required: true
	Threshold *float64 `json:"threshold"`

	// The unit of the threshold, matching the metric definition.
	//
	// If not set the metrics unit is picked.
	// Enum: [BIT BIT_PER_HOUR BIT_PER_MINUTE BIT_PER_SECOND BYTE BYTE_PER_HOUR BYTE_PER_MINUTE BYTE_PER_SECOND CORES COUNT GIBI_BYTE GIGA_BYTE HOUR KIBI_BYTE KIBI_BYTE_PER_HOUR KIBI_BYTE_PER_MINUTE KIBI_BYTE_PER_SECOND KILO_BYTE KILO_BYTE_PER_HOUR KILO_BYTE_PER_MINUTE KILO_BYTE_PER_SECOND MEBI_BYTE MEBI_BYTE_PER_HOUR MEBI_BYTE_PER_MINUTE MEBI_BYTE_PER_SECOND MEGA_BYTE MEGA_BYTE_PER_HOUR MEGA_BYTE_PER_MINUTE MEGA_BYTE_PER_SECOND MICRO_SECOND MILLI_CORES MILLI_SECOND MILLI_SECOND_PER_MINUTE MINUTE NANO_SECOND NANO_SECOND_PER_MINUTE NOT_APPLICABLE PERCENT PER_HOUR PER_MINUTE PER_SECOND PROMILLE RATIO SECOND STATE UNSPECIFIED]
	Unit string `json:"unit,omitempty"`

	// The number of one-minute samples within the evaluation window that must violate the threshold to trigger an event
	// Required: true
	// Minimum: 1
	ViolatingSamples *int32 `json:"violatingSamples"`
}

// AlertingScope gets the alerting scope of this base type
func (m *MetricEvent) AlertingScope() []MetricEventAlertingScope {
	return m.alertingScopeField
}

// SetAlertingScope sets the alerting scope of this base type
func (m *MetricEvent) SetAlertingScope(val []MetricEventAlertingScope) {
	m.alertingScopeField = val
}

// MetricDimensions gets the metric dimensions of this base type
func (m *MetricEvent) MetricDimensions() []MetricEventDimensions {
	return m.metricDimensionsField
}

// SetMetricDimensions sets the metric dimensions of this base type
func (m *MetricEvent) SetMetricDimensions(val []MetricEventDimensions) {
	m.metricDimensionsField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *MetricEvent) UnmarshalJSON(raw []byte) error {
	var data struct {
		AggregationType string `json:"aggregationType,omitempty"`

		AlertCondition *string `json:"alertCondition"`

		AlertingScope json.RawMessage `json:"alertingScope"`

		DealertingSamples *int32 `json:"dealertingSamples"`

		Description *string `json:"description"`

		Enabled *bool `json:"enabled"`

		ID *string `json:"id,omitempty"`

		Metadata *ConfigurationMetadata `json:"metadata,omitempty"`

		MetricDimensions json.RawMessage `json:"metricDimensions"`

		MetricID *string `json:"metricId"`

		Name *string `json:"name"`

		Samples *int32 `json:"samples"`

		Severity string `json:"severity,omitempty"`

		Threshold *float64 `json:"threshold"`

		Unit string `json:"unit,omitempty"`

		ViolatingSamples *int32 `json:"violatingSamples"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var propAlertingScope []MetricEventAlertingScope
	if string(data.AlertingScope) != "null" {
		alertingScope, err := UnmarshalMetricEventAlertingScopeSlice(bytes.NewBuffer(data.AlertingScope), runtime.JSONConsumer())
		if err != nil && err != io.EOF {
			return err
		}
		propAlertingScope = alertingScope
	}

	var propMetricDimensions []MetricEventDimensions
	if string(data.MetricDimensions) != "null" {
		metricDimensions, err := UnmarshalMetricEventDimensionsSlice(bytes.NewBuffer(data.MetricDimensions), runtime.JSONConsumer())
		if err != nil && err != io.EOF {
			return err
		}
		propMetricDimensions = metricDimensions
	}

	var result MetricEvent

	// aggregationType
	result.AggregationType = data.AggregationType

	// alertCondition
	result.AlertCondition = data.AlertCondition

	// alertingScope
	result.alertingScopeField = propAlertingScope

	// dealertingSamples
	result.DealertingSamples = data.DealertingSamples

	// description
	result.Description = data.Description

	// enabled
	result.Enabled = data.Enabled

	// id
	result.ID = data.ID

	// metadata
	result.Metadata = data.Metadata

	// metricDimensions
	result.metricDimensionsField = propMetricDimensions

	// metricId
	result.MetricID = data.MetricID

	// name
	result.Name = data.Name

	// samples
	result.Samples = data.Samples

	// severity
	result.Severity = data.Severity

	// threshold
	result.Threshold = data.Threshold

	// unit
	result.Unit = data.Unit

	// violatingSamples
	result.ViolatingSamples = data.ViolatingSamples

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m MetricEvent) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		AggregationType string `json:"aggregationType,omitempty"`

		AlertCondition *string `json:"alertCondition"`

		DealertingSamples *int32 `json:"dealertingSamples"`

		Description *string `json:"description"`

		Enabled *bool `json:"enabled"`

		ID *string `json:"id,omitempty"`

		Metadata *ConfigurationMetadata `json:"metadata,omitempty"`

		MetricID *string `json:"metricId"`

		Name *string `json:"name"`

		Samples *int32 `json:"samples"`

		Severity string `json:"severity,omitempty"`

		Threshold *float64 `json:"threshold"`

		Unit string `json:"unit,omitempty"`

		ViolatingSamples *int32 `json:"violatingSamples"`
	}{

		AggregationType: m.AggregationType,

		AlertCondition: m.AlertCondition,

		DealertingSamples: m.DealertingSamples,

		Description: m.Description,

		Enabled: m.Enabled,

		ID: m.ID,

		Metadata: m.Metadata,

		MetricID: m.MetricID,

		Name: m.Name,

		Samples: m.Samples,

		Severity: m.Severity,

		Threshold: m.Threshold,

		Unit: m.Unit,

		ViolatingSamples: m.ViolatingSamples,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		AlertingScope []MetricEventAlertingScope `json:"alertingScope"`

		MetricDimensions []MetricEventDimensions `json:"metricDimensions"`
	}{

		AlertingScope: m.alertingScopeField,

		MetricDimensions: m.metricDimensionsField,
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this metric event
func (m *MetricEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAggregationType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAlertCondition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAlertingScope(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDealertingSamples(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetricDimensions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetricID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSamples(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeverity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThreshold(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateViolatingSamples(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var metricEventTypeAggregationTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AVG","COUNT","MAX","MEDIAN","MIN","OF_INTEREST","OF_INTEREST_RATIO","OTHER","OTHER_RATIO","P90","SUM","VALUE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		metricEventTypeAggregationTypePropEnum = append(metricEventTypeAggregationTypePropEnum, v)
	}
}

const (

	// MetricEventAggregationTypeAVG captures enum value "AVG"
	MetricEventAggregationTypeAVG string = "AVG"

	// MetricEventAggregationTypeCOUNT captures enum value "COUNT"
	MetricEventAggregationTypeCOUNT string = "COUNT"

	// MetricEventAggregationTypeMAX captures enum value "MAX"
	MetricEventAggregationTypeMAX string = "MAX"

	// MetricEventAggregationTypeMEDIAN captures enum value "MEDIAN"
	MetricEventAggregationTypeMEDIAN string = "MEDIAN"

	// MetricEventAggregationTypeMIN captures enum value "MIN"
	MetricEventAggregationTypeMIN string = "MIN"

	// MetricEventAggregationTypeOFINTEREST captures enum value "OF_INTEREST"
	MetricEventAggregationTypeOFINTEREST string = "OF_INTEREST"

	// MetricEventAggregationTypeOFINTERESTRATIO captures enum value "OF_INTEREST_RATIO"
	MetricEventAggregationTypeOFINTERESTRATIO string = "OF_INTEREST_RATIO"

	// MetricEventAggregationTypeOTHER captures enum value "OTHER"
	MetricEventAggregationTypeOTHER string = "OTHER"

	// MetricEventAggregationTypeOTHERRATIO captures enum value "OTHER_RATIO"
	MetricEventAggregationTypeOTHERRATIO string = "OTHER_RATIO"

	// MetricEventAggregationTypeP90 captures enum value "P90"
	MetricEventAggregationTypeP90 string = "P90"

	// MetricEventAggregationTypeSUM captures enum value "SUM"
	MetricEventAggregationTypeSUM string = "SUM"

	// MetricEventAggregationTypeVALUE captures enum value "VALUE"
	MetricEventAggregationTypeVALUE string = "VALUE"
)

// prop value enum
func (m *MetricEvent) validateAggregationTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, metricEventTypeAggregationTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MetricEvent) validateAggregationType(formats strfmt.Registry) error {

	if swag.IsZero(m.AggregationType) { // not required
		return nil
	}

	// value enum
	if err := m.validateAggregationTypeEnum("aggregationType", "body", m.AggregationType); err != nil {
		return err
	}

	return nil
}

var metricEventTypeAlertConditionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ABOVE","BELOW"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		metricEventTypeAlertConditionPropEnum = append(metricEventTypeAlertConditionPropEnum, v)
	}
}

const (

	// MetricEventAlertConditionABOVE captures enum value "ABOVE"
	MetricEventAlertConditionABOVE string = "ABOVE"

	// MetricEventAlertConditionBELOW captures enum value "BELOW"
	MetricEventAlertConditionBELOW string = "BELOW"
)

// prop value enum
func (m *MetricEvent) validateAlertConditionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, metricEventTypeAlertConditionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MetricEvent) validateAlertCondition(formats strfmt.Registry) error {

	if err := validate.Required("alertCondition", "body", m.AlertCondition); err != nil {
		return err
	}

	// value enum
	if err := m.validateAlertConditionEnum("alertCondition", "body", *m.AlertCondition); err != nil {
		return err
	}

	return nil
}

func (m *MetricEvent) validateAlertingScope(formats strfmt.Registry) error {

	if swag.IsZero(m.AlertingScope()) { // not required
		return nil
	}

	for i := 0; i < len(m.AlertingScope()); i++ {

		if err := m.alertingScopeField[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("alertingScope" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *MetricEvent) validateDealertingSamples(formats strfmt.Registry) error {

	if err := validate.Required("dealertingSamples", "body", m.DealertingSamples); err != nil {
		return err
	}

	if err := validate.MinimumInt("dealertingSamples", "body", int64(*m.DealertingSamples), 3, false); err != nil {
		return err
	}

	return nil
}

func (m *MetricEvent) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	if err := validate.MinLength("description", "body", string(*m.Description), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", string(*m.Description), 4000); err != nil {
		return err
	}

	return nil
}

func (m *MetricEvent) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *MetricEvent) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.MinLength("id", "body", string(*m.ID), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("id", "body", string(*m.ID), 256); err != nil {
		return err
	}

	if err := validate.Pattern("id", "body", string(*m.ID), `^[a-zA-Z0-9\.\-_]+$`); err != nil {
		return err
	}

	return nil
}

func (m *MetricEvent) validateMetadata(formats strfmt.Registry) error {

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

func (m *MetricEvent) validateMetricDimensions(formats strfmt.Registry) error {

	if swag.IsZero(m.MetricDimensions()) { // not required
		return nil
	}

	for i := 0; i < len(m.MetricDimensions()); i++ {

		if err := m.metricDimensionsField[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metricDimensions" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *MetricEvent) validateMetricID(formats strfmt.Registry) error {

	if err := validate.Required("metricId", "body", m.MetricID); err != nil {
		return err
	}

	return nil
}

func (m *MetricEvent) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 1024); err != nil {
		return err
	}

	return nil
}

func (m *MetricEvent) validateSamples(formats strfmt.Registry) error {

	if err := validate.Required("samples", "body", m.Samples); err != nil {
		return err
	}

	if err := validate.MinimumInt("samples", "body", int64(*m.Samples), 3, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("samples", "body", int64(*m.Samples), 20, false); err != nil {
		return err
	}

	return nil
}

var metricEventTypeSeverityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AVAILABILITY","CUSTOM_ALERT","ERROR","INFO","PERFORMANCE","RESOURCE_CONTENTION"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		metricEventTypeSeverityPropEnum = append(metricEventTypeSeverityPropEnum, v)
	}
}

const (

	// MetricEventSeverityAVAILABILITY captures enum value "AVAILABILITY"
	MetricEventSeverityAVAILABILITY string = "AVAILABILITY"

	// MetricEventSeverityCUSTOMALERT captures enum value "CUSTOM_ALERT"
	MetricEventSeverityCUSTOMALERT string = "CUSTOM_ALERT"

	// MetricEventSeverityERROR captures enum value "ERROR"
	MetricEventSeverityERROR string = "ERROR"

	// MetricEventSeverityINFO captures enum value "INFO"
	MetricEventSeverityINFO string = "INFO"

	// MetricEventSeverityPERFORMANCE captures enum value "PERFORMANCE"
	MetricEventSeverityPERFORMANCE string = "PERFORMANCE"

	// MetricEventSeverityRESOURCECONTENTION captures enum value "RESOURCE_CONTENTION"
	MetricEventSeverityRESOURCECONTENTION string = "RESOURCE_CONTENTION"
)

// prop value enum
func (m *MetricEvent) validateSeverityEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, metricEventTypeSeverityPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MetricEvent) validateSeverity(formats strfmt.Registry) error {

	if swag.IsZero(m.Severity) { // not required
		return nil
	}

	// value enum
	if err := m.validateSeverityEnum("severity", "body", m.Severity); err != nil {
		return err
	}

	return nil
}

func (m *MetricEvent) validateThreshold(formats strfmt.Registry) error {

	if err := validate.Required("threshold", "body", m.Threshold); err != nil {
		return err
	}

	return nil
}

var metricEventTypeUnitPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["BIT","BIT_PER_HOUR","BIT_PER_MINUTE","BIT_PER_SECOND","BYTE","BYTE_PER_HOUR","BYTE_PER_MINUTE","BYTE_PER_SECOND","CORES","COUNT","GIBI_BYTE","GIGA_BYTE","HOUR","KIBI_BYTE","KIBI_BYTE_PER_HOUR","KIBI_BYTE_PER_MINUTE","KIBI_BYTE_PER_SECOND","KILO_BYTE","KILO_BYTE_PER_HOUR","KILO_BYTE_PER_MINUTE","KILO_BYTE_PER_SECOND","MEBI_BYTE","MEBI_BYTE_PER_HOUR","MEBI_BYTE_PER_MINUTE","MEBI_BYTE_PER_SECOND","MEGA_BYTE","MEGA_BYTE_PER_HOUR","MEGA_BYTE_PER_MINUTE","MEGA_BYTE_PER_SECOND","MICRO_SECOND","MILLI_CORES","MILLI_SECOND","MILLI_SECOND_PER_MINUTE","MINUTE","NANO_SECOND","NANO_SECOND_PER_MINUTE","NOT_APPLICABLE","PERCENT","PER_HOUR","PER_MINUTE","PER_SECOND","PROMILLE","RATIO","SECOND","STATE","UNSPECIFIED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		metricEventTypeUnitPropEnum = append(metricEventTypeUnitPropEnum, v)
	}
}

const (

	// MetricEventUnitBIT captures enum value "BIT"
	MetricEventUnitBIT string = "BIT"

	// MetricEventUnitBITPERHOUR captures enum value "BIT_PER_HOUR"
	MetricEventUnitBITPERHOUR string = "BIT_PER_HOUR"

	// MetricEventUnitBITPERMINUTE captures enum value "BIT_PER_MINUTE"
	MetricEventUnitBITPERMINUTE string = "BIT_PER_MINUTE"

	// MetricEventUnitBITPERSECOND captures enum value "BIT_PER_SECOND"
	MetricEventUnitBITPERSECOND string = "BIT_PER_SECOND"

	// MetricEventUnitBYTE captures enum value "BYTE"
	MetricEventUnitBYTE string = "BYTE"

	// MetricEventUnitBYTEPERHOUR captures enum value "BYTE_PER_HOUR"
	MetricEventUnitBYTEPERHOUR string = "BYTE_PER_HOUR"

	// MetricEventUnitBYTEPERMINUTE captures enum value "BYTE_PER_MINUTE"
	MetricEventUnitBYTEPERMINUTE string = "BYTE_PER_MINUTE"

	// MetricEventUnitBYTEPERSECOND captures enum value "BYTE_PER_SECOND"
	MetricEventUnitBYTEPERSECOND string = "BYTE_PER_SECOND"

	// MetricEventUnitCORES captures enum value "CORES"
	MetricEventUnitCORES string = "CORES"

	// MetricEventUnitCOUNT captures enum value "COUNT"
	MetricEventUnitCOUNT string = "COUNT"

	// MetricEventUnitGIBIBYTE captures enum value "GIBI_BYTE"
	MetricEventUnitGIBIBYTE string = "GIBI_BYTE"

	// MetricEventUnitGIGABYTE captures enum value "GIGA_BYTE"
	MetricEventUnitGIGABYTE string = "GIGA_BYTE"

	// MetricEventUnitHOUR captures enum value "HOUR"
	MetricEventUnitHOUR string = "HOUR"

	// MetricEventUnitKIBIBYTE captures enum value "KIBI_BYTE"
	MetricEventUnitKIBIBYTE string = "KIBI_BYTE"

	// MetricEventUnitKIBIBYTEPERHOUR captures enum value "KIBI_BYTE_PER_HOUR"
	MetricEventUnitKIBIBYTEPERHOUR string = "KIBI_BYTE_PER_HOUR"

	// MetricEventUnitKIBIBYTEPERMINUTE captures enum value "KIBI_BYTE_PER_MINUTE"
	MetricEventUnitKIBIBYTEPERMINUTE string = "KIBI_BYTE_PER_MINUTE"

	// MetricEventUnitKIBIBYTEPERSECOND captures enum value "KIBI_BYTE_PER_SECOND"
	MetricEventUnitKIBIBYTEPERSECOND string = "KIBI_BYTE_PER_SECOND"

	// MetricEventUnitKILOBYTE captures enum value "KILO_BYTE"
	MetricEventUnitKILOBYTE string = "KILO_BYTE"

	// MetricEventUnitKILOBYTEPERHOUR captures enum value "KILO_BYTE_PER_HOUR"
	MetricEventUnitKILOBYTEPERHOUR string = "KILO_BYTE_PER_HOUR"

	// MetricEventUnitKILOBYTEPERMINUTE captures enum value "KILO_BYTE_PER_MINUTE"
	MetricEventUnitKILOBYTEPERMINUTE string = "KILO_BYTE_PER_MINUTE"

	// MetricEventUnitKILOBYTEPERSECOND captures enum value "KILO_BYTE_PER_SECOND"
	MetricEventUnitKILOBYTEPERSECOND string = "KILO_BYTE_PER_SECOND"

	// MetricEventUnitMEBIBYTE captures enum value "MEBI_BYTE"
	MetricEventUnitMEBIBYTE string = "MEBI_BYTE"

	// MetricEventUnitMEBIBYTEPERHOUR captures enum value "MEBI_BYTE_PER_HOUR"
	MetricEventUnitMEBIBYTEPERHOUR string = "MEBI_BYTE_PER_HOUR"

	// MetricEventUnitMEBIBYTEPERMINUTE captures enum value "MEBI_BYTE_PER_MINUTE"
	MetricEventUnitMEBIBYTEPERMINUTE string = "MEBI_BYTE_PER_MINUTE"

	// MetricEventUnitMEBIBYTEPERSECOND captures enum value "MEBI_BYTE_PER_SECOND"
	MetricEventUnitMEBIBYTEPERSECOND string = "MEBI_BYTE_PER_SECOND"

	// MetricEventUnitMEGABYTE captures enum value "MEGA_BYTE"
	MetricEventUnitMEGABYTE string = "MEGA_BYTE"

	// MetricEventUnitMEGABYTEPERHOUR captures enum value "MEGA_BYTE_PER_HOUR"
	MetricEventUnitMEGABYTEPERHOUR string = "MEGA_BYTE_PER_HOUR"

	// MetricEventUnitMEGABYTEPERMINUTE captures enum value "MEGA_BYTE_PER_MINUTE"
	MetricEventUnitMEGABYTEPERMINUTE string = "MEGA_BYTE_PER_MINUTE"

	// MetricEventUnitMEGABYTEPERSECOND captures enum value "MEGA_BYTE_PER_SECOND"
	MetricEventUnitMEGABYTEPERSECOND string = "MEGA_BYTE_PER_SECOND"

	// MetricEventUnitMICROSECOND captures enum value "MICRO_SECOND"
	MetricEventUnitMICROSECOND string = "MICRO_SECOND"

	// MetricEventUnitMILLICORES captures enum value "MILLI_CORES"
	MetricEventUnitMILLICORES string = "MILLI_CORES"

	// MetricEventUnitMILLISECOND captures enum value "MILLI_SECOND"
	MetricEventUnitMILLISECOND string = "MILLI_SECOND"

	// MetricEventUnitMILLISECONDPERMINUTE captures enum value "MILLI_SECOND_PER_MINUTE"
	MetricEventUnitMILLISECONDPERMINUTE string = "MILLI_SECOND_PER_MINUTE"

	// MetricEventUnitMINUTE captures enum value "MINUTE"
	MetricEventUnitMINUTE string = "MINUTE"

	// MetricEventUnitNANOSECOND captures enum value "NANO_SECOND"
	MetricEventUnitNANOSECOND string = "NANO_SECOND"

	// MetricEventUnitNANOSECONDPERMINUTE captures enum value "NANO_SECOND_PER_MINUTE"
	MetricEventUnitNANOSECONDPERMINUTE string = "NANO_SECOND_PER_MINUTE"

	// MetricEventUnitNOTAPPLICABLE captures enum value "NOT_APPLICABLE"
	MetricEventUnitNOTAPPLICABLE string = "NOT_APPLICABLE"

	// MetricEventUnitPERCENT captures enum value "PERCENT"
	MetricEventUnitPERCENT string = "PERCENT"

	// MetricEventUnitPERHOUR captures enum value "PER_HOUR"
	MetricEventUnitPERHOUR string = "PER_HOUR"

	// MetricEventUnitPERMINUTE captures enum value "PER_MINUTE"
	MetricEventUnitPERMINUTE string = "PER_MINUTE"

	// MetricEventUnitPERSECOND captures enum value "PER_SECOND"
	MetricEventUnitPERSECOND string = "PER_SECOND"

	// MetricEventUnitPROMILLE captures enum value "PROMILLE"
	MetricEventUnitPROMILLE string = "PROMILLE"

	// MetricEventUnitRATIO captures enum value "RATIO"
	MetricEventUnitRATIO string = "RATIO"

	// MetricEventUnitSECOND captures enum value "SECOND"
	MetricEventUnitSECOND string = "SECOND"

	// MetricEventUnitSTATE captures enum value "STATE"
	MetricEventUnitSTATE string = "STATE"

	// MetricEventUnitUNSPECIFIED captures enum value "UNSPECIFIED"
	MetricEventUnitUNSPECIFIED string = "UNSPECIFIED"
)

// prop value enum
func (m *MetricEvent) validateUnitEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, metricEventTypeUnitPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MetricEvent) validateUnit(formats strfmt.Registry) error {

	if swag.IsZero(m.Unit) { // not required
		return nil
	}

	// value enum
	if err := m.validateUnitEnum("unit", "body", m.Unit); err != nil {
		return err
	}

	return nil
}

func (m *MetricEvent) validateViolatingSamples(formats strfmt.Registry) error {

	if err := validate.Required("violatingSamples", "body", m.ViolatingSamples); err != nil {
		return err
	}

	if err := validate.MinimumInt("violatingSamples", "body", int64(*m.ViolatingSamples), 1, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MetricEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetricEvent) UnmarshalBinary(b []byte) error {
	var res MetricEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
