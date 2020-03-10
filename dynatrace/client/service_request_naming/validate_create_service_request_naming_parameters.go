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

// NewValidateCreateServiceRequestNamingParams creates a new ValidateCreateServiceRequestNamingParams object
// with the default values initialized.
func NewValidateCreateServiceRequestNamingParams() *ValidateCreateServiceRequestNamingParams {
	var ()
	return &ValidateCreateServiceRequestNamingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewValidateCreateServiceRequestNamingParamsWithTimeout creates a new ValidateCreateServiceRequestNamingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewValidateCreateServiceRequestNamingParamsWithTimeout(timeout time.Duration) *ValidateCreateServiceRequestNamingParams {
	var ()
	return &ValidateCreateServiceRequestNamingParams{

		timeout: timeout,
	}
}

// NewValidateCreateServiceRequestNamingParamsWithContext creates a new ValidateCreateServiceRequestNamingParams object
// with the default values initialized, and the ability to set a context for a request
func NewValidateCreateServiceRequestNamingParamsWithContext(ctx context.Context) *ValidateCreateServiceRequestNamingParams {
	var ()
	return &ValidateCreateServiceRequestNamingParams{

		Context: ctx,
	}
}

// NewValidateCreateServiceRequestNamingParamsWithHTTPClient creates a new ValidateCreateServiceRequestNamingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewValidateCreateServiceRequestNamingParamsWithHTTPClient(client *http.Client) *ValidateCreateServiceRequestNamingParams {
	var ()
	return &ValidateCreateServiceRequestNamingParams{
		HTTPClient: client,
	}
}

/*ValidateCreateServiceRequestNamingParams contains all the parameters to send to the API endpoint
for the validate create service request naming operation typically these are written to a http.Request
*/
type ValidateCreateServiceRequestNamingParams struct {

	/*Body
	 The JSON body of the request containing parameters of the new request naming rule.

	You must not specify the ID of the rule!

	*/
	Body *dynatrace.RequestNaming

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the validate create service request naming params
func (o *ValidateCreateServiceRequestNamingParams) WithTimeout(timeout time.Duration) *ValidateCreateServiceRequestNamingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate create service request naming params
func (o *ValidateCreateServiceRequestNamingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate create service request naming params
func (o *ValidateCreateServiceRequestNamingParams) WithContext(ctx context.Context) *ValidateCreateServiceRequestNamingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate create service request naming params
func (o *ValidateCreateServiceRequestNamingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate create service request naming params
func (o *ValidateCreateServiceRequestNamingParams) WithHTTPClient(client *http.Client) *ValidateCreateServiceRequestNamingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate create service request naming params
func (o *ValidateCreateServiceRequestNamingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the validate create service request naming params
func (o *ValidateCreateServiceRequestNamingParams) WithBody(body *dynatrace.RequestNaming) *ValidateCreateServiceRequestNamingParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the validate create service request naming params
func (o *ValidateCreateServiceRequestNamingParams) SetBody(body *dynatrace.RequestNaming) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateCreateServiceRequestNamingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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