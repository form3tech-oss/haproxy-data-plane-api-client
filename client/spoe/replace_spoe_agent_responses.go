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

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/models"
)

// ReplaceSpoeAgentReader is a Reader for the ReplaceSpoeAgent structure.
type ReplaceSpoeAgentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceSpoeAgentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceSpoeAgentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReplaceSpoeAgentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReplaceSpoeAgentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReplaceSpoeAgentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceSpoeAgentOK creates a ReplaceSpoeAgentOK with default headers values
func NewReplaceSpoeAgentOK() *ReplaceSpoeAgentOK {
	return &ReplaceSpoeAgentOK{}
}

/*ReplaceSpoeAgentOK handles this case with default header values.

Spoe agent replaced
*/
type ReplaceSpoeAgentOK struct {
	Payload *models.SpoeAgent
}

func (o *ReplaceSpoeAgentOK) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/spoe/spoe_agents/{name}][%d] replaceSpoeAgentOK  %+v", 200, o.Payload)
}

func (o *ReplaceSpoeAgentOK) GetPayload() *models.SpoeAgent {
	return o.Payload
}

func (o *ReplaceSpoeAgentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SpoeAgent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceSpoeAgentBadRequest creates a ReplaceSpoeAgentBadRequest with default headers values
func NewReplaceSpoeAgentBadRequest() *ReplaceSpoeAgentBadRequest {
	return &ReplaceSpoeAgentBadRequest{
		ConfigurationVersion: 0,
	}
}

/*ReplaceSpoeAgentBadRequest handles this case with default header values.

Bad request
*/
type ReplaceSpoeAgentBadRequest struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *ReplaceSpoeAgentBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/spoe/spoe_agents/{name}][%d] replaceSpoeAgentBadRequest  %+v", 400, o.Payload)
}

func (o *ReplaceSpoeAgentBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceSpoeAgentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceSpoeAgentNotFound creates a ReplaceSpoeAgentNotFound with default headers values
func NewReplaceSpoeAgentNotFound() *ReplaceSpoeAgentNotFound {
	return &ReplaceSpoeAgentNotFound{
		ConfigurationVersion: 0,
	}
}

/*ReplaceSpoeAgentNotFound handles this case with default header values.

The specified resource was not found
*/
type ReplaceSpoeAgentNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *ReplaceSpoeAgentNotFound) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/spoe/spoe_agents/{name}][%d] replaceSpoeAgentNotFound  %+v", 404, o.Payload)
}

func (o *ReplaceSpoeAgentNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceSpoeAgentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceSpoeAgentDefault creates a ReplaceSpoeAgentDefault with default headers values
func NewReplaceSpoeAgentDefault(code int) *ReplaceSpoeAgentDefault {
	return &ReplaceSpoeAgentDefault{
		_statusCode:          code,
		ConfigurationVersion: 0,
	}
}

/*ReplaceSpoeAgentDefault handles this case with default header values.

General Error
*/
type ReplaceSpoeAgentDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

// Code gets the status code for the replace spoe agent default response
func (o *ReplaceSpoeAgentDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceSpoeAgentDefault) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/spoe/spoe_agents/{name}][%d] replaceSpoeAgent default  %+v", o._statusCode, o.Payload)
}

func (o *ReplaceSpoeAgentDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceSpoeAgentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
