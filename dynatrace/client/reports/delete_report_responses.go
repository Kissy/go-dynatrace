// Code generated by go-swagger; DO NOT EDIT.

package reports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteReportReader is a Reader for the DeleteReport structure.
type DeleteReportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteReportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteReportNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteReportNoContent creates a DeleteReportNoContent with default headers values
func NewDeleteReportNoContent() *DeleteReportNoContent {
	return &DeleteReportNoContent{}
}

/*DeleteReportNoContent handles this case with default header values.

Success. The report has been deleted. Response doesn't have a body.
*/
type DeleteReportNoContent struct {
}

func (o *DeleteReportNoContent) Error() string {
	return fmt.Sprintf("[DELETE /reports/{id}][%d] deleteReportNoContent ", 204)
}

func (o *DeleteReportNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
