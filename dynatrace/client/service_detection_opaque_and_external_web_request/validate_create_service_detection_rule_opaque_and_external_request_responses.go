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

// ValidateCreateServiceDetectionRuleOpaqueAndExternalRequestReader is a Reader for the ValidateCreateServiceDetectionRuleOpaqueAndExternalRequest structure.
type ValidateCreateServiceDetectionRuleOpaqueAndExternalRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateCreateServiceDetectionRuleOpaqueAndExternalRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewValidateCreateServiceDetectionRuleOpaqueAndExternalRequestNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewValidateCreateServiceDetectionRuleOpaqueAndExternalRequestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewValidateCreateServiceDetectionRuleOpaqueAndExternalRequestNoContent creates a ValidateCreateServiceDetectionRuleOpaqueAndExternalRequestNoContent with default headers values
func NewValidateCreateServiceDetectionRuleOpaqueAndExternalRequestNoContent() *ValidateCreateServiceDetectionRuleOpaqueAndExternalRequestNoContent {
	return &ValidateCreateServiceDetectionRuleOpaqueAndExternalRequestNoContent{}
}

/*ValidateCreateServiceDetectionRuleOpaqueAndExternalRequestNoContent handles this case with default header values.

Validated. The service detection rule is valid. Response doesn't have a body.
*/
type ValidateCreateServiceDetectionRuleOpaqueAndExternalRequestNoContent struct {
}

func (o *ValidateCreateServiceDetectionRuleOpaqueAndExternalRequestNoContent) Error() string {
	return fmt.Sprintf("[POST /service/detectionRules/OPAQUE_AND_EXTERNAL_WEB_REQUEST/validator][%d] validateCreateServiceDetectionRuleOpaqueAndExternalRequestNoContent ", 204)
}

func (o *ValidateCreateServiceDetectionRuleOpaqueAndExternalRequestNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewValidateCreateServiceDetectionRuleOpaqueAndExternalRequestBadRequest creates a ValidateCreateServiceDetectionRuleOpaqueAndExternalRequestBadRequest with default headers values
func NewValidateCreateServiceDetectionRuleOpaqueAndExternalRequestBadRequest() *ValidateCreateServiceDetectionRuleOpaqueAndExternalRequestBadRequest {
	return &ValidateCreateServiceDetectionRuleOpaqueAndExternalRequestBadRequest{}
}

/*ValidateCreateServiceDetectionRuleOpaqueAndExternalRequestBadRequest handles this case with default header values.

Failed. The input is invalid.
*/
type ValidateCreateServiceDetectionRuleOpaqueAndExternalRequestBadRequest struct {
	Payload *dynatrace.ErrorEnvelope
}

func (o *ValidateCreateServiceDetectionRuleOpaqueAndExternalRequestBadRequest) Error() string {
	return fmt.Sprintf("[POST /service/detectionRules/OPAQUE_AND_EXTERNAL_WEB_REQUEST/validator][%d] validateCreateServiceDetectionRuleOpaqueAndExternalRequestBadRequest  %+v", 400, o.Payload)
}

func (o *ValidateCreateServiceDetectionRuleOpaqueAndExternalRequestBadRequest) GetPayload() *dynatrace.ErrorEnvelope {
	return o.Payload
}

func (o *ValidateCreateServiceDetectionRuleOpaqueAndExternalRequestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.ErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
