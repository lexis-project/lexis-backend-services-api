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

// CheckDecryptionStatusHandlerFunc turns a function with the right signature into a check decryption status handler
type CheckDecryptionStatusHandlerFunc func(CheckDecryptionStatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CheckDecryptionStatusHandlerFunc) Handle(params CheckDecryptionStatusParams) middleware.Responder {
	return fn(params)
}

// CheckDecryptionStatusHandler interface for that can handle valid check decryption status params
type CheckDecryptionStatusHandler interface {
	Handle(CheckDecryptionStatusParams) middleware.Responder
}

// NewCheckDecryptionStatus creates a new http.Handler for the check decryption status operation
func NewCheckDecryptionStatus(ctx *middleware.Context, handler CheckDecryptionStatusHandler) *CheckDecryptionStatus {
	return &CheckDecryptionStatus{Context: ctx, Handler: handler}
}

/*CheckDecryptionStatus swagger:route GET /dataset/encryption/decrypt/{request_id} dataSetManagement checkDecryptionStatus

Check the status of a decryption

Check the status of a decryption

*/
type CheckDecryptionStatus struct {
	Context *middleware.Context
	Handler CheckDecryptionStatusHandler
}

func (o *CheckDecryptionStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCheckDecryptionStatusParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// CheckDecryptionStatusOKBody check decryption status o k body
//
// swagger:model CheckDecryptionStatusOKBody
type CheckDecryptionStatusOKBody struct {

	// status
	// Required: true
	Status *string `json:"status"`

	// Single path for encryption endpoints
	TargetPath string `json:"target_path,omitempty"`
}

// Validate validates this check decryption status o k body
func (o *CheckDecryptionStatusOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CheckDecryptionStatusOKBody) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("checkDecryptionStatusOK"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CheckDecryptionStatusOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckDecryptionStatusOKBody) UnmarshalBinary(b []byte) error {
	var res CheckDecryptionStatusOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
