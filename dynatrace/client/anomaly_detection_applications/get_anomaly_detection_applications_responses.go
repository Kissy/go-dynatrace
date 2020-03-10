// Code generated by go-swagger; DO NOT EDIT.

package anomaly_detection_applications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// GetAnomalyDetectionApplicationsReader is a Reader for the GetAnomalyDetectionApplications structure.
type GetAnomalyDetectionApplicationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAnomalyDetectionApplicationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAnomalyDetectionApplicationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAnomalyDetectionApplicationsOK creates a GetAnomalyDetectionApplicationsOK with default headers values
func NewGetAnomalyDetectionApplicationsOK() *GetAnomalyDetectionApplicationsOK {
	return &GetAnomalyDetectionApplicationsOK{}
}

/*GetAnomalyDetectionApplicationsOK handles this case with default header values.

successful operation
*/
type GetAnomalyDetectionApplicationsOK struct {
	Payload *dynatrace.ApplicationAnomalyDetectionConfig
}

func (o *GetAnomalyDetectionApplicationsOK) Error() string {
	return fmt.Sprintf("[GET /anomalyDetection/applications][%d] getAnomalyDetectionApplicationsOK  %+v", 200, o.Payload)
}

func (o *GetAnomalyDetectionApplicationsOK) GetPayload() *dynatrace.ApplicationAnomalyDetectionConfig {
	return o.Payload
}

func (o *GetAnomalyDetectionApplicationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.ApplicationAnomalyDetectionConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}