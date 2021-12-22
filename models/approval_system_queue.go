// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ApprovalSystemQueue approval system queue
//
// swagger:model ApprovalSystemQueue
type ApprovalSystemQueue struct {

	// cores per node
	CoresPerNode *int64 `json:"CoresPerNode,omitempty"`

	// description
	Description string `json:"Description,omitempty"`

	// ID
	ID *int64 `json:"ID,omitempty"`

	// max wall time
	MaxWallTime *int64 `json:"MaxWallTime,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// number of nodes
	NumberOfNodes *int64 `json:"NumberOfNodes,omitempty"`

	// type
	Type string `json:"Type,omitempty"`
}

// Validate validates this approval system queue
func (m *ApprovalSystemQueue) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ApprovalSystemQueue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApprovalSystemQueue) UnmarshalBinary(b []byte) error {
	var res ApprovalSystemQueue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
