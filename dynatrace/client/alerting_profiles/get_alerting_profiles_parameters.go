// Code generated by go-swagger; DO NOT EDIT.

package alerting_profiles

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

// NewGetAlertingProfilesParams creates a new GetAlertingProfilesParams object
// with the default values initialized.
func NewGetAlertingProfilesParams() *GetAlertingProfilesParams {

	return &GetAlertingProfilesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAlertingProfilesParamsWithTimeout creates a new GetAlertingProfilesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAlertingProfilesParamsWithTimeout(timeout time.Duration) *GetAlertingProfilesParams {

	return &GetAlertingProfilesParams{

		timeout: timeout,
	}
}

// NewGetAlertingProfilesParamsWithContext creates a new GetAlertingProfilesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAlertingProfilesParamsWithContext(ctx context.Context) *GetAlertingProfilesParams {

	return &GetAlertingProfilesParams{

		Context: ctx,
	}
}

// NewGetAlertingProfilesParamsWithHTTPClient creates a new GetAlertingProfilesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAlertingProfilesParamsWithHTTPClient(client *http.Client) *GetAlertingProfilesParams {

	return &GetAlertingProfilesParams{
		HTTPClient: client,
	}
}

/*GetAlertingProfilesParams contains all the parameters to send to the API endpoint
for the get alerting profiles operation typically these are written to a http.Request
*/
type GetAlertingProfilesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get alerting profiles params
func (o *GetAlertingProfilesParams) WithTimeout(timeout time.Duration) *GetAlertingProfilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get alerting profiles params
func (o *GetAlertingProfilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get alerting profiles params
func (o *GetAlertingProfilesParams) WithContext(ctx context.Context) *GetAlertingProfilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get alerting profiles params
func (o *GetAlertingProfilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get alerting profiles params
func (o *GetAlertingProfilesParams) WithHTTPClient(client *http.Client) *GetAlertingProfilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get alerting profiles params
func (o *GetAlertingProfilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAlertingProfilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}