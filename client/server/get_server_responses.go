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

package server

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

// GetServerReader is a Reader for the GetServer structure.
type GetServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetServerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetServerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetServerOK creates a GetServerOK with default headers values
func NewGetServerOK() *GetServerOK {
	return &GetServerOK{}
}

/*GetServerOK handles this case with default header values.

Successful operation
*/
type GetServerOK struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *GetServerOKBody
}

func (o *GetServerOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/servers/{name}][%d] getServerOK  %+v", 200, o.Payload)
}

func (o *GetServerOK) GetPayload() *GetServerOKBody {
	return o.Payload
}

func (o *GetServerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	configurationVersion, err := swag.ConvertInt64(response.GetHeader("Configuration-Version"))
	if err != nil {
		return errors.InvalidType("Configuration-Version", "header", "int64", response.GetHeader("Configuration-Version"))
	}
	o.ConfigurationVersion = configurationVersion

	o.Payload = new(GetServerOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServerNotFound creates a GetServerNotFound with default headers values
func NewGetServerNotFound() *GetServerNotFound {
	return &GetServerNotFound{
		ConfigurationVersion: 0,
	}
}

/*GetServerNotFound handles this case with default header values.

The specified resource was not found
*/
type GetServerNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *GetServerNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/servers/{name}][%d] getServerNotFound  %+v", 404, o.Payload)
}

func (o *GetServerNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetServerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetServerDefault creates a GetServerDefault with default headers values
func NewGetServerDefault(code int) *GetServerDefault {
	return &GetServerDefault{
		_statusCode:          code,
		ConfigurationVersion: 0,
	}
}

/*GetServerDefault handles this case with default header values.

General Error
*/
type GetServerDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

// Code gets the status code for the get server default response
func (o *GetServerDefault) Code() int {
	return o._statusCode
}

func (o *GetServerDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/servers/{name}][%d] getServer default  %+v", o._statusCode, o.Payload)
}

func (o *GetServerDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetServerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetServerOKBody get server o k body
swagger:model GetServerOKBody
*/
type GetServerOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	Data *models.Server `json:"data,omitempty"`
}

// Validate validates this get server o k body
func (o *GetServerOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetServerOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getServerOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetServerOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetServerOKBody) UnmarshalBinary(b []byte) error {
	var res GetServerOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
