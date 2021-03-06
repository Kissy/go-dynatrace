// Code generated by go-swagger; DO NOT EDIT.

package service_request_attributes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// GetServiceRequestAttributesReader is a Reader for the GetServiceRequestAttributes structure.
type GetServiceRequestAttributesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServiceRequestAttributesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServiceRequestAttributesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetServiceRequestAttributesOK creates a GetServiceRequestAttributesOK with default headers values
func NewGetServiceRequestAttributesOK() *GetServiceRequestAttributesOK {
	return &GetServiceRequestAttributesOK{}
}

/*GetServiceRequestAttributesOK handles this case with default header values.

successful operation
*/
type GetServiceRequestAttributesOK struct {
	Payload *dynatrace.StubList
}

func (o *GetServiceRequestAttributesOK) Error() string {
	return fmt.Sprintf("[GET /service/requestAttributes][%d] getServiceRequestAttributesOK  %+v", 200, o.Payload)
}

func (o *GetServiceRequestAttributesOK) GetPayload() *dynatrace.StubList {
	return o.Payload
}

func (o *GetServiceRequestAttributesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.StubList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
