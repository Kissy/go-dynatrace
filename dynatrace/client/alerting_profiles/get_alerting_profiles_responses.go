// Code generated by go-swagger; DO NOT EDIT.

package alerting_profiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// GetAlertingProfilesReader is a Reader for the GetAlertingProfiles structure.
type GetAlertingProfilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlertingProfilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAlertingProfilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAlertingProfilesOK creates a GetAlertingProfilesOK with default headers values
func NewGetAlertingProfilesOK() *GetAlertingProfilesOK {
	return &GetAlertingProfilesOK{}
}

/*GetAlertingProfilesOK handles this case with default header values.

successful operation
*/
type GetAlertingProfilesOK struct {
	Payload *dynatrace.StubList
}

func (o *GetAlertingProfilesOK) Error() string {
	return fmt.Sprintf("[GET /alertingProfiles][%d] getAlertingProfilesOK  %+v", 200, o.Payload)
}

func (o *GetAlertingProfilesOK) GetPayload() *dynatrace.StubList {
	return o.Payload
}

func (o *GetAlertingProfilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.StubList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
