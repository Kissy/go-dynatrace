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

// NewValidateManagementZone2Params creates a new ValidateManagementZone2Params object
// with the default values initialized.
func NewValidateManagementZone2Params() *ValidateManagementZone2Params {
	var ()
	return &ValidateManagementZone2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewValidateManagementZone2ParamsWithTimeout creates a new ValidateManagementZone2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewValidateManagementZone2ParamsWithTimeout(timeout time.Duration) *ValidateManagementZone2Params {
	var ()
	return &ValidateManagementZone2Params{

		timeout: timeout,
	}
}

// NewValidateManagementZone2ParamsWithContext creates a new ValidateManagementZone2Params object
// with the default values initialized, and the ability to set a context for a request
func NewValidateManagementZone2ParamsWithContext(ctx context.Context) *ValidateManagementZone2Params {
	var ()
	return &ValidateManagementZone2Params{

		Context: ctx,
	}
}

// NewValidateManagementZone2ParamsWithHTTPClient creates a new ValidateManagementZone2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewValidateManagementZone2ParamsWithHTTPClient(client *http.Client) *ValidateManagementZone2Params {
	var ()
	return &ValidateManagementZone2Params{
		HTTPClient: client,
	}
}

/*ValidateManagementZone2Params contains all the parameters to send to the API endpoint
for the validate management zone 2 operation typically these are written to a http.Request
*/
type ValidateManagementZone2Params struct {

	/*Body
	  The JSON body of the request, containing the management zone parameters to validate.

	*/
	Body *dynatrace.ManagementZone

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the validate management zone 2 params
func (o *ValidateManagementZone2Params) WithTimeout(timeout time.Duration) *ValidateManagementZone2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate management zone 2 params
func (o *ValidateManagementZone2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate management zone 2 params
func (o *ValidateManagementZone2Params) WithContext(ctx context.Context) *ValidateManagementZone2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate management zone 2 params
func (o *ValidateManagementZone2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate management zone 2 params
func (o *ValidateManagementZone2Params) WithHTTPClient(client *http.Client) *ValidateManagementZone2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate management zone 2 params
func (o *ValidateManagementZone2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the validate management zone 2 params
func (o *ValidateManagementZone2Params) WithBody(body *dynatrace.ManagementZone) *ValidateManagementZone2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the validate management zone 2 params
func (o *ValidateManagementZone2Params) SetBody(body *dynatrace.ManagementZone) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateManagementZone2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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