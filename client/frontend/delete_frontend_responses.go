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

package frontend

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/haproxytech/models/v2"
)

// DeleteFrontendReader is a Reader for the DeleteFrontend structure.
type DeleteFrontendReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFrontendReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewDeleteFrontendAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewDeleteFrontendNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteFrontendNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteFrontendDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteFrontendAccepted creates a DeleteFrontendAccepted with default headers values
func NewDeleteFrontendAccepted() *DeleteFrontendAccepted {
	return &DeleteFrontendAccepted{}
}

/*DeleteFrontendAccepted handles this case with default header values.

Configuration change accepted and reload requested
*/
type DeleteFrontendAccepted struct {
	/*ID of the requested reload
	 */
	ReloadID string
}

func (o *DeleteFrontendAccepted) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/frontends/{name}][%d] deleteFrontendAccepted ", 202)
}

func (o *DeleteFrontendAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Reload-ID
	o.ReloadID = response.GetHeader("Reload-ID")

	return nil
}

// NewDeleteFrontendNoContent creates a DeleteFrontendNoContent with default headers values
func NewDeleteFrontendNoContent() *DeleteFrontendNoContent {
	return &DeleteFrontendNoContent{}
}

/*DeleteFrontendNoContent handles this case with default header values.

Frontend deleted
*/
type DeleteFrontendNoContent struct {
}

func (o *DeleteFrontendNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/frontends/{name}][%d] deleteFrontendNoContent ", 204)
}

func (o *DeleteFrontendNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteFrontendNotFound creates a DeleteFrontendNotFound with default headers values
func NewDeleteFrontendNotFound() *DeleteFrontendNotFound {
	return &DeleteFrontendNotFound{}
}

/*DeleteFrontendNotFound handles this case with default header values.

The specified resource was not found
*/
type DeleteFrontendNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *DeleteFrontendNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/frontends/{name}][%d] deleteFrontendNotFound  %+v", 404, o.Payload)
}

func (o *DeleteFrontendNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFrontendDefault creates a DeleteFrontendDefault with default headers values
func NewDeleteFrontendDefault(code int) *DeleteFrontendDefault {
	return &DeleteFrontendDefault{
		_statusCode: code,
	}
}

/*DeleteFrontendDefault handles this case with default header values.

General Error
*/
type DeleteFrontendDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the delete frontend default response
func (o *DeleteFrontendDefault) Code() int {
	return o._statusCode
}

func (o *DeleteFrontendDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/frontends/{name}][%d] deleteFrontend default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteFrontendDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
