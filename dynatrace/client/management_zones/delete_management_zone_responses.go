// Code generated by go-swagger; DO NOT EDIT.

package management_zones

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteManagementZoneReader is a Reader for the DeleteManagementZone structure.
type DeleteManagementZoneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteManagementZoneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteManagementZoneNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteManagementZoneNoContent creates a DeleteManagementZoneNoContent with default headers values
func NewDeleteManagementZoneNoContent() *DeleteManagementZoneNoContent {
	return &DeleteManagementZoneNoContent{}
}

/*DeleteManagementZoneNoContent handles this case with default header values.

Deleted. Response does not have a body.
*/
type DeleteManagementZoneNoContent struct {
}

func (o *DeleteManagementZoneNoContent) Error() string {
	return fmt.Sprintf("[DELETE /managementZones/{id}][%d] deleteManagementZoneNoContent ", 204)
}

func (o *DeleteManagementZoneNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
