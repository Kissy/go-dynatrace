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

// GetApplicationSymfilesReader is a Reader for the GetApplicationSymfiles structure.
type GetApplicationSymfilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApplicationSymfilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApplicationSymfilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApplicationSymfilesOK creates a GetApplicationSymfilesOK with default headers values
func NewGetApplicationSymfilesOK() *GetApplicationSymfilesOK {
	return &GetApplicationSymfilesOK{}
}

/*GetApplicationSymfilesOK handles this case with default header values.

successful operation
*/
type GetApplicationSymfilesOK struct {
	Payload *dynatrace.SymbolFileList
}

func (o *GetApplicationSymfilesOK) Error() string {
	return fmt.Sprintf("[GET /symfiles/{applicationId}][%d] getApplicationSymfilesOK  %+v", 200, o.Payload)
}

func (o *GetApplicationSymfilesOK) GetPayload() *dynatrace.SymbolFileList {
	return o.Payload
}

func (o *GetApplicationSymfilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.SymbolFileList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}