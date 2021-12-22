// Code generated by go-swagger; DO NOT EDIT.

package data_set_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/lexis-project/lexis-backend-services-api.git/models"
)

// DownloadDatasetOKCode is the HTTP code returned for type DownloadDatasetOK
const DownloadDatasetOKCode int = 200

/*DownloadDatasetOK zip containing dataset

swagger:response downloadDatasetOK
*/
type DownloadDatasetOK struct {

	/*
	  In: Body
	*/
	Payload io.ReadCloser `json:"body,omitempty"`
}

// NewDownloadDatasetOK creates DownloadDatasetOK with default headers values
func NewDownloadDatasetOK() *DownloadDatasetOK {

	return &DownloadDatasetOK{}
}

// WithPayload adds the payload to the download dataset o k response
func (o *DownloadDatasetOK) WithPayload(payload io.ReadCloser) *DownloadDatasetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download dataset o k response
func (o *DownloadDatasetOK) SetPayload(payload io.ReadCloser) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadDatasetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// DownloadDatasetBadRequestCode is the HTTP code returned for type DownloadDatasetBadRequest
const DownloadDatasetBadRequestCode int = 400

/*DownloadDatasetBadRequest Malformed Request

swagger:response downloadDatasetBadRequest
*/
type DownloadDatasetBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewDownloadDatasetBadRequest creates DownloadDatasetBadRequest with default headers values
func NewDownloadDatasetBadRequest() *DownloadDatasetBadRequest {

	return &DownloadDatasetBadRequest{}
}

// WithPayload adds the payload to the download dataset bad request response
func (o *DownloadDatasetBadRequest) WithPayload(payload *models.ErrorResponse) *DownloadDatasetBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download dataset bad request response
func (o *DownloadDatasetBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadDatasetBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DownloadDatasetUnauthorizedCode is the HTTP code returned for type DownloadDatasetUnauthorized
const DownloadDatasetUnauthorizedCode int = 401

/*DownloadDatasetUnauthorized Unauthorized

swagger:response downloadDatasetUnauthorized
*/
type DownloadDatasetUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewDownloadDatasetUnauthorized creates DownloadDatasetUnauthorized with default headers values
func NewDownloadDatasetUnauthorized() *DownloadDatasetUnauthorized {

	return &DownloadDatasetUnauthorized{}
}

// WithPayload adds the payload to the download dataset unauthorized response
func (o *DownloadDatasetUnauthorized) WithPayload(payload *models.ErrorResponse) *DownloadDatasetUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download dataset unauthorized response
func (o *DownloadDatasetUnauthorized) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadDatasetUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DownloadDatasetForbiddenCode is the HTTP code returned for type DownloadDatasetForbidden
const DownloadDatasetForbiddenCode int = 403

/*DownloadDatasetForbidden Forbidden

swagger:response downloadDatasetForbidden
*/
type DownloadDatasetForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewDownloadDatasetForbidden creates DownloadDatasetForbidden with default headers values
func NewDownloadDatasetForbidden() *DownloadDatasetForbidden {

	return &DownloadDatasetForbidden{}
}

// WithPayload adds the payload to the download dataset forbidden response
func (o *DownloadDatasetForbidden) WithPayload(payload *models.ErrorResponse) *DownloadDatasetForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download dataset forbidden response
func (o *DownloadDatasetForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadDatasetForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DownloadDatasetInternalServerErrorCode is the HTTP code returned for type DownloadDatasetInternalServerError
const DownloadDatasetInternalServerErrorCode int = 500

/*DownloadDatasetInternalServerError Internal error processing request

swagger:response downloadDatasetInternalServerError
*/
type DownloadDatasetInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewDownloadDatasetInternalServerError creates DownloadDatasetInternalServerError with default headers values
func NewDownloadDatasetInternalServerError() *DownloadDatasetInternalServerError {

	return &DownloadDatasetInternalServerError{}
}

// WithPayload adds the payload to the download dataset internal server error response
func (o *DownloadDatasetInternalServerError) WithPayload(payload *models.ErrorResponse) *DownloadDatasetInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download dataset internal server error response
func (o *DownloadDatasetInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadDatasetInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DownloadDatasetBadGatewayCode is the HTTP code returned for type DownloadDatasetBadGateway
const DownloadDatasetBadGatewayCode int = 502

/*DownloadDatasetBadGateway Bad Gateway

swagger:response downloadDatasetBadGateway
*/
type DownloadDatasetBadGateway struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewDownloadDatasetBadGateway creates DownloadDatasetBadGateway with default headers values
func NewDownloadDatasetBadGateway() *DownloadDatasetBadGateway {

	return &DownloadDatasetBadGateway{}
}

// WithPayload adds the payload to the download dataset bad gateway response
func (o *DownloadDatasetBadGateway) WithPayload(payload *models.ErrorResponse) *DownloadDatasetBadGateway {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download dataset bad gateway response
func (o *DownloadDatasetBadGateway) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadDatasetBadGateway) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(502)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DownloadDatasetServiceUnavailableCode is the HTTP code returned for type DownloadDatasetServiceUnavailable
const DownloadDatasetServiceUnavailableCode int = 503

/*DownloadDatasetServiceUnavailable unexpected error

swagger:response downloadDatasetServiceUnavailable
*/
type DownloadDatasetServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewDownloadDatasetServiceUnavailable creates DownloadDatasetServiceUnavailable with default headers values
func NewDownloadDatasetServiceUnavailable() *DownloadDatasetServiceUnavailable {

	return &DownloadDatasetServiceUnavailable{}
}

// WithPayload adds the payload to the download dataset service unavailable response
func (o *DownloadDatasetServiceUnavailable) WithPayload(payload *models.ErrorResponse) *DownloadDatasetServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download dataset service unavailable response
func (o *DownloadDatasetServiceUnavailable) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadDatasetServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
