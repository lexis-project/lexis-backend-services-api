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

// NewDecompressParams creates a new DecompressParams object
// no default values defined in spec.
func NewDecompressParams() DecompressParams {

	return DecompressParams{}
}

// DecompressParams contains all the bound params for the decompress operation
// typically these are obtained from a http.Request
//
// swagger:parameters Decompress
type DecompressParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*parameters
	  Required: true
	  In: body
	*/
	Parameters DecompressBody
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDecompressParams() beforehand.
func (o *DecompressParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body DecompressBody
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