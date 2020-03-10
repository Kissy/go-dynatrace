// Code generated by go-swagger; DO NOT EDIT.

package application_detection_rules

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

// NewUpdateApplicationDetectionRuleParams creates a new UpdateApplicationDetectionRuleParams object
// with the default values initialized.
func NewUpdateApplicationDetectionRuleParams() *UpdateApplicationDetectionRuleParams {
	var ()
	return &UpdateApplicationDetectionRuleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateApplicationDetectionRuleParamsWithTimeout creates a new UpdateApplicationDetectionRuleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateApplicationDetectionRuleParamsWithTimeout(timeout time.Duration) *UpdateApplicationDetectionRuleParams {
	var ()
	return &UpdateApplicationDetectionRuleParams{

		timeout: timeout,
	}
}

// NewUpdateApplicationDetectionRuleParamsWithContext creates a new UpdateApplicationDetectionRuleParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateApplicationDetectionRuleParamsWithContext(ctx context.Context) *UpdateApplicationDetectionRuleParams {
	var ()
	return &UpdateApplicationDetectionRuleParams{

		Context: ctx,
	}
}

// NewUpdateApplicationDetectionRuleParamsWithHTTPClient creates a new UpdateApplicationDetectionRuleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateApplicationDetectionRuleParamsWithHTTPClient(client *http.Client) *UpdateApplicationDetectionRuleParams {
	var ()
	return &UpdateApplicationDetectionRuleParams{
		HTTPClient: client,
	}
}

/*UpdateApplicationDetectionRuleParams contains all the parameters to send to the API endpoint
for the update application detection rule operation typically these are written to a http.Request
*/
type UpdateApplicationDetectionRuleParams struct {

	/*Body
	 The JSON body of the request. Contains updated parameters of the application detection rule.

	If the **order** parameter is set, the rule is placed to this position.

	*/
	Body *dynatrace.ApplicationDetectionRuleConfig
	/*ID
	 The ID of the application detection rule to be updated.

	If you set the ID in the body as well, it must match this ID.

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update application detection rule params
func (o *UpdateApplicationDetectionRuleParams) WithTimeout(timeout time.Duration) *UpdateApplicationDetectionRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update application detection rule params
func (o *UpdateApplicationDetectionRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update application detection rule params
func (o *UpdateApplicationDetectionRuleParams) WithContext(ctx context.Context) *UpdateApplicationDetectionRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update application detection rule params
func (o *UpdateApplicationDetectionRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update application detection rule params
func (o *UpdateApplicationDetectionRuleParams) WithHTTPClient(client *http.Client) *UpdateApplicationDetectionRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update application detection rule params
func (o *UpdateApplicationDetectionRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update application detection rule params
func (o *UpdateApplicationDetectionRuleParams) WithBody(body *dynatrace.ApplicationDetectionRuleConfig) *UpdateApplicationDetectionRuleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update application detection rule params
func (o *UpdateApplicationDetectionRuleParams) SetBody(body *dynatrace.ApplicationDetectionRuleConfig) {
	o.Body = body
}

// WithID adds the id to the update application detection rule params
func (o *UpdateApplicationDetectionRuleParams) WithID(id strfmt.UUID) *UpdateApplicationDetectionRuleParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update application detection rule params
func (o *UpdateApplicationDetectionRuleParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateApplicationDetectionRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}