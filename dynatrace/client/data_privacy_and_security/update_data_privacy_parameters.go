// Code generated by go-swagger; DO NOT EDIT.

package data_privacy_and_security

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

// NewUpdateDataPrivacyParams creates a new UpdateDataPrivacyParams object
// with the default values initialized.
func NewUpdateDataPrivacyParams() *UpdateDataPrivacyParams {
	var ()
	return &UpdateDataPrivacyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDataPrivacyParamsWithTimeout creates a new UpdateDataPrivacyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateDataPrivacyParamsWithTimeout(timeout time.Duration) *UpdateDataPrivacyParams {
	var ()
	return &UpdateDataPrivacyParams{

		timeout: timeout,
	}
}

// NewUpdateDataPrivacyParamsWithContext creates a new UpdateDataPrivacyParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateDataPrivacyParamsWithContext(ctx context.Context) *UpdateDataPrivacyParams {
	var ()
	return &UpdateDataPrivacyParams{

		Context: ctx,
	}
}

// NewUpdateDataPrivacyParamsWithHTTPClient creates a new UpdateDataPrivacyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateDataPrivacyParamsWithHTTPClient(client *http.Client) *UpdateDataPrivacyParams {
	var ()
	return &UpdateDataPrivacyParams{
		HTTPClient: client,
	}
}

/*UpdateDataPrivacyParams contains all the parameters to send to the API endpoint
for the update data privacy operation typically these are written to a http.Request
*/
type UpdateDataPrivacyParams struct {

	/*Body*/
	Body *dynatrace.DataPrivacyAndSecurity

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update data privacy params
func (o *UpdateDataPrivacyParams) WithTimeout(timeout time.Duration) *UpdateDataPrivacyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update data privacy params
func (o *UpdateDataPrivacyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update data privacy params
func (o *UpdateDataPrivacyParams) WithContext(ctx context.Context) *UpdateDataPrivacyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update data privacy params
func (o *UpdateDataPrivacyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update data privacy params
func (o *UpdateDataPrivacyParams) WithHTTPClient(client *http.Client) *UpdateDataPrivacyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update data privacy params
func (o *UpdateDataPrivacyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update data privacy params
func (o *UpdateDataPrivacyParams) WithBody(body *dynatrace.DataPrivacyAndSecurity) *UpdateDataPrivacyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update data privacy params
func (o *UpdateDataPrivacyParams) SetBody(body *dynatrace.DataPrivacyAndSecurity) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDataPrivacyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
