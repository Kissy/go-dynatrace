// Code generated by go-swagger; DO NOT EDIT.

package service_custom_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// GetServiceCustomServiceReader is a Reader for the GetServiceCustomService structure.
type GetServiceCustomServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServiceCustomServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServiceCustomServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetServiceCustomServiceOK creates a GetServiceCustomServiceOK with default headers values
func NewGetServiceCustomServiceOK() *GetServiceCustomServiceOK {
	return &GetServiceCustomServiceOK{}
}

/*GetServiceCustomServiceOK handles this case with default header values.

successful operation
*/
type GetServiceCustomServiceOK struct {
	Payload *dynatrace.CustomService
}

func (o *GetServiceCustomServiceOK) Error() string {
	return fmt.Sprintf("[GET /service/customServices/{technology}/{id}][%d] getServiceCustomServiceOK  %+v", 200, o.Payload)
}

func (o *GetServiceCustomServiceOK) GetPayload() *dynatrace.CustomService {
	return o.Payload
}

func (o *GetServiceCustomServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.CustomService)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
