// Code generated by go-swagger; DO NOT EDIT.

package maintenance_windows

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// ValidateCreateConfigurationReader is a Reader for the ValidateCreateConfiguration structure.
type ValidateCreateConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateCreateConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewValidateCreateConfigurationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewValidateCreateConfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewValidateCreateConfigurationNoContent creates a ValidateCreateConfigurationNoContent with default headers values
func NewValidateCreateConfigurationNoContent() *ValidateCreateConfigurationNoContent {
	return &ValidateCreateConfigurationNoContent{}
}

/*ValidateCreateConfigurationNoContent handles this case with default header values.

Validated. The submitted configuration is valid. Response doesn't have a body.
*/
type ValidateCreateConfigurationNoContent struct {
}

func (o *ValidateCreateConfigurationNoContent) Error() string {
	return fmt.Sprintf("[POST /maintenanceWindows/validator][%d] validateCreateConfigurationNoContent ", 204)
}

func (o *ValidateCreateConfigurationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewValidateCreateConfigurationBadRequest creates a ValidateCreateConfigurationBadRequest with default headers values
func NewValidateCreateConfigurationBadRequest() *ValidateCreateConfigurationBadRequest {
	return &ValidateCreateConfigurationBadRequest{}
}

/*ValidateCreateConfigurationBadRequest handles this case with default header values.

Failed. The input is invalid.
*/
type ValidateCreateConfigurationBadRequest struct {
	Payload *dynatrace.ErrorEnvelope
}

func (o *ValidateCreateConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[POST /maintenanceWindows/validator][%d] validateCreateConfigurationBadRequest  %+v", 400, o.Payload)
}

func (o *ValidateCreateConfigurationBadRequest) GetPayload() *dynatrace.ErrorEnvelope {
	return o.Payload
}

func (o *ValidateCreateConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.ErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}