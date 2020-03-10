// Code generated by go-swagger; DO NOT EDIT.

package credential_vault

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

// NewGetCredentialsParams creates a new GetCredentialsParams object
// with the default values initialized.
func NewGetCredentialsParams() *GetCredentialsParams {
	var ()
	return &GetCredentialsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCredentialsParamsWithTimeout creates a new GetCredentialsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCredentialsParamsWithTimeout(timeout time.Duration) *GetCredentialsParams {
	var ()
	return &GetCredentialsParams{

		timeout: timeout,
	}
}

// NewGetCredentialsParamsWithContext creates a new GetCredentialsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCredentialsParamsWithContext(ctx context.Context) *GetCredentialsParams {
	var ()
	return &GetCredentialsParams{

		Context: ctx,
	}
}

// NewGetCredentialsParamsWithHTTPClient creates a new GetCredentialsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCredentialsParamsWithHTTPClient(client *http.Client) *GetCredentialsParams {
	var ()
	return &GetCredentialsParams{
		HTTPClient: client,
	}
}

/*GetCredentialsParams contains all the parameters to send to the API endpoint
for the get credentials operation typically these are written to a http.Request
*/
type GetCredentialsParams struct {

	/*Type
	  Filters the result by the specified credentials type.

	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get credentials params
func (o *GetCredentialsParams) WithTimeout(timeout time.Duration) *GetCredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get credentials params
func (o *GetCredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get credentials params
func (o *GetCredentialsParams) WithContext(ctx context.Context) *GetCredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get credentials params
func (o *GetCredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get credentials params
func (o *GetCredentialsParams) WithHTTPClient(client *http.Client) *GetCredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get credentials params
func (o *GetCredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithType adds the typeVar to the get credentials params
func (o *GetCredentialsParams) WithType(typeVar *string) *GetCredentialsParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get credentials params
func (o *GetCredentialsParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetCredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Type != nil {

		// query param type
		var qrType string
		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {
			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}