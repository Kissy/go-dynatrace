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

// NewUpdateReportParams creates a new UpdateReportParams object
// with the default values initialized.
func NewUpdateReportParams() *UpdateReportParams {
	var ()
	return &UpdateReportParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateReportParamsWithTimeout creates a new UpdateReportParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateReportParamsWithTimeout(timeout time.Duration) *UpdateReportParams {
	var ()
	return &UpdateReportParams{

		timeout: timeout,
	}
}

// NewUpdateReportParamsWithContext creates a new UpdateReportParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateReportParamsWithContext(ctx context.Context) *UpdateReportParams {
	var ()
	return &UpdateReportParams{

		Context: ctx,
	}
}

// NewUpdateReportParamsWithHTTPClient creates a new UpdateReportParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateReportParamsWithHTTPClient(client *http.Client) *UpdateReportParams {
	var ()
	return &UpdateReportParams{
		HTTPClient: client,
	}
}

/*UpdateReportParams contains all the parameters to send to the API endpoint
for the update report operation typically these are written to a http.Request
*/
type UpdateReportParams struct {

	/*Body
	  The JSON body of the request. Contains updated parameters of the report.

	*/
	Body *dynatrace.DashboardReport
	/*ID
	 The ID of the report to be updated.

	The ID in the request body must match this ID.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update report params
func (o *UpdateReportParams) WithTimeout(timeout time.Duration) *UpdateReportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update report params
func (o *UpdateReportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update report params
func (o *UpdateReportParams) WithContext(ctx context.Context) *UpdateReportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update report params
func (o *UpdateReportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update report params
func (o *UpdateReportParams) WithHTTPClient(client *http.Client) *UpdateReportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update report params
func (o *UpdateReportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update report params
func (o *UpdateReportParams) WithBody(body *dynatrace.DashboardReport) *UpdateReportParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update report params
func (o *UpdateReportParams) SetBody(body *dynatrace.DashboardReport) {
	o.Body = body
}

// WithID adds the id to the update report params
func (o *UpdateReportParams) WithID(id string) *UpdateReportParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update report params
func (o *UpdateReportParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateReportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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