// Code generated by go-swagger; DO NOT EDIT.

package workflow_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteWorkflowExecutionHandlerFunc turns a function with the right signature into a delete workflow execution handler
type DeleteWorkflowExecutionHandlerFunc func(DeleteWorkflowExecutionParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteWorkflowExecutionHandlerFunc) Handle(params DeleteWorkflowExecutionParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteWorkflowExecutionHandler interface for that can handle valid delete workflow execution params
type DeleteWorkflowExecutionHandler interface {
	Handle(DeleteWorkflowExecutionParams, interface{}) middleware.Responder
}

// NewDeleteWorkflowExecution creates a new http.Handler for the delete workflow execution operation
func NewDeleteWorkflowExecution(ctx *middleware.Context, handler DeleteWorkflowExecutionHandler) *DeleteWorkflowExecution {
	return &DeleteWorkflowExecution{Context: ctx, Handler: handler}
}

/*DeleteWorkflowExecution swagger:route POST /workflow/{workflowId}/execution/{workflowExecutionId}/remove workflowManagement deleteWorkflowExecution

Delete a LEXIS Workflow Execution.

Removes given LEXIS Workflow Execution.

*/
type DeleteWorkflowExecution struct {
	Context *middleware.Context
	Handler DeleteWorkflowExecutionHandler
}

func (o *DeleteWorkflowExecution) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteWorkflowExecutionParams()

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
