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

// UpdateApplicationDetectionRuleReader is a Reader for the UpdateApplicationDetectionRule structure.
type UpdateApplicationDetectionRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateApplicationDetectionRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateApplicationDetectionRuleCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewUpdateApplicationDetectionRuleNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateApplicationDetectionRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateApplicationDetectionRuleCreated creates a UpdateApplicationDetectionRuleCreated with default headers values
func NewUpdateApplicationDetectionRuleCreated() *UpdateApplicationDetectionRuleCreated {
	return &UpdateApplicationDetectionRuleCreated{}
}

/*UpdateApplicationDetectionRuleCreated handles this case with default header values.

Success. Application detection rule has been created. Response contains the ID of the new rule.
*/
type UpdateApplicationDetectionRuleCreated struct {
	Payload *dynatrace.EntityShortRepresentation
}

func (o *UpdateApplicationDetectionRuleCreated) Error() string {
	return fmt.Sprintf("[PUT /applicationDetectionRules/{id}][%d] updateApplicationDetectionRuleCreated  %+v", 201, o.Payload)
}

func (o *UpdateApplicationDetectionRuleCreated) GetPayload() *dynatrace.EntityShortRepresentation {
	return o.Payload
}

func (o *UpdateApplicationDetectionRuleCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.EntityShortRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateApplicationDetectionRuleNoContent creates a UpdateApplicationDetectionRuleNoContent with default headers values
func NewUpdateApplicationDetectionRuleNoContent() *UpdateApplicationDetectionRuleNoContent {
	return &UpdateApplicationDetectionRuleNoContent{}
}

/*UpdateApplicationDetectionRuleNoContent handles this case with default header values.

Success. Application detection rule has been updated. Response doesn't have a body.
*/
type UpdateApplicationDetectionRuleNoContent struct {
}

func (o *UpdateApplicationDetectionRuleNoContent) Error() string {
	return fmt.Sprintf("[PUT /applicationDetectionRules/{id}][%d] updateApplicationDetectionRuleNoContent ", 204)
}

func (o *UpdateApplicationDetectionRuleNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateApplicationDetectionRuleBadRequest creates a UpdateApplicationDetectionRuleBadRequest with default headers values
func NewUpdateApplicationDetectionRuleBadRequest() *UpdateApplicationDetectionRuleBadRequest {
	return &UpdateApplicationDetectionRuleBadRequest{}
}

/*UpdateApplicationDetectionRuleBadRequest handles this case with default header values.

Failed. The input is invalid.
*/
type UpdateApplicationDetectionRuleBadRequest struct {
	Payload *dynatrace.ErrorEnvelope
}

func (o *UpdateApplicationDetectionRuleBadRequest) Error() string {
	return fmt.Sprintf("[PUT /applicationDetectionRules/{id}][%d] updateApplicationDetectionRuleBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateApplicationDetectionRuleBadRequest) GetPayload() *dynatrace.ErrorEnvelope {
	return o.Payload
}

func (o *UpdateApplicationDetectionRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.ErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
