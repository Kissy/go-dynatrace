// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// CreatePluginEndpointReader is a Reader for the CreatePluginEndpoint structure.
type CreatePluginEndpointReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePluginEndpointReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreatePluginEndpointCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreatePluginEndpointBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreatePluginEndpointCreated creates a CreatePluginEndpointCreated with default headers values
func NewCreatePluginEndpointCreated() *CreatePluginEndpointCreated {
	return &CreatePluginEndpointCreated{}
}

/*CreatePluginEndpointCreated handles this case with default header values.

Success. The plugin endpoint has been created. Response contains the ID of the new endpoint.
*/
type CreatePluginEndpointCreated struct {
	Payload *dynatrace.EntityShortRepresentation
}

func (o *CreatePluginEndpointCreated) Error() string {
	return fmt.Sprintf("[POST /plugins/{id}/endpoints][%d] createPluginEndpointCreated  %+v", 201, o.Payload)
}

func (o *CreatePluginEndpointCreated) GetPayload() *dynatrace.EntityShortRepresentation {
	return o.Payload
}

func (o *CreatePluginEndpointCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.EntityShortRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePluginEndpointBadRequest creates a CreatePluginEndpointBadRequest with default headers values
func NewCreatePluginEndpointBadRequest() *CreatePluginEndpointBadRequest {
	return &CreatePluginEndpointBadRequest{}
}

/*CreatePluginEndpointBadRequest handles this case with default header values.

Failed. The input is invalid.
*/
type CreatePluginEndpointBadRequest struct {
	Payload *dynatrace.ErrorEnvelope
}

func (o *CreatePluginEndpointBadRequest) Error() string {
	return fmt.Sprintf("[POST /plugins/{id}/endpoints][%d] createPluginEndpointBadRequest  %+v", 400, o.Payload)
}

func (o *CreatePluginEndpointBadRequest) GetPayload() *dynatrace.ErrorEnvelope {
	return o.Payload
}

func (o *CreatePluginEndpointBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.ErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
