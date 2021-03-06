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

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// NewValidateCreateManagementZoneParams creates a new ValidateCreateManagementZoneParams object
// with the default values initialized.
func NewValidateCreateManagementZoneParams() *ValidateCreateManagementZoneParams {
	var ()
	return &ValidateCreateManagementZoneParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewValidateCreateManagementZoneParamsWithTimeout creates a new ValidateCreateManagementZoneParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewValidateCreateManagementZoneParamsWithTimeout(timeout time.Duration) *ValidateCreateManagementZoneParams {
	var ()
	return &ValidateCreateManagementZoneParams{

		timeout: timeout,
	}
}

// NewValidateCreateManagementZoneParamsWithContext creates a new ValidateCreateManagementZoneParams object
// with the default values initialized, and the ability to set a context for a request
func NewValidateCreateManagementZoneParamsWithContext(ctx context.Context) *ValidateCreateManagementZoneParams {
	var ()
	return &ValidateCreateManagementZoneParams{

		Context: ctx,
	}
}

// NewValidateCreateManagementZoneParamsWithHTTPClient creates a new ValidateCreateManagementZoneParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewValidateCreateManagementZoneParamsWithHTTPClient(client *http.Client) *ValidateCreateManagementZoneParams {
	var ()
	return &ValidateCreateManagementZoneParams{
		HTTPClient: client,
	}
}

/*ValidateCreateManagementZoneParams contains all the parameters to send to the API endpoint
for the validate create management zone operation typically these are written to a http.Request
*/
type ValidateCreateManagementZoneParams struct {

	/*Body
	  The JSON body of the request, containing the management zone parameters to validate.

	*/
	Body *dynatrace.ManagementZone

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the validate create management zone params
func (o *ValidateCreateManagementZoneParams) WithTimeout(timeout time.Duration) *ValidateCreateManagementZoneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate create management zone params
func (o *ValidateCreateManagementZoneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate create management zone params
func (o *ValidateCreateManagementZoneParams) WithContext(ctx context.Context) *ValidateCreateManagementZoneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate create management zone params
func (o *ValidateCreateManagementZoneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate create management zone params
func (o *ValidateCreateManagementZoneParams) WithHTTPClient(client *http.Client) *ValidateCreateManagementZoneParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate create management zone params
func (o *ValidateCreateManagementZoneParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the validate create management zone params
func (o *ValidateCreateManagementZoneParams) WithBody(body *dynatrace.ManagementZone) *ValidateCreateManagementZoneParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the validate create management zone params
func (o *ValidateCreateManagementZoneParams) SetBody(body *dynatrace.ManagementZone) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateCreateManagementZoneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
