// Code generated by go-swagger; DO NOT EDIT.

package workflow_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteWorkflowHandlerFunc turns a function with the right signature into a delete workflow handler
type DeleteWorkflowHandlerFunc func(DeleteWorkflowParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteWorkflowHandlerFunc) Handle(params DeleteWorkflowParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteWorkflowHandler interface for that can handle valid delete workflow params
type DeleteWorkflowHandler interface {
	Handle(DeleteWorkflowParams, interface{}) middleware.Responder
}

// NewDeleteWorkflow creates a new http.Handler for the delete workflow operation
func NewDeleteWorkflow(ctx *middleware.Context, handler DeleteWorkflowHandler) *DeleteWorkflow {
	return &DeleteWorkflow{Context: ctx, Handler: handler}
}

/*DeleteWorkflow swagger:route DELETE /workflow/{workflowId} workflowManagement deleteWorkflow

Delete LEXIS Workflow on the system

Delete an already existing workflow on the system

*/
type DeleteWorkflow struct {
	Context *middleware.Context
	Handler DeleteWorkflowHandler
}

func (o *DeleteWorkflow) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteWorkflowParams()

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
