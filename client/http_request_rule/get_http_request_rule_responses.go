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

package http_request_rule

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

// GetHTTPRequestRuleReader is a Reader for the GetHTTPRequestRule structure.
type GetHTTPRequestRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHTTPRequestRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHTTPRequestRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetHTTPRequestRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetHTTPRequestRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetHTTPRequestRuleOK creates a GetHTTPRequestRuleOK with default headers values
func NewGetHTTPRequestRuleOK() *GetHTTPRequestRuleOK {
	return &GetHTTPRequestRuleOK{}
}

/*GetHTTPRequestRuleOK handles this case with default header values.

Successful operation
*/
type GetHTTPRequestRuleOK struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *GetHTTPRequestRuleOKBody
}

func (o *GetHTTPRequestRuleOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/http_request_rules/{index}][%d] getHttpRequestRuleOK  %+v", 200, o.Payload)
}

func (o *GetHTTPRequestRuleOK) GetPayload() *GetHTTPRequestRuleOKBody {
	return o.Payload
}

func (o *GetHTTPRequestRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	configurationVersion, err := swag.ConvertInt64(response.GetHeader("Configuration-Version"))
	if err != nil {
		return errors.InvalidType("Configuration-Version", "header", "int64", response.GetHeader("Configuration-Version"))
	}
	o.ConfigurationVersion = configurationVersion

	o.Payload = new(GetHTTPRequestRuleOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHTTPRequestRuleNotFound creates a GetHTTPRequestRuleNotFound with default headers values
func NewGetHTTPRequestRuleNotFound() *GetHTTPRequestRuleNotFound {
	return &GetHTTPRequestRuleNotFound{
		ConfigurationVersion: 0,
	}
}

/*GetHTTPRequestRuleNotFound handles this case with default header values.

The specified resource was not found
*/
type GetHTTPRequestRuleNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *GetHTTPRequestRuleNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/http_request_rules/{index}][%d] getHttpRequestRuleNotFound  %+v", 404, o.Payload)
}

func (o *GetHTTPRequestRuleNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHTTPRequestRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetHTTPRequestRuleDefault creates a GetHTTPRequestRuleDefault with default headers values
func NewGetHTTPRequestRuleDefault(code int) *GetHTTPRequestRuleDefault {
	return &GetHTTPRequestRuleDefault{
		_statusCode:          code,
		ConfigurationVersion: 0,
	}
}

/*GetHTTPRequestRuleDefault handles this case with default header values.

General Error
*/
type GetHTTPRequestRuleDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

// Code gets the status code for the get HTTP request rule default response
func (o *GetHTTPRequestRuleDefault) Code() int {
	return o._statusCode
}

func (o *GetHTTPRequestRuleDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/http_request_rules/{index}][%d] getHTTPRequestRule default  %+v", o._statusCode, o.Payload)
}

func (o *GetHTTPRequestRuleDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHTTPRequestRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetHTTPRequestRuleOKBody get HTTP request rule o k body
swagger:model GetHTTPRequestRuleOKBody
*/
type GetHTTPRequestRuleOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	Data *models.HTTPRequestRule `json:"data,omitempty"`
}

// Validate validates this get HTTP request rule o k body
func (o *GetHTTPRequestRuleOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetHTTPRequestRuleOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getHttpRequestRuleOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetHTTPRequestRuleOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetHTTPRequestRuleOKBody) UnmarshalBinary(b []byte) error {
	var res GetHTTPRequestRuleOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
