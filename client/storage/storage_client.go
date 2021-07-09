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
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new storage API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for storage API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateStorageMapFile creates a managed storage map file with its entries

Creates a managed storage map file with its entries.
*/
func (a *Client) CreateStorageMapFile(params *CreateStorageMapFileParams, authInfo runtime.ClientAuthInfoWriter) (*CreateStorageMapFileCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateStorageMapFileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createStorageMapFile",
		Method:             "POST",
		PathPattern:        "/services/haproxy/storage/maps",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateStorageMapFileReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateStorageMapFileCreated), nil

}

/*
CreateStorageSSLCertificate creates s s l certificate

Creates SSL certificate.
*/
func (a *Client) CreateStorageSSLCertificate(params *CreateStorageSSLCertificateParams, authInfo runtime.ClientAuthInfoWriter) (*CreateStorageSSLCertificateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateStorageSSLCertificateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createStorageSSLCertificate",
		Method:             "POST",
		PathPattern:        "/services/haproxy/storage/ssl_certificates",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateStorageSSLCertificateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateStorageSSLCertificateCreated), nil

}

/*
DeleteStorageMap deletes a managed map file from disk

Deletes a managed map file from disk.
*/
func (a *Client) DeleteStorageMap(params *DeleteStorageMapParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteStorageMapNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteStorageMapParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteStorageMap",
		Method:             "DELETE",
		PathPattern:        "/services/haproxy/storage/maps/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteStorageMapReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteStorageMapNoContent), nil

}

/*
DeleteStorageSSLCertificate deletes s s l certificate from disk

Deletes SSL certificate from disk.
*/
func (a *Client) DeleteStorageSSLCertificate(params *DeleteStorageSSLCertificateParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteStorageSSLCertificateAccepted, *DeleteStorageSSLCertificateNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteStorageSSLCertificateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteStorageSSLCertificate",
		Method:             "DELETE",
		PathPattern:        "/services/haproxy/storage/ssl_certificates/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteStorageSSLCertificateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteStorageSSLCertificateAccepted:
		return value, nil, nil
	case *DeleteStorageSSLCertificateNoContent:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
GetAllStorageMapFiles returns a list of all managed map files

Returns a list of all managed map files
*/
func (a *Client) GetAllStorageMapFiles(params *GetAllStorageMapFilesParams, authInfo runtime.ClientAuthInfoWriter) (*GetAllStorageMapFilesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllStorageMapFilesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAllStorageMapFiles",
		Method:             "GET",
		PathPattern:        "/services/haproxy/storage/maps",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAllStorageMapFilesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAllStorageMapFilesOK), nil

}

/*
GetAllStorageSSLCertificates returns all available s s l certificates on disk

Returns all available SSL certificates on disk.
*/
func (a *Client) GetAllStorageSSLCertificates(params *GetAllStorageSSLCertificatesParams, authInfo runtime.ClientAuthInfoWriter) (*GetAllStorageSSLCertificatesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllStorageSSLCertificatesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAllStorageSSLCertificates",
		Method:             "GET",
		PathPattern:        "/services/haproxy/storage/ssl_certificates",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAllStorageSSLCertificatesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAllStorageSSLCertificatesOK), nil

}

/*
GetOneStorageMap returns the contents of one managed map file from disk

Returns the contents of one managed map file from disk
*/
func (a *Client) GetOneStorageMap(params *GetOneStorageMapParams, authInfo runtime.ClientAuthInfoWriter, writer io.Writer) (*GetOneStorageMapOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOneStorageMapParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOneStorageMap",
		Method:             "GET",
		PathPattern:        "/services/haproxy/storage/maps/{name}",
		ProducesMediaTypes: []string{"application/octet-stream"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetOneStorageMapReader{formats: a.formats, writer: writer},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetOneStorageMapOK), nil

}

/*
GetOneStorageSSLCertificate returns one s s l certificate from disk

Returns one SSL certificate from disk.
*/
func (a *Client) GetOneStorageSSLCertificate(params *GetOneStorageSSLCertificateParams, authInfo runtime.ClientAuthInfoWriter) (*GetOneStorageSSLCertificateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOneStorageSSLCertificateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOneStorageSSLCertificate",
		Method:             "GET",
		PathPattern:        "/services/haproxy/storage/ssl_certificates/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetOneStorageSSLCertificateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetOneStorageSSLCertificateOK), nil

}

/*
ReplaceStorageMapFile replaces contents of a managed map file on disk

Replaces the contents of a managed map file on disk
*/
func (a *Client) ReplaceStorageMapFile(params *ReplaceStorageMapFileParams, authInfo runtime.ClientAuthInfoWriter) (*ReplaceStorageMapFileAccepted, *ReplaceStorageMapFileNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceStorageMapFileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "replaceStorageMapFile",
		Method:             "PUT",
		PathPattern:        "/services/haproxy/storage/maps/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"text/plain"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReplaceStorageMapFileReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ReplaceStorageMapFileAccepted:
		return value, nil, nil
	case *ReplaceStorageMapFileNoContent:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
ReplaceStorageSSLCertificate replaces s s l certificates on disk

Replaces SSL certificate on disk.
*/
func (a *Client) ReplaceStorageSSLCertificate(params *ReplaceStorageSSLCertificateParams, authInfo runtime.ClientAuthInfoWriter) (*ReplaceStorageSSLCertificateOK, *ReplaceStorageSSLCertificateAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceStorageSSLCertificateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "replaceStorageSSLCertificate",
		Method:             "PUT",
		PathPattern:        "/services/haproxy/storage/ssl_certificates/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"text/plain"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReplaceStorageSSLCertificateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ReplaceStorageSSLCertificateOK:
		return value, nil, nil
	case *ReplaceStorageSSLCertificateAccepted:
		return nil, value, nil
	}
	return nil, nil, nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
