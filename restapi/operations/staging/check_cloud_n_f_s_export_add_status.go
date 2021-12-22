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

// CheckCloudNFSExportAddStatusHandlerFunc turns a function with the right signature into a check cloud n f s export add status handler
type CheckCloudNFSExportAddStatusHandlerFunc func(CheckCloudNFSExportAddStatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CheckCloudNFSExportAddStatusHandlerFunc) Handle(params CheckCloudNFSExportAddStatusParams) middleware.Responder {
	return fn(params)
}

// CheckCloudNFSExportAddStatusHandler interface for that can handle valid check cloud n f s export add status params
type CheckCloudNFSExportAddStatusHandler interface {
	Handle(CheckCloudNFSExportAddStatusParams) middleware.Responder
}

// NewCheckCloudNFSExportAddStatus creates a new http.Handler for the check cloud n f s export add status operation
func NewCheckCloudNFSExportAddStatus(ctx *middleware.Context, handler CheckCloudNFSExportAddStatusHandler) *CheckCloudNFSExportAddStatus {
	return &CheckCloudNFSExportAddStatus{Context: ctx, Handler: handler}
}

/*CheckCloudNFSExportAddStatus swagger:route GET /dataset/cloud/add/{param} staging checkCloudNFSExportAddStatus

Check the status of a nfs export add request for the cloud

Check the status of a nfs export add request for the cloud

*/
type CheckCloudNFSExportAddStatus struct {
	Context *middleware.Context
	Handler CheckCloudNFSExportAddStatusHandler
}

func (o *CheckCloudNFSExportAddStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCheckCloudNFSExportAddStatusParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// CheckCloudNFSExportAddStatusOKBody check cloud n f s export add status o k body
//
// swagger:model CheckCloudNFSExportAddStatusOKBody
type CheckCloudNFSExportAddStatusOKBody struct {

	// status
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this check cloud n f s export add status o k body
func (o *CheckCloudNFSExportAddStatusOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CheckCloudNFSExportAddStatusOKBody) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("checkCloudNFSExportAddStatusOK"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CheckCloudNFSExportAddStatusOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckCloudNFSExportAddStatusOKBody) UnmarshalBinary(b []byte) error {
	var res CheckCloudNFSExportAddStatusOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
