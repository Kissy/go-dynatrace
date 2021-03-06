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

// ValidateUpdateAzureCredentialReader is a Reader for the ValidateUpdateAzureCredential structure.
type ValidateUpdateAzureCredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateUpdateAzureCredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewValidateUpdateAzureCredentialNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewValidateUpdateAzureCredentialBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewValidateUpdateAzureCredentialNoContent creates a ValidateUpdateAzureCredentialNoContent with default headers values
func NewValidateUpdateAzureCredentialNoContent() *ValidateUpdateAzureCredentialNoContent {
	return &ValidateUpdateAzureCredentialNoContent{}
}

/*ValidateUpdateAzureCredentialNoContent handles this case with default header values.

Validated. The submitted configuration is valid. The response doesn't have a body.
*/
type ValidateUpdateAzureCredentialNoContent struct {
}

func (o *ValidateUpdateAzureCredentialNoContent) Error() string {
	return fmt.Sprintf("[POST /azure/credentials/{id}/validator][%d] validateUpdateAzureCredentialNoContent ", 204)
}

func (o *ValidateUpdateAzureCredentialNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewValidateUpdateAzureCredentialBadRequest creates a ValidateUpdateAzureCredentialBadRequest with default headers values
func NewValidateUpdateAzureCredentialBadRequest() *ValidateUpdateAzureCredentialBadRequest {
	return &ValidateUpdateAzureCredentialBadRequest{}
}

/*ValidateUpdateAzureCredentialBadRequest handles this case with default header values.

Failed. The input is invalid.
*/
type ValidateUpdateAzureCredentialBadRequest struct {
	Payload *dynatrace.ErrorEnvelope
}

func (o *ValidateUpdateAzureCredentialBadRequest) Error() string {
	return fmt.Sprintf("[POST /azure/credentials/{id}/validator][%d] validateUpdateAzureCredentialBadRequest  %+v", 400, o.Payload)
}

func (o *ValidateUpdateAzureCredentialBadRequest) GetPayload() *dynatrace.ErrorEnvelope {
	return o.Payload
}

func (o *ValidateUpdateAzureCredentialBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.ErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
