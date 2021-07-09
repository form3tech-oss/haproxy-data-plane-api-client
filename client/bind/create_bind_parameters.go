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

package bind

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

	"github.com/haproxytech/models/v2"
)

// NewCreateBindParams creates a new CreateBindParams object
// with the default values initialized.
func NewCreateBindParams() *CreateBindParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &CreateBindParams{
		ForceReload: &forceReloadDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateBindParamsWithTimeout creates a new CreateBindParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateBindParamsWithTimeout(timeout time.Duration) *CreateBindParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &CreateBindParams{
		ForceReload: &forceReloadDefault,

		timeout: timeout,
	}
}

// NewCreateBindParamsWithContext creates a new CreateBindParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateBindParamsWithContext(ctx context.Context) *CreateBindParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &CreateBindParams{
		ForceReload: &forceReloadDefault,

		Context: ctx,
	}
}

// NewCreateBindParamsWithHTTPClient creates a new CreateBindParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateBindParamsWithHTTPClient(client *http.Client) *CreateBindParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &CreateBindParams{
		ForceReload: &forceReloadDefault,
		HTTPClient:  client,
	}
}

/*CreateBindParams contains all the parameters to send to the API endpoint
for the create bind operation typically these are written to a http.Request
*/
type CreateBindParams struct {

	/*Data*/
	Data *models.Bind
	/*ForceReload
	  If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.

	*/
	ForceReload *bool
	/*Frontend
	  Parent frontend name

	*/
	Frontend string
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

// WithTimeout adds the timeout to the create bind params
func (o *CreateBindParams) WithTimeout(timeout time.Duration) *CreateBindParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create bind params
func (o *CreateBindParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create bind params
func (o *CreateBindParams) WithContext(ctx context.Context) *CreateBindParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create bind params
func (o *CreateBindParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create bind params
func (o *CreateBindParams) WithHTTPClient(client *http.Client) *CreateBindParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create bind params
func (o *CreateBindParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the create bind params
func (o *CreateBindParams) WithData(data *models.Bind) *CreateBindParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the create bind params
func (o *CreateBindParams) SetData(data *models.Bind) {
	o.Data = data
}

// WithForceReload adds the forceReload to the create bind params
func (o *CreateBindParams) WithForceReload(forceReload *bool) *CreateBindParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the create bind params
func (o *CreateBindParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithFrontend adds the frontend to the create bind params
func (o *CreateBindParams) WithFrontend(frontend string) *CreateBindParams {
	o.SetFrontend(frontend)
	return o
}

// SetFrontend adds the frontend to the create bind params
func (o *CreateBindParams) SetFrontend(frontend string) {
	o.Frontend = frontend
}

// WithTransactionID adds the transactionID to the create bind params
func (o *CreateBindParams) WithTransactionID(transactionID *string) *CreateBindParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the create bind params
func (o *CreateBindParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the create bind params
func (o *CreateBindParams) WithVersion(version *int64) *CreateBindParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the create bind params
func (o *CreateBindParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *CreateBindParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param frontend
	qrFrontend := o.Frontend
	qFrontend := qrFrontend
	if qFrontend != "" {
		if err := r.SetQueryParam("frontend", qFrontend); err != nil {
			return err
		}
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
