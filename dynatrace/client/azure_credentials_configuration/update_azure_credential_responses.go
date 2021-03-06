// Code generated by go-swagger; DO NOT EDIT.

package azure_credentials_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// UpdateAzureCredentialReader is a Reader for the UpdateAzureCredential structure.
type UpdateAzureCredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAzureCredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateAzureCredentialCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewUpdateAzureCredentialNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateAzureCredentialBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateAzureCredentialCreated creates a UpdateAzureCredentialCreated with default headers values
func NewUpdateAzureCredentialCreated() *UpdateAzureCredentialCreated {
	return &UpdateAzureCredentialCreated{}
}

/*UpdateAzureCredentialCreated handles this case with default header values.

Success. The new Azure credentials configuration has been created. The response body contains the ID of the configuration.
*/
type UpdateAzureCredentialCreated struct {
	Payload *dynatrace.EntityShortRepresentation
}

func (o *UpdateAzureCredentialCreated) Error() string {
	return fmt.Sprintf("[PUT /azure/credentials/{id}][%d] updateAzureCredentialCreated  %+v", 201, o.Payload)
}

func (o *UpdateAzureCredentialCreated) GetPayload() *dynatrace.EntityShortRepresentation {
	return o.Payload
}

func (o *UpdateAzureCredentialCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.EntityShortRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAzureCredentialNoContent creates a UpdateAzureCredentialNoContent with default headers values
func NewUpdateAzureCredentialNoContent() *UpdateAzureCredentialNoContent {
	return &UpdateAzureCredentialNoContent{}
}

/*UpdateAzureCredentialNoContent handles this case with default header values.

Success. The Azure credentials configuration has been updated. Response doesn't have a body.
*/
type UpdateAzureCredentialNoContent struct {
}

func (o *UpdateAzureCredentialNoContent) Error() string {
	return fmt.Sprintf("[PUT /azure/credentials/{id}][%d] updateAzureCredentialNoContent ", 204)
}

func (o *UpdateAzureCredentialNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateAzureCredentialBadRequest creates a UpdateAzureCredentialBadRequest with default headers values
func NewUpdateAzureCredentialBadRequest() *UpdateAzureCredentialBadRequest {
	return &UpdateAzureCredentialBadRequest{}
}

/*UpdateAzureCredentialBadRequest handles this case with default header values.

Failed. The input is invalid.
*/
type UpdateAzureCredentialBadRequest struct {
	Payload *dynatrace.ErrorEnvelope
}

func (o *UpdateAzureCredentialBadRequest) Error() string {
	return fmt.Sprintf("[PUT /azure/credentials/{id}][%d] updateAzureCredentialBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateAzureCredentialBadRequest) GetPayload() *dynatrace.ErrorEnvelope {
	return o.Payload
}

func (o *UpdateAzureCredentialBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.ErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
