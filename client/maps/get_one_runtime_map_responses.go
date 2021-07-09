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

package maps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/haproxytech/models/v2"
)

// GetOneRuntimeMapReader is a Reader for the GetOneRuntimeMap structure.
type GetOneRuntimeMapReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOneRuntimeMapReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOneRuntimeMapOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetOneRuntimeMapNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetOneRuntimeMapDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetOneRuntimeMapOK creates a GetOneRuntimeMapOK with default headers values
func NewGetOneRuntimeMapOK() *GetOneRuntimeMapOK {
	return &GetOneRuntimeMapOK{}
}

/*GetOneRuntimeMapOK handles this case with default header values.

Successful operation
*/
type GetOneRuntimeMapOK struct {
	Payload *models.Map
}

func (o *GetOneRuntimeMapOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/maps/{name}][%d] getOneRuntimeMapOK  %+v", 200, o.Payload)
}

func (o *GetOneRuntimeMapOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Map)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOneRuntimeMapNotFound creates a GetOneRuntimeMapNotFound with default headers values
func NewGetOneRuntimeMapNotFound() *GetOneRuntimeMapNotFound {
	return &GetOneRuntimeMapNotFound{}
}

/*GetOneRuntimeMapNotFound handles this case with default header values.

The specified resource was not found
*/
type GetOneRuntimeMapNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *GetOneRuntimeMapNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/maps/{name}][%d] getOneRuntimeMapNotFound  %+v", 404, o.Payload)
}

func (o *GetOneRuntimeMapNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOneRuntimeMapDefault creates a GetOneRuntimeMapDefault with default headers values
func NewGetOneRuntimeMapDefault(code int) *GetOneRuntimeMapDefault {
	return &GetOneRuntimeMapDefault{
		_statusCode: code,
	}
}

/*GetOneRuntimeMapDefault handles this case with default header values.

General Error
*/
type GetOneRuntimeMapDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get one runtime map default response
func (o *GetOneRuntimeMapDefault) Code() int {
	return o._statusCode
}

func (o *GetOneRuntimeMapDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/runtime/maps/{name}][%d] getOneRuntimeMap default  %+v", o._statusCode, o.Payload)
}

func (o *GetOneRuntimeMapDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
