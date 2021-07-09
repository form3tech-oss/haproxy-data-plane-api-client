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

// CreateSpoeAgentReader is a Reader for the CreateSpoeAgent structure.
type CreateSpoeAgentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSpoeAgentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateSpoeAgentCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateSpoeAgentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateSpoeAgentConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCreateSpoeAgentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateSpoeAgentCreated creates a CreateSpoeAgentCreated with default headers values
func NewCreateSpoeAgentCreated() *CreateSpoeAgentCreated {
	return &CreateSpoeAgentCreated{}
}

/*CreateSpoeAgentCreated handles this case with default header values.

Spoe agent created
*/
type CreateSpoeAgentCreated struct {
	Payload *models.SpoeAgent
}

func (o *CreateSpoeAgentCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/spoe/spoe_agents][%d] createSpoeAgentCreated  %+v", 201, o.Payload)
}

func (o *CreateSpoeAgentCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SpoeAgent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSpoeAgentBadRequest creates a CreateSpoeAgentBadRequest with default headers values
func NewCreateSpoeAgentBadRequest() *CreateSpoeAgentBadRequest {
	return &CreateSpoeAgentBadRequest{}
}

/*CreateSpoeAgentBadRequest handles this case with default header values.

Bad request
*/
type CreateSpoeAgentBadRequest struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateSpoeAgentBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/spoe/spoe_agents][%d] createSpoeAgentBadRequest  %+v", 400, o.Payload)
}

func (o *CreateSpoeAgentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSpoeAgentConflict creates a CreateSpoeAgentConflict with default headers values
func NewCreateSpoeAgentConflict() *CreateSpoeAgentConflict {
	return &CreateSpoeAgentConflict{}
}

/*CreateSpoeAgentConflict handles this case with default header values.

The specified resource already exists
*/
type CreateSpoeAgentConflict struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateSpoeAgentConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/spoe/spoe_agents][%d] createSpoeAgentConflict  %+v", 409, o.Payload)
}

func (o *CreateSpoeAgentConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSpoeAgentDefault creates a CreateSpoeAgentDefault with default headers values
func NewCreateSpoeAgentDefault(code int) *CreateSpoeAgentDefault {
	return &CreateSpoeAgentDefault{
		_statusCode: code,
	}
}

/*CreateSpoeAgentDefault handles this case with default header values.

General Error
*/
type CreateSpoeAgentDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the create spoe agent default response
func (o *CreateSpoeAgentDefault) Code() int {
	return o._statusCode
}

func (o *CreateSpoeAgentDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/spoe/spoe_agents][%d] createSpoeAgent default  %+v", o._statusCode, o.Payload)
}

func (o *CreateSpoeAgentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
