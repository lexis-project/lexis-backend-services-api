// Code generated by go-swagger; DO NOT EDIT.

package workflow_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UploadWorkflowTemplateHandlerFunc turns a function with the right signature into a upload workflow template handler
type UploadWorkflowTemplateHandlerFunc func(UploadWorkflowTemplateParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn UploadWorkflowTemplateHandlerFunc) Handle(params UploadWorkflowTemplateParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// UploadWorkflowTemplateHandler interface for that can handle valid upload workflow template params
type UploadWorkflowTemplateHandler interface {
	Handle(UploadWorkflowTemplateParams, interface{}) middleware.Responder
}

// NewUploadWorkflowTemplate creates a new http.Handler for the upload workflow template operation
func NewUploadWorkflowTemplate(ctx *middleware.Context, handler UploadWorkflowTemplateHandler) *UploadWorkflowTemplate {
	return &UploadWorkflowTemplate{Context: ctx, Handler: handler}
}

/*UploadWorkflowTemplate swagger:route POST /workflow/template/upload workflowManagement uploadWorkflowTemplate

Create a new LEXIS Workflow Template on the system

If successful returns a rest response with the id of the created LEXIS workflow template. If not successful a rest response with an error content is returned.

*/
type UploadWorkflowTemplate struct {
	Context *middleware.Context
	Handler UploadWorkflowTemplateHandler
}

func (o *UploadWorkflowTemplate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUploadWorkflowTemplateParams()

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
