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

// UpdateServiceResourceNamingReader is a Reader for the UpdateServiceResourceNaming structure.
type UpdateServiceResourceNamingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateServiceResourceNamingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateServiceResourceNamingNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateServiceResourceNamingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateServiceResourceNamingNoContent creates a UpdateServiceResourceNamingNoContent with default headers values
func NewUpdateServiceResourceNamingNoContent() *UpdateServiceResourceNamingNoContent {
	return &UpdateServiceResourceNamingNoContent{}
}

/*UpdateServiceResourceNamingNoContent handles this case with default header values.

Success. The configuration is updated. Response does not have a body.
*/
type UpdateServiceResourceNamingNoContent struct {
}

func (o *UpdateServiceResourceNamingNoContent) Error() string {
	return fmt.Sprintf("[PUT /service/resourceNaming][%d] updateServiceResourceNamingNoContent ", 204)
}

func (o *UpdateServiceResourceNamingNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateServiceResourceNamingBadRequest creates a UpdateServiceResourceNamingBadRequest with default headers values
func NewUpdateServiceResourceNamingBadRequest() *UpdateServiceResourceNamingBadRequest {
	return &UpdateServiceResourceNamingBadRequest{}
}

/*UpdateServiceResourceNamingBadRequest handles this case with default header values.

Failed. The input is invalid. See the response body for details.
*/
type UpdateServiceResourceNamingBadRequest struct {
	Payload *dynatrace.ErrorEnvelope
}

func (o *UpdateServiceResourceNamingBadRequest) Error() string {
	return fmt.Sprintf("[PUT /service/resourceNaming][%d] updateServiceResourceNamingBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateServiceResourceNamingBadRequest) GetPayload() *dynatrace.ErrorEnvelope {
	return o.Payload
}

func (o *UpdateServiceResourceNamingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.ErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
