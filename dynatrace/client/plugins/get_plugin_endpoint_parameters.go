// Code generated by go-swagger; DO NOT EDIT.

package plugins

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

// NewGetPluginEndpointParams creates a new GetPluginEndpointParams object
// with the default values initialized.
func NewGetPluginEndpointParams() *GetPluginEndpointParams {
	var ()
	return &GetPluginEndpointParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPluginEndpointParamsWithTimeout creates a new GetPluginEndpointParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPluginEndpointParamsWithTimeout(timeout time.Duration) *GetPluginEndpointParams {
	var ()
	return &GetPluginEndpointParams{

		timeout: timeout,
	}
}

// NewGetPluginEndpointParamsWithContext creates a new GetPluginEndpointParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPluginEndpointParamsWithContext(ctx context.Context) *GetPluginEndpointParams {
	var ()
	return &GetPluginEndpointParams{

		Context: ctx,
	}
}

// NewGetPluginEndpointParamsWithHTTPClient creates a new GetPluginEndpointParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPluginEndpointParamsWithHTTPClient(client *http.Client) *GetPluginEndpointParams {
	var ()
	return &GetPluginEndpointParams{
		HTTPClient: client,
	}
}

/*GetPluginEndpointParams contains all the parameters to send to the API endpoint
for the get plugin endpoint operation typically these are written to a http.Request
*/
type GetPluginEndpointParams struct {

	/*EndpointID
	  The ID of the required endpoint.

	*/
	EndpointID string
	/*ID
	  The ID of the required plugin.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get plugin endpoint params
func (o *GetPluginEndpointParams) WithTimeout(timeout time.Duration) *GetPluginEndpointParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get plugin endpoint params
func (o *GetPluginEndpointParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get plugin endpoint params
func (o *GetPluginEndpointParams) WithContext(ctx context.Context) *GetPluginEndpointParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get plugin endpoint params
func (o *GetPluginEndpointParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get plugin endpoint params
func (o *GetPluginEndpointParams) WithHTTPClient(client *http.Client) *GetPluginEndpointParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get plugin endpoint params
func (o *GetPluginEndpointParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndpointID adds the endpointID to the get plugin endpoint params
func (o *GetPluginEndpointParams) WithEndpointID(endpointID string) *GetPluginEndpointParams {
	o.SetEndpointID(endpointID)
	return o
}

// SetEndpointID adds the endpointId to the get plugin endpoint params
func (o *GetPluginEndpointParams) SetEndpointID(endpointID string) {
	o.EndpointID = endpointID
}

// WithID adds the id to the get plugin endpoint params
func (o *GetPluginEndpointParams) WithID(id string) *GetPluginEndpointParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get plugin endpoint params
func (o *GetPluginEndpointParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetPluginEndpointParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param endpointId
	if err := r.SetPathParam("endpointId", o.EndpointID); err != nil {
		return err
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
