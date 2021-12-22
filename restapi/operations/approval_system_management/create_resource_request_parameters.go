// Code generated by go-swagger; DO NOT EDIT.

package approval_system_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/lexis-project/lexis-backend-services-api.git/models"
)

// NewCreateResourceRequestParams creates a new CreateResourceRequestParams object
// no default values defined in spec.
func NewCreateResourceRequestParams() CreateResourceRequestParams {

	return CreateResourceRequestParams{}
}

// CreateResourceRequestParams contains all the bound params for the create resource request operation
// typically these are obtained from a http.Request
//
// swagger:parameters CreateResourceRequest
type CreateResourceRequestParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*HPC resource request to be created
	  In: body
	*/
	ResourceRequest *models.ApprovalSystemResourceRequest
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCreateResourceRequestParams() beforehand.
func (o *CreateResourceRequestParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.ApprovalSystemResourceRequest
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("resourceRequest", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.ResourceRequest = &body
			}
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
