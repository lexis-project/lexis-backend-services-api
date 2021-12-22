// Code generated by go-swagger; DO NOT EDIT.

package user_org_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/lexis-project/lexis-backend-services-api.git/models"
)

// GetHPCResourceOKCode is the HTTP code returned for type GetHPCResourceOK
const GetHPCResourceOKCode int = 200

/*GetHPCResourceOK HPCResource returned

swagger:response getHPCResourceOK
*/
type GetHPCResourceOK struct {

	/*
	  In: Body
	*/
	Payload *models.HPCResource `json:"body,omitempty"`
}

// NewGetHPCResourceOK creates GetHPCResourceOK with default headers values
func NewGetHPCResourceOK() *GetHPCResourceOK {

	return &GetHPCResourceOK{}
}

// WithPayload adds the payload to the get h p c resource o k response
func (o *GetHPCResourceOK) WithPayload(payload *models.HPCResource) *GetHPCResourceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get h p c resource o k response
func (o *GetHPCResourceOK) SetPayload(payload *models.HPCResource) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHPCResourceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetHPCResourceUnauthorizedCode is the HTTP code returned for type GetHPCResourceUnauthorized
const GetHPCResourceUnauthorizedCode int = 401

/*GetHPCResourceUnauthorized Unauthorized

swagger:response getHPCResourceUnauthorized
*/
type GetHPCResourceUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetHPCResourceUnauthorized creates GetHPCResourceUnauthorized with default headers values
func NewGetHPCResourceUnauthorized() *GetHPCResourceUnauthorized {

	return &GetHPCResourceUnauthorized{}
}

// WithPayload adds the payload to the get h p c resource unauthorized response
func (o *GetHPCResourceUnauthorized) WithPayload(payload *models.ErrorResponse) *GetHPCResourceUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get h p c resource unauthorized response
func (o *GetHPCResourceUnauthorized) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHPCResourceUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetHPCResourceForbiddenCode is the HTTP code returned for type GetHPCResourceForbidden
const GetHPCResourceForbiddenCode int = 403

/*GetHPCResourceForbidden Forbidden

swagger:response getHPCResourceForbidden
*/
type GetHPCResourceForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetHPCResourceForbidden creates GetHPCResourceForbidden with default headers values
func NewGetHPCResourceForbidden() *GetHPCResourceForbidden {

	return &GetHPCResourceForbidden{}
}

// WithPayload adds the payload to the get h p c resource forbidden response
func (o *GetHPCResourceForbidden) WithPayload(payload *models.ErrorResponse) *GetHPCResourceForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get h p c resource forbidden response
func (o *GetHPCResourceForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHPCResourceForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetHPCResourceNotFoundCode is the HTTP code returned for type GetHPCResourceNotFound
const GetHPCResourceNotFoundCode int = 404

/*GetHPCResourceNotFound HPCResource with not found

swagger:response getHPCResourceNotFound
*/
type GetHPCResourceNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetHPCResourceNotFound creates GetHPCResourceNotFound with default headers values
func NewGetHPCResourceNotFound() *GetHPCResourceNotFound {

	return &GetHPCResourceNotFound{}
}

// WithPayload adds the payload to the get h p c resource not found response
func (o *GetHPCResourceNotFound) WithPayload(payload *models.ErrorResponse) *GetHPCResourceNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get h p c resource not found response
func (o *GetHPCResourceNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHPCResourceNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetHPCResourceInternalServerErrorCode is the HTTP code returned for type GetHPCResourceInternalServerError
const GetHPCResourceInternalServerErrorCode int = 500

/*GetHPCResourceInternalServerError unexpected error

swagger:response getHPCResourceInternalServerError
*/
type GetHPCResourceInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetHPCResourceInternalServerError creates GetHPCResourceInternalServerError with default headers values
func NewGetHPCResourceInternalServerError() *GetHPCResourceInternalServerError {

	return &GetHPCResourceInternalServerError{}
}

// WithPayload adds the payload to the get h p c resource internal server error response
func (o *GetHPCResourceInternalServerError) WithPayload(payload *models.ErrorResponse) *GetHPCResourceInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get h p c resource internal server error response
func (o *GetHPCResourceInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHPCResourceInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}