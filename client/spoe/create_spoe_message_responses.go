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

// CreateSpoeMessageReader is a Reader for the CreateSpoeMessage structure.
type CreateSpoeMessageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSpoeMessageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateSpoeMessageCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateSpoeMessageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateSpoeMessageConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCreateSpoeMessageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateSpoeMessageCreated creates a CreateSpoeMessageCreated with default headers values
func NewCreateSpoeMessageCreated() *CreateSpoeMessageCreated {
	return &CreateSpoeMessageCreated{}
}

/*CreateSpoeMessageCreated handles this case with default header values.

Spoe message created
*/
type CreateSpoeMessageCreated struct {
	Payload *models.SpoeMessage
}

func (o *CreateSpoeMessageCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/spoe/spoe_messages][%d] createSpoeMessageCreated  %+v", 201, o.Payload)
}

func (o *CreateSpoeMessageCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SpoeMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSpoeMessageBadRequest creates a CreateSpoeMessageBadRequest with default headers values
func NewCreateSpoeMessageBadRequest() *CreateSpoeMessageBadRequest {
	return &CreateSpoeMessageBadRequest{}
}

/*CreateSpoeMessageBadRequest handles this case with default header values.

Bad request
*/
type CreateSpoeMessageBadRequest struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateSpoeMessageBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/spoe/spoe_messages][%d] createSpoeMessageBadRequest  %+v", 400, o.Payload)
}

func (o *CreateSpoeMessageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSpoeMessageConflict creates a CreateSpoeMessageConflict with default headers values
func NewCreateSpoeMessageConflict() *CreateSpoeMessageConflict {
	return &CreateSpoeMessageConflict{}
}

/*CreateSpoeMessageConflict handles this case with default header values.

The specified resource already exists
*/
type CreateSpoeMessageConflict struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateSpoeMessageConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/spoe/spoe_messages][%d] createSpoeMessageConflict  %+v", 409, o.Payload)
}

func (o *CreateSpoeMessageConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSpoeMessageDefault creates a CreateSpoeMessageDefault with default headers values
func NewCreateSpoeMessageDefault(code int) *CreateSpoeMessageDefault {
	return &CreateSpoeMessageDefault{
		_statusCode: code,
	}
}

/*CreateSpoeMessageDefault handles this case with default header values.

General Error
*/
type CreateSpoeMessageDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the create spoe message default response
func (o *CreateSpoeMessageDefault) Code() int {
	return o._statusCode
}

func (o *CreateSpoeMessageDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/spoe/spoe_messages][%d] createSpoeMessage default  %+v", o._statusCode, o.Payload)
}

func (o *CreateSpoeMessageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
