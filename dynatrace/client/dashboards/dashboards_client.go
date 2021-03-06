// Code generated by go-swagger; DO NOT EDIT.

package dashboards

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new dashboards API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for dashboards API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateDashboard(params *CreateDashboardParams, authInfo runtime.ClientAuthInfoWriter) (*CreateDashboardCreated, error)

	DeleteDashboard(params *DeleteDashboardParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteDashboardNoContent, error)

	GetDashboard(params *GetDashboardParams, authInfo runtime.ClientAuthInfoWriter) (*GetDashboardOK, error)

	GetDashboards(params *GetDashboardsParams, authInfo runtime.ClientAuthInfoWriter) (*GetDashboardsOK, error)

	UpdateDashboard(params *UpdateDashboardParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateDashboardCreated, *UpdateDashboardNoContent, error)

	ValidateCreateDashboard(params *ValidateCreateDashboardParams, authInfo runtime.ClientAuthInfoWriter) (*ValidateCreateDashboardNoContent, error)

	ValidateUpdateDashboard(params *ValidateUpdateDashboardParams, authInfo runtime.ClientAuthInfoWriter) (*ValidateUpdateDashboardNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateDashboard creates a dashboard pipe maturity e a r l y a d o p t e r

  The body must not provide an ID. An ID is assigned automatically by the Dynatrace server.
*/
func (a *Client) CreateDashboard(params *CreateDashboardParams, authInfo runtime.ClientAuthInfoWriter) (*CreateDashboardCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDashboardParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createDashboard",
		Method:             "POST",
		PathPattern:        "/dashboards",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateDashboardReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateDashboardCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createDashboard: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteDashboard deletes the specified dashboard pipe maturity e a r l y a d o p t e r
*/
func (a *Client) DeleteDashboard(params *DeleteDashboardParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteDashboardNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDashboardParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteDashboard",
		Method:             "DELETE",
		PathPattern:        "/dashboards/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteDashboardReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteDashboardNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteDashboard: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDashboard gets the properties of the specified dashboard pipe maturity e a r l y a d o p t e r
*/
func (a *Client) GetDashboard(params *GetDashboardParams, authInfo runtime.ClientAuthInfoWriter) (*GetDashboardOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDashboardParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDashboard",
		Method:             "GET",
		PathPattern:        "/dashboards/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDashboardReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDashboardOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDashboard: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDashboards lists all dashboards of the environment pipe maturity e a r l y a d o p t e r
*/
func (a *Client) GetDashboards(params *GetDashboardsParams, authInfo runtime.ClientAuthInfoWriter) (*GetDashboardsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDashboardsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDashboards",
		Method:             "GET",
		PathPattern:        "/dashboards",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDashboardsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDashboardsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDashboards: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateDashboard updates or creates a dashboard pipe maturity e a r l y a d o p t e r
*/
func (a *Client) UpdateDashboard(params *UpdateDashboardParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateDashboardCreated, *UpdateDashboardNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDashboardParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateDashboard",
		Method:             "PUT",
		PathPattern:        "/dashboards/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateDashboardReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *UpdateDashboardCreated:
		return value, nil, nil
	case *UpdateDashboardNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dashboards: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ValidateCreateDashboard validates the payload for the p o s t dashboards request pipe maturity e a r l y a d o p t e r

  The body must not provide an ID.
*/
func (a *Client) ValidateCreateDashboard(params *ValidateCreateDashboardParams, authInfo runtime.ClientAuthInfoWriter) (*ValidateCreateDashboardNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateCreateDashboardParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "validateCreateDashboard",
		Method:             "POST",
		PathPattern:        "/dashboards/validator",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ValidateCreateDashboardReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ValidateCreateDashboardNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for validateCreateDashboard: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ValidateUpdateDashboard validates the payload for the p u t dashboards id request pipe maturity e a r l y a d o p t e r
*/
func (a *Client) ValidateUpdateDashboard(params *ValidateUpdateDashboardParams, authInfo runtime.ClientAuthInfoWriter) (*ValidateUpdateDashboardNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateUpdateDashboardParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "validateUpdateDashboard",
		Method:             "POST",
		PathPattern:        "/dashboards/{id}/validator",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ValidateUpdateDashboardReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ValidateUpdateDashboardNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for validateUpdateDashboard: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
