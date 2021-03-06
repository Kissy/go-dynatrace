// Code generated by go-swagger; DO NOT EDIT.

package mobile_deobfuscation_and_symbolication

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

// NewUpdateSymfilePinningParams creates a new UpdateSymfilePinningParams object
// with the default values initialized.
func NewUpdateSymfilePinningParams() *UpdateSymfilePinningParams {
	var ()
	return &UpdateSymfilePinningParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSymfilePinningParamsWithTimeout creates a new UpdateSymfilePinningParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateSymfilePinningParamsWithTimeout(timeout time.Duration) *UpdateSymfilePinningParams {
	var ()
	return &UpdateSymfilePinningParams{

		timeout: timeout,
	}
}

// NewUpdateSymfilePinningParamsWithContext creates a new UpdateSymfilePinningParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateSymfilePinningParamsWithContext(ctx context.Context) *UpdateSymfilePinningParams {
	var ()
	return &UpdateSymfilePinningParams{

		Context: ctx,
	}
}

// NewUpdateSymfilePinningParamsWithHTTPClient creates a new UpdateSymfilePinningParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateSymfilePinningParamsWithHTTPClient(client *http.Client) *UpdateSymfilePinningParams {
	var ()
	return &UpdateSymfilePinningParams{
		HTTPClient: client,
	}
}

/*UpdateSymfilePinningParams contains all the parameters to send to the API endpoint
for the update symfile pinning operation typically these are written to a http.Request
*/
type UpdateSymfilePinningParams struct {

	/*ApplicationID
	  The application id used in Dynatrace of the app which should be (un)pinned

	*/
	ApplicationID strfmt.UUID
	/*Body
	  JSON body of the request, containing the pinning status to set.

	*/
	Body *dynatrace.SymbolFilePinning
	/*Os
	  The operating system for which the files will be (un)pinned

	*/
	Os string
	/*PackageName
	  The CFBundleIdentifier (iOS) or the package name (Android) which identifies the app associated with the files to be (un)pinned

	*/
	PackageName string
	/*VersionCode
	  The version code (Android) / CFBundleVersion (iOS) the file belongs to

	*/
	VersionCode string
	/*VersionName
	  The version name (Android) / CFBundleShortVersionString (iOS) the file belongs to

	*/
	VersionName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update symfile pinning params
func (o *UpdateSymfilePinningParams) WithTimeout(timeout time.Duration) *UpdateSymfilePinningParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update symfile pinning params
func (o *UpdateSymfilePinningParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update symfile pinning params
func (o *UpdateSymfilePinningParams) WithContext(ctx context.Context) *UpdateSymfilePinningParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update symfile pinning params
func (o *UpdateSymfilePinningParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update symfile pinning params
func (o *UpdateSymfilePinningParams) WithHTTPClient(client *http.Client) *UpdateSymfilePinningParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update symfile pinning params
func (o *UpdateSymfilePinningParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationID adds the applicationID to the update symfile pinning params
func (o *UpdateSymfilePinningParams) WithApplicationID(applicationID strfmt.UUID) *UpdateSymfilePinningParams {
	o.SetApplicationID(applicationID)
	return o
}

// SetApplicationID adds the applicationId to the update symfile pinning params
func (o *UpdateSymfilePinningParams) SetApplicationID(applicationID strfmt.UUID) {
	o.ApplicationID = applicationID
}

// WithBody adds the body to the update symfile pinning params
func (o *UpdateSymfilePinningParams) WithBody(body *dynatrace.SymbolFilePinning) *UpdateSymfilePinningParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update symfile pinning params
func (o *UpdateSymfilePinningParams) SetBody(body *dynatrace.SymbolFilePinning) {
	o.Body = body
}

// WithOs adds the os to the update symfile pinning params
func (o *UpdateSymfilePinningParams) WithOs(os string) *UpdateSymfilePinningParams {
	o.SetOs(os)
	return o
}

// SetOs adds the os to the update symfile pinning params
func (o *UpdateSymfilePinningParams) SetOs(os string) {
	o.Os = os
}

// WithPackageName adds the packageName to the update symfile pinning params
func (o *UpdateSymfilePinningParams) WithPackageName(packageName string) *UpdateSymfilePinningParams {
	o.SetPackageName(packageName)
	return o
}

// SetPackageName adds the packageName to the update symfile pinning params
func (o *UpdateSymfilePinningParams) SetPackageName(packageName string) {
	o.PackageName = packageName
}

// WithVersionCode adds the versionCode to the update symfile pinning params
func (o *UpdateSymfilePinningParams) WithVersionCode(versionCode string) *UpdateSymfilePinningParams {
	o.SetVersionCode(versionCode)
	return o
}

// SetVersionCode adds the versionCode to the update symfile pinning params
func (o *UpdateSymfilePinningParams) SetVersionCode(versionCode string) {
	o.VersionCode = versionCode
}

// WithVersionName adds the versionName to the update symfile pinning params
func (o *UpdateSymfilePinningParams) WithVersionName(versionName string) *UpdateSymfilePinningParams {
	o.SetVersionName(versionName)
	return o
}

// SetVersionName adds the versionName to the update symfile pinning params
func (o *UpdateSymfilePinningParams) SetVersionName(versionName string) {
	o.VersionName = versionName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSymfilePinningParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param applicationId
	if err := r.SetPathParam("applicationId", o.ApplicationID.String()); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param os
	if err := r.SetPathParam("os", o.Os); err != nil {
		return err
	}

	// path param packageName
	if err := r.SetPathParam("packageName", o.PackageName); err != nil {
		return err
	}

	// path param versionCode
	if err := r.SetPathParam("versionCode", o.VersionCode); err != nil {
		return err
	}

	// path param versionName
	if err := r.SetPathParam("versionName", o.VersionName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
