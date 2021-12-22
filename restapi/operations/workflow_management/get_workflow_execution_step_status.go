// Code generated by go-swagger; DO NOT EDIT.

package workflow_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetWorkflowExecutionStepStatusHandlerFunc turns a function with the right signature into a get workflow execution step status handler
type GetWorkflowExecutionStepStatusHandlerFunc func(GetWorkflowExecutionStepStatusParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetWorkflowExecutionStepStatusHandlerFunc) Handle(params GetWorkflowExecutionStepStatusParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetWorkflowExecutionStepStatusHandler interface for that can handle valid get workflow execution step status params
type GetWorkflowExecutionStepStatusHandler interface {
	Handle(GetWorkflowExecutionStepStatusParams, interface{}) middleware.Responder
}

// NewGetWorkflowExecutionStepStatus creates a new http.Handler for the get workflow execution step status operation
func NewGetWorkflowExecutionStepStatus(ctx *middleware.Context, handler GetWorkflowExecutionStepStatusHandler) *GetWorkflowExecutionStepStatus {
	return &GetWorkflowExecutionStepStatus{Context: ctx, Handler: handler}
}

/*GetWorkflowExecutionStepStatus swagger:route GET /workflow/{workflowId}/execution/{workflowExecutionId}/status workflowManagement getWorkflowExecutionStepStatus

Returns detailed status of Lexis Workflow Execution and its tasks.

Return Task Status' on given LEXIS Workflow Execution

*/
type GetWorkflowExecutionStepStatus struct {
	Context *middleware.Context
	Handler GetWorkflowExecutionStepStatusHandler
}

func (o *GetWorkflowExecutionStepStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetWorkflowExecutionStepStatusParams()

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
