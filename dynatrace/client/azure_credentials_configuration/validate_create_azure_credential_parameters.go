// Code generated by go-swagger; DO NOT EDIT.

package azure_credentials_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// NewValidateCreateAzureCredentialParams creates a new ValidateCreateAzureCredentialParams object
// with the default values initialized.
func NewValidateCreateAzureCredentialParams() *ValidateCreateAzureCredentialParams {
	var ()
	return &ValidateCreateAzureCredentialParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewValidateCreateAzureCredentialParamsWithTimeout creates a new ValidateCreateAzureCredentialParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewValidateCreateAzureCredentialParamsWithTimeout(timeout time.Duration) *ValidateCreateAzureCredentialParams {
	var ()
	return &ValidateCreateAzureCredentialParams{

		timeout: timeout,
	}
}

// NewValidateCreateAzureCredentialParamsWithContext creates a new ValidateCreateAzureCredentialParams object
// with the default values initialized, and the ability to set a context for a request
func NewValidateCreateAzureCredentialParamsWithContext(ctx context.Context) *ValidateCreateAzureCredentialParams {
	var ()
	return &ValidateCreateAzureCredentialParams{

		Context: ctx,
	}
}

// NewValidateCreateAzureCredentialParamsWithHTTPClient creates a new ValidateCreateAzureCredentialParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewValidateCreateAzureCredentialParamsWithHTTPClient(client *http.Client) *ValidateCreateAzureCredentialParams {
	var ()
	return &ValidateCreateAzureCredentialParams{
		HTTPClient: client,
	}
}

/*ValidateCreateAzureCredentialParams contains all the parameters to send to the API endpoint
for the validate create azure credential operation typically these are written to a http.Request
*/
type ValidateCreateAzureCredentialParams struct {

	/*Body
	  The JSON body of the request. Contains the Azure credentials configuration to be validated.

	*/
	Body *dynatrace.AzureCredentials

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the validate create azure credential params
func (o *ValidateCreateAzureCredentialParams) WithTimeout(timeout time.Duration) *ValidateCreateAzureCredentialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate create azure credential params
func (o *ValidateCreateAzureCredentialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate create azure credential params
func (o *ValidateCreateAzureCredentialParams) WithContext(ctx context.Context) *ValidateCreateAzureCredentialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate create azure credential params
func (o *ValidateCreateAzureCredentialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate create azure credential params
func (o *ValidateCreateAzureCredentialParams) WithHTTPClient(client *http.Client) *ValidateCreateAzureCredentialParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate create azure credential params
func (o *ValidateCreateAzureCredentialParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the validate create azure credential params
func (o *ValidateCreateAzureCredentialParams) WithBody(body *dynatrace.AzureCredentials) *ValidateCreateAzureCredentialParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the validate create azure credential params
func (o *ValidateCreateAzureCredentialParams) SetBody(body *dynatrace.AzureCredentials) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateCreateAzureCredentialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
