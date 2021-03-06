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

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// NewValidateCreateMaintenanceWindowParams creates a new ValidateCreateMaintenanceWindowParams object
// with the default values initialized.
func NewValidateCreateMaintenanceWindowParams() *ValidateCreateMaintenanceWindowParams {
	var ()
	return &ValidateCreateMaintenanceWindowParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewValidateCreateMaintenanceWindowParamsWithTimeout creates a new ValidateCreateMaintenanceWindowParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewValidateCreateMaintenanceWindowParamsWithTimeout(timeout time.Duration) *ValidateCreateMaintenanceWindowParams {
	var ()
	return &ValidateCreateMaintenanceWindowParams{

		timeout: timeout,
	}
}

// NewValidateCreateMaintenanceWindowParamsWithContext creates a new ValidateCreateMaintenanceWindowParams object
// with the default values initialized, and the ability to set a context for a request
func NewValidateCreateMaintenanceWindowParamsWithContext(ctx context.Context) *ValidateCreateMaintenanceWindowParams {
	var ()
	return &ValidateCreateMaintenanceWindowParams{

		Context: ctx,
	}
}

// NewValidateCreateMaintenanceWindowParamsWithHTTPClient creates a new ValidateCreateMaintenanceWindowParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewValidateCreateMaintenanceWindowParamsWithHTTPClient(client *http.Client) *ValidateCreateMaintenanceWindowParams {
	var ()
	return &ValidateCreateMaintenanceWindowParams{
		HTTPClient: client,
	}
}

/*ValidateCreateMaintenanceWindowParams contains all the parameters to send to the API endpoint
for the validate create maintenance window operation typically these are written to a http.Request
*/
type ValidateCreateMaintenanceWindowParams struct {

	/*Body
	  The JSON body of the request. Contains parameters of the maintenance window be validated.

	*/
	Body *dynatrace.MaintenanceWindow

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the validate create maintenance window params
func (o *ValidateCreateMaintenanceWindowParams) WithTimeout(timeout time.Duration) *ValidateCreateMaintenanceWindowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate create maintenance window params
func (o *ValidateCreateMaintenanceWindowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate create maintenance window params
func (o *ValidateCreateMaintenanceWindowParams) WithContext(ctx context.Context) *ValidateCreateMaintenanceWindowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate create maintenance window params
func (o *ValidateCreateMaintenanceWindowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate create maintenance window params
func (o *ValidateCreateMaintenanceWindowParams) WithHTTPClient(client *http.Client) *ValidateCreateMaintenanceWindowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate create maintenance window params
func (o *ValidateCreateMaintenanceWindowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the validate create maintenance window params
func (o *ValidateCreateMaintenanceWindowParams) WithBody(body *dynatrace.MaintenanceWindow) *ValidateCreateMaintenanceWindowParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the validate create maintenance window params
func (o *ValidateCreateMaintenanceWindowParams) SetBody(body *dynatrace.MaintenanceWindow) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateCreateMaintenanceWindowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
