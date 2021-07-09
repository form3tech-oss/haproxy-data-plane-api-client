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

// NewDeleteRuntimeMapEntryParams creates a new DeleteRuntimeMapEntryParams object
// with the default values initialized.
func NewDeleteRuntimeMapEntryParams() *DeleteRuntimeMapEntryParams {
	var (
		forceSyncDefault = bool(false)
	)
	return &DeleteRuntimeMapEntryParams{
		ForceSync: &forceSyncDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRuntimeMapEntryParamsWithTimeout creates a new DeleteRuntimeMapEntryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteRuntimeMapEntryParamsWithTimeout(timeout time.Duration) *DeleteRuntimeMapEntryParams {
	var (
		forceSyncDefault = bool(false)
	)
	return &DeleteRuntimeMapEntryParams{
		ForceSync: &forceSyncDefault,

		timeout: timeout,
	}
}

// NewDeleteRuntimeMapEntryParamsWithContext creates a new DeleteRuntimeMapEntryParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteRuntimeMapEntryParamsWithContext(ctx context.Context) *DeleteRuntimeMapEntryParams {
	var (
		forceSyncDefault = bool(false)
	)
	return &DeleteRuntimeMapEntryParams{
		ForceSync: &forceSyncDefault,

		Context: ctx,
	}
}

// NewDeleteRuntimeMapEntryParamsWithHTTPClient creates a new DeleteRuntimeMapEntryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteRuntimeMapEntryParamsWithHTTPClient(client *http.Client) *DeleteRuntimeMapEntryParams {
	var (
		forceSyncDefault = bool(false)
	)
	return &DeleteRuntimeMapEntryParams{
		ForceSync:  &forceSyncDefault,
		HTTPClient: client,
	}
}

/*DeleteRuntimeMapEntryParams contains all the parameters to send to the API endpoint
for the delete runtime map entry operation typically these are written to a http.Request
*/
type DeleteRuntimeMapEntryParams struct {

	/*ForceSync
	  If true, immediately syncs changes to disk

	*/
	ForceSync *bool
	/*ID
	  Map id

	*/
	ID string
	/*Map
	  Mapfile attribute storage_name

	*/
	Map string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete runtime map entry params
func (o *DeleteRuntimeMapEntryParams) WithTimeout(timeout time.Duration) *DeleteRuntimeMapEntryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete runtime map entry params
func (o *DeleteRuntimeMapEntryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete runtime map entry params
func (o *DeleteRuntimeMapEntryParams) WithContext(ctx context.Context) *DeleteRuntimeMapEntryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete runtime map entry params
func (o *DeleteRuntimeMapEntryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete runtime map entry params
func (o *DeleteRuntimeMapEntryParams) WithHTTPClient(client *http.Client) *DeleteRuntimeMapEntryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete runtime map entry params
func (o *DeleteRuntimeMapEntryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForceSync adds the forceSync to the delete runtime map entry params
func (o *DeleteRuntimeMapEntryParams) WithForceSync(forceSync *bool) *DeleteRuntimeMapEntryParams {
	o.SetForceSync(forceSync)
	return o
}

// SetForceSync adds the forceSync to the delete runtime map entry params
func (o *DeleteRuntimeMapEntryParams) SetForceSync(forceSync *bool) {
	o.ForceSync = forceSync
}

// WithID adds the id to the delete runtime map entry params
func (o *DeleteRuntimeMapEntryParams) WithID(id string) *DeleteRuntimeMapEntryParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete runtime map entry params
func (o *DeleteRuntimeMapEntryParams) SetID(id string) {
	o.ID = id
}

// WithMap adds the mapVar to the delete runtime map entry params
func (o *DeleteRuntimeMapEntryParams) WithMap(mapVar string) *DeleteRuntimeMapEntryParams {
	o.SetMap(mapVar)
	return o
}

// SetMap adds the map to the delete runtime map entry params
func (o *DeleteRuntimeMapEntryParams) SetMap(mapVar string) {
	o.Map = mapVar
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRuntimeMapEntryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// query param map
	qrMap := o.Map
	qMap := qrMap
	if qMap != "" {
		if err := r.SetQueryParam("map", qMap); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
