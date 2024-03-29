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

package frontend

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

	"github.com/haproxytech/models"
)

// NewReplaceFrontendParams creates a new ReplaceFrontendParams object
// with the default values initialized.
func NewReplaceFrontendParams() *ReplaceFrontendParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &ReplaceFrontendParams{
		ForceReload: &forceReloadDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceFrontendParamsWithTimeout creates a new ReplaceFrontendParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceFrontendParamsWithTimeout(timeout time.Duration) *ReplaceFrontendParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &ReplaceFrontendParams{
		ForceReload: &forceReloadDefault,

		timeout: timeout,
	}
}

// NewReplaceFrontendParamsWithContext creates a new ReplaceFrontendParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceFrontendParamsWithContext(ctx context.Context) *ReplaceFrontendParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &ReplaceFrontendParams{
		ForceReload: &forceReloadDefault,

		Context: ctx,
	}
}

// NewReplaceFrontendParamsWithHTTPClient creates a new ReplaceFrontendParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceFrontendParamsWithHTTPClient(client *http.Client) *ReplaceFrontendParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &ReplaceFrontendParams{
		ForceReload: &forceReloadDefault,
		HTTPClient:  client,
	}
}

/*ReplaceFrontendParams contains all the parameters to send to the API endpoint
for the replace frontend operation typically these are written to a http.Request
*/
type ReplaceFrontendParams struct {

	/*Data*/
	Data *models.Frontend
	/*ForceReload
	  If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.

	*/
	ForceReload *bool
	/*Name
	  Frontend name

	*/
	Name string
	/*TransactionID
	  ID of the transaction where we want to add the operation. Cannot be used when version is specified.

	*/
	TransactionID *string
	/*Version
	  Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it's own version.

	*/
	Version *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the replace frontend params
func (o *ReplaceFrontendParams) WithTimeout(timeout time.Duration) *ReplaceFrontendParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace frontend params
func (o *ReplaceFrontendParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace frontend params
func (o *ReplaceFrontendParams) WithContext(ctx context.Context) *ReplaceFrontendParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace frontend params
func (o *ReplaceFrontendParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace frontend params
func (o *ReplaceFrontendParams) WithHTTPClient(client *http.Client) *ReplaceFrontendParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace frontend params
func (o *ReplaceFrontendParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the replace frontend params
func (o *ReplaceFrontendParams) WithData(data *models.Frontend) *ReplaceFrontendParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the replace frontend params
func (o *ReplaceFrontendParams) SetData(data *models.Frontend) {
	o.Data = data
}

// WithForceReload adds the forceReload to the replace frontend params
func (o *ReplaceFrontendParams) WithForceReload(forceReload *bool) *ReplaceFrontendParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the replace frontend params
func (o *ReplaceFrontendParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithName adds the name to the replace frontend params
func (o *ReplaceFrontendParams) WithName(name string) *ReplaceFrontendParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the replace frontend params
func (o *ReplaceFrontendParams) SetName(name string) {
	o.Name = name
}

// WithTransactionID adds the transactionID to the replace frontend params
func (o *ReplaceFrontendParams) WithTransactionID(transactionID *string) *ReplaceFrontendParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the replace frontend params
func (o *ReplaceFrontendParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the replace frontend params
func (o *ReplaceFrontendParams) WithVersion(version *int64) *ReplaceFrontendParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the replace frontend params
func (o *ReplaceFrontendParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceFrontendParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
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

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if o.TransactionID != nil {

		// query param transaction_id
		var qrTransactionID string
		if o.TransactionID != nil {
			qrTransactionID = *o.TransactionID
		}
		qTransactionID := qrTransactionID
		if qTransactionID != "" {
			if err := r.SetQueryParam("transaction_id", qTransactionID); err != nil {
				return err
			}
		}

	}

	if o.Version != nil {

		// query param version
		var qrVersion int64
		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := swag.FormatInt64(qrVersion)
		if qVersion != "" {
			if err := r.SetQueryParam("version", qVersion); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
