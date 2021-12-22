// Code generated by go-swagger; DO NOT EDIT.

package usage_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewGetUsageParams creates a new GetUsageParams object
// no default values defined in spec.
func NewGetUsageParams() GetUsageParams {

	return GetUsageParams{}
}

// GetUsageParams contains all the bound params for the get usage operation
// typically these are obtained from a http.Request
//
// swagger:parameters getUsage
type GetUsageParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Datetime from which to get the usage report
	  In: query
	*/
	From *strfmt.DateTime
	/*Id of project to be obtained
	  Required: true
	  In: path
	*/
	ID strfmt.UUID
	/*Datetime until which to get the usage report
	  In: query
	*/
	To *strfmt.DateTime
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetUsageParams() beforehand.
func (o *GetUsageParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qFrom, qhkFrom, _ := qs.GetOK("from")
	if err := o.bindFrom(qFrom, qhkFrom, route.Formats); err != nil {
		res = append(res, err)
	}

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	qTo, qhkTo, _ := qs.GetOK("to")
	if err := o.bindTo(qTo, qhkTo, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindFrom binds and validates parameter From from query.
func (o *GetUsageParams) bindFrom(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	// Format: datetime
	value, err := formats.Parse("datetime", raw)
	if err != nil {
		return errors.InvalidType("from", "query", "strfmt.DateTime", raw)
	}
	o.From = (value.(*strfmt.DateTime))

	if err := o.validateFrom(formats); err != nil {
		return err
	}

	return nil
}

// validateFrom carries on validations for parameter From
func (o *GetUsageParams) validateFrom(formats strfmt.Registry) error {

	if err := validate.FormatOf("from", "query", "datetime", o.From.String(), formats); err != nil {
		return err
	}
	return nil
}

// bindID binds and validates parameter ID from path.
func (o *GetUsageParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("id", "path", "strfmt.UUID", raw)
	}
	o.ID = *(value.(*strfmt.UUID))

	if err := o.validateID(formats); err != nil {
		return err
	}

	return nil
}

// validateID carries on validations for parameter ID
func (o *GetUsageParams) validateID(formats strfmt.Registry) error {

	if err := validate.FormatOf("id", "path", "uuid", o.ID.String(), formats); err != nil {
		return err
	}
	return nil
}

// bindTo binds and validates parameter To from query.
func (o *GetUsageParams) bindTo(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	// Format: datetime
	value, err := formats.Parse("datetime", raw)
	if err != nil {
		return errors.InvalidType("to", "query", "strfmt.DateTime", raw)
	}
	o.To = (value.(*strfmt.DateTime))

	if err := o.validateTo(formats); err != nil {
		return err
	}

	return nil
}

// validateTo carries on validations for parameter To
func (o *GetUsageParams) validateTo(formats strfmt.Registry) error {

	if err := validate.FormatOf("to", "query", "datetime", o.To.String(), formats); err != nil {
		return err
	}
	return nil
}
