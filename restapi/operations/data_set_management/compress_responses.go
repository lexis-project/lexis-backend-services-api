// Code generated by go-swagger; DO NOT EDIT.

package data_set_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/lexis-project/lexis-backend-services-api.git/models"
)

// CompressCreatedCode is the HTTP code returned for type CompressCreated
const CompressCreatedCode int = 201

/*CompressCreated The response code means that the compression has been initiated. Status of the operation can be checked by querying the status.

swagger:response compressCreated
*/
type CompressCreated struct {

	/*
	  In: Body
	*/
	Payload *models.SteeringRequestID `json:"body,omitempty"`
}

// NewCompressCreated creates CompressCreated with default headers values
func NewCompressCreated() *CompressCreated {

	return &CompressCreated{}
}

// WithPayload adds the payload to the compress created response
func (o *CompressCreated) WithPayload(payload *models.SteeringRequestID) *CompressCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the compress created response
func (o *CompressCreated) SetPayload(payload *models.SteeringRequestID) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CompressCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CompressBadRequestCode is the HTTP code returned for type CompressBadRequest
const CompressBadRequestCode int = 400

/*CompressBadRequest This means that there's something wrong in the input parameters and the server couldn't understand the request.

swagger:response compressBadRequest
*/
type CompressBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCompressBadRequest creates CompressBadRequest with default headers values
func NewCompressBadRequest() *CompressBadRequest {

	return &CompressBadRequest{}
}

// WithPayload adds the payload to the compress bad request response
func (o *CompressBadRequest) WithPayload(payload *models.ErrorResponse) *CompressBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the compress bad request response
func (o *CompressBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CompressBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CompressUnauthorizedCode is the HTTP code returned for type CompressUnauthorized
const CompressUnauthorizedCode int = 401

/*CompressUnauthorized This means that the user is not authenticated with keycloak and compression can't be triggered unless the user first log in with a valid user

swagger:response compressUnauthorized
*/
type CompressUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCompressUnauthorized creates CompressUnauthorized with default headers values
func NewCompressUnauthorized() *CompressUnauthorized {

	return &CompressUnauthorized{}
}

// WithPayload adds the payload to the compress unauthorized response
func (o *CompressUnauthorized) WithPayload(payload *models.ErrorResponse) *CompressUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the compress unauthorized response
func (o *CompressUnauthorized) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CompressUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CompressForbiddenCode is the HTTP code returned for type CompressForbidden
const CompressForbiddenCode int = 403

/*CompressForbidden This means that the resource the user is trying to compress is not readable. User doesn't have the correct rights to read the source file.

swagger:response compressForbidden
*/
type CompressForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCompressForbidden creates CompressForbidden with default headers values
func NewCompressForbidden() *CompressForbidden {

	return &CompressForbidden{}
}

// WithPayload adds the payload to the compress forbidden response
func (o *CompressForbidden) WithPayload(payload *models.ErrorResponse) *CompressForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the compress forbidden response
func (o *CompressForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CompressForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CompressNotFoundCode is the HTTP code returned for type CompressNotFound
const CompressNotFoundCode int = 404

/*CompressNotFound This means that the source path on the system doesn't exist.

swagger:response compressNotFound
*/
type CompressNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCompressNotFound creates CompressNotFound with default headers values
func NewCompressNotFound() *CompressNotFound {

	return &CompressNotFound{}
}

// WithPayload adds the payload to the compress not found response
func (o *CompressNotFound) WithPayload(payload *models.ErrorResponse) *CompressNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the compress not found response
func (o *CompressNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CompressNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CompressRequestURITooLongCode is the HTTP code returned for type CompressRequestURITooLong
const CompressRequestURITooLongCode int = 414

/*CompressRequestURITooLong This means that the either the source path is longer than the server is willing to interpret.

swagger:response compressRequestUriTooLong
*/
type CompressRequestURITooLong struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCompressRequestURITooLong creates CompressRequestURITooLong with default headers values
func NewCompressRequestURITooLong() *CompressRequestURITooLong {

	return &CompressRequestURITooLong{}
}

// WithPayload adds the payload to the compress request Uri too long response
func (o *CompressRequestURITooLong) WithPayload(payload *models.ErrorResponse) *CompressRequestURITooLong {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the compress request Uri too long response
func (o *CompressRequestURITooLong) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CompressRequestURITooLong) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(414)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CompressTooManyRequestsCode is the HTTP code returned for type CompressTooManyRequests
const CompressTooManyRequestsCode int = 429

/*CompressTooManyRequests This means that the user has sent too many requests in a given amount of time. The user is advised to wait and send the request again.

swagger:response compressTooManyRequests
*/
type CompressTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCompressTooManyRequests creates CompressTooManyRequests with default headers values
func NewCompressTooManyRequests() *CompressTooManyRequests {

	return &CompressTooManyRequests{}
}

// WithPayload adds the payload to the compress too many requests response
func (o *CompressTooManyRequests) WithPayload(payload *models.ErrorResponse) *CompressTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the compress too many requests response
func (o *CompressTooManyRequests) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CompressTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CompressInternalServerErrorCode is the HTTP code returned for type CompressInternalServerError
const CompressInternalServerErrorCode int = 500

/*CompressInternalServerError This means that something has gone wrong on the burst buffer. The user is advised to wait and send the request again.

swagger:response compressInternalServerError
*/
type CompressInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCompressInternalServerError creates CompressInternalServerError with default headers values
func NewCompressInternalServerError() *CompressInternalServerError {

	return &CompressInternalServerError{}
}

// WithPayload adds the payload to the compress internal server error response
func (o *CompressInternalServerError) WithPayload(payload *models.ErrorResponse) *CompressInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the compress internal server error response
func (o *CompressInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CompressInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
