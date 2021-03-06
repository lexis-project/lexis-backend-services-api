// Code generated by go-swagger; DO NOT EDIT.

package data_set_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

// NewEncryptParams creates a new EncryptParams object
// no default values defined in spec.
func NewEncryptParams() EncryptParams {

	return EncryptParams{}
}

// EncryptParams contains all the bound params for the encrypt operation
// typically these are obtained from a http.Request
//
// swagger:parameters Encrypt
type EncryptParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*parameters
	  Required: true
	  In: body
	*/
	Parameters EncryptBody
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewEncryptParams() beforehand.
func (o *EncryptParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body EncryptBody
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("parameters", "body", ""))
			} else {
				res = append(res, errors.NewParseError("parameters", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Parameters = body
			}
		}
	} else {
		res = append(res, errors.Required("parameters", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
