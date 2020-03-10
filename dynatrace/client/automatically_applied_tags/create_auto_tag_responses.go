// Code generated by go-swagger; DO NOT EDIT.

package automatically_applied_tags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// CreateAutoTagReader is a Reader for the CreateAutoTag structure.
type CreateAutoTagReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAutoTagReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateAutoTagCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateAutoTagBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateAutoTagCreated creates a CreateAutoTagCreated with default headers values
func NewCreateAutoTagCreated() *CreateAutoTagCreated {
	return &CreateAutoTagCreated{}
}

/*CreateAutoTagCreated handles this case with default header values.

Success. The auto-tag has been created. The response body contains the ID of the new auto-tag.
*/
type CreateAutoTagCreated struct {
	Payload *dynatrace.EntityShortRepresentation
}

func (o *CreateAutoTagCreated) Error() string {
	return fmt.Sprintf("[POST /autoTags][%d] createAutoTagCreated  %+v", 201, o.Payload)
}

func (o *CreateAutoTagCreated) GetPayload() *dynatrace.EntityShortRepresentation {
	return o.Payload
}

func (o *CreateAutoTagCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.EntityShortRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAutoTagBadRequest creates a CreateAutoTagBadRequest with default headers values
func NewCreateAutoTagBadRequest() *CreateAutoTagBadRequest {
	return &CreateAutoTagBadRequest{}
}

/*CreateAutoTagBadRequest handles this case with default header values.

Failed. The input is invalid.
*/
type CreateAutoTagBadRequest struct {
	Payload *dynatrace.ErrorEnvelope
}

func (o *CreateAutoTagBadRequest) Error() string {
	return fmt.Sprintf("[POST /autoTags][%d] createAutoTagBadRequest  %+v", 400, o.Payload)
}

func (o *CreateAutoTagBadRequest) GetPayload() *dynatrace.ErrorEnvelope {
	return o.Payload
}

func (o *CreateAutoTagBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.ErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}