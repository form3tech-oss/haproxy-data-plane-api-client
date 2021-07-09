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

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/haproxytech/models/v2"
)

// CreateStorageMapFileReader is a Reader for the CreateStorageMapFile structure.
type CreateStorageMapFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateStorageMapFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateStorageMapFileCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateStorageMapFileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateStorageMapFileConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCreateStorageMapFileDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateStorageMapFileCreated creates a CreateStorageMapFileCreated with default headers values
func NewCreateStorageMapFileCreated() *CreateStorageMapFileCreated {
	return &CreateStorageMapFileCreated{}
}

/*CreateStorageMapFileCreated handles this case with default header values.

Map file created with its entries
*/
type CreateStorageMapFileCreated struct {
	Payload *models.Map
}

func (o *CreateStorageMapFileCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/storage/maps][%d] createStorageMapFileCreated  %+v", 201, o.Payload)
}

func (o *CreateStorageMapFileCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Map)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateStorageMapFileBadRequest creates a CreateStorageMapFileBadRequest with default headers values
func NewCreateStorageMapFileBadRequest() *CreateStorageMapFileBadRequest {
	return &CreateStorageMapFileBadRequest{}
}

/*CreateStorageMapFileBadRequest handles this case with default header values.

Bad request
*/
type CreateStorageMapFileBadRequest struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateStorageMapFileBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/storage/maps][%d] createStorageMapFileBadRequest  %+v", 400, o.Payload)
}

func (o *CreateStorageMapFileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateStorageMapFileConflict creates a CreateStorageMapFileConflict with default headers values
func NewCreateStorageMapFileConflict() *CreateStorageMapFileConflict {
	return &CreateStorageMapFileConflict{}
}

/*CreateStorageMapFileConflict handles this case with default header values.

The specified resource already exists
*/
type CreateStorageMapFileConflict struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateStorageMapFileConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/storage/maps][%d] createStorageMapFileConflict  %+v", 409, o.Payload)
}

func (o *CreateStorageMapFileConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateStorageMapFileDefault creates a CreateStorageMapFileDefault with default headers values
func NewCreateStorageMapFileDefault(code int) *CreateStorageMapFileDefault {
	return &CreateStorageMapFileDefault{
		_statusCode: code,
	}
}

/*CreateStorageMapFileDefault handles this case with default header values.

General Error
*/
type CreateStorageMapFileDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the create storage map file default response
func (o *CreateStorageMapFileDefault) Code() int {
	return o._statusCode
}

func (o *CreateStorageMapFileDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/storage/maps][%d] createStorageMapFile default  %+v", o._statusCode, o.Payload)
}

func (o *CreateStorageMapFileDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
