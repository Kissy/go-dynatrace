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

// CreateApplicationsWebReader is a Reader for the CreateApplicationsWeb structure.
type CreateApplicationsWebReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateApplicationsWebReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateApplicationsWebCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateApplicationsWebBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateApplicationsWebCreated creates a CreateApplicationsWebCreated with default headers values
func NewCreateApplicationsWebCreated() *CreateApplicationsWebCreated {
	return &CreateApplicationsWebCreated{}
}

/*CreateApplicationsWebCreated handles this case with default header values.

Success. The response body contains the ID and name of the new web application.
*/
type CreateApplicationsWebCreated struct {
	Payload *dynatrace.EntityShortRepresentation
}

func (o *CreateApplicationsWebCreated) Error() string {
	return fmt.Sprintf("[POST /applications/web][%d] createApplicationsWebCreated  %+v", 201, o.Payload)
}

func (o *CreateApplicationsWebCreated) GetPayload() *dynatrace.EntityShortRepresentation {
	return o.Payload
}

func (o *CreateApplicationsWebCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.EntityShortRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateApplicationsWebBadRequest creates a CreateApplicationsWebBadRequest with default headers values
func NewCreateApplicationsWebBadRequest() *CreateApplicationsWebBadRequest {
	return &CreateApplicationsWebBadRequest{}
}

/*CreateApplicationsWebBadRequest handles this case with default header values.

Failed. The input is invalid.
*/
type CreateApplicationsWebBadRequest struct {
	Payload *dynatrace.ErrorEnvelope
}

func (o *CreateApplicationsWebBadRequest) Error() string {
	return fmt.Sprintf("[POST /applications/web][%d] createApplicationsWebBadRequest  %+v", 400, o.Payload)
}

func (o *CreateApplicationsWebBadRequest) GetPayload() *dynatrace.ErrorEnvelope {
	return o.Payload
}

func (o *CreateApplicationsWebBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.ErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
