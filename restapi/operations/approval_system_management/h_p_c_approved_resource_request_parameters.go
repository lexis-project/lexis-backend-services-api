// Code generated by go-swagger; DO NOT EDIT.

package approval_system_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewHPCApprovedResourceRequestParams creates a new HPCApprovedResourceRequestParams object
// no default values defined in spec.
func NewHPCApprovedResourceRequestParams() HPCApprovedResourceRequestParams {

	return HPCApprovedResourceRequestParams{}
}

// HPCApprovedResourceRequestParams contains all the bound params for the h p c approved resource request operation
// typically these are obtained from a http.Request
//
// swagger:parameters HPCApprovedResourceRequest
type HPCApprovedResourceRequestParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*String ID of the HPC resource whose approved resource request to be obtained.
	  Required: true
	  In: path
	*/
	HPCResourceID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewHPCApprovedResourceRequestParams() beforehand.
func (o *HPCApprovedResourceRequestParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rHPCResourceID, rhkHPCResourceID, _ := route.Params.GetOK("HPCResourceID")
	if err := o.bindHPCResourceID(rHPCResourceID, rhkHPCResourceID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindHPCResourceID binds and validates parameter HPCResourceID from path.
func (o *HPCApprovedResourceRequestParams) bindHPCResourceID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.HPCResourceID = raw

	return nil
}