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

// NewValidateCreateReportParams creates a new ValidateCreateReportParams object
// with the default values initialized.
func NewValidateCreateReportParams() *ValidateCreateReportParams {
	var ()
	return &ValidateCreateReportParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewValidateCreateReportParamsWithTimeout creates a new ValidateCreateReportParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewValidateCreateReportParamsWithTimeout(timeout time.Duration) *ValidateCreateReportParams {
	var ()
	return &ValidateCreateReportParams{

		timeout: timeout,
	}
}

// NewValidateCreateReportParamsWithContext creates a new ValidateCreateReportParams object
// with the default values initialized, and the ability to set a context for a request
func NewValidateCreateReportParamsWithContext(ctx context.Context) *ValidateCreateReportParams {
	var ()
	return &ValidateCreateReportParams{

		Context: ctx,
	}
}

// NewValidateCreateReportParamsWithHTTPClient creates a new ValidateCreateReportParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewValidateCreateReportParamsWithHTTPClient(client *http.Client) *ValidateCreateReportParams {
	var ()
	return &ValidateCreateReportParams{
		HTTPClient: client,
	}
}

/*ValidateCreateReportParams contains all the parameters to send to the API endpoint
for the validate create report operation typically these are written to a http.Request
*/
type ValidateCreateReportParams struct {

	/*Body
	  The JSON body of the request. Contains the report to be validated.

	*/
	Body *dynatrace.DashboardReport

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the validate create report params
func (o *ValidateCreateReportParams) WithTimeout(timeout time.Duration) *ValidateCreateReportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate create report params
func (o *ValidateCreateReportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate create report params
func (o *ValidateCreateReportParams) WithContext(ctx context.Context) *ValidateCreateReportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate create report params
func (o *ValidateCreateReportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate create report params
func (o *ValidateCreateReportParams) WithHTTPClient(client *http.Client) *ValidateCreateReportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate create report params
func (o *ValidateCreateReportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the validate create report params
func (o *ValidateCreateReportParams) WithBody(body *dynatrace.DashboardReport) *ValidateCreateReportParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the validate create report params
func (o *ValidateCreateReportParams) SetBody(body *dynatrace.DashboardReport) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateCreateReportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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