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

package acl_runtime

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

// NewDeleteServicesHaproxyRuntimeACLFileEntriesIDParams creates a new DeleteServicesHaproxyRuntimeACLFileEntriesIDParams object
// with the default values initialized.
func NewDeleteServicesHaproxyRuntimeACLFileEntriesIDParams() *DeleteServicesHaproxyRuntimeACLFileEntriesIDParams {
	var ()
	return &DeleteServicesHaproxyRuntimeACLFileEntriesIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteServicesHaproxyRuntimeACLFileEntriesIDParamsWithTimeout creates a new DeleteServicesHaproxyRuntimeACLFileEntriesIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteServicesHaproxyRuntimeACLFileEntriesIDParamsWithTimeout(timeout time.Duration) *DeleteServicesHaproxyRuntimeACLFileEntriesIDParams {
	var ()
	return &DeleteServicesHaproxyRuntimeACLFileEntriesIDParams{

		timeout: timeout,
	}
}

// NewDeleteServicesHaproxyRuntimeACLFileEntriesIDParamsWithContext creates a new DeleteServicesHaproxyRuntimeACLFileEntriesIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteServicesHaproxyRuntimeACLFileEntriesIDParamsWithContext(ctx context.Context) *DeleteServicesHaproxyRuntimeACLFileEntriesIDParams {
	var ()
	return &DeleteServicesHaproxyRuntimeACLFileEntriesIDParams{

		Context: ctx,
	}
}

// NewDeleteServicesHaproxyRuntimeACLFileEntriesIDParamsWithHTTPClient creates a new DeleteServicesHaproxyRuntimeACLFileEntriesIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteServicesHaproxyRuntimeACLFileEntriesIDParamsWithHTTPClient(client *http.Client) *DeleteServicesHaproxyRuntimeACLFileEntriesIDParams {
	var ()
	return &DeleteServicesHaproxyRuntimeACLFileEntriesIDParams{
		HTTPClient: client,
	}
}

/*DeleteServicesHaproxyRuntimeACLFileEntriesIDParams contains all the parameters to send to the API endpoint
for the delete services haproxy runtime ACL file entries ID operation typically these are written to a http.Request
*/
type DeleteServicesHaproxyRuntimeACLFileEntriesIDParams struct {

	/*ACLID
	  ACL ID

	*/
	ACLID string
	/*ID
	  File entry ID

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete services haproxy runtime ACL file entries ID params
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDParams) WithTimeout(timeout time.Duration) *DeleteServicesHaproxyRuntimeACLFileEntriesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete services haproxy runtime ACL file entries ID params
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete services haproxy runtime ACL file entries ID params
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDParams) WithContext(ctx context.Context) *DeleteServicesHaproxyRuntimeACLFileEntriesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete services haproxy runtime ACL file entries ID params
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete services haproxy runtime ACL file entries ID params
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDParams) WithHTTPClient(client *http.Client) *DeleteServicesHaproxyRuntimeACLFileEntriesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete services haproxy runtime ACL file entries ID params
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithACLID adds the aCLID to the delete services haproxy runtime ACL file entries ID params
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDParams) WithACLID(aCLID string) *DeleteServicesHaproxyRuntimeACLFileEntriesIDParams {
	o.SetACLID(aCLID)
	return o
}

// SetACLID adds the aclId to the delete services haproxy runtime ACL file entries ID params
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDParams) SetACLID(aCLID string) {
	o.ACLID = aCLID
}

// WithID adds the id to the delete services haproxy runtime ACL file entries ID params
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDParams) WithID(id string) *DeleteServicesHaproxyRuntimeACLFileEntriesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete services haproxy runtime ACL file entries ID params
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteServicesHaproxyRuntimeACLFileEntriesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param acl_id
	qrACLID := o.ACLID
	qACLID := qrACLID
	if qACLID != "" {
		if err := r.SetQueryParam("acl_id", qACLID); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
