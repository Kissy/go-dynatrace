// Code generated by go-swagger; DO NOT EDIT.

package log_monitoring_metrics

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

// NewUpdateCalculatedMetricsLogParams creates a new UpdateCalculatedMetricsLogParams object
// with the default values initialized.
func NewUpdateCalculatedMetricsLogParams() *UpdateCalculatedMetricsLogParams {
	var ()
	return &UpdateCalculatedMetricsLogParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCalculatedMetricsLogParamsWithTimeout creates a new UpdateCalculatedMetricsLogParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateCalculatedMetricsLogParamsWithTimeout(timeout time.Duration) *UpdateCalculatedMetricsLogParams {
	var ()
	return &UpdateCalculatedMetricsLogParams{

		timeout: timeout,
	}
}

// NewUpdateCalculatedMetricsLogParamsWithContext creates a new UpdateCalculatedMetricsLogParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateCalculatedMetricsLogParamsWithContext(ctx context.Context) *UpdateCalculatedMetricsLogParams {
	var ()
	return &UpdateCalculatedMetricsLogParams{

		Context: ctx,
	}
}

// NewUpdateCalculatedMetricsLogParamsWithHTTPClient creates a new UpdateCalculatedMetricsLogParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateCalculatedMetricsLogParamsWithHTTPClient(client *http.Client) *UpdateCalculatedMetricsLogParams {
	var ()
	return &UpdateCalculatedMetricsLogParams{
		HTTPClient: client,
	}
}

/*UpdateCalculatedMetricsLogParams contains all the parameters to send to the API endpoint
for the update calculated metrics log operation typically these are written to a http.Request
*/
type UpdateCalculatedMetricsLogParams struct {

	/*Body
	  The JSON body of the request. Contains the definition of the custom log metric.

	*/
	Body *dynatrace.LogMetricConfig
	/*MetricKey
	 The required key of the custom log metric. The key must have the `calc:log.` prefix.

	The key in the body of the request must match this key.

	*/
	MetricKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update calculated metrics log params
func (o *UpdateCalculatedMetricsLogParams) WithTimeout(timeout time.Duration) *UpdateCalculatedMetricsLogParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update calculated metrics log params
func (o *UpdateCalculatedMetricsLogParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update calculated metrics log params
func (o *UpdateCalculatedMetricsLogParams) WithContext(ctx context.Context) *UpdateCalculatedMetricsLogParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update calculated metrics log params
func (o *UpdateCalculatedMetricsLogParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update calculated metrics log params
func (o *UpdateCalculatedMetricsLogParams) WithHTTPClient(client *http.Client) *UpdateCalculatedMetricsLogParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update calculated metrics log params
func (o *UpdateCalculatedMetricsLogParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update calculated metrics log params
func (o *UpdateCalculatedMetricsLogParams) WithBody(body *dynatrace.LogMetricConfig) *UpdateCalculatedMetricsLogParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update calculated metrics log params
func (o *UpdateCalculatedMetricsLogParams) SetBody(body *dynatrace.LogMetricConfig) {
	o.Body = body
}

// WithMetricKey adds the metricKey to the update calculated metrics log params
func (o *UpdateCalculatedMetricsLogParams) WithMetricKey(metricKey string) *UpdateCalculatedMetricsLogParams {
	o.SetMetricKey(metricKey)
	return o
}

// SetMetricKey adds the metricKey to the update calculated metrics log params
func (o *UpdateCalculatedMetricsLogParams) SetMetricKey(metricKey string) {
	o.MetricKey = metricKey
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCalculatedMetricsLogParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param metricKey
	if err := r.SetPathParam("metricKey", o.MetricKey); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
