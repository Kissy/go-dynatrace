// Code generated by go-swagger; DO NOT EDIT.

package a_w_s_credentials_configuration

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

// NewValidateUpdateAwsCredentialParams creates a new ValidateUpdateAwsCredentialParams object
// with the default values initialized.
func NewValidateUpdateAwsCredentialParams() *ValidateUpdateAwsCredentialParams {
	var ()
	return &ValidateUpdateAwsCredentialParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewValidateUpdateAwsCredentialParamsWithTimeout creates a new ValidateUpdateAwsCredentialParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewValidateUpdateAwsCredentialParamsWithTimeout(timeout time.Duration) *ValidateUpdateAwsCredentialParams {
	var ()
	return &ValidateUpdateAwsCredentialParams{

		timeout: timeout,
	}
}

// NewValidateUpdateAwsCredentialParamsWithContext creates a new ValidateUpdateAwsCredentialParams object
// with the default values initialized, and the ability to set a context for a request
func NewValidateUpdateAwsCredentialParamsWithContext(ctx context.Context) *ValidateUpdateAwsCredentialParams {
	var ()
	return &ValidateUpdateAwsCredentialParams{

		Context: ctx,
	}
}

// NewValidateUpdateAwsCredentialParamsWithHTTPClient creates a new ValidateUpdateAwsCredentialParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewValidateUpdateAwsCredentialParamsWithHTTPClient(client *http.Client) *ValidateUpdateAwsCredentialParams {
	var ()
	return &ValidateUpdateAwsCredentialParams{
		HTTPClient: client,
	}
}

/*ValidateUpdateAwsCredentialParams contains all the parameters to send to the API endpoint
for the validate update aws credential operation typically these are written to a http.Request
*/
type ValidateUpdateAwsCredentialParams struct {

	/*Body
	  The JSON body of the request. Contains the AWS credentials configuration to be validated.

	*/
	Body *dynatrace.AwsCredentialsConfig
	/*ID
	  The ID of the AWS credentials configuration to be validated.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the validate update aws credential params
func (o *ValidateUpdateAwsCredentialParams) WithTimeout(timeout time.Duration) *ValidateUpdateAwsCredentialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate update aws credential params
func (o *ValidateUpdateAwsCredentialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate update aws credential params
func (o *ValidateUpdateAwsCredentialParams) WithContext(ctx context.Context) *ValidateUpdateAwsCredentialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate update aws credential params
func (o *ValidateUpdateAwsCredentialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate update aws credential params
func (o *ValidateUpdateAwsCredentialParams) WithHTTPClient(client *http.Client) *ValidateUpdateAwsCredentialParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate update aws credential params
func (o *ValidateUpdateAwsCredentialParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the validate update aws credential params
func (o *ValidateUpdateAwsCredentialParams) WithBody(body *dynatrace.AwsCredentialsConfig) *ValidateUpdateAwsCredentialParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the validate update aws credential params
func (o *ValidateUpdateAwsCredentialParams) SetBody(body *dynatrace.AwsCredentialsConfig) {
	o.Body = body
}

// WithID adds the id to the validate update aws credential params
func (o *ValidateUpdateAwsCredentialParams) WithID(id string) *ValidateUpdateAwsCredentialParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the validate update aws credential params
func (o *ValidateUpdateAwsCredentialParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateUpdateAwsCredentialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
