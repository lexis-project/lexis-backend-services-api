// Code generated by go-swagger; DO NOT EDIT.

package workflow_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	"github.com/lexis-project/lexis-backend-services-api.git/models"
)

// NewCreateWorkflowExecutionParams creates a new CreateWorkflowExecutionParams object
// no default values defined in spec.
func NewCreateWorkflowExecutionParams() CreateWorkflowExecutionParams {

	return CreateWorkflowExecutionParams{}
}

// CreateWorkflowExecutionParams contains all the bound params for the create workflow execution operation
// typically these are obtained from a http.Request
//
// swagger:parameters createWorkflowExecution
type CreateWorkflowExecutionParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Create LEXIS Workflow Execution request
	  In: body
	*/
	WorkflowExecutionRequest *models.CreateWorkflowExecutionRequest
	/*WorkflowId for WorkflowExecution to be created.
	  Required: true
	  In: path
	*/
	WorkflowID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCreateWorkflowExecutionParams() beforehand.
func (o *CreateWorkflowExecutionParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.CreateWorkflowExecutionRequest
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("workflowExecutionRequest", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.WorkflowExecutionRequest = &body
			}
		}
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

// bindWorkflowID binds and validates parameter WorkflowID from path.
func (o *CreateWorkflowExecutionParams) bindWorkflowID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.WorkflowID = raw

	return nil
}