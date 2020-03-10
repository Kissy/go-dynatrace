// Code generated by go-swagger; DO NOT EDIT.

package frequent_issue_detection

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new frequent issue detection API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for frequent issue detection API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetFrequentIssueDetection(params *GetFrequentIssueDetectionParams, authInfo runtime.ClientAuthInfoWriter) (*GetFrequentIssueDetectionOK, error)

	UpdateFrequentIssueDetection(params *UpdateFrequentIssueDetectionParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateFrequentIssueDetectionNoContent, error)

	ValidateFrequentIssueDetection(params *ValidateFrequentIssueDetectionParams, authInfo runtime.ClientAuthInfoWriter) (*ValidateFrequentIssueDetectionNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetFrequentIssueDetection gets the configuration of frequent issue detection
*/
func (a *Client) GetFrequentIssueDetection(params *GetFrequentIssueDetectionParams, authInfo runtime.ClientAuthInfoWriter) (*GetFrequentIssueDetectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFrequentIssueDetectionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFrequentIssueDetection",
		Method:             "GET",
		PathPattern:        "/frequentIssueDetection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFrequentIssueDetectionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFrequentIssueDetectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getFrequentIssueDetection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateFrequentIssueDetection updates the configuration of frequent issue detection
*/
func (a *Client) UpdateFrequentIssueDetection(params *UpdateFrequentIssueDetectionParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateFrequentIssueDetectionNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateFrequentIssueDetectionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateFrequentIssueDetection",
		Method:             "PUT",
		PathPattern:        "/frequentIssueDetection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateFrequentIssueDetectionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateFrequentIssueDetectionNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateFrequentIssueDetection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ValidateFrequentIssueDetection validates the frequent issue detection configuration for the p u t frequent issue detection request
*/
func (a *Client) ValidateFrequentIssueDetection(params *ValidateFrequentIssueDetectionParams, authInfo runtime.ClientAuthInfoWriter) (*ValidateFrequentIssueDetectionNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateFrequentIssueDetectionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "validateFrequentIssueDetection",
		Method:             "POST",
		PathPattern:        "/frequentIssueDetection/validator",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ValidateFrequentIssueDetectionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ValidateFrequentIssueDetectionNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for validateFrequentIssueDetection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}