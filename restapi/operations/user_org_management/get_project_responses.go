// Code generated by go-swagger; DO NOT EDIT.

package user_org_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/lexis-project/lexis-backend-services-api.git/models"
)

// GetProjectOKCode is the HTTP code returned for type GetProjectOK
const GetProjectOKCode int = 200

/*GetProjectOK project returned

swagger:response getProjectOK
*/
type GetProjectOK struct {

	/*
	  In: Body
	*/
	Payload *models.Project `json:"body,omitempty"`
}

// NewGetProjectOK creates GetProjectOK with default headers values
func NewGetProjectOK() *GetProjectOK {

	return &GetProjectOK{}
}

// WithPayload adds the payload to the get project o k response
func (o *GetProjectOK) WithPayload(payload *models.Project) *GetProjectOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get project o k response
func (o *GetProjectOK) SetPayload(payload *models.Project) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProjectOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetProjectUnauthorizedCode is the HTTP code returned for type GetProjectUnauthorized
const GetProjectUnauthorizedCode int = 401

/*GetProjectUnauthorized Unauthorized

swagger:response getProjectUnauthorized
*/
type GetProjectUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetProjectUnauthorized creates GetProjectUnauthorized with default headers values
func NewGetProjectUnauthorized() *GetProjectUnauthorized {

	return &GetProjectUnauthorized{}
}

// WithPayload adds the payload to the get project unauthorized response
func (o *GetProjectUnauthorized) WithPayload(payload *models.ErrorResponse) *GetProjectUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get project unauthorized response
func (o *GetProjectUnauthorized) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProjectUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetProjectForbiddenCode is the HTTP code returned for type GetProjectForbidden
const GetProjectForbiddenCode int = 403

/*GetProjectForbidden Forbidden

swagger:response getProjectForbidden
*/
type GetProjectForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetProjectForbidden creates GetProjectForbidden with default headers values
func NewGetProjectForbidden() *GetProjectForbidden {

	return &GetProjectForbidden{}
}

// WithPayload adds the payload to the get project forbidden response
func (o *GetProjectForbidden) WithPayload(payload *models.ErrorResponse) *GetProjectForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get project forbidden response
func (o *GetProjectForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProjectForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetProjectNotFoundCode is the HTTP code returned for type GetProjectNotFound
const GetProjectNotFoundCode int = 404

/*GetProjectNotFound project with not found

swagger:response getProjectNotFound
*/
type GetProjectNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetProjectNotFound creates GetProjectNotFound with default headers values
func NewGetProjectNotFound() *GetProjectNotFound {

	return &GetProjectNotFound{}
}

// WithPayload adds the payload to the get project not found response
func (o *GetProjectNotFound) WithPayload(payload *models.ErrorResponse) *GetProjectNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get project not found response
func (o *GetProjectNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProjectNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetProjectInternalServerErrorCode is the HTTP code returned for type GetProjectInternalServerError
const GetProjectInternalServerErrorCode int = 500

/*GetProjectInternalServerError unexpected error

swagger:response getProjectInternalServerError
*/
type GetProjectInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetProjectInternalServerError creates GetProjectInternalServerError with default headers values
func NewGetProjectInternalServerError() *GetProjectInternalServerError {

	return &GetProjectInternalServerError{}
}

// WithPayload adds the payload to the get project internal server error response
func (o *GetProjectInternalServerError) WithPayload(payload *models.ErrorResponse) *GetProjectInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get project internal server error response
func (o *GetProjectInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProjectInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
