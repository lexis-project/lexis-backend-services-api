// Code generated by go-swagger; DO NOT EDIT.

package user_org_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/lexis-project/lexis-backend-services-api.git/models"
)

// DeleteOrganizationOKCode is the HTTP code returned for type DeleteOrganizationOK
const DeleteOrganizationOKCode int = 200

/*DeleteOrganizationOK deleted organization

swagger:response deleteOrganizationOK
*/
type DeleteOrganizationOK struct {

	/*
	  In: Body
	*/
	Payload *models.DeletedResponse `json:"body,omitempty"`
}

// NewDeleteOrganizationOK creates DeleteOrganizationOK with default headers values
func NewDeleteOrganizationOK() *DeleteOrganizationOK {

	return &DeleteOrganizationOK{}
}

// WithPayload adds the payload to the delete organization o k response
func (o *DeleteOrganizationOK) WithPayload(payload *models.DeletedResponse) *DeleteOrganizationOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete organization o k response
func (o *DeleteOrganizationOK) SetPayload(payload *models.DeletedResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteOrganizationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteOrganizationUnauthorizedCode is the HTTP code returned for type DeleteOrganizationUnauthorized
const DeleteOrganizationUnauthorizedCode int = 401

/*DeleteOrganizationUnauthorized Unauthorized

swagger:response deleteOrganizationUnauthorized
*/
type DeleteOrganizationUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewDeleteOrganizationUnauthorized creates DeleteOrganizationUnauthorized with default headers values
func NewDeleteOrganizationUnauthorized() *DeleteOrganizationUnauthorized {

	return &DeleteOrganizationUnauthorized{}
}

// WithPayload adds the payload to the delete organization unauthorized response
func (o *DeleteOrganizationUnauthorized) WithPayload(payload *models.ErrorResponse) *DeleteOrganizationUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete organization unauthorized response
func (o *DeleteOrganizationUnauthorized) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteOrganizationUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteOrganizationForbiddenCode is the HTTP code returned for type DeleteOrganizationForbidden
const DeleteOrganizationForbiddenCode int = 403

/*DeleteOrganizationForbidden Forbidden

swagger:response deleteOrganizationForbidden
*/
type DeleteOrganizationForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewDeleteOrganizationForbidden creates DeleteOrganizationForbidden with default headers values
func NewDeleteOrganizationForbidden() *DeleteOrganizationForbidden {

	return &DeleteOrganizationForbidden{}
}

// WithPayload adds the payload to the delete organization forbidden response
func (o *DeleteOrganizationForbidden) WithPayload(payload *models.ErrorResponse) *DeleteOrganizationForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete organization forbidden response
func (o *DeleteOrganizationForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteOrganizationForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteOrganizationNotFoundCode is the HTTP code returned for type DeleteOrganizationNotFound
const DeleteOrganizationNotFoundCode int = 404

/*DeleteOrganizationNotFound organization with organizationId not found

swagger:response deleteOrganizationNotFound
*/
type DeleteOrganizationNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewDeleteOrganizationNotFound creates DeleteOrganizationNotFound with default headers values
func NewDeleteOrganizationNotFound() *DeleteOrganizationNotFound {

	return &DeleteOrganizationNotFound{}
}

// WithPayload adds the payload to the delete organization not found response
func (o *DeleteOrganizationNotFound) WithPayload(payload *models.ErrorResponse) *DeleteOrganizationNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete organization not found response
func (o *DeleteOrganizationNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteOrganizationNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteOrganizationInternalServerErrorCode is the HTTP code returned for type DeleteOrganizationInternalServerError
const DeleteOrganizationInternalServerErrorCode int = 500

/*DeleteOrganizationInternalServerError unexpected error

swagger:response deleteOrganizationInternalServerError
*/
type DeleteOrganizationInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewDeleteOrganizationInternalServerError creates DeleteOrganizationInternalServerError with default headers values
func NewDeleteOrganizationInternalServerError() *DeleteOrganizationInternalServerError {

	return &DeleteOrganizationInternalServerError{}
}

// WithPayload adds the payload to the delete organization internal server error response
func (o *DeleteOrganizationInternalServerError) WithPayload(payload *models.ErrorResponse) *DeleteOrganizationInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete organization internal server error response
func (o *DeleteOrganizationInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteOrganizationInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}