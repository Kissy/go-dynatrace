// Code generated by go-swagger; DO NOT EDIT.

package application_detection_rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// ValidateUpdateApplicationDetectionRuleReader is a Reader for the ValidateUpdateApplicationDetectionRule structure.
type ValidateUpdateApplicationDetectionRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateUpdateApplicationDetectionRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewValidateUpdateApplicationDetectionRuleNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewValidateUpdateApplicationDetectionRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewValidateUpdateApplicationDetectionRuleNoContent creates a ValidateUpdateApplicationDetectionRuleNoContent with default headers values
func NewValidateUpdateApplicationDetectionRuleNoContent() *ValidateUpdateApplicationDetectionRuleNoContent {
	return &ValidateUpdateApplicationDetectionRuleNoContent{}
}

/*ValidateUpdateApplicationDetectionRuleNoContent handles this case with default header values.

Validated. The submitted configuration is valid. Response doesn't have a body.
*/
type ValidateUpdateApplicationDetectionRuleNoContent struct {
}

func (o *ValidateUpdateApplicationDetectionRuleNoContent) Error() string {
	return fmt.Sprintf("[POST /applicationDetectionRules/{id}/validator][%d] validateUpdateApplicationDetectionRuleNoContent ", 204)
}

func (o *ValidateUpdateApplicationDetectionRuleNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewValidateUpdateApplicationDetectionRuleBadRequest creates a ValidateUpdateApplicationDetectionRuleBadRequest with default headers values
func NewValidateUpdateApplicationDetectionRuleBadRequest() *ValidateUpdateApplicationDetectionRuleBadRequest {
	return &ValidateUpdateApplicationDetectionRuleBadRequest{}
}

/*ValidateUpdateApplicationDetectionRuleBadRequest handles this case with default header values.

Failed. The input is invalid.
*/
type ValidateUpdateApplicationDetectionRuleBadRequest struct {
	Payload *dynatrace.ErrorEnvelope
}

func (o *ValidateUpdateApplicationDetectionRuleBadRequest) Error() string {
	return fmt.Sprintf("[POST /applicationDetectionRules/{id}/validator][%d] validateUpdateApplicationDetectionRuleBadRequest  %+v", 400, o.Payload)
}

func (o *ValidateUpdateApplicationDetectionRuleBadRequest) GetPayload() *dynatrace.ErrorEnvelope {
	return o.Payload
}

func (o *ValidateUpdateApplicationDetectionRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.ErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
