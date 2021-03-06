// Code generated by go-swagger; DO NOT EDIT.

package workflow_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateWorkflowExecutionHandlerFunc turns a function with the right signature into a create workflow execution handler
type CreateWorkflowExecutionHandlerFunc func(CreateWorkflowExecutionParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateWorkflowExecutionHandlerFunc) Handle(params CreateWorkflowExecutionParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CreateWorkflowExecutionHandler interface for that can handle valid create workflow execution params
type CreateWorkflowExecutionHandler interface {
	Handle(CreateWorkflowExecutionParams, interface{}) middleware.Responder
}

// NewCreateWorkflowExecution creates a new http.Handler for the create workflow execution operation
func NewCreateWorkflowExecution(ctx *middleware.Context, handler CreateWorkflowExecutionHandler) *CreateWorkflowExecution {
	return &CreateWorkflowExecution{Context: ctx, Handler: handler}
}

/*CreateWorkflowExecution swagger:route POST /workflow/{workflowId}/execution workflowManagement createWorkflowExecution

TODO: Needs implemented with TOSCA 1.3 Capabilitise. Create a new LEXIS Workflow Execution by providing remaining inputs

Creates a LEXIS workflow execution from a LEXIS workflow.
The name of the resulting workflow execution will be of the form `n-[projectShortName]` where `n` is the index of the workflow execution withing it's workflow.
The workflow execution name takes on the suffix `_cron` or `_scheduled` where it is a cron or scheduled job according to the request parameters.
For example, the workflow execution `2-dummyprj_cron` is the third workflow execution from it's workflow under the project `dummyprj` and it is a cron job.


*/
type CreateWorkflowExecution struct {
	Context *middleware.Context
	Handler CreateWorkflowExecutionHandler
}

func (o *CreateWorkflowExecution) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateWorkflowExecutionParams()

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
