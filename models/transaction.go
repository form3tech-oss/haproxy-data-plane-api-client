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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Transaction Configuration transaction
//
// HAProxy configuration transaction
// swagger:model transaction
type Transaction struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// id
	// Pattern: ^[^\s]+$
	ID string `json:"id,omitempty"`

	// status
	// Enum: [failed outdated in_progress success]
	Status string `json:"status,omitempty"`
}

// Validate validates this transaction
func (m *Transaction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Transaction) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.Pattern("id", "body", string(m.ID), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var transactionTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["failed","outdated","in_progress","success"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		transactionTypeStatusPropEnum = append(transactionTypeStatusPropEnum, v)
	}
}

const (

	// TransactionStatusFailed captures enum value "failed"
	TransactionStatusFailed string = "failed"

	// TransactionStatusOutdated captures enum value "outdated"
	TransactionStatusOutdated string = "outdated"

	// TransactionStatusInProgress captures enum value "in_progress"
	TransactionStatusInProgress string = "in_progress"

	// TransactionStatusSuccess captures enum value "success"
	TransactionStatusSuccess string = "success"
)

// prop value enum
func (m *Transaction) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, transactionTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Transaction) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Transaction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Transaction) UnmarshalBinary(b []byte) error {
	var res Transaction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
