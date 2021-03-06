// Code generated by go-swagger; DO NOT EDIT.

package service_detection_full_web_service

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

// NewCreateServiceDetectionRuleFullServiceParams creates a new CreateServiceDetectionRuleFullServiceParams object
// with the default values initialized.
func NewCreateServiceDetectionRuleFullServiceParams() *CreateServiceDetectionRuleFullServiceParams {
	var (
		positionDefault = string("APPEND")
	)
	return &CreateServiceDetectionRuleFullServiceParams{
		Position: &positionDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateServiceDetectionRuleFullServiceParamsWithTimeout creates a new CreateServiceDetectionRuleFullServiceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateServiceDetectionRuleFullServiceParamsWithTimeout(timeout time.Duration) *CreateServiceDetectionRuleFullServiceParams {
	var (
		positionDefault = string("APPEND")
	)
	return &CreateServiceDetectionRuleFullServiceParams{
		Position: &positionDefault,

		timeout: timeout,
	}
}

// NewCreateServiceDetectionRuleFullServiceParamsWithContext creates a new CreateServiceDetectionRuleFullServiceParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateServiceDetectionRuleFullServiceParamsWithContext(ctx context.Context) *CreateServiceDetectionRuleFullServiceParams {
	var (
		positionDefault = string("APPEND")
	)
	return &CreateServiceDetectionRuleFullServiceParams{
		Position: &positionDefault,

		Context: ctx,
	}
}

// NewCreateServiceDetectionRuleFullServiceParamsWithHTTPClient creates a new CreateServiceDetectionRuleFullServiceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateServiceDetectionRuleFullServiceParamsWithHTTPClient(client *http.Client) *CreateServiceDetectionRuleFullServiceParams {
	var (
		positionDefault = string("APPEND")
	)
	return &CreateServiceDetectionRuleFullServiceParams{
		Position:   &positionDefault,
		HTTPClient: client,
	}
}

/*CreateServiceDetectionRuleFullServiceParams contains all the parameters to send to the API endpoint
for the create service detection rule full service operation typically these are written to a http.Request
*/
type CreateServiceDetectionRuleFullServiceParams struct {

	/*Body
	  The JSON body of the request containing parameters of the new service detection rule.

	 You must not specify the ID of the rule!

	The **order** field is ignored in this request. To enforce a particular order use the `PUT /ruleBasedServiceDetection/FULL_WEB_SERVICE/reorder` request.

	*/
	Body *dynatrace.FullWebServiceRule
	/*Position
	 The position of the new rule:

	* `APPEND`: at the end of the rule list.
	* `PREPEND`: on top of the rule list.



	*/
	Position *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create service detection rule full service params
func (o *CreateServiceDetectionRuleFullServiceParams) WithTimeout(timeout time.Duration) *CreateServiceDetectionRuleFullServiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create service detection rule full service params
func (o *CreateServiceDetectionRuleFullServiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create service detection rule full service params
func (o *CreateServiceDetectionRuleFullServiceParams) WithContext(ctx context.Context) *CreateServiceDetectionRuleFullServiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create service detection rule full service params
func (o *CreateServiceDetectionRuleFullServiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create service detection rule full service params
func (o *CreateServiceDetectionRuleFullServiceParams) WithHTTPClient(client *http.Client) *CreateServiceDetectionRuleFullServiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create service detection rule full service params
func (o *CreateServiceDetectionRuleFullServiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create service detection rule full service params
func (o *CreateServiceDetectionRuleFullServiceParams) WithBody(body *dynatrace.FullWebServiceRule) *CreateServiceDetectionRuleFullServiceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create service detection rule full service params
func (o *CreateServiceDetectionRuleFullServiceParams) SetBody(body *dynatrace.FullWebServiceRule) {
	o.Body = body
}

// WithPosition adds the position to the create service detection rule full service params
func (o *CreateServiceDetectionRuleFullServiceParams) WithPosition(position *string) *CreateServiceDetectionRuleFullServiceParams {
	o.SetPosition(position)
	return o
}

// SetPosition adds the position to the create service detection rule full service params
func (o *CreateServiceDetectionRuleFullServiceParams) SetPosition(position *string) {
	o.Position = position
}

// WriteToRequest writes these params to a swagger request
func (o *CreateServiceDetectionRuleFullServiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.Position != nil {

		// query param position
		var qrPosition string
		if o.Position != nil {
			qrPosition = *o.Position
		}
		qPosition := qrPosition
		if qPosition != "" {
			if err := r.SetQueryParam("position", qPosition); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
