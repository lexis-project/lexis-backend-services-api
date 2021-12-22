// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WorkflowRequest workflow request
//
// swagger:model WorkflowRequest
type WorkflowRequest struct {

	// description
	// Required: true
	Description *string `json:"description"`

	// project ID
	// Required: true
	ProjectID *string `json:"projectID"`

	// workflow ID
	WorkflowID string `json:"workflowID,omitempty"`

	// workflow name
	// Required: true
	WorkflowName *string `json:"workflowName"`

	// workflow template ID
	// Required: true
	WorkflowTemplateID *string `json:"workflowTemplateID"`
}

// Validate validates this workflow request
func (m *WorkflowRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkflowName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkflowTemplateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkflowRequest) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *WorkflowRequest) validateProjectID(formats strfmt.Registry) error {

	if err := validate.Required("projectID", "body", m.ProjectID); err != nil {
		return err
	}

	return nil
}

func (m *WorkflowRequest) validateWorkflowName(formats strfmt.Registry) error {

	if err := validate.Required("workflowName", "body", m.WorkflowName); err != nil {
		return err
	}

	return nil
}

func (m *WorkflowRequest) validateWorkflowTemplateID(formats strfmt.Registry) error {

	if err := validate.Required("workflowTemplateID", "body", m.WorkflowTemplateID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowRequest) UnmarshalBinary(b []byte) error {
	var res WorkflowRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}