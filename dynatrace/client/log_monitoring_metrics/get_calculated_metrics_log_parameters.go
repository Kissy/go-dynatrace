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
)

// NewGetCalculatedMetricsLogParams creates a new GetCalculatedMetricsLogParams object
// with the default values initialized.
func NewGetCalculatedMetricsLogParams() *GetCalculatedMetricsLogParams {
	var ()
	return &GetCalculatedMetricsLogParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCalculatedMetricsLogParamsWithTimeout creates a new GetCalculatedMetricsLogParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCalculatedMetricsLogParamsWithTimeout(timeout time.Duration) *GetCalculatedMetricsLogParams {
	var ()
	return &GetCalculatedMetricsLogParams{

		timeout: timeout,
	}
}

// NewGetCalculatedMetricsLogParamsWithContext creates a new GetCalculatedMetricsLogParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCalculatedMetricsLogParamsWithContext(ctx context.Context) *GetCalculatedMetricsLogParams {
	var ()
	return &GetCalculatedMetricsLogParams{

		Context: ctx,
	}
}

// NewGetCalculatedMetricsLogParamsWithHTTPClient creates a new GetCalculatedMetricsLogParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCalculatedMetricsLogParamsWithHTTPClient(client *http.Client) *GetCalculatedMetricsLogParams {
	var ()
	return &GetCalculatedMetricsLogParams{
		HTTPClient: client,
	}
}

/*GetCalculatedMetricsLogParams contains all the parameters to send to the API endpoint
for the get calculated metrics log operation typically these are written to a http.Request
*/
type GetCalculatedMetricsLogParams struct {

	/*MetricKey
	  The key of the required custom log metric.

	*/
	MetricKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get calculated metrics log params
func (o *GetCalculatedMetricsLogParams) WithTimeout(timeout time.Duration) *GetCalculatedMetricsLogParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get calculated metrics log params
func (o *GetCalculatedMetricsLogParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get calculated metrics log params
func (o *GetCalculatedMetricsLogParams) WithContext(ctx context.Context) *GetCalculatedMetricsLogParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get calculated metrics log params
func (o *GetCalculatedMetricsLogParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get calculated metrics log params
func (o *GetCalculatedMetricsLogParams) WithHTTPClient(client *http.Client) *GetCalculatedMetricsLogParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get calculated metrics log params
func (o *GetCalculatedMetricsLogParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMetricKey adds the metricKey to the get calculated metrics log params
func (o *GetCalculatedMetricsLogParams) WithMetricKey(metricKey string) *GetCalculatedMetricsLogParams {
	o.SetMetricKey(metricKey)
	return o
}

// SetMetricKey adds the metricKey to the get calculated metrics log params
func (o *GetCalculatedMetricsLogParams) SetMetricKey(metricKey string) {
	o.MetricKey = metricKey
}

// WriteToRequest writes these params to a swagger request
func (o *GetCalculatedMetricsLogParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param metricKey
	if err := r.SetPathParam("metricKey", o.MetricKey); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}