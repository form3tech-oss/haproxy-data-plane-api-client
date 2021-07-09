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

package backend

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new backend API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for backend API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateBackend adds a backend

Adds a new backend to the configuration file.
*/
func (a *Client) CreateBackend(params *CreateBackendParams, authInfo runtime.ClientAuthInfoWriter) (*CreateBackendCreated, *CreateBackendAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateBackendParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createBackend",
		Method:             "POST",
		PathPattern:        "/services/haproxy/configuration/backends",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateBackendReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateBackendCreated:
		return value, nil, nil
	case *CreateBackendAccepted:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
DeleteBackend deletes a backend

Deletes a backend from the configuration by it's name.
*/
func (a *Client) DeleteBackend(params *DeleteBackendParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteBackendAccepted, *DeleteBackendNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteBackendParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteBackend",
		Method:             "DELETE",
		PathPattern:        "/services/haproxy/configuration/backends/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteBackendReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteBackendAccepted:
		return value, nil, nil
	case *DeleteBackendNoContent:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
GetBackend returns a backend

Returns one backend configuration by it's name.
*/
func (a *Client) GetBackend(params *GetBackendParams, authInfo runtime.ClientAuthInfoWriter) (*GetBackendOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBackendParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getBackend",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/backends/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBackendReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetBackendOK), nil

}

/*
GetBackends returns an array of backends

Returns an array of all configured backends.
*/
func (a *Client) GetBackends(params *GetBackendsParams, authInfo runtime.ClientAuthInfoWriter) (*GetBackendsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBackendsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getBackends",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/backends",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBackendsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetBackendsOK), nil

}

/*
ReplaceBackend replaces a backend

Replaces a backend configuration by it's name.
*/
func (a *Client) ReplaceBackend(params *ReplaceBackendParams, authInfo runtime.ClientAuthInfoWriter) (*ReplaceBackendOK, *ReplaceBackendAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceBackendParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "replaceBackend",
		Method:             "PUT",
		PathPattern:        "/services/haproxy/configuration/backends/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReplaceBackendReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ReplaceBackendOK:
		return value, nil, nil
	case *ReplaceBackendAccepted:
		return nil, value, nil
	}
	return nil, nil, nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
