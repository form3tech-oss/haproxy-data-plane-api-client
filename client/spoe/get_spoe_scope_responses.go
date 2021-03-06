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

// GetSpoeScopeReader is a Reader for the GetSpoeScope structure.
type GetSpoeScopeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSpoeScopeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSpoeScopeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetSpoeScopeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetSpoeScopeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSpoeScopeOK creates a GetSpoeScopeOK with default headers values
func NewGetSpoeScopeOK() *GetSpoeScopeOK {
	return &GetSpoeScopeOK{}
}

/*GetSpoeScopeOK handles this case with default header values.

Successful operation
*/
type GetSpoeScopeOK struct {
	/*Spoe configuration file version
	 */
	ConfigurationVersion int64

	Payload *GetSpoeScopeOKBody
}

func (o *GetSpoeScopeOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/spoe/spoe_scopes/{name}][%d] getSpoeScopeOK  %+v", 200, o.Payload)
}

func (o *GetSpoeScopeOK) GetPayload() *GetSpoeScopeOKBody {
	return o.Payload
}

func (o *GetSpoeScopeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	configurationVersion, err := swag.ConvertInt64(response.GetHeader("Configuration-Version"))
	if err != nil {
		return errors.InvalidType("Configuration-Version", "header", "int64", response.GetHeader("Configuration-Version"))
	}
	o.ConfigurationVersion = configurationVersion

	o.Payload = new(GetSpoeScopeOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpoeScopeNotFound creates a GetSpoeScopeNotFound with default headers values
func NewGetSpoeScopeNotFound() *GetSpoeScopeNotFound {
	return &GetSpoeScopeNotFound{
		ConfigurationVersion: 0,
	}
}

/*GetSpoeScopeNotFound handles this case with default header values.

The specified resource was not found
*/
type GetSpoeScopeNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *GetSpoeScopeNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/spoe/spoe_scopes/{name}][%d] getSpoeScopeNotFound  %+v", 404, o.Payload)
}

func (o *GetSpoeScopeNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSpoeScopeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSpoeScopeDefault creates a GetSpoeScopeDefault with default headers values
func NewGetSpoeScopeDefault(code int) *GetSpoeScopeDefault {
	return &GetSpoeScopeDefault{
		_statusCode:          code,
		ConfigurationVersion: 0,
	}
}

/*GetSpoeScopeDefault handles this case with default header values.

General Error
*/
type GetSpoeScopeDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

// Code gets the status code for the get spoe scope default response
func (o *GetSpoeScopeDefault) Code() int {
	return o._statusCode
}

func (o *GetSpoeScopeDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/spoe/spoe_scopes/{name}][%d] getSpoeScope default  %+v", o._statusCode, o.Payload)
}

func (o *GetSpoeScopeDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSpoeScopeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetSpoeScopeOKBody get spoe scope o k body
swagger:model GetSpoeScopeOKBody
*/
type GetSpoeScopeOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	// Required: true
	Data models.SpoeScope `json:"data"`
}

// Validate validates this get spoe scope o k body
func (o *GetSpoeScopeOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSpoeScopeOKBody) validateData(formats strfmt.Registry) error {

	if err := o.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getSpoeScopeOK" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetSpoeScopeOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSpoeScopeOKBody) UnmarshalBinary(b []byte) error {
	var res GetSpoeScopeOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
