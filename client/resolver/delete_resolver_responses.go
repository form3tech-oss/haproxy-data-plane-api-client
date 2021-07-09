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

package resolver

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/haproxytech/models/v2"
)

// DeleteResolverReader is a Reader for the DeleteResolver structure.
type DeleteResolverReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteResolverReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewDeleteResolverAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewDeleteResolverNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteResolverNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteResolverDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteResolverAccepted creates a DeleteResolverAccepted with default headers values
func NewDeleteResolverAccepted() *DeleteResolverAccepted {
	return &DeleteResolverAccepted{}
}

/*DeleteResolverAccepted handles this case with default header values.

Configuration change accepted and reload requested
*/
type DeleteResolverAccepted struct {
	/*ID of the requested reload
	 */
	ReloadID string
}

func (o *DeleteResolverAccepted) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/resolvers/{name}][%d] deleteResolverAccepted ", 202)
}

func (o *DeleteResolverAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Reload-ID
	o.ReloadID = response.GetHeader("Reload-ID")

	return nil
}

// NewDeleteResolverNoContent creates a DeleteResolverNoContent with default headers values
func NewDeleteResolverNoContent() *DeleteResolverNoContent {
	return &DeleteResolverNoContent{}
}

/*DeleteResolverNoContent handles this case with default header values.

Resolver deleted
*/
type DeleteResolverNoContent struct {
}

func (o *DeleteResolverNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/resolvers/{name}][%d] deleteResolverNoContent ", 204)
}

func (o *DeleteResolverNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteResolverNotFound creates a DeleteResolverNotFound with default headers values
func NewDeleteResolverNotFound() *DeleteResolverNotFound {
	return &DeleteResolverNotFound{}
}

/*DeleteResolverNotFound handles this case with default header values.

The specified resource was not found
*/
type DeleteResolverNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *DeleteResolverNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/resolvers/{name}][%d] deleteResolverNotFound  %+v", 404, o.Payload)
}

func (o *DeleteResolverNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteResolverDefault creates a DeleteResolverDefault with default headers values
func NewDeleteResolverDefault(code int) *DeleteResolverDefault {
	return &DeleteResolverDefault{
		_statusCode: code,
	}
}

/*DeleteResolverDefault handles this case with default header values.

General Error
*/
type DeleteResolverDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the delete resolver default response
func (o *DeleteResolverDefault) Code() int {
	return o._statusCode
}

func (o *DeleteResolverDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/resolvers/{name}][%d] deleteResolver default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteResolverDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
