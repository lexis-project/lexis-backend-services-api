// Code generated by go-swagger; DO NOT EDIT.

package data_set_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewOptionsDatasetUploadParams creates a new OptionsDatasetUploadParams object
// no default values defined in spec.
func NewOptionsDatasetUploadParams() OptionsDatasetUploadParams {

	return OptionsDatasetUploadParams{}
}

// OptionsDatasetUploadParams contains all the bound params for the options dataset upload operation
// typically these are obtained from a http.Request
//
// swagger:parameters OptionsDatasetUpload
type OptionsDatasetUploadParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Protocol version
	  Required: true
	  In: header
	*/
	TusResumable string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewOptionsDatasetUploadParams() beforehand.
func (o *OptionsDatasetUploadParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := o.bindTusResumable(r.Header[http.CanonicalHeaderKey("Tus-Resumable")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindTusResumable binds and validates parameter TusResumable from header.
func (o *OptionsDatasetUploadParams) bindTusResumable(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("Tus-Resumable", "header", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("Tus-Resumable", "header", raw); err != nil {
		return err
	}

	o.TusResumable = raw

	if err := o.validateTusResumable(formats); err != nil {
		return err
	}

	return nil
}

// validateTusResumable carries on validations for parameter TusResumable
func (o *OptionsDatasetUploadParams) validateTusResumable(formats strfmt.Registry) error {

	if err := validate.EnumCase("Tus-Resumable", "header", o.TusResumable, []interface{}{"1.0.0"}, true); err != nil {
		return err
	}

	return nil
}