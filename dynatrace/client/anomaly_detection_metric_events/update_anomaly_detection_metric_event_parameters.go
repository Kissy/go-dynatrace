// Code generated by go-swagger; DO NOT EDIT.

package anomaly_detection_metric_events

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

// NewUpdateAnomalyDetectionMetricEventParams creates a new UpdateAnomalyDetectionMetricEventParams object
// with the default values initialized.
func NewUpdateAnomalyDetectionMetricEventParams() *UpdateAnomalyDetectionMetricEventParams {
	var ()
	return &UpdateAnomalyDetectionMetricEventParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAnomalyDetectionMetricEventParamsWithTimeout creates a new UpdateAnomalyDetectionMetricEventParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateAnomalyDetectionMetricEventParamsWithTimeout(timeout time.Duration) *UpdateAnomalyDetectionMetricEventParams {
	var ()
	return &UpdateAnomalyDetectionMetricEventParams{

		timeout: timeout,
	}
}

// NewUpdateAnomalyDetectionMetricEventParamsWithContext creates a new UpdateAnomalyDetectionMetricEventParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateAnomalyDetectionMetricEventParamsWithContext(ctx context.Context) *UpdateAnomalyDetectionMetricEventParams {
	var ()
	return &UpdateAnomalyDetectionMetricEventParams{

		Context: ctx,
	}
}

// NewUpdateAnomalyDetectionMetricEventParamsWithHTTPClient creates a new UpdateAnomalyDetectionMetricEventParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateAnomalyDetectionMetricEventParamsWithHTTPClient(client *http.Client) *UpdateAnomalyDetectionMetricEventParams {
	var ()
	return &UpdateAnomalyDetectionMetricEventParams{
		HTTPClient: client,
	}
}

/*UpdateAnomalyDetectionMetricEventParams contains all the parameters to send to the API endpoint
for the update anomaly detection metric event operation typically these are written to a http.Request
*/
type UpdateAnomalyDetectionMetricEventParams struct {

	/*Body
	  The JSON body of the request. Contains updated parameters of the metric event.

	*/
	Body *dynatrace.MetricEvent
	/*ID
	 The ID of the metric event to be updated.

	If you also set the ID in the body, it must match this ID.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update anomaly detection metric event params
func (o *UpdateAnomalyDetectionMetricEventParams) WithTimeout(timeout time.Duration) *UpdateAnomalyDetectionMetricEventParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update anomaly detection metric event params
func (o *UpdateAnomalyDetectionMetricEventParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update anomaly detection metric event params
func (o *UpdateAnomalyDetectionMetricEventParams) WithContext(ctx context.Context) *UpdateAnomalyDetectionMetricEventParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update anomaly detection metric event params
func (o *UpdateAnomalyDetectionMetricEventParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update anomaly detection metric event params
func (o *UpdateAnomalyDetectionMetricEventParams) WithHTTPClient(client *http.Client) *UpdateAnomalyDetectionMetricEventParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update anomaly detection metric event params
func (o *UpdateAnomalyDetectionMetricEventParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update anomaly detection metric event params
func (o *UpdateAnomalyDetectionMetricEventParams) WithBody(body *dynatrace.MetricEvent) *UpdateAnomalyDetectionMetricEventParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update anomaly detection metric event params
func (o *UpdateAnomalyDetectionMetricEventParams) SetBody(body *dynatrace.MetricEvent) {
	o.Body = body
}

// WithID adds the id to the update anomaly detection metric event params
func (o *UpdateAnomalyDetectionMetricEventParams) WithID(id string) *UpdateAnomalyDetectionMetricEventParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update anomaly detection metric event params
func (o *UpdateAnomalyDetectionMetricEventParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAnomalyDetectionMetricEventParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
