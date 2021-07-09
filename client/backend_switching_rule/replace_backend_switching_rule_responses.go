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

package backend_switching_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/haproxytech/models/v2"
)

// ReplaceBackendSwitchingRuleReader is a Reader for the ReplaceBackendSwitchingRule structure.
type ReplaceBackendSwitchingRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceBackendSwitchingRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewReplaceBackendSwitchingRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewReplaceBackendSwitchingRuleAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewReplaceBackendSwitchingRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewReplaceBackendSwitchingRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewReplaceBackendSwitchingRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceBackendSwitchingRuleOK creates a ReplaceBackendSwitchingRuleOK with default headers values
func NewReplaceBackendSwitchingRuleOK() *ReplaceBackendSwitchingRuleOK {
	return &ReplaceBackendSwitchingRuleOK{}
}

/*ReplaceBackendSwitchingRuleOK handles this case with default header values.

Backend Switching Rule replaced
*/
type ReplaceBackendSwitchingRuleOK struct {
	Payload *models.BackendSwitchingRule
}

func (o *ReplaceBackendSwitchingRuleOK) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/backend_switching_rules/{index}][%d] replaceBackendSwitchingRuleOK  %+v", 200, o.Payload)
}

func (o *ReplaceBackendSwitchingRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BackendSwitchingRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceBackendSwitchingRuleAccepted creates a ReplaceBackendSwitchingRuleAccepted with default headers values
func NewReplaceBackendSwitchingRuleAccepted() *ReplaceBackendSwitchingRuleAccepted {
	return &ReplaceBackendSwitchingRuleAccepted{}
}

/*ReplaceBackendSwitchingRuleAccepted handles this case with default header values.

Configuration change accepted and reload requested
*/
type ReplaceBackendSwitchingRuleAccepted struct {
	/*ID of the requested reload
	 */
	ReloadID string

	Payload *models.BackendSwitchingRule
}

func (o *ReplaceBackendSwitchingRuleAccepted) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/backend_switching_rules/{index}][%d] replaceBackendSwitchingRuleAccepted  %+v", 202, o.Payload)
}

func (o *ReplaceBackendSwitchingRuleAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Reload-ID
	o.ReloadID = response.GetHeader("Reload-ID")

	o.Payload = new(models.BackendSwitchingRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceBackendSwitchingRuleBadRequest creates a ReplaceBackendSwitchingRuleBadRequest with default headers values
func NewReplaceBackendSwitchingRuleBadRequest() *ReplaceBackendSwitchingRuleBadRequest {
	return &ReplaceBackendSwitchingRuleBadRequest{}
}

/*ReplaceBackendSwitchingRuleBadRequest handles this case with default header values.

Bad request
*/
type ReplaceBackendSwitchingRuleBadRequest struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ReplaceBackendSwitchingRuleBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/backend_switching_rules/{index}][%d] replaceBackendSwitchingRuleBadRequest  %+v", 400, o.Payload)
}

func (o *ReplaceBackendSwitchingRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceBackendSwitchingRuleNotFound creates a ReplaceBackendSwitchingRuleNotFound with default headers values
func NewReplaceBackendSwitchingRuleNotFound() *ReplaceBackendSwitchingRuleNotFound {
	return &ReplaceBackendSwitchingRuleNotFound{}
}

/*ReplaceBackendSwitchingRuleNotFound handles this case with default header values.

The specified resource was not found
*/
type ReplaceBackendSwitchingRuleNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ReplaceBackendSwitchingRuleNotFound) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/backend_switching_rules/{index}][%d] replaceBackendSwitchingRuleNotFound  %+v", 404, o.Payload)
}

func (o *ReplaceBackendSwitchingRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceBackendSwitchingRuleDefault creates a ReplaceBackendSwitchingRuleDefault with default headers values
func NewReplaceBackendSwitchingRuleDefault(code int) *ReplaceBackendSwitchingRuleDefault {
	return &ReplaceBackendSwitchingRuleDefault{
		_statusCode: code,
	}
}

/*ReplaceBackendSwitchingRuleDefault handles this case with default header values.

General Error
*/
type ReplaceBackendSwitchingRuleDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the replace backend switching rule default response
func (o *ReplaceBackendSwitchingRuleDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceBackendSwitchingRuleDefault) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/backend_switching_rules/{index}][%d] replaceBackendSwitchingRule default  %+v", o._statusCode, o.Payload)
}

func (o *ReplaceBackendSwitchingRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
