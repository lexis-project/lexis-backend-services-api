// Code generated by go-swagger; DO NOT EDIT.

package workflow_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/lexis-project/lexis-backend-services-api.git/models"
)

// NewCreateWorkflowExecutionsParams creates a new CreateWorkflowExecutionsParams object
// with the default values initialized.
func NewCreateWorkflowExecutionsParams() *CreateWorkflowExecutionsParams {
	var ()
	return &CreateWorkflowExecutionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateWorkflowExecutionsParamsWithTimeout creates a new CreateWorkflowExecutionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateWorkflowExecutionsParamsWithTimeout(timeout time.Duration) *CreateWorkflowExecutionsParams {
	var ()
	return &CreateWorkflowExecutionsParams{

		timeout: timeout,
	}
}

// NewCreateWorkflowExecutionsParamsWithContext creates a new CreateWorkflowExecutionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateWorkflowExecutionsParamsWithContext(ctx context.Context) *CreateWorkflowExecutionsParams {
	var ()
	return &CreateWorkflowExecutionsParams{

		Context: ctx,
	}
}

// NewCreateWorkflowExecutionsParamsWithHTTPClient creates a new CreateWorkflowExecutionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateWorkflowExecutionsParamsWithHTTPClient(client *http.Client) *CreateWorkflowExecutionsParams {
	var ()
	return &CreateWorkflowExecutionsParams{
		HTTPClient: client,
	}
}

/*CreateWorkflowExecutionsParams contains all the parameters to send to the API endpoint
for the create workflow executions operation typically these are written to a http.Request
*/
type CreateWorkflowExecutionsParams struct {

	/*WorkflowExecutionRequests
	  Create LEXIS Workflow Execution requests

	*/
	WorkflowExecutionRequests []*models.CreateWorkflowExecutionRequest
	/*WorkflowID
	  WorkflowId for WorkflowExecution to be created.

	*/
	WorkflowID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create workflow executions params
func (o *CreateWorkflowExecutionsParams) WithTimeout(timeout time.Duration) *CreateWorkflowExecutionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create workflow executions params
func (o *CreateWorkflowExecutionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create workflow executions params
func (o *CreateWorkflowExecutionsParams) WithContext(ctx context.Context) *CreateWorkflowExecutionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create workflow executions params
func (o *CreateWorkflowExecutionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create workflow executions params
func (o *CreateWorkflowExecutionsParams) WithHTTPClient(client *http.Client) *CreateWorkflowExecutionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create workflow executions params
func (o *CreateWorkflowExecutionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithWorkflowExecutionRequests adds the workflowExecutionRequests to the create workflow executions params
func (o *CreateWorkflowExecutionsParams) WithWorkflowExecutionRequests(workflowExecutionRequests []*models.CreateWorkflowExecutionRequest) *CreateWorkflowExecutionsParams {
	o.SetWorkflowExecutionRequests(workflowExecutionRequests)
	return o
}

// SetWorkflowExecutionRequests adds the workflowExecutionRequests to the create workflow executions params
func (o *CreateWorkflowExecutionsParams) SetWorkflowExecutionRequests(workflowExecutionRequests []*models.CreateWorkflowExecutionRequest) {
	o.WorkflowExecutionRequests = workflowExecutionRequests
}

// WithWorkflowID adds the workflowID to the create workflow executions params
func (o *CreateWorkflowExecutionsParams) WithWorkflowID(workflowID string) *CreateWorkflowExecutionsParams {
	o.SetWorkflowID(workflowID)
	return o
}

// SetWorkflowID adds the workflowId to the create workflow executions params
func (o *CreateWorkflowExecutionsParams) SetWorkflowID(workflowID string) {
	o.WorkflowID = workflowID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateWorkflowExecutionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.WorkflowExecutionRequests != nil {
		if err := r.SetBodyParam(o.WorkflowExecutionRequests); err != nil {
			return err
		}
	}

	// path param workflowId
	if err := r.SetPathParam("workflowId", o.WorkflowID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
