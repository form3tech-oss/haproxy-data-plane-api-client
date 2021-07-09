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

// CreateSpoeReader is a Reader for the CreateSpoe structure.
type CreateSpoeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSpoeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateSpoeCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateSpoeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateSpoeConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCreateSpoeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateSpoeCreated creates a CreateSpoeCreated with default headers values
func NewCreateSpoeCreated() *CreateSpoeCreated {
	return &CreateSpoeCreated{}
}

/*CreateSpoeCreated handles this case with default header values.

SPOE file created with its entries
*/
type CreateSpoeCreated struct {
	Payload string
}

func (o *CreateSpoeCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/spoe/spoe_files][%d] createSpoeCreated  %+v", 201, o.Payload)
}

func (o *CreateSpoeCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSpoeBadRequest creates a CreateSpoeBadRequest with default headers values
func NewCreateSpoeBadRequest() *CreateSpoeBadRequest {
	return &CreateSpoeBadRequest{}
}

/*CreateSpoeBadRequest handles this case with default header values.

Bad request
*/
type CreateSpoeBadRequest struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateSpoeBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/spoe/spoe_files][%d] createSpoeBadRequest  %+v", 400, o.Payload)
}

func (o *CreateSpoeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSpoeConflict creates a CreateSpoeConflict with default headers values
func NewCreateSpoeConflict() *CreateSpoeConflict {
	return &CreateSpoeConflict{}
}

/*CreateSpoeConflict handles this case with default header values.

The specified resource already exists
*/
type CreateSpoeConflict struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateSpoeConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/spoe/spoe_files][%d] createSpoeConflict  %+v", 409, o.Payload)
}

func (o *CreateSpoeConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSpoeDefault creates a CreateSpoeDefault with default headers values
func NewCreateSpoeDefault(code int) *CreateSpoeDefault {
	return &CreateSpoeDefault{
		_statusCode: code,
	}
}

/*CreateSpoeDefault handles this case with default header values.

General Error
*/
type CreateSpoeDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the create spoe default response
func (o *CreateSpoeDefault) Code() int {
	return o._statusCode
}

func (o *CreateSpoeDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/spoe/spoe_files][%d] createSpoe default  %+v", o._statusCode, o.Payload)
}

func (o *CreateSpoeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
