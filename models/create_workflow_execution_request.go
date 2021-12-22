// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateWorkflowExecutionRequest create workflow execution request
//
// swagger:model CreateWorkflowExecutionRequest
type CreateWorkflowExecutionRequest struct {

	// input files
	InputFiles []*InputFile `json:"inputFiles"`

	// input parameters
	InputParameters map[string]interface{} `json:"inputParameters,omitempty"`

	// is batch job
	IsBatchJob bool `json:"isBatchJob,omitempty"`

	// is cron job
	IsCronJob bool `json:"isCronJob,omitempty"`

	// is scheduled job
	IsScheduledJob bool `json:"isScheduledJob,omitempty"`

	// workflow template ID
	WorkflowTemplateID string `json:"workflowTemplateID,omitempty"`
}

// Validate validates this create workflow execution request
func (m *CreateWorkflowExecutionRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInputFiles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateWorkflowExecutionRequest) validateInputFiles(formats strfmt.Registry) error {

	if swag.IsZero(m.InputFiles) { // not required
		return nil
	}

	for i := 0; i < len(m.InputFiles); i++ {
		if swag.IsZero(m.InputFiles[i]) { // not required
			continue
		}

		if m.InputFiles[i] != nil {
			if err := m.InputFiles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inputFiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateWorkflowExecutionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateWorkflowExecutionRequest) UnmarshalBinary(b []byte) error {
	var res CreateWorkflowExecutionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
