// Code generated by go-swagger; DO NOT EDIT.

package node_pools

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new node pools API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for node pools API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateOrUpdateNodePool creates update node pool

Create/update a node pool.
*/
func (a *Client) CreateOrUpdateNodePool(params *CreateOrUpdateNodePoolParams, authInfo runtime.ClientAuthInfoWriter) (*CreateOrUpdateNodePoolOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateOrUpdateNodePoolParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createOrUpdateNodePool",
		Method:             "PUT",
		PathPattern:        "/kubernetes-clusters/{cluster_id}/node-pools/{node_pool_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateOrUpdateNodePoolReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateOrUpdateNodePoolOK), nil

}

/*
DeleteNodePool deletes node pool

Deletes node pool.
*/
func (a *Client) DeleteNodePool(params *DeleteNodePoolParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteNodePoolNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNodePoolParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteNodePool",
		Method:             "DELETE",
		PathPattern:        "/kubernetes-clusters/{cluster_id}/node-pools/{node_pool_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteNodePoolReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteNodePoolNoContent), nil

}

/*
ListNodePools lists node pools

List all node pools of a cluster.
*/
func (a *Client) ListNodePools(params *ListNodePoolsParams, authInfo runtime.ClientAuthInfoWriter) (*ListNodePoolsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListNodePoolsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listNodePools",
		Method:             "GET",
		PathPattern:        "/kubernetes-clusters/{cluster_id}/node-pools",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListNodePoolsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListNodePoolsOK), nil

}

/*
UpdateNodePool updates node pool

Update a node pool.
*/
func (a *Client) UpdateNodePool(params *UpdateNodePoolParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateNodePoolOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateNodePoolParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateNodePool",
		Method:             "PATCH",
		PathPattern:        "/kubernetes-clusters/{cluster_id}/node-pools/{node_pool_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateNodePoolReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateNodePoolOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}