// Code generated by go-swagger; DO NOT EDIT.

package staging

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DuplicateHandlerFunc turns a function with the right signature into a duplicate handler
type DuplicateHandlerFunc func(DuplicateParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DuplicateHandlerFunc) Handle(params DuplicateParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DuplicateHandler interface for that can handle valid duplicate params
type DuplicateHandler interface {
	Handle(DuplicateParams, interface{}) middleware.Responder
}

// NewDuplicate creates a new http.Handler for the duplicate operation
func NewDuplicate(ctx *middleware.Context, handler DuplicateHandler) *Duplicate {
	return &Duplicate{Context: ctx, Handler: handler}
}

/*Duplicate swagger:route POST /dataset/duplicate staging dataSetManagement duplicate

Duplicate a dataset or subdataset (by enqueuing the request for latter processing)

Duplicate a dataset or subdataset (by enqueuing the request for latter processing)

*/
type Duplicate struct {
	Context *middleware.Context
	Handler DuplicateHandler
}

func (o *Duplicate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDuplicateParams()

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

// DuplicateBody duplicate body
//
// swagger:model DuplicateBody
type DuplicateBody struct {

	// source path
	// Required: true
	SourcePath *string `json:"source_path"`

	// Possible values are: "lrz_iRODS", "it4i_iRODS"
	//
	// Required: true
	SourceSystem *string `json:"source_system"`

	// target path
	TargetPath string `json:"target_path,omitempty"`

	// Possible values are: "lrz_iRODS", "it4i_iRODS"
	//
	TargetSystem string `json:"target_system,omitempty"`

	// Title of the new dataset (if not provided, the new title will
	// be "Copy of " and the old dataset title).
	//
	Title string `json:"title,omitempty"`
}

// Validate validates this duplicate body
func (o *DuplicateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSourcePath(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSourceSystem(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DuplicateBody) validateSourcePath(formats strfmt.Registry) error {

	if err := validate.Required("parameters"+"."+"source_path", "body", o.SourcePath); err != nil {
		return err
	}

	return nil
}

func (o *DuplicateBody) validateSourceSystem(formats strfmt.Registry) error {

	if err := validate.Required("parameters"+"."+"source_system", "body", o.SourceSystem); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DuplicateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DuplicateBody) UnmarshalBinary(b []byte) error {
	var res DuplicateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
