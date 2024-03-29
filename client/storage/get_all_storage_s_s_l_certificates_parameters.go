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
)

// NewGetAllStorageSSLCertificatesParams creates a new GetAllStorageSSLCertificatesParams object
// with the default values initialized.
func NewGetAllStorageSSLCertificatesParams() *GetAllStorageSSLCertificatesParams {

	return &GetAllStorageSSLCertificatesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllStorageSSLCertificatesParamsWithTimeout creates a new GetAllStorageSSLCertificatesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAllStorageSSLCertificatesParamsWithTimeout(timeout time.Duration) *GetAllStorageSSLCertificatesParams {

	return &GetAllStorageSSLCertificatesParams{

		timeout: timeout,
	}
}

// NewGetAllStorageSSLCertificatesParamsWithContext creates a new GetAllStorageSSLCertificatesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAllStorageSSLCertificatesParamsWithContext(ctx context.Context) *GetAllStorageSSLCertificatesParams {

	return &GetAllStorageSSLCertificatesParams{

		Context: ctx,
	}
}

// NewGetAllStorageSSLCertificatesParamsWithHTTPClient creates a new GetAllStorageSSLCertificatesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAllStorageSSLCertificatesParamsWithHTTPClient(client *http.Client) *GetAllStorageSSLCertificatesParams {

	return &GetAllStorageSSLCertificatesParams{
		HTTPClient: client,
	}
}

/*GetAllStorageSSLCertificatesParams contains all the parameters to send to the API endpoint
for the get all storage s s l certificates operation typically these are written to a http.Request
*/
type GetAllStorageSSLCertificatesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get all storage s s l certificates params
func (o *GetAllStorageSSLCertificatesParams) WithTimeout(timeout time.Duration) *GetAllStorageSSLCertificatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all storage s s l certificates params
func (o *GetAllStorageSSLCertificatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all storage s s l certificates params
func (o *GetAllStorageSSLCertificatesParams) WithContext(ctx context.Context) *GetAllStorageSSLCertificatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all storage s s l certificates params
func (o *GetAllStorageSSLCertificatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all storage s s l certificates params
func (o *GetAllStorageSSLCertificatesParams) WithHTTPClient(client *http.Client) *GetAllStorageSSLCertificatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all storage s s l certificates params
func (o *GetAllStorageSSLCertificatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllStorageSSLCertificatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
