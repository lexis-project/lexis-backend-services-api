// Code generated by go-swagger; DO NOT EDIT.

package approval_system_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateApprovedResourceRequestHandlerFunc turns a function with the right signature into a create approved resource request handler
type CreateApprovedResourceRequestHandlerFunc func(CreateApprovedResourceRequestParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateApprovedResourceRequestHandlerFunc) Handle(params CreateApprovedResourceRequestParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CreateApprovedResourceRequestHandler interface for that can handle valid create approved resource request params
type CreateApprovedResourceRequestHandler interface {
	Handle(CreateApprovedResourceRequestParams, interface{}) middleware.Responder
}

// NewCreateApprovedResourceRequest creates a new http.Handler for the create approved resource request operation
func NewCreateApprovedResourceRequest(ctx *middleware.Context, handler CreateApprovedResourceRequestHandler) *CreateApprovedResourceRequest {
	return &CreateApprovedResourceRequest{Context: ctx, Handler: handler}
}

/*CreateApprovedResourceRequest swagger:route POST /approval_system/approvedResourceRequest approvalSystemManagement createApprovedResourceRequest

Create HPC approved resource request

Creates a new HPC approved resource request in approval system

*/
type CreateApprovedResourceRequest struct {
	Context *middleware.Context
	Handler CreateApprovedResourceRequestHandler
}

func (o *CreateApprovedResourceRequest) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateApprovedResourceRequestParams()

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
