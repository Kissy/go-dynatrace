// Code generated by go-swagger; DO NOT EDIT.

package automatically_applied_tags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new automatically applied tags API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for automatically applied tags API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateAutoTag(params *CreateAutoTagParams, authInfo runtime.ClientAuthInfoWriter) (*CreateAutoTagCreated, error)

	DeleteAutoTag(params *DeleteAutoTagParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteAutoTagNoContent, error)

	GetAutoTag(params *GetAutoTagParams, authInfo runtime.ClientAuthInfoWriter) (*GetAutoTagOK, error)

	GetAutoTags(params *GetAutoTagsParams, authInfo runtime.ClientAuthInfoWriter) (*GetAutoTagsOK, error)

	UpdateAutoTag(params *UpdateAutoTagParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateAutoTagCreated, *UpdateAutoTagNoContent, error)

	ValidateCreateAutoTag(params *ValidateCreateAutoTagParams, authInfo runtime.ClientAuthInfoWriter) (*ValidateCreateAutoTagNoContent, error)

	ValidateUpdateAutoTag(params *ValidateUpdateAutoTagParams, authInfo runtime.ClientAuthInfoWriter) (*ValidateUpdateAutoTagNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateAutoTag creates a new auto tag

  The body must not provide an ID. An ID is assigned automatically by the Dynatrace server.
*/
func (a *Client) CreateAutoTag(params *CreateAutoTagParams, authInfo runtime.ClientAuthInfoWriter) (*CreateAutoTagCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAutoTagParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createAutoTag",
		Method:             "POST",
		PathPattern:        "/autoTags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateAutoTagReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateAutoTagCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createAutoTag: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteAutoTag deletes the specified auto tag
*/
func (a *Client) DeleteAutoTag(params *DeleteAutoTagParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteAutoTagNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAutoTagParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteAutoTag",
		Method:             "DELETE",
		PathPattern:        "/autoTags/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAutoTagReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteAutoTagNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteAutoTag: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAutoTag gets the properties of the specified auto tag
*/
func (a *Client) GetAutoTag(params *GetAutoTagParams, authInfo runtime.ClientAuthInfoWriter) (*GetAutoTagOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAutoTagParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAutoTag",
		Method:             "GET",
		PathPattern:        "/autoTags/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAutoTagReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAutoTagOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAutoTag: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAutoTags lists all configured auto tags
*/
func (a *Client) GetAutoTags(params *GetAutoTagsParams, authInfo runtime.ClientAuthInfoWriter) (*GetAutoTagsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAutoTagsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAutoTags",
		Method:             "GET",
		PathPattern:        "/autoTags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAutoTagsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAutoTagsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAutoTags: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateAutoTag updates an existing auto tag

  If the auto-tag with the specified ID does not exist, a new auto-tag is created.
*/
func (a *Client) UpdateAutoTag(params *UpdateAutoTagParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateAutoTagCreated, *UpdateAutoTagNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateAutoTagParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateAutoTag",
		Method:             "PUT",
		PathPattern:        "/autoTags/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateAutoTagReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *UpdateAutoTagCreated:
		return value, nil, nil
	case *UpdateAutoTagNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for automatically_applied_tags: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ValidateCreateAutoTag validates the payload for the p o s t auto tags request
*/
func (a *Client) ValidateCreateAutoTag(params *ValidateCreateAutoTagParams, authInfo runtime.ClientAuthInfoWriter) (*ValidateCreateAutoTagNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateCreateAutoTagParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "validateCreateAutoTag",
		Method:             "POST",
		PathPattern:        "/autoTags/validator",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ValidateCreateAutoTagReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ValidateCreateAutoTagNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for validateCreateAutoTag: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ValidateUpdateAutoTag validates the payload for the p u t auto tags id request
*/
func (a *Client) ValidateUpdateAutoTag(params *ValidateUpdateAutoTagParams, authInfo runtime.ClientAuthInfoWriter) (*ValidateUpdateAutoTagNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateUpdateAutoTagParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "validateUpdateAutoTag",
		Method:             "POST",
		PathPattern:        "/autoTags/{id}/validator",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ValidateUpdateAutoTagReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ValidateUpdateAutoTagNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for validateUpdateAutoTag: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}