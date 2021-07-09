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

package spoe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/haproxytech/models/v2"
)

// DeleteSpoeGroupReader is a Reader for the DeleteSpoeGroup structure.
type DeleteSpoeGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSpoeGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteSpoeGroupNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteSpoeGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteSpoeGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteSpoeGroupNoContent creates a DeleteSpoeGroupNoContent with default headers values
func NewDeleteSpoeGroupNoContent() *DeleteSpoeGroupNoContent {
	return &DeleteSpoeGroupNoContent{}
}

/*DeleteSpoeGroupNoContent handles this case with default header values.

Spoe group deleted
*/
type DeleteSpoeGroupNoContent struct {
}

func (o *DeleteSpoeGroupNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/spoe/spoe_groups/{name}][%d] deleteSpoeGroupNoContent ", 204)
}

func (o *DeleteSpoeGroupNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSpoeGroupNotFound creates a DeleteSpoeGroupNotFound with default headers values
func NewDeleteSpoeGroupNotFound() *DeleteSpoeGroupNotFound {
	return &DeleteSpoeGroupNotFound{}
}

/*DeleteSpoeGroupNotFound handles this case with default header values.

The specified resource was not found
*/
type DeleteSpoeGroupNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *DeleteSpoeGroupNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/spoe/spoe_groups/{name}][%d] deleteSpoeGroupNotFound  %+v", 404, o.Payload)
}

func (o *DeleteSpoeGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSpoeGroupDefault creates a DeleteSpoeGroupDefault with default headers values
func NewDeleteSpoeGroupDefault(code int) *DeleteSpoeGroupDefault {
	return &DeleteSpoeGroupDefault{
		_statusCode: code,
	}
}

/*DeleteSpoeGroupDefault handles this case with default header values.

General Error
*/
type DeleteSpoeGroupDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the delete spoe group default response
func (o *DeleteSpoeGroupDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSpoeGroupDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/spoe/spoe_groups/{name}][%d] deleteSpoeGroup default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSpoeGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
