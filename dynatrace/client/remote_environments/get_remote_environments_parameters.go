// Code generated by go-swagger; DO NOT EDIT.

package remote_environments

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

// NewGetRemoteEnvironmentsParams creates a new GetRemoteEnvironmentsParams object
// with the default values initialized.
func NewGetRemoteEnvironmentsParams() *GetRemoteEnvironmentsParams {

	return &GetRemoteEnvironmentsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRemoteEnvironmentsParamsWithTimeout creates a new GetRemoteEnvironmentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRemoteEnvironmentsParamsWithTimeout(timeout time.Duration) *GetRemoteEnvironmentsParams {

	return &GetRemoteEnvironmentsParams{

		timeout: timeout,
	}
}

// NewGetRemoteEnvironmentsParamsWithContext creates a new GetRemoteEnvironmentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRemoteEnvironmentsParamsWithContext(ctx context.Context) *GetRemoteEnvironmentsParams {

	return &GetRemoteEnvironmentsParams{

		Context: ctx,
	}
}

// NewGetRemoteEnvironmentsParamsWithHTTPClient creates a new GetRemoteEnvironmentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRemoteEnvironmentsParamsWithHTTPClient(client *http.Client) *GetRemoteEnvironmentsParams {

	return &GetRemoteEnvironmentsParams{
		HTTPClient: client,
	}
}

/*GetRemoteEnvironmentsParams contains all the parameters to send to the API endpoint
for the get remote environments operation typically these are written to a http.Request
*/
type GetRemoteEnvironmentsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get remote environments params
func (o *GetRemoteEnvironmentsParams) WithTimeout(timeout time.Duration) *GetRemoteEnvironmentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get remote environments params
func (o *GetRemoteEnvironmentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get remote environments params
func (o *GetRemoteEnvironmentsParams) WithContext(ctx context.Context) *GetRemoteEnvironmentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get remote environments params
func (o *GetRemoteEnvironmentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get remote environments params
func (o *GetRemoteEnvironmentsParams) WithHTTPClient(client *http.Client) *GetRemoteEnvironmentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get remote environments params
func (o *GetRemoteEnvironmentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetRemoteEnvironmentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
