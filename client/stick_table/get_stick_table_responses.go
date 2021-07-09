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

package stick_table

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/haproxytech/models/v2"
)

// GetStickTableReader is a Reader for the GetStickTable structure.
type GetStickTableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStickTableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetStickTableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetStickTableNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetStickTableDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetStickTableOK creates a GetStickTableOK with default headers values
func NewGetStickTableOK() *GetStickTableOK {
	return &GetStickTableOK{}
}

/*GetStickTableOK handles this case with default header values.

Successful operation
*/
type GetStickTableOK struct {
	Payload *models.StickTable
}

func (o *GetStickTableOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/stick_tables/{name}][%d] getStickTableOK  %+v", 200, o.Payload)
}

func (o *GetStickTableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StickTable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStickTableNotFound creates a GetStickTableNotFound with default headers values
func NewGetStickTableNotFound() *GetStickTableNotFound {
	return &GetStickTableNotFound{}
}

/*GetStickTableNotFound handles this case with default header values.

The specified resource was not found
*/
type GetStickTableNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *GetStickTableNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/stick_tables/{name}][%d] getStickTableNotFound  %+v", 404, o.Payload)
}

func (o *GetStickTableNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStickTableDefault creates a GetStickTableDefault with default headers values
func NewGetStickTableDefault(code int) *GetStickTableDefault {
	return &GetStickTableDefault{
		_statusCode: code,
	}
}

/*GetStickTableDefault handles this case with default header values.

General Error
*/
type GetStickTableDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get stick table default response
func (o *GetStickTableDefault) Code() int {
	return o._statusCode
}

func (o *GetStickTableDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/stick_tables/{name}][%d] getStickTable default  %+v", o._statusCode, o.Payload)
}

func (o *GetStickTableDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
