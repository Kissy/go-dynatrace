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
)

// NewDeleteAwsCredentialParams creates a new DeleteAwsCredentialParams object
// with the default values initialized.
func NewDeleteAwsCredentialParams() *DeleteAwsCredentialParams {
	var ()
	return &DeleteAwsCredentialParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAwsCredentialParamsWithTimeout creates a new DeleteAwsCredentialParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAwsCredentialParamsWithTimeout(timeout time.Duration) *DeleteAwsCredentialParams {
	var ()
	return &DeleteAwsCredentialParams{

		timeout: timeout,
	}
}

// NewDeleteAwsCredentialParamsWithContext creates a new DeleteAwsCredentialParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAwsCredentialParamsWithContext(ctx context.Context) *DeleteAwsCredentialParams {
	var ()
	return &DeleteAwsCredentialParams{

		Context: ctx,
	}
}

// NewDeleteAwsCredentialParamsWithHTTPClient creates a new DeleteAwsCredentialParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAwsCredentialParamsWithHTTPClient(client *http.Client) *DeleteAwsCredentialParams {
	var ()
	return &DeleteAwsCredentialParams{
		HTTPClient: client,
	}
}

/*DeleteAwsCredentialParams contains all the parameters to send to the API endpoint
for the delete aws credential operation typically these are written to a http.Request
*/
type DeleteAwsCredentialParams struct {

	/*ID
	  The ID of AWS credentials configuration to be deleted.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete aws credential params
func (o *DeleteAwsCredentialParams) WithTimeout(timeout time.Duration) *DeleteAwsCredentialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete aws credential params
func (o *DeleteAwsCredentialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete aws credential params
func (o *DeleteAwsCredentialParams) WithContext(ctx context.Context) *DeleteAwsCredentialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete aws credential params
func (o *DeleteAwsCredentialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete aws credential params
func (o *DeleteAwsCredentialParams) WithHTTPClient(client *http.Client) *DeleteAwsCredentialParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete aws credential params
func (o *DeleteAwsCredentialParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete aws credential params
func (o *DeleteAwsCredentialParams) WithID(id string) *DeleteAwsCredentialParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete aws credential params
func (o *DeleteAwsCredentialParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAwsCredentialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
