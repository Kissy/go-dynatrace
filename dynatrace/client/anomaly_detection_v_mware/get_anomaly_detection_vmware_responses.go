// Code generated by go-swagger; DO NOT EDIT.

package anomaly_detection_v_mware

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// GetAnomalyDetectionVmwareReader is a Reader for the GetAnomalyDetectionVmware structure.
type GetAnomalyDetectionVmwareReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAnomalyDetectionVmwareReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAnomalyDetectionVmwareOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAnomalyDetectionVmwareOK creates a GetAnomalyDetectionVmwareOK with default headers values
func NewGetAnomalyDetectionVmwareOK() *GetAnomalyDetectionVmwareOK {
	return &GetAnomalyDetectionVmwareOK{}
}

/*GetAnomalyDetectionVmwareOK handles this case with default header values.

successful operation
*/
type GetAnomalyDetectionVmwareOK struct {
	Payload *dynatrace.VMwareAnomalyDetectionConfig
}

func (o *GetAnomalyDetectionVmwareOK) Error() string {
	return fmt.Sprintf("[GET /anomalyDetection/vmware][%d] getAnomalyDetectionVmwareOK  %+v", 200, o.Payload)
}

func (o *GetAnomalyDetectionVmwareOK) GetPayload() *dynatrace.VMwareAnomalyDetectionConfig {
	return o.Payload
}

func (o *GetAnomalyDetectionVmwareOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.VMwareAnomalyDetectionConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
