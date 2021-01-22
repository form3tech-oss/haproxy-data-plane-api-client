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

// CreateStorageSSLCertificateReader is a Reader for the CreateStorageSSLCertificate structure.
type CreateStorageSSLCertificateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateStorageSSLCertificateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateStorageSSLCertificateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateStorageSSLCertificateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateStorageSSLCertificateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateStorageSSLCertificateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateStorageSSLCertificateCreated creates a CreateStorageSSLCertificateCreated with default headers values
func NewCreateStorageSSLCertificateCreated() *CreateStorageSSLCertificateCreated {
	return &CreateStorageSSLCertificateCreated{}
}

/*CreateStorageSSLCertificateCreated handles this case with default header values.

SSL certificate created
*/
type CreateStorageSSLCertificateCreated struct {
	Payload *models.SslCertificate
}

func (o *CreateStorageSSLCertificateCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/storage/ssl_certificates][%d] createStorageSSLCertificateCreated  %+v", 201, o.Payload)
}

func (o *CreateStorageSSLCertificateCreated) GetPayload() *models.SslCertificate {
	return o.Payload
}

func (o *CreateStorageSSLCertificateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SslCertificate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateStorageSSLCertificateBadRequest creates a CreateStorageSSLCertificateBadRequest with default headers values
func NewCreateStorageSSLCertificateBadRequest() *CreateStorageSSLCertificateBadRequest {
	return &CreateStorageSSLCertificateBadRequest{
		ConfigurationVersion: 0,
	}
}

/*CreateStorageSSLCertificateBadRequest handles this case with default header values.

Bad request
*/
type CreateStorageSSLCertificateBadRequest struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *CreateStorageSSLCertificateBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/storage/ssl_certificates][%d] createStorageSSLCertificateBadRequest  %+v", 400, o.Payload)
}

func (o *CreateStorageSSLCertificateBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateStorageSSLCertificateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateStorageSSLCertificateConflict creates a CreateStorageSSLCertificateConflict with default headers values
func NewCreateStorageSSLCertificateConflict() *CreateStorageSSLCertificateConflict {
	return &CreateStorageSSLCertificateConflict{
		ConfigurationVersion: 0,
	}
}

/*CreateStorageSSLCertificateConflict handles this case with default header values.

The specified resource already exists
*/
type CreateStorageSSLCertificateConflict struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *CreateStorageSSLCertificateConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/storage/ssl_certificates][%d] createStorageSSLCertificateConflict  %+v", 409, o.Payload)
}

func (o *CreateStorageSSLCertificateConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateStorageSSLCertificateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateStorageSSLCertificateDefault creates a CreateStorageSSLCertificateDefault with default headers values
func NewCreateStorageSSLCertificateDefault(code int) *CreateStorageSSLCertificateDefault {
	return &CreateStorageSSLCertificateDefault{
		_statusCode:          code,
		ConfigurationVersion: 0,
	}
}

/*CreateStorageSSLCertificateDefault handles this case with default header values.

General Error
*/
type CreateStorageSSLCertificateDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

// Code gets the status code for the create storage s s l certificate default response
func (o *CreateStorageSSLCertificateDefault) Code() int {
	return o._statusCode
}

func (o *CreateStorageSSLCertificateDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/storage/ssl_certificates][%d] createStorageSSLCertificate default  %+v", o._statusCode, o.Payload)
}

func (o *CreateStorageSSLCertificateDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateStorageSSLCertificateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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