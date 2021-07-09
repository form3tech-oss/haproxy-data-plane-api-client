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

// AddPayloadRuntimeMapReader is a Reader for the AddPayloadRuntimeMap structure.
type AddPayloadRuntimeMapReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddPayloadRuntimeMapReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewAddPayloadRuntimeMapCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAddPayloadRuntimeMapBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewAddPayloadRuntimeMapDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddPayloadRuntimeMapCreated creates a AddPayloadRuntimeMapCreated with default headers values
func NewAddPayloadRuntimeMapCreated() *AddPayloadRuntimeMapCreated {
	return &AddPayloadRuntimeMapCreated{}
}

/*AddPayloadRuntimeMapCreated handles this case with default header values.

Map payload added
*/
type AddPayloadRuntimeMapCreated struct {
	Payload models.MapEntries
}

func (o *AddPayloadRuntimeMapCreated) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/runtime/maps/{name}][%d] addPayloadRuntimeMapCreated  %+v", 201, o.Payload)
}

func (o *AddPayloadRuntimeMapCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddPayloadRuntimeMapBadRequest creates a AddPayloadRuntimeMapBadRequest with default headers values
func NewAddPayloadRuntimeMapBadRequest() *AddPayloadRuntimeMapBadRequest {
	return &AddPayloadRuntimeMapBadRequest{}
}

/*AddPayloadRuntimeMapBadRequest handles this case with default header values.

Bad request
*/
type AddPayloadRuntimeMapBadRequest struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *AddPayloadRuntimeMapBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/runtime/maps/{name}][%d] addPayloadRuntimeMapBadRequest  %+v", 400, o.Payload)
}

func (o *AddPayloadRuntimeMapBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddPayloadRuntimeMapDefault creates a AddPayloadRuntimeMapDefault with default headers values
func NewAddPayloadRuntimeMapDefault(code int) *AddPayloadRuntimeMapDefault {
	return &AddPayloadRuntimeMapDefault{
		_statusCode: code,
	}
}

/*AddPayloadRuntimeMapDefault handles this case with default header values.

General Error
*/
type AddPayloadRuntimeMapDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the add payload runtime map default response
func (o *AddPayloadRuntimeMapDefault) Code() int {
	return o._statusCode
}

func (o *AddPayloadRuntimeMapDefault) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/runtime/maps/{name}][%d] addPayloadRuntimeMap default  %+v", o._statusCode, o.Payload)
}

func (o *AddPayloadRuntimeMapDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
