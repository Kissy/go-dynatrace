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

// NewSubscribeReportParams creates a new SubscribeReportParams object
// with the default values initialized.
func NewSubscribeReportParams() *SubscribeReportParams {
	var ()
	return &SubscribeReportParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSubscribeReportParamsWithTimeout creates a new SubscribeReportParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSubscribeReportParamsWithTimeout(timeout time.Duration) *SubscribeReportParams {
	var ()
	return &SubscribeReportParams{

		timeout: timeout,
	}
}

// NewSubscribeReportParamsWithContext creates a new SubscribeReportParams object
// with the default values initialized, and the ability to set a context for a request
func NewSubscribeReportParamsWithContext(ctx context.Context) *SubscribeReportParams {
	var ()
	return &SubscribeReportParams{

		Context: ctx,
	}
}

// NewSubscribeReportParamsWithHTTPClient creates a new SubscribeReportParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSubscribeReportParamsWithHTTPClient(client *http.Client) *SubscribeReportParams {
	var ()
	return &SubscribeReportParams{
		HTTPClient: client,
	}
}

/*SubscribeReportParams contains all the parameters to send to the API endpoint
for the subscribe report operation typically these are written to a http.Request
*/
type SubscribeReportParams struct {

	/*Body
	  The JSON body of the request. Contains a list of new subscribers.

	*/
	Body *dynatrace.ReportSubscriptions
	/*ID
	  The ID of the report to subscribe to.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the subscribe report params
func (o *SubscribeReportParams) WithTimeout(timeout time.Duration) *SubscribeReportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the subscribe report params
func (o *SubscribeReportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the subscribe report params
func (o *SubscribeReportParams) WithContext(ctx context.Context) *SubscribeReportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the subscribe report params
func (o *SubscribeReportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the subscribe report params
func (o *SubscribeReportParams) WithHTTPClient(client *http.Client) *SubscribeReportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the subscribe report params
func (o *SubscribeReportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the subscribe report params
func (o *SubscribeReportParams) WithBody(body *dynatrace.ReportSubscriptions) *SubscribeReportParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the subscribe report params
func (o *SubscribeReportParams) SetBody(body *dynatrace.ReportSubscriptions) {
	o.Body = body
}

// WithID adds the id to the subscribe report params
func (o *SubscribeReportParams) WithID(id string) *SubscribeReportParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the subscribe report params
func (o *SubscribeReportParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SubscribeReportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
