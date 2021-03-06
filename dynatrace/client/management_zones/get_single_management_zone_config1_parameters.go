// Code generated by go-swagger; DO NOT EDIT.

package management_zones

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetSingleManagementZoneConfig1Params creates a new GetSingleManagementZoneConfig1Params object
// with the default values initialized.
func NewGetSingleManagementZoneConfig1Params() *GetSingleManagementZoneConfig1Params {
	var (
		includeProcessGroupReferencesDefault = bool(false)
	)
	return &GetSingleManagementZoneConfig1Params{
		IncludeProcessGroupReferences: &includeProcessGroupReferencesDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSingleManagementZoneConfig1ParamsWithTimeout creates a new GetSingleManagementZoneConfig1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSingleManagementZoneConfig1ParamsWithTimeout(timeout time.Duration) *GetSingleManagementZoneConfig1Params {
	var (
		includeProcessGroupReferencesDefault = bool(false)
	)
	return &GetSingleManagementZoneConfig1Params{
		IncludeProcessGroupReferences: &includeProcessGroupReferencesDefault,

		timeout: timeout,
	}
}

// NewGetSingleManagementZoneConfig1ParamsWithContext creates a new GetSingleManagementZoneConfig1Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetSingleManagementZoneConfig1ParamsWithContext(ctx context.Context) *GetSingleManagementZoneConfig1Params {
	var (
		includeProcessGroupReferencesDefault = bool(false)
	)
	return &GetSingleManagementZoneConfig1Params{
		IncludeProcessGroupReferences: &includeProcessGroupReferencesDefault,

		Context: ctx,
	}
}

// NewGetSingleManagementZoneConfig1ParamsWithHTTPClient creates a new GetSingleManagementZoneConfig1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSingleManagementZoneConfig1ParamsWithHTTPClient(client *http.Client) *GetSingleManagementZoneConfig1Params {
	var (
		includeProcessGroupReferencesDefault = bool(false)
	)
	return &GetSingleManagementZoneConfig1Params{
		IncludeProcessGroupReferences: &includeProcessGroupReferencesDefault,
		HTTPClient:                    client,
	}
}

/*GetSingleManagementZoneConfig1Params contains all the parameters to send to the API endpoint
for the get single management zone config 1 operation typically these are written to a http.Request
*/
type GetSingleManagementZoneConfig1Params struct {

	/*ID
	  The ID of the required management zone.

	*/
	ID string
	/*IncludeProcessGroupReferences
	 Flag to include process group references to the response.

	Process group references aren't compatible across environments. When this flag is set to false, conditions with process group references will be removed. If that leads to a rule having no conditions, the entire rule will be removed.

	*/
	IncludeProcessGroupReferences *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get single management zone config 1 params
func (o *GetSingleManagementZoneConfig1Params) WithTimeout(timeout time.Duration) *GetSingleManagementZoneConfig1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get single management zone config 1 params
func (o *GetSingleManagementZoneConfig1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get single management zone config 1 params
func (o *GetSingleManagementZoneConfig1Params) WithContext(ctx context.Context) *GetSingleManagementZoneConfig1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get single management zone config 1 params
func (o *GetSingleManagementZoneConfig1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get single management zone config 1 params
func (o *GetSingleManagementZoneConfig1Params) WithHTTPClient(client *http.Client) *GetSingleManagementZoneConfig1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get single management zone config 1 params
func (o *GetSingleManagementZoneConfig1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get single management zone config 1 params
func (o *GetSingleManagementZoneConfig1Params) WithID(id string) *GetSingleManagementZoneConfig1Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the get single management zone config 1 params
func (o *GetSingleManagementZoneConfig1Params) SetID(id string) {
	o.ID = id
}

// WithIncludeProcessGroupReferences adds the includeProcessGroupReferences to the get single management zone config 1 params
func (o *GetSingleManagementZoneConfig1Params) WithIncludeProcessGroupReferences(includeProcessGroupReferences *bool) *GetSingleManagementZoneConfig1Params {
	o.SetIncludeProcessGroupReferences(includeProcessGroupReferences)
	return o
}

// SetIncludeProcessGroupReferences adds the includeProcessGroupReferences to the get single management zone config 1 params
func (o *GetSingleManagementZoneConfig1Params) SetIncludeProcessGroupReferences(includeProcessGroupReferences *bool) {
	o.IncludeProcessGroupReferences = includeProcessGroupReferences
}

// WriteToRequest writes these params to a swagger request
func (o *GetSingleManagementZoneConfig1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.IncludeProcessGroupReferences != nil {

		// query param includeProcessGroupReferences
		var qrIncludeProcessGroupReferences bool
		if o.IncludeProcessGroupReferences != nil {
			qrIncludeProcessGroupReferences = *o.IncludeProcessGroupReferences
		}
		qIncludeProcessGroupReferences := swag.FormatBool(qrIncludeProcessGroupReferences)
		if qIncludeProcessGroupReferences != "" {
			if err := r.SetQueryParam("includeProcessGroupReferences", qIncludeProcessGroupReferences); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
