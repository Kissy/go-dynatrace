// Code generated by go-swagger; DO NOT EDIT.

package service_request_naming

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// CreateServiceRequestNamingReader is a Reader for the CreateServiceRequestNaming structure.
type CreateServiceRequestNamingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateServiceRequestNamingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateServiceRequestNamingCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateServiceRequestNamingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateServiceRequestNamingCreated creates a CreateServiceRequestNamingCreated with default headers values
func NewCreateServiceRequestNamingCreated() *CreateServiceRequestNamingCreated {
	return &CreateServiceRequestNamingCreated{}
}

/*CreateServiceRequestNamingCreated handles this case with default header values.

Success. The request naming has been created. Response contains the new service's ID.
*/
type CreateServiceRequestNamingCreated struct {
	Payload *dynatrace.EntityShortRepresentation
}

func (o *CreateServiceRequestNamingCreated) Error() string {
	return fmt.Sprintf("[POST /service/requestNaming][%d] createServiceRequestNamingCreated  %+v", 201, o.Payload)
}

func (o *CreateServiceRequestNamingCreated) GetPayload() *dynatrace.EntityShortRepresentation {
	return o.Payload
}

func (o *CreateServiceRequestNamingCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.EntityShortRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateServiceRequestNamingBadRequest creates a CreateServiceRequestNamingBadRequest with default headers values
func NewCreateServiceRequestNamingBadRequest() *CreateServiceRequestNamingBadRequest {
	return &CreateServiceRequestNamingBadRequest{}
}

/*CreateServiceRequestNamingBadRequest handles this case with default header values.

Failed. The input is invalid.
*/
type CreateServiceRequestNamingBadRequest struct {
	Payload *dynatrace.ErrorEnvelope
}

func (o *CreateServiceRequestNamingBadRequest) Error() string {
	return fmt.Sprintf("[POST /service/requestNaming][%d] createServiceRequestNamingBadRequest  %+v", 400, o.Payload)
}

func (o *CreateServiceRequestNamingBadRequest) GetPayload() *dynatrace.ErrorEnvelope {
	return o.Payload
}

func (o *CreateServiceRequestNamingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.ErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
