// Code generated by go-swagger; DO NOT EDIT.

package service_request_naming

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

// NewValidateUpdateServiceResourceNamingParams creates a new ValidateUpdateServiceResourceNamingParams object
// with the default values initialized.
func NewValidateUpdateServiceResourceNamingParams() *ValidateUpdateServiceResourceNamingParams {
	var ()
	return &ValidateUpdateServiceResourceNamingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewValidateUpdateServiceResourceNamingParamsWithTimeout creates a new ValidateUpdateServiceResourceNamingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewValidateUpdateServiceResourceNamingParamsWithTimeout(timeout time.Duration) *ValidateUpdateServiceResourceNamingParams {
	var ()
	return &ValidateUpdateServiceResourceNamingParams{

		timeout: timeout,
	}
}

// NewValidateUpdateServiceResourceNamingParamsWithContext creates a new ValidateUpdateServiceResourceNamingParams object
// with the default values initialized, and the ability to set a context for a request
func NewValidateUpdateServiceResourceNamingParamsWithContext(ctx context.Context) *ValidateUpdateServiceResourceNamingParams {
	var ()
	return &ValidateUpdateServiceResourceNamingParams{

		Context: ctx,
	}
}

// NewValidateUpdateServiceResourceNamingParamsWithHTTPClient creates a new ValidateUpdateServiceResourceNamingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewValidateUpdateServiceResourceNamingParamsWithHTTPClient(client *http.Client) *ValidateUpdateServiceResourceNamingParams {
	var ()
	return &ValidateUpdateServiceResourceNamingParams{
		HTTPClient: client,
	}
}

/*ValidateUpdateServiceResourceNamingParams contains all the parameters to send to the API endpoint
for the validate update service resource naming operation typically these are written to a http.Request
*/
type ValidateUpdateServiceResourceNamingParams struct {

	/*Body*/
	Body *dynatrace.ResourceNaming

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the validate update service resource naming params
func (o *ValidateUpdateServiceResourceNamingParams) WithTimeout(timeout time.Duration) *ValidateUpdateServiceResourceNamingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate update service resource naming params
func (o *ValidateUpdateServiceResourceNamingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate update service resource naming params
func (o *ValidateUpdateServiceResourceNamingParams) WithContext(ctx context.Context) *ValidateUpdateServiceResourceNamingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate update service resource naming params
func (o *ValidateUpdateServiceResourceNamingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate update service resource naming params
func (o *ValidateUpdateServiceResourceNamingParams) WithHTTPClient(client *http.Client) *ValidateUpdateServiceResourceNamingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate update service resource naming params
func (o *ValidateUpdateServiceResourceNamingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the validate update service resource naming params
func (o *ValidateUpdateServiceResourceNamingParams) WithBody(body *dynatrace.ResourceNaming) *ValidateUpdateServiceResourceNamingParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the validate update service resource naming params
func (o *ValidateUpdateServiceResourceNamingParams) SetBody(body *dynatrace.ResourceNaming) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateUpdateServiceResourceNamingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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