// Code generated by go-swagger; DO NOT EDIT.

package anomaly_detection_applications

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

// NewGetAnomalyDetectionApplicationsParams creates a new GetAnomalyDetectionApplicationsParams object
// with the default values initialized.
func NewGetAnomalyDetectionApplicationsParams() *GetAnomalyDetectionApplicationsParams {

	return &GetAnomalyDetectionApplicationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAnomalyDetectionApplicationsParamsWithTimeout creates a new GetAnomalyDetectionApplicationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAnomalyDetectionApplicationsParamsWithTimeout(timeout time.Duration) *GetAnomalyDetectionApplicationsParams {

	return &GetAnomalyDetectionApplicationsParams{

		timeout: timeout,
	}
}

// NewGetAnomalyDetectionApplicationsParamsWithContext creates a new GetAnomalyDetectionApplicationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAnomalyDetectionApplicationsParamsWithContext(ctx context.Context) *GetAnomalyDetectionApplicationsParams {

	return &GetAnomalyDetectionApplicationsParams{

		Context: ctx,
	}
}

// NewGetAnomalyDetectionApplicationsParamsWithHTTPClient creates a new GetAnomalyDetectionApplicationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAnomalyDetectionApplicationsParamsWithHTTPClient(client *http.Client) *GetAnomalyDetectionApplicationsParams {

	return &GetAnomalyDetectionApplicationsParams{
		HTTPClient: client,
	}
}

/*GetAnomalyDetectionApplicationsParams contains all the parameters to send to the API endpoint
for the get anomaly detection applications operation typically these are written to a http.Request
*/
type GetAnomalyDetectionApplicationsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get anomaly detection applications params
func (o *GetAnomalyDetectionApplicationsParams) WithTimeout(timeout time.Duration) *GetAnomalyDetectionApplicationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get anomaly detection applications params
func (o *GetAnomalyDetectionApplicationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get anomaly detection applications params
func (o *GetAnomalyDetectionApplicationsParams) WithContext(ctx context.Context) *GetAnomalyDetectionApplicationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get anomaly detection applications params
func (o *GetAnomalyDetectionApplicationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get anomaly detection applications params
func (o *GetAnomalyDetectionApplicationsParams) WithHTTPClient(client *http.Client) *GetAnomalyDetectionApplicationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get anomaly detection applications params
func (o *GetAnomalyDetectionApplicationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAnomalyDetectionApplicationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
