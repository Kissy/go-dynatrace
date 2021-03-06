// Code generated by go-swagger; DO NOT EDIT.

package log_monitoring_metrics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteCalculatedMetricsLogReader is a Reader for the DeleteCalculatedMetricsLog structure.
type DeleteCalculatedMetricsLogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCalculatedMetricsLogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteCalculatedMetricsLogNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteCalculatedMetricsLogNoContent creates a DeleteCalculatedMetricsLogNoContent with default headers values
func NewDeleteCalculatedMetricsLogNoContent() *DeleteCalculatedMetricsLogNoContent {
	return &DeleteCalculatedMetricsLogNoContent{}
}

/*DeleteCalculatedMetricsLogNoContent handles this case with default header values.

Deleted. Response doesn't have a body.
*/
type DeleteCalculatedMetricsLogNoContent struct {
}

func (o *DeleteCalculatedMetricsLogNoContent) Error() string {
	return fmt.Sprintf("[DELETE /calculatedMetrics/log/{metricKey}][%d] deleteCalculatedMetricsLogNoContent ", 204)
}

func (o *DeleteCalculatedMetricsLogNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
