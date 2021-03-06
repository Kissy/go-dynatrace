// Code generated by go-swagger; DO NOT EDIT.

package web_application_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// UpdateApplicationsWebDataPrivacyReader is a Reader for the UpdateApplicationsWebDataPrivacy structure.
type UpdateApplicationsWebDataPrivacyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateApplicationsWebDataPrivacyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateApplicationsWebDataPrivacyNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateApplicationsWebDataPrivacyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateApplicationsWebDataPrivacyNoContent creates a UpdateApplicationsWebDataPrivacyNoContent with default headers values
func NewUpdateApplicationsWebDataPrivacyNoContent() *UpdateApplicationsWebDataPrivacyNoContent {
	return &UpdateApplicationsWebDataPrivacyNoContent{}
}

/*UpdateApplicationsWebDataPrivacyNoContent handles this case with default header values.

Success. Data privacy settings have been updated. Response doesn't have a body.
*/
type UpdateApplicationsWebDataPrivacyNoContent struct {
}

func (o *UpdateApplicationsWebDataPrivacyNoContent) Error() string {
	return fmt.Sprintf("[PUT /applications/web/{id}/dataPrivacy][%d] updateApplicationsWebDataPrivacyNoContent ", 204)
}

func (o *UpdateApplicationsWebDataPrivacyNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateApplicationsWebDataPrivacyBadRequest creates a UpdateApplicationsWebDataPrivacyBadRequest with default headers values
func NewUpdateApplicationsWebDataPrivacyBadRequest() *UpdateApplicationsWebDataPrivacyBadRequest {
	return &UpdateApplicationsWebDataPrivacyBadRequest{}
}

/*UpdateApplicationsWebDataPrivacyBadRequest handles this case with default header values.

Failed. The input is invalid.
*/
type UpdateApplicationsWebDataPrivacyBadRequest struct {
	Payload *dynatrace.ErrorEnvelope
}

func (o *UpdateApplicationsWebDataPrivacyBadRequest) Error() string {
	return fmt.Sprintf("[PUT /applications/web/{id}/dataPrivacy][%d] updateApplicationsWebDataPrivacyBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateApplicationsWebDataPrivacyBadRequest) GetPayload() *dynatrace.ErrorEnvelope {
	return o.Payload
}

func (o *UpdateApplicationsWebDataPrivacyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.ErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
