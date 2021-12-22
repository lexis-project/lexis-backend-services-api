// Code generated by go-swagger; DO NOT EDIT.

package data_set_management

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

// DecompressHandlerFunc turns a function with the right signature into a decompress handler
type DecompressHandlerFunc func(DecompressParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DecompressHandlerFunc) Handle(params DecompressParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DecompressHandler interface for that can handle valid decompress params
type DecompressHandler interface {
	Handle(DecompressParams, interface{}) middleware.Responder
}

// NewDecompress creates a new http.Handler for the decompress operation
func NewDecompress(ctx *middleware.Context, handler DecompressHandler) *Decompress {
	return &Decompress{Context: ctx, Handler: handler}
}

/*Decompress swagger:route POST /dataset/encryption/decompress dataSetManagement decompress

Decompress a dataset or subdataset (by enqueuing the request for latter processing)

Decompress a dataset or subdataset (by enqueuing the request for latter processing)
If you have a tuple [project, access, internalID] and the current user, the corresponding path should be calculated by calculating the md5 hash of the project, and then:
 - For public datasets: "public/proj"+hash+"/"+internalID
 - For user datasets: "user/proj"+hash+"/"+user+"/"+internalID
 - For project datasets: "project/proj"+hash+"/"+internalID


*/
type Decompress struct {
	Context *middleware.Context
	Handler DecompressHandler
}

func (o *Decompress) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDecompressParams()

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

// DecompressBody decompress body
//
// swagger:model DecompressBody
type DecompressBody struct {

	// source path
	// Required: true
	SourcePath *string `json:"source_path"`

	// source system
	// Required: true
	SourceSystem *string `json:"source_system"`
}

// Validate validates this decompress body
func (o *DecompressBody) Validate(formats strfmt.Registry) error {
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

func (o *DecompressBody) validateSourcePath(formats strfmt.Registry) error {

	if err := validate.Required("parameters"+"."+"source_path", "body", o.SourcePath); err != nil {
		return err
	}

	return nil
}

func (o *DecompressBody) validateSourceSystem(formats strfmt.Registry) error {

	if err := validate.Required("parameters"+"."+"source_system", "body", o.SourceSystem); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DecompressBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DecompressBody) UnmarshalBinary(b []byte) error {
	var res DecompressBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
