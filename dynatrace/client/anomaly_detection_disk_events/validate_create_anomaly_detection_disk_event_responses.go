// Code generated by go-swagger; DO NOT EDIT.

package anomaly_detection_disk_events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// ValidateCreateAnomalyDetectionDiskEventReader is a Reader for the ValidateCreateAnomalyDetectionDiskEvent structure.
type ValidateCreateAnomalyDetectionDiskEventReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateCreateAnomalyDetectionDiskEventReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewValidateCreateAnomalyDetectionDiskEventNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewValidateCreateAnomalyDetectionDiskEventBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewValidateCreateAnomalyDetectionDiskEventNoContent creates a ValidateCreateAnomalyDetectionDiskEventNoContent with default headers values
func NewValidateCreateAnomalyDetectionDiskEventNoContent() *ValidateCreateAnomalyDetectionDiskEventNoContent {
	return &ValidateCreateAnomalyDetectionDiskEventNoContent{}
}

/*ValidateCreateAnomalyDetectionDiskEventNoContent handles this case with default header values.

Validated. The submitted disk event rule is valid. Response doesn't have a body.
*/
type ValidateCreateAnomalyDetectionDiskEventNoContent struct {
}

func (o *ValidateCreateAnomalyDetectionDiskEventNoContent) Error() string {
	return fmt.Sprintf("[POST /anomalyDetection/diskEvents/validator][%d] validateCreateAnomalyDetectionDiskEventNoContent ", 204)
}

func (o *ValidateCreateAnomalyDetectionDiskEventNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewValidateCreateAnomalyDetectionDiskEventBadRequest creates a ValidateCreateAnomalyDetectionDiskEventBadRequest with default headers values
func NewValidateCreateAnomalyDetectionDiskEventBadRequest() *ValidateCreateAnomalyDetectionDiskEventBadRequest {
	return &ValidateCreateAnomalyDetectionDiskEventBadRequest{}
}

/*ValidateCreateAnomalyDetectionDiskEventBadRequest handles this case with default header values.

Failed. The input is invalid.
*/
type ValidateCreateAnomalyDetectionDiskEventBadRequest struct {
	Payload *dynatrace.ErrorEnvelope
}

func (o *ValidateCreateAnomalyDetectionDiskEventBadRequest) Error() string {
	return fmt.Sprintf("[POST /anomalyDetection/diskEvents/validator][%d] validateCreateAnomalyDetectionDiskEventBadRequest  %+v", 400, o.Payload)
}

func (o *ValidateCreateAnomalyDetectionDiskEventBadRequest) GetPayload() *dynatrace.ErrorEnvelope {
	return o.Payload
}

func (o *ValidateCreateAnomalyDetectionDiskEventBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.ErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
