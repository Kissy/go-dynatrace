// Code generated by go-swagger; DO NOT EDIT.

package service_request_attributes

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

// NewValidateUpdateServiceRequestAttributeParams creates a new ValidateUpdateServiceRequestAttributeParams object
// with the default values initialized.
func NewValidateUpdateServiceRequestAttributeParams() *ValidateUpdateServiceRequestAttributeParams {
	var ()
	return &ValidateUpdateServiceRequestAttributeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewValidateUpdateServiceRequestAttributeParamsWithTimeout creates a new ValidateUpdateServiceRequestAttributeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewValidateUpdateServiceRequestAttributeParamsWithTimeout(timeout time.Duration) *ValidateUpdateServiceRequestAttributeParams {
	var ()
	return &ValidateUpdateServiceRequestAttributeParams{

		timeout: timeout,
	}
}

// NewValidateUpdateServiceRequestAttributeParamsWithContext creates a new ValidateUpdateServiceRequestAttributeParams object
// with the default values initialized, and the ability to set a context for a request
func NewValidateUpdateServiceRequestAttributeParamsWithContext(ctx context.Context) *ValidateUpdateServiceRequestAttributeParams {
	var ()
	return &ValidateUpdateServiceRequestAttributeParams{

		Context: ctx,
	}
}

// NewValidateUpdateServiceRequestAttributeParamsWithHTTPClient creates a new ValidateUpdateServiceRequestAttributeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewValidateUpdateServiceRequestAttributeParamsWithHTTPClient(client *http.Client) *ValidateUpdateServiceRequestAttributeParams {
	var ()
	return &ValidateUpdateServiceRequestAttributeParams{
		HTTPClient: client,
	}
}

/*ValidateUpdateServiceRequestAttributeParams contains all the parameters to send to the API endpoint
for the validate update service request attribute operation typically these are written to a http.Request
*/
type ValidateUpdateServiceRequestAttributeParams struct {

	/*Body*/
	Body *dynatrace.RequestAttribute
	/*ID
	  The ID of the request attribute to be validated.

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the validate update service request attribute params
func (o *ValidateUpdateServiceRequestAttributeParams) WithTimeout(timeout time.Duration) *ValidateUpdateServiceRequestAttributeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate update service request attribute params
func (o *ValidateUpdateServiceRequestAttributeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate update service request attribute params
func (o *ValidateUpdateServiceRequestAttributeParams) WithContext(ctx context.Context) *ValidateUpdateServiceRequestAttributeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate update service request attribute params
func (o *ValidateUpdateServiceRequestAttributeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate update service request attribute params
func (o *ValidateUpdateServiceRequestAttributeParams) WithHTTPClient(client *http.Client) *ValidateUpdateServiceRequestAttributeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate update service request attribute params
func (o *ValidateUpdateServiceRequestAttributeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the validate update service request attribute params
func (o *ValidateUpdateServiceRequestAttributeParams) WithBody(body *dynatrace.RequestAttribute) *ValidateUpdateServiceRequestAttributeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the validate update service request attribute params
func (o *ValidateUpdateServiceRequestAttributeParams) SetBody(body *dynatrace.RequestAttribute) {
	o.Body = body
}

// WithID adds the id to the validate update service request attribute params
func (o *ValidateUpdateServiceRequestAttributeParams) WithID(id strfmt.UUID) *ValidateUpdateServiceRequestAttributeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the validate update service request attribute params
func (o *ValidateUpdateServiceRequestAttributeParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateUpdateServiceRequestAttributeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
