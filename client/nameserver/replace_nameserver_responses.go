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

package nameserver

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/models"
)

// ReplaceNameserverReader is a Reader for the ReplaceNameserver structure.
type ReplaceNameserverReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceNameserverReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceNameserverOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewReplaceNameserverAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReplaceNameserverBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReplaceNameserverNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReplaceNameserverDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceNameserverOK creates a ReplaceNameserverOK with default headers values
func NewReplaceNameserverOK() *ReplaceNameserverOK {
	return &ReplaceNameserverOK{}
}

/*ReplaceNameserverOK handles this case with default header values.

Nameserver replaced
*/
type ReplaceNameserverOK struct {
	Payload *models.Nameserver
}

func (o *ReplaceNameserverOK) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/nameservers/{name}][%d] replaceNameserverOK  %+v", 200, o.Payload)
}

func (o *ReplaceNameserverOK) GetPayload() *models.Nameserver {
	return o.Payload
}

func (o *ReplaceNameserverOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Nameserver)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceNameserverAccepted creates a ReplaceNameserverAccepted with default headers values
func NewReplaceNameserverAccepted() *ReplaceNameserverAccepted {
	return &ReplaceNameserverAccepted{}
}

/*ReplaceNameserverAccepted handles this case with default header values.

Configuration change accepted and reload requested
*/
type ReplaceNameserverAccepted struct {
	/*ID of the requested reload
	 */
	ReloadID string

	Payload *models.Nameserver
}

func (o *ReplaceNameserverAccepted) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/nameservers/{name}][%d] replaceNameserverAccepted  %+v", 202, o.Payload)
}

func (o *ReplaceNameserverAccepted) GetPayload() *models.Nameserver {
	return o.Payload
}

func (o *ReplaceNameserverAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Reload-ID
	o.ReloadID = response.GetHeader("Reload-ID")

	o.Payload = new(models.Nameserver)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceNameserverBadRequest creates a ReplaceNameserverBadRequest with default headers values
func NewReplaceNameserverBadRequest() *ReplaceNameserverBadRequest {
	return &ReplaceNameserverBadRequest{
		ConfigurationVersion: 0,
	}
}

/*ReplaceNameserverBadRequest handles this case with default header values.

Bad request
*/
type ReplaceNameserverBadRequest struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *ReplaceNameserverBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/nameservers/{name}][%d] replaceNameserverBadRequest  %+v", 400, o.Payload)
}

func (o *ReplaceNameserverBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceNameserverBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	configurationVersion, err := swag.ConvertInt64(response.GetHeader("Configuration-Version"))
	if err != nil {
		return errors.InvalidType("Configuration-Version", "header", "int64", response.GetHeader("Configuration-Version"))
	}
	o.ConfigurationVersion = configurationVersion

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceNameserverNotFound creates a ReplaceNameserverNotFound with default headers values
func NewReplaceNameserverNotFound() *ReplaceNameserverNotFound {
	return &ReplaceNameserverNotFound{
		ConfigurationVersion: 0,
	}
}

/*ReplaceNameserverNotFound handles this case with default header values.

The specified resource was not found
*/
type ReplaceNameserverNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *ReplaceNameserverNotFound) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/nameservers/{name}][%d] replaceNameserverNotFound  %+v", 404, o.Payload)
}

func (o *ReplaceNameserverNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceNameserverNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	configurationVersion, err := swag.ConvertInt64(response.GetHeader("Configuration-Version"))
	if err != nil {
		return errors.InvalidType("Configuration-Version", "header", "int64", response.GetHeader("Configuration-Version"))
	}
	o.ConfigurationVersion = configurationVersion

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceNameserverDefault creates a ReplaceNameserverDefault with default headers values
func NewReplaceNameserverDefault(code int) *ReplaceNameserverDefault {
	return &ReplaceNameserverDefault{
		_statusCode:          code,
		ConfigurationVersion: 0,
	}
}

/*ReplaceNameserverDefault handles this case with default header values.

General Error
*/
type ReplaceNameserverDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

// Code gets the status code for the replace nameserver default response
func (o *ReplaceNameserverDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceNameserverDefault) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/nameservers/{name}][%d] replaceNameserver default  %+v", o._statusCode, o.Payload)
}

func (o *ReplaceNameserverDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceNameserverDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	configurationVersion, err := swag.ConvertInt64(response.GetHeader("Configuration-Version"))
	if err != nil {
		return errors.InvalidType("Configuration-Version", "header", "int64", response.GetHeader("Configuration-Version"))
	}
	o.ConfigurationVersion = configurationVersion

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
