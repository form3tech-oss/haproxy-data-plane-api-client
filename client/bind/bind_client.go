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
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new bind API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for bind API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateBind adds a new bind

Adds a new bind in the specified frontend in the configuration file.
*/
func (a *Client) CreateBind(params *CreateBindParams, authInfo runtime.ClientAuthInfoWriter) (*CreateBindCreated, *CreateBindAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateBindParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createBind",
		Method:             "POST",
		PathPattern:        "/services/haproxy/configuration/binds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateBindReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateBindCreated:
		return value, nil, nil
	case *CreateBindAccepted:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
DeleteBind deletes a bind

Deletes a bind configuration by it's name in the specified frontend.
*/
func (a *Client) DeleteBind(params *DeleteBindParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteBindAccepted, *DeleteBindNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteBindParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteBind",
		Method:             "DELETE",
		PathPattern:        "/services/haproxy/configuration/binds/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteBindReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteBindAccepted:
		return value, nil, nil
	case *DeleteBindNoContent:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
GetBind returns one bind

Returns one bind configuration by it's name in the specified frontend.
*/
func (a *Client) GetBind(params *GetBindParams, authInfo runtime.ClientAuthInfoWriter) (*GetBindOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBindParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getBind",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/binds/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBindReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetBindOK), nil

}

/*
GetBinds returns an array of binds

Returns an array of all binds that are configured in specified frontend.
*/
func (a *Client) GetBinds(params *GetBindsParams, authInfo runtime.ClientAuthInfoWriter) (*GetBindsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBindsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getBinds",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/binds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBindsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetBindsOK), nil

}

/*
ReplaceBind replaces a bind

Replaces a bind configuration by it's name in the specified frontend.
*/
func (a *Client) ReplaceBind(params *ReplaceBindParams, authInfo runtime.ClientAuthInfoWriter) (*ReplaceBindOK, *ReplaceBindAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceBindParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "replaceBind",
		Method:             "PUT",
		PathPattern:        "/services/haproxy/configuration/binds/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReplaceBindReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ReplaceBindOK:
		return value, nil, nil
	case *ReplaceBindAccepted:
		return nil, value, nil
	}
	return nil, nil, nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
