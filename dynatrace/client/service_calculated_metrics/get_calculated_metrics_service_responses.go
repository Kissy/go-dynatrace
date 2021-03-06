// Code generated by go-swagger; DO NOT EDIT.

package service_calculated_metrics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// GetCalculatedMetricsServiceReader is a Reader for the GetCalculatedMetricsService structure.
type GetCalculatedMetricsServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCalculatedMetricsServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCalculatedMetricsServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCalculatedMetricsServiceOK creates a GetCalculatedMetricsServiceOK with default headers values
func NewGetCalculatedMetricsServiceOK() *GetCalculatedMetricsServiceOK {
	return &GetCalculatedMetricsServiceOK{}
}

/*GetCalculatedMetricsServiceOK handles this case with default header values.

successful operation
*/
type GetCalculatedMetricsServiceOK struct {
	Payload *dynatrace.CalculatedServiceMetric
}

func (o *GetCalculatedMetricsServiceOK) Error() string {
	return fmt.Sprintf("[GET /calculatedMetrics/service/{metricKey}][%d] getCalculatedMetricsServiceOK  %+v", 200, o.Payload)
}

func (o *GetCalculatedMetricsServiceOK) GetPayload() *dynatrace.CalculatedServiceMetric {
	return o.Payload
}

func (o *GetCalculatedMetricsServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.CalculatedServiceMetric)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
