// Code generated by go-swagger; DO NOT EDIT.

package reports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// GetReportsReader is a Reader for the GetReports structure.
type GetReportsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReportsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReportsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetReportsOK creates a GetReportsOK with default headers values
func NewGetReportsOK() *GetReportsOK {
	return &GetReportsOK{}
}

/*GetReportsOK handles this case with default header values.

successful operation
*/
type GetReportsOK struct {
	Payload *dynatrace.ReportStubList
}

func (o *GetReportsOK) Error() string {
	return fmt.Sprintf("[GET /reports][%d] getReportsOK  %+v", 200, o.Payload)
}

func (o *GetReportsOK) GetPayload() *dynatrace.ReportStubList {
	return o.Payload
}

func (o *GetReportsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.ReportStubList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
