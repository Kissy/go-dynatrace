// Code generated by go-swagger; DO NOT EDIT.

package service_detection_opaque_and_external_web_request

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// GetServiceDetectionRuleOpaqueAndExternalRequestsReader is a Reader for the GetServiceDetectionRuleOpaqueAndExternalRequests structure.
type GetServiceDetectionRuleOpaqueAndExternalRequestsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServiceDetectionRuleOpaqueAndExternalRequestsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServiceDetectionRuleOpaqueAndExternalRequestsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetServiceDetectionRuleOpaqueAndExternalRequestsOK creates a GetServiceDetectionRuleOpaqueAndExternalRequestsOK with default headers values
func NewGetServiceDetectionRuleOpaqueAndExternalRequestsOK() *GetServiceDetectionRuleOpaqueAndExternalRequestsOK {
	return &GetServiceDetectionRuleOpaqueAndExternalRequestsOK{}
}

/*GetServiceDetectionRuleOpaqueAndExternalRequestsOK handles this case with default header values.

Success. The response contains the ordered list of rules.
*/
type GetServiceDetectionRuleOpaqueAndExternalRequestsOK struct {
	Payload *dynatrace.StubList
}

func (o *GetServiceDetectionRuleOpaqueAndExternalRequestsOK) Error() string {
	return fmt.Sprintf("[GET /service/detectionRules/OPAQUE_AND_EXTERNAL_WEB_REQUEST][%d] getServiceDetectionRuleOpaqueAndExternalRequestsOK  %+v", 200, o.Payload)
}

func (o *GetServiceDetectionRuleOpaqueAndExternalRequestsOK) GetPayload() *dynatrace.StubList {
	return o.Payload
}

func (o *GetServiceDetectionRuleOpaqueAndExternalRequestsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.StubList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
