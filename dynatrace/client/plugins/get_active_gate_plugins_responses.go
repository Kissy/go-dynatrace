// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// GetActiveGatePluginsReader is a Reader for the GetActiveGatePlugins structure.
type GetActiveGatePluginsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetActiveGatePluginsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetActiveGatePluginsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetActiveGatePluginsOK creates a GetActiveGatePluginsOK with default headers values
func NewGetActiveGatePluginsOK() *GetActiveGatePluginsOK {
	return &GetActiveGatePluginsOK{}
}

/*GetActiveGatePluginsOK handles this case with default header values.

Success. The response contains IDs of ActiveGate plugin modules. Use them to configure plugin endpoints.
*/
type GetActiveGatePluginsOK struct {
	Payload *dynatrace.StubList
}

func (o *GetActiveGatePluginsOK) Error() string {
	return fmt.Sprintf("[GET /plugins/activeGatePluginModules][%d] getActiveGatePluginsOK  %+v", 200, o.Payload)
}

func (o *GetActiveGatePluginsOK) GetPayload() *dynatrace.StubList {
	return o.Payload
}

func (o *GetActiveGatePluginsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.StubList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
