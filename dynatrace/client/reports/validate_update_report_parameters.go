// Code generated by go-swagger; DO NOT EDIT.

package reports

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

// NewValidateUpdateReportParams creates a new ValidateUpdateReportParams object
// with the default values initialized.
func NewValidateUpdateReportParams() *ValidateUpdateReportParams {
	var ()
	return &ValidateUpdateReportParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewValidateUpdateReportParamsWithTimeout creates a new ValidateUpdateReportParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewValidateUpdateReportParamsWithTimeout(timeout time.Duration) *ValidateUpdateReportParams {
	var ()
	return &ValidateUpdateReportParams{

		timeout: timeout,
	}
}

// NewValidateUpdateReportParamsWithContext creates a new ValidateUpdateReportParams object
// with the default values initialized, and the ability to set a context for a request
func NewValidateUpdateReportParamsWithContext(ctx context.Context) *ValidateUpdateReportParams {
	var ()
	return &ValidateUpdateReportParams{

		Context: ctx,
	}
}

// NewValidateUpdateReportParamsWithHTTPClient creates a new ValidateUpdateReportParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewValidateUpdateReportParamsWithHTTPClient(client *http.Client) *ValidateUpdateReportParams {
	var ()
	return &ValidateUpdateReportParams{
		HTTPClient: client,
	}
}

/*ValidateUpdateReportParams contains all the parameters to send to the API endpoint
for the validate update report operation typically these are written to a http.Request
*/
type ValidateUpdateReportParams struct {

	/*Body
	  The JSON body of the request. Contains the report to be validated.

	*/
	Body *dynatrace.DashboardReport
	/*ID
	 The ID of the report to be validated.

	The ID in the request body must match this ID.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the validate update report params
func (o *ValidateUpdateReportParams) WithTimeout(timeout time.Duration) *ValidateUpdateReportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate update report params
func (o *ValidateUpdateReportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate update report params
func (o *ValidateUpdateReportParams) WithContext(ctx context.Context) *ValidateUpdateReportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate update report params
func (o *ValidateUpdateReportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate update report params
func (o *ValidateUpdateReportParams) WithHTTPClient(client *http.Client) *ValidateUpdateReportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate update report params
func (o *ValidateUpdateReportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the validate update report params
func (o *ValidateUpdateReportParams) WithBody(body *dynatrace.DashboardReport) *ValidateUpdateReportParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the validate update report params
func (o *ValidateUpdateReportParams) SetBody(body *dynatrace.DashboardReport) {
	o.Body = body
}

// WithID adds the id to the validate update report params
func (o *ValidateUpdateReportParams) WithID(id string) *ValidateUpdateReportParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the validate update report params
func (o *ValidateUpdateReportParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateUpdateReportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
