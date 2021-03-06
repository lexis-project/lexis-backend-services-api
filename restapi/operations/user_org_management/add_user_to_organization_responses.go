// Code generated by go-swagger; DO NOT EDIT.

package user_org_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/lexis-project/lexis-backend-services-api.git/models"
)

// AddUserToOrganizationOKCode is the HTTP code returned for type AddUserToOrganizationOK
const AddUserToOrganizationOKCode int = 200

/*AddUserToOrganizationOK user updated

swagger:response addUserToOrganizationOK
*/
type AddUserToOrganizationOK struct {

	/*
	  In: Body
	*/
	Payload *models.OKResponse `json:"body,omitempty"`
}

// NewAddUserToOrganizationOK creates AddUserToOrganizationOK with default headers values
func NewAddUserToOrganizationOK() *AddUserToOrganizationOK {

	return &AddUserToOrganizationOK{}
}

// WithPayload adds the payload to the add user to organization o k response
func (o *AddUserToOrganizationOK) WithPayload(payload *models.OKResponse) *AddUserToOrganizationOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add user to organization o k response
func (o *AddUserToOrganizationOK) SetPayload(payload *models.OKResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddUserToOrganizationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddUserToOrganizationUnauthorizedCode is the HTTP code returned for type AddUserToOrganizationUnauthorized
const AddUserToOrganizationUnauthorizedCode int = 401

/*AddUserToOrganizationUnauthorized Authorization error

swagger:response addUserToOrganizationUnauthorized
*/
type AddUserToOrganizationUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewAddUserToOrganizationUnauthorized creates AddUserToOrganizationUnauthorized with default headers values
func NewAddUserToOrganizationUnauthorized() *AddUserToOrganizationUnauthorized {

	return &AddUserToOrganizationUnauthorized{}
}

// WithPayload adds the payload to the add user to organization unauthorized response
func (o *AddUserToOrganizationUnauthorized) WithPayload(payload *models.ErrorResponse) *AddUserToOrganizationUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add user to organization unauthorized response
func (o *AddUserToOrganizationUnauthorized) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddUserToOrganizationUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddUserToOrganizationForbiddenCode is the HTTP code returned for type AddUserToOrganizationForbidden
const AddUserToOrganizationForbiddenCode int = 403

/*AddUserToOrganizationForbidden Authorization error

swagger:response addUserToOrganizationForbidden
*/
type AddUserToOrganizationForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewAddUserToOrganizationForbidden creates AddUserToOrganizationForbidden with default headers values
func NewAddUserToOrganizationForbidden() *AddUserToOrganizationForbidden {

	return &AddUserToOrganizationForbidden{}
}

// WithPayload adds the payload to the add user to organization forbidden response
func (o *AddUserToOrganizationForbidden) WithPayload(payload *models.ErrorResponse) *AddUserToOrganizationForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add user to organization forbidden response
func (o *AddUserToOrganizationForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddUserToOrganizationForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddUserToOrganizationNotFoundCode is the HTTP code returned for type AddUserToOrganizationNotFound
const AddUserToOrganizationNotFoundCode int = 404

/*AddUserToOrganizationNotFound organization with organizationId not found

swagger:response addUserToOrganizationNotFound
*/
type AddUserToOrganizationNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewAddUserToOrganizationNotFound creates AddUserToOrganizationNotFound with default headers values
func NewAddUserToOrganizationNotFound() *AddUserToOrganizationNotFound {

	return &AddUserToOrganizationNotFound{}
}

// WithPayload adds the payload to the add user to organization not found response
func (o *AddUserToOrganizationNotFound) WithPayload(payload *models.ErrorResponse) *AddUserToOrganizationNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add user to organization not found response
func (o *AddUserToOrganizationNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddUserToOrganizationNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddUserToOrganizationInternalServerErrorCode is the HTTP code returned for type AddUserToOrganizationInternalServerError
const AddUserToOrganizationInternalServerErrorCode int = 500

/*AddUserToOrganizationInternalServerError unexpected error

swagger:response addUserToOrganizationInternalServerError
*/
type AddUserToOrganizationInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewAddUserToOrganizationInternalServerError creates AddUserToOrganizationInternalServerError with default headers values
func NewAddUserToOrganizationInternalServerError() *AddUserToOrganizationInternalServerError {

	return &AddUserToOrganizationInternalServerError{}
}

// WithPayload adds the payload to the add user to organization internal server error response
func (o *AddUserToOrganizationInternalServerError) WithPayload(payload *models.ErrorResponse) *AddUserToOrganizationInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add user to organization internal server error response
func (o *AddUserToOrganizationInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddUserToOrganizationInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
