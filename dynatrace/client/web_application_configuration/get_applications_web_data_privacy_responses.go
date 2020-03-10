// Code generated by go-swagger; DO NOT EDIT.

package web_application_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// GetApplicationsWebDataPrivacyReader is a Reader for the GetApplicationsWebDataPrivacy structure.
type GetApplicationsWebDataPrivacyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApplicationsWebDataPrivacyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApplicationsWebDataPrivacyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApplicationsWebDataPrivacyOK creates a GetApplicationsWebDataPrivacyOK with default headers values
func NewGetApplicationsWebDataPrivacyOK() *GetApplicationsWebDataPrivacyOK {
	return &GetApplicationsWebDataPrivacyOK{}
}

/*GetApplicationsWebDataPrivacyOK handles this case with default header values.

successful operation
*/
type GetApplicationsWebDataPrivacyOK struct {
	Payload *dynatrace.ApplicationDataPrivacy
}

func (o *GetApplicationsWebDataPrivacyOK) Error() string {
	return fmt.Sprintf("[GET /applications/web/{id}/dataPrivacy][%d] getApplicationsWebDataPrivacyOK  %+v", 200, o.Payload)
}

func (o *GetApplicationsWebDataPrivacyOK) GetPayload() *dynatrace.ApplicationDataPrivacy {
	return o.Payload
}

func (o *GetApplicationsWebDataPrivacyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.ApplicationDataPrivacy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}