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

package bind

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

// DeleteBindReader is a Reader for the DeleteBind structure.
type DeleteBindReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteBindReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteBindAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteBindNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteBindNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteBindDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteBindAccepted creates a DeleteBindAccepted with default headers values
func NewDeleteBindAccepted() *DeleteBindAccepted {
	return &DeleteBindAccepted{}
}

/*DeleteBindAccepted handles this case with default header values.

Configuration change accepted and reload requested
*/
type DeleteBindAccepted struct {
	/*ID of the requested reload
	 */
	ReloadID string
}

func (o *DeleteBindAccepted) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/binds/{name}][%d] deleteBindAccepted ", 202)
}

func (o *DeleteBindAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Reload-ID
	o.ReloadID = response.GetHeader("Reload-ID")

	return nil
}

// NewDeleteBindNoContent creates a DeleteBindNoContent with default headers values
func NewDeleteBindNoContent() *DeleteBindNoContent {
	return &DeleteBindNoContent{}
}

/*DeleteBindNoContent handles this case with default header values.

Bind deleted
*/
type DeleteBindNoContent struct {
}

func (o *DeleteBindNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/binds/{name}][%d] deleteBindNoContent ", 204)
}

func (o *DeleteBindNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteBindNotFound creates a DeleteBindNotFound with default headers values
func NewDeleteBindNotFound() *DeleteBindNotFound {
	return &DeleteBindNotFound{
		ConfigurationVersion: 0,
	}
}

/*DeleteBindNotFound handles this case with default header values.

The specified resource was not found
*/
type DeleteBindNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *DeleteBindNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/binds/{name}][%d] deleteBindNotFound  %+v", 404, o.Payload)
}

func (o *DeleteBindNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteBindNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteBindDefault creates a DeleteBindDefault with default headers values
func NewDeleteBindDefault(code int) *DeleteBindDefault {
	return &DeleteBindDefault{
		_statusCode:          code,
		ConfigurationVersion: 0,
	}
}

/*DeleteBindDefault handles this case with default header values.

General Error
*/
type DeleteBindDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

// Code gets the status code for the delete bind default response
func (o *DeleteBindDefault) Code() int {
	return o._statusCode
}

func (o *DeleteBindDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/binds/{name}][%d] deleteBind default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteBindDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteBindDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
