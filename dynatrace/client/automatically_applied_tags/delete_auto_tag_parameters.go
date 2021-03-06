// Code generated by go-swagger; DO NOT EDIT.

package automatically_applied_tags

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

// NewDeleteAutoTagParams creates a new DeleteAutoTagParams object
// with the default values initialized.
func NewDeleteAutoTagParams() *DeleteAutoTagParams {
	var ()
	return &DeleteAutoTagParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAutoTagParamsWithTimeout creates a new DeleteAutoTagParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAutoTagParamsWithTimeout(timeout time.Duration) *DeleteAutoTagParams {
	var ()
	return &DeleteAutoTagParams{

		timeout: timeout,
	}
}

// NewDeleteAutoTagParamsWithContext creates a new DeleteAutoTagParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAutoTagParamsWithContext(ctx context.Context) *DeleteAutoTagParams {
	var ()
	return &DeleteAutoTagParams{

		Context: ctx,
	}
}

// NewDeleteAutoTagParamsWithHTTPClient creates a new DeleteAutoTagParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAutoTagParamsWithHTTPClient(client *http.Client) *DeleteAutoTagParams {
	var ()
	return &DeleteAutoTagParams{
		HTTPClient: client,
	}
}

/*DeleteAutoTagParams contains all the parameters to send to the API endpoint
for the delete auto tag operation typically these are written to a http.Request
*/
type DeleteAutoTagParams struct {

	/*ID
	  The ID of the auto-tag to be deleted.

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete auto tag params
func (o *DeleteAutoTagParams) WithTimeout(timeout time.Duration) *DeleteAutoTagParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete auto tag params
func (o *DeleteAutoTagParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete auto tag params
func (o *DeleteAutoTagParams) WithContext(ctx context.Context) *DeleteAutoTagParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete auto tag params
func (o *DeleteAutoTagParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete auto tag params
func (o *DeleteAutoTagParams) WithHTTPClient(client *http.Client) *DeleteAutoTagParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete auto tag params
func (o *DeleteAutoTagParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete auto tag params
func (o *DeleteAutoTagParams) WithID(id strfmt.UUID) *DeleteAutoTagParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete auto tag params
func (o *DeleteAutoTagParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAutoTagParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
