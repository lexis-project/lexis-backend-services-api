// Code generated by go-swagger; DO NOT EDIT.

package data_set_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"encoding/json"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/lexis-project/lexis-backend-services-api.git/models"
)

// CreateDatasetHandlerFunc turns a function with the right signature into a create dataset handler
type CreateDatasetHandlerFunc func(CreateDatasetParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateDatasetHandlerFunc) Handle(params CreateDatasetParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CreateDatasetHandler interface for that can handle valid create dataset params
type CreateDatasetHandler interface {
	Handle(CreateDatasetParams, interface{}) middleware.Responder
}

// NewCreateDataset creates a new http.Handler for the create dataset operation
func NewCreateDataset(ctx *middleware.Context, handler CreateDatasetHandler) *CreateDataset {
	return &CreateDataset{Context: ctx, Handler: handler}
}

/*CreateDataset swagger:route POST /dataset dataSetManagement createDataset

Create a dataset

Creates a new dataset

*/
type CreateDataset struct {
	Context *middleware.Context
	Handler CreateDatasetHandler
}

func (o *CreateDataset) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateDatasetParams()

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

// CreateDatasetBody create dataset body
//
// swagger:model CreateDatasetBody
type CreateDatasetBody struct {

	// access
	// Required: true
	Access models.AccessMode `json:"access"`

	// comp
	// Enum: [yes no]
	Comp string `json:"comp,omitempty"`

	// compress method
	CompressMethod models.CompressMethod `json:"compress_method,omitempty"`

	// enc
	// Enum: [yes no]
	Enc string `json:"enc,omitempty"`

	// - If pushmethod is directupload, json-escaped, base64-encoded
	// file or zip (depending on compress_method).
	// - If pushmethod is tus, url from the Location returned by the
	// previous tus call.
	//
	File string `json:"file,omitempty"`

	// internal ID
	InternalID string `json:"internalID,omitempty"`

	// metadata
	Metadata *models.DatasetMetadata `json:"metadata,omitempty"`

	// name of the file if compress_method is file, or if using pushmethod tus.
	Name string `json:"name,omitempty"`

	// path within the dataset, without a starting slash. An empty
	// string pushes to the root of the dataset (i.e. inside the
	// directory named <internalID> in the iRODS backend).
	//
	Path string `json:"path,omitempty"`

	// project
	// Required: true
	Project *string `json:"project"`

	// push method
	// Required: true
	PushMethod models.PushMethod `json:"push_method"`

	// zone
	Zone string `json:"zone,omitempty"`
}

// Validate validates this create dataset body
func (o *CreateDatasetBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateComp(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCompressMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateEnc(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePushMethod(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateDatasetBody) validateAccess(formats strfmt.Registry) error {

	if err := o.Access.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("dataset" + "." + "access")
		}
		return err
	}

	return nil
}

var createDatasetBodyTypeCompPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["yes","no"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createDatasetBodyTypeCompPropEnum = append(createDatasetBodyTypeCompPropEnum, v)
	}
}

const (

	// CreateDatasetBodyCompYes captures enum value "yes"
	CreateDatasetBodyCompYes string = "yes"

	// CreateDatasetBodyCompNo captures enum value "no"
	CreateDatasetBodyCompNo string = "no"
)

// prop value enum
func (o *CreateDatasetBody) validateCompEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createDatasetBodyTypeCompPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *CreateDatasetBody) validateComp(formats strfmt.Registry) error {

	if swag.IsZero(o.Comp) { // not required
		return nil
	}

	// value enum
	if err := o.validateCompEnum("dataset"+"."+"comp", "body", o.Comp); err != nil {
		return err
	}

	return nil
}

func (o *CreateDatasetBody) validateCompressMethod(formats strfmt.Registry) error {

	if swag.IsZero(o.CompressMethod) { // not required
		return nil
	}

	if err := o.CompressMethod.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("dataset" + "." + "compress_method")
		}
		return err
	}

	return nil
}

var createDatasetBodyTypeEncPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["yes","no"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createDatasetBodyTypeEncPropEnum = append(createDatasetBodyTypeEncPropEnum, v)
	}
}

const (

	// CreateDatasetBodyEncYes captures enum value "yes"
	CreateDatasetBodyEncYes string = "yes"

	// CreateDatasetBodyEncNo captures enum value "no"
	CreateDatasetBodyEncNo string = "no"
)

// prop value enum
func (o *CreateDatasetBody) validateEncEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createDatasetBodyTypeEncPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *CreateDatasetBody) validateEnc(formats strfmt.Registry) error {

	if swag.IsZero(o.Enc) { // not required
		return nil
	}

	// value enum
	if err := o.validateEncEnum("dataset"+"."+"enc", "body", o.Enc); err != nil {
		return err
	}

	return nil
}

func (o *CreateDatasetBody) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(o.Metadata) { // not required
		return nil
	}

	if o.Metadata != nil {
		if err := o.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dataset" + "." + "metadata")
			}
			return err
		}
	}

	return nil
}

func (o *CreateDatasetBody) validateProject(formats strfmt.Registry) error {

	if err := validate.Required("dataset"+"."+"project", "body", o.Project); err != nil {
		return err
	}

	return nil
}

func (o *CreateDatasetBody) validatePushMethod(formats strfmt.Registry) error {

	if err := o.PushMethod.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("dataset" + "." + "push_method")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateDatasetBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateDatasetBody) UnmarshalBinary(b []byte) error {
	var res CreateDatasetBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}