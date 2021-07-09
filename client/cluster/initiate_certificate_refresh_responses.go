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

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/haproxytech/models/v2"
)

// InitiateCertificateRefreshReader is a Reader for the InitiateCertificateRefresh structure.
type InitiateCertificateRefreshReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InitiateCertificateRefreshReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewInitiateCertificateRefreshOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewInitiateCertificateRefreshForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewInitiateCertificateRefreshDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewInitiateCertificateRefreshOK creates a InitiateCertificateRefreshOK with default headers values
func NewInitiateCertificateRefreshOK() *InitiateCertificateRefreshOK {
	return &InitiateCertificateRefreshOK{}
}

/*InitiateCertificateRefreshOK handles this case with default header values.

refresh activated
*/
type InitiateCertificateRefreshOK struct {
}

func (o *InitiateCertificateRefreshOK) Error() string {
	return fmt.Sprintf("[POST /cluster/certificate][%d] initiateCertificateRefreshOK ", 200)
}

func (o *InitiateCertificateRefreshOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInitiateCertificateRefreshForbidden creates a InitiateCertificateRefreshForbidden with default headers values
func NewInitiateCertificateRefreshForbidden() *InitiateCertificateRefreshForbidden {
	return &InitiateCertificateRefreshForbidden{}
}

/*InitiateCertificateRefreshForbidden handles this case with default header values.

refresh not possible
*/
type InitiateCertificateRefreshForbidden struct {
}

func (o *InitiateCertificateRefreshForbidden) Error() string {
	return fmt.Sprintf("[POST /cluster/certificate][%d] initiateCertificateRefreshForbidden ", 403)
}

func (o *InitiateCertificateRefreshForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInitiateCertificateRefreshDefault creates a InitiateCertificateRefreshDefault with default headers values
func NewInitiateCertificateRefreshDefault(code int) *InitiateCertificateRefreshDefault {
	return &InitiateCertificateRefreshDefault{
		_statusCode: code,
	}
}

/*InitiateCertificateRefreshDefault handles this case with default header values.

General Error
*/
type InitiateCertificateRefreshDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the initiate certificate refresh default response
func (o *InitiateCertificateRefreshDefault) Code() int {
	return o._statusCode
}

func (o *InitiateCertificateRefreshDefault) Error() string {
	return fmt.Sprintf("[POST /cluster/certificate][%d] initiateCertificateRefresh default  %+v", o._statusCode, o.Payload)
}

func (o *InitiateCertificateRefreshDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
