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

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v2/models"
)

// ReplaceAWSRegionReader is a Reader for the ReplaceAWSRegion structure.
type ReplaceAWSRegionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceAWSRegionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewReplaceAWSRegionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewReplaceAWSRegionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewReplaceAWSRegionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewReplaceAWSRegionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceAWSRegionOK creates a ReplaceAWSRegionOK with default headers values
func NewReplaceAWSRegionOK() *ReplaceAWSRegionOK {
	return &ReplaceAWSRegionOK{}
}

/*ReplaceAWSRegionOK handles this case with default header values.

Resource updated
*/
type ReplaceAWSRegionOK struct {
	Payload *models.AwsRegion
}

func (o *ReplaceAWSRegionOK) Error() string {
	return fmt.Sprintf("[PUT /service_discovery/aws/{id}][%d] replaceAWSRegionOK  %+v", 200, o.Payload)
}

func (o *ReplaceAWSRegionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AwsRegion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceAWSRegionBadRequest creates a ReplaceAWSRegionBadRequest with default headers values
func NewReplaceAWSRegionBadRequest() *ReplaceAWSRegionBadRequest {
	return &ReplaceAWSRegionBadRequest{}
}

/*ReplaceAWSRegionBadRequest handles this case with default header values.

Bad request
*/
type ReplaceAWSRegionBadRequest struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ReplaceAWSRegionBadRequest) Error() string {
	return fmt.Sprintf("[PUT /service_discovery/aws/{id}][%d] replaceAWSRegionBadRequest  %+v", 400, o.Payload)
}

func (o *ReplaceAWSRegionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceAWSRegionNotFound creates a ReplaceAWSRegionNotFound with default headers values
func NewReplaceAWSRegionNotFound() *ReplaceAWSRegionNotFound {
	return &ReplaceAWSRegionNotFound{}
}

/*ReplaceAWSRegionNotFound handles this case with default header values.

The specified resource was not found
*/
type ReplaceAWSRegionNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ReplaceAWSRegionNotFound) Error() string {
	return fmt.Sprintf("[PUT /service_discovery/aws/{id}][%d] replaceAWSRegionNotFound  %+v", 404, o.Payload)
}

func (o *ReplaceAWSRegionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceAWSRegionDefault creates a ReplaceAWSRegionDefault with default headers values
func NewReplaceAWSRegionDefault(code int) *ReplaceAWSRegionDefault {
	return &ReplaceAWSRegionDefault{
		_statusCode: code,
	}
}

/*ReplaceAWSRegionDefault handles this case with default header values.

General Error
*/
type ReplaceAWSRegionDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the replace a w s region default response
func (o *ReplaceAWSRegionDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceAWSRegionDefault) Error() string {
	return fmt.Sprintf("[PUT /service_discovery/aws/{id}][%d] replaceAWSRegion default  %+v", o._statusCode, o.Payload)
}

func (o *ReplaceAWSRegionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
