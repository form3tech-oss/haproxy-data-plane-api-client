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

package service_discovery

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

// NewGetAWSRegionParams creates a new GetAWSRegionParams object
// with the default values initialized.
func NewGetAWSRegionParams() *GetAWSRegionParams {
	var ()
	return &GetAWSRegionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAWSRegionParamsWithTimeout creates a new GetAWSRegionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAWSRegionParamsWithTimeout(timeout time.Duration) *GetAWSRegionParams {
	var ()
	return &GetAWSRegionParams{

		timeout: timeout,
	}
}

// NewGetAWSRegionParamsWithContext creates a new GetAWSRegionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAWSRegionParamsWithContext(ctx context.Context) *GetAWSRegionParams {
	var ()
	return &GetAWSRegionParams{

		Context: ctx,
	}
}

// NewGetAWSRegionParamsWithHTTPClient creates a new GetAWSRegionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAWSRegionParamsWithHTTPClient(client *http.Client) *GetAWSRegionParams {
	var ()
	return &GetAWSRegionParams{
		HTTPClient: client,
	}
}

/*GetAWSRegionParams contains all the parameters to send to the API endpoint
for the get a w s region operation typically these are written to a http.Request
*/
type GetAWSRegionParams struct {

	/*ID
	  AWS region id

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get a w s region params
func (o *GetAWSRegionParams) WithTimeout(timeout time.Duration) *GetAWSRegionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get a w s region params
func (o *GetAWSRegionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get a w s region params
func (o *GetAWSRegionParams) WithContext(ctx context.Context) *GetAWSRegionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get a w s region params
func (o *GetAWSRegionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get a w s region params
func (o *GetAWSRegionParams) WithHTTPClient(client *http.Client) *GetAWSRegionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get a w s region params
func (o *GetAWSRegionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get a w s region params
func (o *GetAWSRegionParams) WithID(id string) *GetAWSRegionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get a w s region params
func (o *GetAWSRegionParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetAWSRegionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
