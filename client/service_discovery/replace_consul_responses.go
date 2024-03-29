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
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/models"
)

// ReplaceConsulReader is a Reader for the ReplaceConsul structure.
type ReplaceConsulReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceConsulReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceConsulOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReplaceConsulBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReplaceConsulNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReplaceConsulDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceConsulOK creates a ReplaceConsulOK with default headers values
func NewReplaceConsulOK() *ReplaceConsulOK {
	return &ReplaceConsulOK{}
}

/*ReplaceConsulOK handles this case with default header values.

Consul server replaced
*/
type ReplaceConsulOK struct {
	Payload *models.Consul
}

func (o *ReplaceConsulOK) Error() string {
	return fmt.Sprintf("[PUT /service_discovery/consul/{id}][%d] replaceConsulOK  %+v", 200, o.Payload)
}

func (o *ReplaceConsulOK) GetPayload() *models.Consul {
	return o.Payload
}

func (o *ReplaceConsulOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Consul)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceConsulBadRequest creates a ReplaceConsulBadRequest with default headers values
func NewReplaceConsulBadRequest() *ReplaceConsulBadRequest {
	return &ReplaceConsulBadRequest{
		ConfigurationVersion: 0,
	}
}

/*ReplaceConsulBadRequest handles this case with default header values.

Bad request
*/
type ReplaceConsulBadRequest struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *ReplaceConsulBadRequest) Error() string {
	return fmt.Sprintf("[PUT /service_discovery/consul/{id}][%d] replaceConsulBadRequest  %+v", 400, o.Payload)
}

func (o *ReplaceConsulBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceConsulBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceConsulNotFound creates a ReplaceConsulNotFound with default headers values
func NewReplaceConsulNotFound() *ReplaceConsulNotFound {
	return &ReplaceConsulNotFound{
		ConfigurationVersion: 0,
	}
}

/*ReplaceConsulNotFound handles this case with default header values.

The specified resource was not found
*/
type ReplaceConsulNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *ReplaceConsulNotFound) Error() string {
	return fmt.Sprintf("[PUT /service_discovery/consul/{id}][%d] replaceConsulNotFound  %+v", 404, o.Payload)
}

func (o *ReplaceConsulNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceConsulNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceConsulDefault creates a ReplaceConsulDefault with default headers values
func NewReplaceConsulDefault(code int) *ReplaceConsulDefault {
	return &ReplaceConsulDefault{
		_statusCode:          code,
		ConfigurationVersion: 0,
	}
}

/*ReplaceConsulDefault handles this case with default header values.

General Error
*/
type ReplaceConsulDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

// Code gets the status code for the replace consul default response
func (o *ReplaceConsulDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceConsulDefault) Error() string {
	return fmt.Sprintf("[PUT /service_discovery/consul/{id}][%d] replaceConsul default  %+v", o._statusCode, o.Payload)
}

func (o *ReplaceConsulDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceConsulDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
