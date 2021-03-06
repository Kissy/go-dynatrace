// Code generated by go-swagger; DO NOT EDIT.

package frequent_issue_detection

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

// NewUpdateFrequentIssueDetectionParams creates a new UpdateFrequentIssueDetectionParams object
// with the default values initialized.
func NewUpdateFrequentIssueDetectionParams() *UpdateFrequentIssueDetectionParams {
	var ()
	return &UpdateFrequentIssueDetectionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateFrequentIssueDetectionParamsWithTimeout creates a new UpdateFrequentIssueDetectionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateFrequentIssueDetectionParamsWithTimeout(timeout time.Duration) *UpdateFrequentIssueDetectionParams {
	var ()
	return &UpdateFrequentIssueDetectionParams{

		timeout: timeout,
	}
}

// NewUpdateFrequentIssueDetectionParamsWithContext creates a new UpdateFrequentIssueDetectionParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateFrequentIssueDetectionParamsWithContext(ctx context.Context) *UpdateFrequentIssueDetectionParams {
	var ()
	return &UpdateFrequentIssueDetectionParams{

		Context: ctx,
	}
}

// NewUpdateFrequentIssueDetectionParamsWithHTTPClient creates a new UpdateFrequentIssueDetectionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateFrequentIssueDetectionParamsWithHTTPClient(client *http.Client) *UpdateFrequentIssueDetectionParams {
	var ()
	return &UpdateFrequentIssueDetectionParams{
		HTTPClient: client,
	}
}

/*UpdateFrequentIssueDetectionParams contains all the parameters to send to the API endpoint
for the update frequent issue detection operation typically these are written to a http.Request
*/
type UpdateFrequentIssueDetectionParams struct {

	/*Body
	  The JSON body of the request, containing parameters of the frequent issue detection configuration

	*/
	Body *dynatrace.FrequentIssueDetectionConfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update frequent issue detection params
func (o *UpdateFrequentIssueDetectionParams) WithTimeout(timeout time.Duration) *UpdateFrequentIssueDetectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update frequent issue detection params
func (o *UpdateFrequentIssueDetectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update frequent issue detection params
func (o *UpdateFrequentIssueDetectionParams) WithContext(ctx context.Context) *UpdateFrequentIssueDetectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update frequent issue detection params
func (o *UpdateFrequentIssueDetectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update frequent issue detection params
func (o *UpdateFrequentIssueDetectionParams) WithHTTPClient(client *http.Client) *UpdateFrequentIssueDetectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update frequent issue detection params
func (o *UpdateFrequentIssueDetectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update frequent issue detection params
func (o *UpdateFrequentIssueDetectionParams) WithBody(body *dynatrace.FrequentIssueDetectionConfig) *UpdateFrequentIssueDetectionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update frequent issue detection params
func (o *UpdateFrequentIssueDetectionParams) SetBody(body *dynatrace.FrequentIssueDetectionConfig) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateFrequentIssueDetectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
