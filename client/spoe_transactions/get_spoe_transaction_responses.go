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

package spoe_transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/haproxytech/models/v2"
)

// GetSpoeTransactionReader is a Reader for the GetSpoeTransaction structure.
type GetSpoeTransactionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSpoeTransactionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSpoeTransactionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetSpoeTransactionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetSpoeTransactionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSpoeTransactionOK creates a GetSpoeTransactionOK with default headers values
func NewGetSpoeTransactionOK() *GetSpoeTransactionOK {
	return &GetSpoeTransactionOK{}
}

/*GetSpoeTransactionOK handles this case with default header values.

Successful operation
*/
type GetSpoeTransactionOK struct {
	Payload *models.SpoeTransaction
}

func (o *GetSpoeTransactionOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/spoe_transactions/{id}][%d] getSpoeTransactionOK  %+v", 200, o.Payload)
}

func (o *GetSpoeTransactionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SpoeTransaction)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpoeTransactionNotFound creates a GetSpoeTransactionNotFound with default headers values
func NewGetSpoeTransactionNotFound() *GetSpoeTransactionNotFound {
	return &GetSpoeTransactionNotFound{}
}

/*GetSpoeTransactionNotFound handles this case with default header values.

The specified resource was not found
*/
type GetSpoeTransactionNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *GetSpoeTransactionNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/spoe_transactions/{id}][%d] getSpoeTransactionNotFound  %+v", 404, o.Payload)
}

func (o *GetSpoeTransactionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpoeTransactionDefault creates a GetSpoeTransactionDefault with default headers values
func NewGetSpoeTransactionDefault(code int) *GetSpoeTransactionDefault {
	return &GetSpoeTransactionDefault{
		_statusCode: code,
	}
}

/*GetSpoeTransactionDefault handles this case with default header values.

General Error
*/
type GetSpoeTransactionDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get spoe transaction default response
func (o *GetSpoeTransactionDefault) Code() int {
	return o._statusCode
}

func (o *GetSpoeTransactionDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/spoe_transactions/{id}][%d] getSpoeTransaction default  %+v", o._statusCode, o.Payload)
}

func (o *GetSpoeTransactionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
