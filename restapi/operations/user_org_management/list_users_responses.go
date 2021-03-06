// Code generated by go-swagger; DO NOT EDIT.

package user_org_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/lexis-project/lexis-backend-services-api.git/models"
)

// ListUsersOKCode is the HTTP code returned for type ListUsersOK
const ListUsersOKCode int = 200

/*ListUsersOK list of users returned

swagger:response listUsersOK
*/
type ListUsersOK struct {

	/*
	  In: Body
	*/
	Payload []*models.User `json:"body,omitempty"`
}

// NewListUsersOK creates ListUsersOK with default headers values
func NewListUsersOK() *ListUsersOK {

	return &ListUsersOK{}
}

// WithPayload adds the payload to the list users o k response
func (o *ListUsersOK) WithPayload(payload []*models.User) *ListUsersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list users o k response
func (o *ListUsersOK) SetPayload(payload []*models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListUsersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.User, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListUsersUnauthorizedCode is the HTTP code returned for type ListUsersUnauthorized
const ListUsersUnauthorizedCode int = 401

/*ListUsersUnauthorized Unauthorized

swagger:response listUsersUnauthorized
*/
type ListUsersUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewListUsersUnauthorized creates ListUsersUnauthorized with default headers values
func NewListUsersUnauthorized() *ListUsersUnauthorized {

	return &ListUsersUnauthorized{}
}

// WithPayload adds the payload to the list users unauthorized response
func (o *ListUsersUnauthorized) WithPayload(payload *models.ErrorResponse) *ListUsersUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list users unauthorized response
func (o *ListUsersUnauthorized) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListUsersUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListUsersForbiddenCode is the HTTP code returned for type ListUsersForbidden
const ListUsersForbiddenCode int = 403

/*ListUsersForbidden Forbidden

swagger:response listUsersForbidden
*/
type ListUsersForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewListUsersForbidden creates ListUsersForbidden with default headers values
func NewListUsersForbidden() *ListUsersForbidden {

	return &ListUsersForbidden{}
}

// WithPayload adds the payload to the list users forbidden response
func (o *ListUsersForbidden) WithPayload(payload *models.ErrorResponse) *ListUsersForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list users forbidden response
func (o *ListUsersForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListUsersForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListUsersInternalServerErrorCode is the HTTP code returned for type ListUsersInternalServerError
const ListUsersInternalServerErrorCode int = 500

/*ListUsersInternalServerError unexpected error

swagger:response listUsersInternalServerError
*/
type ListUsersInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewListUsersInternalServerError creates ListUsersInternalServerError with default headers values
func NewListUsersInternalServerError() *ListUsersInternalServerError {

	return &ListUsersInternalServerError{}
}

// WithPayload adds the payload to the list users internal server error response
func (o *ListUsersInternalServerError) WithPayload(payload *models.ErrorResponse) *ListUsersInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list users internal server error response
func (o *ListUsersInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListUsersInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
