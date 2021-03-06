// Code generated by go-swagger; DO NOT EDIT.

package mobile_deobfuscation_and_symbolication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// UpdateSymfilePinningReader is a Reader for the UpdateSymfilePinning structure.
type UpdateSymfilePinningReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSymfilePinningReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateSymfilePinningNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateSymfilePinningBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateSymfilePinningNoContent creates a UpdateSymfilePinningNoContent with default headers values
func NewUpdateSymfilePinningNoContent() *UpdateSymfilePinningNoContent {
	return &UpdateSymfilePinningNoContent{}
}

/*UpdateSymfilePinningNoContent handles this case with default header values.

Success. The pinning status of the filed has been updated. Response doesn't have a body.
*/
type UpdateSymfilePinningNoContent struct {
}

func (o *UpdateSymfilePinningNoContent) Error() string {
	return fmt.Sprintf("[PUT /symfiles/{applicationId}/{packageName}/{os}/{versionCode}/{versionName}/pinning][%d] updateSymfilePinningNoContent ", 204)
}

func (o *UpdateSymfilePinningNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateSymfilePinningBadRequest creates a UpdateSymfilePinningBadRequest with default headers values
func NewUpdateSymfilePinningBadRequest() *UpdateSymfilePinningBadRequest {
	return &UpdateSymfilePinningBadRequest{}
}

/*UpdateSymfilePinningBadRequest handles this case with default header values.

Failed. The input is invalid.
*/
type UpdateSymfilePinningBadRequest struct {
	Payload *dynatrace.ErrorEnvelope
}

func (o *UpdateSymfilePinningBadRequest) Error() string {
	return fmt.Sprintf("[PUT /symfiles/{applicationId}/{packageName}/{os}/{versionCode}/{versionName}/pinning][%d] updateSymfilePinningBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateSymfilePinningBadRequest) GetPayload() *dynatrace.ErrorEnvelope {
	return o.Payload
}

func (o *UpdateSymfilePinningBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.ErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
