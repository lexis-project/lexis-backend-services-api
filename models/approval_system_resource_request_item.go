// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ApprovalSystemResourceRequestItem approval system resource request item
//
// swagger:model ApprovalSystemResourceRequestItem
type ApprovalSystemResourceRequestItem struct {

	// cluster ID
	ClusterID *int64 `json:"ClusterID,omitempty"`

	// cluster name
	ClusterName string `json:"ClusterName,omitempty"`

	// queue ID
	QueueID *int64 `json:"QueueID,omitempty"`

	// queue name
	QueueName string `json:"QueueName,omitempty"`
}

// Validate validates this approval system resource request item
func (m *ApprovalSystemResourceRequestItem) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ApprovalSystemResourceRequestItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApprovalSystemResourceRequestItem) UnmarshalBinary(b []byte) error {
	var res ApprovalSystemResourceRequestItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}