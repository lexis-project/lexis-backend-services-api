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

// EncryptHandlerFunc turns a function with the right signature into a encrypt handler
type EncryptHandlerFunc func(EncryptParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn EncryptHandlerFunc) Handle(params EncryptParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// EncryptHandler interface for that can handle valid encrypt params
type EncryptHandler interface {
	Handle(EncryptParams, interface{}) middleware.Responder
}

// NewEncrypt creates a new http.Handler for the encrypt operation
func NewEncrypt(ctx *middleware.Context, handler EncryptHandler) *Encrypt {
	return &Encrypt{Context: ctx, Handler: handler}
}

/*Encrypt swagger:route POST /dataset/encryption/encrypt dataSetManagement encrypt

Encrypt a dataset or subdataset (by enqueuing the request for latter processing)

Encrypt a dataset or subdataset (by enqueuing the request for latter processing)
If you have a tuple [project, access, internalID] and the current user, the corresponding path should be calculated by calculating the md5 hash of the project, and then:
 - For public datasets: "public/proj"+hash+"/"+internalID
 - For user datasets: "user/proj"+hash+"/"+user+"/"+internalID
 - For project datasets: "project/proj"+hash+"/"+internalID


*/
type Encrypt struct {
	Context *middleware.Context
	Handler EncryptHandler
}

func (o *Encrypt) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewEncryptParams()

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

// EncryptBody encrypt body
//
// swagger:model EncryptBody
type EncryptBody struct {

	// project
	// Required: true
	Project *string `json:"project"`

	// source path
	// Required: true
	SourcePath *string `json:"source_path"`

	// source system
	// Required: true
	SourceSystem *string `json:"source_system"`
}

// Validate validates this encrypt body
func (o *EncryptBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateProject(formats); err != nil {
		res = append(res, err)
	}

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

func (o *EncryptBody) validateProject(formats strfmt.Registry) error {

	if err := validate.Required("parameters"+"."+"project", "body", o.Project); err != nil {
		return err
	}

	return nil
}

func (o *EncryptBody) validateSourcePath(formats strfmt.Registry) error {

	if err := validate.Required("parameters"+"."+"source_path", "body", o.SourcePath); err != nil {
		return err
	}

	return nil
}

func (o *EncryptBody) validateSourceSystem(formats strfmt.Registry) error {

	if err := validate.Required("parameters"+"."+"source_system", "body", o.SourceSystem); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *EncryptBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *EncryptBody) UnmarshalBinary(b []byte) error {
	var res EncryptBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
