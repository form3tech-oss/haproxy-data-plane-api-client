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

package log_target

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

// NewGetLogTargetParams creates a new GetLogTargetParams object
// with the default values initialized.
func NewGetLogTargetParams() *GetLogTargetParams {
	var ()
	return &GetLogTargetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLogTargetParamsWithTimeout creates a new GetLogTargetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLogTargetParamsWithTimeout(timeout time.Duration) *GetLogTargetParams {
	var ()
	return &GetLogTargetParams{

		timeout: timeout,
	}
}

// NewGetLogTargetParamsWithContext creates a new GetLogTargetParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLogTargetParamsWithContext(ctx context.Context) *GetLogTargetParams {
	var ()
	return &GetLogTargetParams{

		Context: ctx,
	}
}

// NewGetLogTargetParamsWithHTTPClient creates a new GetLogTargetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLogTargetParamsWithHTTPClient(client *http.Client) *GetLogTargetParams {
	var ()
	return &GetLogTargetParams{
		HTTPClient: client,
	}
}

/*GetLogTargetParams contains all the parameters to send to the API endpoint
for the get log target operation typically these are written to a http.Request
*/
type GetLogTargetParams struct {

	/*Index
	  Log Target Index

	*/
	Index int64
	/*ParentName
	  Parent name

	*/
	ParentName string
	/*ParentType
	  Parent type

	*/
	ParentType string
	/*TransactionID
	  ID of the transaction where we want to add the operation. Cannot be used when version is specified.

	*/
	TransactionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get log target params
func (o *GetLogTargetParams) WithTimeout(timeout time.Duration) *GetLogTargetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get log target params
func (o *GetLogTargetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get log target params
func (o *GetLogTargetParams) WithContext(ctx context.Context) *GetLogTargetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get log target params
func (o *GetLogTargetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get log target params
func (o *GetLogTargetParams) WithHTTPClient(client *http.Client) *GetLogTargetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get log target params
func (o *GetLogTargetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIndex adds the index to the get log target params
func (o *GetLogTargetParams) WithIndex(index int64) *GetLogTargetParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the get log target params
func (o *GetLogTargetParams) SetIndex(index int64) {
	o.Index = index
}

// WithParentName adds the parentName to the get log target params
func (o *GetLogTargetParams) WithParentName(parentName string) *GetLogTargetParams {
	o.SetParentName(parentName)
	return o
}

// SetParentName adds the parentName to the get log target params
func (o *GetLogTargetParams) SetParentName(parentName string) {
	o.ParentName = parentName
}

// WithParentType adds the parentType to the get log target params
func (o *GetLogTargetParams) WithParentType(parentType string) *GetLogTargetParams {
	o.SetParentType(parentType)
	return o
}

// SetParentType adds the parentType to the get log target params
func (o *GetLogTargetParams) SetParentType(parentType string) {
	o.ParentType = parentType
}

// WithTransactionID adds the transactionID to the get log target params
func (o *GetLogTargetParams) WithTransactionID(transactionID *string) *GetLogTargetParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the get log target params
func (o *GetLogTargetParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLogTargetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param index
	if err := r.SetPathParam("index", swag.FormatInt64(o.Index)); err != nil {
		return err
	}

	// query param parent_name
	qrParentName := o.ParentName
	qParentName := qrParentName
	if qParentName != "" {
		if err := r.SetQueryParam("parent_name", qParentName); err != nil {
			return err
		}
	}

	// query param parent_type
	qrParentType := o.ParentType
	qParentType := qrParentType
	if qParentType != "" {
		if err := r.SetQueryParam("parent_type", qParentType); err != nil {
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
