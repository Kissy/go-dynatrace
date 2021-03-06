// Code generated by go-swagger; DO NOT EDIT.

package r_u_m_metrics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// ValidateUpdateCalculatedMetricsRumReader is a Reader for the ValidateUpdateCalculatedMetricsRum structure.
type ValidateUpdateCalculatedMetricsRumReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateUpdateCalculatedMetricsRumReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewValidateUpdateCalculatedMetricsRumNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewValidateUpdateCalculatedMetricsRumBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewValidateUpdateCalculatedMetricsRumNoContent creates a ValidateUpdateCalculatedMetricsRumNoContent with default headers values
func NewValidateUpdateCalculatedMetricsRumNoContent() *ValidateUpdateCalculatedMetricsRumNoContent {
	return &ValidateUpdateCalculatedMetricsRumNoContent{}
}

/*ValidateUpdateCalculatedMetricsRumNoContent handles this case with default header values.

Validated. The submitted rum metric is valid. Response doesn't have a body
*/
type ValidateUpdateCalculatedMetricsRumNoContent struct {
}

func (o *ValidateUpdateCalculatedMetricsRumNoContent) Error() string {
	return fmt.Sprintf("[POST /calculatedMetrics/rum/{metricKey}/validator][%d] validateUpdateCalculatedMetricsRumNoContent ", 204)
}

func (o *ValidateUpdateCalculatedMetricsRumNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewValidateUpdateCalculatedMetricsRumBadRequest creates a ValidateUpdateCalculatedMetricsRumBadRequest with default headers values
func NewValidateUpdateCalculatedMetricsRumBadRequest() *ValidateUpdateCalculatedMetricsRumBadRequest {
	return &ValidateUpdateCalculatedMetricsRumBadRequest{}
}

/*ValidateUpdateCalculatedMetricsRumBadRequest handles this case with default header values.

Failed. The input is invalid.
*/
type ValidateUpdateCalculatedMetricsRumBadRequest struct {
	Payload *dynatrace.ErrorEnvelope
}

func (o *ValidateUpdateCalculatedMetricsRumBadRequest) Error() string {
	return fmt.Sprintf("[POST /calculatedMetrics/rum/{metricKey}/validator][%d] validateUpdateCalculatedMetricsRumBadRequest  %+v", 400, o.Payload)
}

func (o *ValidateUpdateCalculatedMetricsRumBadRequest) GetPayload() *dynatrace.ErrorEnvelope {
	return o.Payload
}

func (o *ValidateUpdateCalculatedMetricsRumBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.ErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
