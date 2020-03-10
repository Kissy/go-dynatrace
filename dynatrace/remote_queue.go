// Code generated by go-swagger; DO NOT EDIT.

package dynatrace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RemoteQueue Define a local definition of a remote queue owned by another queue manager. The local definition can be made visible in one or more clusters.
// swagger:model RemoteQueue
type RemoteQueue struct {

	// The local definition of the remote queue is visible in these [clusters](https://www.ibm.com/support/knowledgecenter/en/SSFKSJ_7.5.0/com.ibm.mq.pro.doc/q002750_.htm). The queue manager must be part of these clusters.
	// Required: true
	// Max Items: 10000
	// Min Items: 0
	ClusterVisibility []string `json:"clusterVisibility"`

	// The name of the local definition of the remote queue.
	// Required: true
	// Max Length: 500
	// Min Length: 1
	LocalQueue *string `json:"localQueue"`

	// The name of the remote queue.
	// Required: true
	// Max Length: 500
	// Min Length: 1
	RemoteQueue *string `json:"remoteQueue"`

	// The name of the remote queue manager.
	// Required: true
	// Max Length: 500
	// Min Length: 1
	RemoteQueueManager *string `json:"remoteQueueManager"`
}

// Validate validates this remote queue
func (m *RemoteQueue) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterVisibility(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalQueue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemoteQueue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemoteQueueManager(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RemoteQueue) validateClusterVisibility(formats strfmt.Registry) error {

	if err := validate.Required("clusterVisibility", "body", m.ClusterVisibility); err != nil {
		return err
	}

	iClusterVisibilitySize := int64(len(m.ClusterVisibility))

	if err := validate.MinItems("clusterVisibility", "body", iClusterVisibilitySize, 0); err != nil {
		return err
	}

	if err := validate.MaxItems("clusterVisibility", "body", iClusterVisibilitySize, 10000); err != nil {
		return err
	}

	return nil
}

func (m *RemoteQueue) validateLocalQueue(formats strfmt.Registry) error {

	if err := validate.Required("localQueue", "body", m.LocalQueue); err != nil {
		return err
	}

	if err := validate.MinLength("localQueue", "body", string(*m.LocalQueue), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("localQueue", "body", string(*m.LocalQueue), 500); err != nil {
		return err
	}

	return nil
}

func (m *RemoteQueue) validateRemoteQueue(formats strfmt.Registry) error {

	if err := validate.Required("remoteQueue", "body", m.RemoteQueue); err != nil {
		return err
	}

	if err := validate.MinLength("remoteQueue", "body", string(*m.RemoteQueue), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("remoteQueue", "body", string(*m.RemoteQueue), 500); err != nil {
		return err
	}

	return nil
}

func (m *RemoteQueue) validateRemoteQueueManager(formats strfmt.Registry) error {

	if err := validate.Required("remoteQueueManager", "body", m.RemoteQueueManager); err != nil {
		return err
	}

	if err := validate.MinLength("remoteQueueManager", "body", string(*m.RemoteQueueManager), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("remoteQueueManager", "body", string(*m.RemoteQueueManager), 500); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RemoteQueue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RemoteQueue) UnmarshalBinary(b []byte) error {
	var res RemoteQueue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}