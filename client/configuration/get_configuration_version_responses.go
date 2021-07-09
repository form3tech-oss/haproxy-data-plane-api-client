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

package configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/haproxytech/models/v2"
)

// GetConfigurationVersionReader is a Reader for the GetConfigurationVersion structure.
type GetConfigurationVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConfigurationVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetConfigurationVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetConfigurationVersionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetConfigurationVersionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetConfigurationVersionOK creates a GetConfigurationVersionOK with default headers values
func NewGetConfigurationVersionOK() *GetConfigurationVersionOK {
	return &GetConfigurationVersionOK{}
}

/*GetConfigurationVersionOK handles this case with default header values.

Configuration version
*/
type GetConfigurationVersionOK struct {
	Payload int64
}

func (o *GetConfigurationVersionOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/version][%d] getConfigurationVersionOK  %+v", 200, o.Payload)
}

func (o *GetConfigurationVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationVersionNotFound creates a GetConfigurationVersionNotFound with default headers values
func NewGetConfigurationVersionNotFound() *GetConfigurationVersionNotFound {
	return &GetConfigurationVersionNotFound{}
}

/*GetConfigurationVersionNotFound handles this case with default header values.

The specified resource was not found
*/
type GetConfigurationVersionNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *GetConfigurationVersionNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/version][%d] getConfigurationVersionNotFound  %+v", 404, o.Payload)
}

func (o *GetConfigurationVersionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationVersionDefault creates a GetConfigurationVersionDefault with default headers values
func NewGetConfigurationVersionDefault(code int) *GetConfigurationVersionDefault {
	return &GetConfigurationVersionDefault{
		_statusCode: code,
	}
}

/*GetConfigurationVersionDefault handles this case with default header values.

General Error
*/
type GetConfigurationVersionDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get configuration version default response
func (o *GetConfigurationVersionDefault) Code() int {
	return o._statusCode
}

func (o *GetConfigurationVersionDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/version][%d] getConfigurationVersion default  %+v", o._statusCode, o.Payload)
}

func (o *GetConfigurationVersionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
