// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeletePluginEndpointReader is a Reader for the DeletePluginEndpoint structure.
type DeletePluginEndpointReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePluginEndpointReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeletePluginEndpointNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeletePluginEndpointNoContent creates a DeletePluginEndpointNoContent with default headers values
func NewDeletePluginEndpointNoContent() *DeletePluginEndpointNoContent {
	return &DeletePluginEndpointNoContent{}
}

/*DeletePluginEndpointNoContent handles this case with default header values.

Deleted. Response doesn't have a body.
*/
type DeletePluginEndpointNoContent struct {
}

func (o *DeletePluginEndpointNoContent) Error() string {
	return fmt.Sprintf("[DELETE /plugins/{id}/endpoints/{endpointId}][%d] deletePluginEndpointNoContent ", 204)
}

func (o *DeletePluginEndpointNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}