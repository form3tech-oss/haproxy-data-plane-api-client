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

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new server API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for server API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateServer adds a new server

Adds a new server in the specified backend in the configuration file.
*/
func (a *Client) CreateServer(params *CreateServerParams, authInfo runtime.ClientAuthInfoWriter) (*CreateServerCreated, *CreateServerAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateServerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createServer",
		Method:             "POST",
		PathPattern:        "/services/haproxy/configuration/servers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateServerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateServerCreated:
		return value, nil, nil
	case *CreateServerAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateServerDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteServer deletes a server

Deletes a server configuration by it's name in the specified backend.
*/
func (a *Client) DeleteServer(params *DeleteServerParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteServerAccepted, *DeleteServerNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteServerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteServer",
		Method:             "DELETE",
		PathPattern:        "/services/haproxy/configuration/servers/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteServerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteServerAccepted:
		return value, nil, nil
	case *DeleteServerNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteServerDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetServer returns one server

Returns one server configuration by it's name in the specified backend.
*/
func (a *Client) GetServer(params *GetServerParams, authInfo runtime.ClientAuthInfoWriter) (*GetServerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getServer",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/servers/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetServerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetServerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetServerDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetServers returns an array of servers

Returns an array of all servers that are configured in specified backend.
*/
func (a *Client) GetServers(params *GetServersParams, authInfo runtime.ClientAuthInfoWriter) (*GetServersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getServers",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/servers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetServersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetServersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetServersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ReplaceServer replaces a server

Replaces a server configuration by it's name in the specified backend.
*/
func (a *Client) ReplaceServer(params *ReplaceServerParams, authInfo runtime.ClientAuthInfoWriter) (*ReplaceServerOK, *ReplaceServerAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceServerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "replaceServer",
		Method:             "PUT",
		PathPattern:        "/services/haproxy/configuration/servers/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReplaceServerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ReplaceServerOK:
		return value, nil, nil
	case *ReplaceServerAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReplaceServerDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}