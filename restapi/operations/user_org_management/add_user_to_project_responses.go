// Code generated by go-swagger; DO NOT EDIT.

package user_org_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/lexis-project/lexis-backend-services-api.git/models"
)

// AddUserToProjectOKCode is the HTTP code returned for type AddUserToProjectOK
const AddUserToProjectOKCode int = 200

/*AddUserToProjectOK user updated

swagger:response addUserToProjectOK
*/
type AddUserToProjectOK struct {

	/*
	  In: Body
	*/
	Payload *models.OKResponse `json:"body,omitempty"`
}

// NewAddUserToProjectOK creates AddUserToProjectOK with default headers values
func NewAddUserToProjectOK() *AddUserToProjectOK {

	return &AddUserToProjectOK{}
}

// WithPayload adds the payload to the add user to project o k response
func (o *AddUserToProjectOK) WithPayload(payload *models.OKResponse) *AddUserToProjectOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add user to project o k response
func (o *AddUserToProjectOK) SetPayload(payload *models.OKResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddUserToProjectOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddUserToProjectUnauthorizedCode is the HTTP code returned for type AddUserToProjectUnauthorized
const AddUserToProjectUnauthorizedCode int = 401

/*AddUserToProjectUnauthorized Authorization error

swagger:response addUserToProjectUnauthorized
*/
type AddUserToProjectUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewAddUserToProjectUnauthorized creates AddUserToProjectUnauthorized with default headers values
func NewAddUserToProjectUnauthorized() *AddUserToProjectUnauthorized {

	return &AddUserToProjectUnauthorized{}
}

// WithPayload adds the payload to the add user to project unauthorized response
func (o *AddUserToProjectUnauthorized) WithPayload(payload *models.ErrorResponse) *AddUserToProjectUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add user to project unauthorized response
func (o *AddUserToProjectUnauthorized) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddUserToProjectUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddUserToProjectForbiddenCode is the HTTP code returned for type AddUserToProjectForbidden
const AddUserToProjectForbiddenCode int = 403

/*AddUserToProjectForbidden Authorization error

swagger:response addUserToProjectForbidden
*/
type AddUserToProjectForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewAddUserToProjectForbidden creates AddUserToProjectForbidden with default headers values
func NewAddUserToProjectForbidden() *AddUserToProjectForbidden {

	return &AddUserToProjectForbidden{}
}

// WithPayload adds the payload to the add user to project forbidden response
func (o *AddUserToProjectForbidden) WithPayload(payload *models.ErrorResponse) *AddUserToProjectForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add user to project forbidden response
func (o *AddUserToProjectForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddUserToProjectForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddUserToProjectNotFoundCode is the HTTP code returned for type AddUserToProjectNotFound
const AddUserToProjectNotFoundCode int = 404

/*AddUserToProjectNotFound project with not found

swagger:response addUserToProjectNotFound
*/
type AddUserToProjectNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewAddUserToProjectNotFound creates AddUserToProjectNotFound with default headers values
func NewAddUserToProjectNotFound() *AddUserToProjectNotFound {

	return &AddUserToProjectNotFound{}
}

// WithPayload adds the payload to the add user to project not found response
func (o *AddUserToProjectNotFound) WithPayload(payload *models.ErrorResponse) *AddUserToProjectNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add user to project not found response
func (o *AddUserToProjectNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddUserToProjectNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddUserToProjectInternalServerErrorCode is the HTTP code returned for type AddUserToProjectInternalServerError
const AddUserToProjectInternalServerErrorCode int = 500

/*AddUserToProjectInternalServerError unexpected error

swagger:response addUserToProjectInternalServerError
*/
type AddUserToProjectInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewAddUserToProjectInternalServerError creates AddUserToProjectInternalServerError with default headers values
func NewAddUserToProjectInternalServerError() *AddUserToProjectInternalServerError {

	return &AddUserToProjectInternalServerError{}
}

// WithPayload adds the payload to the add user to project internal server error response
func (o *AddUserToProjectInternalServerError) WithPayload(payload *models.ErrorResponse) *AddUserToProjectInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add user to project internal server error response
func (o *AddUserToProjectInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddUserToProjectInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}