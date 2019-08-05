// Code generated by go-swagger; DO NOT EDIT.

package node_pool_config_items

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new node pool config items API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for node pool config items API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddOrUpdateNodePoolConfigItem adds update config item

Add/update a configuration item unique to the node pool.
*/
func (a *Client) AddOrUpdateNodePoolConfigItem(params *AddOrUpdateNodePoolConfigItemParams, authInfo runtime.ClientAuthInfoWriter) (*AddOrUpdateNodePoolConfigItemOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddOrUpdateNodePoolConfigItemParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addOrUpdateNodePoolConfigItem",
		Method:             "PUT",
		PathPattern:        "/kubernetes-clusters/{cluster_id}/node-pools/{node_pool_name}/config-items/{config_key}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddOrUpdateNodePoolConfigItemReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddOrUpdateNodePoolConfigItemOK), nil

}

/*
DeleteNodePoolConfigItem deletes config item

Deletes config item.
*/
func (a *Client) DeleteNodePoolConfigItem(params *DeleteNodePoolConfigItemParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteNodePoolConfigItemNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNodePoolConfigItemParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteNodePoolConfigItem",
		Method:             "DELETE",
		PathPattern:        "/kubernetes-clusters/{cluster_id}/node-pools/{node_pool_name}/config-items/{config_key}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteNodePoolConfigItemReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteNodePoolConfigItemNoContent), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}