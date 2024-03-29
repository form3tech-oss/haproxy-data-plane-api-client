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

package tcp_request_rule

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

// NewReplaceTCPRequestRuleParams creates a new ReplaceTCPRequestRuleParams object
// with the default values initialized.
func NewReplaceTCPRequestRuleParams() *ReplaceTCPRequestRuleParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &ReplaceTCPRequestRuleParams{
		ForceReload: &forceReloadDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceTCPRequestRuleParamsWithTimeout creates a new ReplaceTCPRequestRuleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceTCPRequestRuleParamsWithTimeout(timeout time.Duration) *ReplaceTCPRequestRuleParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &ReplaceTCPRequestRuleParams{
		ForceReload: &forceReloadDefault,

		timeout: timeout,
	}
}

// NewReplaceTCPRequestRuleParamsWithContext creates a new ReplaceTCPRequestRuleParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceTCPRequestRuleParamsWithContext(ctx context.Context) *ReplaceTCPRequestRuleParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &ReplaceTCPRequestRuleParams{
		ForceReload: &forceReloadDefault,

		Context: ctx,
	}
}

// NewReplaceTCPRequestRuleParamsWithHTTPClient creates a new ReplaceTCPRequestRuleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceTCPRequestRuleParamsWithHTTPClient(client *http.Client) *ReplaceTCPRequestRuleParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &ReplaceTCPRequestRuleParams{
		ForceReload: &forceReloadDefault,
		HTTPClient:  client,
	}
}

/*ReplaceTCPRequestRuleParams contains all the parameters to send to the API endpoint
for the replace TCP request rule operation typically these are written to a http.Request
*/
type ReplaceTCPRequestRuleParams struct {

	/*Data*/
	Data *models.TCPRequestRule
	/*ForceReload
	  If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.

	*/
	ForceReload *bool
	/*Index
	  TCP Request Rule Index

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
	/*Version
	  Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it's own version.

	*/
	Version *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the replace TCP request rule params
func (o *ReplaceTCPRequestRuleParams) WithTimeout(timeout time.Duration) *ReplaceTCPRequestRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace TCP request rule params
func (o *ReplaceTCPRequestRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace TCP request rule params
func (o *ReplaceTCPRequestRuleParams) WithContext(ctx context.Context) *ReplaceTCPRequestRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace TCP request rule params
func (o *ReplaceTCPRequestRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace TCP request rule params
func (o *ReplaceTCPRequestRuleParams) WithHTTPClient(client *http.Client) *ReplaceTCPRequestRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace TCP request rule params
func (o *ReplaceTCPRequestRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the replace TCP request rule params
func (o *ReplaceTCPRequestRuleParams) WithData(data *models.TCPRequestRule) *ReplaceTCPRequestRuleParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the replace TCP request rule params
func (o *ReplaceTCPRequestRuleParams) SetData(data *models.TCPRequestRule) {
	o.Data = data
}

// WithForceReload adds the forceReload to the replace TCP request rule params
func (o *ReplaceTCPRequestRuleParams) WithForceReload(forceReload *bool) *ReplaceTCPRequestRuleParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the replace TCP request rule params
func (o *ReplaceTCPRequestRuleParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithIndex adds the index to the replace TCP request rule params
func (o *ReplaceTCPRequestRuleParams) WithIndex(index int64) *ReplaceTCPRequestRuleParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the replace TCP request rule params
func (o *ReplaceTCPRequestRuleParams) SetIndex(index int64) {
	o.Index = index
}

// WithParentName adds the parentName to the replace TCP request rule params
func (o *ReplaceTCPRequestRuleParams) WithParentName(parentName string) *ReplaceTCPRequestRuleParams {
	o.SetParentName(parentName)
	return o
}

// SetParentName adds the parentName to the replace TCP request rule params
func (o *ReplaceTCPRequestRuleParams) SetParentName(parentName string) {
	o.ParentName = parentName
}

// WithParentType adds the parentType to the replace TCP request rule params
func (o *ReplaceTCPRequestRuleParams) WithParentType(parentType string) *ReplaceTCPRequestRuleParams {
	o.SetParentType(parentType)
	return o
}

// SetParentType adds the parentType to the replace TCP request rule params
func (o *ReplaceTCPRequestRuleParams) SetParentType(parentType string) {
	o.ParentType = parentType
}

// WithTransactionID adds the transactionID to the replace TCP request rule params
func (o *ReplaceTCPRequestRuleParams) WithTransactionID(transactionID *string) *ReplaceTCPRequestRuleParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the replace TCP request rule params
func (o *ReplaceTCPRequestRuleParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the replace TCP request rule params
func (o *ReplaceTCPRequestRuleParams) WithVersion(version *int64) *ReplaceTCPRequestRuleParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the replace TCP request rule params
func (o *ReplaceTCPRequestRuleParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceTCPRequestRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
