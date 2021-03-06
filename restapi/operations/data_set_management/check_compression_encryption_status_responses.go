// Code generated by go-swagger; DO NOT EDIT.

package data_set_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/lexis-project/lexis-backend-services-api.git/models"
)

// CheckCompressionEncryptionStatusOKCode is the HTTP code returned for type CheckCompressionEncryptionStatusOK
const CheckCompressionEncryptionStatusOKCode int = 200

/*CheckCompressionEncryptionStatusOK This means that the status has been returned to the user in the response body.

swagger:response checkCompressionEncryptionStatusOK
*/
type CheckCompressionEncryptionStatusOK struct {

	/*
	  In: Body
	*/
	Payload *CheckCompressionEncryptionStatusOKBody `json:"body,omitempty"`
}

// NewCheckCompressionEncryptionStatusOK creates CheckCompressionEncryptionStatusOK with default headers values
func NewCheckCompressionEncryptionStatusOK() *CheckCompressionEncryptionStatusOK {

	return &CheckCompressionEncryptionStatusOK{}
}

// WithPayload adds the payload to the check compression encryption status o k response
func (o *CheckCompressionEncryptionStatusOK) WithPayload(payload *CheckCompressionEncryptionStatusOKBody) *CheckCompressionEncryptionStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check compression encryption status o k response
func (o *CheckCompressionEncryptionStatusOK) SetPayload(payload *CheckCompressionEncryptionStatusOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckCompressionEncryptionStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckCompressionEncryptionStatusBadRequestCode is the HTTP code returned for type CheckCompressionEncryptionStatusBadRequest
const CheckCompressionEncryptionStatusBadRequestCode int = 400

/*CheckCompressionEncryptionStatusBadRequest This means that the request ID given by the user is incorrect.

swagger:response checkCompressionEncryptionStatusBadRequest
*/
type CheckCompressionEncryptionStatusBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCheckCompressionEncryptionStatusBadRequest creates CheckCompressionEncryptionStatusBadRequest with default headers values
func NewCheckCompressionEncryptionStatusBadRequest() *CheckCompressionEncryptionStatusBadRequest {

	return &CheckCompressionEncryptionStatusBadRequest{}
}

// WithPayload adds the payload to the check compression encryption status bad request response
func (o *CheckCompressionEncryptionStatusBadRequest) WithPayload(payload *models.ErrorResponse) *CheckCompressionEncryptionStatusBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check compression encryption status bad request response
func (o *CheckCompressionEncryptionStatusBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckCompressionEncryptionStatusBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckCompressionEncryptionStatusUnauthorizedCode is the HTTP code returned for type CheckCompressionEncryptionStatusUnauthorized
const CheckCompressionEncryptionStatusUnauthorizedCode int = 401

/*CheckCompressionEncryptionStatusUnauthorized This means that the user is not authenticated with keycloak and compression with encryption can't be triggered unless the user first log in with a valid user

swagger:response checkCompressionEncryptionStatusUnauthorized
*/
type CheckCompressionEncryptionStatusUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCheckCompressionEncryptionStatusUnauthorized creates CheckCompressionEncryptionStatusUnauthorized with default headers values
func NewCheckCompressionEncryptionStatusUnauthorized() *CheckCompressionEncryptionStatusUnauthorized {

	return &CheckCompressionEncryptionStatusUnauthorized{}
}

// WithPayload adds the payload to the check compression encryption status unauthorized response
func (o *CheckCompressionEncryptionStatusUnauthorized) WithPayload(payload *models.ErrorResponse) *CheckCompressionEncryptionStatusUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check compression encryption status unauthorized response
func (o *CheckCompressionEncryptionStatusUnauthorized) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckCompressionEncryptionStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckCompressionEncryptionStatusNotFoundCode is the HTTP code returned for type CheckCompressionEncryptionStatusNotFound
const CheckCompressionEncryptionStatusNotFoundCode int = 404

/*CheckCompressionEncryptionStatusNotFound This means that the ID doesn't exist and thus a status can't be returned.

swagger:response checkCompressionEncryptionStatusNotFound
*/
type CheckCompressionEncryptionStatusNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCheckCompressionEncryptionStatusNotFound creates CheckCompressionEncryptionStatusNotFound with default headers values
func NewCheckCompressionEncryptionStatusNotFound() *CheckCompressionEncryptionStatusNotFound {

	return &CheckCompressionEncryptionStatusNotFound{}
}

// WithPayload adds the payload to the check compression encryption status not found response
func (o *CheckCompressionEncryptionStatusNotFound) WithPayload(payload *models.ErrorResponse) *CheckCompressionEncryptionStatusNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check compression encryption status not found response
func (o *CheckCompressionEncryptionStatusNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckCompressionEncryptionStatusNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckCompressionEncryptionStatusRequestURITooLongCode is the HTTP code returned for type CheckCompressionEncryptionStatusRequestURITooLong
const CheckCompressionEncryptionStatusRequestURITooLongCode int = 414

/*CheckCompressionEncryptionStatusRequestURITooLong This means that the the request ID is longer than the server is willing to interpret.

swagger:response checkCompressionEncryptionStatusRequestUriTooLong
*/
type CheckCompressionEncryptionStatusRequestURITooLong struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCheckCompressionEncryptionStatusRequestURITooLong creates CheckCompressionEncryptionStatusRequestURITooLong with default headers values
func NewCheckCompressionEncryptionStatusRequestURITooLong() *CheckCompressionEncryptionStatusRequestURITooLong {

	return &CheckCompressionEncryptionStatusRequestURITooLong{}
}

// WithPayload adds the payload to the check compression encryption status request Uri too long response
func (o *CheckCompressionEncryptionStatusRequestURITooLong) WithPayload(payload *models.ErrorResponse) *CheckCompressionEncryptionStatusRequestURITooLong {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check compression encryption status request Uri too long response
func (o *CheckCompressionEncryptionStatusRequestURITooLong) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckCompressionEncryptionStatusRequestURITooLong) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(414)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckCompressionEncryptionStatusInternalServerErrorCode is the HTTP code returned for type CheckCompressionEncryptionStatusInternalServerError
const CheckCompressionEncryptionStatusInternalServerErrorCode int = 500

/*CheckCompressionEncryptionStatusInternalServerError This means that something has gone wrong on the burst buffer. The user is advised to wait and send the request again.

swagger:response checkCompressionEncryptionStatusInternalServerError
*/
type CheckCompressionEncryptionStatusInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCheckCompressionEncryptionStatusInternalServerError creates CheckCompressionEncryptionStatusInternalServerError with default headers values
func NewCheckCompressionEncryptionStatusInternalServerError() *CheckCompressionEncryptionStatusInternalServerError {

	return &CheckCompressionEncryptionStatusInternalServerError{}
}

// WithPayload adds the payload to the check compression encryption status internal server error response
func (o *CheckCompressionEncryptionStatusInternalServerError) WithPayload(payload *models.ErrorResponse) *CheckCompressionEncryptionStatusInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check compression encryption status internal server error response
func (o *CheckCompressionEncryptionStatusInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckCompressionEncryptionStatusInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
