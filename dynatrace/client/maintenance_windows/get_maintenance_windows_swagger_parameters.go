// Code generated by go-swagger; DO NOT EDIT.

package maintenance_windows

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

// NewGetMaintenanceWindowsParams creates a new GetMaintenanceWindowsParams object
// with the default values initialized.
func NewGetMaintenanceWindowsParams() *GetMaintenanceWindowsParams {

	return &GetMaintenanceWindowsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMaintenanceWindowsParamsWithTimeout creates a new GetMaintenanceWindowsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMaintenanceWindowsParamsWithTimeout(timeout time.Duration) *GetMaintenanceWindowsParams {

	return &GetMaintenanceWindowsParams{

		timeout: timeout,
	}
}

// NewGetMaintenanceWindowsParamsWithContext creates a new GetMaintenanceWindowsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMaintenanceWindowsParamsWithContext(ctx context.Context) *GetMaintenanceWindowsParams {

	return &GetMaintenanceWindowsParams{

		Context: ctx,
	}
}

// NewGetMaintenanceWindowsParamsWithHTTPClient creates a new GetMaintenanceWindowsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMaintenanceWindowsParamsWithHTTPClient(client *http.Client) *GetMaintenanceWindowsParams {

	return &GetMaintenanceWindowsParams{
		HTTPClient: client,
	}
}

/*GetMaintenanceWindowsParams contains all the parameters to send to the API endpoint
for the get maintenance windows operation typically these are written to a http.Request
*/
type GetMaintenanceWindowsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get maintenance windows params
func (o *GetMaintenanceWindowsParams) WithTimeout(timeout time.Duration) *GetMaintenanceWindowsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get maintenance windows params
func (o *GetMaintenanceWindowsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get maintenance windows params
func (o *GetMaintenanceWindowsParams) WithContext(ctx context.Context) *GetMaintenanceWindowsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get maintenance windows params
func (o *GetMaintenanceWindowsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get maintenance windows params
func (o *GetMaintenanceWindowsParams) WithHTTPClient(client *http.Client) *GetMaintenanceWindowsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get maintenance windows params
func (o *GetMaintenanceWindowsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetMaintenanceWindowsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
