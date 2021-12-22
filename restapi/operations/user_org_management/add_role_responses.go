// Code generated by go-swagger; DO NOT EDIT.

package user_org_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/lexis-project/lexis-backend-services-api.git/models"
)

// AddRoleOKCode is the HTTP code returned for type AddRoleOK
const AddRoleOKCode int = 200

/*AddRoleOK Role added successfully

swagger:response addRoleOK
*/
type AddRoleOK struct {

	/*
	  In: Body
	*/
	Payload *models.AuthorizationResponse `json:"body,omitempty"`
}

// NewAddRoleOK creates AddRoleOK with default headers values
func NewAddRoleOK() *AddRoleOK {

	return &AddRoleOK{}
}

// WithPayload adds the payload to the add role o k response
func (o *AddRoleOK) WithPayload(payload *models.AuthorizationResponse) *AddRoleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add role o k response
func (o *AddRoleOK) SetPayload(payload *models.AuthorizationResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddRoleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddRoleUnauthorizedCode is the HTTP code returned for type AddRoleUnauthorized
const AddRoleUnauthorizedCode int = 401

/*AddRoleUnauthorized Unauthorized

swagger:response addRoleUnauthorized
*/
type AddRoleUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewAddRoleUnauthorized creates AddRoleUnauthorized with default headers values
func NewAddRoleUnauthorized() *AddRoleUnauthorized {

	return &AddRoleUnauthorized{}
}

// WithPayload adds the payload to the add role unauthorized response
func (o *AddRoleUnauthorized) WithPayload(payload *models.ErrorResponse) *AddRoleUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add role unauthorized response
func (o *AddRoleUnauthorized) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddRoleUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddRoleForbiddenCode is the HTTP code returned for type AddRoleForbidden
const AddRoleForbiddenCode int = 403

/*AddRoleForbidden Forbidden

swagger:response addRoleForbidden
*/
type AddRoleForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewAddRoleForbidden creates AddRoleForbidden with default headers values
func NewAddRoleForbidden() *AddRoleForbidden {

	return &AddRoleForbidden{}
}

// WithPayload adds the payload to the add role forbidden response
func (o *AddRoleForbidden) WithPayload(payload *models.ErrorResponse) *AddRoleForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add role forbidden response
func (o *AddRoleForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddRoleForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddRoleInternalServerErrorCode is the HTTP code returned for type AddRoleInternalServerError
const AddRoleInternalServerErrorCode int = 500

/*AddRoleInternalServerError unexpected error

swagger:response addRoleInternalServerError
*/
type AddRoleInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewAddRoleInternalServerError creates AddRoleInternalServerError with default headers values
func NewAddRoleInternalServerError() *AddRoleInternalServerError {

	return &AddRoleInternalServerError{}
}

// WithPayload adds the payload to the add role internal server error response
func (o *AddRoleInternalServerError) WithPayload(payload *models.ErrorResponse) *AddRoleInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add role internal server error response
func (o *AddRoleInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddRoleInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}