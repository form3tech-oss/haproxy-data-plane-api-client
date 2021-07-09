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

package service_discovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v2/models"
)

// GetAWSRegionsReader is a Reader for the GetAWSRegions structure.
type GetAWSRegionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAWSRegionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAWSRegionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetAWSRegionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAWSRegionsOK creates a GetAWSRegionsOK with default headers values
func NewGetAWSRegionsOK() *GetAWSRegionsOK {
	return &GetAWSRegionsOK{}
}

/*GetAWSRegionsOK handles this case with default header values.

Successful operation
*/
type GetAWSRegionsOK struct {
	Payload *GetAWSRegionsOKBody
}

func (o *GetAWSRegionsOK) Error() string {
	return fmt.Sprintf("[GET /service_discovery/aws][%d] getAWSRegionsOK  %+v", 200, o.Payload)
}

func (o *GetAWSRegionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetAWSRegionsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAWSRegionsDefault creates a GetAWSRegionsDefault with default headers values
func NewGetAWSRegionsDefault(code int) *GetAWSRegionsDefault {
	return &GetAWSRegionsDefault{
		_statusCode: code,
	}
}

/*GetAWSRegionsDefault handles this case with default header values.

General Error
*/
type GetAWSRegionsDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get a w s regions default response
func (o *GetAWSRegionsDefault) Code() int {
	return o._statusCode
}

func (o *GetAWSRegionsDefault) Error() string {
	return fmt.Sprintf("[GET /service_discovery/aws][%d] getAWSRegions default  %+v", o._statusCode, o.Payload)
}

func (o *GetAWSRegionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetAWSRegionsOKBody get a w s regions o k body
swagger:model GetAWSRegionsOKBody
*/
type GetAWSRegionsOKBody struct {

	// data
	// Required: true
	Data models.AwsRegions `json:"data"`
}

// Validate validates this get a w s regions o k body
func (o *GetAWSRegionsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAWSRegionsOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getAWSRegionsOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if err := o.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getAWSRegionsOK" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAWSRegionsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAWSRegionsOKBody) UnmarshalBinary(b []byte) error {
	var res GetAWSRegionsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
