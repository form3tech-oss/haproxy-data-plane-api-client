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

package storage

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

// DeleteStorageSSLCertificateReader is a Reader for the DeleteStorageSSLCertificate structure.
type DeleteStorageSSLCertificateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteStorageSSLCertificateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteStorageSSLCertificateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteStorageSSLCertificateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteStorageSSLCertificateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteStorageSSLCertificateNoContent creates a DeleteStorageSSLCertificateNoContent with default headers values
func NewDeleteStorageSSLCertificateNoContent() *DeleteStorageSSLCertificateNoContent {
	return &DeleteStorageSSLCertificateNoContent{}
}

/*DeleteStorageSSLCertificateNoContent handles this case with default header values.

SSL certificate deleted
*/
type DeleteStorageSSLCertificateNoContent struct {
}

func (o *DeleteStorageSSLCertificateNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/storage/ssl_certificates/{name}][%d] deleteStorageSSLCertificateNoContent ", 204)
}

func (o *DeleteStorageSSLCertificateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteStorageSSLCertificateNotFound creates a DeleteStorageSSLCertificateNotFound with default headers values
func NewDeleteStorageSSLCertificateNotFound() *DeleteStorageSSLCertificateNotFound {
	return &DeleteStorageSSLCertificateNotFound{
		ConfigurationVersion: 0,
	}
}

/*DeleteStorageSSLCertificateNotFound handles this case with default header values.

The specified resource was not found
*/
type DeleteStorageSSLCertificateNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *DeleteStorageSSLCertificateNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/storage/ssl_certificates/{name}][%d] deleteStorageSSLCertificateNotFound  %+v", 404, o.Payload)
}

func (o *DeleteStorageSSLCertificateNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteStorageSSLCertificateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteStorageSSLCertificateDefault creates a DeleteStorageSSLCertificateDefault with default headers values
func NewDeleteStorageSSLCertificateDefault(code int) *DeleteStorageSSLCertificateDefault {
	return &DeleteStorageSSLCertificateDefault{
		_statusCode:          code,
		ConfigurationVersion: 0,
	}
}

/*DeleteStorageSSLCertificateDefault handles this case with default header values.

General Error
*/
type DeleteStorageSSLCertificateDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

// Code gets the status code for the delete storage s s l certificate default response
func (o *DeleteStorageSSLCertificateDefault) Code() int {
	return o._statusCode
}

func (o *DeleteStorageSSLCertificateDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/storage/ssl_certificates/{name}][%d] deleteStorageSSLCertificate default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteStorageSSLCertificateDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteStorageSSLCertificateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
