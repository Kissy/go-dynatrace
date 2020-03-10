// Code generated by go-swagger; DO NOT EDIT.

package cloud_foundry_credentials_configuration

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

// NewValidateUpdateCloudFoundryCredentialParams creates a new ValidateUpdateCloudFoundryCredentialParams object
// with the default values initialized.
func NewValidateUpdateCloudFoundryCredentialParams() *ValidateUpdateCloudFoundryCredentialParams {
	var ()
	return &ValidateUpdateCloudFoundryCredentialParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewValidateUpdateCloudFoundryCredentialParamsWithTimeout creates a new ValidateUpdateCloudFoundryCredentialParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewValidateUpdateCloudFoundryCredentialParamsWithTimeout(timeout time.Duration) *ValidateUpdateCloudFoundryCredentialParams {
	var ()
	return &ValidateUpdateCloudFoundryCredentialParams{

		timeout: timeout,
	}
}

// NewValidateUpdateCloudFoundryCredentialParamsWithContext creates a new ValidateUpdateCloudFoundryCredentialParams object
// with the default values initialized, and the ability to set a context for a request
func NewValidateUpdateCloudFoundryCredentialParamsWithContext(ctx context.Context) *ValidateUpdateCloudFoundryCredentialParams {
	var ()
	return &ValidateUpdateCloudFoundryCredentialParams{

		Context: ctx,
	}
}

// NewValidateUpdateCloudFoundryCredentialParamsWithHTTPClient creates a new ValidateUpdateCloudFoundryCredentialParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewValidateUpdateCloudFoundryCredentialParamsWithHTTPClient(client *http.Client) *ValidateUpdateCloudFoundryCredentialParams {
	var ()
	return &ValidateUpdateCloudFoundryCredentialParams{
		HTTPClient: client,
	}
}

/*ValidateUpdateCloudFoundryCredentialParams contains all the parameters to send to the API endpoint
for the validate update cloud foundry credential operation typically these are written to a http.Request
*/
type ValidateUpdateCloudFoundryCredentialParams struct {

	/*Body
	  `name` must be unique. `password` can be omitted for updates. See the constraints for the PUT /cloudFoundry/credentials/{id} request.

	*/
	Body *dynatrace.CloudFoundryCredentials
	/*ID
	  The ID of the Cloud Foundry foundation credentials.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the validate update cloud foundry credential params
func (o *ValidateUpdateCloudFoundryCredentialParams) WithTimeout(timeout time.Duration) *ValidateUpdateCloudFoundryCredentialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate update cloud foundry credential params
func (o *ValidateUpdateCloudFoundryCredentialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate update cloud foundry credential params
func (o *ValidateUpdateCloudFoundryCredentialParams) WithContext(ctx context.Context) *ValidateUpdateCloudFoundryCredentialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate update cloud foundry credential params
func (o *ValidateUpdateCloudFoundryCredentialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate update cloud foundry credential params
func (o *ValidateUpdateCloudFoundryCredentialParams) WithHTTPClient(client *http.Client) *ValidateUpdateCloudFoundryCredentialParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate update cloud foundry credential params
func (o *ValidateUpdateCloudFoundryCredentialParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the validate update cloud foundry credential params
func (o *ValidateUpdateCloudFoundryCredentialParams) WithBody(body *dynatrace.CloudFoundryCredentials) *ValidateUpdateCloudFoundryCredentialParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the validate update cloud foundry credential params
func (o *ValidateUpdateCloudFoundryCredentialParams) SetBody(body *dynatrace.CloudFoundryCredentials) {
	o.Body = body
}

// WithID adds the id to the validate update cloud foundry credential params
func (o *ValidateUpdateCloudFoundryCredentialParams) WithID(id string) *ValidateUpdateCloudFoundryCredentialParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the validate update cloud foundry credential params
func (o *ValidateUpdateCloudFoundryCredentialParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateUpdateCloudFoundryCredentialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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