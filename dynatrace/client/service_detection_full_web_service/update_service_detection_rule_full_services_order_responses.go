// Code generated by go-swagger; DO NOT EDIT.

package service_detection_full_web_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// UpdateServiceDetectionRuleFullServicesOrderReader is a Reader for the UpdateServiceDetectionRuleFullServicesOrder structure.
type UpdateServiceDetectionRuleFullServicesOrderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateServiceDetectionRuleFullServicesOrderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateServiceDetectionRuleFullServicesOrderNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateServiceDetectionRuleFullServicesOrderBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateServiceDetectionRuleFullServicesOrderNoContent creates a UpdateServiceDetectionRuleFullServicesOrderNoContent with default headers values
func NewUpdateServiceDetectionRuleFullServicesOrderNoContent() *UpdateServiceDetectionRuleFullServicesOrderNoContent {
	return &UpdateServiceDetectionRuleFullServicesOrderNoContent{}
}

/*UpdateServiceDetectionRuleFullServicesOrderNoContent handles this case with default header values.

Success. Service detection rules have been reordered. Response doesn't have a body.
*/
type UpdateServiceDetectionRuleFullServicesOrderNoContent struct {
}

func (o *UpdateServiceDetectionRuleFullServicesOrderNoContent) Error() string {
	return fmt.Sprintf("[PUT /service/detectionRules/FULL_WEB_SERVICE/order][%d] updateServiceDetectionRuleFullServicesOrderNoContent ", 204)
}

func (o *UpdateServiceDetectionRuleFullServicesOrderNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateServiceDetectionRuleFullServicesOrderBadRequest creates a UpdateServiceDetectionRuleFullServicesOrderBadRequest with default headers values
func NewUpdateServiceDetectionRuleFullServicesOrderBadRequest() *UpdateServiceDetectionRuleFullServicesOrderBadRequest {
	return &UpdateServiceDetectionRuleFullServicesOrderBadRequest{}
}

/*UpdateServiceDetectionRuleFullServicesOrderBadRequest handles this case with default header values.

Failed. The input is invalid.
*/
type UpdateServiceDetectionRuleFullServicesOrderBadRequest struct {
	Payload *dynatrace.ErrorEnvelope
}

func (o *UpdateServiceDetectionRuleFullServicesOrderBadRequest) Error() string {
	return fmt.Sprintf("[PUT /service/detectionRules/FULL_WEB_SERVICE/order][%d] updateServiceDetectionRuleFullServicesOrderBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateServiceDetectionRuleFullServicesOrderBadRequest) GetPayload() *dynatrace.ErrorEnvelope {
	return o.Payload
}

func (o *UpdateServiceDetectionRuleFullServicesOrderBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.ErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
