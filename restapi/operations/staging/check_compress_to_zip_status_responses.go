// Code generated by go-swagger; DO NOT EDIT.

package staging

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/lexis-project/lexis-backend-services-api.git/models"
)

// CheckCompressToZipStatusOKCode is the HTTP code returned for type CheckCompressToZipStatusOK
const CheckCompressToZipStatusOKCode int = 200

/*CheckCompressToZipStatusOK This means that the status has been returned to the user in the response body.

swagger:response checkCompressToZipStatusOK
*/
type CheckCompressToZipStatusOK struct {

	/*
	  In: Body
	*/
	Payload *CheckCompressToZipStatusOKBody `json:"body,omitempty"`
}

// NewCheckCompressToZipStatusOK creates CheckCompressToZipStatusOK with default headers values
func NewCheckCompressToZipStatusOK() *CheckCompressToZipStatusOK {

	return &CheckCompressToZipStatusOK{}
}

// WithPayload adds the payload to the check compress to zip status o k response
func (o *CheckCompressToZipStatusOK) WithPayload(payload *CheckCompressToZipStatusOKBody) *CheckCompressToZipStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check compress to zip status o k response
func (o *CheckCompressToZipStatusOK) SetPayload(payload *CheckCompressToZipStatusOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckCompressToZipStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckCompressToZipStatusBadRequestCode is the HTTP code returned for type CheckCompressToZipStatusBadRequest
const CheckCompressToZipStatusBadRequestCode int = 400

/*CheckCompressToZipStatusBadRequest This means that the request ID given by the user is incorrect.

swagger:response checkCompressToZipStatusBadRequest
*/
type CheckCompressToZipStatusBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCheckCompressToZipStatusBadRequest creates CheckCompressToZipStatusBadRequest with default headers values
func NewCheckCompressToZipStatusBadRequest() *CheckCompressToZipStatusBadRequest {

	return &CheckCompressToZipStatusBadRequest{}
}

// WithPayload adds the payload to the check compress to zip status bad request response
func (o *CheckCompressToZipStatusBadRequest) WithPayload(payload *models.ErrorResponse) *CheckCompressToZipStatusBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check compress to zip status bad request response
func (o *CheckCompressToZipStatusBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckCompressToZipStatusBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckCompressToZipStatusUnauthorizedCode is the HTTP code returned for type CheckCompressToZipStatusUnauthorized
const CheckCompressToZipStatusUnauthorizedCode int = 401

/*CheckCompressToZipStatusUnauthorized This means that the user is not authenticated with keycloak and data transfer can't be triggered unless the user first log in with a valid user

swagger:response checkCompressToZipStatusUnauthorized
*/
type CheckCompressToZipStatusUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCheckCompressToZipStatusUnauthorized creates CheckCompressToZipStatusUnauthorized with default headers values
func NewCheckCompressToZipStatusUnauthorized() *CheckCompressToZipStatusUnauthorized {

	return &CheckCompressToZipStatusUnauthorized{}
}

// WithPayload adds the payload to the check compress to zip status unauthorized response
func (o *CheckCompressToZipStatusUnauthorized) WithPayload(payload *models.ErrorResponse) *CheckCompressToZipStatusUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check compress to zip status unauthorized response
func (o *CheckCompressToZipStatusUnauthorized) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckCompressToZipStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckCompressToZipStatusNotFoundCode is the HTTP code returned for type CheckCompressToZipStatusNotFound
const CheckCompressToZipStatusNotFoundCode int = 404

/*CheckCompressToZipStatusNotFound This means that the ID doesn't exist and thus a status can't be returned.

swagger:response checkCompressToZipStatusNotFound
*/
type CheckCompressToZipStatusNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCheckCompressToZipStatusNotFound creates CheckCompressToZipStatusNotFound with default headers values
func NewCheckCompressToZipStatusNotFound() *CheckCompressToZipStatusNotFound {

	return &CheckCompressToZipStatusNotFound{}
}

// WithPayload adds the payload to the check compress to zip status not found response
func (o *CheckCompressToZipStatusNotFound) WithPayload(payload *models.ErrorResponse) *CheckCompressToZipStatusNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check compress to zip status not found response
func (o *CheckCompressToZipStatusNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckCompressToZipStatusNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckCompressToZipStatusRequestURITooLongCode is the HTTP code returned for type CheckCompressToZipStatusRequestURITooLong
const CheckCompressToZipStatusRequestURITooLongCode int = 414

/*CheckCompressToZipStatusRequestURITooLong This means that the the request ID is longer than the server is willing to interpret.

swagger:response checkCompressToZipStatusRequestUriTooLong
*/
type CheckCompressToZipStatusRequestURITooLong struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCheckCompressToZipStatusRequestURITooLong creates CheckCompressToZipStatusRequestURITooLong with default headers values
func NewCheckCompressToZipStatusRequestURITooLong() *CheckCompressToZipStatusRequestURITooLong {

	return &CheckCompressToZipStatusRequestURITooLong{}
}

// WithPayload adds the payload to the check compress to zip status request Uri too long response
func (o *CheckCompressToZipStatusRequestURITooLong) WithPayload(payload *models.ErrorResponse) *CheckCompressToZipStatusRequestURITooLong {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check compress to zip status request Uri too long response
func (o *CheckCompressToZipStatusRequestURITooLong) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckCompressToZipStatusRequestURITooLong) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(414)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckCompressToZipStatusInternalServerErrorCode is the HTTP code returned for type CheckCompressToZipStatusInternalServerError
const CheckCompressToZipStatusInternalServerErrorCode int = 500

/*CheckCompressToZipStatusInternalServerError This means that something has gone wrong on the data transfer steering server. The user is advised to wait and send the request again.

swagger:response checkCompressToZipStatusInternalServerError
*/
type CheckCompressToZipStatusInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCheckCompressToZipStatusInternalServerError creates CheckCompressToZipStatusInternalServerError with default headers values
func NewCheckCompressToZipStatusInternalServerError() *CheckCompressToZipStatusInternalServerError {

	return &CheckCompressToZipStatusInternalServerError{}
}

// WithPayload adds the payload to the check compress to zip status internal server error response
func (o *CheckCompressToZipStatusInternalServerError) WithPayload(payload *models.ErrorResponse) *CheckCompressToZipStatusInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check compress to zip status internal server error response
func (o *CheckCompressToZipStatusInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckCompressToZipStatusInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
