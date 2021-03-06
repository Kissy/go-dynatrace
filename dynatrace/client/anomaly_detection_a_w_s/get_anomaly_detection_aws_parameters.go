// Code generated by go-swagger; DO NOT EDIT.

package anomaly_detection_a_w_s

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

// NewGetAnomalyDetectionAwsParams creates a new GetAnomalyDetectionAwsParams object
// with the default values initialized.
func NewGetAnomalyDetectionAwsParams() *GetAnomalyDetectionAwsParams {

	return &GetAnomalyDetectionAwsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAnomalyDetectionAwsParamsWithTimeout creates a new GetAnomalyDetectionAwsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAnomalyDetectionAwsParamsWithTimeout(timeout time.Duration) *GetAnomalyDetectionAwsParams {

	return &GetAnomalyDetectionAwsParams{

		timeout: timeout,
	}
}

// NewGetAnomalyDetectionAwsParamsWithContext creates a new GetAnomalyDetectionAwsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAnomalyDetectionAwsParamsWithContext(ctx context.Context) *GetAnomalyDetectionAwsParams {

	return &GetAnomalyDetectionAwsParams{

		Context: ctx,
	}
}

// NewGetAnomalyDetectionAwsParamsWithHTTPClient creates a new GetAnomalyDetectionAwsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAnomalyDetectionAwsParamsWithHTTPClient(client *http.Client) *GetAnomalyDetectionAwsParams {

	return &GetAnomalyDetectionAwsParams{
		HTTPClient: client,
	}
}

/*GetAnomalyDetectionAwsParams contains all the parameters to send to the API endpoint
for the get anomaly detection aws operation typically these are written to a http.Request
*/
type GetAnomalyDetectionAwsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get anomaly detection aws params
func (o *GetAnomalyDetectionAwsParams) WithTimeout(timeout time.Duration) *GetAnomalyDetectionAwsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get anomaly detection aws params
func (o *GetAnomalyDetectionAwsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get anomaly detection aws params
func (o *GetAnomalyDetectionAwsParams) WithContext(ctx context.Context) *GetAnomalyDetectionAwsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get anomaly detection aws params
func (o *GetAnomalyDetectionAwsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get anomaly detection aws params
func (o *GetAnomalyDetectionAwsParams) WithHTTPClient(client *http.Client) *GetAnomalyDetectionAwsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get anomaly detection aws params
func (o *GetAnomalyDetectionAwsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAnomalyDetectionAwsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
