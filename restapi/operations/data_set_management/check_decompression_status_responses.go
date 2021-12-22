// Code generated by go-swagger; DO NOT EDIT.

package data_set_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/lexis-project/lexis-backend-services-api.git/models"
)

// CheckDecompressionStatusOKCode is the HTTP code returned for type CheckDecompressionStatusOK
const CheckDecompressionStatusOKCode int = 200

/*CheckDecompressionStatusOK This means that the status has been returned to the user in the response body.

swagger:response checkDecompressionStatusOK
*/
type CheckDecompressionStatusOK struct {

	/*
	  In: Body
	*/
	Payload *CheckDecompressionStatusOKBody `json:"body,omitempty"`
}

// NewCheckDecompressionStatusOK creates CheckDecompressionStatusOK with default headers values
func NewCheckDecompressionStatusOK() *CheckDecompressionStatusOK {

	return &CheckDecompressionStatusOK{}
}

// WithPayload adds the payload to the check decompression status o k response
func (o *CheckDecompressionStatusOK) WithPayload(payload *CheckDecompressionStatusOKBody) *CheckDecompressionStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check decompression status o k response
func (o *CheckDecompressionStatusOK) SetPayload(payload *CheckDecompressionStatusOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckDecompressionStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckDecompressionStatusBadRequestCode is the HTTP code returned for type CheckDecompressionStatusBadRequest
const CheckDecompressionStatusBadRequestCode int = 400

/*CheckDecompressionStatusBadRequest This means that the request ID given by the user is incorrect.

swagger:response checkDecompressionStatusBadRequest
*/
type CheckDecompressionStatusBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCheckDecompressionStatusBadRequest creates CheckDecompressionStatusBadRequest with default headers values
func NewCheckDecompressionStatusBadRequest() *CheckDecompressionStatusBadRequest {

	return &CheckDecompressionStatusBadRequest{}
}

// WithPayload adds the payload to the check decompression status bad request response
func (o *CheckDecompressionStatusBadRequest) WithPayload(payload *models.ErrorResponse) *CheckDecompressionStatusBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check decompression status bad request response
func (o *CheckDecompressionStatusBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckDecompressionStatusBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckDecompressionStatusUnauthorizedCode is the HTTP code returned for type CheckDecompressionStatusUnauthorized
const CheckDecompressionStatusUnauthorizedCode int = 401

/*CheckDecompressionStatusUnauthorized This means that the user is not authenticated with keycloak and decompression can't be triggered unless the user first log in with a valid user

swagger:response checkDecompressionStatusUnauthorized
*/
type CheckDecompressionStatusUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCheckDecompressionStatusUnauthorized creates CheckDecompressionStatusUnauthorized with default headers values
func NewCheckDecompressionStatusUnauthorized() *CheckDecompressionStatusUnauthorized {

	return &CheckDecompressionStatusUnauthorized{}
}

// WithPayload adds the payload to the check decompression status unauthorized response
func (o *CheckDecompressionStatusUnauthorized) WithPayload(payload *models.ErrorResponse) *CheckDecompressionStatusUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check decompression status unauthorized response
func (o *CheckDecompressionStatusUnauthorized) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckDecompressionStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckDecompressionStatusNotFoundCode is the HTTP code returned for type CheckDecompressionStatusNotFound
const CheckDecompressionStatusNotFoundCode int = 404

/*CheckDecompressionStatusNotFound This means that the ID doesn't exist and thus a status can't be returned.

swagger:response checkDecompressionStatusNotFound
*/
type CheckDecompressionStatusNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCheckDecompressionStatusNotFound creates CheckDecompressionStatusNotFound with default headers values
func NewCheckDecompressionStatusNotFound() *CheckDecompressionStatusNotFound {

	return &CheckDecompressionStatusNotFound{}
}

// WithPayload adds the payload to the check decompression status not found response
func (o *CheckDecompressionStatusNotFound) WithPayload(payload *models.ErrorResponse) *CheckDecompressionStatusNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check decompression status not found response
func (o *CheckDecompressionStatusNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckDecompressionStatusNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckDecompressionStatusRequestURITooLongCode is the HTTP code returned for type CheckDecompressionStatusRequestURITooLong
const CheckDecompressionStatusRequestURITooLongCode int = 414

/*CheckDecompressionStatusRequestURITooLong This means that the the request ID is longer than the server is willing to interpret.

swagger:response checkDecompressionStatusRequestUriTooLong
*/
type CheckDecompressionStatusRequestURITooLong struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCheckDecompressionStatusRequestURITooLong creates CheckDecompressionStatusRequestURITooLong with default headers values
func NewCheckDecompressionStatusRequestURITooLong() *CheckDecompressionStatusRequestURITooLong {

	return &CheckDecompressionStatusRequestURITooLong{}
}

// WithPayload adds the payload to the check decompression status request Uri too long response
func (o *CheckDecompressionStatusRequestURITooLong) WithPayload(payload *models.ErrorResponse) *CheckDecompressionStatusRequestURITooLong {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check decompression status request Uri too long response
func (o *CheckDecompressionStatusRequestURITooLong) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckDecompressionStatusRequestURITooLong) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(414)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckDecompressionStatusInternalServerErrorCode is the HTTP code returned for type CheckDecompressionStatusInternalServerError
const CheckDecompressionStatusInternalServerErrorCode int = 500

/*CheckDecompressionStatusInternalServerError This means that something has gone wrong on the burst buffer. The user is advised to wait and send the request again.

swagger:response checkDecompressionStatusInternalServerError
*/
type CheckDecompressionStatusInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCheckDecompressionStatusInternalServerError creates CheckDecompressionStatusInternalServerError with default headers values
func NewCheckDecompressionStatusInternalServerError() *CheckDecompressionStatusInternalServerError {

	return &CheckDecompressionStatusInternalServerError{}
}

// WithPayload adds the payload to the check decompression status internal server error response
func (o *CheckDecompressionStatusInternalServerError) WithPayload(payload *models.ErrorResponse) *CheckDecompressionStatusInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check decompression status internal server error response
func (o *CheckDecompressionStatusInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckDecompressionStatusInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}