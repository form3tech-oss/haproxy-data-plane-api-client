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

package log_target

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/haproxytech/models"
)

// GetLogTargetsReader is a Reader for the GetLogTargets structure.
type GetLogTargetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLogTargetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLogTargetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLogTargetsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLogTargetsOK creates a GetLogTargetsOK with default headers values
func NewGetLogTargetsOK() *GetLogTargetsOK {
	return &GetLogTargetsOK{}
}

/*GetLogTargetsOK handles this case with default header values.

Successful operation
*/
type GetLogTargetsOK struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *GetLogTargetsOKBody
}

func (o *GetLogTargetsOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/log_targets][%d] getLogTargetsOK  %+v", 200, o.Payload)
}

func (o *GetLogTargetsOK) GetPayload() *GetLogTargetsOKBody {
	return o.Payload
}

func (o *GetLogTargetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	configurationVersion, err := swag.ConvertInt64(response.GetHeader("Configuration-Version"))
	if err != nil {
		return errors.InvalidType("Configuration-Version", "header", "int64", response.GetHeader("Configuration-Version"))
	}
	o.ConfigurationVersion = configurationVersion

	o.Payload = new(GetLogTargetsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogTargetsDefault creates a GetLogTargetsDefault with default headers values
func NewGetLogTargetsDefault(code int) *GetLogTargetsDefault {
	return &GetLogTargetsDefault{
		_statusCode:          code,
		ConfigurationVersion: 0,
	}
}

/*GetLogTargetsDefault handles this case with default header values.

General Error
*/
type GetLogTargetsDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

// Code gets the status code for the get log targets default response
func (o *GetLogTargetsDefault) Code() int {
	return o._statusCode
}

func (o *GetLogTargetsDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/log_targets][%d] getLogTargets default  %+v", o._statusCode, o.Payload)
}

func (o *GetLogTargetsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLogTargetsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetLogTargetsOKBody get log targets o k body
swagger:model GetLogTargetsOKBody
*/
type GetLogTargetsOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	// Required: true
	Data models.LogTargets `json:"data"`
}

// Validate validates this get log targets o k body
func (o *GetLogTargetsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetLogTargetsOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getLogTargetsOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if err := o.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getLogTargetsOK" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetLogTargetsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLogTargetsOKBody) UnmarshalBinary(b []byte) error {
	var res GetLogTargetsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
