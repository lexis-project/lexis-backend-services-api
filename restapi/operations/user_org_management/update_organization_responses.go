// Code generated by go-swagger; DO NOT EDIT.

package user_org_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/lexis-project/lexis-backend-services-api.git/models"
)

// UpdateOrganizationOKCode is the HTTP code returned for type UpdateOrganizationOK
const UpdateOrganizationOKCode int = 200

/*UpdateOrganizationOK updated organization

swagger:response updateOrganizationOK
*/
type UpdateOrganizationOK struct {

	/*
	  In: Body
	*/
	Payload *models.Organization `json:"body,omitempty"`
}

// NewUpdateOrganizationOK creates UpdateOrganizationOK with default headers values
func NewUpdateOrganizationOK() *UpdateOrganizationOK {

	return &UpdateOrganizationOK{}
}

// WithPayload adds the payload to the update organization o k response
func (o *UpdateOrganizationOK) WithPayload(payload *models.Organization) *UpdateOrganizationOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update organization o k response
func (o *UpdateOrganizationOK) SetPayload(payload *models.Organization) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateOrganizationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateOrganizationUnauthorizedCode is the HTTP code returned for type UpdateOrganizationUnauthorized
const UpdateOrganizationUnauthorizedCode int = 401

/*UpdateOrganizationUnauthorized Unauthorized

swagger:response updateOrganizationUnauthorized
*/
type UpdateOrganizationUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewUpdateOrganizationUnauthorized creates UpdateOrganizationUnauthorized with default headers values
func NewUpdateOrganizationUnauthorized() *UpdateOrganizationUnauthorized {

	return &UpdateOrganizationUnauthorized{}
}

// WithPayload adds the payload to the update organization unauthorized response
func (o *UpdateOrganizationUnauthorized) WithPayload(payload *models.ErrorResponse) *UpdateOrganizationUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update organization unauthorized response
func (o *UpdateOrganizationUnauthorized) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateOrganizationUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateOrganizationForbiddenCode is the HTTP code returned for type UpdateOrganizationForbidden
const UpdateOrganizationForbiddenCode int = 403

/*UpdateOrganizationForbidden Forbidden

swagger:response updateOrganizationForbidden
*/
type UpdateOrganizationForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewUpdateOrganizationForbidden creates UpdateOrganizationForbidden with default headers values
func NewUpdateOrganizationForbidden() *UpdateOrganizationForbidden {

	return &UpdateOrganizationForbidden{}
}

// WithPayload adds the payload to the update organization forbidden response
func (o *UpdateOrganizationForbidden) WithPayload(payload *models.ErrorResponse) *UpdateOrganizationForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update organization forbidden response
func (o *UpdateOrganizationForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateOrganizationForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateOrganizationNotFoundCode is the HTTP code returned for type UpdateOrganizationNotFound
const UpdateOrganizationNotFoundCode int = 404

/*UpdateOrganizationNotFound organization with organizationId not found

swagger:response updateOrganizationNotFound
*/
type UpdateOrganizationNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewUpdateOrganizationNotFound creates UpdateOrganizationNotFound with default headers values
func NewUpdateOrganizationNotFound() *UpdateOrganizationNotFound {

	return &UpdateOrganizationNotFound{}
}

// WithPayload adds the payload to the update organization not found response
func (o *UpdateOrganizationNotFound) WithPayload(payload *models.ErrorResponse) *UpdateOrganizationNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update organization not found response
func (o *UpdateOrganizationNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateOrganizationNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateOrganizationInternalServerErrorCode is the HTTP code returned for type UpdateOrganizationInternalServerError
const UpdateOrganizationInternalServerErrorCode int = 500

/*UpdateOrganizationInternalServerError unexpected error

swagger:response updateOrganizationInternalServerError
*/
type UpdateOrganizationInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewUpdateOrganizationInternalServerError creates UpdateOrganizationInternalServerError with default headers values
func NewUpdateOrganizationInternalServerError() *UpdateOrganizationInternalServerError {

	return &UpdateOrganizationInternalServerError{}
}

// WithPayload adds the payload to the update organization internal server error response
func (o *UpdateOrganizationInternalServerError) WithPayload(payload *models.ErrorResponse) *UpdateOrganizationInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update organization internal server error response
func (o *UpdateOrganizationInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateOrganizationInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
