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

package spoe

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

// NewDeleteSpoeAgentParams creates a new DeleteSpoeAgentParams object
// with the default values initialized.
func NewDeleteSpoeAgentParams() *DeleteSpoeAgentParams {
	var ()
	return &DeleteSpoeAgentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSpoeAgentParamsWithTimeout creates a new DeleteSpoeAgentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSpoeAgentParamsWithTimeout(timeout time.Duration) *DeleteSpoeAgentParams {
	var ()
	return &DeleteSpoeAgentParams{

		timeout: timeout,
	}
}

// NewDeleteSpoeAgentParamsWithContext creates a new DeleteSpoeAgentParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSpoeAgentParamsWithContext(ctx context.Context) *DeleteSpoeAgentParams {
	var ()
	return &DeleteSpoeAgentParams{

		Context: ctx,
	}
}

// NewDeleteSpoeAgentParamsWithHTTPClient creates a new DeleteSpoeAgentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSpoeAgentParamsWithHTTPClient(client *http.Client) *DeleteSpoeAgentParams {
	var ()
	return &DeleteSpoeAgentParams{
		HTTPClient: client,
	}
}

/*DeleteSpoeAgentParams contains all the parameters to send to the API endpoint
for the delete spoe agent operation typically these are written to a http.Request
*/
type DeleteSpoeAgentParams struct {

	/*Name
	  Spoe agent name

	*/
	Name string
	/*Scope
	  Spoe scope

	*/
	Scope string
	/*Spoe
	  Spoe file name

	*/
	Spoe string
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

// WithTimeout adds the timeout to the delete spoe agent params
func (o *DeleteSpoeAgentParams) WithTimeout(timeout time.Duration) *DeleteSpoeAgentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete spoe agent params
func (o *DeleteSpoeAgentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete spoe agent params
func (o *DeleteSpoeAgentParams) WithContext(ctx context.Context) *DeleteSpoeAgentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete spoe agent params
func (o *DeleteSpoeAgentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete spoe agent params
func (o *DeleteSpoeAgentParams) WithHTTPClient(client *http.Client) *DeleteSpoeAgentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete spoe agent params
func (o *DeleteSpoeAgentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete spoe agent params
func (o *DeleteSpoeAgentParams) WithName(name string) *DeleteSpoeAgentParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete spoe agent params
func (o *DeleteSpoeAgentParams) SetName(name string) {
	o.Name = name
}

// WithScope adds the scope to the delete spoe agent params
func (o *DeleteSpoeAgentParams) WithScope(scope string) *DeleteSpoeAgentParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the delete spoe agent params
func (o *DeleteSpoeAgentParams) SetScope(scope string) {
	o.Scope = scope
}

// WithSpoe adds the spoe to the delete spoe agent params
func (o *DeleteSpoeAgentParams) WithSpoe(spoe string) *DeleteSpoeAgentParams {
	o.SetSpoe(spoe)
	return o
}

// SetSpoe adds the spoe to the delete spoe agent params
func (o *DeleteSpoeAgentParams) SetSpoe(spoe string) {
	o.Spoe = spoe
}

// WithTransactionID adds the transactionID to the delete spoe agent params
func (o *DeleteSpoeAgentParams) WithTransactionID(transactionID *string) *DeleteSpoeAgentParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the delete spoe agent params
func (o *DeleteSpoeAgentParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the delete spoe agent params
func (o *DeleteSpoeAgentParams) WithVersion(version *int64) *DeleteSpoeAgentParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete spoe agent params
func (o *DeleteSpoeAgentParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSpoeAgentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// query param scope
	qrScope := o.Scope
	qScope := qrScope
	if qScope != "" {
		if err := r.SetQueryParam("scope", qScope); err != nil {
			return err
		}
	}

	// query param spoe
	qrSpoe := o.Spoe
	qSpoe := qrSpoe
	if qSpoe != "" {
		if err := r.SetQueryParam("spoe", qSpoe); err != nil {
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
