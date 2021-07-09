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

package maps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewClearRuntimeMapParams creates a new ClearRuntimeMapParams object
// with the default values initialized.
func NewClearRuntimeMapParams() *ClearRuntimeMapParams {
	var (
		forceSyncDefault = bool(false)
	)
	return &ClearRuntimeMapParams{
		ForceSync: &forceSyncDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewClearRuntimeMapParamsWithTimeout creates a new ClearRuntimeMapParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewClearRuntimeMapParamsWithTimeout(timeout time.Duration) *ClearRuntimeMapParams {
	var (
		forceSyncDefault = bool(false)
	)
	return &ClearRuntimeMapParams{
		ForceSync: &forceSyncDefault,

		timeout: timeout,
	}
}

// NewClearRuntimeMapParamsWithContext creates a new ClearRuntimeMapParams object
// with the default values initialized, and the ability to set a context for a request
func NewClearRuntimeMapParamsWithContext(ctx context.Context) *ClearRuntimeMapParams {
	var (
		forceSyncDefault = bool(false)
	)
	return &ClearRuntimeMapParams{
		ForceSync: &forceSyncDefault,

		Context: ctx,
	}
}

// NewClearRuntimeMapParamsWithHTTPClient creates a new ClearRuntimeMapParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewClearRuntimeMapParamsWithHTTPClient(client *http.Client) *ClearRuntimeMapParams {
	var (
		forceSyncDefault = bool(false)
	)
	return &ClearRuntimeMapParams{
		ForceSync:  &forceSyncDefault,
		HTTPClient: client,
	}
}

/*ClearRuntimeMapParams contains all the parameters to send to the API endpoint
for the clear runtime map operation typically these are written to a http.Request
*/
type ClearRuntimeMapParams struct {

	/*ForceDelete
	  If true, deletes file from disk

	*/
	ForceDelete *bool
	/*ForceSync
	  If true, immediately syncs changes to disk

	*/
	ForceSync *bool
	/*Name
	  Map file name

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the clear runtime map params
func (o *ClearRuntimeMapParams) WithTimeout(timeout time.Duration) *ClearRuntimeMapParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the clear runtime map params
func (o *ClearRuntimeMapParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the clear runtime map params
func (o *ClearRuntimeMapParams) WithContext(ctx context.Context) *ClearRuntimeMapParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the clear runtime map params
func (o *ClearRuntimeMapParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the clear runtime map params
func (o *ClearRuntimeMapParams) WithHTTPClient(client *http.Client) *ClearRuntimeMapParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the clear runtime map params
func (o *ClearRuntimeMapParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForceDelete adds the forceDelete to the clear runtime map params
func (o *ClearRuntimeMapParams) WithForceDelete(forceDelete *bool) *ClearRuntimeMapParams {
	o.SetForceDelete(forceDelete)
	return o
}

// SetForceDelete adds the forceDelete to the clear runtime map params
func (o *ClearRuntimeMapParams) SetForceDelete(forceDelete *bool) {
	o.ForceDelete = forceDelete
}

// WithForceSync adds the forceSync to the clear runtime map params
func (o *ClearRuntimeMapParams) WithForceSync(forceSync *bool) *ClearRuntimeMapParams {
	o.SetForceSync(forceSync)
	return o
}

// SetForceSync adds the forceSync to the clear runtime map params
func (o *ClearRuntimeMapParams) SetForceSync(forceSync *bool) {
	o.ForceSync = forceSync
}

// WithName adds the name to the clear runtime map params
func (o *ClearRuntimeMapParams) WithName(name string) *ClearRuntimeMapParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the clear runtime map params
func (o *ClearRuntimeMapParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ClearRuntimeMapParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ForceDelete != nil {

		// query param forceDelete
		var qrForceDelete bool
		if o.ForceDelete != nil {
			qrForceDelete = *o.ForceDelete
		}
		qForceDelete := swag.FormatBool(qrForceDelete)
		if qForceDelete != "" {
			if err := r.SetQueryParam("forceDelete", qForceDelete); err != nil {
				return err
			}
		}

	}

	if o.ForceSync != nil {

		// query param force_sync
		var qrForceSync bool
		if o.ForceSync != nil {
			qrForceSync = *o.ForceSync
		}
		qForceSync := swag.FormatBool(qrForceSync)
		if qForceSync != "" {
			if err := r.SetQueryParam("force_sync", qForceSync); err != nil {
				return err
			}
		}

	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
