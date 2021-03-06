// Code generated by go-swagger; DO NOT EDIT.

package anomaly_detection_database_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// UpdateAnomalyDetectionDatabaseServicesReader is a Reader for the UpdateAnomalyDetectionDatabaseServices structure.
type UpdateAnomalyDetectionDatabaseServicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAnomalyDetectionDatabaseServicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateAnomalyDetectionDatabaseServicesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateAnomalyDetectionDatabaseServicesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateAnomalyDetectionDatabaseServicesNoContent creates a UpdateAnomalyDetectionDatabaseServicesNoContent with default headers values
func NewUpdateAnomalyDetectionDatabaseServicesNoContent() *UpdateAnomalyDetectionDatabaseServicesNoContent {
	return &UpdateAnomalyDetectionDatabaseServicesNoContent{}
}

/*UpdateAnomalyDetectionDatabaseServicesNoContent handles this case with default header values.

Success. Configuration has been updated. Response doesn't have a body.
*/
type UpdateAnomalyDetectionDatabaseServicesNoContent struct {
}

func (o *UpdateAnomalyDetectionDatabaseServicesNoContent) Error() string {
	return fmt.Sprintf("[PUT /anomalyDetection/databaseServices][%d] updateAnomalyDetectionDatabaseServicesNoContent ", 204)
}

func (o *UpdateAnomalyDetectionDatabaseServicesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateAnomalyDetectionDatabaseServicesBadRequest creates a UpdateAnomalyDetectionDatabaseServicesBadRequest with default headers values
func NewUpdateAnomalyDetectionDatabaseServicesBadRequest() *UpdateAnomalyDetectionDatabaseServicesBadRequest {
	return &UpdateAnomalyDetectionDatabaseServicesBadRequest{}
}

/*UpdateAnomalyDetectionDatabaseServicesBadRequest handles this case with default header values.

Failed. The input is invalid.
*/
type UpdateAnomalyDetectionDatabaseServicesBadRequest struct {
	Payload *dynatrace.ErrorEnvelope
}

func (o *UpdateAnomalyDetectionDatabaseServicesBadRequest) Error() string {
	return fmt.Sprintf("[PUT /anomalyDetection/databaseServices][%d] updateAnomalyDetectionDatabaseServicesBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateAnomalyDetectionDatabaseServicesBadRequest) GetPayload() *dynatrace.ErrorEnvelope {
	return o.Payload
}

func (o *UpdateAnomalyDetectionDatabaseServicesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.ErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
