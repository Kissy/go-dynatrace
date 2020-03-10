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

// UpdateApplicationsWebDefaultReader is a Reader for the UpdateApplicationsWebDefault structure.
type UpdateApplicationsWebDefaultReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateApplicationsWebDefaultReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateApplicationsWebDefaultNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateApplicationsWebDefaultBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateApplicationsWebDefaultNoContent creates a UpdateApplicationsWebDefaultNoContent with default headers values
func NewUpdateApplicationsWebDefaultNoContent() *UpdateApplicationsWebDefaultNoContent {
	return &UpdateApplicationsWebDefaultNoContent{}
}

/*UpdateApplicationsWebDefaultNoContent handles this case with default header values.

Success. Configuration has been updated. Response doesn't have a body.
*/
type UpdateApplicationsWebDefaultNoContent struct {
}

func (o *UpdateApplicationsWebDefaultNoContent) Error() string {
	return fmt.Sprintf("[PUT /applications/web/default][%d] updateApplicationsWebDefaultNoContent ", 204)
}

func (o *UpdateApplicationsWebDefaultNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateApplicationsWebDefaultBadRequest creates a UpdateApplicationsWebDefaultBadRequest with default headers values
func NewUpdateApplicationsWebDefaultBadRequest() *UpdateApplicationsWebDefaultBadRequest {
	return &UpdateApplicationsWebDefaultBadRequest{}
}

/*UpdateApplicationsWebDefaultBadRequest handles this case with default header values.

Failed. The input is invalid.
*/
type UpdateApplicationsWebDefaultBadRequest struct {
	Payload *dynatrace.ErrorEnvelope
}

func (o *UpdateApplicationsWebDefaultBadRequest) Error() string {
	return fmt.Sprintf("[PUT /applications/web/default][%d] updateApplicationsWebDefaultBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateApplicationsWebDefaultBadRequest) GetPayload() *dynatrace.ErrorEnvelope {
	return o.Payload
}

func (o *UpdateApplicationsWebDefaultBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.ErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}