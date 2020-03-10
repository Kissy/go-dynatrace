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

// GetKubernetesCredentialReader is a Reader for the GetKubernetesCredential structure.
type GetKubernetesCredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKubernetesCredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKubernetesCredentialOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetKubernetesCredentialOK creates a GetKubernetesCredentialOK with default headers values
func NewGetKubernetesCredentialOK() *GetKubernetesCredentialOK {
	return &GetKubernetesCredentialOK{}
}

/*GetKubernetesCredentialOK handles this case with default header values.

successful operation
*/
type GetKubernetesCredentialOK struct {
	Payload *dynatrace.KubernetesCredentials
}

func (o *GetKubernetesCredentialOK) Error() string {
	return fmt.Sprintf("[GET /kubernetes/credentials/{id}][%d] getKubernetesCredentialOK  %+v", 200, o.Payload)
}

func (o *GetKubernetesCredentialOK) GetPayload() *dynatrace.KubernetesCredentials {
	return o.Payload
}

func (o *GetKubernetesCredentialOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.KubernetesCredentials)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}