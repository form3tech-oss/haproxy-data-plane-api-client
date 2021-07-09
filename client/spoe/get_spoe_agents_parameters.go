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

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetSpoeAgentsParams creates a new GetSpoeAgentsParams object
// with the default values initialized.
func NewGetSpoeAgentsParams() *GetSpoeAgentsParams {
	var ()
	return &GetSpoeAgentsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSpoeAgentsParamsWithTimeout creates a new GetSpoeAgentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSpoeAgentsParamsWithTimeout(timeout time.Duration) *GetSpoeAgentsParams {
	var ()
	return &GetSpoeAgentsParams{

		timeout: timeout,
	}
}

// NewGetSpoeAgentsParamsWithContext creates a new GetSpoeAgentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSpoeAgentsParamsWithContext(ctx context.Context) *GetSpoeAgentsParams {
	var ()
	return &GetSpoeAgentsParams{

		Context: ctx,
	}
}

// NewGetSpoeAgentsParamsWithHTTPClient creates a new GetSpoeAgentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSpoeAgentsParamsWithHTTPClient(client *http.Client) *GetSpoeAgentsParams {
	var ()
	return &GetSpoeAgentsParams{
		HTTPClient: client,
	}
}

/*GetSpoeAgentsParams contains all the parameters to send to the API endpoint
for the get spoe agents operation typically these are written to a http.Request
*/
type GetSpoeAgentsParams struct {

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get spoe agents params
func (o *GetSpoeAgentsParams) WithTimeout(timeout time.Duration) *GetSpoeAgentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get spoe agents params
func (o *GetSpoeAgentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get spoe agents params
func (o *GetSpoeAgentsParams) WithContext(ctx context.Context) *GetSpoeAgentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get spoe agents params
func (o *GetSpoeAgentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get spoe agents params
func (o *GetSpoeAgentsParams) WithHTTPClient(client *http.Client) *GetSpoeAgentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get spoe agents params
func (o *GetSpoeAgentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithScope adds the scope to the get spoe agents params
func (o *GetSpoeAgentsParams) WithScope(scope string) *GetSpoeAgentsParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the get spoe agents params
func (o *GetSpoeAgentsParams) SetScope(scope string) {
	o.Scope = scope
}

// WithSpoe adds the spoe to the get spoe agents params
func (o *GetSpoeAgentsParams) WithSpoe(spoe string) *GetSpoeAgentsParams {
	o.SetSpoe(spoe)
	return o
}

// SetSpoe adds the spoe to the get spoe agents params
func (o *GetSpoeAgentsParams) SetSpoe(spoe string) {
	o.Spoe = spoe
}

// WithTransactionID adds the transactionID to the get spoe agents params
func (o *GetSpoeAgentsParams) WithTransactionID(transactionID *string) *GetSpoeAgentsParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the get spoe agents params
func (o *GetSpoeAgentsParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSpoeAgentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
