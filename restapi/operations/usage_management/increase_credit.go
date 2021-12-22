// Code generated by go-swagger; DO NOT EDIT.

package usage_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// IncreaseCreditHandlerFunc turns a function with the right signature into a increase credit handler
type IncreaseCreditHandlerFunc func(IncreaseCreditParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn IncreaseCreditHandlerFunc) Handle(params IncreaseCreditParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// IncreaseCreditHandler interface for that can handle valid increase credit params
type IncreaseCreditHandler interface {
	Handle(IncreaseCreditParams, interface{}) middleware.Responder
}

// NewIncreaseCredit creates a new http.Handler for the increase credit operation
func NewIncreaseCredit(ctx *middleware.Context, handler IncreaseCreditHandler) *IncreaseCredit {
	return &IncreaseCredit{Context: ctx, Handler: handler}
}

/*IncreaseCredit swagger:route POST /accounting/{id}/manage/{medium}/increase/{amount} usageManagement increaseCredit

Insert a new reseller in the system.

*/
type IncreaseCredit struct {
	Context *middleware.Context
	Handler IncreaseCreditHandler
}

func (o *IncreaseCredit) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewIncreaseCreditParams()

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
