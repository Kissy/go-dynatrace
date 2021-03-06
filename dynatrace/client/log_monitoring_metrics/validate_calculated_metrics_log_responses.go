// Code generated by go-swagger; DO NOT EDIT.

package log_monitoring_metrics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// ValidateCalculatedMetricsLogReader is a Reader for the ValidateCalculatedMetricsLog structure.
type ValidateCalculatedMetricsLogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateCalculatedMetricsLogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewValidateCalculatedMetricsLogNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewValidateCalculatedMetricsLogBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewValidateCalculatedMetricsLogNoContent creates a ValidateCalculatedMetricsLogNoContent with default headers values
func NewValidateCalculatedMetricsLogNoContent() *ValidateCalculatedMetricsLogNoContent {
	return &ValidateCalculatedMetricsLogNoContent{}
}

/*ValidateCalculatedMetricsLogNoContent handles this case with default header values.

Validated. The submitted configuration is valid. Response doesn't have a body.
*/
type ValidateCalculatedMetricsLogNoContent struct {
}

func (o *ValidateCalculatedMetricsLogNoContent) Error() string {
	return fmt.Sprintf("[POST /calculatedMetrics/log/{metricKey}/validator][%d] validateCalculatedMetricsLogNoContent ", 204)
}

func (o *ValidateCalculatedMetricsLogNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewValidateCalculatedMetricsLogBadRequest creates a ValidateCalculatedMetricsLogBadRequest with default headers values
func NewValidateCalculatedMetricsLogBadRequest() *ValidateCalculatedMetricsLogBadRequest {
	return &ValidateCalculatedMetricsLogBadRequest{}
}

/*ValidateCalculatedMetricsLogBadRequest handles this case with default header values.

Failed. The input is invalid.
*/
type ValidateCalculatedMetricsLogBadRequest struct {
	Payload *dynatrace.ErrorEnvelope
}

func (o *ValidateCalculatedMetricsLogBadRequest) Error() string {
	return fmt.Sprintf("[POST /calculatedMetrics/log/{metricKey}/validator][%d] validateCalculatedMetricsLogBadRequest  %+v", 400, o.Payload)
}

func (o *ValidateCalculatedMetricsLogBadRequest) GetPayload() *dynatrace.ErrorEnvelope {
	return o.Payload
}

func (o *ValidateCalculatedMetricsLogBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.ErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
