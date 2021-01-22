// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewCreateStorageSSLCertificateParams creates a new CreateStorageSSLCertificateParams object
// with the default values initialized.
func NewCreateStorageSSLCertificateParams() *CreateStorageSSLCertificateParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &CreateStorageSSLCertificateParams{
		ForceReload: &forceReloadDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateStorageSSLCertificateParamsWithTimeout creates a new CreateStorageSSLCertificateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateStorageSSLCertificateParamsWithTimeout(timeout time.Duration) *CreateStorageSSLCertificateParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &CreateStorageSSLCertificateParams{
		ForceReload: &forceReloadDefault,

		timeout: timeout,
	}
}

// NewCreateStorageSSLCertificateParamsWithContext creates a new CreateStorageSSLCertificateParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateStorageSSLCertificateParamsWithContext(ctx context.Context) *CreateStorageSSLCertificateParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &CreateStorageSSLCertificateParams{
		ForceReload: &forceReloadDefault,

		Context: ctx,
	}
}

// NewCreateStorageSSLCertificateParamsWithHTTPClient creates a new CreateStorageSSLCertificateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateStorageSSLCertificateParamsWithHTTPClient(client *http.Client) *CreateStorageSSLCertificateParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &CreateStorageSSLCertificateParams{
		ForceReload: &forceReloadDefault,
		HTTPClient:  client,
	}
}

/*CreateStorageSSLCertificateParams contains all the parameters to send to the API endpoint
for the create storage s s l certificate operation typically these are written to a http.Request
*/
type CreateStorageSSLCertificateParams struct {

	/*FileUpload
	  The SSL certificate to upload

	*/
	FileUpload runtime.NamedReadCloser
	/*ForceReload
	  If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.

	*/
	ForceReload *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create storage s s l certificate params
func (o *CreateStorageSSLCertificateParams) WithTimeout(timeout time.Duration) *CreateStorageSSLCertificateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create storage s s l certificate params
func (o *CreateStorageSSLCertificateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create storage s s l certificate params
func (o *CreateStorageSSLCertificateParams) WithContext(ctx context.Context) *CreateStorageSSLCertificateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create storage s s l certificate params
func (o *CreateStorageSSLCertificateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create storage s s l certificate params
func (o *CreateStorageSSLCertificateParams) WithHTTPClient(client *http.Client) *CreateStorageSSLCertificateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create storage s s l certificate params
func (o *CreateStorageSSLCertificateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFileUpload adds the fileUpload to the create storage s s l certificate params
func (o *CreateStorageSSLCertificateParams) WithFileUpload(fileUpload runtime.NamedReadCloser) *CreateStorageSSLCertificateParams {
	o.SetFileUpload(fileUpload)
	return o
}

// SetFileUpload adds the fileUpload to the create storage s s l certificate params
func (o *CreateStorageSSLCertificateParams) SetFileUpload(fileUpload runtime.NamedReadCloser) {
	o.FileUpload = fileUpload
}

// WithForceReload adds the forceReload to the create storage s s l certificate params
func (o *CreateStorageSSLCertificateParams) WithForceReload(forceReload *bool) *CreateStorageSSLCertificateParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the create storage s s l certificate params
func (o *CreateStorageSSLCertificateParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WriteToRequest writes these params to a swagger request
func (o *CreateStorageSSLCertificateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FileUpload != nil {

		if o.FileUpload != nil {

			// form file param file_upload
			if err := r.SetFileParam("file_upload", o.FileUpload); err != nil {
				return err
			}

		}

	}

	if o.ForceReload != nil {

		// query param force_reload
		var qrForceReload bool
		if o.ForceReload != nil {
			qrForceReload = *o.ForceReload
		}
		qForceReload := swag.FormatBool(qrForceReload)
		if qForceReload != "" {
			if err := r.SetQueryParam("force_reload", qForceReload); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}