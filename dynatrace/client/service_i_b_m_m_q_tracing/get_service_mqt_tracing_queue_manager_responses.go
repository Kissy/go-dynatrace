// Code generated by go-swagger; DO NOT EDIT.

package service_i_b_m_m_q_tracing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// GetServiceMqtTracingQueueManagerReader is a Reader for the GetServiceMqtTracingQueueManager structure.
type GetServiceMqtTracingQueueManagerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServiceMqtTracingQueueManagerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServiceMqtTracingQueueManagerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetServiceMqtTracingQueueManagerOK creates a GetServiceMqtTracingQueueManagerOK with default headers values
func NewGetServiceMqtTracingQueueManagerOK() *GetServiceMqtTracingQueueManagerOK {
	return &GetServiceMqtTracingQueueManagerOK{}
}

/*GetServiceMqtTracingQueueManagerOK handles this case with default header values.

successful operation
*/
type GetServiceMqtTracingQueueManagerOK struct {
	Payload *dynatrace.QueueManager
}

func (o *GetServiceMqtTracingQueueManagerOK) Error() string {
	return fmt.Sprintf("[GET /service/ibmMQTracing/queueManager/{name}][%d] getServiceMqtTracingQueueManagerOK  %+v", 200, o.Payload)
}

func (o *GetServiceMqtTracingQueueManagerOK) GetPayload() *dynatrace.QueueManager {
	return o.Payload
}

func (o *GetServiceMqtTracingQueueManagerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.QueueManager)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}