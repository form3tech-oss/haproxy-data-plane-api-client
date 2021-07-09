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

package http_response_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/haproxytech/models/v2"
)

// ReplaceHTTPResponseRuleReader is a Reader for the ReplaceHTTPResponseRule structure.
type ReplaceHTTPResponseRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceHTTPResponseRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewReplaceHTTPResponseRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewReplaceHTTPResponseRuleAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewReplaceHTTPResponseRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewReplaceHTTPResponseRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewReplaceHTTPResponseRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceHTTPResponseRuleOK creates a ReplaceHTTPResponseRuleOK with default headers values
func NewReplaceHTTPResponseRuleOK() *ReplaceHTTPResponseRuleOK {
	return &ReplaceHTTPResponseRuleOK{}
}

/*ReplaceHTTPResponseRuleOK handles this case with default header values.

HTTP Response Rule replaced
*/
type ReplaceHTTPResponseRuleOK struct {
	Payload *models.HTTPResponseRule
}

func (o *ReplaceHTTPResponseRuleOK) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/http_response_rules/{index}][%d] replaceHttpResponseRuleOK  %+v", 200, o.Payload)
}

func (o *ReplaceHTTPResponseRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPResponseRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceHTTPResponseRuleAccepted creates a ReplaceHTTPResponseRuleAccepted with default headers values
func NewReplaceHTTPResponseRuleAccepted() *ReplaceHTTPResponseRuleAccepted {
	return &ReplaceHTTPResponseRuleAccepted{}
}

/*ReplaceHTTPResponseRuleAccepted handles this case with default header values.

Configuration change accepted and reload requested
*/
type ReplaceHTTPResponseRuleAccepted struct {
	/*ID of the requested reload
	 */
	ReloadID string

	Payload *models.HTTPResponseRule
}

func (o *ReplaceHTTPResponseRuleAccepted) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/http_response_rules/{index}][%d] replaceHttpResponseRuleAccepted  %+v", 202, o.Payload)
}

func (o *ReplaceHTTPResponseRuleAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Reload-ID
	o.ReloadID = response.GetHeader("Reload-ID")

	o.Payload = new(models.HTTPResponseRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceHTTPResponseRuleBadRequest creates a ReplaceHTTPResponseRuleBadRequest with default headers values
func NewReplaceHTTPResponseRuleBadRequest() *ReplaceHTTPResponseRuleBadRequest {
	return &ReplaceHTTPResponseRuleBadRequest{}
}

/*ReplaceHTTPResponseRuleBadRequest handles this case with default header values.

Bad request
*/
type ReplaceHTTPResponseRuleBadRequest struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ReplaceHTTPResponseRuleBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/http_response_rules/{index}][%d] replaceHttpResponseRuleBadRequest  %+v", 400, o.Payload)
}

func (o *ReplaceHTTPResponseRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceHTTPResponseRuleNotFound creates a ReplaceHTTPResponseRuleNotFound with default headers values
func NewReplaceHTTPResponseRuleNotFound() *ReplaceHTTPResponseRuleNotFound {
	return &ReplaceHTTPResponseRuleNotFound{}
}

/*ReplaceHTTPResponseRuleNotFound handles this case with default header values.

The specified resource was not found
*/
type ReplaceHTTPResponseRuleNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ReplaceHTTPResponseRuleNotFound) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/http_response_rules/{index}][%d] replaceHttpResponseRuleNotFound  %+v", 404, o.Payload)
}

func (o *ReplaceHTTPResponseRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceHTTPResponseRuleDefault creates a ReplaceHTTPResponseRuleDefault with default headers values
func NewReplaceHTTPResponseRuleDefault(code int) *ReplaceHTTPResponseRuleDefault {
	return &ReplaceHTTPResponseRuleDefault{
		_statusCode: code,
	}
}

/*ReplaceHTTPResponseRuleDefault handles this case with default header values.

General Error
*/
type ReplaceHTTPResponseRuleDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the replace HTTP response rule default response
func (o *ReplaceHTTPResponseRuleDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceHTTPResponseRuleDefault) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/http_response_rules/{index}][%d] replaceHTTPResponseRule default  %+v", o._statusCode, o.Payload)
}

func (o *ReplaceHTTPResponseRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
