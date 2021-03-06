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

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// NewUpdatePluginEndpointParams creates a new UpdatePluginEndpointParams object
// with the default values initialized.
func NewUpdatePluginEndpointParams() *UpdatePluginEndpointParams {
	var ()
	return &UpdatePluginEndpointParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePluginEndpointParamsWithTimeout creates a new UpdatePluginEndpointParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdatePluginEndpointParamsWithTimeout(timeout time.Duration) *UpdatePluginEndpointParams {
	var ()
	return &UpdatePluginEndpointParams{

		timeout: timeout,
	}
}

// NewUpdatePluginEndpointParamsWithContext creates a new UpdatePluginEndpointParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdatePluginEndpointParamsWithContext(ctx context.Context) *UpdatePluginEndpointParams {
	var ()
	return &UpdatePluginEndpointParams{

		Context: ctx,
	}
}

// NewUpdatePluginEndpointParamsWithHTTPClient creates a new UpdatePluginEndpointParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdatePluginEndpointParamsWithHTTPClient(client *http.Client) *UpdatePluginEndpointParams {
	var ()
	return &UpdatePluginEndpointParams{
		HTTPClient: client,
	}
}

/*UpdatePluginEndpointParams contains all the parameters to send to the API endpoint
for the update plugin endpoint operation typically these are written to a http.Request
*/
type UpdatePluginEndpointParams struct {

	/*Body
	  The JSON body of the request. Contains updated parameters of the plugin endpoint.

	*/
	Body *dynatrace.RemotePluginEndpoint
	/*EndpointID
	  The ID of the endpoint to be updated.

	If you set the endpoint ID in the body as well, it must match this ID.

	*/
	EndpointID string
	/*ID
	  The ID of the plugin where you want to update an endpoint.

	If you set the plugin ID in the body as well, it must match this ID.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update plugin endpoint params
func (o *UpdatePluginEndpointParams) WithTimeout(timeout time.Duration) *UpdatePluginEndpointParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update plugin endpoint params
func (o *UpdatePluginEndpointParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update plugin endpoint params
func (o *UpdatePluginEndpointParams) WithContext(ctx context.Context) *UpdatePluginEndpointParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update plugin endpoint params
func (o *UpdatePluginEndpointParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update plugin endpoint params
func (o *UpdatePluginEndpointParams) WithHTTPClient(client *http.Client) *UpdatePluginEndpointParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update plugin endpoint params
func (o *UpdatePluginEndpointParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update plugin endpoint params
func (o *UpdatePluginEndpointParams) WithBody(body *dynatrace.RemotePluginEndpoint) *UpdatePluginEndpointParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update plugin endpoint params
func (o *UpdatePluginEndpointParams) SetBody(body *dynatrace.RemotePluginEndpoint) {
	o.Body = body
}

// WithEndpointID adds the endpointID to the update plugin endpoint params
func (o *UpdatePluginEndpointParams) WithEndpointID(endpointID string) *UpdatePluginEndpointParams {
	o.SetEndpointID(endpointID)
	return o
}

// SetEndpointID adds the endpointId to the update plugin endpoint params
func (o *UpdatePluginEndpointParams) SetEndpointID(endpointID string) {
	o.EndpointID = endpointID
}

// WithID adds the id to the update plugin endpoint params
func (o *UpdatePluginEndpointParams) WithID(id string) *UpdatePluginEndpointParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update plugin endpoint params
func (o *UpdatePluginEndpointParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePluginEndpointParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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
