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

package backend_switching_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new backend switching rule API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for backend switching rule API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateBackendSwitchingRule(params *CreateBackendSwitchingRuleParams, authInfo runtime.ClientAuthInfoWriter) (*CreateBackendSwitchingRuleCreated, *CreateBackendSwitchingRuleAccepted, error)

	DeleteBackendSwitchingRule(params *DeleteBackendSwitchingRuleParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteBackendSwitchingRuleAccepted, *DeleteBackendSwitchingRuleNoContent, error)

	GetBackendSwitchingRule(params *GetBackendSwitchingRuleParams, authInfo runtime.ClientAuthInfoWriter) (*GetBackendSwitchingRuleOK, error)

	GetBackendSwitchingRules(params *GetBackendSwitchingRulesParams, authInfo runtime.ClientAuthInfoWriter) (*GetBackendSwitchingRulesOK, error)

	ReplaceBackendSwitchingRule(params *ReplaceBackendSwitchingRuleParams, authInfo runtime.ClientAuthInfoWriter) (*ReplaceBackendSwitchingRuleOK, *ReplaceBackendSwitchingRuleAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateBackendSwitchingRule adds a new backend switching rule

  Adds a new Backend Switching Rule of the specified type in the specified frontend.
*/
func (a *Client) CreateBackendSwitchingRule(params *CreateBackendSwitchingRuleParams, authInfo runtime.ClientAuthInfoWriter) (*CreateBackendSwitchingRuleCreated, *CreateBackendSwitchingRuleAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateBackendSwitchingRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createBackendSwitchingRule",
		Method:             "POST",
		PathPattern:        "/services/haproxy/configuration/backend_switching_rules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateBackendSwitchingRuleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateBackendSwitchingRuleCreated:
		return value, nil, nil
	case *CreateBackendSwitchingRuleAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateBackendSwitchingRuleDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteBackendSwitchingRule deletes a backend switching rule

  Deletes a Backend Switching Rule configuration by it's index from the specified frontend.
*/
func (a *Client) DeleteBackendSwitchingRule(params *DeleteBackendSwitchingRuleParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteBackendSwitchingRuleAccepted, *DeleteBackendSwitchingRuleNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteBackendSwitchingRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteBackendSwitchingRule",
		Method:             "DELETE",
		PathPattern:        "/services/haproxy/configuration/backend_switching_rules/{index}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteBackendSwitchingRuleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteBackendSwitchingRuleAccepted:
		return value, nil, nil
	case *DeleteBackendSwitchingRuleNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteBackendSwitchingRuleDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetBackendSwitchingRule returns one backend switching rule

  Returns one Backend Switching Rule configuration by it's index in the specified frontend.
*/
func (a *Client) GetBackendSwitchingRule(params *GetBackendSwitchingRuleParams, authInfo runtime.ClientAuthInfoWriter) (*GetBackendSwitchingRuleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBackendSwitchingRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getBackendSwitchingRule",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/backend_switching_rules/{index}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBackendSwitchingRuleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBackendSwitchingRuleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetBackendSwitchingRuleDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetBackendSwitchingRules returns an array of all backend switching rules

  Returns all Backend Switching Rules that are configured in specified frontend.
*/
func (a *Client) GetBackendSwitchingRules(params *GetBackendSwitchingRulesParams, authInfo runtime.ClientAuthInfoWriter) (*GetBackendSwitchingRulesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBackendSwitchingRulesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getBackendSwitchingRules",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/backend_switching_rules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBackendSwitchingRulesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBackendSwitchingRulesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetBackendSwitchingRulesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ReplaceBackendSwitchingRule replaces a backend switching rule

  Replaces a Backend Switching Rule configuration by it's index in the specified frontend.
*/
func (a *Client) ReplaceBackendSwitchingRule(params *ReplaceBackendSwitchingRuleParams, authInfo runtime.ClientAuthInfoWriter) (*ReplaceBackendSwitchingRuleOK, *ReplaceBackendSwitchingRuleAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceBackendSwitchingRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "replaceBackendSwitchingRule",
		Method:             "PUT",
		PathPattern:        "/services/haproxy/configuration/backend_switching_rules/{index}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReplaceBackendSwitchingRuleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ReplaceBackendSwitchingRuleOK:
		return value, nil, nil
	case *ReplaceBackendSwitchingRuleAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReplaceBackendSwitchingRuleDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
