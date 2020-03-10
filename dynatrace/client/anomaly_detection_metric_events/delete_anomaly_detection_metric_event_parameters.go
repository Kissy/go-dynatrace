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
)

// NewDeleteAnomalyDetectionMetricEventParams creates a new DeleteAnomalyDetectionMetricEventParams object
// with the default values initialized.
func NewDeleteAnomalyDetectionMetricEventParams() *DeleteAnomalyDetectionMetricEventParams {
	var ()
	return &DeleteAnomalyDetectionMetricEventParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAnomalyDetectionMetricEventParamsWithTimeout creates a new DeleteAnomalyDetectionMetricEventParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAnomalyDetectionMetricEventParamsWithTimeout(timeout time.Duration) *DeleteAnomalyDetectionMetricEventParams {
	var ()
	return &DeleteAnomalyDetectionMetricEventParams{

		timeout: timeout,
	}
}

// NewDeleteAnomalyDetectionMetricEventParamsWithContext creates a new DeleteAnomalyDetectionMetricEventParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAnomalyDetectionMetricEventParamsWithContext(ctx context.Context) *DeleteAnomalyDetectionMetricEventParams {
	var ()
	return &DeleteAnomalyDetectionMetricEventParams{

		Context: ctx,
	}
}

// NewDeleteAnomalyDetectionMetricEventParamsWithHTTPClient creates a new DeleteAnomalyDetectionMetricEventParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAnomalyDetectionMetricEventParamsWithHTTPClient(client *http.Client) *DeleteAnomalyDetectionMetricEventParams {
	var ()
	return &DeleteAnomalyDetectionMetricEventParams{
		HTTPClient: client,
	}
}

/*DeleteAnomalyDetectionMetricEventParams contains all the parameters to send to the API endpoint
for the delete anomaly detection metric event operation typically these are written to a http.Request
*/
type DeleteAnomalyDetectionMetricEventParams struct {

	/*ID
	  The ID of the metric event to be deleted.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete anomaly detection metric event params
func (o *DeleteAnomalyDetectionMetricEventParams) WithTimeout(timeout time.Duration) *DeleteAnomalyDetectionMetricEventParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete anomaly detection metric event params
func (o *DeleteAnomalyDetectionMetricEventParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete anomaly detection metric event params
func (o *DeleteAnomalyDetectionMetricEventParams) WithContext(ctx context.Context) *DeleteAnomalyDetectionMetricEventParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete anomaly detection metric event params
func (o *DeleteAnomalyDetectionMetricEventParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete anomaly detection metric event params
func (o *DeleteAnomalyDetectionMetricEventParams) WithHTTPClient(client *http.Client) *DeleteAnomalyDetectionMetricEventParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete anomaly detection metric event params
func (o *DeleteAnomalyDetectionMetricEventParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete anomaly detection metric event params
func (o *DeleteAnomalyDetectionMetricEventParams) WithID(id string) *DeleteAnomalyDetectionMetricEventParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete anomaly detection metric event params
func (o *DeleteAnomalyDetectionMetricEventParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAnomalyDetectionMetricEventParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}