// Code generated by go-swagger; DO NOT EDIT.

package workflow_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CancelWorkflowExecutionHandlerFunc turns a function with the right signature into a cancel workflow execution handler
type CancelWorkflowExecutionHandlerFunc func(CancelWorkflowExecutionParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CancelWorkflowExecutionHandlerFunc) Handle(params CancelWorkflowExecutionParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CancelWorkflowExecutionHandler interface for that can handle valid cancel workflow execution params
type CancelWorkflowExecutionHandler interface {
	Handle(CancelWorkflowExecutionParams, interface{}) middleware.Responder
}

// NewCancelWorkflowExecution creates a new http.Handler for the cancel workflow execution operation
func NewCancelWorkflowExecution(ctx *middleware.Context, handler CancelWorkflowExecutionHandler) *CancelWorkflowExecution {
	return &CancelWorkflowExecution{Context: ctx, Handler: handler}
}

/*CancelWorkflowExecution swagger:route DELETE /workflow/{workflowId}/execution/{workflowExecutionId} workflowManagement cancelWorkflowExecution

Cancel a LEXIS Workflow Execution.

Cancels given LEXIS Workflow Execution.

*/
type CancelWorkflowExecution struct {
	Context *middleware.Context
	Handler CancelWorkflowExecutionHandler
}

func (o *CancelWorkflowExecution) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCancelWorkflowExecutionParams()

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
