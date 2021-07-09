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

package http_response_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new http response rule API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for http response rule API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateHTTPResponseRule adds a new HTTP response rule

Adds a new HTTP Response Rule of the specified type in the specified parent.
*/
func (a *Client) CreateHTTPResponseRule(params *CreateHTTPResponseRuleParams, authInfo runtime.ClientAuthInfoWriter) (*CreateHTTPResponseRuleCreated, *CreateHTTPResponseRuleAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateHTTPResponseRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createHTTPResponseRule",
		Method:             "POST",
		PathPattern:        "/services/haproxy/configuration/http_response_rules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateHTTPResponseRuleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateHTTPResponseRuleCreated:
		return value, nil, nil
	case *CreateHTTPResponseRuleAccepted:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
DeleteHTTPResponseRule deletes a HTTP response rule

Deletes a HTTP Response Rule configuration by it's index from the specified parent.
*/
func (a *Client) DeleteHTTPResponseRule(params *DeleteHTTPResponseRuleParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteHTTPResponseRuleAccepted, *DeleteHTTPResponseRuleNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteHTTPResponseRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteHTTPResponseRule",
		Method:             "DELETE",
		PathPattern:        "/services/haproxy/configuration/http_response_rules/{index}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteHTTPResponseRuleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteHTTPResponseRuleAccepted:
		return value, nil, nil
	case *DeleteHTTPResponseRuleNoContent:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
GetHTTPResponseRule returns one HTTP response rule

Returns one HTTP Response Rule configuration by it's index in the specified parent.
*/
func (a *Client) GetHTTPResponseRule(params *GetHTTPResponseRuleParams, authInfo runtime.ClientAuthInfoWriter) (*GetHTTPResponseRuleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHTTPResponseRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getHTTPResponseRule",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/http_response_rules/{index}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetHTTPResponseRuleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetHTTPResponseRuleOK), nil

}

/*
GetHTTPResponseRules returns an array of all HTTP response rules

Returns all HTTP Response Rules that are configured in specified parent.
*/
func (a *Client) GetHTTPResponseRules(params *GetHTTPResponseRulesParams, authInfo runtime.ClientAuthInfoWriter) (*GetHTTPResponseRulesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHTTPResponseRulesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getHTTPResponseRules",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/http_response_rules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetHTTPResponseRulesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetHTTPResponseRulesOK), nil

}

/*
ReplaceHTTPResponseRule replaces a HTTP response rule

Replaces a HTTP Response Rule configuration by it's index in the specified parent.
*/
func (a *Client) ReplaceHTTPResponseRule(params *ReplaceHTTPResponseRuleParams, authInfo runtime.ClientAuthInfoWriter) (*ReplaceHTTPResponseRuleOK, *ReplaceHTTPResponseRuleAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceHTTPResponseRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "replaceHTTPResponseRule",
		Method:             "PUT",
		PathPattern:        "/services/haproxy/configuration/http_response_rules/{index}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReplaceHTTPResponseRuleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ReplaceHTTPResponseRuleOK:
		return value, nil, nil
	case *ReplaceHTTPResponseRuleAccepted:
		return nil, value, nil
	}
	return nil, nil, nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
