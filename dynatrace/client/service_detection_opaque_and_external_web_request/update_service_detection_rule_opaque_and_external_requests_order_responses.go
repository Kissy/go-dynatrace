// Code generated by go-swagger; DO NOT EDIT.

package service_detection_opaque_and_external_web_request

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// UpdateServiceDetectionRuleOpaqueAndExternalRequestsOrderReader is a Reader for the UpdateServiceDetectionRuleOpaqueAndExternalRequestsOrder structure.
type UpdateServiceDetectionRuleOpaqueAndExternalRequestsOrderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateServiceDetectionRuleOpaqueAndExternalRequestsOrderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateServiceDetectionRuleOpaqueAndExternalRequestsOrderNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateServiceDetectionRuleOpaqueAndExternalRequestsOrderBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateServiceDetectionRuleOpaqueAndExternalRequestsOrderNoContent creates a UpdateServiceDetectionRuleOpaqueAndExternalRequestsOrderNoContent with default headers values
func NewUpdateServiceDetectionRuleOpaqueAndExternalRequestsOrderNoContent() *UpdateServiceDetectionRuleOpaqueAndExternalRequestsOrderNoContent {
	return &UpdateServiceDetectionRuleOpaqueAndExternalRequestsOrderNoContent{}
}

/*UpdateServiceDetectionRuleOpaqueAndExternalRequestsOrderNoContent handles this case with default header values.

Success. Service detection rules have been reordered. Response doesn't have a body.
*/
type UpdateServiceDetectionRuleOpaqueAndExternalRequestsOrderNoContent struct {
}

func (o *UpdateServiceDetectionRuleOpaqueAndExternalRequestsOrderNoContent) Error() string {
	return fmt.Sprintf("[PUT /service/detectionRules/OPAQUE_AND_EXTERNAL_WEB_REQUEST/order][%d] updateServiceDetectionRuleOpaqueAndExternalRequestsOrderNoContent ", 204)
}

func (o *UpdateServiceDetectionRuleOpaqueAndExternalRequestsOrderNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateServiceDetectionRuleOpaqueAndExternalRequestsOrderBadRequest creates a UpdateServiceDetectionRuleOpaqueAndExternalRequestsOrderBadRequest with default headers values
func NewUpdateServiceDetectionRuleOpaqueAndExternalRequestsOrderBadRequest() *UpdateServiceDetectionRuleOpaqueAndExternalRequestsOrderBadRequest {
	return &UpdateServiceDetectionRuleOpaqueAndExternalRequestsOrderBadRequest{}
}

/*UpdateServiceDetectionRuleOpaqueAndExternalRequestsOrderBadRequest handles this case with default header values.

Failed. The input is invalid.
*/
type UpdateServiceDetectionRuleOpaqueAndExternalRequestsOrderBadRequest struct {
	Payload *dynatrace.ErrorEnvelope
}

func (o *UpdateServiceDetectionRuleOpaqueAndExternalRequestsOrderBadRequest) Error() string {
	return fmt.Sprintf("[PUT /service/detectionRules/OPAQUE_AND_EXTERNAL_WEB_REQUEST/order][%d] updateServiceDetectionRuleOpaqueAndExternalRequestsOrderBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateServiceDetectionRuleOpaqueAndExternalRequestsOrderBadRequest) GetPayload() *dynatrace.ErrorEnvelope {
	return o.Payload
}

func (o *UpdateServiceDetectionRuleOpaqueAndExternalRequestsOrderBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.ErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}