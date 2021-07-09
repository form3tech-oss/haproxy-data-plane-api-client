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

package tcp_request_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new tcp request rule API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for tcp request rule API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateTCPRequestRule adds a new TCP request rule

Adds a new TCP Request Rule of the specified type in the specified parent.
*/
func (a *Client) CreateTCPRequestRule(params *CreateTCPRequestRuleParams, authInfo runtime.ClientAuthInfoWriter) (*CreateTCPRequestRuleCreated, *CreateTCPRequestRuleAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateTCPRequestRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createTCPRequestRule",
		Method:             "POST",
		PathPattern:        "/services/haproxy/configuration/tcp_request_rules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateTCPRequestRuleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateTCPRequestRuleCreated:
		return value, nil, nil
	case *CreateTCPRequestRuleAccepted:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
DeleteTCPRequestRule deletes a TCP request rule

Deletes a TCP Request Rule configuration by it's index from the specified parent.
*/
func (a *Client) DeleteTCPRequestRule(params *DeleteTCPRequestRuleParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteTCPRequestRuleAccepted, *DeleteTCPRequestRuleNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTCPRequestRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteTCPRequestRule",
		Method:             "DELETE",
		PathPattern:        "/services/haproxy/configuration/tcp_request_rules/{index}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteTCPRequestRuleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteTCPRequestRuleAccepted:
		return value, nil, nil
	case *DeleteTCPRequestRuleNoContent:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
GetTCPRequestRule returns one TCP request rule

Returns one TCP Request Rule configuration by it's index in the specified parent.
*/
func (a *Client) GetTCPRequestRule(params *GetTCPRequestRuleParams, authInfo runtime.ClientAuthInfoWriter) (*GetTCPRequestRuleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTCPRequestRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTCPRequestRule",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/tcp_request_rules/{index}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTCPRequestRuleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTCPRequestRuleOK), nil

}

/*
GetTCPRequestRules returns an array of all TCP request rules

Returns all TCP Request Rules that are configured in specified parent and parent type.
*/
func (a *Client) GetTCPRequestRules(params *GetTCPRequestRulesParams, authInfo runtime.ClientAuthInfoWriter) (*GetTCPRequestRulesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTCPRequestRulesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTCPRequestRules",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/tcp_request_rules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTCPRequestRulesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTCPRequestRulesOK), nil

}

/*
ReplaceTCPRequestRule replaces a TCP request rule

Replaces a TCP Request Rule configuration by it's index in the specified parent.
*/
func (a *Client) ReplaceTCPRequestRule(params *ReplaceTCPRequestRuleParams, authInfo runtime.ClientAuthInfoWriter) (*ReplaceTCPRequestRuleOK, *ReplaceTCPRequestRuleAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceTCPRequestRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "replaceTCPRequestRule",
		Method:             "PUT",
		PathPattern:        "/services/haproxy/configuration/tcp_request_rules/{index}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReplaceTCPRequestRuleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ReplaceTCPRequestRuleOK:
		return value, nil, nil
	case *ReplaceTCPRequestRuleAccepted:
		return nil, value, nil
	}
	return nil, nil, nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
