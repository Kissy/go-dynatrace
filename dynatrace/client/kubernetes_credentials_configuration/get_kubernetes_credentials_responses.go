// Code generated by go-swagger; DO NOT EDIT.

package kubernetes_credentials_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// GetKubernetesCredentialsReader is a Reader for the GetKubernetesCredentials structure.
type GetKubernetesCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKubernetesCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKubernetesCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetKubernetesCredentialsOK creates a GetKubernetesCredentialsOK with default headers values
func NewGetKubernetesCredentialsOK() *GetKubernetesCredentialsOK {
	return &GetKubernetesCredentialsOK{}
}

/*GetKubernetesCredentialsOK handles this case with default header values.

successful operation
*/
type GetKubernetesCredentialsOK struct {
	Payload *dynatrace.StubList
}

func (o *GetKubernetesCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /kubernetes/credentials][%d] getKubernetesCredentialsOK  %+v", 200, o.Payload)
}

func (o *GetKubernetesCredentialsOK) GetPayload() *dynatrace.StubList {
	return o.Payload
}

func (o *GetKubernetesCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.StubList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
