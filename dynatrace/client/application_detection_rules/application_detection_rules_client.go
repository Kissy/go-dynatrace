// Code generated by go-swagger; DO NOT EDIT.

package application_detection_rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new application detection rules API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for application detection rules API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateApplicationDetectionRule(params *CreateApplicationDetectionRuleParams, authInfo runtime.ClientAuthInfoWriter) (*CreateApplicationDetectionRuleCreated, error)

	DeleteApplicationDetectionRule(params *DeleteApplicationDetectionRuleParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteApplicationDetectionRuleNoContent, error)

	GetApplicationDetectionRule(params *GetApplicationDetectionRuleParams, authInfo runtime.ClientAuthInfoWriter) (*GetApplicationDetectionRuleOK, error)

	GetApplicationDetectionRules(params *GetApplicationDetectionRulesParams, authInfo runtime.ClientAuthInfoWriter) (*GetApplicationDetectionRulesOK, error)

	UpdateApplicationDetectionRule(params *UpdateApplicationDetectionRuleParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateApplicationDetectionRuleCreated, *UpdateApplicationDetectionRuleNoContent, error)

	UpdateApplicationDetectionRulesOrder(params *UpdateApplicationDetectionRulesOrderParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateApplicationDetectionRulesOrderNoContent, error)

	ValidateCreateApplicationDetectionRule(params *ValidateCreateApplicationDetectionRuleParams, authInfo runtime.ClientAuthInfoWriter) (*ValidateCreateApplicationDetectionRuleNoContent, error)

	ValidateUpdateApplicationDetectionRule(params *ValidateUpdateApplicationDetectionRuleParams, authInfo runtime.ClientAuthInfoWriter) (*ValidateUpdateApplicationDetectionRuleNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateApplicationDetectionRule creates a new application detection rule

  The body must not provide an ID. An ID is assigned automatically by the Dynatrace server.

You can only create detection rules for an existing application. If you need to create a rule for an application that doesn't exist yet, [create an application first](https://www.dynatrace.com/support/help/shortlink/api-config-web-app-post-web-app) and then configure detection rules for it.
*/
func (a *Client) CreateApplicationDetectionRule(params *CreateApplicationDetectionRuleParams, authInfo runtime.ClientAuthInfoWriter) (*CreateApplicationDetectionRuleCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateApplicationDetectionRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createApplicationDetectionRule",
		Method:             "POST",
		PathPattern:        "/applicationDetectionRules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateApplicationDetectionRuleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateApplicationDetectionRuleCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createApplicationDetectionRule: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteApplicationDetectionRule deletes the specified application detection rule
*/
func (a *Client) DeleteApplicationDetectionRule(params *DeleteApplicationDetectionRuleParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteApplicationDetectionRuleNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteApplicationDetectionRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteApplicationDetectionRule",
		Method:             "DELETE",
		PathPattern:        "/applicationDetectionRules/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteApplicationDetectionRuleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteApplicationDetectionRuleNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteApplicationDetectionRule: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetApplicationDetectionRule gets the parameters of the specified application detection rule
*/
func (a *Client) GetApplicationDetectionRule(params *GetApplicationDetectionRuleParams, authInfo runtime.ClientAuthInfoWriter) (*GetApplicationDetectionRuleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetApplicationDetectionRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getApplicationDetectionRule",
		Method:             "GET",
		PathPattern:        "/applicationDetectionRules/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetApplicationDetectionRuleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetApplicationDetectionRuleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getApplicationDetectionRule: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetApplicationDetectionRules lists all available application detection rules
*/
func (a *Client) GetApplicationDetectionRules(params *GetApplicationDetectionRulesParams, authInfo runtime.ClientAuthInfoWriter) (*GetApplicationDetectionRulesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetApplicationDetectionRulesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getApplicationDetectionRules",
		Method:             "GET",
		PathPattern:        "/applicationDetectionRules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetApplicationDetectionRulesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetApplicationDetectionRulesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getApplicationDetectionRules: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateApplicationDetectionRule updates the specified application detection rule

  If the application detection rule with the specified ID doesn't exist, a new application is created and appended to the end of the rule list.

If the **order** parameter is set for an existing rule, the request uses this value. Otherwise it keeps the existing order of rules.
*/
func (a *Client) UpdateApplicationDetectionRule(params *UpdateApplicationDetectionRuleParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateApplicationDetectionRuleCreated, *UpdateApplicationDetectionRuleNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateApplicationDetectionRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateApplicationDetectionRule",
		Method:             "PUT",
		PathPattern:        "/applicationDetectionRules/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateApplicationDetectionRuleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *UpdateApplicationDetectionRuleCreated:
		return value, nil, nil
	case *UpdateApplicationDetectionRuleNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for application_detection_rules: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateApplicationDetectionRulesOrder reorders the application detection rules

  This request reorders the application detection rules according to the submitted list of IDs. Application detection rules not present in the body of the request will retain their relative ordering but are placed *after* all those present in the request.
*/
func (a *Client) UpdateApplicationDetectionRulesOrder(params *UpdateApplicationDetectionRulesOrderParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateApplicationDetectionRulesOrderNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateApplicationDetectionRulesOrderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateApplicationDetectionRulesOrder",
		Method:             "PUT",
		PathPattern:        "/applicationDetectionRules/order",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateApplicationDetectionRulesOrderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateApplicationDetectionRulesOrderNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateApplicationDetectionRulesOrder: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ValidateCreateApplicationDetectionRule validates the payload for the p o s t application detection request
*/
func (a *Client) ValidateCreateApplicationDetectionRule(params *ValidateCreateApplicationDetectionRuleParams, authInfo runtime.ClientAuthInfoWriter) (*ValidateCreateApplicationDetectionRuleNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateCreateApplicationDetectionRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "validateCreateApplicationDetectionRule",
		Method:             "POST",
		PathPattern:        "/applicationDetectionRules/validator",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ValidateCreateApplicationDetectionRuleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ValidateCreateApplicationDetectionRuleNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for validateCreateApplicationDetectionRule: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ValidateUpdateApplicationDetectionRule validates the payload for the p u t application detection id request
*/
func (a *Client) ValidateUpdateApplicationDetectionRule(params *ValidateUpdateApplicationDetectionRuleParams, authInfo runtime.ClientAuthInfoWriter) (*ValidateUpdateApplicationDetectionRuleNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateUpdateApplicationDetectionRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "validateUpdateApplicationDetectionRule",
		Method:             "POST",
		PathPattern:        "/applicationDetectionRules/{id}/validator",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ValidateUpdateApplicationDetectionRuleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ValidateUpdateApplicationDetectionRuleNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for validateUpdateApplicationDetectionRule: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
