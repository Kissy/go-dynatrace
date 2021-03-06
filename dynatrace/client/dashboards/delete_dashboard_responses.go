// Code generated by go-swagger; DO NOT EDIT.

package dashboards

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteDashboardReader is a Reader for the DeleteDashboard structure.
type DeleteDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteDashboardNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteDashboardNoContent creates a DeleteDashboardNoContent with default headers values
func NewDeleteDashboardNoContent() *DeleteDashboardNoContent {
	return &DeleteDashboardNoContent{}
}

/*DeleteDashboardNoContent handles this case with default header values.

Success. The dashboard has been deleted. Response doesn't have a body.
*/
type DeleteDashboardNoContent struct {
}

func (o *DeleteDashboardNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dashboards/{id}][%d] deleteDashboardNoContent ", 204)
}

func (o *DeleteDashboardNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
