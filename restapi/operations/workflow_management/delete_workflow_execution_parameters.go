// Code generated by go-swagger; DO NOT EDIT.

package workflow_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewDeleteWorkflowExecutionParams creates a new DeleteWorkflowExecutionParams object
// no default values defined in spec.
func NewDeleteWorkflowExecutionParams() DeleteWorkflowExecutionParams {

	return DeleteWorkflowExecutionParams{}
}

// DeleteWorkflowExecutionParams contains all the bound params for the delete workflow execution operation
// typically these are obtained from a http.Request
//
// swagger:parameters deleteWorkflowExecution
type DeleteWorkflowExecutionParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*WorkflowExecutionID for WorkflowExecution to be removed
	  Required: true
	  In: path
	*/
	WorkflowExecutionID string
	/*workflowID for LEXIS Workflow Execution to be removed
	  Required: true
	  In: path
	*/
	WorkflowID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteWorkflowExecutionParams() beforehand.
func (o *DeleteWorkflowExecutionParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rWorkflowExecutionID, rhkWorkflowExecutionID, _ := route.Params.GetOK("workflowExecutionId")
	if err := o.bindWorkflowExecutionID(rWorkflowExecutionID, rhkWorkflowExecutionID, route.Formats); err != nil {
		res = append(res, err)
	}

	rWorkflowID, rhkWorkflowID, _ := route.Params.GetOK("workflowId")
	if err := o.bindWorkflowID(rWorkflowID, rhkWorkflowID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindWorkflowExecutionID binds and validates parameter WorkflowExecutionID from path.
func (o *DeleteWorkflowExecutionParams) bindWorkflowExecutionID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.WorkflowExecutionID = raw

	return nil
}

// bindWorkflowID binds and validates parameter WorkflowID from path.
func (o *DeleteWorkflowExecutionParams) bindWorkflowID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.WorkflowID = raw

	return nil
}
