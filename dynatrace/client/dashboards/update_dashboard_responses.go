// Code generated by go-swagger; DO NOT EDIT.

package dashboards

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// UpdateDashboardReader is a Reader for the UpdateDashboard structure.
type UpdateDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateDashboardCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewUpdateDashboardNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateDashboardBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateDashboardCreated creates a UpdateDashboardCreated with default headers values
func NewUpdateDashboardCreated() *UpdateDashboardCreated {
	return &UpdateDashboardCreated{}
}

/*UpdateDashboardCreated handles this case with default header values.

Success. The new dashboard has been created. Response doesn't have a body.
*/
type UpdateDashboardCreated struct {
	Payload *dynatrace.EntityShortRepresentation
}

func (o *UpdateDashboardCreated) Error() string {
	return fmt.Sprintf("[PUT /dashboards/{id}][%d] updateDashboardCreated  %+v", 201, o.Payload)
}

func (o *UpdateDashboardCreated) GetPayload() *dynatrace.EntityShortRepresentation {
	return o.Payload
}

func (o *UpdateDashboardCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.EntityShortRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDashboardNoContent creates a UpdateDashboardNoContent with default headers values
func NewUpdateDashboardNoContent() *UpdateDashboardNoContent {
	return &UpdateDashboardNoContent{}
}

/*UpdateDashboardNoContent handles this case with default header values.

Success. The dashboard has been updated. Response doesn't have a body
*/
type UpdateDashboardNoContent struct {
}

func (o *UpdateDashboardNoContent) Error() string {
	return fmt.Sprintf("[PUT /dashboards/{id}][%d] updateDashboardNoContent ", 204)
}

func (o *UpdateDashboardNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateDashboardBadRequest creates a UpdateDashboardBadRequest with default headers values
func NewUpdateDashboardBadRequest() *UpdateDashboardBadRequest {
	return &UpdateDashboardBadRequest{}
}

/*UpdateDashboardBadRequest handles this case with default header values.

Failed. The input is invalid.
*/
type UpdateDashboardBadRequest struct {
	Payload *dynatrace.ErrorEnvelope
}

func (o *UpdateDashboardBadRequest) Error() string {
	return fmt.Sprintf("[PUT /dashboards/{id}][%d] updateDashboardBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateDashboardBadRequest) GetPayload() *dynatrace.ErrorEnvelope {
	return o.Payload
}

func (o *UpdateDashboardBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.ErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
