// Code generated by go-swagger; DO NOT EDIT.

package workflow_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateWorkflowExecutionsHandlerFunc turns a function with the right signature into a create workflow executions handler
type CreateWorkflowExecutionsHandlerFunc func(CreateWorkflowExecutionsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateWorkflowExecutionsHandlerFunc) Handle(params CreateWorkflowExecutionsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CreateWorkflowExecutionsHandler interface for that can handle valid create workflow executions params
type CreateWorkflowExecutionsHandler interface {
	Handle(CreateWorkflowExecutionsParams, interface{}) middleware.Responder
}

// NewCreateWorkflowExecutions creates a new http.Handler for the create workflow executions operation
func NewCreateWorkflowExecutions(ctx *middleware.Context, handler CreateWorkflowExecutionsHandler) *CreateWorkflowExecutions {
	return &CreateWorkflowExecutions{Context: ctx, Handler: handler}
}

/*CreateWorkflowExecutions swagger:route POST /workflow/{workflowId}/executions workflowManagement createWorkflowExecutions

Create a batch of LEXIS Workflow Executions from a LEXIS Workflow

Create a batch of LEXIS WorkflowExecutions from a LEXIS Workflow given an array of WorkflowExecution parameters

*/
type CreateWorkflowExecutions struct {
	Context *middleware.Context
	Handler CreateWorkflowExecutionsHandler
}

func (o *CreateWorkflowExecutions) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateWorkflowExecutionsParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
